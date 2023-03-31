package ui

import (
	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/html"
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

	Header    html.HTMLstring // optional header to display on top of the message
	CanDelete bool            // set to true to display the delete button and allow user to delete the message
	Message   html.HTMLstring // message to display within the notification
}

func (_msg Message) Template(*html.DataState) (_t html.SnippetTemplate) {
	_t.TagName = "div"
	_t.Attributes = `class="message"`
	if _msg.Header != "" {
		var delhtml html.HTMLstring
		if _msg.CanDelete {
			console.Warnf("TargetId=%s", _msg.Id())
			delhtml = `<ick-delete TargetID='` + html.HTMLstring(_msg.Id()) + `'/>`
		}
		_t.Body = `<div class="message-header"><p>` + _msg.Header + `</p>` + delhtml + `</div>`
	}
	if _msg.Message != "" {
		_t.Body += `<div class="message-body">` + _msg.Message + `</div>`
	}
	return
}
