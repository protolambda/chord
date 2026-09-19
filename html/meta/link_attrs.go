package meta

import (
	"github.com/protolambda/chord/core/attr"
)

// Href sets the href attribute.
func Href(v string) attr.Node { return attr.Name("href").Value(v) }

// Rel sets the rel attribute.
func Rel(v string) attr.Node { return attr.Name("rel").Value(v) }

// Type sets the type attribute.
func Type(v string) attr.Node { return attr.Name("type").Value(v) }

// Media sets the media attribute.
func Media(v string) attr.Node { return attr.Name("media").Value(v) }

// Sizes sets the sizes attribute.
func Sizes(v string) attr.Node { return attr.Name("sizes").Value(v) }

// Crossorigin sets the crossorigin attribute.
func Crossorigin(v string) attr.Node { return attr.Name("crossorigin").Value(v) }

// Integrity sets the integrity attribute.
func Integrity(v string) attr.Node { return attr.Name("integrity").Value(v) }

// Referrerpolicy sets the referrerpolicy attribute.
func Referrerpolicy(v string) attr.Node { return attr.Name("referrerpolicy").Value(v) }
