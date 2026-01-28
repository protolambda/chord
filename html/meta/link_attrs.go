package meta

import "github.com/protolambda/chord/core"

// Href sets the href attribute.
func Href(v string) core.Node { return core.Attribute("href", v) }

// Rel sets the rel attribute.
func Rel(v string) core.Node { return core.Attribute("rel", v) }

// Type sets the type attribute.
func Type(v string) core.Node { return core.Attribute("type", v) }

// Media sets the media attribute.
func Media(v string) core.Node { return core.Attribute("media", v) }

// Sizes sets the sizes attribute.
func Sizes(v string) core.Node { return core.Attribute("sizes", v) }

// Crossorigin sets the crossorigin attribute.
func Crossorigin(v string) core.Node { return core.Attribute("crossorigin", v) }

// Integrity sets the integrity attribute.
func Integrity(v string) core.Node { return core.Attribute("integrity", v) }

// Referrerpolicy sets the referrerpolicy attribute.
func Referrerpolicy(v string) core.Node { return core.Attribute("referrerpolicy", v) }
