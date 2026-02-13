package on_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/elem"
	"github.com/protolambda/chord/html/embed/img"
	"github.com/protolambda/chord/html/form/button"
	"github.com/protolambda/chord/html/form/input"
	"github.com/protolambda/chord/html/group/div"
	"github.com/protolambda/chord/html/on"
	"github.com/protolambda/chord/html/text"
)

func ExampleClick() {
	core.Dump(button.Button(on.Click("handleClick()"))(text.Text("Click me")))
	// Output: <button onclick="handleClick()">Click me</button>
}

func ExampleDblClick() {
	core.Dump(div.Div(on.DblClick("handleDblClick()")))
	// Output: <div ondblclick="handleDblClick()"></div>
}

func ExampleMouseOver() {
	core.Dump(div.Div(on.MouseOver("showTooltip()"), on.MouseOut("hideTooltip()")))
	// Output: <div onmouseover="showTooltip()" onmouseout="hideTooltip()"></div>
}

func ExampleSubmit() {
	core.Dump(elem.New("form", on.Submit("return validate()")))
	// Output: <form onsubmit="return validate()"></form>
}

func ExampleChange() {
	core.Dump(input.Input(input.Type("text"), on.Change("handleChange(this)")))
	// Output: <input type="text" onchange="handleChange(this)"/>
}

func ExampleInput() {
	core.Dump(input.Input(input.Type("text"), on.Input("filterResults(this.value)")))
	// Output: <input type="text" oninput="filterResults(this.value)"/>
}

func ExampleFocus() {
	core.Dump(input.Input(on.Focus("highlight()"), on.Blur("unhighlight()")))
	// Output: <input onfocus="highlight()" onblur="unhighlight()"/>
}

func ExampleKeyDown() {
	core.Dump(input.Input(on.KeyDown("handleKey(event)")))
	// Output: <input onkeydown="handleKey(event)"/>
}

func ExampleKeyUp() {
	core.Dump(input.Input(on.KeyUp("search(this.value)")))
	// Output: <input onkeyup="search(this.value)"/>
}

func ExampleLoad() {
	core.Dump(img.Img(img.Src("pic.jpg"), on.Load("imageLoaded()")))
	// Output: <img src="pic.jpg" onload="imageLoaded()"/>
}

func ExampleError() {
	core.Dump(img.Img(img.Src("pic.jpg"), on.Error("handleError()")))
	// Output: <img src="pic.jpg" onerror="handleError()"/>
}
