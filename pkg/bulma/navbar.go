package bulma

import (
	"io"
	"net/url"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-navbar", &Navbar{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type NAVBARITEM_TYPE string

const (
	NAVBARIT_DIVIDER = "divider" // creates a divider between two navbar items. other navbar properties are ignored.
	NAVBARIT_BRAND   = "brand"   // item stacked in the brand area of the navbar.
	NAVBARIT_START   = "start"   // item stacked at the start (left) of the navbar, this is the default behaviour
	NAVBARIT_END     = "end"     // item stacked at the end (right) of the navbar,
)

type NavbarItem struct {
	html.HTMLSnippet

	// The Item Type defines the location of the item in the navbar or if it's a simple divider.
	// If ItemType is empty, NAVBARIT_START is used for rendering.
	ItemType NAVBARITEM_TYPE

	// HRef defines the optional associated url link.
	// If HRef is defined the item become an anchor link <a>, otherwise it's a <div>
	// HRef can be nil. Usually it's created calling NavbarItem.TryParseHRef
	HRef *url.URL

	// ImageSrc defines an optional image to display at the begining of the Item
	ImageSrc *url.URL // the url for the source of the image

	// Item Content
	Content html.HTMLComposer

	items []*NavbarItem // list of navbar items
}

// Ensure NavbarItem implements HTMLComposer interface
var _ html.HTMLComposer = (*NavbarItem)(nil)

// BuildTag builds the tag used to render the html element.
func (item *NavbarItem) BuildTag(tag *html.Tag) {
	if item.ItemType == NAVBARIT_DIVIDER {
		tag.SetTagName("hr")
	} else {
		if item.HRef != nil && item.HRef.String() != "" {
			tag.SetTagName("a")
		} else {
			tag.SetTagName("div")
		}
	}

	amap := tag.AttributeMap
	if item.ItemType == NAVBARIT_DIVIDER {
		amap.AddClasses("navbar-divider")
	} else {
		amap.AddClasses("navbar-item")
		if item.HRef != nil {
			if href := item.HRef.String(); href != "" {
				amap.SetAttribute("href", href)
			}
		}
	}
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (item *NavbarItem) RenderContent(out io.Writer) error {
	if item.ItemType != NAVBARIT_DIVIDER {
		if item.ImageSrc != nil {
			imgsrc := item.ImageSrc.String()
			html.WriteStringsIf(imgsrc != "", out, `<img src="`, imgsrc, `" width="auto" height="28">`)
		}
		item.RenderChildSnippet(out, item.Content)
	}
	return nil
}

func (item *NavbarItem) AddItems(items ...*NavbarItem) *NavbarItem {
	item.items = append(item.items, items...)
	return item
}

// ParseHRef tries to parse rawUrl to HRef ignoring error.
func (item *NavbarItem) ParseHRef(rawUrl string) *NavbarItem {
	item.HRef, _ = url.Parse(rawUrl)
	return item
}

// ParseImageSrc tries to parse rawUrl to image src ignoring error.
func (item *NavbarItem) ParseImageSrc(rawUrl string) *NavbarItem {
	item.ImageSrc, _ = url.Parse(rawUrl)
	return item
}

// Navbar is an UISnippet registered with the ick-tag `ick-navbar`.
type Navbar struct {
	html.HTMLSnippet

	items []*NavbarItem // list of navbar items

	// TODO: handle Navbar active item

	// Styling
	IsTransparent bool
	HasShadow     bool
}

// Ensure Navbar implements HTMLComposer interface
var _ html.HTMLComposer = (*Navbar)(nil)

// AddItems adds items to the navbar .
func (_nav *Navbar) AddItems(items ...*NavbarItem) *Navbar {
	_nav.items = append(_nav.items, items...)
	return _nav
}

// BuildTag builds the tag used to render the html element.
func (nav *Navbar) BuildTag(tag *html.Tag) {
	tag.SetTagName("nav").
		SetAttribute("role", "navigation").
		AddClasses("navbar").
		AddClassesIf(nav.IsTransparent, "is-transparent").
		AddClassesIf(nav.HasShadow, "has-shadow")
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (nav *Navbar) RenderContent(out io.Writer) error {
	// brand area
	html.WriteString(out, `<div class="navbar-brand">`)

	// brand items
	for _, item := range nav.items {
		nav.RenderChildSnippetIf(item.ItemType == NAVBARIT_BRAND, out, item)
	}
	// burger
	html.WriteStrings(out, `<a class="navbar-burger" role="button">`, `<span></span><span></span><span></span>`, `</a>`)
	html.WriteString(out, `</div>`)

	// menu area
	menuid := nav.Id() + `menu`
	html.WriteStrings(out, `<div class="navbar-menu" id="`, menuid, `">`)

	html.WriteStrings(out, `<div class="navbar-start">`)
	for _, item := range nav.items {
		nav.RenderChildSnippetIf(item.ItemType == NAVBARIT_START, out, item)
	}
	html.WriteString(out, `</div>`)

	html.WriteStrings(out, `<div class="navbar-end">`)
	for _, item := range nav.items {
		nav.RenderChildSnippetIf(item.ItemType == NAVBARIT_END, out, item)
	}
	html.WriteString(out, `</div>`)

	html.WriteString(out, `</div>`) // navbar-menu

	return nil
}

// func (_nav *Navbar) Template(*DataState) (_t SnippetTemplate) {
// 	_t.TagName = "nav"
// 	_t.Attributes = `class="navbar" role="navigation"`
// 	if _nav.IsTransparent {
// 		_nav.SetClasses("is-transparent")
// 	}
// 	if _nav.HasShadow {
// 		_nav.SetClasses("has-shadow")
// 	}

// 	// brand
// 	_t.Body = `<div class="navbar-brand">`

// 	for _, item := range _nav.items {
// 		if item.Brand {
// 			_t.Body += _nav.RenderChildSnippet(&item)
// 		}
// 	}

// 	// _t.Body += `<a class="navbar-item" href="https://bulma.io">
// 	// 			<img src="https://bulma.io/images/bulma-logo.png" width="112" height="28">
// 	// 			</a>`

// 	// burger
// 	_t.Body += `<a class="navbar-burger" role="button">`
// 	_t.Body += `<span></span><span></span><span></span>`
// 	_t.Body += `</a>`

// 	_t.Body += `</div>` //brand

// 	// menu
// 	_t.Body += `<div class="navbar-menu" id="` + String(_nav.Id()) + `menu">`

// 	_t.Body += `<div class="navbar-start">`
// 	for _, item := range _nav.items {
// 		if !item.Brand && !item.End {
// 			_t.Body += _nav.RenderChildSnippet(&item)
// 		}
// 	}
// 	_t.Body += `</div>`

// 	_t.Body += `<div class="navbar-end">`
// 	for _, item := range _nav.items {
// 		if !item.Brand && item.End {
// 			_t.Body += _nav.RenderChildSnippet(&item)
// 		}
// 	}
// 	_t.Body += `</div>`

// 	_t.Body += `</div>` // navbar-menu

// 	return _t
// }
