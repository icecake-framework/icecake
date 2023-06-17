package html

import "io"

func init() {
	RegisterComposer("ick-card", &Card{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type Card struct {
	HTMLSnippet

	Title       HTMLString // optional title to display on the head of the card
	Image       *Image     // optional image
	Content     HTMLString // any html content to render within the body of the card
	FooterItem1 HTMLString // optional Footer 1 of 3 items max
	FooterItem2 HTMLString // optional Footer 1 of 3 items max
	FooterItem3 HTMLString // optional Footer 1 of 3 items max
}

// Ensure Card implements HTMLComposer interface
var _ HTMLComposer = (*Card)(nil)

func (card *Card) Tag() *Tag {
	card.tag.SetName("div").Attributes().AddClasses("card")
	return &card.tag
}

func (card *Card) WriteBody(out io.Writer) error {

	WriteStringsIf(card.Title != "", out, `<header class="card-header">`, `<p class="card-header-title">`, card.Title.String(), `</p>`, `</header>`)

	if card.Image != nil {
		// FIXME
		card.WriteChildSnippet(out, card.Image)
		//io.WriteString(out, card.RenderChildSnippet(card.Image))
	}

	WriteStringsIf(card.Content != "", out, `<div class="card-content">`, string(card.Content), `</div>`)

	if card.FooterItem1 != "" || card.FooterItem2 != "" || card.FooterItem3 != "" {
		WriteString(out, `<div class="card-footer">`)
		WriteStringsIf(card.FooterItem1 != "", out, `<span class="card-footer-item">`, card.FooterItem1.String(), `</span>`)
		WriteStringsIf(card.FooterItem2 != "", out, `<span class="card-footer-item">`, card.FooterItem2.String(), `</span>`)
		WriteStringsIf(card.FooterItem3 != "", out, `<span class="card-footer-item">`, card.FooterItem3.String(), `</span>`)
		WriteString(out, `</div>`)
	}
	return nil
}
