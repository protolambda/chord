package meta

import (
	"github.com/protolambda/chord/core/attrib"
)

// Href sets the href attribute.
func Href(v string) attrib.Node { return attrib.KV("href", v) }

// Rel sets the rel attribute.
func Rel(v string) attrib.Node { return attrib.KV("rel", v) }

// Type sets the type attribute.
func Type(v string) attrib.Node { return attrib.KV("type", v) }

// Media sets the media attribute.
func Media(v string) attrib.Node { return attrib.KV("media", v) }

// Sizes sets the sizes attribute.
func Sizes(v string) attrib.Node { return attrib.KV("sizes", v) }

// Crossorigin sets the crossorigin attribute.
func Crossorigin(v string) attrib.Node { return attrib.KV("crossorigin", v) }

// Integrity sets the integrity attribute.
func Integrity(v string) attrib.Node { return attrib.KV("integrity", v) }

// Referrerpolicy sets the referrerpolicy attribute.
func Referrerpolicy(v string) attrib.Node { return attrib.KV("referrerpolicy", v) }
