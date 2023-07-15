package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/dom"
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

	_btnreset.OnClick = func() { RenderBoxUsage() }

	// wrap the back rendered navbar and reset btn
	dom.MountId(&ickui.ICKNavbar{}, "topbar")
	dom.MountId(_btnreset, "btnreset")

	RenderBoxUsage()

	// let's go
	fmt.Println("Icecake initialized. Listening browser events")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-running">running</span>`)
	}
	<-c
}

func RenderBoxUsage() {

	// clear boxusage content
	boxusage := dom.Id("boxusage")
	boxusage.InsertText(dom.INSERT_BODY, "")

	u1 := ickui.InputField("in-u1", "", "e.g. bob").
		SetLabel("Enter your pseudo").
		SetHelp("Only alphabetic letters")
	dom.Id("boxusage").InsertSnippet(dom.INSERT_LAST_CHILD, u1)

	u2 := ickui.InputField("in-u2", "", "password").
		SetHidden(true).
		SetCanToggleVisibility(true).
		SetLabel("Enter your password").
		SetHelp("Must be 12 characters long or more and must contains lowercase, uppercase, digit and symbol")

	dom.Id("boxusage").InsertSnippet(dom.INSERT_LAST_CHILD, u2)

	_btnreset.SetDisabled(true)
}
