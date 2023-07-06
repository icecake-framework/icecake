package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/html"
)

const (
	linkBulmaMessage string = `<a href="https://bulma.io/documentation/components/message">bulma Message</a>`
)

type SectionBulmaMessage struct{ SectionIcecakeDoc }

func (sec *SectionBulmaMessage) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Message</h2>`+
		`<p>bulma.Message is an icecake snippet providing the HTML rendering for a `, linkBulmaMessage, ` with extra features and usefull Go APIs.</p>`)

	// usages
	u1 := bulma.Button(*html.ToHTML("reset"), "btnreset", "", `class="mb-3"`).
		SetColor(bulma.COLOR_PRIMARY).
		SetOutlined(true).
		SetDisabled(true)
	u2 := html.Div(`class="box mr-5"`).
		SetId("boxusage").
		AddContent(bulma.Spinner())
	html.Render(out, nil, u1, u2)

	// apis
	html.WriteString(out, `<h3>bulma.ICKMessage API</h3>`+
		`<p><code>Message(content html.HTMLComposer) *ICKMessage</code> is the main Message factory.</p>`+
		`<p><code>.SetHeader(header html.HTMLString, candelete bool)</code> set a header with or without the delete button.</p>`+
		`<p>The Message is an HTMLSnippet so you can use <code>AddContent</code> to setup the content of the message.</p>`+
		`<p><code>.AddContent(cmp html.HTMLComposer) error</code> Stack content inside.</p>`)

	// rendering
	html.WriteString(out, `<h3>Rendering</h3>`)
	html.WriteString(out, `<div class="box mr-5">`)
	r1 := bulma.Message(html.ToHTML("This is a simple message."))
	r2 := bulma.Message(html.ToHTML("This is a message with a header.")).SetHeader(*html.ToHTML("Icecake Message"))
	r3 := bulma.Message(html.ToHTML("This is a message with the delete button.")).SetHeader(*html.ToHTML("Icecake Message")).SetDeletable("msgr3")
	r4 := bulma.Message(nil).SetHeader(*html.ToHTML("Only header")).SetDeletable("msgr4")
	html.Render(out, nil, r1, r2, r3, r4)
	html.WriteString(out, `</div>`)

	// styling
	html.WriteString(out, `<h3>Styling</h3>`)
	html.WriteString(out, `<div class="box mr-5">`)

	styl1 := bulma.Message(html.ToHTML("Make use of bulma.COLOR property.<br>COLOR = COLOR_PRIMARY")).SetHeader(*html.ToHTML("Icecake Message")).
		SetColor(bulma.COLOR_INFO)
	styl2 := bulma.Message(html.ToHTML("Make use of bulma.SIZE property.<br>SIZE = SIZE_SMALL")).SetHeader(*html.ToHTML("Icecake Message")).
		SetSize(bulma.SIZE_SMALL)
	html.Render(out, nil, styl1, styl2)

	html.WriteString(out, `</div>`)

	return nil
}
