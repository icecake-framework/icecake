package ui

import (
	_ "embed"

	"github.com/sunraylab/icecake/pkg/dom"
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
	dom.UISnippet

	// the message to display within the notification
	Message html.String

	// Notify includes a programmable Delete Button.
	//
	// Delete is a local variable rather than an embedded struct to avoid AddliSteners interface conflict.
	// Notify implements AddliSteners interface via the UISnippet embedded.
	Delete Delete

	// TODO: handle toast style
	// Toast bool
}

func (_cmp *Notify) Template(*html.DataState) (t html.SnippetTemplate) {

	_cmp.Delete.TargetID = _cmp.Id()

	t.TagName = "div"
	t.Attributes = `class="notification"`
	t.Body = _cmp.HTML(&_cmp.Delete) + _cmp.Message
	return
}
