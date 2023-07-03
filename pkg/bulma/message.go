package bulma

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/lolorenzo777/verbose"
)

func init() {
	html.RegisterComposer("ick-message", &ICKMessage{})
}

// bulma.Message is an icecake snippet providing the HTML rendering for a [bulma message].
// The card is an HTMLSnippet. Use AddContent to setup the content of the body of the message.
//
//	Use `<ick-message/>` for inline rendering.
//
// [bulma message]: https://bulma.io/documentation/components/message/
type ICKMessage struct {
	html.HTMLSnippet

	// optional header to display on top of the message
	Header html.HTMLString

	// set to true to display the delete button and allow user to delete the message
	CanDelete bool

	// COLOR define the color of the message
	COLOR

	// SIZE define the size of the message
	SIZE
}

// Ensure Message implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*ICKMessage)(nil)

func Message(cmp html.HTMLComposer) *ICKMessage {
	msg := new(ICKMessage)
	msg.AddContent(cmp)
	return msg
}

// SetHeader set a message header
func (msg *ICKMessage) SetHeader(header html.HTMLString, candelete bool) *ICKMessage {
	msg.Header = header
	msg.CanDelete = candelete
	return msg
}

// SetColor set a message color
func (msg *ICKMessage) SetColor(c COLOR) *ICKMessage {
	msg.COLOR = c
	return msg
}

// SetSize set the size of the message
func (msg *ICKMessage) SetSize(s SIZE) *ICKMessage {
	msg.SIZE = s
	return msg
}

// BuildTag returns tag <div class="message {classes}" {attributes}>
func (msg *ICKMessage) BuildTag() html.Tag {
	msg.Tag().
		SetTagName("div").
		AddClass("message").
		PickClass(COLOR_OPTIONS, string(msg.COLOR)).
		PickClass(SIZE_OPTIONS, string(msg.SIZE))
	return *msg.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (msg *ICKMessage) RenderContent(out io.Writer) error {

	if !msg.Header.IsEmpty() {
		html.WriteString(out, `<div class="message-header">`)
		msg.RenderChilds(out, html.Span().AddContent(&msg.Header))
		if msg.CanDelete {
			verbose.Debug("Message can delete TargetId=%s", msg.Id())
			html.Render(out, nil, html.ToHTML(`<ick-delete TargetID='`+msg.Id()+`'/>`))
		}
		html.WriteString(out, `</div>`)
	}

	if msg.HasContent() {
		html.WriteString(out, `<div class="message-body">`)
		msg.HTMLSnippet.RenderContent(out)
		html.WriteString(out, `</div>`)
	}

	return nil
}
