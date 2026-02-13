package figure_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/html/embed/img"
	"github.com/protolambda/chord/html/group/figure"
	"github.com/protolambda/chord/html/text"
)

func ExampleFigure() {
	core.Dump(figure.Figure()(
		img.Img(img.Src("photo.jpg"), img.Alt("A photo")),
		figure.Figcaption()(text.Text("Photo caption")),
	))
	// Output: <figure><img src="photo.jpg" alt="A photo"/><figcaption>Photo caption</figcaption></figure>
}

func ExampleFigcaption() {
	core.Dump(figure.Figcaption()(text.Text("Figure 1: Example")))
	// Output: <figcaption>Figure 1: Example</figcaption>
}
