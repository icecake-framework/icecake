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
	html.WriteString(out, `<h2>Bulma Message</h2>`)
	html.WriteString(out, `<p>bulma.Message is an icecake snippet providing the HTML rendering for a `, linkBulmaMessage, ` with extra features and usefull Go APIs.</p>`)

	// apis
	html.WriteString(out, `<h3>bulma.ICKMessage API</h3>`)
	html.WriteString(out, `<p><code>Message(cmp html.HTMLComposer) *ICKMessage</code> is the Message factory.</p>`)
	html.WriteString(out, `<p><code>.SetHeader(header html.HTMLString, candelete bool)</code> set a header with or without the delete button.</p>`)
	html.WriteString(out, `<p>The Message is an HTMLSnippet so you can use <code>AddContent</code> to setup the content of the message.</p>`)
	html.WriteString(out, `<p><code>.AddContent(cmp html.HTMLComposer) error</code> Stack content inside.</p>`)

	// rendering
	html.WriteString(out, `<h3>Usage</h3>`)
	sec.RenderChilds(out, bulma.Button(*html.ToHTML("reset"), "btnreset", "", `class="mb-3"`).SetColor(bulma.COLOR_PRIMARY).SetOutlined(true).SetDisabled(true))
	html.WriteString(out, `<div id="boxusage" class="box mr-5">`)
	html.Render(out, nil, bulma.Spinner())
	html.WriteString(out, `</div>`)

	// rendering
	html.WriteString(out, `<h3>Rendering</h3>`)
	html.WriteString(out, `<div class="box mr-5">`)
	r1 := bulma.Message(html.ToHTML("This is a simple message."))
	r2 := bulma.Message(html.ToHTML("This is a message with a header.")).SetHeader(*html.ToHTML("Icecake Message"), false)
	r3 := bulma.Message(html.ToHTML("This is a message with the delete button.")).SetHeader(*html.ToHTML("Icecake Message"), true)
	r4 := bulma.Message(nil).SetHeader(*html.ToHTML("Only header"), true)
	html.Render(out, sec, r1, r2, r3, r4)
	html.WriteString(out, `</div>`)

	// styling
	html.WriteString(out, `<h3>Styling</h3>`)
	html.WriteString(out, `<div class="box mr-5">`)
	styl1 := bulma.Message(html.ToHTML("Make use of bulma.COLOR property.<br>COLOR = COLOR_PRIMARY")).SetHeader(*html.ToHTML("Icecake Message"), false).
		SetColor(bulma.COLOR_INFO)
	styl2 := bulma.Message(html.ToHTML("Make use of bulma.SIZE property.<br>SIZE = SIZE_SMALL")).SetHeader(*html.ToHTML("Icecake Message"), false).
		SetSize(bulma.SIZE_SMALL)

	html.Render(out, sec, styl1, styl2)
	html.WriteString(out, `</div>`)

	// states

	return nil
}
