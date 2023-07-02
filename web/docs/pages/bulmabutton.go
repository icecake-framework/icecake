package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/html"
)

const (
	linkBulmaButton string = `<a href="https://bulma.io/documentation/components/button">bulma Button</a>`
)

type SectionBulmaButton struct {
	SectionIcecakeDoc
}

func (cmp *SectionBulmaButton) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Icecake Bulma Button</h2>`)
	html.WriteString(out, `<p>bulma.Button is an icecake snippet providing the HTML rendering for a `, linkBulmaButton, ` with extra features and usefull Go APIs.</p>`)

	html.WriteString(out, `<h3>bulma.Button API</h3>`)
	html.WriteString(out, `<p><strong>Title HTMLString</strong> The title of the Button. Can be a simple text or a more complex html string.</p>`)
	html.WriteString(out, `<p><strong>HRef *url.URL</strong> HRef defines the associated url link. HRef can be nil. If HRef is defined then the rendered element is a &lt;a&gt; tag, otherwise it's a &lt;button&gt; tag.</p>`)

	// usage
	html.WriteString(out, `<h3>Usage</h3>`)

	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uA1 := bulma.NewButton(*html.ToHTML("Click Link"), "", "#")
	uA2 := bulma.NewButton(*html.ToHTML("Trigger Event"), "uA2", "")
	html.Render(out, cmp, uA1, uA2)
	html.WriteString(out, `</div>`)

	// stylling
	html.WriteString(out, `<h3>Stylling</h3>`)
	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uB0 := bulma.NewButton(*html.ToHTML("Default"), "", "")
	uB1 := bulma.NewButton(*html.ToHTML("Primary color"), "", "").SetColor(bulma.COLOR_PRIMARY)
	uB2 := bulma.NewButton(*html.ToHTML("Light color"), "", "").SetColor(bulma.COLOR_PRIMARY).SetLight(true)
	uB3 := bulma.NewButton(*html.ToHTML("Link color"), "", "").SetColor(bulma.COLOR_LINK)
	uB4 := bulma.NewButton(*html.ToHTML("Outlined"), "", "").SetColor(bulma.COLOR_PRIMARY).SetOutlined(true)
	uB5 := bulma.NewButton(*html.ToHTML("Rounded"), "", "").SetColor(bulma.COLOR_PRIMARY).SetRounded(true)
	html.Render(out, cmp, uB0, uB1, uB2, uB3, uB4, uB5)
	html.WriteString(out, `</div>`)

	// States
	html.WriteString(out, `<h3>States</h3>`)
	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uC1 := bulma.NewButton(*html.ToHTML("Standard"), "", "").SetColor(bulma.COLOR_PRIMARY)
	uC2 := bulma.NewButton(*html.ToHTML("Loading"), "", "").SetColor(bulma.COLOR_PRIMARY).SetLoading(true)
	uC3 := bulma.NewButton(*html.ToHTML("Disabled"), "", "").SetColor(bulma.COLOR_PRIMARY).SetDisabled(true)
	html.Render(out, cmp, uC1, uC2, uC3)
	html.WriteString(out, `</div>`)

	return nil
}
