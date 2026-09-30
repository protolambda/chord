package core

import (
	"context"
	"errors"
	"fmt"
	"html"
	"io"
	"strings"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/core/internal/walk"
)

// ErrRenderOutput indicates that the output rejected a rendered string.
var ErrRenderOutput = errors.New("failed to write render output")

// Render evaluates the element tree once and writes it as HTML to out.
//
// Text and logical attribute values are escaped here; raw content is written
// verbatim, and so is text in HTML script and style elements, which browsers
// do not unescape (see [elem.Text]). Errors carry the location of the
// failing node, and wrap the original cause. Output failures wrap
// [ErrRenderOutput]. The context is checked while walking the tree, so a
// canceled context stops rendering.
func Render(ctx context.Context, v elem.Node, out io.StringWriter, opts ...Option) error {
	cfg := &renderConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	w := &htmlWriter{
		out:    out,
		frames: []htmlFrame{{indent: cfg.Indent}},
	}
	if err := walk.Walk(ctx, v, w); err != nil {
		return fmt.Errorf("render: %w", err)
	}
	if err := w.finish(); err != nil {
		return fmt.Errorf("render: %w", err)
	}
	return nil
}

// htmlWriter is the streaming HTML receiver. Its frame stack tracks the
// indentation state of the currently open elements; frames[0] is the root.
type htmlWriter struct {
	out    io.StringWriter
	frames []htmlFrame
}

// lastChild classifies the most recent child written into a frame.
type lastChild uint8

const (
	lastNone   lastChild = iota // no child yet
	lastBlock                   // a block-level element
	lastInline                  // text, raw content, a comment, or another element
)

type htmlFrame struct {
	tag  string
	void bool
	// literal is set for script and style elements, whose text is written
	// verbatim.
	literal bool
	// newlineFirst is set for elements whose first newline the HTML parser
	// drops.
	newlineFirst bool
	// wrote is set once content was written after the start tag.
	wrote bool
	// indent is whether whitespace may be added between this frame's
	// children: indentation is enabled and the content is not preformatted.
	indent bool
	// block is whether the element is block-level, as a child of its parent.
	block bool
	// depth is the indentation depth of this frame's children.
	depth int
	// last is the kind of the most recent child.
	last lastChild
}

var _ walk.Receiver = (*htmlWriter)(nil)

func (w *htmlWriter) top() *htmlFrame {
	return &w.frames[len(w.frames)-1]
}

// newline writes a line break and the indentation of depth.
func (w *htmlWriter) newline(depth int) error {
	return w.write("\n", strings.Repeat("  ", depth))
}

func (w *htmlWriter) Open(el walk.Element) error {
	parent := w.top()
	parent.wrote = true
	root := len(w.frames) == 1
	block := parent.indent && blockLevel(el.Tag, parent.tag)
	// Indentation only goes between block-level elements, and between a
	// block-level element and the start of its parent: browsers render no
	// whitespace there. The output does not start with a line break.
	if block && (parent.last == lastBlock || parent.last == lastNone && !root) {
		if err := w.newline(parent.depth); err != nil {
			return err
		}
	}
	if err := w.write("<", el.Tag); err != nil {
		return err
	}
	for _, a := range el.Attrs {
		if err := w.write(" ", a.Key); err != nil {
			return err
		}
		switch a.Kind {
		case attr.KindValue:
			if err := w.write(`="`, html.EscapeString(a.Val), `"`); err != nil {
				return err
			}
		case attr.KindRawValue:
			if err := w.write(`="`, a.Val, `"`); err != nil {
				return err
			}
		}
	}
	if el.Void {
		if err := w.write("/>"); err != nil {
			return err
		}
		w.frames = append(w.frames, htmlFrame{tag: el.Tag, void: true, block: block})
		return nil
	}
	if err := w.write(">"); err != nil {
		return err
	}
	w.frames = append(w.frames, htmlFrame{
		tag:          el.Tag,
		literal:      el.Literal,
		newlineFirst: dropsLeadingNewline(el.Tag),
		indent:       parent.indent && !preformatted(el.Tag),
		block:        block,
		depth:        parent.depth + 1,
	})
	return nil
}

// leaf writes inline content: no whitespace is added around it.
func (w *htmlWriter) leaf(content string) error {
	f := w.top()
	f.last = lastInline
	if content != "" {
		f.wrote = true
	}
	return w.write(content)
}

func (w *htmlWriter) Text(v string) error {
	f := w.top()
	if f.newlineFirst && !f.wrote && v != "" && (v[0] == '\n' || v[0] == '\r') {
		// The parser drops one newline right after the start tag, also
		// when written as CR LF or CR, which it reads as LF.
		if err := w.write("\n"); err != nil {
			return err
		}
	}
	if f.literal {
		return w.leaf(v)
	}
	return w.leaf(html.EscapeString(v))
}

func (w *htmlWriter) Raw(v string) error {
	return w.leaf(v)
}

func (w *htmlWriter) Comment(v string) error {
	return w.leaf("<!-- " + html.EscapeString(v) + " -->")
}

// Doctype writes the declaration. For indentation it counts as a block-level
// element: the parser ignores whitespace around it.
func (w *htmlWriter) Doctype(name string) error {
	root := w.top()
	if root.indent && root.last == lastBlock {
		if err := w.newline(0); err != nil {
			return err
		}
	}
	if err := w.write("<!DOCTYPE ", name, ">"); err != nil {
		return err
	}
	root.last = lastBlock
	return nil
}

func (w *htmlWriter) Close() error {
	f := *w.top()
	w.frames = w.frames[:len(w.frames)-1]
	if !f.void {
		if f.indent && f.last == lastBlock {
			if err := w.newline(f.depth - 1); err != nil {
				return err
			}
		}
		if err := w.write("</", f.tag, ">"); err != nil {
			return err
		}
	}
	if f.block {
		w.top().last = lastBlock
	} else {
		w.top().last = lastInline
	}
	return nil
}

// finish ends indented output that ends with a block-level element with a
// line break.
func (w *htmlWriter) finish() error {
	if root := w.top(); root.indent && root.last == lastBlock {
		return w.write("\n")
	}
	return nil
}

func (w *htmlWriter) write(values ...string) error {
	for _, value := range values {
		n, err := w.out.WriteString(value)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrRenderOutput, err)
		}
		if n != len(value) {
			return fmt.Errorf("%w: %w", ErrRenderOutput, io.ErrShortWrite)
		}
	}
	return nil
}
