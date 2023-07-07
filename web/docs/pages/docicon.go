package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaIcon     string = `<a href="https://bulma.io/documentation/elements/icon">bulma Button</a>`
	hrefICKIcon_Git   string = href_GitPkg + `/ick/icon.go`
	hrefICKIcon_GitUI string = href_GitPkg + `/ick/ickui/icon.go`
	hrefICKIcon_Go    string = href_GoPkg + `/ick#ICKIcon`
	hrefICKIcon_GoUI  string = href_GoPkg + `/ick/ickui#ICKIcon`
)

type SectionDocIcon struct {
	SectionDocIcecake
}

func (sec *SectionDocIcon) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Icon", "icon.go", "ICKIcon")
	html.WriteString(out, `<div class="block">`+
		`<p>ICKIcon is an icecake snippet providing the HTML rendering for a `, linkBulmaIcon, ` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// usages
	html.WriteString(out, `<div class="box spaceout">`)
	// uA1 := ick.Button("Click Link").ParseHRef("#")
	// uA2 := ick.Button("Trigger Event").SetId("uA2")
	// html.Render(out, cmp, uA1, uA2)
	html.WriteString(out, `</div>`)

	// styling
	html.Render(out, nil, ick.Title(4, "Styling"))
	html.WriteString(out, `<div class="box spaceout">`)
	s1 := ick.Icon("bi bi-rocket")
	s2 := ick.Icon("bi bi-rocket", `style="font-size:Smaller;"`)
	s3 := ick.Icon("bi bi-rocket", `style="font-size:larger;"`)
	s4 := ick.Icon("bi bi-rocket").SetText("rocket")
	s5 := ick.Icon("bi bi-rocket", `style="font-size:Smaller;"`).SetText("rocket")
	s6 := ick.Icon("bi bi-rocket", `style="font-size:larger;"`).SetText("rocket")
	s7 := ick.Icon("bi bi-rocket").SetColor(ick.TXTCOLOR_DANGER)
	s8 := ick.Icon("bi bi-rocket").SetColor(ick.TXTCOLOR_SUCCESS)
	s9 := ick.Icon("bi bi-rocket")
	html.Render(out, sec, s1, s2, s3, s4, s5, s6, s7, s8, s9)
	html.WriteString(out, `</div>`)

	return nil
}
