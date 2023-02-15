// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example2.
//
// It's compiled into a '.wasm' file with the build_ex2 task
package main

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/dom"
	icecake "github.com/sunraylab/icecake/pkg/framework"
)

var count int

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	renderCount()

	icecake.GetButtonById("btn-ex2").AddMouseEvent(dom.MOUSE_ONCLICK, OnClickBtnEx2)

	icecake.GetButtonById("btn-lightmode").AddMouseEvent(dom.MOUSE_ONCLICK, OnClickBtnLightMode)
	icecake.GetButtonById("btn-darkmode").AddMouseEvent(dom.MOUSE_ONCLICK, OnClickBtnDarkMode)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

/******************************************************************************
* browser event handlers
******************************************************************************/

func OnClickBtnEx2(event *dom.MouseEvent, target *dom.HTMLElement) {
	count++
	renderCount()
}

func OnClickBtnLightMode(event *dom.MouseEvent, target *dom.HTMLElement) {
	icecake.DocumentBody().ClassList().Remove("dark")
}

func OnClickBtnDarkMode(event *dom.MouseEvent, target *dom.HTMLElement) {
	icecake.DocumentBody().ClassList().Set("dark")
}

/******************************************************************************
* UI update
******************************************************************************/

func renderCount() {

	// 1st solution:
	icecake.DocumentBody().RenderChildrenValue("count1", "%v", count)

	// 2nd solution: render a value inside an elem selected by it's id
	icecake.GetElementById("count1").RenderValue("%v", count)
}
