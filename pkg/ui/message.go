package ui

import (
	ick "github.com/sunraylab/icecake/pkg/icecake2"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	ick.RegisterComposer("ick-message", Message{})
}

type Message struct {
	Snippet

	Header    ick.HTML // optional header to display on top of the message
	CanDelete bool     // set to true to display the delete button and allow user to delete the message
	Message   ick.HTML // message to display within the notification
}

func (_msg Message) Template(*ick.DataState) (_t ick.SnippetTemplate) {
	_t.TagName = "div"
	_t.Attributes = `class="message"`
	if _msg.Header != "" {
		var delhtml ick.HTML
		if _msg.CanDelete {
			delhtml = `<ick-delete TargetID='` + _msg.HtmlSnippet.Id() + `'/>`
		}
		_t.Body = `<div class="message-header"><p>` + _msg.Header + `</p>` + delhtml + `</div>`
	}
	_t.Body += `<div class="message-body">` + _msg.Message + `</div>`
	return
}
