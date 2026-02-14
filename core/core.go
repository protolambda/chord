package core

import (
	"context"
	"fmt"
	"iter"
	"strings"

	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/core/elem"
)

// Render renders an element tree to the given string builder.
func Render(ctx context.Context, v elem.Node, out *strings.Builder, opts ...Option) error {
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

func renderElemObj(ctx context.Context, obj elem.Obj, out *strings.Builder, cfg *renderConfig, depth int, seen map[string]struct{}) error {
	indent := ""
	if cfg.Indent {
		indent = strings.Repeat("  ", depth)
	}

	// Raw content (including comments).
	if obj.Raw != "" {
		out.WriteString(indent)
		out.WriteString(obj.Raw)
		if cfg.Indent {
			out.WriteString("\n")
		}
		return nil
	}

	if obj.Tag == "" {
		return fmt.Errorf("expected element to have a tag")
	}

	out.WriteString(indent)
	out.WriteString("<")
	out.WriteString(obj.Tag)

	// Flatten and render attributes.
	if err := renderAttribs(ctx, obj.Attribs, out, seen); err != nil {
		return err
	}

	if obj.Void {
		out.WriteString("/>")
		if cfg.Indent {
			out.WriteString("\n")
		}
		return nil
	}

	out.WriteString(">")

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
		out.WriteString("\n")
	}
	for i, child := range children {
		if err := renderElemObj(ctx, child, out, cfgInner, depth+1, seen); err != nil {
			return fmt.Errorf("failed to render child %d (tag %q): %w", i, child.Tag, err)
		}
	}
	if cfgInner.Indent && len(children) > 0 {
		out.WriteString(indent)
	}

	out.WriteString("</")
	out.WriteString(obj.Tag)
	out.WriteString(">")
	if cfg.Indent {
		out.WriteString("\n")
	}
	return nil
}

// renderAttribs flattens and renders all attributes.
func renderAttribs(ctx context.Context, attribs attr.Seq, out *strings.Builder, seen map[string]struct{}) error {
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
		out.WriteString(" ")
		out.WriteString(attr.Key)
		if !attr.Bool {
			out.WriteString("=\"")
			switch attr.Key {
			case "class":
				out.WriteString(classes)
			case "style":
				out.WriteString(styles)
			default:
				out.WriteString(attr.Val)
			}
			out.WriteString("\"")
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
