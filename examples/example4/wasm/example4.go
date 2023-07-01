// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example4.
//
// It's compiled into a '.wasm' file with the build_ex2 task
package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
)

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// msg1 := &ui.Message{
	// 	Header:  "1st message: <strong>Hello World</strong>",
	// 	Message: "This is the default message layout. It may contains <a href='#'>link</a> or any other HTML content. It can't be deleted by the user.",
	// }
	// ui.RenderSnippet(webapp.ChildById("msg-container"), msg1, nil)

	msg2 := &bulma.Message{}
	msg2.Header = *html.ToHTML("2nd message")
	msg2.CanDelete = true
	msg2.Msg = *html.ToHTML("This second message use the BULMA <i>is-info</i> color class. The <i>CanDelete</i> property is set to true so the user can delete the message.")
	msg2.Tag().AddClasses("is-info")
	dom.Id("msg-container").InsertSnippet(dom.INSERT_LAST_CHILD, msg2, nil)

	// msg3 := &ui.Message{
	// 	Message: "<strong>3rd message:</strong>&nbsp;this third message don't have header",
	// }
	// msg3.SetClasses("is-warning")
	// ui.RenderSnippet(webapp.ChildById("msg-container"), msg3, nil)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
