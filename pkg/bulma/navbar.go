package bulma

import (
	"io"
	"net/url"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-navbar", &Navbar{})
}

type NAVBARITEM_TYPE string

const (
	NAVBARIT_DIVIDER = "divider" // creates a divider between two navbar items. other navbar properties are ignored.
	NAVBARIT_BRAND   = "brand"   // item stacked in the brand area of the navbar.
	NAVBARIT_START   = "start"   // item stacked at the start (left) of the navbar, this is the default behaviour
	NAVBARIT_END     = "end"     // item stacked at the end (right) of the navbar,
)

// bulma.NavbarItem is an icecake snippet providing the HTML rendering for a [bulma navbar item].
//
// Can't be used for inline rendering.
//
// [bulma navbar item]: https://bulma.io/documentation/components/navbar/#navbar-item
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

	// Highlight this item
	IsActive bool

	items []*NavbarItem // list of navbar items
}

// Ensure NavbarItem implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*NavbarItem)(nil)

// BuildTag builds the tag used to render the html element.
// The Navbar Item tag depends on the item properties:
//   - it's <hr> for a NAVBARIT_DIVIDER item type, otherwise
//   - it's <a> when an HRef is provided,
//   - it's <div> in other cases
func (item *NavbarItem) BuildTag(tag *html.Tag) {
	if item.ItemType == NAVBARIT_DIVIDER {
		tag.SetTagName("hr").AddClasses("navbar-divider")
	} else {
		tag.AddClasses("navbar-item").SetClassesIf(item.IsActive, "is-active")
		if item.HRef != nil {
			tag.SetTagName("a").SetURL("href", item.HRef)
		} else {
			tag.SetTagName("div")
		}
	}
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (item *NavbarItem) RenderContent(out io.Writer) error {
	if item.ItemType != NAVBARIT_DIVIDER {
		if item.ImageSrc != nil {
			img := html.NewSnippet("img", `width="auto" height="28"`)
			img.Tag().SetURL("src", item.ImageSrc)
			item.RenderChilds(out, img)
		}
		item.RenderChilds(out, item.Content)
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

	// Styling
	IsTransparent bool
	HasShadow     bool
}

// Ensure Navbar implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*Navbar)(nil)

func (nav *Navbar) Item(index int) *NavbarItem {
	if index < 0 || index >= len(nav.items) {
		return nil
	}
	return nav.items[index]
}

// AddItems adds items to the navbar .
func (nav *Navbar) AddItems(items ...*NavbarItem) *Navbar {
	nav.items = append(nav.items, items...)
	return nav
}

// BuildTag builds the tag used to render the html element.
func (nav *Navbar) BuildTag(tag *html.Tag) {
	tag.SetTagName("nav").
		SetAttribute("role", "navigation").
		AddClasses("navbar").
		SetClassesIf(nav.IsTransparent, "is-transparent").
		SetClassesIf(nav.HasShadow, "has-shadow")
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (nav *Navbar) RenderContent(out io.Writer) error {
	// brand area
	html.WriteString(out, `<div class="navbar-brand">`)

	// brand items
	for _, item := range nav.items {
		nav.RenderChildsIf(item.ItemType == NAVBARIT_BRAND, out, item)
	}
	// burger
	html.WriteStrings(out, `<a class="navbar-burger" role="button">`, `<span></span><span></span><span></span>`, `</a>`)
	html.WriteString(out, `</div>`)

	// menu area
	menuid := nav.Id() + `menu`
	html.WriteStrings(out, `<div class="navbar-menu" id="`, menuid, `">`)

	html.WriteStrings(out, `<div class="navbar-start">`)
	for _, item := range nav.items {
		nav.RenderChildsIf(item.ItemType == NAVBARIT_START, out, item)
	}
	html.WriteString(out, `</div>`)

	html.WriteStrings(out, `<div class="navbar-end">`)
	for _, item := range nav.items {
		nav.RenderChildsIf(item.ItemType == NAVBARIT_END, out, item)
	}
	html.WriteString(out, `</div>`)

	html.WriteString(out, `</div>`) // navbar-menu

	return nil
}

// func (nav *Navbar) Template(*DataState) (_t SnippetTemplate) {
// 	_t.TagName = "nav"
// 	_t.Attributes = `class="navbar" role="navigation"`
// 	if nav.IsTransparent {
// 		nav.SetClasses("is-transparent")
// 	}
// 	if nav.HasShadow {
// 		nav.SetClasses("has-shadow")
// 	}

// 	// brand
// 	_t.Body = `<div class="navbar-brand">`

// 	for _, item := range nav.items {
// 		if item.Brand {
// 			_t.Body += nav.RenderChildSnippet(&item)
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
// 	_t.Body += `<div class="navbar-menu" id="` + String(nav.Id()) + `menu">`

// 	_t.Body += `<div class="navbar-start">`
// 	for _, item := range nav.items {
// 		if !item.Brand && !item.End {
// 			_t.Body += nav.RenderChildSnippet(&item)
// 		}
// 	}
// 	_t.Body += `</div>`

// 	_t.Body += `<div class="navbar-end">`
// 	for _, item := range nav.items {
// 		if !item.Brand && item.End {
// 			_t.Body += nav.RenderChildSnippet(&item)
// 		}
// 	}
// 	_t.Body += `</div>`

// 	_t.Body += `</div>` // navbar-menu

// 	return _t
// }
