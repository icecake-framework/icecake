package docs

import (
	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/html"
)

func MyNavbar() *bulma.Navbar {
	nav := &bulma.Navbar{HasShadow: true}
	nav.Tag().SetId("topbar")
	nav.AddItem("", bulma.NAVBARIT_BRAND, html.ToHTML(`<span class="title pl-2">Icecake</span>`)).ParseHRef("/").ParseImageSrc("/assets/icecake-color.svg")
	nav.AddItem("home", bulma.NAVBARIT_START, html.ToHTML(`Home`)).ParseHRef("/")
	nav.AddItem("docs", bulma.NAVBARIT_START, html.ToHTML(`Docs`)).ParseHRef("/docs.html")

	btngit := &bulma.Button{Title: *html.ToHTML("GitHub")}
	btngit.ParseHRef("https://github.com/icecake-framework/icecake")
	btngit.SetColor(bulma.COLOR_PRIMARY).SetOutlined(true)

	nav.AddItem("", bulma.NAVBARIT_END, btngit)
	nav.AddItem("", bulma.NAVBARIT_END, html.ToHTML("<small>Alpha</small>"))
	return nav
}
