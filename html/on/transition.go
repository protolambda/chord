package on

import (
	"github.com/protolambda/chord/core/attr"
)

// TransitionStart sets the ontransitionstart event handler.
func TransitionStart(v string) attr.Node { return attr.KV("ontransitionstart", v) }

// TransitionEnd sets the ontransitionend event handler.
func TransitionEnd(v string) attr.Node { return attr.KV("ontransitionend", v) }

// TransitionRun sets the ontransitionrun event handler.
func TransitionRun(v string) attr.Node { return attr.KV("ontransitionrun", v) }

// TransitionCancel sets the ontransitioncancel event handler.
func TransitionCancel(v string) attr.Node { return attr.KV("ontransitioncancel", v) }
