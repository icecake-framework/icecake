package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
)

// This main package contains the web assembly source code for the icecake example.
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	msg1 := &bulma.Message{Msg: *html.ToHTML("This is a simple message without header.")}
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, msg1, nil)

	msg2 := &bulma.Message{
		Header: *html.ToHTML("simple message"),
		Msg:    *html.ToHTML("This is a simple message with a header."),
	}
	msg2.Tag().AddClasses("is-info")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, msg2, nil)

	msg3 := &bulma.Message{
		Header:    *html.ToHTML("message with delete button"),
		CanDelete: true,
		Msg:       *html.ToHTML("This message use the BULMA <i>is-warning</i> color class. The <i>CanDelete</i> property is set to true so the user can delete the message."),
	}
	msg3.Tag().AddClasses("is-warning")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, msg3, nil)

	msg4 := &bulma.Message{
		Header:    *html.ToHTML("only header"),
		CanDelete: true,
	}
	msg4.Tag().AddClasses("is-success")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, msg4, nil)
	msg4.Tag().SwitchClass("is-success", "is-danger")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, msg4, nil)

	msg5 := &bulma.Message{}
	msg5.Tag().AddClasses("is-danger")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, msg5, nil)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
