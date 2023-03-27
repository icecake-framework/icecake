package ui

import (
	"github.com/sunraylab/icecake/pkg/ick"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	ick.RegisterComposer("ick-message", &Message{})
}

type Message struct {
	Snippet

	Header    ick.HTMLstring // optional header to display on top of the message
	CanDelete bool           // set to true to display the delete button and allow user to delete the message
	Message   ick.HTMLstring // message to display within the notification
}

func (_msg Message) Template(*ick.DataState) (_t ick.SnippetTemplate) {
	_t.TagName = "div"
	_t.Attributes = `class="message"`
	if _msg.Header != "" {
		var delhtml ick.HTMLstring
		if _msg.CanDelete {
			delhtml = `<ick-delete TargetID='` + ick.HTMLstring(_msg.Id()) + `'/>`
		}
		_t.Body = `<div class="message-header"><p>` + _msg.Header + `</p>` + delhtml + `</div>`
	}
	_t.Body += `<div class="message-body">` + _msg.Message + `</div>`
	return
}
