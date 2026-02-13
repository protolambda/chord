package area_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/embed/area"
	"github.com/protolambda/chord/html/embed/img"
)

func ExampleMap() {
	core.Dump(area.Map(attr.ID("sitemap"))(
		area.Area(attrib.KV("shape", "rect"), attrib.KV("coords", "0,0,100,100"), attrib.KV("href", "/home")),
	))
	// Output: <map id="sitemap"><area shape="rect" coords="0,0,100,100" href="/home"/></map>
}

func ExampleArea() {
	core.Dump(area.Area(attrib.KV("shape", "circle"), attrib.KV("coords", "50,50,25"), attrib.KV("href", "/link")))
	// Output: <area shape="circle" coords="50,50,25" href="/link"/>
}

func Example_imagemap() {
	core.Dump(elem.Bundle{
		img.Img(img.Src("map.png"), img.Usemap("#nav")),
		area.Map(attrib.KV("name", "nav"))(
			area.Area(attrib.KV("shape", "rect"), attrib.KV("coords", "0,0,50,50"), attrib.KV("href", "/a")),
		),
	})
	// Output: <img src="map.png" usemap="#nav"/><map name="nav"><area shape="rect" coords="0,0,50,50" href="/a"/></map>
}
