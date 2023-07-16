package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

const (
	linkBulmaIcon string = `<a href="https://bulma.io/documentation/elements/icon">bulma Icon</a>`
)

type SectionDocIcon struct {
	SectionDocIcecake
}

func (sec *SectionDocIcon) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Icon", "icon.go", "ICKIcon")

	ickcore.RenderString(out, `<div class="block">`+
		`<p>ICKIcon is an icecake snippet providing the HTML rendering for a `, linkBulmaIcon, ` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// styling
	ickcore.RenderChild(out, sec, ick.Title(4, "Styling"))
	ickcore.RenderString(out, `<div class="box spaceout">`)
	ickcore.RenderChild(out, sec,
		ick.Icon("bi bi-rocket"),
		ick.Icon("bi bi-rocket", `style="font-size:Smaller;"`),
		ick.Icon("bi bi-rocket", `style="font-size:larger;"`),
		ick.Icon("bi bi-rocket").SetText("rocket"),
		ick.Icon("bi bi-rocket", `style="font-size:Smaller;"`).SetText("rocket"),
		ick.Icon("bi bi-rocket", `style="font-size:larger;"`).SetText("rocket"),
		ick.Icon("bi bi-rocket").SetColor(ick.TXTCOLOR_DANGER),
		ick.Icon("bi bi-rocket").SetColor(ick.TXTCOLOR_SUCCESS),
		ick.Icon("bi bi-rocket"))
	ickcore.RenderString(out, `</div>`)

	return nil
}
