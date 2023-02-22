package ick

import (
	"bytes"
	"errors"
	"fmt"
	"syscall/js"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/yuin/goldmark"
)

// InsertComponent
func InsertComponent(_newcmp any, _id string) (_err error) {

	// get the element
	elem := GetElementById(_id)
	if !elem.IsDefined() {
		return fmt.Errorf("id %q not found in the DOM or not an <Element>", _id)
	}

	var cmpelem *Element
	switch compounder := _newcmp.(type) {
	case HtmlCompounder:

		// name the component
		name := elem.TagName() + "/" + elem.Id()

		// unfold and render html for a compounder
		cmpelem, _err = CreateCompoundElement(compounder)
		if _err == nil {
			html, _ := unfoldComponents(name, compounder.Template(), compounder, 0)
			cmpelem.SetInnerHTML(html)

			//elem.InsertAdjacentHTML(WI_INSIDEFIRST, html)
			elem.PrependNodes(&cmpelem.Node)

		} else {
			ConsoleWarnf(_err.Error())
			return _err
		}

		// wrap this new html element to th _cmp
		switch wrapper := _newcmp.(type) {
		case JSWrapper:
			if typ := wrapper.JSValue().Type(); typ == js.TypeNull || typ == js.TypeUndefined {
				// fmt.Println("_newcmp is a Element")
				wrapper.Wrap(cmpelem.JSValue())
			} else {
				return fmt.Errorf("compounder %q has already been inserted", _id)
			}
		default:
			return fmt.Errorf("compounder %q is not an Element", _id)
		}

		// TODO: add style

		// addlisteners
		switch listener := _newcmp.(type) {
		case HtmlListener:
			// fmt.Println("_newcmp is a listener")
			listener.AddListeners()
		}
	default:
		return errors.New("InsertComponent failed: _newcmp is not a compounder")
	}

	return nil
}

/*****************************************************************************
* ICElement
******************************************************************************/

// ICElement is an extension of the HTMLElement
type ICElement struct {
	Element
}

func CastICElement(_value js.Value) *ICElement {
	if _value.Type() != js.TypeObject {
		ConsoleErrorf("casting ICElement failed")
		return new(ICElement)
	}
	cast := new(ICElement)
	cast.Wrap(_value)
	return cast
}

// GetICElementById returns an ICElement corresponding to the _id if it exists into the DOM,
// otherwhise returns an undefined ICElement.
func GetElementById(_id string) *ICElement {
	_id = helper.Normalize(_id)
	jse := GetDocument().JSValue().Call("getElementById", _id)
	if etyp := jse.Type(); etyp != js.TypeNull && etyp != js.TypeUndefined {
		elem := new(ICElement)
		elem.Wrap(jse)
		return elem
	}
	ConsoleWarnf("GetElementById failed: %q not found, or not a <Element>", _id)
	return new(ICElement)
}

// SetInnerValue set the innext text of the element with a formated value.
// The format string follow the fmt rules: https://pkg.go.dev/fmt#hdr-Printing
func (_elem *ICElement) RenderValue(format string, _value ...any) {
	if !_elem.IsDefined() {
		return
	}
	text := fmt.Sprintf(format, _value...)
	_elem.SetInnerText(text)
}

// RenderHtml set inner HTML with the htmlTemplate executed with the _data and unfolding components if any
func (_elem *ICElement) RenderHtml(_unsafeHtmlTemplate string, _data any) {
	if !_elem.IsDefined() {
		return
	}
	name := _elem.TagName() + "/" + _elem.Id()
	html, _ := unfoldComponents(name, _unsafeHtmlTemplate, _data, 0)
	_elem.SetInnerHTML(html)
}

// RenderMarkdown process _mdtxt markdown source file and convert it to an HTML string,
// then use it as an HTML template to render it with data and components.
//
// Returns an error if the markdown processor fails.
func (_elem *ICElement) RenderMarkdown(_mdtxt string, _data any, _options ...goldmark.Option) error {
	if !_elem.IsDefined() {
		return nil
	}
	name := _elem.TagName() + ":" + _elem.Id()
	md := goldmark.New(_options...)
	var buf bytes.Buffer
	err := md.Convert([]byte(_mdtxt), &buf)
	if err != nil {
		ConsoleWarnf("RenderMarkdown has error: %s", err.Error())
		return err
	}

	html, _ := unfoldComponents(name, buf.String(), _data, 0)
	_elem.SetInnerHTML(html)
	return nil
}

// RenderNamedValue look recursively for any _elem children having the "data-ic-namedvalue" token matching _name
// and render inner text with the _value
func (_elem *ICElement) RenderChildrenValue(_name string, _format string, _value ...any) {
	if !_elem.IsDefined() {
		return
	}
	_name = helper.Normalize(_name)
	text := fmt.Sprintf(_format, _value...)

	children := _elem.FilteredChildren(NT_ELEMENT, 99, func(_node *Node) bool {
		dataset := CastICElement(_node.JSValue()).Attributes().Dataset()
		for i := 0; i < dataset.Count(); i++ {
			if dataset.At(i).Name() == "data-ic-namedvalue" && dataset.At(i).Value() == _name {
				return true
			}
		}
		return false
	})

	for _, node := range children {
		CastICElement(node.JSValue()).RenderValue(text)
	}
}
