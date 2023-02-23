// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example2.
//
// It's compiled into a '.wasm' file with the build_ex2 task
package main

import (
	"fmt"

	_ "embed"

	"github.com/sunraylab/icecake/pkg/html"
	ick "github.com/sunraylab/icecake/pkg/icecake"
	"github.com/sunraylab/icecake/pkg/markdown"
)

//go:embed "introduction.md"
var introduction string

var gcount float64

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// proceed with localstorage
	gdark := ick.App().Browser().LocalStorage().GetBool("darkmode")
	updateDarkMode(gdark)

	markdown.RenderMarkdown(ick.GetElementById("introduction"), introduction, nil)

	// init UI with a first update
	updateUI()

	// add simple event hendling
	html.GetButtonById("btn-ex2").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtnEx2)
	html.GetButtonById("btn-lightmode").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtnLightMode)
	html.GetButtonById("btn-darkmode").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtnDarkMode)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

/******************************************************************************
* browser event handlers
******************************************************************************/

func OnClickBtnEx2(event *ick.MouseEvent, target *ick.Element) {
	gcount += 0.5
	updateUI()
}

func OnClickBtnLightMode(event *ick.MouseEvent, target *ick.Element) {
	updateDarkMode(false)
}

func OnClickBtnDarkMode(event *ick.MouseEvent, target *ick.Element) {
	updateDarkMode(true)
}

/******************************************************************************
* UI update
******************************************************************************/

func updateUI() {

	// 1st solution: render a value for any `data-ic-namedvalue="count1"` inside "sectionbody"
	ick.GetElementById("sectionbody").RenderChildrenValue("count1", "%v", gcount)

	// 2nd solution: render a value inside an elem selected by it's id
	ick.GetElementById("count1").RenderValue("%.1f", gcount)
}

func updateDarkMode(dark bool) {
	if dark {
		ick.App().Body().ClassList().Set("dark")
	} else {
		ick.App().Body().ClassList().Remove("dark")
	}
	html.GetButtonById("btn-lightmode").SetDisabled(!dark)
	html.GetButtonById("btn-darkmode").SetDisabled(dark)

	ick.App().Browser().LocalStorage().Set("darkmode", fmt.Sprintf("%v", dark))
}
