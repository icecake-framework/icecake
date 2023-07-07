package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaCard string = `<a href="https://bulma.io/documentation/components/card">Bulma Card</a>`
)

type SectionDocCard struct{ SectionDocIcecake }

func (sec *SectionDocCard) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Card", "card.go", "ICKCard")
	html.WriteString(out, `<div class="block">`+
		`<p>ICKCard is an icecake snippet providing the HTML rendering for a `+linkBulmaCard+` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// usages
	html.WriteString(out, `<div class="box is-flex spaceout">`)
	u1 := ick.Card(html.ToHTML(`<div class="title">Very Good Cake</div>`))
	u1.Tag().AddStyle("width: 150px;")
	u2 := ick.Card(html.ToHTML("Nice ice cake !")).
		SetTitle("Hello World").
		SetImage(*ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_R2by1, `class="m-0"`)).
		AddFooterItem(*html.ToHTML("<a href='/'>home</a>"))
	u3 := ick.Card(html.ToHTML("Nice cake")).
		SetImage(*ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_R2by1, `class="m-0"`))
	html.Render(out, nil, u1, u2, u3)
	html.WriteString(out, `</div>`)

	return nil
}
