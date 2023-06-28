package bulma

import (
	"io"
	"net/url"

	"github.com/huandu/go-clone"
	"github.com/icecake-framework/icecake/pkg/html"
)

// func init() {
// 	html.RegisterComposer("ick-navbar", &Navbar{})
// }

type NAVBARITEM_TYPE string

const (
	NAVBARIT_DIVIDER NAVBARITEM_TYPE = "divider" // creates a divider between two navbar items. other navbar properties are ignored.
	NAVBARIT_BRAND   NAVBARITEM_TYPE = "brand"   // item stacked in the brand area of the navbar.
	NAVBARIT_START   NAVBARITEM_TYPE = "start"   // item stacked at the start (left) of the navbar, this is the default behaviour
	NAVBARIT_END     NAVBARITEM_TYPE = "end"     // item stacked at the end (right) of the navbar,
)

// bulma.NavbarItem is an icecake snippet providing the HTML rendering for a [bulma navbar item].
//
// Can't be used for inline rendering.
//
// [bulma navbar item]: https://bulma.io/documentation/components/navbar/#navbar-item
type NavbarItem struct {
	html.HTMLSnippet

	// Optional Key allows to access a specific navbaritem, whatever it's level in the hierarchy, directly from the navbar.
	Key string

	// The Item Type defines the location of the item in the navbar or if it's a simple divider.
	// If Type is empty, NAVBARIT_START is used for rendering.
	Type NAVBARITEM_TYPE

	// Item Content
	Content html.HTMLComposer

	// HRef defines the optional associated url link.
	// If HRef is defined the item become an anchor link <a>, otherwise it's a <div>
	// HRef can be nil. Usually it's created calling NavbarItem.TryParseHRef
	HRef *url.URL

	// ImageSrc defines an optional image to display at the begining of the Item
	ImageSrc *url.URL // the url for the source of the image

	// Highlight this item
	IsActive bool

	items []*NavbarItem // list of navbar items
}

// Ensure NavbarItem implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*NavbarItem)(nil)

// Clone clones this navbar and all its items and subitem, keeping their attributes their item index and their key.
func (navi NavbarItem) Clone() *NavbarItem {
	c := new(NavbarItem)
	c.Key = navi.Key
	c.Type = navi.Type

	if navi.Content != nil {
		copy := clone.Clone(navi.Content)
		c.Content = copy.(html.HTMLComposer)
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

	c.items = make([]*NavbarItem, len(navi.items))
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
func (navi *NavbarItem) BuildTag(tag *html.Tag) {
	if navi.Type == NAVBARIT_DIVIDER {
		tag.SetTagName("hr")
		tag.PickClass("navbar-divider navbar-item", "navbar-divider")
	} else {
		tag.PickClass("navbar-divider navbar-item", "navbar-item").SetClassesIf(navi.IsActive, "is-active")
		if navi.HRef != nil {
			tag.SetTagName("a").SetURL("href", navi.HRef)
		} else {
			tag.SetTagName("div").RemoveAttribute("href")
		}
	}
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (navi *NavbarItem) RenderContent(out io.Writer) error {
	if navi.Type != NAVBARIT_DIVIDER {
		if navi.ImageSrc != nil {
			img := html.NewSnippet("img", `width="auto" height="28"`)
			img.Tag().SetURL("src", navi.ImageSrc)
			navi.RenderChilds(out, img)
		}
		navi.RenderChilds(out, navi.Content)
	}
	return nil
}

// AddItem adds the item as a subitem within the navbar item
func (navi *NavbarItem) AddItem(key string, itmtyp NAVBARITEM_TYPE, content html.HTMLComposer) *NavbarItem {
	itm := new(NavbarItem)
	itm.Key = key
	itm.Type = itmtyp
	itm.Content = content
	itm.Meta().LinkParent(navi)
	navi.items = append(navi.items, itm)
	return itm
}

// At returns the item at a given index.
// returns nil if index is out of range.
func (navi *NavbarItem) At(index int) *NavbarItem {
	if index < 0 || index >= len(navi.items) {
		return nil
	}
	return navi.items[index]
}

// Item returns the first item found with the given key, walking through all levels.
// returns nil if key is not found
func (navi *NavbarItem) Item(key string) *NavbarItem {
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
func (navi *NavbarItem) ParseHRef(rawUrl string) *NavbarItem {
	navi.HRef, _ = url.Parse(rawUrl)
	return navi
}

// ParseImageSrc tries to parse rawUrl to image src ignoring error.
func (navi *NavbarItem) ParseImageSrc(rawUrl string) *NavbarItem {
	navi.ImageSrc, _ = url.Parse(rawUrl)
	return navi
}

// bulma.Navbar is an icecake snippet providing the HTML rendering for a [bulma navbar].
//
// Can't be used for inline rendering.
//
// [bulma navbar]: https://bulma.io/documentation/components/navbar
type Navbar struct {
	html.HTMLSnippet

	items []*NavbarItem // list of navbar items

	// Styling
	IsTransparent bool
	HasShadow     bool
}

// Ensure Navbar implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*Navbar)(nil)

// Clone clones this navbar and all its items and subitem, keeping their attributes their item index and their key.
func (src Navbar) Clone() *Navbar {
	to := new(Navbar)
	to.HTMLSnippet = *src.HTMLSnippet.Clone()
	to.IsTransparent = src.IsTransparent
	to.HasShadow = src.HasShadow
	to.items = make([]*NavbarItem, len(src.items))
	for i, itm := range src.items {
		to.items[i] = itm.Clone()
	}
	return to
}

// AddItem adds the item to the navbar
func (nav *Navbar) AddItem(key string, itmtyp NAVBARITEM_TYPE, content html.HTMLComposer) *NavbarItem {
	itm := new(NavbarItem)
	itm.Key = key
	itm.Type = itmtyp
	itm.Content = content
	itm.Meta().LinkParent(nav)
	nav.items = append(nav.items, itm)
	return itm
}

// At returns the item at a given index.
// returns nil if index is out of range.
func (nav *Navbar) At(index int) *NavbarItem {
	if index < 0 || index >= len(nav.items) {
		return nil
	}
	return nav.items[index]
}

// SetActiveItem look for the key item (or subitem) and sets its IsActive flag.
// warning: does not unset other actve items if any.
func (nav *Navbar) SetActiveItem(key string) *Navbar {
	itm := nav.Item(key)
	if itm != nil {
		itm.IsActive = true
	}
	return nav
}

// Item returns the first item found with the given key, walking through all levels.
// returns nil if key is not found
func (nav *Navbar) Item(key string) *NavbarItem {
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
		nav.RenderChildsIf(item.Type == NAVBARIT_BRAND, out, item)
	}
	// burger
	html.WriteStrings(out, `<a class="navbar-burger" role="button">`, `<span></span><span></span><span></span>`, `</a>`)
	html.WriteString(out, `</div>`)

	// menu area
	// the burger id is required for flipping it
	html.WriteStrings(out, `<div class="navbar-menu">`)

	html.WriteStrings(out, `<div class="navbar-start">`)
	for _, item := range nav.items {
		nav.RenderChildsIf(item.Type == NAVBARIT_START, out, item)
	}
	html.WriteString(out, `</div>`)

	html.WriteStrings(out, `<div class="navbar-end">`)
	for _, item := range nav.items {
		nav.RenderChildsIf(item.Type == NAVBARIT_END, out, item)
	}
	html.WriteString(out, `</div>`)

	html.WriteString(out, `</div>`) // navbar-menu

	return nil
}
