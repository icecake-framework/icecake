package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

const (
	linkBulmaCard string = `<a href="https://bulma.io/documentation/components/card">Bulma Card</a>`
)

type SectionDocCard struct{ SectionDocIcecake }

func (sec *SectionDocCard) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Card", "card.go", "ICKCard")
	ickcore.RenderString(out, `<div class="block">`+
		`<p>ICKCard is an icecake snippet providing the HTML rendering for a `+linkBulmaCard+` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// usages
	ickcore.RenderString(out, `<div class="box is-flex spaceout">`)
	ickcore.RenderChild(out, sec,
		ick.Card(ickcore.ToHTML(`<div class="title">Very Good Cake</div>`), `style="width: 150px;"`),
		ick.Card(ickcore.ToHTML("Nice ice cake !")).
			SetTitle("Hello World").
			SetImage(*ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_R2by1, `class="m-0"`)).
			AddFooterItem(*ickcore.ToHTML("<a href='/'>home</a>")),
		ick.Card(ickcore.ToHTML("Nice cake")).
			SetImage(*ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_R2by1, `class="m-0"`)))
	ickcore.RenderString(out, `</div>`)

	return nil
}
