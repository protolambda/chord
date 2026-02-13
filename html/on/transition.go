package on

import (
	"github.com/protolambda/chord/core/attrib"
)

// TransitionStart sets the ontransitionstart event handler.
func TransitionStart(v string) attrib.Node { return attrib.KV("ontransitionstart", v) }

// TransitionEnd sets the ontransitionend event handler.
func TransitionEnd(v string) attrib.Node { return attrib.KV("ontransitionend", v) }

// TransitionRun sets the ontransitionrun event handler.
func TransitionRun(v string) attrib.Node { return attrib.KV("ontransitionrun", v) }

// TransitionCancel sets the ontransitioncancel event handler.
func TransitionCancel(v string) attrib.Node { return attrib.KV("ontransitioncancel", v) }
