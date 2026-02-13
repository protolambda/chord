package text

import (
	"github.com/protolambda/chord/core/attrib"
)

// Href sets the href attribute for anchor elements.
func Href(v string) attrib.Node { return attrib.KV("href", v) }

// Target sets the target attribute for anchor elements.
func Target(v string) attrib.Node { return attrib.KV("target", v) }

// Rel sets the rel attribute for anchor elements.
func Rel(v string) attrib.Node { return attrib.KV("rel", v) }

// Download sets the download attribute for anchor elements.
func Download(v string) attrib.Node { return attrib.KV("download", v) }

// Hreflang sets the hreflang attribute for anchor elements.
func Hreflang(v string) attrib.Node { return attrib.KV("hreflang", v) }

// Ping sets the ping attribute for anchor elements.
func Ping(v string) attrib.Node { return attrib.KV("ping", v) }
