package ickui

import (
	"github.com/icecake-framework/icecake/pkg/clock"
	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/js"
)

type ICKDelete struct {
	ick.ICKDelete
	dom.UI

	// The TargetID will be automatically removed after the clock Timeout duration if not zero.
	// The timer starts when the delete button is rendered (call to addlisteners).
	clock.Clock

	// OnDelete, if it is set, it's called when the deletion occurs and after the targetId has been removed.
	OnDelete func(*ICKDelete)
}

// Ensure Button implements UIComposer interface
var _ dom.UIComposer = (*ICKDelete)(nil)

func Delete(targetid string) *ICKDelete {
	del := new(ICKDelete)
	del.ICKDelete = *ick.Delete(targetid)
	return del
}

// Wrap implements the JSValueWrapper to enable wrapping of a dom.Element usually
// to wrap embedded component instantiated during unfolding an html string.
// Does not need to be overloaded by the component embedding UISnippet.
func (del *ICKDelete) Wrap(jsvp js.JSValueProvider) {
	if del.UI.Wrap(jsvp); !del.DOM.IsInDOM() {
		console.Errorf("ICKDelete.Wrap: failed")
		return
	}
	t, has := del.DOM.Attribute("data-targetid")
	console.Warnf("ICKDelete.Wrap: data-targetid=%q", t)
	if has {
		del.TargetId = t
	}
}

func (del *ICKDelete) AddListeners() {
	console.Warnf("ICKDelete.AddListeners: targetid=%q", del.TargetId)
	if del.TargetId != "" {
		del.DOM.AddMouseEvent(event.MOUSE_ONCLICK, func(*event.MouseEvent, *dom.Element) {
			del.RemoveTarget()
		})
		del.Clock.Start(del.RemoveTarget)
	} else {
		console.Warnf("ICKDelete.AddListeners: missing TargetID")
	}
}

func (del *ICKDelete) RemoveListeners() {
	del.DOM.RemoveListeners()
}

// RemoveTarget stops the timer and the ticker and removes the TargetID from the DOM
func (del *ICKDelete) RemoveTarget() {
	del.Stop()
	dom.Id(del.TargetId).Remove()
	if del.OnDelete != nil {
		del.OnDelete(del)
	}
}
