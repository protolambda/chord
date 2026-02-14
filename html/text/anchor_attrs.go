package text

import (
	"github.com/protolambda/chord/core/attr"
)

// Href sets the href attribute for anchor elements.
func Href(v string) attr.Node { return attr.KV("href", v) }

// Target sets the target attribute for anchor elements.
func Target(v string) attr.Node { return attr.KV("target", v) }

// Rel sets the rel attribute for anchor elements.
func Rel(v string) attr.Node { return attr.KV("rel", v) }

// Download sets the download attribute for anchor elements.
func Download(v string) attr.Node { return attr.KV("download", v) }

// Hreflang sets the hreflang attribute for anchor elements.
func Hreflang(v string) attr.Node { return attr.KV("hreflang", v) }

// Ping sets the ping attribute for anchor elements.
func Ping(v string) attr.Node { return attr.KV("ping", v) }
