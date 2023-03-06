// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example2.
//
// It's compiled into a '.wasm' file with the build_ex2 task
package main

import (
	"fmt"

	_ "embed"

	"github.com/sunraylab/icecake/pkg/extensions/markdown"
	ick "github.com/sunraylab/icecake/pkg/icecake"
	"github.com/sunraylab/icecake/pkg/ui"
)

//go:embed "readme.md"
var readme string

var webapp *ick.WebApp
var gcount float64

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	webapp = ick.NewWebApp()

	// proceed with localstorage
	gdark := webapp.Browser().LocalStorage().GetBool("darkmode")
	updateDarkMode(gdark)

	markdown.RenderMarkdown(webapp.ChildById("readme"), readme, nil)

	// init UI with a first update
	updateUI()

	// add simple event hendling
	ui.ButtonById("btn-ex2").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtnEx2)
	ui.ButtonById("btn-lightmode").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtnLightMode)
	ui.ButtonById("btn-darkmode").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtnDarkMode)

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
	webapp.ChildById("sectionbody").RenderChildrenValue("count1", "%v", gcount)

	// 2nd solution: render a value inside an elem selected by it's id
	webapp.ChildById("count1").RenderValue("%.1f", gcount)
}

func updateDarkMode(dark bool) {
	if dark {
		webapp.Body().Classes().SetTokens("dark")
	} else {
		webapp.Body().Classes().RemoveTokens("dark")
	}
	ui.ButtonById("btn-lightmode").SetDisabled(!dark)
	ui.ButtonById("btn-darkmode").SetDisabled(dark)

	webapp.Browser().LocalStorage().Set("darkmode", fmt.Sprintf("%v", dark))
}
