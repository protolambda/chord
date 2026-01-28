package core

import (
	"slices"
)

// Noop creates a node that doesn't render any output
func Noop() Node {
	return Obj{
		Key:      "",
		Val:      "",
		SubNodes: nil,
	}
}

// Bundle creates a node that renders all the given nodes
func Bundle(nodes ...Node) Node {
	return Obj{
		Key:      "",
		Val:      "",
		SubNodes: slices.Values(nodes),
	}
}

// Raw creates a node that renders the given string directly in the output
func Raw(v string) Node {
	return Obj{
		IsElement: true,
		Key:       "RAW", // avoid squashing of this element
		Val:       v,
	}
}

// Attribute creates a node that applies the given key-value
// pair as attribute to the currently scoped element.
func Attribute(k, v string) Node {
	return Obj{
		IsElement: false,
		Key:       k,
		Val:       v,
		Void:      false,
		SubNodes:  nil,
	}
}

// BoolAttribute is like an Attribute without explicit value.
// Just the presence of the attribute counts for HTML as truthy.
func BoolAttribute(k string) Node {
	return Obj{
		IsElement: false,
		Key:       k,
		Val:       "",
		Void:      true,
		SubNodes:  nil,
	}
}

// Element creates a node that embeds as element in the scope.
// The Element itself can have child-nodes.
func Element(name string, children ...Node) Node {
	return Obj{
		IsElement: true,
		Key:       name,
		Val:       "",
		Void:      false,
		SubNodes:  slices.Values(children),
	}
}

// VoidElement creates a node that embeds as element in the scope.
// This type of element has no child elements, it's self-closing, e.g. a `<br/>`.
func VoidElement(name string, opts ...Node) Node {
	return Obj{
		IsElement: true,
		Key:       name,
		Val:       "",
		Void:      true,
		SubNodes:  slices.Values(opts),
	}
}
