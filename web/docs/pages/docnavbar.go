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
	html.WriteString(out, `<h2>Navbar</h2>`)
	html.WriteString(out, `<p>ICKNavbar is an icecake snippet providing the HTML rendering for a `, linkBulmaNavbar, `, with extra features and usefull go wasm APIs.</p>`)
	html.WriteString(out, `<p>The navbar is an horizontal bar containing items and sub items.`)
	html.WriteString(out, `The navbar is splitted in three areas: the brand area, the start area stacked on the left, and the end area stacked on the right.</p>`)

	// example 1
	ex1 := new(ick.ICKNavbar)
	ex1.Tag().AddStyle("border: solid 1px;")
	ex1.AddItem("", ick.NAVBARIT_BRAND, html.ToHTML("BRAND"))
	ex1.AddItem("", ick.NAVBARIT_START, html.ToHTML("Home"))
	ex1.AddItem("", ick.NAVBARIT_START, html.ToHTML("Second Item"))
	ex1.AddItem("", ick.NAVBARIT_END, html.ToHTML("Last Item"))
	html.Render(out, cmp, ex1)

	html.WriteString(out, `<h3>ICKNavbar APIs</h3>`)

	html.WriteString(out, `<p><strong>IsTransparent</strong> renders a transparent navbar</p>`)
	html.WriteString(out, `<p><strong>HasShadow</strong> renders a shadow below the navbar</p>`)

	html.WriteString(out, `<h3>Styling Properties</h3>`)

	return nil
}
