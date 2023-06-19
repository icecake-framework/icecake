package html

import (
	"bytes"
	"fmt"
	"io"
)

func ExampleTag() {

	tag := NewTag("div", ParseAttributes(`class="example"`))

	out := new(bytes.Buffer)
	tag.RenderOpening(out)
	io.WriteString(out, "content")
	tag.RenderClosing(out)
	fmt.Println(out.String())

	highlight := true
	tag.Attributes().SwitchClass("example", "example2").SetName("ex").AddClassesIf(highlight, "lightcolor")
	out.Reset()
	tag.RenderOpening(out)
	io.WriteString(out, "content")
	tag.RenderClosing(out)
	fmt.Println(out.String())

	// output:
	// <div class="example">content</div>
	// <div name="ex" class="example2 lightcolor">content</div>
}
