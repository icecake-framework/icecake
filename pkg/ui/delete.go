package ui

import (
	"github.com/icecake-framework/icecake/pkg/clock"
	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/html"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	html.RegisterComposer("ick-delete", &Delete{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type Delete struct {
	dom.UISnippet

	// the element id to remove from the DOM
	TargetID string

	// The TargetID will be automatically removed after the clock Timeout duration if not zero.
	// The timer starts when the delete button is rendered (call to addlisteners).
	clock.Clock

	// OnDelete, if set, is called when the deletion occurs and after the targetId has been removed
	OnDelete func(*Delete)
}

func (_del *Delete) Template(_data *html.DataState) (_t html.SnippetTemplate) {
	_t.TagName = "Button"
	_t.Attributes = `class="delete" aria-label='delete'`
	return
}

func (_del *Delete) AddListeners() {
	if _del.TargetID != "" {
		_del.DOM.AddMouseEvent(event.MOUSE_ONCLICK, func(*event.MouseEvent, *dom.Element) {
			_del.RemoveTarget()
		})
		_del.Clock.Start(_del.RemoveTarget)

	} else {
		console.Warnf("missing TargetID for the ic-delete component")
	}
}

// RemoveTarget stops the timer and the ticker and removes the TargetID from the DOM
func (_del *Delete) RemoveTarget() {
	_del.Stop()
	dom.Id(_del.TargetID).Remove()
	if _del.OnDelete != nil {
		_del.OnDelete(_del)
	}
}
