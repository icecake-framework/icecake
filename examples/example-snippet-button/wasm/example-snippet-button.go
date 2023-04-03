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

	btn = make([]*ui.Button, 0)

	b0 := &ui.Button{Title: "Click here"}
	b0.SetClasses("m-2 is-link is-light")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, b0, nil)
	b0.DOM.AddMouseEvent(event.MOUSE_ONCLICK, OnClick)
	btn = append(btn, b0)

	b1 := &ui.Button{Title: "Toggle Rounded"}
	b1.SetClasses("m-2 is-link is-light").SetAttribute("data-example", 1)
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, b1, nil)
	b1.DOM.AddMouseEvent(event.MOUSE_ONCLICK, OnClickExample)
	btn = append(btn, b1)

	b2 := &ui.Button{Title: "Toggle Outlined"}
	b2.SetClasses("m-2 is-link is-light").SetAttribute("data-example", 2)
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, b2, nil)
	b2.DOM.AddMouseEvent(event.MOUSE_ONCLICK, OnClickExample)
	btn = append(btn, b2)

	b3 := &ui.Button{Title: "Toggle Loading"}
	b3.SetClasses("m-2 is-link is-light").SetAttribute("data-example", 3)
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, b3, nil)
	b3.DOM.AddMouseEvent(event.MOUSE_ONCLICK, OnClickExample)
	btn = append(btn, b3)

	b4 := &ui.Button{Title: "Toggle Disabled"}
	b4.SetClasses("m-2 is-link is-light").SetAttribute("data-example", 4)
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, b4, nil)
	b4.DOM.AddMouseEvent(event.MOUSE_ONCLICK, OnClickExample)
	btn = append(btn, b4)

	b5 := &ui.Button{Title: "Go To Home", ButtonType: ui.BTN_TYPE_A}
	b5.HRef.Path = "/"
	b5.SetClasses("m-2 is-link is-light").SetAttribute("data-example", 5)
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, b5, nil)
	btn = append(btn, b5)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

func OnClick(_evt *event.MouseEvent, _elem *dom.Element) {
	browser.Win().Alert("clicked")
}

func OnClickExample(_evt *event.MouseEvent, _elem *dom.Element) {
	example, _ := _elem.Attribute("data-example")
	for i, b := range btn {
		if b != nil {
			switch example {
			case "1":
				b.SetRounded(!b.IsRounded)
			case "2":
				b.SetOutlined(!b.IsOutlined)
			case "3":
				if i != 3 {
					b.SetLoading(!b.IsLoading)
				}
			case "4":
				if i != 4 {
					b.SetDisabled(!b.IsDisabled())
				}
			default:
				browser.Win().Alert("example " + example + " clicked")
			}
		}
	}
}
