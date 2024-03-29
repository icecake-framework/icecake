package ick

import (
	"io"
	"net/url"

	"github.com/huandu/go-clone"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

type NAVBARITEM_TYPE string

const (
	NAVBARIT_DIVIDER NAVBARITEM_TYPE = "divider" // creates a divider between two navbar items. other navbar properties are ignored.
	NAVBARIT_BRAND   NAVBARITEM_TYPE = "brand"   // item stacked in the brand area of the navbar.
	NAVBARIT_START   NAVBARITEM_TYPE = "start"   // item stacked at the start (left) of the navbar, this is the default behaviour
	NAVBARIT_END     NAVBARITEM_TYPE = "end"     // item stacked at the end (right) of the navbar,
)

// ICKNavbarItem is an icecake snippet providing the HTML rendering for a [bulma navbar item].
//
// [bulma navbar item]: https://bulma.io/documentation/components/navbar/#navbar-item
type ICKNavbarItem struct {
	ickcore.BareSnippet

	// Optional Key allows to access a specific navbaritem, whatever it's level in the hierarchy, directly from the navbar.
	Key string

	// The Item Type defines the location of the item in the navbar or if it's a simple divider.
	// If Type is empty, NAVBARIT_START is used for rendering.
	Type NAVBARITEM_TYPE

	// Item Content
	Content ickcore.ContentComposer

	// HRef defines the optional associated url link.
	// If HRef is defined the item become an anchor link <a>, otherwise it's a <div>
	// HRef can be nil. Usually it's created calling NavbarItem.TryParseHRef
	HRef *url.URL

	// ImageSrc defines an optional image to display at the begining of the Item
	ImageSrc *url.URL // the url for the source of the image

	// Highlight this item
	IsActive bool

	items []*ICKNavbarItem // list of navbar items
}

// Ensuring ICKNavbarItem implements the right interface
var _ ickcore.ContentComposer = (*ICKNavbarItem)(nil)
var _ ickcore.TagBuilder = (*ICKNavbarItem)(nil)

// Clone clones this navbar and all its items and subitem, keeping their attributes their item index and their key.
func (navi ICKNavbarItem) Clone() *ICKNavbarItem {
	c := new(ICKNavbarItem)
	c.BareSnippet = *navi.BareSnippet.Clone()
	c.Key = navi.Key
	c.Type = navi.Type

	if navi.Content != nil {
		copy := clone.Clone(navi.Content)
		c.Content = copy.(ickcore.ContentComposer)
	}

	if navi.HRef != nil {
		c.HRef = new(url.URL)
		*c.HRef = *navi.HRef
	}
	if navi.ImageSrc != nil {
		c.ImageSrc = new(url.URL)
		*c.ImageSrc = *navi.ImageSrc
	}
	c.IsActive = navi.IsActive

	c.items = make([]*ICKNavbarItem, len(navi.items))
	for i, itm := range navi.items {
		c.items[i] = itm.Clone()
	}
	return c
}

// BuildTag builds the tag used to render the html element.
// The Navbar Item tag depends on the item properties:
//   - it's <hr> for a NAVBARIT_DIVIDER item type, otherwise
//   - it's <a> when an HRef is provided,
//   - it's <div> in other cases
func (navi *ICKNavbarItem) BuildTag() ickcore.Tag {
	if navi.Type == NAVBARIT_DIVIDER {
		navi.Tag().
			SetTagName("hr").
			PickClass("navbar-divider navbar-item", "navbar-divider")
	} else {
		navi.Tag().PickClass("navbar-divider navbar-item", "navbar-item").SetClassIf(navi.IsActive, "is-active")
		if navi.HRef != nil {
			navi.Tag().SetTagName("a").SetURL("href", navi.HRef)
		} else {
			navi.Tag().SetTagName("div").RemoveAttribute("href")
		}
	}
	return *navi.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (navi *ICKNavbarItem) RenderContent(out io.Writer) error {
	if navi.Type != NAVBARIT_DIVIDER {
		if navi.ImageSrc != nil {
			img := Elem("img", `width="auto" height="28"`)
			img.Tag().SetURL("src", navi.ImageSrc)
			ickcore.RenderChild(out, navi, img)
		}
		ickcore.RenderChild(out, navi, navi.Content)
	}
	return nil
}

// AddItem adds the item as a subitem within the navbar item
func (navi *ICKNavbarItem) AddItem(key string, itmtyp NAVBARITEM_TYPE, content ickcore.ContentComposer) *ICKNavbarItem {
	itm := new(ICKNavbarItem)
	itm.Key = key
	itm.Type = itmtyp
	itm.Content = content
	navi.items = append(navi.items, itm)
	return itm
}

// At returns the item at a given index.
// returns nil if index is out of range.
func (navi *ICKNavbarItem) At(index int) *ICKNavbarItem {
	if index < 0 || index >= len(navi.items) {
		return nil
	}
	return navi.items[index]
}

// Item returns the first item found with the given key, walking through all levels.
// returns nil if key is not found
func (navi *ICKNavbarItem) Item(key string) *ICKNavbarItem {
	for _, itm := range navi.items {
		if itm.Key == key {
			return itm
		}
		if found := itm.Item(key); found != nil {
			return found
		}
	}
	return nil
}

// ParseHRef tries to parse rawUrl to HRef ignoring error.
func (navi *ICKNavbarItem) ParseHRef(rawUrl string) *ICKNavbarItem {
	navi.HRef, _ = url.Parse(rawUrl)
	return navi
}

func (navi *ICKNavbarItem) SetHRef(u url.URL) *ICKNavbarItem {
	navi.HRef = &u
	return navi
}

// ParseImageSrc tries to parse rawUrl to image src ignoring error.
func (navi *ICKNavbarItem) ParseImageSrc(rawUrl string) *ICKNavbarItem {
	navi.ImageSrc, _ = url.Parse(rawUrl)
	return navi
}

func (navi *ICKNavbarItem) SetImageSrc(u url.URL) *ICKNavbarItem {
	navi.ImageSrc = &u
	return navi
}

// ICKNavbar is an icecake snippet providing the HTML rendering for a [bulma navbar].
//
// [bulma navbar]: https://bulma.io/documentation/components/navbar
type ICKNavbar struct {
	ickcore.BareSnippet

	items []*ICKNavbarItem // list of navbar items

	// Styling properties
	IsTransparent bool // renders a transparent navbar
	HasShadow     bool // renders a shadow below the navbar
}

// Ensuring ICKNavbar implements the right interface
var _ ickcore.ContentComposer = (*ICKNavbar)(nil)
var _ ickcore.TagBuilder = (*ICKNavbar)(nil)

func NavBar(attrs ...string) *ICKNavbar {
	n := new(ICKNavbar)
	n.Tag().ParseAttributes(attrs...)
	return n
}

// Clone clones this navbar and all its items and subitem, keeping their attributes their item index and their key.
func (src ICKNavbar) Clone() *ICKNavbar {
	c := new(ICKNavbar)
	c.BareSnippet = *src.BareSnippet.Clone()
	c.IsTransparent = src.IsTransparent
	c.HasShadow = src.HasShadow
	c.items = make([]*ICKNavbarItem, len(src.items))
	for i, itm := range src.items {
		c.items[i] = itm.Clone()
	}
	return c
}

// AddItem adds the item to the navbar
func (nav *ICKNavbar) AddItem(key string, itmtyp NAVBARITEM_TYPE, content ickcore.ContentComposer) *ICKNavbarItem {
	itm := new(ICKNavbarItem)
	itm.Key = key
	itm.Type = itmtyp
	itm.Content = content
	nav.items = append(nav.items, itm)
	return itm
}

// At returns the item at a given index.
// returns nil if index is out of range.
func (nav *ICKNavbar) At(index int) *ICKNavbarItem {
	if index < 0 || index >= len(nav.items) {
		return nil
	}
	return nav.items[index]
}

// SetActiveItem look for the key item (or subitem) and sets its IsActive flag.
// warning: does not unset other actve items if any.
func (nav *ICKNavbar) SetActiveItem(key string) *ICKNavbar {
	itm := nav.Item(key)
	if itm != nil {
		itm.IsActive = true
	}
	return nav
}

// Item returns the first item found with the given key, walking through all levels.
// returns nil if key is not found
func (nav *ICKNavbar) Item(key string) *ICKNavbarItem {
	for _, itm := range nav.items {
		if itm.Key == key {
			return itm
		}
		if found := itm.Item(key); found != nil {
			return found
		}
	}
	return nil
}

func (nav *ICKNavbar) NeedRendering() bool {
	return len(nav.items) > 0
}

// BuildTag builds the tag used to render the html element.
func (nav *ICKNavbar) BuildTag() ickcore.Tag {
	nav.Tag().
		SetTagName("nav").
		SetAttribute("role", "navigation").
		AddClass("navbar").
		SetClassIf(nav.IsTransparent, "is-transparent").
		SetClassIf(nav.HasShadow, "has-shadow")
	return *nav.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (nav *ICKNavbar) RenderContent(out io.Writer) error {
	// brand area
	ickcore.RenderString(out, `<div class="navbar-brand">`)

	// brand items
	for _, item := range nav.items {
		if item.Type == NAVBARIT_BRAND {
			ickcore.RenderChild(out, nav, item)
		}
	}
	// burger
	ickcore.RenderString(out, `<a class="navbar-burger" role="button">`, `<span></span><span></span><span></span>`, `</a>`)
	ickcore.RenderString(out, `</div>`)

	// menu area
	// the burger id is required for flipping it
	ickcore.RenderString(out, `<div class="navbar-menu">`)

	ickcore.RenderString(out, `<div class="navbar-start">`)
	for _, item := range nav.items {
		if item.Type == NAVBARIT_START {
			ickcore.RenderChild(out, nav, item)
		}
	}
	ickcore.RenderString(out, `</div>`)

	ickcore.RenderString(out, `<div class="navbar-end">`)
	for _, item := range nav.items {
		if item.Type == NAVBARIT_END {
			ickcore.RenderChild(out, nav, item)
		}
	}
	ickcore.RenderString(out, `</div>`)

	ickcore.RenderString(out, `</div>`) // navbar-menu

	return nil
}
