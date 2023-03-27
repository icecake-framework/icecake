package ui

import (
	_ "embed"

	"github.com/sunraylab/icecake/pkg/ick"
)

/******************************************************************************
* Component
******************************************************************************/

//go:embed "notify.css"
var notifycss string

func init() {
	ick.RegisterComposer("ick-notify", &Notify{})
}

type Notify struct {
	Snippet

	// the message to display within the notification
	Message ick.HTMLstring

	Delete // TODO use en embedded the delete sub-component

}

func (s Notify) Template(*ick.DataState) (t ick.SnippetTemplate) {
	t.TagName = "div"
	t.Attributes = `class="notification"`
	//	t.Body = `<ick-delete/>` + s.Message
	t.Body = ick.HTML(&s.Delete) + s.Message
	return
}
