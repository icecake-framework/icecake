package main

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/ui"
)

// This main package contains the web assembly source code for the icecake example.
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	dom.MountCSSLinks()

	card1 := &ui.Card{
		Title:       "Hello World",
		Content:     "Nice cake",
		FooterItem1: "<a href='/'>home</a>",
	}
	card1.SetStyle("width: 350px;").SetClasses("mr-5")
	card1.Image = ui.NewImage("/icecake.jpg", ui.IMG_SQUARE)
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, card1, nil)

	card2 := &ui.Card{Content: "Nice cake"}
	card2.SetStyle("width: 128px;").SetClasses("mr-5")
	card2.Image = ui.NewImage("/icecake.jpg", ui.IMG_SQUARE)
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, card2, nil)

	card3 := &ui.Card{Content: "Very Nice cake"}
	card3.SetClasses("mr-5")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, card3, nil)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
