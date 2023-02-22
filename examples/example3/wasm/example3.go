// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example2.
//
// It's compiled into a '.wasm' file with the build_ex2 task
package main

import (
	"fmt"
	"reflect"

	_ "embed"

	"github.com/sunraylab/icecake/pkg/html"
	ick "github.com/sunraylab/icecake/pkg/icecake"
)

//go:embed "readme.md"
var readme string

//var notif *CmpNotif

func init() {
	ick.RegisterComponentType("ick-notiftoast", reflect.TypeOf(NotificationToast{}))
}

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// render introduction
	ick.GetElementById("introduction").RenderMarkdown(readme, nil)

	// add simple event hendling
	html.GetButtonById("btn0").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtn0)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

/******************************************************************************
* browser event handlers
******************************************************************************/

func OnClickBtn0(event *ick.MouseEvent, target *ick.Element) {

	//	where := icecake.GetElementById("notif0") //.BuildComponent("<ick-notif/>", nil)

	// create the component and init its data
	toast := &NotificationToast{
		Notification: `Primar lorem ipsum dolor sit amet, consectetur adipiscing elit lorem ipsum dolor. 
		<strong>Pellentesque risus mi</strong>, tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. 
		Nullam <a>gravida purus diam</a>, et dictum felis venenatis efficitur.`,
	}
	toast.Color = "is-warning is-light"
	err := ick.InsertComponent(toast, "notif0")
	if err != nil {
		ick.ConsoleErrorf(err.Error())
	}
}

/******************************************************************************
* Component
******************************************************************************/

type NotificationToast struct {
	ick.Element

	Color        string
	Notification string
}

func (c *NotificationToast) Envelope() (_tagname string, _classname string) {
	return "div", "notification {{.Color}}"
}

func (c *NotificationToast) Template() (_html string) {
	return `<button class="delete"></button>
			{{.Notification}}`
}

func (c *NotificationToast) AddListeners() {
	btndel := c.SelectorQueryFirst(".delete")
	btndel.AddMouseEvent(ick.MOUSE_ONCLICK, func(*ick.MouseEvent, *ick.Element) {
		c.Remove()
	})
}
