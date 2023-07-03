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
	html.WriteString(out, `<h2>Bulma Card</h2>`)
	html.WriteString(out, `<p>bulma.Card is an icecake snippet providing the HTML rendering for a `, linkBulmaCard, ` with extra features and usefull Go APIs.</p>`)

	// api
	html.WriteString(out, `<h3>bulma.Card API</h3>`)
	html.WriteString(out, `<p>The card is an HTMLSnippet. Use 'SetBody()' to setup the body of the card.</p>`)
	html.WriteString(out, `<p><strong>Title HTMLString</strong> Optional title to display in the head of the card. Can be a simple text or a more complex html string.</p>`)
	html.WriteString(out, `<p><strong>Image *Image</strong> Optional image to display on top of the card.</p>`)

	// usages
	html.WriteString(out, `<h3>Usage</h3>`)
	html.WriteString(out, `<div class="box is-flex spaceout mr-5">`)

	card3 := bulma.NewCard()
	card3.AddContent(html.ToHTML(`<p class="title">Very Good Cake</p>`))
	card3.Tag().AddStyle("width: 150px;")
	cmp.RenderChilds(out, card3)

	card1 := bulma.NewCard()
	card1.SetTitle(*html.ToHTML("Hello World")).
		SetImage(*bulma.NewImage("/assets/icecake.jpg", bulma.IMG_128x128)).
		AddFooterItem(*html.ToHTML("<a href='/'>home</a>")).
		AddContent(html.ToHTML("Nice cake"))
	cmp.RenderChilds(out, card1)

	card2 := bulma.NewCard()
	card2.SetImage(*bulma.NewImage("/assets/icecake.jpg", bulma.IMG_64x64)).
		AddContent(html.ToHTML("Nice cake"))
	cmp.RenderChilds(out, card2)

	html.WriteString(out, `</div>`)

	return nil
}
