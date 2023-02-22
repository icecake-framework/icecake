// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example2.
//
// It's compiled into a '.wasm' file with the build_ex2 task
package main

import (
	"fmt"

	_ "embed"

	"github.com/sunraylab/icecake/pkg/dom"
	icecake "github.com/sunraylab/icecake/pkg/framework"
	"github.com/sunraylab/icecake/pkg/html"
)

//go:embed "introduction.md"
var introduction string

var gcount float64
var gdark bool

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// proceed with localstorage
	gdark := dom.GetWindow().LocalStorage().GetBool("darkmode")
	updateDarkMode(gdark)

	icecake.GetElementById("introduction").RenderMarkdown(introduction, nil)

	// init UI with a first update
	updateUI()

	// add simple event hendling
	html.GetButtonById("btn-ex2").AddMouseEvent(dom.MOUSE_ONCLICK, OnClickBtnEx2)
	html.GetButtonById("btn-lightmode").AddMouseEvent(dom.MOUSE_ONCLICK, OnClickBtnLightMode)
	html.GetButtonById("btn-darkmode").AddMouseEvent(dom.MOUSE_ONCLICK, OnClickBtnDarkMode)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

/******************************************************************************
* browser event handlers
******************************************************************************/

func OnClickBtnEx2(event *dom.MouseEvent, target *dom.Element) {
	gcount += 0.5
	updateUI()
}

func OnClickBtnLightMode(event *dom.MouseEvent, target *dom.Element) {
	updateDarkMode(false)
}

func OnClickBtnDarkMode(event *dom.MouseEvent, target *dom.Element) {
	updateDarkMode(true)
}

/******************************************************************************
* UI update
******************************************************************************/

func updateUI() {

	// 1st solution: render a value for any `data-ic-namedvalue="count1"` inside "sectionbody"
	icecake.GetElementById("sectionbody").RenderChildrenValue("count1", "%v", gcount)

	// 2nd solution: render a value inside an elem selected by it's id
	icecake.GetElementById("count1").RenderValue("%.1f", gcount)
}

func updateDarkMode(dark bool) {
	if dark {
		dom.GetDocument().Body().ClassList().Set("dark")
	} else {
		dom.GetDocument().Body().ClassList().Remove("dark")
	}
	html.GetButtonById("btn-lightmode").SetDisabled(!dark)
	html.GetButtonById("btn-darkmode").SetDisabled(dark)

	dom.GetWindow().LocalStorage().Set("darkmode", fmt.Sprintf("%v", dark))
}
