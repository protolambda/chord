package on

import (
	"github.com/protolambda/chord/core/attr"
)

// TransitionStart sets the ontransitionstart event handler.
func TransitionStart(v string) attr.Node { return attr.Name("ontransitionstart").Value(v) }

// TransitionEnd sets the ontransitionend event handler.
func TransitionEnd(v string) attr.Node { return attr.Name("ontransitionend").Value(v) }

// TransitionRun sets the ontransitionrun event handler.
func TransitionRun(v string) attr.Node { return attr.Name("ontransitionrun").Value(v) }

// TransitionCancel sets the ontransitioncancel event handler.
func TransitionCancel(v string) attr.Node { return attr.Name("ontransitioncancel").Value(v) }
