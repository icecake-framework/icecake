package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/bulmaui"
	"github.com/icecake-framework/icecake/pkg/dom"
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

	// burger := dom.Id("topbar")
	// burger.AddMouseEvent(event.MOUSE_ONCLICK, OnClickBurger)
	bulmaui.WrapNavbar("topbar")

	// let's go
	fmt.Println("Icecake initialized. Listening browser events")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-running">running</span>`)
	}
	<-c
}