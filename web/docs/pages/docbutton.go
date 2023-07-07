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
	html.Render(out, nil, ick.Title(3, "Button"))
	html.WriteString(out, `<div class="block">`+
		`<p>ICKButton is an icecake snippet providing the HTML rendering for a `, linkBulmaButton, ` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// usages
	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uA1 := ick.Button("Click Link").ParseHRef("#")
	uA2 := ick.Button("Trigger Event").SetId("uA2")
	html.Render(out, cmp, uA1, uA2)
	html.WriteString(out, `</div>`)

	// apis
	html.Render(out, nil, ick.Title(4, "ICKButton APIs"))
	html.WriteString(out, `<div class="block">`+
		`<p><code>Title HTMLString</code> The title of the Button. Can be a simple text or a more complex html string.</p>`+
		`<p><code>HRef *url.URL</code> HRef defines the associated url link. HRef can be nil. If HRef is defined then the rendered element is a &lt;a&gt; tag, otherwise it's a &lt;button&gt; tag.</p>`+
		`</div>`)

	// styling
	html.Render(out, nil, ick.Title(4, "Styling"))
	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uB0 := ick.Button("Default")
	uB1 := ick.Button("Primary color").SetColor(ick.COLOR_PRIMARY)
	uB2 := ick.Button("Light color").SetColor(ick.COLOR_PRIMARY).SetLight(true)
	uB3 := ick.Button("Link color").SetColor(ick.COLOR_LINK)
	uB4 := ick.Button("Outlined").SetColor(ick.COLOR_PRIMARY).SetOutlined(true)
	uB5 := ick.Button("Rounded").SetColor(ick.COLOR_PRIMARY).SetRounded(true)
	uB6 := ick.Button("Small").SetColor(ick.COLOR_PRIMARY).SetSize(ick.SIZE_SMALL)
	uB7 := ick.Button("Large").SetColor(ick.COLOR_PRIMARY).SetSize(ick.SIZE_LARGE)
	html.Render(out, cmp, uB0, uB1, uB2, uB3, uB4, uB5, uB6, uB7)
	html.WriteString(out, `</div>`)

	// states
	html.Render(out, nil, ick.Title(4, "State"))
	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uC1 := ick.Button("Standard").SetColor(ick.COLOR_PRIMARY)
	uC2 := ick.Button("Loading").SetColor(ick.COLOR_PRIMARY).SetLoading(true)
	uC3 := ick.Button("Disabled").SetColor(ick.COLOR_PRIMARY).SetDisabled(true)
	html.Render(out, cmp, uC1, uC2, uC3)
	html.WriteString(out, `</div>`)

	return nil
}
