package img_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/embed/img"
)

func ExampleImg() {
	core.Dump(img.Img(img.Src("photo.jpg"), img.Alt("A photo")))
	// Output: <img src="photo.jpg" alt="A photo"/>
}

func ExampleImg_responsive() {
	core.Dump(img.Img(img.Src("small.jpg"), img.Srcset("large.jpg 2x"), img.Alt("Responsive")))
	// Output: <img src="small.jpg" srcset="large.jpg 2x" alt="Responsive"/>
}

func ExampleImg_lazy() {
	core.Dump(img.Img(img.Src("big.jpg"), img.Loading("lazy"), img.Alt("Lazy loaded")))
	// Output: <img src="big.jpg" loading="lazy" alt="Lazy loaded"/>
}

func ExamplePicture() {
	core.Dump(img.Picture(
		img.Source(core.Attribute("media", "(min-width: 800px)"), core.Attribute("srcset", "large.jpg")),
		img.Img(img.Src("small.jpg"), img.Alt("Art directed")),
	))
	// Output: <picture><source media="(min-width: 800px)" srcset="large.jpg"/><img src="small.jpg" alt="Art directed"/></picture>
}
