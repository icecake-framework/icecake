package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaButton string = `<a href="https://ick.io/documentation/components/button">bulma Button</a>`
)

type SectionDocButton struct {
	SectionDocIcecake
}

func (cmp *SectionDocButton) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Button</h2>`)
	html.WriteString(out, `<p>ick.Button is an icecake snippet providing the HTML rendering for a `, linkBulmaButton, ` with extra features and usefull Go APIs.</p>`)

	// usages
	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uA1 := ick.Button(*html.ToHTML("Click Link"), "", "#")
	uA2 := ick.Button(*html.ToHTML("Trigger Event"), "uA2", "")
	html.Render(out, cmp, uA1, uA2)
	html.WriteString(out, `</div>`)

	// apis
	html.WriteString(out, `<h3>ick.Button API</h3>`)
	html.WriteString(out, `<p><strong>Title HTMLString</strong> The title of the Button. Can be a simple text or a more complex html string.</p>`)
	html.WriteString(out, `<p><strong>HRef *url.URL</strong> HRef defines the associated url link. HRef can be nil. If HRef is defined then the rendered element is a &lt;a&gt; tag, otherwise it's a &lt;button&gt; tag.</p>`)

	// styling
	html.WriteString(out, `<h3>Styling</h3>`)
	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uB0 := ick.Button(*html.ToHTML("Default"), "", "")
	uB1 := ick.Button(*html.ToHTML("Primary color"), "", "").SetColor(ick.COLOR_PRIMARY)
	uB2 := ick.Button(*html.ToHTML("Light color"), "", "").SetColor(ick.COLOR_PRIMARY).SetLight(true)
	uB3 := ick.Button(*html.ToHTML("Link color"), "", "").SetColor(ick.COLOR_LINK)
	uB4 := ick.Button(*html.ToHTML("Outlined"), "", "").SetColor(ick.COLOR_PRIMARY).SetOutlined(true)
	uB5 := ick.Button(*html.ToHTML("Rounded"), "", "").SetColor(ick.COLOR_PRIMARY).SetRounded(true)
	html.Render(out, cmp, uB0, uB1, uB2, uB3, uB4, uB5)
	html.WriteString(out, `</div>`)

	// states
	html.WriteString(out, `<h3>States</h3>`)
	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uC1 := ick.Button(*html.ToHTML("Standard"), "", "").SetColor(ick.COLOR_PRIMARY)
	uC2 := ick.Button(*html.ToHTML("Loading"), "", "").SetColor(ick.COLOR_PRIMARY).SetLoading(true)
	uC3 := ick.Button(*html.ToHTML("Disabled"), "", "").SetColor(ick.COLOR_PRIMARY).SetDisabled(true)
	html.Render(out, cmp, uC1, uC2, uC3)
	html.WriteString(out, `</div>`)

	return nil
}
