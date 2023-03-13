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

	"github.com/sunraylab/icecake/pkg/clock"
	"github.com/sunraylab/icecake/pkg/extensions/markdown"
	ick "github.com/sunraylab/icecake/pkg/icecake"
	"github.com/sunraylab/icecake/pkg/ui"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/renderer/html"
)

//go:embed "readme.md"
var readme string

var webapp *ick.WebApp

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// we must call the icecake webapp factory once
	webapp = ick.NewWebApp()

	// render readme
	markdown.RenderMarkdown(webapp.ChildById("readme"), readme, nil,
		goldmark.WithRendererOptions(
			html.WithUnsafe()),
		goldmark.WithExtensions(
			highlighting.Highlighting,
		))

	// add global event handling
	webapp.ChildById("btn1").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtn1)
	webapp.ChildById("btn2").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtn2)
	webapp.ChildById("btn3").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtn3)
	webapp.ChildById("btn4").AddMouseEvent(ick.MOUSE_ONCLICK, OnClickBtn4)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

/******************************************************************************
* browser event handlers
******************************************************************************/

func OnClickBtn1(event *ick.MouseEvent, target *ick.Element) {

	// instantiate the Notify component and init its data
	notif := &ui.Notify{
		Message: `This is a typical notification message <strong>including html <a href="#">link</a>.</strong> Use the closing button on the right corner to remove this notification.`,
	}
	notif.Classes().AddTokens("is-warning", "is-light")

	// Insert the component into the DOM
	webapp.ChildById("notif_container").RenderComponent(notif, nil)
}

func OnClickBtn2(event *ick.MouseEvent, target *ick.Element) {

	// instantiate the Notify component and init its data
	notif := new(ui.Notify)
	notif.Message = `This message will be automatically removed in <strong><span id="{{.Id}}-timeleft"></span> seconds</strong>, unless you close it before. 😀`
	notif.Timeout = time.Second * 7
	notif.Classes().AddTokens("is-danger", "is-light")
	notif.Attributes().SetAttribute("role", "alert")
	notif.Tic = func(clk *clock.Clock) {
		fmt.Println("tic", notif.Id())
		s := math.Round(notif.Delete.TimeLeft().Seconds())
		idtl := notif.Id() + "-timeleft"
		webapp.ChildById(idtl).RenderValue("%v", s)
	}

	// Insert the component into the DOM
	webapp.ChildById("notif_container").RenderComponent(notif, nil)
}

func OnClickBtn3(event *ick.MouseEvent, target *ick.Element) {

	// instantiate the Notify component and init its data
	notif := &ui.Notify{
		Message: `This is a toast notification`,
	}
	notif.Delete.Clock.Timeout = time.Second * 3
	notif.Classes().AddTokens("is-success", "toast")

	// Insert the component into the DOM
	webapp.ChildById("toast_container").RenderComponent(notif, nil)
}

func OnClickBtn4(event *ick.MouseEvent, target *ick.Element) {

	// define the HTML template
	html := `<div class="box">
	<p class="pb-2">This is an html template object embedding the &lt;ick-notify&gt; element.</p>
	<div class="block">
		<ick-notify Message="This message comes from the Notify Component <strong>embedded into an html template</strong>."
		class="is-info is-light"
		role="success"/>
	</div>
	</box>`

	// Insert the component into the DOM
	webapp.ChildById("ex3_container").RenderTemplate(html, nil)
}
