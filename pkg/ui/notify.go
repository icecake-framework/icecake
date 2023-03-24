package ui

import (
	_ "embed"

	ick "github.com/sunraylab/icecake/pkg/icecake2"
)

/******************************************************************************
* Component
******************************************************************************/

//go:embed "notify.css"
var notifycss string

func init() {
	ick.RegisterComposer("ick-notify", Notify{})
}

type Notify struct {
	Snippet

	// the message to display within the notification
	Message ick.HTML

	Delete // TODO use en embedded the delete sub-component

}

func (_snippet Notify) Template(*ick.DataState) (_t ick.SnippetTemplate) {
	_t.TagName = "div"
	_t.Attributes = `class="notification"`
	_t.Body = `<ick-delete/>` + _snippet.Message
	return
}
