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
// verbatim. Errors carry the location of the failing node, and wrap the
// original cause. Output failures wrap [ErrRenderOutput]. The context is
// checked while walking the tree, so a canceled context stops rendering.
func Render(ctx context.Context, v elem.Node, out io.StringWriter, opts ...Option) error {
	cfg := &renderConfig{}
	for _, opt := range opts {
		opt(cfg)
	}
	w := &htmlWriter{
		out: out,
		// The root frame never emits a leading newline: it starts as if it
		// already has content.
		frames: []htmlFrame{{indent: cfg.Indent, hasContent: true}},
	}
	if err := walk.Walk(ctx, v, w); err != nil {
		return fmt.Errorf("render: %w", err)
	}
	return nil
}

// htmlWriter is the streaming HTML receiver. Its frame stack tracks the
// indentation state of the currently open elements.
type htmlWriter struct {
	out    io.StringWriter
	frames []htmlFrame
}

type htmlFrame struct {
	tag  string
	void bool
	// indent is whether this frame's content is indented on its own lines.
	indent bool
	// depth is the indentation depth of this frame's content.
	depth int
	// hasContent is set once any content was written inside the frame.
	hasContent bool
}

var _ walk.Receiver = (*htmlWriter)(nil)

func (w *htmlWriter) top() *htmlFrame {
	return &w.frames[len(w.frames)-1]
}

// beginContent writes the newline and indentation that precede a piece of
// content in the current frame.
func (w *htmlWriter) beginContent() error {
	parent := w.top()
	if parent.indent {
		if !parent.hasContent {
			if err := w.write("\n"); err != nil {
				return err
			}
		}
		if err := w.write(strings.Repeat("  ", parent.depth)); err != nil {
			return err
		}
	}
	parent.hasContent = true
	return nil
}

// endContent writes the line ending that follows a piece of content in the
// current frame.
func (w *htmlWriter) endContent() error {
	if w.top().indent {
		return w.write("\n")
	}
	return nil
}

func (w *htmlWriter) Open(el walk.Element) error {
	if err := w.beginContent(); err != nil {
		return err
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
		if err := w.endContent(); err != nil {
			return err
		}
		w.frames = append(w.frames, htmlFrame{tag: el.Tag, void: true})
		return nil
	}
	if err := w.write(">"); err != nil {
		return err
	}
	parent := w.top()
	w.frames = append(w.frames, htmlFrame{
		tag:    el.Tag,
		indent: parent.indent && el.Tag != "pre", // never reformat preformatted content
		depth:  parent.depth + 1,
	})
	return nil
}

func (w *htmlWriter) leaf(content string) error {
	if err := w.beginContent(); err != nil {
		return err
	}
	if err := w.write(content); err != nil {
		return err
	}
	return w.endContent()
}

func (w *htmlWriter) Text(v string) error {
	return w.leaf(html.EscapeString(v))
}

func (w *htmlWriter) Raw(v string) error {
	return w.leaf(v)
}

func (w *htmlWriter) Comment(v string) error {
	return w.leaf("<!-- " + html.EscapeString(v) + " -->")
}

func (w *htmlWriter) Close() error {
	f := *w.top()
	w.frames = w.frames[:len(w.frames)-1]
	if f.void {
		return nil
	}
	if f.indent && f.hasContent {
		if err := w.write(strings.Repeat("  ", w.top().depth)); err != nil {
			return err
		}
	}
	if err := w.write("</", f.tag, ">"); err != nil {
		return err
	}
	return w.endContent()
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
