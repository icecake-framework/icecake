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
	html.WriteString(out, `<div class="block">`+
		`<p>ICKMessage is an icecake snippet providing the HTML rendering for a `, linkBulmaMessage, ` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// usages
	ux := html.Snippet("div", `id="boxusage" class="box"`).AddContent(ick.Spinner())
	btnreset := ick.Button("reset", `class="mb-5"`).
		SetId("btnreset").
		SetColor(ick.COLOR_PRIMARY).
		SetOutlined(true).
		SetDisabled(true)
	html.Render(out, nil, ux, btnreset)

	// rendering
	html.Render(out, nil, ick.Title(4, "Rendering"))
	html.WriteString(out, `<div class="box">`)
	r1 := ick.Message(html.ToHTML("This is a simple message."))
	r2 := ick.Message(html.ToHTML("This is a message with a header.")).SetHeader(*html.ToHTML("Icecake Message"))
	r3 := ick.Message(html.ToHTML("This is a message with the delete button.")).SetHeader(*html.ToHTML("Icecake Message")).SetDeletable("msgr3")
	r4 := ick.Message(nil).SetHeader(*html.ToHTML("Only header")).SetDeletable("msgr4")
	html.Render(out, nil, r1, r2, r3, r4)
	html.WriteString(out, `</div>`)

	// styling
	html.Render(out, nil, ick.Title(4, "Styling"))
	html.WriteString(out, `<div class="box">`)
	styl1 := ick.Message(html.ToHTML("Make use of ick.COLOR property.<br>COLOR = COLOR_PRIMARY")).SetHeader(*html.ToHTML("Icecake Message")).
		SetColor(ick.COLOR_INFO)
	styl2 := ick.Message(html.ToHTML("Make use of ick.SIZE property.<br>SIZE = SIZE_SMALL")).SetHeader(*html.ToHTML("Icecake Message")).
		SetSize(ick.SIZE_SMALL)
	html.Render(out, nil, styl1, styl2)
	html.WriteString(out, `</div>`)

	return nil
}
