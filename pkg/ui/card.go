package ui

import (
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	html.RegisterComposer("ick-card", &Card{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type Card struct {
	dom.UISnippet

	Title       html.String // optional title to display on the head of the card
	*Image                  // optional image
	Content     html.String // any html content to render within the body of the card
	FooterItem1 html.String // optional Footer 1 of 3 items max
	FooterItem2 html.String // optional Footer 1 of 3 items max
	FooterItem3 html.String // optional Footer 1 of 3 items max
}

func (_card *Card) Template(*html.DataState) (_t html.SnippetTemplate) {
	_t.TagName = "div"
	_t.Attributes = `class="card"`

	if _card.Title != "" {
		_t.Body += `<header class="card-header"><p class="card-header-title">` + _card.Title + `</p>`
		_t.Body += `</header>`
	}

	if _card.Image != nil {
		_t.Body += _card.RenderHTML(_card.Image)
	}

	if _card.Content != "" {
		_t.Body += `<div class="card-content">` + _card.Content + `</div>`
	}

	if _card.FooterItem1 != "" || _card.FooterItem2 != "" || _card.FooterItem3 != "" {
		_t.Body += `<div class="card-footer">`
		if _card.FooterItem1 != "" {
			_t.Body += `<span class="card-footer-item">` + _card.FooterItem1 + `</span>`
		}
		if _card.FooterItem2 != "" {
			_t.Body += `<span class="card-footer-item">` + _card.FooterItem2 + `</span>`
		}
		if _card.FooterItem3 != "" {
			_t.Body += `<span class="card-footer-item">` + _card.FooterItem3 + `</span>`
		}
		_t.Body += `</div>`
	}

	return _t
}
