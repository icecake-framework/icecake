// Copyright 2022 by lolorenzo77. All rights reserved.
// Use of this source code is governed by MIT licence that can be found in the LICENSE file.

// this main package contains the web assembly source code.
// It's compiled into a '.wasm' file with "GOOS=js GOARCH=wasm go build -o ../webapp/main.wasm"
package main

import (
	"fmt"
	"os"

	"github.com/icecake-framework/icecake/pkg/icksdk"
)

// the main func is required by the wasm GO builder
// outputs will appears in the console of the browser
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// Check Server Health
	if !icksdk.ApiGetHealth() {
		fmt.Println("Go/WASM stopped")
		os.Exit(1)
	}

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
