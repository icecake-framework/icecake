package bulmaui

import (
	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
)

type Delete struct {
	bulma.Delete
	DOM dom.Element
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
func (del *Delete) RemoveTarget() {
	del.Stop()
	dom.Id(del.TargetID).Remove()
	if del.OnDelete != nil {
		del.OnDelete(&del.Delete)
	}
}
