package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
)

// This main package contains the web assembly source code for the icecake example.
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	dom.MountCSSLinks()

	msg1 := &html.Message{Message: "This is a simple message without header."}
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, msg1, nil)

	msg2 := &html.Message{
		Header:  "simple message",
		Message: "This is a simple message with a header.",
	}
	msg2.Tag().Attributes().SetClasses("is-info")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, msg2, nil)

	msg3 := &html.Message{
		Header:    "message with delete button",
		CanDelete: true,
		Message:   "This message use the BULMA <i>is-warning</i> color class. The <i>CanDelete</i> property is set to true so the user can delete the message.",
	}
	msg3.Tag().Attributes().SetClasses("is-warning")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, msg3, nil)

	msg4 := &html.Message{
		Header:    "only header",
		CanDelete: true,
	}
	msg4.Tag().Attributes().SetClasses("is-success")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, msg4, nil)
	msg4.Tag().Attributes().SwitchClasses("is-success", "is-danger")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, msg4, nil)

	msg5 := &html.Message{}
	msg5.Tag().Attributes().SetClasses("is-danger")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, msg5, nil)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
