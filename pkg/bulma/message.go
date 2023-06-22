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

// Tag Builder used by the rendering functions.
func (msg *Message) BuildTag(tag *html.Tag) {
	tag.SetTagName("div").AddClasses("message")
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (msg *Message) RenderContent(out io.Writer) error {

	if !msg.Header.IsEmpty() {

		html.WriteString(out, `<div class="message-header">`)
		html.WriteStrings(out, `<p>`)
		html.RenderHTML(out, nil, msg.Header)
		html.WriteString(out, `</p>`)

		if msg.CanDelete {
			verbose.Printf(verbose.WARNING, "Message can delete TargetId=%s", msg.Id())
			html.RenderHTML(out, nil, html.HTML(`<ick-delete TargetID='`+msg.Id()+`'/>`))
		}
		html.WriteString(out, `</div>`)
	}

	html.RenderHTMLIf(!msg.Message.IsEmpty(), out, msg, html.HTML(`<div class="message-body">`), msg.Message, html.HTML(`</div>`))

	return nil
}
