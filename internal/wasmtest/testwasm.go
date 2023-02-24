package main

import (
	"fmt"
)

func main() {
	c := make(chan struct{})
	fmt.Println("Go/WASM loaded, tests started.")

	fmt.Println("starting tests")

	// TestAttributes()
	TestNode()

	fmt.Println("starting ended")

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
