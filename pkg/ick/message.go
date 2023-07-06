package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/lolorenzo777/verbose"
)

func init() {
	html.RegisterComposer("ick-message", &ICKMessage{})
}

// ICKMessage is an icecake snippet providing the HTML rendering for a [bulma message].
// The message is an HTMLSnippet. Use AddContent to setup the content of the body of the message.
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

// Ensure Message implements HTMLComposer interface
var _ html.HTMLComposer = (*ICKMessage)(nil)

func Message(cnt html.HTMLContentComposer) *ICKMessage {
	msg := new(ICKMessage)
	msg.AddContent(cnt)
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
		msg.RenderChild(out, html.Snippet("span").AddContent(&msg.Header))
		if msg.CanDelete {
			id := msg.Id()
			if id == "" {
				verbose.Debug("ICKMessage: Rendering Deletable Message without TargetId")
				msg.RenderChild(out, html.ToHTML(`<ick-delete/>`))
			} else {
				msg.RenderChild(out, html.ToHTML(`<ick-delete id="del`+id+`" TargetId='`+id+`'/>`))
			}
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

/******************************************************************************/

// SetHeader set a message header
func (msg *ICKMessage) SetHeader(header html.HTMLString) *ICKMessage {
	msg.Header = header
	return msg
}

// SetDeletable make this message delatable by rendering the delete button.
// A deletable message must have an id. Dos nothing if id is empty.
func (msg *ICKMessage) SetDeletable(id string) *ICKMessage {
	if id != "" {
		msg.Tag().SetId(id)
		msg.CanDelete = true
	}
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
