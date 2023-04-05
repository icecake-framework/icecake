package ui

import (
	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/html"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	html.RegisterComposer("ick-message", &Message{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type Message struct {
	dom.UISnippet

	Header    html.String // optional header to display on top of the message
	Message   html.String // message to display within the notification
	CanDelete bool        // set to true to display the delete button and allow user to delete the message
}

func (_msg *Message) Template(*html.DataState) (_t html.SnippetTemplate) {
	_t.TagName = "div"
	_t.Attributes = `class="message"`
	if _msg.Header != "" {
		var delhtml html.String
		if _msg.CanDelete {
			console.Warnf("TargetId=%s", _msg.Id())
			delhtml = `<ick-delete TargetID='` + html.String(_msg.Id()) + `'/>`
		}
		_t.Body = `<div class="message-header"><p>` + _msg.Header + `</p>` + delhtml + `</div>`
	}
	if _msg.Message != "" {
		_t.Body += `<div class="message-body">` + _msg.Message + `</div>`
	}
	return
}
