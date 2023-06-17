package html

import (
	"io"

	"github.com/sunraylab/verbose"
)

func init() {
	RegisterComposer("ick-message", &Message{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type Message struct {
	HTMLSnippet

	Header    HTMLString // optional header to display on top of the message
	Message   HTMLString // message to display within the message
	CanDelete bool       // set to true to display the delete button and allow user to delete the message
}

// Ensure Message implements HTMLComposer interface
var _ HTMLComposer = (*Message)(nil)

func (msg *Message) Tag() *Tag {
	msg.tag.SetName("div")
	msg.tag.Attributes().AddClasses("message")
	return &msg.tag
}

func (msg *Message) WriteBody(out io.Writer) error {
	if msg.Header != "" {
		var delhtml string
		if msg.CanDelete {
			verbose.Printf(verbose.WARNING, "Message can delete TargetId=%s", msg.Id())
			delhtml = `<ick-delete TargetID='` + msg.Id() + `'/>`
		}
		WriteStrings(out, `<div class="message-header"><p>`, string(msg.Header), `</p>`, delhtml, `</div>`)
	}
	WriteStringsIf(msg.Message != "", out, `<div class="message-body">`, string(msg.Message), `</div>`)
	return nil
}
