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
	html.HTMLSnippet

	// Optional title to display in the head of the card
	Title html.HTMLString

	// Optional image to display on top of the card
	Image *ICKImage

	// optional Footer items
	footerItem []html.HTMLString
}

// Ensure Card implements HTMLComposer interface
var _ html.HTMLComposer = (*ICKCard)(nil)

// Card main factory
func Card(content html.HTMLContentComposer) *ICKCard {
	c := new(ICKCard)
	c.footerItem = make([]html.HTMLString, 0)
	c.AddContent(content)
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
		html.WriteString(out, `<header class="card-header">`, `<p class="card-header-title">`)
		card.RenderChild(out, &card.Title)
		html.WriteString(out, `</p></header>`)
	}

	if card.Image != nil {
		html.WriteString(out, `<div class="card-image">`)
		card.RenderChild(out, card.Image)
		html.WriteString(out, `</div>`)
	}

	if card.HasContent() {
		html.WriteString(out, `<div class="card-content">`)
		card.HTMLSnippet.RenderContent(out)
		html.WriteString(out, `</div>`)
	}

	if len(card.footerItem) > 0 {
		html.WriteString(out, `<div class="card-footer">`)
		for _, item := range card.footerItem {
			html.WriteString(out, `<span class="card-footer-item">`)
			card.RenderChild(out, &item)
			html.WriteString(out, `</span>`)
		}
		html.WriteString(out, `</div>`)
	}
	return nil
}
