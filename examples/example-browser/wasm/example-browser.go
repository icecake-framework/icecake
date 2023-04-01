package main

import "fmt"

// This main package contains the web assembly source code for the icecake example.
// It's compiled into a '.wasm' file with the build_ex1 task.
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
