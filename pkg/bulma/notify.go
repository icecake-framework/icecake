package bulma

import (
	_ "embed"
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

/******************************************************************************
* Component
******************************************************************************/

//go:embed "notify.css"
var notifycss string

func init() {
	html.RegisterComposer("ick-notify", &Notify{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type Notify struct {
	html.HTMLSnippet

	// the message to display within the notification
	Message html.HTMLString

	// Notify includes a programmable Delete Button.
	//
	// Delete is a local variable rather than an embedded struct to avoid AddliSteners interface conflict.
	// Notify implements AddliSteners interface via the UISnippet embedded.
	Delete Delete

	// TODO: handle Notify toast style
	// Toast bool
}

// Ensure Notify implements HTMLComposer interface
var _ html.HTMLComposer = (*Notify)(nil)

func (notify *Notify) BuildTag(tag *html.Tag) {
	tag.SetName("div")
	tag.Attributes().AddClasses("notification")
}

func (notify *Notify) RenderContent(out io.Writer) error {
	notify.Delete.TargetID = notify.Id()

	notify.RenderChildSnippet(out, &notify.Delete)

	notify.RenderChildHTML(out, notify.Message)

	return nil
}
