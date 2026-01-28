package ba

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
)

// TextStart creates a Bootstrap text-start class.
func TextStart() core.Node { return attr.Class("text-start") }

// TextCenter creates a Bootstrap text-center class.
func TextCenter() core.Node { return attr.Class("text-center") }

// TextEnd creates a Bootstrap text-end class.
func TextEnd() core.Node { return attr.Class("text-end") }

// TextWrap creates a Bootstrap text-wrap class.
func TextWrap() core.Node { return attr.Class("text-wrap") }

// TextNowrap creates a Bootstrap text-nowrap class.
func TextNowrap() core.Node { return attr.Class("text-nowrap") }

// TextBreak creates a Bootstrap text-break class.
func TextBreak() core.Node { return attr.Class("text-break") }

// TextLowercase creates a Bootstrap text-lowercase class.
func TextLowercase() core.Node { return attr.Class("text-lowercase") }

// TextUppercase creates a Bootstrap text-uppercase class.
func TextUppercase() core.Node { return attr.Class("text-uppercase") }

// TextCapitalize creates a Bootstrap text-capitalize class.
func TextCapitalize() core.Node { return attr.Class("text-capitalize") }

// TextMuted creates a Bootstrap text-body-secondary class (formerly text-muted).
func TextMuted() core.Node { return attr.Class("text-body-secondary") }

// TextPrimary creates a Bootstrap text-primary class.
func TextPrimary() core.Node { return attr.Class("text-primary") }

// TextSecondary creates a Bootstrap text-secondary class.
func TextSecondary() core.Node { return attr.Class("text-secondary") }

// TextSuccess creates a Bootstrap text-success class.
func TextSuccess() core.Node { return attr.Class("text-success") }

// TextDanger creates a Bootstrap text-danger class.
func TextDanger() core.Node { return attr.Class("text-danger") }

// TextWarning creates a Bootstrap text-warning class.
func TextWarning() core.Node { return attr.Class("text-warning") }

// TextInfo creates a Bootstrap text-info class.
func TextInfo() core.Node { return attr.Class("text-info") }

// TextLight creates a Bootstrap text-light class.
func TextLight() core.Node { return attr.Class("text-light") }

// TextDark creates a Bootstrap text-dark class.
func TextDark() core.Node { return attr.Class("text-dark") }

// TextBody creates a Bootstrap text-body class.
func TextBody() core.Node { return attr.Class("text-body") }

// TextWhite creates a Bootstrap text-white class.
func TextWhite() core.Node { return attr.Class("text-white") }

// FW creates a Bootstrap fw-{weight} class.
func FW(weight string) core.Node { return attr.Class("fw-" + weight) }

// FWBold creates a Bootstrap fw-bold class.
func FWBold() core.Node { return attr.Class("fw-bold") }

// FWNormal creates a Bootstrap fw-normal class.
func FWNormal() core.Node { return attr.Class("fw-normal") }

// FWLight creates a Bootstrap fw-light class.
func FWLight() core.Node { return attr.Class("fw-light") }

// FSItalic creates a Bootstrap fst-italic class.
func FSItalic() core.Node { return attr.Class("fst-italic") }

// FSNormal creates a Bootstrap fst-normal class.
func FSNormal() core.Node { return attr.Class("fst-normal") }

// LH creates a Bootstrap lh-{v} (line-height) class.
func LH(v string) core.Node { return attr.Class("lh-" + v) }

// FontMonospace creates a Bootstrap font-monospace class.
func FontMonospace() core.Node { return attr.Class("font-monospace") }
