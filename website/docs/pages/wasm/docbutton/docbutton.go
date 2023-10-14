package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/browser"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ick/ickui"
)

var _btnreset = &ickui.ICKButton{}

// This main package contains web assembly source code.
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded. Icecake initializing...")

	ests := dom.Id("icecake-status")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-initializing">initializing</span>`)
	}

	// front rendering
	clicked := 0
	boxusage := dom.Id("boxusage")
	uF := ickui.Button("Click Counter 0")
	uF.OnClick = func() {
		_btnreset.SetDisabled(false)
		clicked++
		if clicked == 5 {
			uF.SetColor(ick.COLOR_SUCCESS).SetOutlined(true)
			uF.SetIcon(*ick.Icon("bi bi-check-lg"), false)
		}
		uF.SetTitle(fmt.Sprintf("Click Counter %v", clicked))
	}
	boxusage.InsertSnippet(dom.INSERT_BODY, uF)

	_btnreset.OnClick = func() {
		_btnreset.SetDisabled(true)
		clicked = 0
		uF.SetColor(ick.COLOR_NONE).SetOutlined(false)
		uF.SetIcon(ick.ICKIcon{}, false)
		uF.SetTitle(fmt.Sprintf("Click Counter %v", clicked))
	}

	// wrap the back rendered navbar and reset btn
	dom.MountId(&ickui.ICKNavbar{}, "topbar")
	dom.MountId(_btnreset, "btnreset")

	// static button
	uS := &ickui.ICKButton{}
	uS.OnClick = func() { browser.Win().Alert("clicked") }
	dom.MountId(uS, "uA2")

	// let's go
	fmt.Println("Icecake initialized. Listening browser events")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-running">running</span>`)
	}
	<-c
}
