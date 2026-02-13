package elem

import "context"

// Scope is returned by non-void element constructors.
// Call it with child elements to produce an elem.Node.
// It also implements elem.Node directly (evaluates with no children).
type Scope func(children ...Node) Node

// Eval implements Elem for Scope: an unopened Scope itself counts as an element without content.
func (s Scope) Eval(ctx context.Context) (Obj, error) {
	return s().Eval(ctx)
}
