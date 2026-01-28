// Package hx provides HTMX v1 attributes.
package hx1

import "github.com/protolambda/chord/core"

// Get sets the hx-get attribute for HTTP GET requests.
func Get(v string) core.Node { return core.Attribute("hx-get", v) }

// Post sets the hx-post attribute for HTTP POST requests.
func Post(v string) core.Node { return core.Attribute("hx-post", v) }

// Put sets the hx-put attribute for HTTP PUT requests.
func Put(v string) core.Node { return core.Attribute("hx-put", v) }

// Patch sets the hx-patch attribute for HTTP PATCH requests.
func Patch(v string) core.Node { return core.Attribute("hx-patch", v) }

// Delete sets the hx-delete attribute for HTTP DELETE requests.
func Delete(v string) core.Node { return core.Attribute("hx-delete", v) }
