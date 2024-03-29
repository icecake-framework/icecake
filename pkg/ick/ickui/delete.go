package ickui

import (
	"errors"

	"github.com/icecake-framework/icecake/pkg/clock"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/js"
	"github.com/lolorenzo777/verbose"
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

func Delete(id string, targetid string) *ICKDelete {
	del := new(ICKDelete)
	del.ICKDelete = *ick.Delete(id, targetid)
	return del
}

func (del *ICKDelete) SetSize(sz ick.SIZE) *ICKDelete {
	del.ICKDelete.SetSize(sz)
	if del.DOM.IsInDOM() {
		del.DOM.PickClass(ick.SIZE_OPTIONS, string(del.SIZE))
	}
	return del
}

/******************************************************************************/

// Wrap implements the JSValueWrapper to enable wrapping of a dom.Element usually
// to wrap embedded component instantiated during unfolding an html string.
// Does not need to be overloaded by the component embedding UISnippet.
func (del *ICKDelete) Wrap(jsvp js.JSValueProvider) {
	if del.UI.Wrap(jsvp); !del.DOM.IsInDOM() {
		verbose.Error("ICKDelete.Wrap:", errors.New("not in DOM"))
		return
	}
	t, has := del.DOM.Attribute("data-targetid")
	verbose.Debug("ICKDelete.Wrap: data-targetid=%q", t)
	if has {
		del.TargetId = t
	}
}

func (del *ICKDelete) AddListeners() {
	verbose.Debug("ICKDelete.AddListeners: targetid=%q", del.TargetId)
	if del.TargetId != "" {
		del.DOM.AddMouseEvent(event.MOUSE_ONCLICK, func(*event.MouseEvent, *dom.Element) {
			del.RemoveTarget()
		})
		del.Clock.Start(del.RemoveTarget)
	} else {
		verbose.Println(verbose.WARNING, "ICKDelete.AddListeners: missing TargetID")
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
