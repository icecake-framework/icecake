package dom

import (
	"bytes"

	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/html"
	"github.com/sunraylab/icecake/pkg/js"
)

// type Listener interface {
// 	AddListeners()
// 	RemoveListeners()
// }

type UIComposer interface {
	html.HTMLComposer
	js.JSValueWrapper
	Mount()
	AddListeners()
	RemoveListeners()
	UnMount()
}

/*****************************************************************************/

// UISnippet combines an htmlSnippet allowing html rendering of ick-name tags in different ways, and
// a wrapped dom.Element allowing event listening and other direct DOM interactions.
type UISnippet struct {
	html.HTMLSnippet
	DOM Element
}

// Mount does nothing by default. Can be implemented by the component embedding UISnippet.
func (_s *UISnippet) Mount() {}

// UnMount does nothing by default. Can be implemented by the component embedding UISnippet.
func (_s *UISnippet) UnMount() {}

// AddListeners does nothing by default. Can be implemented by the component embedding UISnippet.
func (_s *UISnippet) AddListeners() {}

// Wrap implements the JSValueWrapper to enable wrapping of a dom.Element usually
// to wrap embedded component instantiated during unfolding an html string.
// Does not need to be overloaded by the component embedding UISnippet.
func (_s *UISnippet) Wrap(_jsvp js.JSValueProvider) {
	if _s.DOM.Value().Truthy() {
		console.Warnf("wrapping snippet %q to the already wrapped element %q", _s.Id(), _s.DOM.Id())
	}
	_s.DOM.JSValue = _jsvp.Value()
}

// RemoveListeners remove all event handlers attached to this UISnippet Element.
// If RemoveListeners is implemented by the component embedding UISnippet then the UISnippet one should be called.
// Usually RemoveListeners does not need to be overloaded because every listeners added to the Element are automatically removed.
func (_s *UISnippet) RemoveListeners() {
	_s.DOM.RemoveListeners()
}

// Html builds and unfolds the UIcomposer and returns its cStringng HTML string
func (_s *UISnippet) HTML(_snippet UIComposer) (_html html.String) {
	// render the html element and body, unfolding sub components
	out := new(bytes.Buffer)
	id, err := html.WriteHTMLSnippet(out, _snippet, nil)
	if err == nil {
		_s.Embed(id, _snippet)
		_html = html.String(out.String())
	}
	return _html
}

// InsertSnippet
func (_parent *UISnippet) InsertSnippet(_where INSERT_WHERE, _snippet any, _data *html.DataState) (_id string, _err error) {
	if _parent == nil || !_parent.DOM.IsDefined() {
		return "", console.Errorf("Snippet:InsertSnippetfailed on undefined _parent")
	}

	_parent.DOM.InsertSnippet(_where, _snippet, _data)
	_parent.AddListeners()
	return _id, nil
}

func (_s *UISnippet) SetDisabled(_disable bool) {
	_s.HTMLSnippet.SetDisabled(_disable)
	if _s.DOM.IsInDOM() {
		_s.DOM.SetDisabled(_disable)
	}
}

/******************************************************************************
* Private Area
*******************************************************************************/

// mountDeepSnippet wraps _elem to the _snippet add its listeners and call the customized Mount function.
// mountDeepSnippet is called recursively for every embedded components of the _snippet.
func mountDeepSnippet(_snippet UIComposer, _elem *Element) (_err error) {
	//DEBUG: console.Warnf("mouting %s(%s)", _snippet.Id(), reflect.TypeOf(_snippet).String())
	_snippet.Wrap(_elem)
	_snippet.AddListeners()
	_snippet.Mount()

	if embedded := _snippet.Embedded(); embedded != nil {
		// DEBUG: console.Warnf("scanning %+v", embedded)
		for subid, sub := range embedded {
			// look everywhere in the DOM
			if sube := Id(subid); sube != nil {
				if cmp, ok := sub.(UIComposer); ok {
					// DEBUG: console.Warnf("wrapping %+v", w)
					_err = mountDeepSnippet(cmp, sube)
				}
			}
		}
	}
	return _err
}

// FIXME: must call unmount somewhere
// unmountDeepSnippet remove listeners anc all Unmount recusrively for every embedded components
func unmountDeepSnippet(_snippet UIComposer) {
	_snippet.RemoveListeners()
	_snippet.UnMount()

	if embedded := _snippet.Embedded(); embedded != nil {
		// DEBUG: console.Warnf("scanning %+v", embedded)
		for _, sub := range embedded {
			if cmp, ok := sub.(UIComposer); ok {
				// DEBUG: console.Warnf("wrapping %+v", w)
				unmountDeepSnippet(cmp)
			}
		}
	}
}
