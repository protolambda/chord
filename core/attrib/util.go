package attrib

import (
	"context"
	"iter"
	"slices"
)

// Noop produces an attribute that is not rendered.
func Noop() Node {
	return Obj{}
}

// Bundle is a slice of attributes.
// These attributes will be rendered adjacent to each other.
type Bundle []Node

func (b Bundle) Eval(ctx context.Context) (Obj, error) {
	return Obj{
		Sub: slices.Values(b),
	}, nil
}

// Seq is a sequence of attributes.
// These attributes will be rendered adjacent to each other.
type Seq iter.Seq[Node]

func (s Seq) Eval(ctx context.Context) (Obj, error) {
	return Obj{
		Sub: iter.Seq[Node](s),
	}, nil
}

// Cons is a functional convenience for producing a sequence of attributes:
// starting with the head attribute, followed by the tail attributes.
// These attributes will be rendered adjacent to each other.
func Cons(head Node, tail ...Node) Node {
	return Seq(func(yield func(Node) bool) {
		if !yield(head) {
			return
		}
		for _, t := range tail {
			if !yield(t) {
				return
			}
		}
	})
}

// If returns v if x is true, and returns a Noop otherwise.
func If(x bool, v Node) Node {
	if x {
		return v
	} else {
		return Noop()
	}
}

// IfElse returns trueCase if x is true, and returns falseCase otherwise.
func IfElse(x bool, trueCase Node, falseCase Node) Node {
	if x {
		return trueCase
	} else {
		return falseCase
	}
}

// Fn is a type of node that generates the output dynamically.
type Fn func(ctx context.Context) (Node, error)

var _ Node = Fn(nil)

func (fn Fn) Eval(ctx context.Context) (Obj, error) {
	out, err := fn(ctx)
	if err != nil {
		return Obj{}, err
	}
	return out.Eval(ctx)
}
