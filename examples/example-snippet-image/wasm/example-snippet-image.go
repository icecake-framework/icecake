package main

import (
	"fmt"
	"net/url"

	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ui"
)

// This main package contains the web assembly source code for the icecake example.
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	dom.MountCSSLinks()

	img1 := &ui.Image{Size: ui.IMG_96x96}
	img1.SetClasses("mr-4")
	img1.Src, _ = url.Parse("/icecake.jpg")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, img1, nil)

	img2 := &ui.Image{Size: ui.IMG_96x96, IsRounded: true}
	img2.SetClasses("mr-4")
	img2.Src, _ = url.Parse("/icecake.jpg")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, img2, nil)

	img3 := &ui.Image{Size: ui.IMG_96x96}
	img3.SetClasses("mr-4")
	img3.Src, _ = url.Parse("/icecake.svg")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, img3, nil)

	img4 := html.String(`<ick-image Size="96x96" Src="/icecake.svg"/>`)
	dom.Id("content").InsertHTML(dom.INSERT_LAST_CHILD, img4, nil)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
