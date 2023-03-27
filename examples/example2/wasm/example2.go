// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example2.
//
// It's compiled into a '.wasm' file with the build_ex2 task
package main

import (
	"fmt"

	_ "embed"

	"github.com/sunraylab/icecake/pkg/browser"
	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/event"
	"github.com/sunraylab/icecake/pkg/extensions/markdown"
	"github.com/sunraylab/icecake/pkg/ui"
)

//go:embed "readme.md"
var readme string

// var webapp *dom.WebApp
var gcount float64

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// webapp = dom.NewWebApp()

	// proceed with localstorage
	gdark := browser.Win().LocalStorage().GetBool("darkmode")
	updateDarkMode(gdark)

	markdown.RenderMarkdown(dom.Id("readme"), readme, nil)

	// init UI with a first update
	updateUI()

	// add simple event hendling
	ui.ButtonById("btn-ex2").AddMouseEvent(event.MOUSE_ONCLICK, OnClickBtnEx2)
	ui.ButtonById("btn-lightmode").AddMouseEvent(event.MOUSE_ONCLICK, OnClickBtnLightMode)
	ui.ButtonById("btn-darkmode").AddMouseEvent(event.MOUSE_ONCLICK, OnClickBtnDarkMode)

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

	// 1st solution: render a value for any `data-ic-namedvalue="count1"` inside "sectionbody"
	dom.Id("sectionbody").RenderChildrenValue("count1", "%v", gcount)

	// 2nd solution: render a value inside an elem selected by it's id
	dom.Id("count1").RenderValue("%.1f", gcount)
}

func updateDarkMode(dark bool) {
	if dark {
		dom.Doc().Body().Classes().AddTokens("dark")
	} else {
		dom.Doc().Body().Classes().RemoveTokens("dark")
	}
	ui.ButtonById("btn-lightmode").SetDisabled(!dark)
	ui.ButtonById("btn-darkmode").SetDisabled(dark)

	browser.Win().LocalStorage().Set("darkmode", fmt.Sprintf("%v", dark))
}
