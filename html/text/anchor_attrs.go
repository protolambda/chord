package text

import (
	"github.com/protolambda/chord/core/attr"
)

// Href sets the href attribute for anchor elements.
func Href(v string) attr.Node { return attr.Name("href").Value(v) }

// Target sets the target attribute for anchor elements.
func Target(v string) attr.Node { return attr.Name("target").Value(v) }

// Rel sets the rel attribute for anchor elements.
func Rel(v string) attr.Node { return attr.Name("rel").Value(v) }

// Download sets the download attribute for anchor elements.
func Download(v string) attr.Node { return attr.Name("download").Value(v) }

// Hreflang sets the hreflang attribute for anchor elements.
func Hreflang(v string) attr.Node { return attr.Name("hreflang").Value(v) }

// Ping sets the ping attribute for anchor elements.
func Ping(v string) attr.Node { return attr.Name("ping").Value(v) }
