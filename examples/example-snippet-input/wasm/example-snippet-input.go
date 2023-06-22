package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ui"
)

var in5 *ui.InputField

// This main package contains the web assembly source code for the icecake example.
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	in1 := &ui.InputField{}
	in1.PlaceHolder = "Very simple"
	in1.Tag().AddClasses("mr-4")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, in1, nil)

	in2 := &ui.InputField{}
	in2.Label = html.HTML("Name")
	in2.PlaceHolder = "Text input"
	in2.Help = html.HTML("With a label, a placeholder, and a help")
	in2.Tag().AddClasses("mr-4")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, in2, nil)

	in3 := &ui.InputField{}
	in3.Label = html.HTML("Username")
	in3.PlaceHolder = "Text input"
	in3.Help = html.HTML("Rounded style")
	in3.Value = "my name"
	in3.IsRounded = true
	in3.Tag().AddClasses("mr-4")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, in3, nil)

	in4 := &ui.InputField{}
	in4.Label = html.HTML("Loading")
	in4.PlaceHolder = "Text input"
	in4.State = bulma.INPUT_LOADING
	in4.Help = html.HTML("With loading state")
	in4.Tag().AddClasses("mr-4")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, in4, nil)

	in5 = &ui.InputField{}
	in5.Label = html.HTML("eMail")
	in5.PlaceHolder = "email address"
	in5.Help = html.HTML("Enter a valid email address")
	in5.Tag().AddClasses("mr-4")
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
