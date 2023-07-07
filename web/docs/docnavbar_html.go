package docs

import (
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

func DocNavbar() *ick.ICKNavbar {
	nav := &ick.ICKNavbar{HasShadow: true}
	nav.Tag().SetId("topbar")
	nav.AddItem("", ick.NAVBARIT_BRAND, html.ToHTML(`<span class="title pl-2">Icecake</span>`)).ParseHRef("/").ParseImageSrc("/assets/icecake-color.svg")
	nav.AddItem("home", ick.NAVBARIT_START, html.ToHTML(`Home`)).ParseHRef("/")
	nav.AddItem("docs", ick.NAVBARIT_START, html.ToHTML(`Docs`)).ParseHRef("/docoverview.html")

	btngit := ick.Button("GitHub").ParseHRef("https://github.com/icecake-framework/icecake")
	btngit.SetColor(ick.COLOR_PRIMARY).SetOutlined(true)

	nav.AddItem("", ick.NAVBARIT_END, btngit)
	nav.AddItem("", ick.NAVBARIT_END, html.ToHTML("<small>Alpha</small>"))
	return nav
}
