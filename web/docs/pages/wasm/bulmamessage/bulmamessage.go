package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/bulmaui"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
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

	boxusage := dom.Id("boxusage")
	r1 := bulma.Message(html.ToHTML("This is for informations.")).SetColor(bulma.COLOR_INFO)
	r2 := bulma.Message(html.ToHTML("Click on the delete button to close this warning.")).SetColor(bulma.COLOR_WARNING).SetHeader(*html.ToHTML("Warning"), true)
	boxusage.InsertSnippet(dom.INSERT_BODY, r1, r2)

	// let's go
	fmt.Println("Icecake initialized. Listening browser events")
	if ests.IsDefined() {
		ests.InsertRawHTML(dom.INSERT_BODY, `<span class="ick-running">running</span>`)
	}
	<-c
}
