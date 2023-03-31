package main

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/ui"
)

// This main package contains the web assembly source code for the icecake example.
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	msg1 := &ui.Notify{
		Message: "This is a simple notification.",
	}
	dom.Id("content").RenderSnippet(dom.INSERT_LAST_CHILD, msg1, nil)

	msg2 := &ui.Notify{
		Message: "This is another simple notification.",
	}
	msg2.SetClasses("is-success is-light")
	dom.Id("content").RenderSnippet(dom.INSERT_LAST_CHILD, msg2, nil)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
