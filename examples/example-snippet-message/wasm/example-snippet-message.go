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

	msg1 := &ui.Message{
		Message: "This is a simple message without header.",
	}
	dom.Id("content").RenderSnippet(dom.INSERT_LAST_CHILD, msg1, nil)

	msg2 := &ui.Message{
		Header:  "simple message",
		Message: "This is a simple message with a header.",
	}
	msg2.SetClasses("is-info")
	dom.Id("content").RenderSnippet(dom.INSERT_LAST_CHILD, msg2, nil)

	msg3 := &ui.Message{
		Header:    "message with delete button",
		CanDelete: true,
		Message:   "This message use the BULMA <i>is-warning</i> color class. The <i>CanDelete</i> property is set to true so the user can delete the message.",
	}
	msg3.SetClasses("is-warning")
	dom.Id("content").RenderSnippet(dom.INSERT_LAST_CHILD, msg3, nil)

	msg4 := &ui.Message{
		Header:    "only header",
		CanDelete: true,
	}
	msg4.SetClasses("is-success")
	dom.Id("content").RenderSnippet(dom.INSERT_LAST_CHILD, msg4, nil)
	msg4.SwitchClass("is-success", "is-danger")
	dom.Id("content").RenderSnippet(dom.INSERT_LAST_CHILD, msg4, nil)

	msg5 := &ui.Message{}
	msg5.SetClasses("is-danger")
	dom.Id("content").RenderSnippet(dom.INSERT_LAST_CHILD, msg5, nil)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}