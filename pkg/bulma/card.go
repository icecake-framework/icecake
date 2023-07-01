package bulma

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-card", &Card{})
}

type Card struct {
	html.HTMLSnippet

	Title   html.HTMLString   // optional title to display on the head of the card
	Image   *Image            // optional image
	Content html.HTMLComposer // any html content to render within the body of the card

	FooterItem []html.HTMLString // optional Footer 1 of 3 items max
}

// Ensure Card implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*Card)(nil)

// BuildTag builds the tag used to render the html element.
// Card Tag is a simple <div class="card"></div>
func (card *Card) BuildTag(tag *html.Tag) {
	tag.SetTagName("div").AddClasses("card")
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// Card rendering renders the optional header withe the Title, the optional Image, the content, and a slice of footers
func (card *Card) RenderContent(out io.Writer) error {

	if !card.Title.IsEmpty() {
		html.WriteStrings(out, `<header class="card-header">`, `<p class="card-header-title">`)
		card.RenderChilds(out, &card.Title)
		html.WriteString(out, `</p></header>`)
	}

	card.RenderChildsIf(card.Image != nil, out, card.Image)

	if card.Content != nil {
		html.WriteString(out, `<div class="card-content">`)
		card.RenderChilds(out, card.Content)
		html.WriteString(out, `</div>`)
	}

	if len(card.FooterItem) > 0 {
		html.WriteString(out, `<div class="card-footer">`)
		for _, item := range card.FooterItem {
			html.WriteString(out, `<span class="card-footer-item">`)
			card.RenderChilds(out, &item)
			html.WriteString(out, `</span>`)
		}
		html.WriteString(out, `</div>`)
	}
	return nil
}
