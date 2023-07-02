package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/browser"
	"github.com/icecake-framework/icecake/pkg/bulmaui"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
)

// This main package contains the web assembly source code for makedocs
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded. Icecake initializing...")

	ests := dom.Id("icecake-status")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-initializing">initializing</span>`)
	}

	bulmaui.WrapNavbar("topbar")

	// static button
	uA2 := bulmaui.WrapButton("uA2")
	uA2.DOM.AddMouseEvent(event.MOUSE_ONCLICK, func(*event.MouseEvent, *dom.Element) {
		browser.Win().Alert("clicked")
	})

	// let's go
	fmt.Println("Icecake initialized. Listening browser events")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-running">running</span>`)
	}
	<-c
}
