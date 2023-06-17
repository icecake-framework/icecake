package html

import (
	_ "embed"
	"io"
)

/******************************************************************************
* Component
******************************************************************************/

//go:embed "notify.css"
var notifycss string

func init() {
	RegisterComposer("ick-notify", &Notify{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type Notify struct {
	HTMLSnippet

	// the message to display within the notification
	Message HTMLString

	// Notify includes a programmable Delete Button.
	//
	// Delete is a local variable rather than an embedded struct to avoid AddliSteners interface conflict.
	// Notify implements AddliSteners interface via the UISnippet embedded.
	Delete Delete

	// TODO: handle toast style
	// Toast bool
}

// Ensure Notify implements HTMLComposer interface
var _ HTMLComposer = (*Notify)(nil)

func (notify *Notify) Tag() *Tag {
	notify.tag.SetName("div")
	notify.tag.Attributes().AddClasses("notification")
	return &notify.tag
}

func (notify *Notify) WriteBody(out io.Writer) error {
	notify.Delete.TargetID = notify.Id()

	notify.WriteChildSnippet(out, &notify.Delete)

	WriteString(out, string(notify.Message))

	return nil
}
