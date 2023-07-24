package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

func init() {
	ickcore.RegisterComposer("ick-card", &ICKCard{})
}

// The card is an HTMLSnippet. Use AddContent to setup the content of the card
type ICKCard struct {
	ickcore.BareSnippet

	// Optional title to display in the head of the card
	Title ickcore.HTMLString

	// the body of the card
	Body ICKElem // rendered as <div class="card-content">

	// Optional image to display on top of the card
	Image *ICKImage

	// optional Footer items
	footerItem []ickcore.HTMLString
}

// Ensuring ICKCard implements the right interface
var _ ickcore.ContentComposer = (*ICKCard)(nil)
var _ ickcore.TagBuilder = (*ICKCard)(nil)

// Card main factory
func Card(content ickcore.ContentComposer, attrs ...string) *ICKCard {
	c := new(ICKCard)
	c.Tag().ParseAttributes(attrs...)
	c.footerItem = make([]ickcore.HTMLString, 0)
	c.Body.Append(content)
	return c
}

func (card *ICKCard) SetTitle(title string) *ICKCard {
	card.Title = *ickcore.ToHTML(title)
	return card
}
func (card *ICKCard) SetImage(image ICKImage) *ICKCard {
	card.Image = &image
	return card
}
func (card *ICKCard) AddFooterItem(item ickcore.HTMLString) *ICKCard {
	card.footerItem = append(card.footerItem, item)
	return card
}

/******************************************************************************/

// BuildTag builds the tag used to render the html element.
// Card Tag is a simple <div class="card"></div>
func (card *ICKCard) BuildTag() ickcore.Tag {
	card.Tag().SetTagName("div").AddClass("card")
	card.Body.Tag().SetTagName("div").AddClass("card-content")
	return *card.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// Card rendering renders the optional header withe the Title, the optional Image, the content, and a slice of footers
func (card *ICKCard) RenderContent(out io.Writer) error {

	if card.Title.NeedRendering() {
		ickcore.RenderString(out, `<header class="card-header">`, `<p class="card-header-title">`)
		ickcore.RenderChild(out, card, &card.Title)
		ickcore.RenderString(out, `</p></header>`)
	}

	if card.Image != nil {
		ickcore.RenderString(out, `<div class="card-image">`)
		ickcore.RenderChild(out, card, card.Image)
		ickcore.RenderString(out, `</div>`)
	}

	ickcore.RenderChild(out, card, &card.Body)

	if len(card.footerItem) > 0 {
		ickcore.RenderString(out, `<div class="card-footer">`)
		for _, item := range card.footerItem {
			ickcore.RenderString(out, `<span class="card-footer-item">`)
			ickcore.RenderChild(out, card, &item)
			ickcore.RenderString(out, `</span>`)
		}
		ickcore.RenderString(out, `</div>`)
	}
	return nil
}
