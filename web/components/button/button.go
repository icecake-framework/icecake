package button

import (
	"fmt"
	"log"
	"syscall/js"

	browser "github.com/sunraylab/icecake/pkg/webclientsdk"
)

type ICButton struct {
	*browser.Element
	step int
}

// ICButtonFromJS is casting a js.Value into ICButton.
func Cast(value js.Value) *ICButton {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &ICButton{}
	ret.Element = browser.MakeElementFromJS(value)
	return ret
}

func (icb *ICButton) Render() {
	icb.Element.ClassList().SetTokens("button").ToDOM()
	//icb.Element.SetAttribute("role", "button")
	icb.Element.SetTextContent("click here")
	icb.SetEvent()
}

func (icb *ICButton) SetEvent() {

	evt := browser.NewEventListenerFunc(icb.OnClick)
	icb.Element.AddEventListener("click", evt)
}

func (icb *ICButton) OnClick(event *browser.Event) {
	log.Println("Button Clicked")

	attr := icb.OwnerDocument().GetElementById("idtest").Attributes()
	divcontent := icb.OwnerDocument().GetElementById("testcontent")

	e := browser.GetDocument().CreateElement("p")
	str := fmt.Sprintf("step %v: %s", icb.step, attr.String())
	e.SetInnerHTML(str)
	log.Println(str)
	divcontent.AppendChild(&e.Node)
	icb.step++

	switch icb.step {
	case 1:
		attr.Remove("disabled")
	case 2:
		attr.Set("role", "button2")
	case 3:
		attr.Set("show", "true")
	}
}
