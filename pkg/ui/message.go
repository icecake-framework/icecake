package ui

import (
	ick "github.com/sunraylab/icecake/pkg/icecake"
	wick "github.com/sunraylab/icecake/pkg/wicecake"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	ick.TheCmpReg.RegisterComponent(Message{})
}

type Message struct {
	wick.UIComponent // embedded Component, with default implementation of composer interfaces

	Header    string // optional header to display on top of the message
	CanDelete bool   // set to true to display the delete button and allow user to delete the message
	Message   string // message to display within the notification
}

func (*Message) RegisterName() string {
	return "ick-message"
}

func (*Message) Container(_compid string) (_tagname string, _contclasses string, _contattrs string, _contstyle string) {
	return "div", "ick-message message", "", ""
}

func (_msg *Message) Body() (_html string) {
	if _msg.Header != "" {
		var delhtml string
		if _msg.CanDelete {
			//			delhtml = `<button class="delete" aria-label="delete"></button>`
			delhtml = `<ick-delete TargetID='{{.Id}}'/>`
		}
		_html += `<div class="message-header"><p>` + _msg.Header + `</p>` + delhtml + `</div>`
	}
	_html += `<div class="message-body">` + _msg.Message + `</div>`
	return _html
}
