package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaNavbar string = `<a href="https://ick.io/documentation/components/navbar">bulma navbar</a>`
)

type SectionDocNavbar struct{ SectionDocIcecake }

func (cmp *SectionDocNavbar) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Navbar</h2>`+
		`<p>ICKNavbar is an icecake snippet providing the HTML rendering for a `, linkBulmaNavbar, `, with extra features and usefull go wasm APIs.</p>`+
		`<p>The navbar is an horizontal bar containing items and sub items.`+
		`The navbar is splitted in three areas: the brand area, the start area stacked on the left, and the end area stacked on the right.</p>`)

	// usages
	html.WriteString(out, `<div id="boxusage" class="box mr-5">`)
	us1 := new(ick.ICKNavbar)
	us1.Tag().AddStyle("border: solid 1px;")
	us1.AddItem("", ick.NAVBARIT_BRAND, html.ToHTML("BRAND"))
	us1.AddItem("", ick.NAVBARIT_START, html.ToHTML("Home"))
	us1.AddItem("", ick.NAVBARIT_START, html.ToHTML("Second Item"))
	us1.AddItem("", ick.NAVBARIT_END, html.ToHTML("Last Item"))
	html.Render(out, cmp, us1)
	html.WriteString(out, `</div>`)

	html.WriteString(out, `<h3>ICKNavbar APIs</h3>`+
		`<p><strong>IsTransparent</strong> renders a transparent navbar</p>`+
		`<p><strong>HasShadow</strong> renders a shadow below the navbar</p>`)

	html.WriteString(out, `<h3>Styling Properties</h3>`)

	return nil
}
