// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example1.
//
// It's compiled into a '.wasm' file with the build_ex1 task
package main

import (
	"fmt"

	_ "embed"

	"github.com/sunraylab/icecake/pkg/extensions/markdown"
	"github.com/sunraylab/icecake/pkg/ick"
	"github.com/sunraylab/icecake/pkg/ui"
	wick "github.com/sunraylab/icecake/pkg/wicecake"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

//go:embed readme.md
var readme string

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	doc := wick.GetDocument()

	// 1. demonstrate the use of the go HTML templating package to build page content directly on the front-end.
	htmlTemplate := `Hello <strong>%s</strong>!`
	ui.RenderHtml(doc.ChildById("ex1a"), ick.HTML(fmt.Sprintf(htmlTemplate, "Bob")), nil)
	ui.RenderHtml(doc.ChildById("ex1b"), ick.HTML(fmt.Sprintf(htmlTemplate, "Alice")), nil)

	// To see what happend with a wrong html element ID,
	// open the console on the browser side.
	ui.RenderHtml(doc.ChildById("ex1c"), ick.HTML(fmt.Sprintf(htmlTemplate, "Carole")), nil)

	// 2. demonstrate how to generate HTML content from a markdown source, directly on the front-side.
	markdown.RenderMarkdown(doc.ChildById("ex1d"), "### Markdown\nHello **John**", nil)

	// Text source is embedded in the compiled wasm code with the //go:embed compiler directive
	ick.RegisterSnippet("ick-icecake-brand", "", "<span class='brand'>Icecake</span>")
	markdown.RenderMarkdown(doc.ChildById("readme"), readme, nil,
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
