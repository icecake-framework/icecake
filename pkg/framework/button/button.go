package button

import (
	"log"
	"syscall/js"

	"github.com/sunraylab/icecake/pkg/webclientsdk"
)

type ICButton struct {
	*webclientsdk.Element
}

// ICButtonFromJS is casting a js.Value into ICButton.
func Cast(value js.Value) *ICButton {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &ICButton{}
	ret.Element = webclientsdk.MakeElementFromJS(value)
	return ret
}

func (icb *ICButton) Render() {
	icb.Element.ClassList().SetTokens("button").ToDOM()
	//icb.Element.SetAttribute("role", "button")
	icb.Element.SetTextContent("click here")
	icb.SetEvent()
}

func (icb *ICButton) SetEvent() {

	evt := webclientsdk.NewEventListenerFunc(icb.OnClick)
	icb.Element.AddEventListener("click", evt)
}

func (icb *ICButton) OnClick(event *webclientsdk.Event) {
	log.Println("Button Clicked")

	divtest := icb.OwnerDocument().GetElementById("idtest")
	log.Println(divtest.Attributes().String())
}
