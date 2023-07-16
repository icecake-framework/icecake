package ickui

import (
	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

type ICKMessage struct {
	ick.ICKMessage
	dom.UI

	BtnDelete ICKDelete // The delete button snippet created only if candelete is true
}

// Ensure ICKMessage implements UIComposer interface
var _ dom.UIComposer = (*ICKMessage)(nil)

// Message factory
func Message(cnt ickcore.ContentComposer) *ICKMessage {
	msg := new(ICKMessage)
	msg.ICKMessage = *ick.Message(cnt)
	return msg
}

/******************************************************************************/

// SetColor set a message color. Immediate effect to the DOM.
func (msg *ICKMessage) SetColor(c ick.COLOR) *ICKMessage {
	msg.ICKMessage.SetColor(c)
	if msg.DOM.IsInDOM() {
		msg.DOM.PickClass(ick.COLOR_OPTIONS, string(msg.COLOR))
	}
	return msg
}

// SetSize set the size of the message. Immediate effect to the DOM.
func (msg *ICKMessage) SetSize(s ick.SIZE) *ICKMessage {
	msg.ICKMessage.SetSize(s)
	if msg.DOM.IsInDOM() {
		msg.DOM.PickClass(ick.SIZE_OPTIONS, string(msg.SIZE))
	}
	return msg
}

/******************************************************************************/

// AddListeners adds the listener to the embedded delete button, if any.
func (msg *ICKMessage) AddListeners() {
	console.Warnf("ICKMessage.AddListeners")

	if !msg.DOM.IsInDOM() {
		console.Errorf("ICKMessage.AddListeners NOT IN DOM")
	}

	dom.TryMountId(&msg.BtnDelete, msg.Tag().SubId("btndel"))

}

// RemoveListeners remove delete button listeners
func (msg *ICKMessage) RemoveListeners() {
	msg.BtnDelete.RemoveListeners()
	msg.UI.RemoveListeners()
}
