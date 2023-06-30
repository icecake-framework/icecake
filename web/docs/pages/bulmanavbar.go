package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/html"
)

const (
	linkBulmaNavbar string = `<a href="https://bulma.io/documentation/components/navbar">bulma navbar</a>`
)

type SectionBulmaNavbar struct{ SectionIcecakeDoc }

func (cmp *SectionBulmaNavbar) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Navbar</h2>`)
	html.WriteStrings(out, `<p>bulma.Navbar is an icecake snippet providing the HTML rendering for a `, linkBulmaNavbar, `.</p>`)
	html.WriteStrings(out, `<p>The navbar is an horizontal bar containing items and sub items.`)
	html.WriteStrings(out, `The navbar is splitted in three areas: the brand area, the start area stacked on the left, and the end area stacked on the right.</p>`)

	navex1 := bulma.Navbar{}
	navex1.Tag().SetStyle("border: solid 1px;")
	navex1.AddItem("", bulma.NAVBARIT_BRAND, html.ToHTML("BRAND"))
	navex1.AddItem("", bulma.NAVBARIT_START, html.ToHTML("Home"))
	navex1.AddItem("", bulma.NAVBARIT_START, html.ToHTML("Second Item"))
	navex1.AddItem("", bulma.NAVBARIT_END, html.ToHTML("Last Item"))
	html.Render(out, cmp, &navex1)

	html.WriteString(out, `<h3>Styling Properties</h3>`)

	html.WriteString(out, `<p><strong>IsTransparent</strong> renders a transparent navbar</p>`)
	html.WriteString(out, `<p><strong>HasShadow</strong> renders a shadow below the navbar</p>`)

	html.WriteString(out, `<h3>Styling Properties</h3>`)

	return nil
}
