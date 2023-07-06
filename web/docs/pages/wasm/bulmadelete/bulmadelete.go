package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	bulmaui "github.com/icecake-framework/icecake/pkg/ick/ickui"
)

// This main package contains web assembly source code.
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded. Icecake initializing...")

	ests := dom.Id("icecake-status")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-initializing">initializing</span>`)
	}

	dom.MountId(&bulmaui.ICKNavbar{}, "topbar")

	dom.Id("btnreset").AddMouseEvent(event.MOUSE_ONCLICK, OnBtnReset)

	// let's go
	fmt.Println("Icecake initialized. Listening browser events")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-running">running</span>`)
	}
	<-c
}

func OnBtnReset(*event.MouseEvent, *dom.Element) {

}

func RenderUsage() {

}
