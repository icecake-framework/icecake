package ick

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
	html.BareSnippet

	// the message to display within the notification
	Message html.HTMLString

	// Notify includes a programmable Delete Button.
	// Delete is a local variable rather than an embedded struct to avoid AddliSteners interface conflict.
	// Notify implements AddliSteners interface via the UISnippet embedded.
	Delete ICKDelete

	// TODO: ick.Notify - handle toast style
	// Toast bool
}

// Ensuring ICKNotify implements the right interface
var _ html.ElementComposer = (*Notify)(nil)

// BuildTag builds the tag used to render the html element.
// Notify tag is a simple <div class="notification"></div>
func (notify *Notify) BuildTag() html.Tag {
	notify.Tag().SetTagName("div").AddClass("notification")
	return *notify.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (notify *Notify) RenderContent(out io.Writer) error {
	notify.Delete.TargetId = notify.Id()
	html.RenderChild(out, notify, &notify.Delete)
	html.RenderChild(out, notify, &notify.Message)
	return nil
}
