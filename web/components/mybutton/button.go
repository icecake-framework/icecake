package mybutton

import (
	"syscall/js"

	browser "github.com/sunraylab/icecake/pkg/dom"
)

type MyButton struct {
	*browser.Element
}

// Cast is casting a js.Value into MyButton.
func Cast(value js.Value) *MyButton {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &MyButton{}
	ret.Element = browser.CastElement(value)
	return ret
}

func (_btn *MyButton) Render() {
	_btn.Element.ClassList().Set("button")
	_btn.Element.SetTextContent("Test")
}
