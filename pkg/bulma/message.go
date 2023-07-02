package bulma

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/sunraylab/verbose"
)

func init() {
	html.RegisterComposer("ick-message", &Message{})
}

// bulma.Message is an icecake snippet providing the HTML rendering for a [bulma message].
//
//	Use `<ick-message/>` for inline rendering.
//
// [bulma message]: https://bulma.io/documentation/components/message/
type Message struct {
	html.HTMLSnippet

	Header    html.HTMLString // optional header to display on top of the message
	Msg       html.HTMLString // message to display within the message
	CanDelete bool            // set to true to display the delete button and allow user to delete the message
}

// Ensure Message implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*Message)(nil)

// Tag Builder used by the rendering functions.
func (msg *Message) BuildTag(tag *html.Tag) {
	tag.SetTagName("div").AddClass("message")
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (msg *Message) RenderContent(out io.Writer) error {

	if !msg.Header.IsEmpty() {
		html.WriteString(out, `<div class="message-header">`)
		msg.RenderChilds(out, html.P().SetBody(&msg.Header))

		if msg.CanDelete {
			verbose.Debug("Message can delete TargetId=%s", msg.Id())
			html.Render(out, nil, html.ToHTML(`<ick-delete TargetID='`+msg.Id()+`'/>`))
		}
		html.WriteString(out, `</div>`)
	}

	if !msg.Msg.IsEmpty() {
		msg.RenderChilds(out, html.Div(`class="message-body"`).SetBody(&msg.Msg))
	}

	return nil
}
