package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/bulmaui"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/lolorenzo777/verbose"
)

var _btnreset = &bulmaui.ICKButton{}

// This main package contains the web assembly source code for makedocs
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded. Icecake initializing...")
	verbose.IsOn = true
	verbose.IsDebugging = true

	ests := dom.Id("icecake-status")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-initializing">initializing</span>`)
	}

	// wrap the back rendered navbar
	dom.WrapId(&bulmaui.ICKNavbar{}, "topbar").AddListeners()
	dom.WrapId(_btnreset, "btnreset")
	_btnreset.DOM.AddMouseEvent(event.MOUSE_ONCLICK, OnReset)

	// front rendering
	boxusage := dom.Id("boxusage")
	boxusage.InsertText(dom.INSERT_BODY, "")
	u1 := bulmaui.Message(html.ToHTML("This is an informative message.")).SetColor(bulma.COLOR_INFO)
	boxusage.InsertSnippet(dom.INSERT_LAST_CHILD, u1)

	ResetBoxUsage()

	// let's go
	fmt.Println("Icecake initialized. Listening browser events")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-running">running</span>`)
	}
	<-c
}

func OnDeleteU2(del *bulmaui.ICKDelete) {
	_btnreset.SetDisabled(false)
}

func OnReset(evt *event.MouseEvent, e *dom.Element) {
	ResetBoxUsage()
}

func ResetBoxUsage() {
	u2 := bulmaui.Message(html.ToHTML("Click on the delete button to close this warning.")).SetColor(bulma.COLOR_WARNING)
	u2.SetHeader(*html.ToHTML("Warning")).SetDeletable("msgu2")
	u2.BtnDelete.OnDelete = OnDeleteU2
	dom.Id("boxusage").InsertSnippet(dom.INSERT_LAST_CHILD, u2)
	_btnreset.SetDisabled(true)
}
