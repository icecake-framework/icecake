// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example2.
//
// It's compiled into a '.wasm' file with the build_ex2 task
package main

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/browser"
	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/event"
)

var gcount float64

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	dom.MountCSSLinks()

	// proceed with localstorage
	gdark := browser.LocalStorage().GetBool("darkmode")
	updateDarkMode(gdark)

	// init UI with a first update
	updateUI()

	// add simple event hendling
	dom.Id("btn-ex2").AddMouseEvent(event.MOUSE_ONCLICK, OnClickBtnEx2)
	dom.Id("btn-lightmode").AddMouseEvent(event.MOUSE_ONCLICK, OnClickBtnLightMode)
	dom.Id("btn-darkmode").AddMouseEvent(event.MOUSE_ONCLICK, OnClickBtnDarkMode)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

/******************************************************************************
* browser event handlers
******************************************************************************/

func OnClickBtnEx2(event *event.MouseEvent, target *dom.Element) {
	gcount += 0.5
	updateUI()
}

func OnClickBtnLightMode(event *event.MouseEvent, target *dom.Element) {
	updateDarkMode(false)
}

func OnClickBtnDarkMode(event *event.MouseEvent, target *dom.Element) {
	updateDarkMode(true)
}

/******************************************************************************
* UI update
******************************************************************************/

func updateUI() {
	// render a value for any `data-ic-namedvalue="count1"` inside "sectionbody"
	children := dom.Id("sectionbody").ChildrenByData("data-ick-namedvalue", "count1")
	for _, e := range children {
		e.InsertText(dom.INSERT_BODY, "%.1f", gcount)
	}

	// render a value inside an elem selected by it's id
	dom.Id("count1").InsertText(dom.INSERT_BODY, "%.1f", gcount)
}

func updateDarkMode(dark bool) {
	if dark {
		dom.Doc().Body().SetClasses("dark")
	} else {
		dom.Doc().Body().RemoveClasses("dark")
	}
	sdark := "false"
	if dark {
		sdark = "true"
	}
	dom.Id("btn-lightmode").SetAttribute("dark", sdark)
	dom.Id("btn-darkmode").SetAttribute("dark", sdark)

	browser.LocalStorage().Set("darkmode", fmt.Sprintf("%v", dark))
}
