package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-card", &ICKCard{})
}

// The card is an HTMLSnippet. Use AddContent to setup the content of the card
type ICKCard struct {
	html.BareSnippet

	// Optional title to display in the head of the card
	Title html.HTMLString

	// the body of the card
	Body html.ContentStack

	// Optional image to display on top of the card
	Image *ICKImage

	// optional Footer items
	footerItem []html.HTMLString
}

// Ensuring ICKCard implements the right interface
var _ html.ElementComposer = (*ICKCard)(nil)

// Card main factory
func Card(content html.ContentComposer, attrs ...string) *ICKCard {
	c := new(ICKCard)
	c.Tag().ParseAttributes(attrs...)
	c.footerItem = make([]html.HTMLString, 0)
	c.Body.Push(content)
	return c
}

func (card *ICKCard) SetTitle(title string) *ICKCard {
	card.Title = *html.ToHTML(title)
	return card
}
func (card *ICKCard) SetImage(image ICKImage) *ICKCard {
	card.Image = &image
	return card
}
func (card *ICKCard) AddFooterItem(item html.HTMLString) *ICKCard {
	card.footerItem = append(card.footerItem, item)
	return card
}

/******************************************************************************/

// BuildTag builds the tag used to render the html element.
// Card Tag is a simple <div class="card"></div>
func (card *ICKCard) BuildTag() html.Tag {
	card.Tag().SetTagName("div").AddClass("card")
	return *card.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// Card rendering renders the optional header withe the Title, the optional Image, the content, and a slice of footers
func (card *ICKCard) RenderContent(out io.Writer) error {

	if !card.Title.IsEmpty() {
		html.RenderString(out, `<header class="card-header">`, `<p class="card-header-title">`)
		html.RenderChild(out, card, &card.Title)
		html.RenderString(out, `</p></header>`)
	}

	if card.Image != nil {
		html.RenderString(out, `<div class="card-image">`)
		html.RenderChild(out, card, card.Image)
		html.RenderString(out, `</div>`)
	}

	if card.Body.HasContent() {
		html.RenderString(out, `<div class="card-content">`)
		card.Body.RenderStack(out, card)
		html.RenderString(out, `</div>`)
	}

	if len(card.footerItem) > 0 {
		html.RenderString(out, `<div class="card-footer">`)
		for _, item := range card.footerItem {
			html.RenderString(out, `<span class="card-footer-item">`)
			html.RenderChild(out, card, &item)
			html.RenderString(out, `</span>`)
		}
		html.RenderString(out, `</div>`)
	}
	return nil
}
