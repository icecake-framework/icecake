package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/browser"
	"github.com/icecake-framework/icecake/pkg/dom"
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

	// static button
	uA2 := &bulmaui.ICKButton{}
	uA2.OnClick = func() { browser.Win().Alert("clicked") }
	dom.MountId(&bulmaui.ICKNavbar{}, "uA2")

	// let's go
	fmt.Println("Icecake initialized. Listening browser events")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-running">running</span>`)
	}
	<-c
}
