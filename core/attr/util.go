package attr

import (
	"context"
	"fmt"
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
		Kind: KindBundle,
		Sub:  slices.Values(b),
	}, nil
}

// Seq is a sequence of attributes.
// These attributes will be rendered adjacent to each other.
// The sequence is iterated once per evaluation; a single-use sequence is
// therefore only safe in a node graph that is evaluated once.
// A nil Seq is empty.
type Seq iter.Seq[Node]

func (s Seq) Eval(ctx context.Context) (Obj, error) {
	return Obj{
		Kind: KindBundle,
		Sub:  iter.Seq[Node](s),
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
// The function must return a non-nil node or an error; return [Noop] for no
// attribute. A nil node, or a nil Fn, fails evaluation with [ErrNilNode].
type Fn func(ctx context.Context) (Node, error)

var _ Node = Fn(nil)

func (fn Fn) Eval(ctx context.Context) (Obj, error) {
	if fn == nil {
		return Obj{}, fmt.Errorf("%w: Fn is nil", ErrNilNode)
	}
	out, err := fn(ctx)
	if err != nil {
		return Obj{}, err
	}
	if out == nil {
		return Obj{}, fmt.Errorf("%w returned by Fn", ErrNilNode)
	}
	return out.Eval(ctx)
}
