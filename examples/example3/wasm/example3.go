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

	"github.com/sunraylab/icecake/pkg/dom"
	icecake "github.com/sunraylab/icecake/pkg/framework"
)

//go:embed "readme.md"
var readme string

var notif *CmpNotif

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// render introduction
	icecake.GetElementById("introduction").RenderMarkdown(readme, nil)

	icecake.RegisterComponentType("ick-notif", reflect.TypeOf(CmpNotif{}))

	// add simple event hendling
	icecake.GetButtonById("btn0").AddMouseEvent(dom.MOUSE_ONCLICK, OnClickBtn0)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

/******************************************************************************
* browser event handlers
******************************************************************************/

func OnClickBtn0(event *dom.MouseEvent, target *dom.HTMLElement) {

	//	icecake.GetElementById("notif0").RenderHtml("<ick-notif/>", nil)

	//	where := icecake.GetElementById("notif0") //.BuildComponent("<ick-notif/>", nil)

	// create the component and init it's data
	notif = &CmpNotif{}
	icecake.InsertComponent("notif0", notif)
}

/******************************************************************************
* Component
******************************************************************************/

type CmpNotif struct {
	dom.HTMLElement
}

func (c *CmpNotif) Template() string {
	return `
	<div class="notification is-warning is-light">
		<button class="delete"></button>
		Primar lorem ipsum dolor sit amet, consectetur
		adipiscing elit lorem ipsum dolor. <strong>Pellentesque risus mi</strong>, tempus quis placerat ut, porta nec nulla. Vestibulum rhoncus ac ex sit amet fringilla. Nullam gravida purus diam, et dictum <a>felis venenatis</a> efficitur.
  	</div>`
}

func (c *CmpNotif) AddListeners() {
	fmt.Println("**AddListeners**")
	qs := c.SelectorQueryFirst(".delete")

	fmt.Println("**qs:", qs)

	btndel := dom.CastHTMLElement(qs.JSValue())

	fmt.Println("**btndel: ", btndel)

	btndel.AddMouseEvent(dom.MOUSE_ONCLICK, func(*dom.MouseEvent, *dom.HTMLElement) {
		c.Remove()
	})
}
