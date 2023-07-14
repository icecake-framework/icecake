package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/dom"
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

	ii := make([]*ickui.ICKInputField, 9)
	ii[0] = ickui.InputField("in-u0", "", "input")
	ii[1] = ickui.InputField("in-u1", "initial value", "input with initial value")
	ii[2] = ickui.InputField("in-u2", "readonly", "input readonly").SetReadOnly(true)
	ii[3] = ickui.InputField("in-u3", "", "loading input").SetState(ickui.INPUT_LOADING)
	ii[4] = ickui.InputField("in-u4", "", "disabled input").SetState(ickui.INPUT_DISABLED)
	ii[5] = ickui.InputField("in-u5", "", "static readonly input").SetState(ickui.INPUT_STATIC).SetReadOnly(true)
	ii[6] = ickui.InputField("in-u6", "", "email").
		SetIcon(*ick.Icon("bi bi-envelope-at").SetColor(ick.TXTCOLOR_INFO_DARK), false).
		SetIcon(*ick.Icon("bi bi-info-circle").SetColor(ick.TXTCOLOR_INFO_DARK), true)
	ii[7] = ickui.InputField("in-u8", "", "password").SetHidden(true)
	ii[8] = ickui.InputField("in-u7", "", "e.g. bob").
		SetLabel("Enter your pseudo").
		SetHelp("Only alphabetic letters")

	for _, c := range ii {
		dom.Id("boxusage").InsertSnippet(dom.INSERT_LAST_CHILD, c)
	}

	_btnreset.SetDisabled(true)
}
