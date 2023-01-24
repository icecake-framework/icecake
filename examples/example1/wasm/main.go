// Copyright 2022 by lolorenzo77. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code.
// It's compiled into a '.wasm' file with "GOOS=js GOARCH=wasm go build -o ../webapp/main.wasm"
package main

import (
	"fmt"
	"os"

	_ "embed"

	"bytes"

	"github.com/sunraylab/icecake/pkg/spasdk"
	browser "github.com/sunraylab/icecake/pkg/webclientsdk"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

//go:embed hello.md
var hellotxt string

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// Check Server Health
	if !spasdk.ApiGetHealth() {
		fmt.Println("Go/WASM stopped")
		os.Exit(1)
	}

	// convert markdown content to HTML
	md := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(hellotxt), &buf); err != nil {
		panic(err)
	}
	doc := browser.GetDocument()
	p := doc.CreateElement("p").SetInnerHTML(buf.String())
	doc.Body().AppendChild(&p.Node)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
