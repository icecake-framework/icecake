// Copyright 2023 by lolorenzo777. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code for the icecake example1.
//
// It's compiled into a '.wasm' file with the build_ex1 task
package main

import (
	"fmt"

	_ "embed"

	_ "github.com/icecake-framework/icecake/pkg/bulmaui" // automatic registering of the ui components

	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/extensions/markdown"
	"github.com/icecake-framework/icecake/pkg/html"
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

	// register a tiny html snippet
	tiny := html.NewSnippet("a", `class="brand" href="https://icecake.net"`).AddContent(html.ToHTML("<strong>Icecake</strong>"))
	html.RegisterComposer("ick-icecake-brand", tiny)

	// Text source is embedded in the compiled wasm code with the //go:embed compiler directive
	// 2. demonstrate how to generate HTML content from a markdown source, directly on the front-side.
	markdown.RenderIn(dom.Id("readme"), readme, nil,
		goldmark.WithRendererOptions(
			mdhtml.WithUnsafe(),
		),
	)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
