package mybutton

import (
	"fmt"
	"log"
	"syscall/js"
	"time"

	browser "github.com/sunraylab/icecake/pkg/webclientsdk"
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
	ret.Element = browser.MakeElementFromJS(value)
	return ret
}

func (_btn *MyButton) Render() {
	_btn.Element.ClassList().SetTokens("button").ToDOM()
	_btn.Element.SetTextContent("Test")
	_btn.SetEvent()
}

func (_btn *MyButton) SetEvent() {
	evt := browser.NewEventListenerFunc(_btn.OnClickTest)
	_btn.Element.AddEventListener("click", evt)
}

func (_btn *MyButton) OnClickTest(event *browser.Event) {
	log.Println("Button Test Clicked")
	_btn.Attributes().Set("disabled", "")

	divcontent := _btn.Doc().GetElementById("testcontent")

	attrs := _btn.Doc().GetElementById("idtest").Attributes()
	fmt.Println("attrs.OwnerElement: ", attrs.OwnerElement.TagName())

	str := fmt.Sprintf("1: %s", attrs.String())
	p := _btn.Doc().CreateElement("p").SetInnerHTML(str)
	divcontent.AppendChild(&p.Node)

	attrs.Set("show", "true")
	str = fmt.Sprintf("2: %s", attrs.String())
	p = _btn.Doc().CreateElement("p").SetInnerHTML(str)
	divcontent.AppendChild(&p.Node)

	attrs.Remove("disabled")
	str = fmt.Sprintf("3: %s", attrs.String())
	p = _btn.Doc().CreateElement("p").SetInnerHTML(str)
	divcontent.AppendChild(&p.Node)

	attrs.Set("role", "button2")
	str = fmt.Sprintf("4: %s", attrs.String())
	p = _btn.Doc().CreateElement("p").SetInnerHTML(str)
	divcontent.AppendChild(&p.Node)

	attrs = _btn.Doc().GetElementById("idtest").Attributes()
	str = fmt.Sprintf("TOT: %s", attrs.String())
	p = _btn.Doc().CreateElement("p").SetInnerHTML(str)
	divcontent.AppendChild(&p.Node)

	attrs = _btn.Doc().GetElementById("idtest2").Attributes()
	str = fmt.Sprintf("TEST2: %s", attrs.String())
	p = _btn.Doc().CreateElement("p").SetInnerHTML(str)
	divcontent.AppendChild(&p.Node)

	attrs.Set("class", "title is-2")
	str = fmt.Sprintf("TEST2: %s", attrs.String())
	p = _btn.Doc().CreateElement("p").SetInnerHTML(str)
	divcontent.AppendChild(&p.Node)

	txt := attrs.Get("data-text")
	str = fmt.Sprintf("TEST2: %s", txt)
	p = _btn.Doc().CreateElement("p").SetInnerHTML(str)
	divcontent.AppendChild(&p.Node)

	attrs = _btn.Doc().GetElementById("idtest3").Attributes()
	str = fmt.Sprintf("TEST3: %s", attrs.String())
	p = _btn.Doc().CreateElement("p").SetInnerHTML(str)
	divcontent.AppendChild(&p.Node)

	go func() {
		time.Sleep(time.Second * 3)
		attrs.Toggle("hidden", "show")

		time.Sleep(time.Second * 3)
		attrs.Toggle("hidden", "show")

		time.Sleep(time.Second * 3)
		attrs.Toggle("hidden", "")

		time.Sleep(time.Second * 3)
		attrs.Toggle("hidden", "")

		time.Sleep(time.Second * 3)
		attrs.Toggle("hidden", "")
	}()
}
