package bulma

import (
	_ "embed"
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

//go:embed "notify.css"
var notifycss string

func init() {
	html.RegisterComposer("ick-notify", &Notify{})
	html.RequireCSSStyle("ick-notify", notifycss)
}

type Notify struct {
	html.HTMLSnippet

	// the message to display within the notification
	Message html.HTMLString

	// Notify includes a programmable Delete Button.
	// Delete is a local variable rather than an embedded struct to avoid AddliSteners interface conflict.
	// Notify implements AddliSteners interface via the UISnippet embedded.
	Delete Delete

	// TODO: handle Notify toast style
	// Toast bool
}

// Ensure Notify implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*Notify)(nil)

// BuildTag builds the tag used to render the html element.
// Notify tag is a simple <div class="notification"></div>
func (notify *Notify) BuildTag() html.Tag {
	notify.Tag().SetTagName("div").AddClass("notification")
	return *notify.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (notify *Notify) RenderContent(out io.Writer) error {
	notify.Delete.TargetID = notify.Id()
	notify.RenderChilds(out, &notify.Delete)
	notify.RenderChilds(out, &notify.Message)
	return nil
}
