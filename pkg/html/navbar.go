package html

import (
	"io"
	"net/url"
)

func init() {
	RegisterComposer("ick-navbar", &Navbar{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type NAVBARITEM_TYPE string

const (
	NAVBARIT_DIVIDER = "divider" // creates a divider between two navbar items. other navbar properties are ignored.
	NAVBARIT_BRAND   = "brand"   // item stacked in the brand area of the navbar.
	NAVBARIT_START   = "start"   // item stacked at the start (left) of the navbar, this is the default behaviour
	NAVBARIT_END     = "end"     // item stacked at the end (right) of the navbar,
)

type NavbarItem struct {
	HTMLSnippet

	// The Item Type defines the location of the item in the navbar or if it's a simple divider.
	// If ItemType is empty, NAVBARIT_START is used for rendering.
	ItemType NAVBARITEM_TYPE

	// HRef defines the optional associated url link.
	// If HRef is defined the item become an anchor link <a>, otherwise it's a <div>
	// HRef can be nil. Usually it's created calling NavbarItem.TryParseHRef
	HRef *url.URL

	// ImageSrc defines an optional image to display at the begining of the Item
	ImageSrc *url.URL // the url for the source of the image

	// Body
	Body HTMLString
}

// Ensure NavbarItem implements HTMLComposer interface
var _ HTMLComposer = (*NavbarItem)(nil)

func (item *NavbarItem) Tag() *Tag {
	if item.ItemType == NAVBARIT_DIVIDER {
		item.tag.SetName("hr")
	} else {
		if item.HRef != nil && item.HRef.String() != "" {
			item.tag.SetName("a")
		} else {
			item.tag.SetName("div")
		}
	}

	amap := item.tag.Attributes()
	if item.ItemType == NAVBARIT_DIVIDER {
		amap.SetClasses("navbar-divider")
	} else {
		amap.SetClasses("navbar-item")
		if item.HRef != nil {
			if href := item.HRef.String(); href != "" {
				amap.SetAttribute("href", href, true)
			}
		}
	}

	return &item.tag
}

func (item *NavbarItem) WriteBody(out io.Writer) error {
	if item.ItemType != NAVBARIT_DIVIDER {
		if item.ImageSrc != nil {
			imgsrc := item.ImageSrc.String()
			WriteStringsIf(imgsrc != "", out, `<img src="`, imgsrc, `" width="auto" height="28">`)
		}
		item.UnfoldHTML(out, item.Body, nil)
	}
	return nil
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
	HTMLSnippet

	items []*NavbarItem

	// Styling
	IsTransparent bool
	HasShadow     bool
}

// Ensure Navbar implements HTMLComposer interface
var _ HTMLComposer = (*Navbar)(nil)

func (_nav *Navbar) AddItems(_items ...*NavbarItem) *Navbar {
	for _, item := range _items {
		_nav.items = append(_nav.items, item)
	}
	return _nav
}

func (nav *Navbar) Tag() *Tag {
	nav.tag.SetName("nav")
	amap := nav.tag.Attributes()
	amap.SetClasses("navbar").setAttribute("role", "navigation", true)
	amap.SetClassesIf(nav.IsTransparent, "is-transparent")
	amap.SetClassesIf(nav.HasShadow, "has-shadow")
	return &nav.tag
}

func (nav *Navbar) WriteBody(out io.Writer) error {
	// brand area
	WriteString(out, `<div class="navbar-brand">`)

	// brand items
	for _, item := range nav.items {
		nav.WriteChildSnippetIf(item.ItemType == NAVBARIT_BRAND, out, item)
	}
	// burger
	WriteStrings(out, `<a class="navbar-burger" role="button">`, `<span></span><span></span><span></span>`, `</a>`)
	WriteString(out, `</div>`)

	// menu area
	menuid := nav.Id() + `menu`
	WriteStrings(out, `<div class="navbar-menu" id="`, menuid, `">`)

	WriteStrings(out, `<div class="navbar-start">`)
	for _, item := range nav.items {
		nav.WriteChildSnippetIf(item.ItemType == NAVBARIT_START, out, item)
	}
	WriteString(out, `</div>`)

	WriteStrings(out, `<div class="navbar-end">`)
	for _, item := range nav.items {
		nav.WriteChildSnippetIf(item.ItemType == NAVBARIT_END, out, item)
	}
	WriteString(out, `</div>`)

	WriteString(out, `</div>`) // navbar-menu

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
