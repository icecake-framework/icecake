package ui

import (
	"github.com/sunraylab/icecake/pkg/clock"
	"github.com/sunraylab/icecake/pkg/errors"
	ick "github.com/sunraylab/icecake/pkg/icecake"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	ick.App.RegisterComponent("ick-delete", Delete{}, "")
}

type Delete struct {
	ick.UIComponent

	// the element id to remove from the DOM
	TargetID string

	// The TargetID will be automatically removed after the clock Timeout duration, if not zero.
	// The timer starts when the delete button is rendered (call to addlisteners).
	clock.Clock
}

func (_del *Delete) Container() (_tagname string, _classes string, _attrs string) {
	return "button", "ick-delete delete", "aria-label='delete'"
}

func (_del *Delete) AddListeners() {
	if _del.TargetID != "" {
		_del.AddMouseEvent(ick.MOUSE_ONCLICK, func(*ick.MouseEvent, *ick.Element) {
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
	_del.UIComponent.Remove()
	ick.App.ChildById(_del.TargetID).Remove()
}
