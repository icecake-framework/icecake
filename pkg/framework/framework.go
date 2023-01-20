package framework

import (
	"fmt"
	"syscall/js"

	"github.com/sunraylab/icecake/pkg/webclientsdk"
)

type ICButton struct {
	*webclientsdk.Element
}

// ICButtonFromJS is casting a js.Value into Document.
func ICButtonFromJS(value js.Value) *ICButton {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &ICButton{}
	ret.Element = webclientsdk.ElementFromJS(value)
	return ret
}

func (icb *ICButton) Render() {
	// TODO test element nit nul and the type of the value
	icb.Element.SetClassName("button")
	icb.Element.SetAttribute("role", "button")
	icb.Element.SetTextContent("ok")
}

func RenderComponents() {
	// rechercher

	fmt.Println("Component rendering.")

	coll := webclientsdk.GetDocument().GetElementsByTagName("ic-button")
	if coll != nil {
		for i := uint(0); i < coll.Length(); i++ {
			e := coll.Item(i)
			icbutton := ICButtonFromJS(e.JSValue())
			icbutton.Render()
		}
	}

}
