package ickui

import (
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/ick"
)

type ICKTagLabel struct {
	ick.ICKTagLabel
	dom.UI

	BtnDelete ICKDelete // The delete button snippet created only if candelete is true
}

// Ensure ICKTagLabel implements the right interfaces
var _ dom.UIComposer = (*ICKTagLabel)(nil)

// TagLabel factory
func TagLabel(text string, c ick.COLOR, attrs ...string) *ICKTagLabel {
	msg := new(ICKTagLabel)
	msg.ICKTagLabel = *ick.TagLabel(text, c, attrs...)
	return msg
}

/******************************************************************************/

// AddListeners adds the listener to the embedded delete button, if any.
func (t *ICKTagLabel) AddListeners() {
	// DEBUG: console.Warnf("ICKTagLabel.AddListeners")
	t.BtnDelete.TargetId = t.BtnDelete.Tag().Id()
	dom.TryMountId(&t.BtnDelete, t.Tag().SubId("btndel"))
}

// RemoveListeners remove delete button listeners
func (t *ICKTagLabel) RemoveListeners() {
	t.BtnDelete.RemoveListeners()
	t.UI.RemoveListeners()
}
