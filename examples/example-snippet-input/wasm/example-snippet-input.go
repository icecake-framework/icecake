package main

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/event"
	"github.com/sunraylab/icecake/pkg/ui"
)

var in5 *ui.InputField

// This main package contains the web assembly source code for the icecake example.
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	dom.MountCSSLinks()

	in1 := &ui.InputField{
		PlaceHolder: "Very simple",
	}
	in1.SetClasses("mr-4")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, in1, nil)

	in2 := &ui.InputField{
		Label:       "Name",
		PlaceHolder: "Text input",
		Help:        "With a label, a placeholder, and a help",
	}
	in2.SetClasses("mr-4")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, in2, nil)

	in3 := &ui.InputField{
		Label:       "Username",
		PlaceHolder: "Text input",
		Help:        "Rounded style",
		Value:       "my name",
		IsRounded:   true,
	}
	in3.SetClasses("mr-4")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, in3, nil)

	in4 := &ui.InputField{
		Label:       "Loading",
		PlaceHolder: "Text input",
		State:       ui.INPUT_LOADING,
		Help:        "With loading state",
	}
	in4.SetClasses("mr-4")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, in4, nil)

	in5 = &ui.InputField{
		Label:       "eMail",
		PlaceHolder: "email address",
		Help:        "Enter a valid email address",
	}
	in5.SetClasses("mr-4")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, in5, nil)
	in5.DOM.AddInputEvent(event.INPUT_ONINPUT, OnInput)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

func OnInput(_event *event.InputEvent, _target *dom.Element) {
	v := _target.JSValue.String()
	console.Warnf("targetvalue: %s, cmpvalue: %s", v, in5.Value)
}
