package ui

import (
	ick "github.com/sunraylab/icecake/pkg/icecake"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	ick.App.RegisterComponent("ick-message", Message{}, "")
}

type Message struct {
	ick.UIComponent // embedded Component, with default implementation of composer interfaces

	Header    string // optional header to display on top of the message
	CanDelete bool   // set to true to display the delete button and allow user to delete the message
	Message   string // message to display within the notification
}

func (c *Message) Container(_compid string) (_tagname string, _classes string, _attrs string) {
	return "div", "ick-message message", ""
}

func (c *Message) Body() (_html string) {
	if c.Header != "" {
		var delhtml string
		if c.CanDelete {
			//			delhtml = `<button class="delete" aria-label="delete"></button>`
			delhtml = `<ick-delete TargetID='{{.Id}}'/>`
		}
		_html += `<div class="message-header"><p>` + c.Header + `</p>` + delhtml + `</div>`
	}
	_html += `<div class="message-body">` + c.Message + `</div>`
	return _html
}
