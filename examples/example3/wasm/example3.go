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
	toast := &NotificationToast{
		Message: `Primar lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. 
		<strong>Pellentesque risus mi</strong>, tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. 
		Nullam <a>gravida purus diam</a>, et dictum felis venenatis efficitur.`,
		ColorClass: "is-warning is-light",
	}

	// Insert the component into the DOM
	if _, err := ick.GetElementById("notif_container").InsertNewComponent(toast); err != nil {
		ick.ConsoleErrorf(err.Error())
	}
}

func OnClickBtna(event *ick.MouseEvent, target *ick.Element) {

	// instantiate the NotificationToast component and init its data
	toast := &NotificationToast{
		Message:    `lorem ipsum dolor sit amet`,
		ColorClass: "is-danger",
		Timeout:    time.Second * 4,
	}

	// Insert the component into the DOM
	if _, err := ick.GetElementById("notif_container").InsertNewComponent(toast); err != nil {
		ick.ConsoleErrorf(err.Error())
	}
}

func OnClickBtns(event *ick.MouseEvent, target *ick.Element) {

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

	html := `<div class="box">I'm in a box.<div class="block">
	<ick-notiftoast message="This is the message" colorclass="is-info"/>
	</div></box>`

	// Insert the component into the DOM
	if err := ick.GetElementById("inside_container").RenderHtml(html, nil); err != nil {
		ick.ConsoleErrorf(err.Error())
	}
}
