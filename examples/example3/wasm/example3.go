// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example3.
//
// It's compiled into a '.wasm' file with the build_ex2 task
package main

import (
	"fmt"
	"time"

	_ "embed"

	"github.com/sunraylab/icecake/pkg/errors"
	"github.com/sunraylab/icecake/pkg/extensions/markdown"
	ick "github.com/sunraylab/icecake/pkg/icecake"
	"github.com/sunraylab/icecake/pkg/uicomponent"
)

//go:embed "readme.md"
var readme string

// our own APP with global app data that can be used within components templates
type myApp struct {
	*ick.WebApp // embedded icecake app structure

	LastMsgNumber int // the last message number used by our components in this example. Must be scopped outside the package (so start with an uppercase) to enable reading by html template
}

var app myApp

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// we must call the icecake webapp factory once
	app.WebApp = ick.NewWebApp()

	// we can make some app init
	app.LastMsgNumber = 0

	// render introduction
	markdown.RenderMarkdown(app.ChildById("introduction"), readme, nil)

	// app.Body().FilteredChildren(ick.NT_ELEMENT, 9, func(n *ick.Node) bool {
	// 	e := ick.CastElement(n.JSValue())
	// 	switch tagname := e.TagName(); tagname {
	// 	case "H1":
	// 		e.ClassList().Set("title is-1")
	// 	case "H2":
	// 		e.ClassList().Set("title is-2")
	// 	case "H3":
	// 		e.ClassList().Set("title is-3")
	// 	case "H4":
	// 		e.ClassList().Set("title is-4")
	// 	default:
	// 		ick.ConsoleWarnf("%s\n", tagname)
	// 	}
	// 	return false
	// })

	// add simple event hendling
	app.ChildById("btnw").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtnw)
	app.ChildById("btna").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtna)
	app.ChildById("btns").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtns)
	app.ChildById("btni").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtni)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

/******************************************************************************
* browser event handlers
******************************************************************************/

func OnClickBtnw(event *ick.MouseEvent, target *ick.Element) {

	// instantiate the NotificationToast component and init its data
	toast := &uicomponent.Notify{
		Message: `This is a typical notification message <strong>including html</strong> event html links.<br/><a href="#">link1</a>&nbsp;|&nbsp;<a href="#">link2</a>`,
		// ColorClass: "is-warning is-light",
	}
	toast.SetClasses("is-warning is-light")

	// Insert the component into the DOM
	app.LastMsgNumber += 1
	if _, err := app.ChildById("notif_container").InsertNewComponent(toast, app); err != nil {
		errors.ConsoleErrorf(err.Error())
	}
}

func OnClickBtna(event *ick.MouseEvent, target *ick.Element) {

	// instantiate the NotificationToast component and init its data
	toast := &uicomponent.Notify{
		Message: `This message will be automatically removed in <strong>4 seconds</strong>, unless you close it before. 😀`,
		// ColorClass: "is-danger is-light",
		Timeout: time.Second * 4,
	}
	toast.SetClasses("is-danger is-light")
	//TODO: toast.Attributes().Set("class", "myclass")

	// Insert the component into the DOM
	app.LastMsgNumber += 1
	if _, err := app.ChildById("notif_container").InsertNewComponent(toast, app); err != nil {
		errors.ConsoleErrorf(err.Error())
	}
}

func OnClickBtns(event *ick.MouseEvent, target *ick.Element) {

	app.LastMsgNumber += 1

	// instantiate the NotificationToast component and init its data
	// toast := &NotificationToast{
	// 	Message:    `Single`,
	// 	ColorClass: "is-primary",
	// }

	// // Insert the component into the DOM
	// singlenotifid, err := ick.GetElementById("notif_container").InsertNewComponent(toast)
	// if err != nil {
	// 	ick.ConsoleErrorf(err.Error())
	// }
}

func OnClickBtni(event *ick.MouseEvent, target *ick.Element) {

	html := `<div class="box"><p class="pb-2">This is an html template object embedding the &lt;ick-notify&gt; element.</p><div class="block">
	<ick-notify 
	Message="This message comes from the Notify Component <strong>embedded into an html template</strong>."
	ColorClass="is-info" 
	role="alert"/>
	</div></box>`

	// Insert the component into the DOM
	app.LastMsgNumber += 1
	if err := app.ChildById("inside_container").RenderHtml(html, nil); err != nil {
		errors.ConsoleErrorf(err.Error())
	}

}
