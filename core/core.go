package core

import (
	"context"
	"errors"
	"fmt"
	"io"
	"iter"
	"strings"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// ErrRenderOutput indicates that the output rejected a rendered string.
var ErrRenderOutput = errors.New("failed to write render output")

// Render renders an element tree to the given string writer.
func Render(ctx context.Context, v elem.Node, out io.StringWriter, opts ...Option) error {
	cfg := &renderConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	obj, err := v.Eval(ctx)
	if err != nil {
		return fmt.Errorf("failed to load root element: %w", err)
	}

	// We allocate this map once, and reuse it during rendering of every element.
	seen := make(map[string]struct{})

	// Squash the root to support rendering multiple adjacent elements.
	for o, squashErr := range squashElemObj(ctx, obj) {
		if squashErr != nil {
			return fmt.Errorf("failed to open: %w", squashErr)
		}
		if err := renderElemObj(ctx, o, out, cfg, 0, seen); err != nil {
			return fmt.Errorf("failed to render: %w", err)
		}
	}
	return nil
}

func renderElemObj(ctx context.Context, obj elem.Obj, out io.StringWriter, cfg *renderConfig, depth int, seen map[string]struct{}) error {
	indent := ""
	if cfg.Indent {
		indent = strings.Repeat("  ", depth)
	}

	// Raw content (including comments).
	if obj.Raw != "" {
		if err := writeStrings(out, indent, obj.Raw); err != nil {
			return err
		}
		if cfg.Indent {
			return writeStrings(out, "\n")
		}
		return nil
	}

	if obj.Tag == "" {
		return fmt.Errorf("expected element to have a tag")
	}

	if err := writeStrings(out, indent, "<", obj.Tag); err != nil {
		return err
	}

	// Flatten and render attributes.
	if err := renderAttribs(ctx, obj.Attribs, out, seen); err != nil {
		return err
	}

	if obj.Void {
		ending := "/>"
		if cfg.Indent {
			ending += "\n"
		}
		return writeStrings(out, ending)
	}

	if err := writeStrings(out, ">"); err != nil {
		return err
	}

	// Disable indentation within <pre> tags.
	cfgInner := cfg
	if obj.Tag == "pre" {
		cpy := *cfg
		cpy.Indent = false
		cfgInner = &cpy
	}

	// Collect and flatten children.
	var children []elem.Obj
	if obj.Children != nil {
		for child := range obj.Children {
			childObj, err := child.Eval(ctx)
			if err != nil {
				return fmt.Errorf("failed to load child: %w", err)
			}
			for o, squashErr := range squashElemObj(ctx, childObj) {
				if squashErr != nil {
					return fmt.Errorf("failed to open child: %w", squashErr)
				}
				children = append(children, o)
			}
		}
	}

	if cfgInner.Indent && len(children) > 0 {
		if err := writeStrings(out, "\n"); err != nil {
			return err
		}
	}
	for i, child := range children {
		if err := renderElemObj(ctx, child, out, cfgInner, depth+1, seen); err != nil {
			return fmt.Errorf("failed to render child %d (tag %q): %w", i, child.Tag, err)
		}
	}
	if cfgInner.Indent && len(children) > 0 {
		if err := writeStrings(out, indent); err != nil {
			return err
		}
	}

	if err := writeStrings(out, "</", obj.Tag, ">"); err != nil {
		return err
	}
	if cfg.Indent {
		return writeStrings(out, "\n")
	}
	return nil
}

// renderAttribs flattens and renders all attributes.
func renderAttribs(ctx context.Context, attribs attr.Seq, out io.StringWriter, seen map[string]struct{}) error {
	if attribs == nil {
		return nil
	}

	// reset the attributes seen map
	clear(seen)

	var classes string
	var styles string
	var collected []attr.Obj

	for a := range attribs {
		if err := flattenAttrib(ctx, a, &classes, &styles, &collected, seen); err != nil {
			return err
		}
	}

	for _, attr := range collected {
		if err := writeStrings(out, " ", attr.Key); err != nil {
			return err
		}
		if !attr.Bool {
			value := attr.Val
			switch attr.Key {
			case "class":
				value = classes
			case "style":
				value = styles
			}
			if err := writeStrings(out, "=\"", value, "\""); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeStrings(out io.StringWriter, values ...string) error {
	for _, value := range values {
		n, err := out.WriteString(value)
		if err != nil {
			return fmt.Errorf("%w: %w", ErrRenderOutput, err)
		}
		if n != len(value) {
			return fmt.Errorf("%w: %w", ErrRenderOutput, io.ErrShortWrite)
		}
	}
	return nil
}

// flattenAttrib recursively flattens attribute bundles and collects attributes.
func flattenAttrib(ctx context.Context, a attr.Node, classes,
	styles *string, collected *[]attr.Obj, seen map[string]struct{}) error {
	obj, err := a.Eval(ctx)
	if err != nil {
		return fmt.Errorf("failed to eval attribute: %w", err)
	}

	// Noop.
	if obj.Key == "" && obj.Sub == nil {
		return nil
	}

	// Bundle: flatten sub-attributes.
	if obj.Key == "" && obj.Sub != nil {
		for sub := range obj.Sub {
			if err := flattenAttrib(ctx, sub, classes, styles, collected, seen); err != nil {
				return err
			}
		}
		return nil
	}

	// Regular attribute.
	_, alreadySeen := seen[obj.Key]
	if !alreadySeen {
		seen[obj.Key] = struct{}{}
		*collected = append(*collected, obj)
	}

	switch obj.Key {
	case "class":
		if alreadySeen {
			*classes += " "
		}
		*classes += obj.Val
	case "style":
		if alreadySeen {
			*styles += ";"
		}
		*styles += obj.Val
	default:
		if alreadySeen {
			return fmt.Errorf("duplicate attribute %q", obj.Key)
		}
	}

	return nil
}

// squashElemObj yields non-bundle elem.Objs by recursively flattening bundles.
func squashElemObj(ctx context.Context, obj elem.Obj) iter.Seq2[elem.Obj, error] {
	return func(yield func(elem.Obj, error) bool) {
		// Non-empty tag or raw content: this is a real element.
		if obj.Tag != "" || obj.Raw != "" {
			yield(obj, nil)
			return
		}
		// Noop: no tag, no raw, no children.
		if obj.Children == nil {
			return
		}
		// Bundle: flatten children.
		for child := range obj.Children {
			childObj, err := child.Eval(ctx)
			if err != nil {
				if !yield(elem.Obj{}, err) {
					return
				}
				return
			}
			for o, squashErr := range squashElemObj(ctx, childObj) {
				if !yield(o, squashErr) {
					return
				}
			}
		}
	}
}
