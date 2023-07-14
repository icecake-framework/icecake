package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ickcore"
	"github.com/lolorenzo777/verbose"
)

func init() {
	ickcore.RegisterComposer("ick-message", &ICKMessage{})
}

// ICKMessage is an icecake snippet providing the HTML rendering for a [bulma message].
//
// [bulma message]: https://bulma.io/documentation/components/message/
type ICKMessage struct {
	ickcore.BareSnippet

	// optional header to display on top of the message
	Header ickcore.HTMLString

	// the body of the message
	Msg ickcore.ContentStack

	// set to true to display the delete button and allow user to delete the message
	CanDelete bool

	// COLOR define the color of the message
	COLOR

	// SIZE define the size of the message
	SIZE

	//TODO: ICKMessage - handle Notify style (toast)
}

// Ensuring ICKMessage implements the right interface
var _ ickcore.ContentComposer = (*ICKMessage)(nil)
var _ ickcore.TagBuilder = (*ICKMessage)(nil)

func Message(cnt ickcore.ContentComposer) *ICKMessage {
	msg := new(ICKMessage)
	msg.Msg.Push(cnt)
	return msg
}

// SetHeader set a message header
func (msg *ICKMessage) SetHeader(header ickcore.HTMLString) *ICKMessage {
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

/******************************************************************************/

// BuildTag returns tag <div class="message {classes}" {attributes}>
func (msg *ICKMessage) BuildTag() ickcore.Tag {
	msg.Tag().
		SetTagName("div").
		AddClass("message").
		PickClass(COLOR_OPTIONS, string(msg.COLOR)).
		PickClass(SIZE_OPTIONS, string(msg.SIZE))
	return *msg.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (msg *ICKMessage) RenderContent(out io.Writer) error {
	if msg.Header.NeedRendering() {
		ickcore.RenderString(out, `<div class="message-header">`)
		ickcore.RenderChild(out, msg, Elem("span", "", &msg.Header))
		if msg.CanDelete {
			msg.RenderDelete(out)
		}
		ickcore.RenderString(out, `</div>`)
	}

	if msg.Msg.NeedRendering() {
		ickcore.RenderString(out, `<div class="message-body pr-4 is-flex is-align-items-top is-justify-content-space-between">`)
		ickcore.RenderString(out, `<span>`)
		msg.Msg.RenderStack(out, msg)
		ickcore.RenderString(out, `</span>`)
		if msg.CanDelete && !msg.Header.NeedRendering() {
			msg.RenderDelete(out)
		}
		ickcore.RenderString(out, `</div>`)
	}

	return nil
}

func (msg *ICKMessage) RenderDelete(out io.Writer) error {
	id := msg.Tag().Id()
	if id == "" {
		verbose.Debug("ICKMessage: Rendering Deletable Message without TargetId")
		ickcore.RenderChild(out, msg, ickcore.ToHTML(`<ick-delete/>`))
	} else {
		ickcore.RenderChild(out, msg, ickcore.ToHTML(`<ick-delete id="del`+id+`" TargetId='`+id+`'/>`))
	}
	return nil
}
