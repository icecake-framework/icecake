// Copyright 2022 by lolorenzo77. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code.
// It's compiled into a '.wasm' file with "GOOS=js GOARCH=wasm go build -o ../webapp/main.wasm"
package main

import (
	"fmt"
	"os"

	_ "embed"

	// "github.com/sunraylab/icecake/pkg/dom"
	// icecake "github.com/sunraylab/icecake/pkg/framework"
	"github.com/sunraylab/icecake/pkg/spasdk"
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

	// doc := dom.GetDocument()

	// data := struct{ Name string }{
	// 	Name: "Bob",
	// }
	// icecake.RenderElement(`example0 : Hello <strong>{{.Name}}</strong>!`, data, doc.ChildById("example0"))

	// icecake.RenderElement(`example1 : <ic-ex1 />`, data, doc.ChildById("example1"))

	// icecake.RenderElement(`example2 : <ic-ex2 />`, data, doc.ChildById("example2"))

	// icecake.RenderElement(`example3 : <ic-ex3 />`, data, doc.ChildById("example3"))

	// icecake.RenderElement(`example4 : <ic-ex4 />`, data, doc.ChildById("example4"))

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
