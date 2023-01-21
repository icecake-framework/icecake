// Copyright 2022 by lolorenzo77. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code.
// It's compiled into a '.wasm' file with "GOOS=js GOARCH=wasm go build -o ../webapp/main.wasm"
package main

import (
	"fmt"
	"os"

	"github.com/sunraylab/icecake/pkg/spasdk"
	browser "github.com/sunraylab/icecake/pkg/webclientsdk"
	"github.com/sunraylab/icecake/web/components/button"
)

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

	coll := browser.GetDocument().GetElementsByTagName("ic-button")
	if coll != nil {
		for i := uint(0); i < coll.Length(); i++ {
			e := coll.Item(i)
			icb := button.Cast(e.JSValue())
			icb.Render()
		}
	}

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
	fmt.Println("Go/WASM ended")
}
