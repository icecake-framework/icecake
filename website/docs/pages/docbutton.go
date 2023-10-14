package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

const (
	linkBulmaButton string = `<a href="https://bulma.io/documentation/components/button">bulma Button</a>`
)

type SectionDocButton struct {
	SectionDocIcecake
}

func (sec *SectionDocButton) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Button", "button.go", "ICKButton")
	ickcore.RenderString(out, `<div class="block">`+
		`<p>ICKButton is an icecake snippet providing the HTML rendering for a `, linkBulmaButton, ` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// usages
	ickcore.RenderString(out, `<div class="box spaceout">`)
	ickcore.RenderChild(out, sec,
		ick.Button("Click Link").ParseHRef("/"),
		ick.Button("Trigger Event").SetId("uA2"))
	ickcore.RenderString(out, `</div>`)

	ickcore.RenderChild(out, sec,
		ick.Elem("div", `id="boxusage" class="box"`, ick.Spinner()),
		ick.Button("reset", `class="mb-5"`).
			SetId("btnreset").
			SetColor(ick.COLOR_PRIMARY).
			SetOutlined(true).
			SetDisabled(true))

	// styling
	ickcore.RenderChild(out, sec, ick.Title(4, "Styling"))
	ickcore.RenderString(out, `<div class="box spaceout">`)
	ickcore.RenderChild(out, sec,
		ick.Button("Default"),
		ick.Button("Primary color").SetColor(ick.COLOR_PRIMARY),
		ick.Button("Light color").SetColor(*ick.Color(ick.COLOR_PRIMARY).SetLight(true)),
		ick.Button("Link color").SetColor(ick.COLOR_LINK),
		ick.Button("Outlined").SetColor(ick.COLOR_PRIMARY).SetOutlined(true),
		ick.Button("Rounded").SetColor(ick.COLOR_PRIMARY).SetRounded(true),
		ick.Button("Small").SetSize(ick.SIZE_SMALL),
		ick.Button("Large").SetSize(ick.SIZE_LARGE),
		ick.Button("with opening icons").SetIcon(*ick.Icon("bi bi-check2-square"), false),
		ick.Button("with closing icons").SetIcon(*ick.Icon("bi bi-box-arrow-up-right"), true))
	ickcore.RenderString(out, `</div>`)

	// states
	ickcore.RenderChild(out, sec, ick.Title(4, "State"))
	ickcore.RenderString(out, `<div class="box spaceout">`)
	ickcore.RenderChild(out, sec,
		ick.Button("Standard").SetColor(ick.COLOR_PRIMARY),
		ick.Button("Loading").SetColor(ick.COLOR_PRIMARY).SetLoading(true),
		ick.Button("Disabled").SetColor(ick.COLOR_PRIMARY).SetDisabled(true))
	ickcore.RenderString(out, `</div>`)

	return nil
}
