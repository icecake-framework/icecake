package main

import (
	"fmt"
	"math"
	"time"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/clock"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/registry"
	"github.com/icecake-framework/icecake/pkg/ui"
)

// This main package contains the web assembly source code for the icecake example.
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// 1st notification example
	// simplest
	notif1 := &ui.Notify{Notify: bulma.Notify{
		Message: html.String("This is a simple notification message. Use the closing button on the right corner to remove this notification."),
	}}
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, notif1, nil)

	// 2nd notification example
	// simple wuth custom classes
	notif2 := &ui.Notify{Notify: bulma.Notify{
		Message: html.String("This is another simple notification."),
	}}
	notif2.Tag().Attributes().AddClasses("is-success is-light")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, notif2, nil)

	// 3rd notification message
	// autoclosing with a timeout
	notif3 := &ui.Notify{Notify: bulma.Notify{
		Message: html.String(`This message will be automatically removed in few seconds, unless you close it before. 😀`)}}
	notif3.Delete.Timeout = time.Second * 5
	notif3.Tag().Attributes().AddClasses("is-danger is-light").SetAttribute("role", "alert")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, notif3, nil)

	// 4th notification message
	// autoclosing anf display the ticker
	idtimeleft := registry.GetUniqueId("timeleft")
	notif4 := &ui.Notify{}
	notif4.Message = html.String(`This message will be automatically removed in <strong><span id="` + idtimeleft + `"></span> seconds</strong>`)
	notif4.Tag().Attributes().AddClasses("is-warning is-light")
	//notif4.Delete.OnDelete = OnCloseNotif
	notif4.Delete.Timeout = time.Second * 10
	notif4.Delete.Tic = func(clk *clock.Clock) {
		s := math.Round(notif4.Delete.TimeLeft().Seconds())
		dom.Id(idtimeleft).InsertText(dom.INSERT_BODY, "%v", s)
	}
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, notif4, nil)

	// 5th notification message
	// embedded into another html
	h := `<ick-notify Message="This notify component is <strong>embedded into an html string</strong>." class="is-info is-light" role="success"/>`
	dom.Id("content").InsertHTML(dom.INSERT_LAST_CHILD, html.String(h), nil)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

// func OnCloseNotif(*ui.Delete) {
// 	console.Warnf("OnCloseNotif called")
// }
