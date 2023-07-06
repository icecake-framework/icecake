package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/browser"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/ick/ickui"
)

// This main package contains web assembly source code.
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded. Icecake initializing...")

	ests := dom.Id("icecake-status")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-initializing">initializing</span>`)
	}

	dom.MountId(&ickui.ICKNavbar{}, "topbar")

	// static button
	uA2 := &ickui.ICKButton{}
	uA2.OnClick = func() { browser.Win().Alert("clicked") }
	dom.MountId(&ickui.ICKNavbar{}, "uA2")

	// let's go
	fmt.Println("Icecake initialized. Listening browser events")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-running">running</span>`)
	}
	<-c
}
