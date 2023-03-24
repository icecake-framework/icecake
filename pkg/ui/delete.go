package ui

import (
	"github.com/sunraylab/icecake/pkg/clock"
	"github.com/sunraylab/icecake/pkg/errors"
	"github.com/sunraylab/icecake/pkg/ick"
	wick "github.com/sunraylab/icecake/pkg/wicecake"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	ick.RegisterComposer("ick-delete", Delete{})
}

type Delete struct {
	Snippet

	// the element id to remove from the DOM
	TargetID string

	// The TargetID will be automatically removed after the clock Timeout duration, if not zero.
	// The timer starts when the delete button is rendered (call to addlisteners).
	clock.Clock
}

func (_del *Delete) Template(_data *ick.DataState) (_t ick.SnippetTemplate) {
	_t.TagName = "Button"
	_t.Attributes = `class="delete" aria-label='delete'`
	return
}

func (_del *Delete) Listeners() {
	if _del.TargetID != "" {
		_del.DOM.AddMouseEvent(wick.MOUSE_ONCLICK, func(*wick.MouseEvent, *wick.Element) {
			_del.Remove()
		})
		_del.Clock.Start(_del.Remove)

	} else {
		errors.ConsoleWarnf("missing TargetID for the ic-delete component")
	}
}

// Remove stops the timer and tocker, removes the delete component from the DOM
// and removes also the TargetID from the DOM
func (_del *Delete) Remove() {
	_del.Stop()
	_del.DOM.Remove()
	wick.App.ChildById(_del.TargetID).Remove()
}
