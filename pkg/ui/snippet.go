package ui

import (
	"bytes"
	"fmt"

	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/ick"
	wick "github.com/sunraylab/icecake/pkg/wicecake"
)

type Listener interface {
	AddListeners()
}

type Composer interface {
	ick.HtmlComposer
	Listener
	Wrap(wick.JSValueProvider)
	UpdateUI()
}

/*****************************************************************************/

type Snippet struct {
	ick.HtmlSnippet
	DOM wick.Element

	//MountClasses    *Classes    // classes added to the component during the mounting stage
	//MountAttributes *Attributes // attributes addes to the component during the mounting stage
	//UpdateUI        func(any)   // Optional function called to update UI
}

// func (c *Snippet) UpdateUI() {
// 	// DEBUG:
// 	fmt.Printf("UIComponent.UpdateUI does nothing by default\n")
// }

func (c *Snippet) AddListeners() {
	// DEBUG:
	fmt.Printf("UIComponent.Listeners is empty\n")
}

// RenderHtml unfolding components if any
// The element must be in the DOM to
func RenderHtml(_elem *wick.Element, _body ick.HTML, _data *ick.DataState) (_err error) {
	if !_elem.IsDefined() || !_elem.IsInDOM() {
		return fmt.Errorf("unable to render Html on nil element or for an element not into the DOM")
	}

	out := new(bytes.Buffer)
	_err = ick.RenderHtmlBody(out, _body, _data)
	if _err == nil {
		_elem.SetInnerHTML(out.String())
		// TODO: loop over embedded component to add listeners

		// TODO: showUnfoldedComponents(unfoldedCmps)
	}
	return _err
}

// RenderComponent
func RenderSnippet(_elem *wick.Element, _cmp any, _data *ick.DataState) (_err error) {
	if !_elem.IsDefined() {
		return console.Errorf("RenderComponent: failed on undefined element")
	}

	out := new(bytes.Buffer)
	_err = ick.RenderHtmlSnippet(out, _cmp, _data)
	if _err == nil {
		_elem.SetInnerHTML(out.String())
		// TODO: loop over embedded component
		if l, ok := _cmp.(Listener); ok {
			l.AddListeners()
		}
	}
	return nil
}
