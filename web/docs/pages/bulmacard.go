package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaCard string = `<a href="https://ick.io/documentation/components/card">Card</a>`
)

type SectionBulmaCard struct{ SectionIcecakeDoc }

func (cmp *SectionBulmaCard) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Ick Card</h2>`+
		`<p>ick.Card is an icecake snippet providing the HTML rendering for a `+linkBulmaCard+` with extra features and usefull Go APIs.</p>`)

	// api
	html.WriteString(out, `<h3>ick.Card API</h3>`+
		`<p>The Card is an HTMLSnippet. Use <code>AddContent()</code> to setup the content of the card.</p>`+
		`<p><strong>Title HTMLString</strong> Optional title to display in the head of the card. Can be a simple text or a more complex html string.</p>`+
		`<p><strong>Image *Image</strong> Optional image to display on top of the card.</p>`)

	// usages
	html.WriteString(out, `<h3>Usage</h3>`+
		`<div class="box is-flex spaceout mr-5">`)

	u1 := ick.Card(html.ToHTML(`<div class="title">Very Good Cake</div>`))
	u1.Tag().AddStyle("width: 150px;")
	html.Render(out, nil, u1)

	html.Render(out, nil, ick.Card(html.ToHTML("Nice ice cake !")).
		SetTitle(*html.ToHTML("Hello World")).
		SetImage(*ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_R2by1, `class="m-0"`)).
		AddFooterItem(*html.ToHTML("<a href='/'>home</a>")))

	html.Render(out, nil, ick.Card(html.ToHTML("Nice cake")).
		SetImage(*ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_R2by1, `class="m-0"`)))

	html.WriteString(out, `</div>`)

	return nil
}
