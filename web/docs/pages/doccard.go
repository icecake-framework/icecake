package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaCard string = `<a href="https://ick.io/documentation/components/card">Bulma Card</a>`
)

type SectionDocCard struct{ SectionDocIcecake }

func (cmp *SectionDocCard) RenderContent(out io.Writer) error {
	html.Render(out, nil, ick.Title(3, "Card"))
	html.WriteString(out, `<div class="block">`+
		`<p>ICKCard is an icecake snippet providing the HTML rendering for a `+linkBulmaCard+` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// usages
	html.WriteString(out, `<div class="box is-flex spaceout mr-5">`)
	u1 := ick.Card(html.ToHTML(`<div class="title">Very Good Cake</div>`))
	u1.Tag().AddStyle("width: 150px;")
	u2 := ick.Card(html.ToHTML("Nice ice cake !")).
		SetTitle(*html.ToHTML("Hello World")).
		SetImage(*ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_R2by1, `class="m-0"`)).
		AddFooterItem(*html.ToHTML("<a href='/'>home</a>"))
	u3 := ick.Card(html.ToHTML("Nice cake")).
		SetImage(*ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_R2by1, `class="m-0"`))
	html.Render(out, nil, u1, u2, u3)
	html.WriteString(out, `</div>`)

	// apis
	html.Render(out, nil, ick.Title(4, "ICKCard APIs"))
	html.WriteString(out, `<div class="block">`+
		`<p>ICKCard embeds an HTMLSnippet. Use <code>AddContent()</code> to setup its content.</p>`+
		`<p><code>Title HTMLString</code> Optional title to display in the head of the card. Can be a simple text or a more complex html string.</p>`+
		`<p><code>Image *Image</code> Optional image to display on top of the card.</p>`+
		`</div>`)

	return nil
}
