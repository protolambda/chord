// Package hx provides HTMX v1 attributes.
package hx1

import (
	"github.com/protolambda/chord/core/attrib"
)

// Get sets the hx-get attribute for HTTP GET requests.
func Get(v string) attrib.Node { return attrib.KV("hx-get", v) }

// Post sets the hx-post attribute for HTTP POST requests.
func Post(v string) attrib.Node { return attrib.KV("hx-post", v) }

// Put sets the hx-put attribute for HTTP PUT requests.
func Put(v string) attrib.Node { return attrib.KV("hx-put", v) }

// Patch sets the hx-patch attribute for HTTP PATCH requests.
func Patch(v string) attrib.Node { return attrib.KV("hx-patch", v) }

// Delete sets the hx-delete attribute for HTTP DELETE requests.
func Delete(v string) attrib.Node { return attrib.KV("hx-delete", v) }
