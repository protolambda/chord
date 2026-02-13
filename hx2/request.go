package hx2

import (
	"github.com/protolambda/chord/core/attrib"
)

// Include sets the hx-include attribute to include additional elements in request.
func Include(v string) attrib.Node { return attrib.KV("hx-include", v) }

// Params sets the hx-params attribute to filter request parameters.
func Params(v string) attrib.Node { return attrib.KV("hx-params", v) }

// Vals sets the hx-vals attribute to add values to the request.
func Vals(v string) attrib.Node { return attrib.KV("hx-vals", v) }

// Headers sets the hx-headers attribute to add headers to the request.
func Headers(v string) attrib.Node { return attrib.KV("hx-headers", v) }

// Request sets the hx-request attribute to configure request behavior.
func Request(v string) attrib.Node { return attrib.KV("hx-request", v) }

// Sync sets the hx-sync attribute to synchronize AJAX requests.
func Sync(v string) attrib.Node { return attrib.KV("hx-sync", v) }
