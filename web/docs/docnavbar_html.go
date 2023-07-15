package docs

import (
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

func DocNavbar(pg *ick.Page) *ick.ICKNavbar {
	nav := &ick.ICKNavbar{HasShadow: true}
	nav.Tag().SetId("topbar")
	nav.AddItem("", ick.NAVBARIT_BRAND, ickcore.ToHTML(`<span class="title pl-2">Icecake</span>`)).
		SetHRef(*pg.ToAbsURL("/")).
		SetImageSrc(*pg.ToAbsURL("/assets/icecake-color.svg"))
	nav.AddItem("home", ick.NAVBARIT_START, ickcore.ToHTML(`Home`)).
		SetHRef(*pg.ToAbsURL("/"))
	nav.AddItem("docs", ick.NAVBARIT_START, ickcore.ToHTML(`Docs`)).
		SetHRef(*pg.ToAbsURL("/docoverview.html"))

	btngit := ick.Button("GitHub").ParseHRef("https://github.com/icecake-framework/icecake")
	btngit.SetColor(ick.COLOR_LINK).SetOutlined(true)

	nav.AddItem("", ick.NAVBARIT_END, btngit)
	nav.AddItem("", ick.NAVBARIT_END, ickcore.ToHTML("<small>Alpha</small>"))
	return nav
}
