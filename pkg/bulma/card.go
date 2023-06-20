package bulma

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-card", &Card{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type Card struct {
	html.HTMLSnippet

	Title       html.HTMLString // optional title to display on the head of the card
	Image       *Image          // optional image
	Content     html.HTMLString // any html content to render within the body of the card
	FooterItem1 html.HTMLString // optional Footer 1 of 3 items max
	FooterItem2 html.HTMLString // optional Footer 1 of 3 items max
	FooterItem3 html.HTMLString // optional Footer 1 of 3 items max
}

// Ensure Card implements HTMLComposer interface
var _ html.HTMLComposer = (*Card)(nil)

func (card *Card) BuildTag(tag *html.Tag) {
	tag.SetName("div").Attributes().AddClasses("card")
}

func (card *Card) RenderContent(out io.Writer) error {

	html.WriteStringsIf(!card.Title.IsEmpty(), out, `<header class="card-header">`, `<p class="card-header-title">`, card.Title.String(), `</p>`, `</header>`)

	if card.Image != nil {
		// FIXME
		card.RenderChildSnippet(out, card.Image)
		//io.WriteString(out, card.RenderChildSnippet(card.Image))
	}

	html.WriteStringsIf(!card.Content.IsEmpty(), out, `<div class="card-content">`, card.Content.String(), `</div>`)

	if !card.FooterItem1.IsEmpty() || !card.FooterItem2.IsEmpty() || !card.FooterItem3.IsEmpty() {
		// FIXME
		html.WriteString(out, `<div class="card-footer">`)
		html.WriteStringsIf(!card.FooterItem1.IsEmpty(), out, `<span class="card-footer-item">`, card.FooterItem1.String(), `</span>`)
		html.WriteStringsIf(!card.FooterItem2.IsEmpty(), out, `<span class="card-footer-item">`, card.FooterItem2.String(), `</span>`)
		html.WriteStringsIf(!card.FooterItem3.IsEmpty(), out, `<span class="card-footer-item">`, card.FooterItem3.String(), `</span>`)
		html.WriteString(out, `</div>`)
	}
	return nil
}
