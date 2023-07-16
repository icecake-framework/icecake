package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ick/ickui"
	"github.com/icecake-framework/icecake/pkg/ickcore"
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

	_btnreset.OnClick = func() { ResetBoxUsage() }

	// wrap the back rendered navbar and reset btn
	dom.MountId(&ickui.ICKNavbar{}, "topbar")
	dom.MountId(_btnreset, "btnreset")

	// front rendering
	ResetBoxUsage()

	// let's go
	fmt.Println("Icecake initialized. Listening browser events")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-running">running</span>`)
	}
	<-c
}

func ResetBoxUsage() {
	btndel := ickui.Delete("btndelu1", "deleteme")
	btndel.OnDelete = OnDeleteU1
	u1 := ick.Elem("div", `id="deleteme"`, ickcore.ToHTML("Click on the delete button to delete this text &rarr; "), btndel)
	dom.Id("boxusage").InsertSnippet(dom.INSERT_BODY, u1)

	_btnreset.SetDisabled(true)
}

func OnDeleteU1(del *ickui.ICKDelete) {
	_btnreset.SetDisabled(false)
}
