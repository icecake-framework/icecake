package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ick/ickui"
	"github.com/icecake-framework/icecake/pkg/ickcore"
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
	u0 := ickui.Message(ickcore.ToHTML("This is an informative message.")).SetColor(ick.COLOR_INFO)
	boxusage.InsertSnippet(dom.INSERT_LAST_CHILD, u0)

	ResetBoxUsage()

	// let's go
	fmt.Println("Icecake initialized. Listening browser events")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-running">running</span>`)
	}
	<-c
}

func OnDeleteU1(del *ickui.ICKDelete) {
	_btnreset.SetDisabled(false)
}

func ResetBoxUsage() {
	u1 := ickui.Message(ickcore.ToHTML("Click on the delete button to close this warning.")).SetColor(ick.COLOR_WARNING)
	u1.SetHeader(*ickcore.ToHTML("Warning")).SetDeletable("msgu2")
	u1.BtnDelete.OnDelete = OnDeleteU1
	dom.Id("boxusage").InsertSnippet(dom.INSERT_LAST_CHILD, u1)
	_btnreset.SetDisabled(true)
}
