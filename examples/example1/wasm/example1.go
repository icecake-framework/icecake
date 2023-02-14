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
	icecake "github.com/sunraylab/icecake/pkg/framework"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

//go:embed hello1.md
var mymarkdown string

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	doc := dom.GetDocument()

	// 1. demonstrate the use of the go HTML templating package to build page content directly on the front-end.
	var data1 struct{ Name string }
	htmlTemplate := `Hello <strong>{{.Name}}</strong>!`

	data1.Name = "Bob"
	icecake.RenderElemHtml(htmlTemplate, data1, doc.ChildById("ex1a"))

	data1.Name = "Alice"
	icecake.RenderElemHtml(htmlTemplate, data1, doc.ChildById("ex1b"))

	// To see what happend with a wrong html element ID,
	// open the console on the browser side.
	data1.Name = "Carol"
	icecake.RenderElemHtml(htmlTemplate, data1, doc.ChildById("ex1c"))

	// 2. demonstrate how to generate HTML content from a markdown source, directly on the front-side.
	data1.Name = "John"
	icecake.RenderElemMarkdown("### Markdown\nHello **{{.Name}}**", data1, doc.ChildById("ex1d"))

	// Text source is embedded in the compiled wasm code with the //go:embed compiler directive
	var data2 struct{ Brand string }
	data2.Brand = "<span class='brand'>Icecake</span>"
	icecake.RenderElemMarkdown(mymarkdown, data2, doc.ChildById("ex1e"),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
