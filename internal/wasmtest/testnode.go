package main

import (
	"fmt"

	ick "github.com/sunraylab/icecake/pkg/icecake"
)

func TestNode() {

	app := ick.NewWebApp()

	e := app.CreateElement("DIV")
	e.SetId("tstnode1")
	if e.IsInDOM() {
		fmt.Println("1: want false, get true")
	}

	econtainer := app.ChildById("test-container")
	econtainer.AppendChild(&e.Node)
	if !e.IsInDOM() {
		fmt.Println("2: want true, get false")
	}

	if new(ick.Element).IsInDOM() {
		fmt.Println("3: want false, get true")
	}

}
