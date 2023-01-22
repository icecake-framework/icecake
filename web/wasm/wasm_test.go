package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/sunraylab/icecake/pkg/spasdk"
	browser "github.com/sunraylab/icecake/pkg/webclientsdk"
)

func TestXxx(t *testing.T) {

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
			// e := coll.Item(i)
			// icb := mybutton.Cast(e.JSValue())
			// icb.Render()
		}
	}

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
	fmt.Println("Go/WASM ended")
	t.Fail()

}
