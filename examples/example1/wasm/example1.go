// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example1.
//
// It's compiled into a '.wasm' file with the build_ex1 task
package main

import (
	"fmt"

	_ "embed"

	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/extensions/markdown"
	"github.com/sunraylab/icecake/pkg/html"
	"github.com/sunraylab/icecake/pkg/ick"
	"github.com/yuin/goldmark"
	mdhtml "github.com/yuin/goldmark/renderer/html"
)

//go:embed readme.md
var readme string

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// 1. demonstrate the use of the go HTML templating package to build page content directly on the front-end.
	htmlTemplate := `Hello <strong>%s</strong>!`
	dom.Id("ex1a").InsertHTML(dom.INSERT_BODY, html.String(fmt.Sprintf(htmlTemplate, "Bob")), nil)
	dom.Id("ex1b").InsertHTML(dom.INSERT_BODY, html.String(fmt.Sprintf(htmlTemplate, "Alice")), nil)

	// To see what happend with a wrong html element ID,
	// open the console on the browser side.
	dom.Id("ex1c").InsertHTML(dom.INSERT_BODY, html.String(fmt.Sprintf(htmlTemplate, "Carole")), nil)

	// 2. demonstrate how to generate HTML content from a markdown source, directly on the front-side.
	markdown.RenderMarkdown(dom.Id("ex1d"), "### Markdown\nHello **John**", nil)

	// Text source is embedded in the compiled wasm code with the //go:embed compiler directive
	ick.RegisterDefaultSnippet("ick-icecake-brand", "", "<span class='brand'>Icecake</span>")
	markdown.RenderMarkdown(dom.Id("readme"), readme, nil,
		goldmark.WithRendererOptions(
			mdhtml.WithUnsafe(),
		),
	)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
