package text

import "github.com/protolambda/chord/core"

// Href sets the href attribute for anchor elements.
func Href(v string) core.Node { return core.Attribute("href", v) }

// Target sets the target attribute for anchor elements.
func Target(v string) core.Node { return core.Attribute("target", v) }

// Rel sets the rel attribute for anchor elements.
func Rel(v string) core.Node { return core.Attribute("rel", v) }

// Download sets the download attribute for anchor elements.
func Download(v string) core.Node { return core.Attribute("download", v) }

// Hreflang sets the hreflang attribute for anchor elements.
func Hreflang(v string) core.Node { return core.Attribute("hreflang", v) }

// Ping sets the ping attribute for anchor elements.
func Ping(v string) core.Node { return core.Attribute("ping", v) }
