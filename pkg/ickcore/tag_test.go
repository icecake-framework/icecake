package ickcore

import (
	"bytes"
	"fmt"
	"io"
)

func ExampleTag() {

	render := func(w io.Writer, tag Tag, s string) {
		tag.RenderOpening(w)
		io.WriteString(w, s)
		tag.RenderClosing(w)
	}

	// create a new div tag
	tag := NewTag("div", `class="example dark"`)

	out := new(bytes.Buffer)
	render(out, *tag, "example1")
	fmt.Println(out.String())

	// update div tag
	active := true
	tag.SwitchClass("dark", "light").SetClassIf(active, "is-active")

	out.Reset()
	render(out, *tag, "example2")
	fmt.Println(out.String())

	// output:
	// <div class="example dark">example1</div>
	// <div class="example light is-active">example2</div>
}
