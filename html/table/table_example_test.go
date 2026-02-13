package table_test

import (
	"github.com/protolambda/chord/core"
	"github.com/protolambda/chord/core/attrib"
	"github.com/protolambda/chord/html/table"
	"github.com/protolambda/chord/html/text"
)

func ExampleTable() {
	core.Dump(table.Table()(
		table.Thead()(table.TR()(table.TH()(text.Text("Name")))),
		table.Tbody()(table.TR()(table.TD()(text.Text("Alice")))),
	))
	// Output: <table><thead><tr><th>Name</th></tr></thead><tbody><tr><td>Alice</td></tr></tbody></table>
}

func ExampleCaption() {
	core.Dump(table.Table()(table.Caption()(text.Text("Users"))))
	// Output: <table><caption>Users</caption></table>
}

func ExampleColgroup() {
	core.Dump(table.Colgroup()(table.Col(attrib.KV("span", "2"))))
	// Output: <colgroup><col span="2"/></colgroup>
}

func ExampleTfoot() {
	core.Dump(table.Tfoot()(table.TR()(table.TD()(text.Text("Total")))))
	// Output: <tfoot><tr><td>Total</td></tr></tfoot>
}

func ExampleTH_scope() {
	core.Dump(table.TH(table.Scope("col"))(text.Text("Header")))
	// Output: <th scope="col">Header</th>
}

func ExampleTD_colspan() {
	core.Dump(table.TD(table.Colspan("2"))(text.Text("Merged")))
	// Output: <td colspan="2">Merged</td>
}
