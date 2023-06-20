package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
)

// This main package contains the web assembly source code for the icecake example.
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	e := dom.CreateElement("div").SetId("testid")
	e.InsertText(dom.INSERT_BODY, "Hello world")
	dom.Id("content").InsertElement(dom.INSERT_AFTER_ME, e)

	print("tagname", e.TagName())
	print("id", e.Id())

	e.SetClasses("myclass1 myclass2 myclass3 myclass6")
	e.SetClasses("")
	e.SwitchClasses("myclass3 myclass1", "myclass2 myclass4")
	print("classes", e.Classes())
	print("has class myclass1", sprintBool(e.HasClass("myclass1")))
	print("has class myclass2", sprintBool(e.HasClass("myclass2")))

	e.RemoveClasses("myclass2")
	print("remove classes myclass2", e.Classes())

	e.ResetClasses("myclass5")
	print("reset myclass5", e.Classes())

	e.ResetClasses("")
	print("reset empty class", e.Classes())

	e.SetStyle("color:red;")
	print("style red", "color")

	e.SetTabIndex(1)
	print("tabindex", e.TabIndex())

	e.SetDisabled(true)
	print("Disabled", e.IsDisabled())

	e.CreateAttribute("astring1", "string")
	e.CreateAttribute("astring2", "string")
	e.CreateAttribute("anum", 123)
	e.CreateAttribute("abool", false)
	e.SetAttribute("anum", 456)
	e.CreateAttribute("anum", 789)
	e.RemoveAttribute("astring1")
	e.ToggleAttribute("abool2")
	print("attributes", e.Attributes())

	e.ToggleAttribute("abool2")
	print("attributes", e.Attributes())

	print("access key", "H")
	e.SetAccessKey("H")

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}

func print(_title string, _value any) {
	v := fmt.Sprintf("%v", _value)
	dom.Id("content").InsertRawHTML(dom.INSERT_LAST_CHILD, *html.NewString(_title + ": <strong>" + v + "</strong><br/>"))
}

func sprintBool(f bool) string {
	if f {
		return "true"
	}
	return "false"
}
