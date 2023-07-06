package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ick/ickui"
	"github.com/lolorenzo777/verbose"
)

var _btnreset = &ickui.ICKButton{}

// This main package contains web assembly source code.
func main() {
	c := make(chan struct{})
	fmt.Println("Go/WASM loaded. Icecake initializing...")
	verbose.IsOn = true
	verbose.IsDebugging = true

	ests := dom.Id("icecake-status")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-initializing">initializing</span>`)
	}

	_btnreset.OnClick = func() { ResetBoxUsage() }

	// wrap the back rendered navbar and reset btn
	dom.MountId(&ickui.ICKNavbar{}, "topbar")
	dom.MountId(_btnreset, "btnreset")

	// front rendering
	boxusage := dom.Id("boxusage")
	boxusage.InsertText(dom.INSERT_BODY, "")
	u1 := ickui.Message(html.ToHTML("This is an informative message.")).SetColor(ick.COLOR_INFO)
	boxusage.InsertSnippet(dom.INSERT_LAST_CHILD, u1)

	ResetBoxUsage()

	// let's go
	fmt.Println("Icecake initialized. Listening browser events")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-running">running</span>`)
	}
	<-c
}

func OnDeleteU2(del *ickui.ICKDelete) {
	_btnreset.SetDisabled(false)
}

func ResetBoxUsage() {
	u2 := ickui.Message(html.ToHTML("Click on the delete button to close this warning.")).SetColor(ick.COLOR_WARNING)
	u2.SetHeader(*html.ToHTML("Warning")).SetDeletable("msgu2")
	u2.BtnDelete.OnDelete = OnDeleteU2
	dom.Id("boxusage").InsertSnippet(dom.INSERT_LAST_CHILD, u2)
	_btnreset.SetDisabled(true)
}
