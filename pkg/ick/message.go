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
//
// [bulma message]: https://bulma.io/documentation/components/message/
type ICKMessage struct {
	html.BareSnippet

	// optional header to display on top of the message
	Header html.HTMLString

	// the body of the message
	Msg html.ContentStack

	// set to true to display the delete button and allow user to delete the message
	CanDelete bool

	// COLOR define the color of the message
	COLOR

	// SIZE define the size of the message
	SIZE
}

// Ensuring ICKMessage implements the right interface
var _ html.ElementComposer = (*ICKMessage)(nil)

func Message(cnt html.ContentComposer) *ICKMessage {
	msg := new(ICKMessage)
	msg.Msg.Push(cnt)
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
		html.RenderString(out, `<div class="message-header">`)
		html.RenderChild(out, msg, html.Snippet("span", "", &msg.Header))
		if msg.CanDelete {
			id := msg.Id()
			if id == "" {
				verbose.Debug("ICKMessage: Rendering Deletable Message without TargetId")
				html.RenderChild(out, msg, html.ToHTML(`<ick-delete/>`))
			} else {
				html.RenderChild(out, msg, html.ToHTML(`<ick-delete id="del`+id+`" TargetId='`+id+`'/>`))
			}
		}
		html.RenderString(out, `</div>`)
	}

	if msg.Msg.HasContent() {
		html.RenderString(out, `<div class="message-body">`)
		msg.Msg.RenderStack(out, msg)
		html.RenderString(out, `</div>`)
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
