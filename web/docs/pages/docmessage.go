package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

const (
	linkBulmaMessage string = `<a href="https://bulma.io/documentation/components/message">bulma Message</a>`
)

type SectionDocMessage struct{ SectionDocIcecake }

func (sec *SectionDocMessage) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Message", "message.go", "ICKMessage")

	ickcore.RenderString(out, `<div class="block">`+
		`<p>ICKMessage is an icecake snippet providing the HTML rendering for a `, linkBulmaMessage, ` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// usages
	ickcore.RenderChild(out, sec,
		ick.Elem("div", `id="boxusage" class="box"`, ick.Spinner()),
		ick.Button("reset", `class="mb-5"`).
			SetId("btnreset").
			SetColor(ick.COLOR_PRIMARY).
			SetOutlined(true).
			SetDisabled(true))

	// rendering
	ickcore.RenderChild(out, sec, ick.Title(4, "Rendering"))
	ickcore.RenderString(out, `<div class="box">`)
	ickcore.RenderChild(out, sec,
		ick.Message(ickcore.ToHTML("This is a simple message.")),
		ick.Message(ickcore.ToHTML("This is a message with a header.")).
			SetHeader(*ickcore.ToHTML("Icecake Message")),
		ick.Message(ickcore.ToHTML("This is a message with the delete button.")).
			SetHeader(*ickcore.ToHTML("Icecake Message")).
			SetDeletable("msgr3"),
		ick.Message(nil).SetHeader(*ickcore.ToHTML("Only header")).
			SetDeletable("msgr4"),
		ick.Message(ickcore.ToHTML("This is a deletable message.<br>Two lines")).
			SetDeletable("msgr5"))
	ickcore.RenderString(out, `</div>`)

	// styling
	ickcore.RenderChild(out, sec, ick.Title(4, "Styling"))
	ickcore.RenderString(out, `<div class="box">`)
	ickcore.RenderChild(out, sec,
		ick.Message(ickcore.ToHTML("Make use of ick.COLOR property.<br>COLOR = COLOR_PRIMARY")).
			SetHeader(*ickcore.ToHTML("Icecake Message")).
			SetColor(ick.COLOR_INFO),
		ick.Message(ickcore.ToHTML("Make use of ick.SIZE property.<br>SIZE = SIZE_SMALL")).
			SetHeader(*ickcore.ToHTML("Icecake Message")).
			SetSize(ick.SIZE_SMALL))
	ickcore.RenderString(out, `</div>`)

	return nil
}
