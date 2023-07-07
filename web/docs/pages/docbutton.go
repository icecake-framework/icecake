package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaButton string = `<a href="https://bulma.io/documentation/components/button">bulma Button</a>`
)

type SectionDocButton struct {
	SectionDocIcecake
}

func (sec *SectionDocButton) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Button", "button.go", "ICKButton")
	html.WriteString(out, `<div class="block">`+
		`<p>ICKButton is an icecake snippet providing the HTML rendering for a `, linkBulmaButton, ` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// usages
	html.WriteString(out, `<div class="box spaceout">`)
	uA1 := ick.Button("Click Link").ParseHRef("/")
	uA2 := ick.Button("Trigger Event").SetId("uA2")
	html.Render(out, sec, uA1, uA2)
	html.WriteString(out, `</div>`)

	// styling
	html.Render(out, nil, ick.Title(4, "Styling"))
	html.WriteString(out, `<div class="box spaceout">`)
	s0 := ick.Button("Default")
	s1 := ick.Button("Primary color").SetColor(ick.COLOR_PRIMARY)
	s2 := ick.Button("Light color").SetColor(*ick.Color(ick.COLOR_PRIMARY).SetLight(true))
	s3 := ick.Button("Link color").SetColor(ick.COLOR_LINK)
	s4 := ick.Button("Outlined").SetColor(ick.COLOR_PRIMARY).SetOutlined(true)
	s5 := ick.Button("Rounded").SetColor(ick.COLOR_PRIMARY).SetRounded(true)
	s6 := ick.Button("Small").SetSize(ick.SIZE_SMALL)
	s7 := ick.Button("Large").SetSize(ick.SIZE_LARGE)
	s8 := ick.Button("with opening icons").SetIcon(*ick.Icon("bi bi-check2-square"), false)
	s9 := ick.Button("with closing icons").SetIcon(*ick.Icon("bi bi-box-arrow-up-right"), true)
	html.Render(out, sec, s0, s1, s2, s3, s4, s5, s6, s7, s8, s9)
	html.WriteString(out, `</div>`)

	// states
	html.Render(out, nil, ick.Title(4, "State"))
	html.WriteString(out, `<div class="box spaceout">`)
	st1 := ick.Button("Standard").SetColor(ick.COLOR_PRIMARY)
	st2 := ick.Button("Loading").SetColor(ick.COLOR_PRIMARY).SetLoading(true)
	st3 := ick.Button("Disabled").SetColor(ick.COLOR_PRIMARY).SetDisabled(true)
	html.Render(out, sec, st1, st2, st3)
	html.WriteString(out, `</div>`)

	return nil
}
