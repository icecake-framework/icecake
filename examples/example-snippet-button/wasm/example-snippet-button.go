package main

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/browser"
	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/event"
	"github.com/sunraylab/icecake/pkg/ui"
)

var btn []*ui.Button

// This main package contains the web assembly source code for the icecake example.
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	btn = make([]*ui.Button, 10)
	btn[0] = &ui.Button{Title: "Click here"}
	btn[0].SetClasses("m-2 is-link is-light")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, btn[0], nil)
	btn[0].DOM.AddMouseEvent(event.MOUSE_ONCLICK, OnClick)

	btn[1] = &ui.Button{Title: "Toggle Rounded"}
	btn[1].SetClasses("m-2 is-link is-light")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, btn[1], nil)
	btn[1].DOM.AddMouseEvent(event.MOUSE_ONCLICK, OnClickRounded)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

func OnClick(_evt *event.MouseEvent, _elem *dom.Element) {
	browser.Win().Alert("clicked")
}

func OnClickRounded(_evt *event.MouseEvent, _elem *dom.Element) {
	for _, b := range btn {
		if b != nil {
			b.Rounded(!b.IsRounded)
		}
	}
}
