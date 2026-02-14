package meta

import (
	"github.com/protolambda/chord/core/attr"
)

// Href sets the href attribute.
func Href(v string) attr.Node { return attr.KV("href", v) }

// Rel sets the rel attribute.
func Rel(v string) attr.Node { return attr.KV("rel", v) }

// Type sets the type attribute.
func Type(v string) attr.Node { return attr.KV("type", v) }

// Media sets the media attribute.
func Media(v string) attr.Node { return attr.KV("media", v) }

// Sizes sets the sizes attribute.
func Sizes(v string) attr.Node { return attr.KV("sizes", v) }

// Crossorigin sets the crossorigin attribute.
func Crossorigin(v string) attr.Node { return attr.KV("crossorigin", v) }

// Integrity sets the integrity attribute.
func Integrity(v string) attr.Node { return attr.KV("integrity", v) }

// Referrerpolicy sets the referrerpolicy attribute.
func Referrerpolicy(v string) attr.Node { return attr.KV("referrerpolicy", v) }
