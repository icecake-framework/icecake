package ui

import (
	"bytes"
	"fmt"

	"github.com/sunraylab/icecake/pkg/errors"
	ick "github.com/sunraylab/icecake/pkg/icecake2"
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
	wick.Element

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

// RenderTemplate set inner HTML with the htmlTemplate executed with the _data and unfolding components if any
// The element must be in the DOM to
func RenderSnippetBody(_elem *wick.Element, _body ick.HTML, _data *ick.DataState) (_err error) {
	if !_elem.IsDefined() || !_elem.IsInDOM() {
		return fmt.Errorf("unable to render Html on nil element or for an element not into the DOM")
	}

	out := new(bytes.Buffer)
	_err = ick.RenderHtmlSnippetBody(out, _body, _data)
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
		return errors.ConsoleErrorf("RenderComponent: failed on undefined element")
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
