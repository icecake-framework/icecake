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

	"github.com/sunraylab/icecake/pkg/html"
	ick "github.com/sunraylab/icecake/pkg/icecake"
	"github.com/sunraylab/icecake/pkg/markdown"
)

//go:embed "readme.md"
var readme string

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// render introduction
	markdown.RenderMarkdown(ick.GetElementById("introduction"), readme, nil)

	ick.GData["msgnumber"] = 0

	// add simple event hendling
	html.GetButtonById("btnw").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtnw)
	html.GetButtonById("btna").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtna)
	html.GetButtonById("btns").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtns)
	html.GetButtonById("btni").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtni)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

/******************************************************************************
* browser event handlers
******************************************************************************/

func OnClickBtnw(event *ick.MouseEvent, target *ick.Element) {

	// instantiate the NotificationToast component and init its data
	toast := &Notify{
		Message:    `This is a typical notification message <strong>including html</strong> event html links.<br/><a href="#">link1</a>&nbsp;|&nbsp;<a href="#">link2</a>`,
		ColorClass: "is-warning is-light",
	}

	// Insert the component into the DOM
	ick.GData["msgnumber"] = ick.GData["msgnumber"].(int) + 1
	if _, err := ick.GetElementById("notif_container").InsertNewComponent(toast); err != nil {
		ick.ConsoleErrorf(err.Error())
	}
}

func OnClickBtna(event *ick.MouseEvent, target *ick.Element) {

	// instantiate the NotificationToast component and init its data
	toast := &Notify{
		Message:    `This message will be automatically removed in <strong>4 seconds</strong>, unless you close it before. 😀`,
		ColorClass: "is-danger is-light",
		Timeout:    time.Second * 4,
	}
	//TODO: toast.Attributes().Set("class", "myclass")

	// Insert the component into the DOM
	ick.GData["msgnumber"] = ick.GData["msgnumber"].(int) + 1
	if _, err := ick.GetElementById("notif_container").InsertNewComponent(toast); err != nil {
		ick.ConsoleErrorf(err.Error())
	}
}

func OnClickBtns(event *ick.MouseEvent, target *ick.Element) {

	ick.GData["msgnumber"] = ick.GData["msgnumber"].(int) + 1

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
	ick.GData["msgnumber"] = ick.GData["msgnumber"].(int) + 1
	if err := ick.GetElementById("inside_container").RenderHtml(html, nil); err != nil {
		ick.ConsoleErrorf(err.Error())
	}

}
