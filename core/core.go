package core

import (
	"context"
	"fmt"
	"iter"
	"strings"
)

// Obj represents a raw element/attribute (one of the two)
type Obj struct {
	// IsElement as attribute: false, as element: true.
	IsElement bool
	// Key as attribute: attribute-key, as element: element-type.
	// If Key is empty, then Obj is just considered as a bundle of other Obj that all apply to the parent node.
	// Assumed to be safe name/attribute format.
	Key string
	// Val as attribute: attribute value, as element: raw HTML (key/children/void are ignored).
	// Assumed to be html-escaped already.
	Val string
	// Void as attribute: bool-style, as element: no children.
	Void bool
	// SubNodes as attribute: bundled attributes, as element: the child-elements.
	SubNodes iter.Seq[Node]
}

// Eval for Obj just returns the Obj itself, as a valid but "static" node.
func (obj Obj) Eval(ctx context.Context) (Obj, error) {
	return obj, nil
}

// Node is a node in the rendering graph.
// This can be an element-attribute or element-child.
// It's basically an option that applies to the parent node.
// Nodes can be nested and bundled further, see Bundle.
type Node interface {
	Eval(ctx context.Context) (Obj, error)
}

func Render(ctx context.Context, v Node, out *strings.Builder, opts ...Option) error {
	cfg := &renderConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	root, err := v.Eval(ctx)
	if err != nil {
		return fmt.Errorf("failed to load root element: %w", err)
	}
	tmp := make(map[string]struct{})

	// We open the root, since we can represent multiple adjacent HTML elements
	// (no wrapping parent element)
	for o, subErr := range squashOpenNodes(ctx, root) {
		if subErr != nil {
			return fmt.Errorf("failed to open: %w", subErr)
		}
		err := renderObj(ctx, o, out, tmp, cfg, 0)
		if err != nil {
			return fmt.Errorf("failed to render: %w", err)
		}
	}
	return nil
}

func renderObj(ctx context.Context, root Obj, out *strings.Builder, tmpSeenAttribs map[string]struct{}, cfg *renderConfig, depth int) error {
	if !root.IsElement {
		return fmt.Errorf("expected node to be an element")
	}

	if root.Key == "" {
		return fmt.Errorf("expected element to have a key")
	}

	indent := ""
	if cfg.Indent {
		indent = strings.Repeat("  ", depth)
	}

	if root.Val != "" {
		out.WriteString(indent)
		out.WriteString(root.Val)
		if cfg.Indent {
			out.WriteString("\n")
		}
		return nil
	}

	out.WriteString(indent)
	out.WriteString("<")
	out.WriteString(root.Key)

	var classes string
	var styles string
	var attributes, children []Obj

	clear(tmpSeenAttribs)

	for o, subErr := range openSubNodes(ctx, root) {
		if subErr != nil {
			return fmt.Errorf("failed to open: %w", subErr)
		}
		childObj, err := o.Eval(ctx)
		if err != nil {
			return fmt.Errorf("failed to load: %w", err)
		}
		if childObj.IsElement {
			children = append(children, childObj)
		} else {
			if childObj.SubNodes != nil {
				return fmt.Errorf("attribute (%q) cannot have sub-elements", childObj.Key)
			}
			_, seen := tmpSeenAttribs[childObj.Key]
			if !seen {
				tmpSeenAttribs[childObj.Key] = struct{}{}
				attributes = append(attributes, childObj)
			}
			switch childObj.Key {
			case "class":
				if seen {
					classes += " "
				}
				classes += childObj.Val
			case "style":
				if seen {
					styles += ";"
				}
				styles += childObj.Val
			default:
				if seen {
					return fmt.Errorf("duplicate attribute %q", childObj.Key)
				}
			}
		}
	}
	clear(tmpSeenAttribs)

	for _, attr := range attributes {
		out.WriteString(" ")
		out.WriteString(attr.Key)
		if !attr.Void {
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

	if root.Void {
		if len(children) > 0 {
			return fmt.Errorf("unexpected sub-elements in void element: %v", children)
		}
		out.WriteString("/>")
		if cfg.Indent {
			out.WriteString("\n")
		}
	} else {
		out.WriteString(">")
		if cfg.Indent && len(children) > 0 {
			out.WriteString("\n")
		}
		for i, child := range children {
			if err := renderObj(ctx, child, out, tmpSeenAttribs, cfg, depth+1); err != nil {
				return fmt.Errorf("failed to render sub-element %d (type %q): %w", i, child.Key, err)
			}
		}
		if cfg.Indent && len(children) > 0 {
			out.WriteString(indent)
		}
		out.WriteString("</")
		out.WriteString(root.Key)
		out.WriteString(">")
		if cfg.Indent {
			out.WriteString("\n")
		}
	}
	return nil
}

// openSubNodes opens all sub-nodes
func openSubNodes(ctx context.Context, from Obj) iter.Seq2[Obj, error] {
	return func(yield func(Obj, error) bool) {
		if from.Key == "" {
			panic("cannot open sub nodes of a node without a key")
		}
		if from.SubNodes == nil {
			return
		}
		for subNode := range from.SubNodes {
			out, err := subNode.Eval(ctx)
			if err != nil {
				if !yield(Obj{}, err) {
					return
				}
				return
			}
			for got, err := range squashOpenNodes(ctx, out) {
				if !yield(got, err) {
					return
				}
			}
		}
	}
}

// squashOpenNodes squashes the object structure by yielding only objects with a non-empty Key.
func squashOpenNodes(ctx context.Context, from Obj) iter.Seq2[Obj, error] {
	return func(yield func(Obj, error) bool) {
		if from.Key != "" {
			// If Key is not empty, then this is a legitimate sub-node to handle.
			if !yield(from, nil) {
				return
			}
		} else {
			if from.SubNodes == nil {
				return
			}
			// If Key is empty, then squash,
			// by evaluating the sub-nodes, and yielding the results of opening those.
			for subNode := range from.SubNodes {
				out, err := subNode.Eval(ctx)
				if err != nil {
					if !yield(Obj{}, err) {
						return
					}
					return
				}
				for got, err := range squashOpenNodes(ctx, out) {
					if !yield(got, err) {
						return
					}
				}
			}
		}
	}
}
