package object_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attr"
	"github.com/protolambda/chord/html/embed/object"
)

func ExampleObject() {
	core.Dump(object.Object(object.Data("movie.swf"), object.Type("application/x-shockwave-flash")))
	// Output: <object data="movie.swf" type="application/x-shockwave-flash"></object>
}

func ExampleEmbed() {
	core.Dump(object.Embed(attr.KV("src", "plugin.swf"), object.Type("application/x-shockwave-flash")))
	// Output: <embed src="plugin.swf" type="application/x-shockwave-flash"/>
}

func ExampleParam() {
	core.Dump(object.Object(object.Data("movie.swf"))(
		object.Param(object.Name("autoplay"), attr.KV("value", "true")),
	))
	// Output: <object data="movie.swf"><param name="autoplay" value="true"/></object>
}
