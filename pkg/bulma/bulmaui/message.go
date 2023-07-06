package bulmaui

import (
	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
)

type ICKMessage struct {
	bulma.ICKMessage
	dom.UI

	BtnDelete ICKDelete // The delete button snippet created only if candelete is true
}

// Ensure Card implements HTMLTagComposer interface
var _ dom.UIComposer = (*ICKMessage)(nil)

// Message factory
func Message(cnt html.HTMLContentComposer) *ICKMessage {
	msg := new(ICKMessage)
	msg.ICKMessage = *bulma.Message(cnt)
	return msg
}

/******************************************************************************/

// SetColor set a message color
func (msg *ICKMessage) SetColor(c bulma.COLOR) *ICKMessage {
	msg.ICKMessage.SetColor(c)
	if msg.DOM.IsInDOM() {
		msg.DOM.PickClass(bulma.COLOR_OPTIONS, string(msg.COLOR))
	}
	return msg
}

// SetSize set the size of the message
func (msg *ICKMessage) SetSize(s bulma.SIZE) *ICKMessage {
	msg.ICKMessage.SetSize(s)
	if msg.DOM.IsInDOM() {
		msg.DOM.PickClass(bulma.SIZE_OPTIONS, string(msg.SIZE))
	}
	return msg
}

/******************************************************************************/

// AddListeners adds the listener to the embedded delete button, if any.
func (msg *ICKMessage) AddListeners() {
	console.Warnf("ICKMessage.AddListeners")
	if msg.DOM.Id() != dom.UNDEFINED_NODE {
		msg.BtnDelete.RemoveListeners()
		// Mount the button only if it's in the DOM
		if btndelid := "del" + msg.DOM.Id(); dom.Doc().IsInDOM(btndelid) {
			if err := dom.TryMountId(&msg.BtnDelete, btndelid); err != nil {
				console.Errorf("ICKMessage.AddListeners: %s", err.Error())
			}
		}
	}
}

func (msg *ICKMessage) RemoveListeners() {
	msg.BtnDelete.RemoveListeners()
	msg.UI.RemoveListeners()
}
