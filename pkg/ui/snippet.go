package ui

import (
	"bytes"

	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/event"
	"github.com/sunraylab/icecake/pkg/html"
	"github.com/sunraylab/icecake/pkg/js"
)

type Composer interface {
	html.HtmlComposer
	js.JSValueWrapper
	event.Listener
}

/*****************************************************************************/

type Snippet struct {
	html.HtmlSnippet
	DOM dom.Element

	//MountClasses    *Classes    // classes added to the component during the mounting stage
	//MountAttributes *Attributes // attributes addes to the component during the mounting stage
	//UpdateUI        func(any)   // Optional function called to update UI
}

// AddListeners, by default, call AddListeners for every embedded components into HtmlSnippet
func (_s *Snippet) AddListeners() {
	embedded := _s.Embedded()
	if embedded != nil {
		for _, e := range embedded {
			if l, ok := e.(event.Listener); ok {
				l.AddListeners()
			}
		}
	} else {
		// DEBUG:
		console.Warnf("Snippet.AddListeners for %q is empty", _s.Id())
	}
}

func (_s *Snippet) Wrap(_jsvp js.JSValueProvider) {
	_s.DOM.Wrap(_jsvp)
}

func (_s *Snippet) HTML(_snippet Composer) (_html html.HTMLstring) {
	// render the html element and body, unfolding sub components
	out := new(bytes.Buffer)
	id, err := html.WriteHtmlSnippet(out, _snippet, nil)
	if err == nil {
		_s.Embed(id, _snippet)
		_html = html.HTMLstring(out.String())
	}
	return _html
}
