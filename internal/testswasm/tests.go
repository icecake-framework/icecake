package main

import (
	"flag"
	"fmt"
	"testing"
)

func main() {
	c := make(chan struct{})
	fmt.Println("Go/WASM loaded, tests started.")

	testing.Init()
	flag.Parse()
	flag.Set("test.v", "true")

	testing.Main(nil,
		[]testing.InternalTest{
			{"Tests JSValue", TestJSValue},
			{"Tests Node", TestNode},
			{"Tests browser", TestBrowser},
		}, nil, nil)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
