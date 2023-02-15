package icecake

import (
	"bytes"
	"fmt"
	"syscall/js"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/yuin/goldmark"
)

func DocumentBody() *ICElement {
	// TODO: error handling
	return &ICElement{HTMLElement: dom.GetDocument().Body()}
}

// func DocChildById(_id string) *Element {
// 	child := dom.GetDocument().ChildById(_id)
// 	if child == nil || child.NodeType() != dom.NT_ELEMENT {
// 		return nil
// 	}
// 	return &Element{HTMLElement: dom.NewHTMLElementFromJS(child.JSValue())}
// }

/*****************************************************************************/

// // scan the Documents's Body to look for components and to render them.
// func RenderBody() {
// 	body := ICElement{HTMLElement: dom.Doc().Body()}
// 	body.RenderHtml(body.InnerHTML(), GData)
// }

// func RenderValueElem(_elem *ICElement, _value any) {
// 	if _elem == nil {
// 		dom.ConsoleError("RenderValue failed: Unable to render nil element")
// 		return
// 	}
// 	_elem.RenderValue(_value)
// }

// func RenderValueElem(_dome *dom.Element, _value any) {
// 	if _dome == nil {
// 		dom.ConsoleError("RenderValue failed: Unable to render nil element")
// 		return
// 	}
// 	elem := NewElementFromJS(_dome.JSValue())
// 	elem.RenderValue(_value)
// }

// func RenderValueById(_id string, _value any) {
// 	elem := Doc().ChildById(_id)
// 	if elem == nil || elem.NodeType() != dom.NT_ELEMENT {
// 		return
// 	}
// 	&ICElement{HTMLElement: dom.CastHTMLElement(child.JSValue())}

// 	elem.RenderValue(_value)
// }

// func RenderHtmlElem(_elem *Element, _unsafeHtmlTemplate string, _data any) {
// 	if _elem == nil {
// 		dom.ConsoleError("RenderElement failed: Unable to render nil element")
// 		return
// 	}
// 	_elem.RenderHtml(_unsafeHtmlTemplate, _data)
// }

// func RenderMarkdownElem(_elem *Element, _mdtxt string, _data any, _options ...goldmark.Option) {
// 	if _elem == nil {
// 		dom.ConsoleError("RenderMarkdown failed: Unable to render nil element")
// 		return
// 	}
// 	_elem.RenderMarkdown(_mdtxt, _data, _options...)
// }

/*****************************************************************************
* ICElement
******************************************************************************/

// ICElement is an extension of the dom.HTMLElement
type ICElement struct {
	*dom.HTMLElement
}

func CastICElement(_value js.Value) *ICElement {
	return &ICElement{HTMLElement: dom.CastHTMLElement(_value)}
}

// GetICElementById returns an ICElement corresponding to the _id if it exists into the DOM,
// otherwhise returns an undefined ICElement.
func GetElementById(_id string) *ICElement {
	child := dom.GetDocument().ChildById(_id)
	if child.NodeType() != dom.NT_ELEMENT {
		return &ICElement{}
	}
	return CastICElement(child.JSValue())
}

// SetInnerValue set the innext text of the element with a formated value.
// The format string follow the fmt rules: https://pkg.go.dev/fmt#hdr-Printing
func (_elem *ICElement) RenderValue(format string, _value ...any) {
	text := fmt.Sprintf(format, _value...)
	_elem.SetInnerText(text)
}

// RenderHtml set inner HTML with the htmlTemplate executed with the _data and unfolding components if any
func (_elem *ICElement) RenderHtml(_unsafeHtmlTemplate string, _data any) {
	name := _elem.TagName() + "/" + _elem.Id()
	html, _ := unfoldComponents(name, _unsafeHtmlTemplate, _data, 0)
	_elem.SetInnerHTML(html)
}

// RenderMarkdown process _mdtxt markdown source file and convert it to an HTML string,
// then use it as an HTML template to render it with data and components.
//
// Returns an error if the markdown processor fails.
func (_elem *ICElement) RenderMarkdown(_mdtxt string, _data any, _options ...goldmark.Option) error {
	name := _elem.TagName() + "/" + _elem.Id()
	md := goldmark.New(_options...)
	var buf bytes.Buffer
	err := md.Convert([]byte(_mdtxt), &buf)
	if err != nil {
		dom.ConsoleWarn("RenderMarkdown failed: %s", err.Error())
		return err
	}

	html, _ := unfoldComponents(name, buf.String(), _data, 0)
	_elem.SetInnerHTML(html)
	return nil
}

// RenderNamedValue look recursively for any _elem children having the "data-ic-namedvalue" token matching _name
// and render inner text with the _value
func (_elem *ICElement) RenderChildrenValue(_name string, _format string, _value ...any) {
	_name = helper.Normalize(_name)
	text := fmt.Sprintf(_format, _value...)

	//fmt.Println("text: ", text)

	children := _elem.FilteredChildren(dom.NT_ELEMENT, 1, func(_node *dom.Node) bool {
		elem := CastICElement(_node.JSValue())
		fmt.Println(" element:", elem.TagName())
		attrs := elem.Attributes()
		fmt.Println(" attributes:", attrs.Count(), attrs.String())
		dataset := attrs.Dataset()
		fmt.Println(" dataset:", dataset.Count(), dataset.String())
		for i := 0; i < dataset.Count(); i++ {
			if dataset.At(i).Name() == "data-ic-namedvalue" && dataset.At(i).Value() == _name {
				return true
			}
		}
		return false
	})

	fmt.Println("top filter ended")
	for _, node := range children {
		CastICElement(node.JSValue()).RenderValue(_name, text, children)
	}
}

/*****************************************************************************
* ICButton
******************************************************************************/

// ICButton is an extension of the dom.HTMLElement
type ICButton struct {
	*dom.HTMLButton
}

func CastICButton(_value js.Value) *ICButton {
	return &ICButton{HTMLButton: dom.CastHTMLButton(_value)}
}

// GetButtonById returns an ICElement corresponding to the _id if it exists into the DOM,
// otherwhise returns an undefined ICElement.
func GetButtonById(_id string) *ICButton {
	child := dom.GetDocument().ChildById(_id)
	return CastICButton(child.JSValue())
}
