package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaMessage string = `<a href="https://ick.io/documentation/components/message">bulma Message</a>`
)

type SectionDocMessage struct{ SectionDocIcecake }

func (sec *SectionDocMessage) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Message</h2>`+
		`<p>ICKMessage is an icecake snippet providing the HTML rendering for a `, linkBulmaMessage, ` with extra features and usefull Go APIs.</p>`)

	// usages
	u1 := ick.Button(*html.ToHTML("reset"), "btnreset", "", `class="mb-3"`).
		SetColor(ick.COLOR_PRIMARY).
		SetOutlined(true).
		SetDisabled(true)
	u2 := html.Div(`class="box mr-5"`).
		SetId("boxusage").
		AddContent(ick.Spinner())
	html.Render(out, nil, u1, u2)

	// apis
	html.WriteString(out, `<h3>ick.ICKMessage API</h3>`+
		`<p><code>Message(content html.HTMLComposer) *ICKMessage</code> is the main Message factory.</p>`+
		`<p><code>.SetHeader(header html.HTMLString, candelete bool)</code> set a header with or without the delete button.</p>`+
		`<p>The Message is an HTMLSnippet so you can use <code>AddContent</code> to setup the content of the message.</p>`+
		`<p><code>.AddContent(cmp html.HTMLComposer) error</code> Stack content inside.</p>`)

	// rendering
	html.WriteString(out, `<h3>Rendering</h3>`)
	html.WriteString(out, `<div class="box mr-5">`)
	r1 := ick.Message(html.ToHTML("This is a simple message."))
	r2 := ick.Message(html.ToHTML("This is a message with a header.")).SetHeader(*html.ToHTML("Icecake Message"))
	r3 := ick.Message(html.ToHTML("This is a message with the delete button.")).SetHeader(*html.ToHTML("Icecake Message")).SetDeletable("msgr3")
	r4 := ick.Message(nil).SetHeader(*html.ToHTML("Only header")).SetDeletable("msgr4")
	html.Render(out, nil, r1, r2, r3, r4)
	html.WriteString(out, `</div>`)

	// styling
	html.WriteString(out, `<h3>Styling</h3>`)
	html.WriteString(out, `<div class="box mr-5">`)

	styl1 := ick.Message(html.ToHTML("Make use of ick.COLOR property.<br>COLOR = COLOR_PRIMARY")).SetHeader(*html.ToHTML("Icecake Message")).
		SetColor(ick.COLOR_INFO)
	styl2 := ick.Message(html.ToHTML("Make use of ick.SIZE property.<br>SIZE = SIZE_SMALL")).SetHeader(*html.ToHTML("Icecake Message")).
		SetSize(ick.SIZE_SMALL)
	html.Render(out, nil, styl1, styl2)

	html.WriteString(out, `</div>`)

	return nil
}