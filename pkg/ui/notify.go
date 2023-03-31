package ui

import (
	_ "embed"

	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/html"
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
	Message html.HTMLstring

	Delete // TODO use en embedded the delete sub-component

}

func (s Notify) Template(*html.DataState) (t html.SnippetTemplate) {
	t.TagName = "div"
	t.Attributes = `class="notification"`
	//	t.Body = `<ick-delete/>` + s.Message
	s.Delete.TargetID = s.Id()
	console.Warnf("delete target id: %s", s.Id())
	t.Body = s.HTML(&s.Delete) + s.Message
	return
}
