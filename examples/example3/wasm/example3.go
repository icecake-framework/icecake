// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example3.
//
// It's compiled into a '.wasm' file with the build_ex2 task
package main

import (
	"fmt"
	"math"
	"time"

	_ "embed"

	"github.com/icecake-framework/icecake/pkg/clock"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/extensions/markdown"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/registry"
	"github.com/icecake-framework/icecake/pkg/ui"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	mdhtml "github.com/yuin/goldmark/renderer/html"
)

//go:embed "readme.md"
var readme string

// var webapp *dom.WebApp

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// we must call the icecake webapp factory once
	// webapp = dom.NewWebApp()

	// render readme
	markdown.RenderIn(dom.Id("readme"), readme, nil,
		goldmark.WithRendererOptions(
			mdhtml.WithUnsafe()),
		goldmark.WithExtensions(
			highlighting.Highlighting,
		))

	// add global event handling
	dom.Id("btn1").AddMouseEvent(event.MOUSE_ONCLICK, OnClickBtn1)
	dom.Id("btn2").AddMouseEvent(event.MOUSE_ONCLICK, OnClickBtn2)
	dom.Id("btn3").AddMouseEvent(event.MOUSE_ONCLICK, OnClickBtn3)
	dom.Id("btn4").AddMouseEvent(event.MOUSE_ONCLICK, OnClickBtn4)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

/******************************************************************************
* browser event handlers
******************************************************************************/

func OnClickBtn1(event *event.MouseEvent, target *dom.Element) {

	// instantiate the Notify component and init its data
	notif := &ui.Notify{}
	notif.Message = `This is a typical notification message <strong>including html <a href="#">link</a>.</strong> Use the closing button on the right corner to remove this notification.`
	notif.Tag().Attributes().SetClasses("is-warning is-light")

	// Insert the component into the DOM
	dom.Id("notif_container").InsertSnippet(dom.INSERT_LAST_CHILD, notif, nil)
}

func OnClickBtn2(event *event.MouseEvent, target *dom.Element) {

	// instantiate the Notify component and init its data
	idtimeleft := registry.GetUniqueId("timeleft")
	notif := new(ui.Notify)
	notif.Message = `This message will be automatically removed in <strong><span id="` + html.HTMLString(idtimeleft) + `"></span> seconds</strong>, unless you close it before. 😀`
	notif.Delete.Timeout = time.Second * 7
	notif.Tag().Attributes().SetClasses("is-danger is-light").SetAttribute("role", "alert", true)
	notif.Delete.Tic = func(clk *clock.Clock) {
		s := math.Round(notif.Delete.TimeLeft().Seconds())
		dom.Id(idtimeleft).InsertText(dom.INSERT_BODY, "%v", s)
	}

	// Insert the component into the DOM
	dom.Id("notif_container").InsertSnippet(dom.INSERT_LAST_CHILD, notif, nil)
}

func OnClickBtn3(event *event.MouseEvent, target *dom.Element) {

	// instantiate the Notify component and init its data
	notif := &ui.Notify{}
	notif.Message = `This is a toast notification`
	notif.Delete.Clock.Timeout = time.Second * 3
	notif.Tag().Attributes().SetClasses("is-success toast")

	// Insert the component into the DOM
	dom.Id("toast_container").InsertSnippet(dom.INSERT_LAST_CHILD, notif, nil)
}

func OnClickBtn4(event *event.MouseEvent, target *dom.Element) {

	// define the HTML template
	h := `<div class="box">
	<p class="pb-2">This is an html template object embedding the &lt;ick-notify&gt; element.</p>
	<div class="block">
		<ick-notify Message="This message comes from the Notify Component <strong>embedded into an html template</strong>."
		class="is-info is-light"
		role="success"/>
	</div>
	</box>`

	// Insert the component into the DOM
	dom.Id("ex3_container").InsertHTML(dom.INSERT_LAST_CHILD, html.HTMLString(h), nil)
}
