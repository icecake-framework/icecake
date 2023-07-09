package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaNavbar string = `<a href="https://bulma.io/documentation/components/navbar">bulma navbar</a>`
)

type SectionDocNavbar struct{ SectionDocIcecake }

func (sec *SectionDocNavbar) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Navbar", "Navbar.go", "ICKNavbar")

	html.RenderString(out, `<div class="block">`+
		`<p>ICKNavbar is an icecake snippet providing the HTML rendering for a `, linkBulmaNavbar, `, with extra features and usefull go wasm APIs.</p>`+
		`<p>The navbar is an horizontal bar containing items and sub items.`+
		`The navbar is splitted in three areas: the brand area, the start area stacked on the left, and the end area stacked on the right.</p>`+
		`</div>`)

	// usages
	html.RenderString(out, `<div id="boxusage" class="box mr-5">`)
	bar := ick.NavBar(`style="border: solid 1px;"`)
	bar.AddItem("", ick.NAVBARIT_BRAND, html.ToHTML("BRAND"))
	bar.AddItem("", ick.NAVBARIT_START, html.ToHTML("Home"))
	bar.AddItem("", ick.NAVBARIT_START, html.ToHTML("Second Item"))
	bar.AddItem("", ick.NAVBARIT_END, html.ToHTML("Last Item"))
	html.RenderChild(out, sec, bar)
	html.RenderString(out, `</div>`)

	//Styling
	html.RenderChild(out, sec, ick.Title(4, "Styling"))
	html.RenderString(out, `<div class="block">`+
		`</div>`)

	return nil
}
