package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/html"
)

const (
	linkBulmaCard string = `<a href="https://bulma.io/documentation/components/card">bulma Card</a>`
)

type SectionBulmaCard struct{ SectionIcecakeDoc }

func (cmp *SectionBulmaCard) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Card</h2>`+
		`<p>bulma.Card is an icecake snippet providing the HTML rendering for a `+linkBulmaCard+` with extra features and usefull Go APIs.</p>`)

	// api
	html.WriteString(out, `<h3>bulma.Card API</h3>`+
		`<p>The Card is an HTMLSnippet. Use <code>AddContent()</code> to setup the content of the card.</p>`+
		`<p><strong>Title HTMLString</strong> Optional title to display in the head of the card. Can be a simple text or a more complex html string.</p>`+
		`<p><strong>Image *Image</strong> Optional image to display on top of the card.</p>`)

	// usages
	html.WriteString(out, `<h3>Usage</h3>`+
		`<div class="box is-flex spaceout mr-5">`)

	u1 := bulma.Card(html.ToHTML(`<p class="title">Very Good Cake</p>`))
	u1.Tag().AddStyle("width: 150px;")
	html.Render(out, nil, u1)

	html.Render(out, nil, bulma.Card(html.ToHTML("Nice cake")).
		SetTitle(*html.ToHTML("Hello World")).
		SetImage(*bulma.NewImage("/assets/icecake.jpg", bulma.IMG_128x128)).
		AddFooterItem(*html.ToHTML("<a href='/'>home</a>")))

	html.Render(out, nil, bulma.Card(html.ToHTML("Nice cake")).
		SetImage(*bulma.NewImage("/assets/icecake.jpg", bulma.IMG_64x64)))

	html.WriteString(out, `</div>`)

	return nil
}
