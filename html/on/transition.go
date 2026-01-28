package on

import "github.com/protolambda/chord/core"

// TransitionStart sets the ontransitionstart event handler.
func TransitionStart(v string) core.Node { return core.Attribute("ontransitionstart", v) }

// TransitionEnd sets the ontransitionend event handler.
func TransitionEnd(v string) core.Node { return core.Attribute("ontransitionend", v) }

// TransitionRun sets the ontransitionrun event handler.
func TransitionRun(v string) core.Node { return core.Attribute("ontransitionrun", v) }

// TransitionCancel sets the ontransitioncancel event handler.
func TransitionCancel(v string) core.Node { return core.Attribute("ontransitioncancel", v) }
