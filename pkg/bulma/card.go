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

	Title       html.HTMLString   // optional title to display on the head of the card
	Image       *Image            // optional image
	Content     html.HTMLComposer // any html content to render within the body of the card
	FooterItem1 html.HTMLComposer // optional Footer 1 of 3 items max
	FooterItem2 html.HTMLComposer // optional Footer 1 of 3 items max
	FooterItem3 html.HTMLComposer // optional Footer 1 of 3 items max
}

// Ensure Card implements HTMLComposer interface
var _ html.HTMLComposer = (*Card)(nil)

// BuildTag builds the tag used to render the html element.
func (card *Card) BuildTag(tag *html.Tag) {
	tag.SetTagName("div").AddClasses("card")
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (card *Card) RenderContent(out io.Writer) error {

	html.RenderHTMLIf(!card.Title.IsEmpty(), out, card, html.HTML(`<header class="card-header">`), html.HTML(`<p class="card-header-title">`), card.Title, html.HTML(`</p></header>`))

	card.RenderChildSnippetIf(card.Image != nil, out, card.Image)

	card.RenderChildSnippetIf(card.Content != nil, out, html.NewSnippet("div", `class="card-content"`).StackContent(card.Content))

	if card.FooterItem1 != nil || card.FooterItem2 != nil || card.FooterItem3 != nil {
		html.WriteString(out, `<div class="card-footer">`)
		card.RenderChildSnippetIf(card.FooterItem1 != nil, out, html.NewSnippet("span", `class="card-footer-item"`).StackContent(card.FooterItem1))
		card.RenderChildSnippetIf(card.FooterItem2 != nil, out, html.NewSnippet("span", `class="card-footer-item"`).StackContent(card.FooterItem2))
		card.RenderChildSnippetIf(card.FooterItem3 != nil, out, html.NewSnippet("span", `class="card-footer-item"`).StackContent(card.FooterItem3))
		html.WriteString(out, `</div>`)
	}
	return nil
}
