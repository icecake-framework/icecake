package dom

import (
	"bytes"

	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/html"
	"github.com/sunraylab/icecake/pkg/js"
)

type Listener interface {
	AddListeners()
	RemoveListeners()
}

type UIComposer interface {
	html.HTMLComposer
	js.JSValueWrapper
	Listener
}

/*****************************************************************************/

// UISnippet combines an htmlSnippet allowing html rendering of ick-name tags in different ways, and
// a wrapped dom.Element allowing event listening and other direct DOM interactions.
type UISnippet struct {
	html.HTMLSnippet
	DOM Element
}

// AddListeners, by default, listen every embedded components into HtmlSnippet
// AddListeners, by default, call AddListeners for every embedded components into HtmlSnippet
func (_s *UISnippet) AddListeners() {
	console.Warnf("Snippet.AddListeners for %q", _s.Id())
	embedded := _s.Embedded()
	if embedded != nil {
		for _, e := range embedded {
			if l, ok := e.(Listener); ok {
				l.AddListeners()
			}
		}
	} else {
		// DEBUG:
		console.Warnf("Snippet.AddListeners for %q is empty", _s.Id())
	}
}

// Wrap implements the JSValueWrapper to enable wrapping of a dom.Element usually
// to wrap embedded component instantiated during unfolding an html string.
func (_s *UISnippet) Wrap(_jsvp js.JSValueProvider) {
	_s.DOM.Wrap(_jsvp)
}

func (_s *UISnippet) RemoveListeners() {
	_s.DOM.RemoveListeners()
}

// Html builds and unfolds the UIcomposer and returns its cStringng HTML string
func (_s *UISnippet) HTML(_snippet UIComposer) (_html html.String) {
	// render the html element and body, unfolding sub components
	out := new(bytes.Buffer)
	id, err := html.WriteHtmlSnippet(out, _snippet, nil)
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
