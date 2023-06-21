package bulma

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/sunraylab/verbose"
)

func init() {
	html.RegisterComposer("ick-message", &Message{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type Message struct {
	html.HTMLSnippet

	Header    html.HTMLString // optional header to display on top of the message
	Message   html.HTMLString // message to display within the message
	CanDelete bool            // set to true to display the delete button and allow user to delete the message
}

// Ensure Message implements HTMLComposer interface
var _ html.HTMLComposer = (*Message)(nil)

func (msg *Message) BuildTag(tag *html.Tag) {
	tag.SetName("div").Attributes().AddClasses("message")
}

func (msg *Message) RenderContent(out io.Writer) error {
	if !msg.Header.IsEmpty() {
		var delhtml html.HTMLString
		if msg.CanDelete {
			verbose.Printf(verbose.WARNING, "Message can delete TargetId=%s", msg.Id())
			delhtml = html.HTML(`<ick-delete TargetID='` + msg.Id() + `'/>`)
		}
		html.RenderHTML(out, msg, html.HTML(`<div class="message-header"><p>`), msg.Header, html.HTML(`</p>`), delhtml, html.HTML(`</div>`))
	}
	html.RenderHTMLIf(!msg.Message.IsEmpty(), out, msg, html.HTML(`<div class="message-body">`), msg.Message, html.HTML(`</div>`))
	return nil
}
