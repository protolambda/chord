package area_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/attr"
	"github.com/protolambda/chord/html/embed/area"
	"github.com/protolambda/chord/html/embed/img"
)

func ExampleMap() {
	core.Dump(area.Map(attr.ID("sitemap"),
		area.Area(core.Attribute("shape", "rect"), core.Attribute("coords", "0,0,100,100"), core.Attribute("href", "/home")),
	))
	// Output: <map id="sitemap"><area shape="rect" coords="0,0,100,100" href="/home"/></map>
}

func ExampleArea() {
	core.Dump(area.Area(core.Attribute("shape", "circle"), core.Attribute("coords", "50,50,25"), core.Attribute("href", "/link")))
	// Output: <area shape="circle" coords="50,50,25" href="/link"/>
}

func Example_imagemap() {
	core.Dump(core.Bundle(
		img.Img(img.Src("map.png"), img.Usemap("#nav")),
		area.Map(core.Attribute("name", "nav"),
			area.Area(core.Attribute("shape", "rect"), core.Attribute("coords", "0,0,50,50"), core.Attribute("href", "/a")),
		),
	))
	// Output: <img src="map.png" usemap="#nav"/><map name="nav"><area shape="rect" coords="0,0,50,50" href="/a"/></map>
}
