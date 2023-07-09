package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaMessage string = `<a href="https://bulma.io/documentation/components/message">bulma Message</a>`
)

type SectionDocMessage struct{ SectionDocIcecake }

func (sec *SectionDocMessage) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Message", "message.go", "ICKMessage")

	html.RenderString(out, `<div class="block">`+
		`<p>ICKMessage is an icecake snippet providing the HTML rendering for a `, linkBulmaMessage, ` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// usages
	html.RenderChild(out, sec,
		html.Snippet("div", `id="boxusage" class="box"`, ick.Spinner()),
		ick.Button("reset", `class="mb-5"`).
			SetId("btnreset").
			SetColor(ick.COLOR_PRIMARY).
			SetOutlined(true).
			SetDisabled(true))

	// rendering
	html.RenderChild(out, sec, ick.Title(4, "Rendering"))
	html.RenderString(out, `<div class="box">`)
	html.RenderChild(out, sec,
		ick.Message(html.ToHTML("This is a simple message.")),
		ick.Message(html.ToHTML("This is a message with a header.")).
			SetHeader(*html.ToHTML("Icecake Message")),
		ick.Message(html.ToHTML("This is a message with the delete button.")).
			SetHeader(*html.ToHTML("Icecake Message")).
			SetDeletable("msgr3"),
		ick.Message(nil).SetHeader(*html.ToHTML("Only header")).
			SetDeletable("msgr4"))
	html.RenderString(out, `</div>`)

	// styling
	html.RenderChild(out, sec, ick.Title(4, "Styling"))
	html.RenderString(out, `<div class="box">`)
	html.RenderChild(out, sec,
		ick.Message(html.ToHTML("Make use of ick.COLOR property.<br>COLOR = COLOR_PRIMARY")).
			SetHeader(*html.ToHTML("Icecake Message")).
			SetColor(ick.COLOR_INFO),
		ick.Message(html.ToHTML("Make use of ick.SIZE property.<br>SIZE = SIZE_SMALL")).
			SetHeader(*html.ToHTML("Icecake Message")).
			SetSize(ick.SIZE_SMALL))
	html.RenderString(out, `</div>`)

	return nil
}
