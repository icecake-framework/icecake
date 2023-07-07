package ick

import (
	"io"
	"net/url"

	"github.com/icecake-framework/icecake/pkg/html"
)

// func init() {
// 	html.RegisterComposer("ick-menu", &Menu{})
// }

type MENUITEM_TYPE string

const (
	MENUIT_LABEL      MENUITEM_TYPE = "label"  // informative label, a <p> tag at the top level
	MENUIT_LINK       MENUITEM_TYPE = "link"   // interactive menu item, a <li><a> tag at the 1st level
	MENUIT_NESTEDLINK MENUITEM_TYPE = "nested" // interactive menu item nested in a second level, a <li><a> tag at a 2nd level
	MENUIT_FOOTER     MENUITEM_TYPE = "footer" // informative text, a <p> tag inserted in the foot of the menu
)

// IckMenuItem is an icecake snippet providing the HTML rendering for a [bulma navbar item].
//
// Can't be used for inline rendering.
//
// [bulma navbar item]: https://bulma.io/documentation/components/navbar/#navbar-item
type IckMenuItem struct {
	html.BareSnippet

	// Optional Key allows to access a specific navbaritem, whatever it's level in the hierarchy, directly from the navbar.
	Key string

	// The Item Type defines the location of the item in the navbar or if it's a simple divider.
	// If Type is empty, NAVBARIT_START is used for rendering.
	Type MENUITEM_TYPE

	// Item Text
	Text string

	// HRef defines the optional associated url link.
	// If HRef is defined the item become an anchor link <a>, otherwise it's a <div>
	// HRef can be nil. Usually it's created calling MenuItem.ParseHRef
	HRef *url.URL

	// Highlight this item
	IsActive bool
}

// Ensuring ICKMenuItem implements the right interface
var _ html.ElementComposer = (*IckMenuItem)(nil)

// ParseHRef tries to parse rawUrl to HRef ignoring error.
func (mnui *IckMenuItem) ParseHRef(rawUrl string) *IckMenuItem {
	mnui.HRef, _ = url.Parse(rawUrl)
	return mnui
}

// Clone clones this navbar and all its items and subitem, keeping their attributes their item index and their key.
func (mnui IckMenuItem) Clone() *IckMenuItem {
	m := new(IckMenuItem)
	m.Key = mnui.Key
	m.Type = mnui.Type
	m.Text = mnui.Text
	if mnui.HRef != nil {
		m.HRef = new(url.URL)
		*m.HRef = *mnui.HRef
	}
	m.IsActive = mnui.IsActive
	return m
}

// BuildTag builds the tag used to render the html element.
func (mnui *IckMenuItem) BuildTag() html.Tag {
	if mnui.Type == MENUIT_LABEL {
		mnui.Tag().SetTagName("p").AddClass("menu-label")
	} else {
		mnui.Tag().SetTagName("li")

	}
	mnui.Tag().SetAttributeIf(mnui.Key != "", "data-key", mnui.Key)
	return *mnui.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (mnui *IckMenuItem) RenderContent(out io.Writer) error {
	switch mnui.Type {
	case MENUIT_LABEL:
		html.WriteString(out, mnui.Text)
	case MENUIT_FOOTER:
		mnui.RenderChild(out, html.ToHTML(mnui.Text))
	default:
		item := A().SetHRef(mnui.HRef)
		item.Push(html.ToHTML(mnui.Text))
		item.Tag().SetClassIf(mnui.IsActive, "is-active")
		mnui.RenderChild(out, item)
	}
	return nil
}

// IckMenu is an icecake snippet providing the HTML rendering for a [bulma menu].
//
// [bulma menu]: https://bulma.io/documentation/components/menu
type IckMenu struct {
	html.BareSnippet

	menuTag html.Tag // menu tag: nav, aside, menu. <menu> is used if nothing is specified. Cna be used to setup some classes like "is-small"

	items []*IckMenuItem // list of Menu items
}

// Ensuring ICKMenu implements the right interface
var _ html.ElementComposer = (*IckMenu)(nil)

// Clone clones this Menu and all its items and subitem, keeping their attributes their item index and their key.
func (src IckMenu) Clone() *IckMenu {
	clone := new(IckMenu)
	clone.BareSnippet = *src.BareSnippet.Clone()
	clone.menuTag = *src.menuTag.Clone()
	clone.items = make([]*IckMenuItem, len(src.items))
	for i, itm := range src.items {
		clone.items[i] = itm.Clone()
	}
	return clone
}

func (mnu *IckMenu) MenuTag() *html.Tag {
	if mnu.menuTag.AttributeMap == nil {
		mnu.menuTag.AttributeMap = make(html.AttributeMap)
	}
	return &mnu.menuTag
}

// SetActiveItem look for the key item (or subitem) and sets its IsActive flag.
// warning: does not unset other actve items if any.
func (mnu *IckMenu) SetActiveItem(key string) *IckMenu {
	itm := mnu.Item(key)
	if itm != nil {
		itm.IsActive = true
	}
	return mnu
}

// AddItem adds the item to the Menu
func (mnu *IckMenu) AddItem(key string, itmtyp MENUITEM_TYPE, txt string) *IckMenuItem {
	itm := new(IckMenuItem)
	itm.Key = key
	itm.Type = itmtyp
	itm.Text = txt
	mnu.items = append(mnu.items, itm)
	return itm
}

// At returns the item at a given index.
// returns nil if index is out of range.
func (mnu *IckMenu) At(index int) *IckMenuItem {
	if index < 0 || index >= len(mnu.items) {
		return nil
	}
	return mnu.items[index]
}

// Item returns the first item found with the given key, walking through all levels.
// returns nil if key is not found
func (mnu *IckMenu) Item(key string) *IckMenuItem {
	for _, itm := range mnu.items {
		if itm.Key == key {
			return itm
		}
	}
	return nil
}

// BuildTag builds the tag used to render the html element.
func (mnu *IckMenu) BuildTag() html.Tag {
	mnu.Tag().SetTagName("div")

	// set style height if there's a footer
	for _, item := range mnu.items {
		if item.Type == MENUIT_FOOTER {
			mnu.Tag().AddClass("is-flex is-flex-direction-column is-justify-content-space-between")
			mnu.Tag().AddStyle("height:100%;")
			break
		}
	}
	return *mnu.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (mnu *IckMenu) RenderContent(out io.Writer) error {
	mnutag := mnu.menuTag.Clone()
	if tagname, _ := mnutag.TagName(); tagname == "" {
		mnutag.SetTagName("menu")
	}
	mnutag.AddClass("menu")
	mnutag.SetAttribute("role", "navigation")
	mnutag.RenderOpening(out)

	lastlevel := 0
	for _, item := range mnu.items {
		switch item.Type {
		case MENUIT_LABEL:
			switch lastlevel {
			case 0:
			case 1: // close upper list
				html.WriteString(out, "</ul>")
			case 2: // close upper lists
				html.WriteString(out, "</li></ul></ul>")
			}
			lastlevel = 0
		case MENUIT_LINK:
			switch lastlevel {
			case 0: // open 1st list
				html.WriteString(out, `<ul class="menu-list">`)
			case 1:
			case 2: // close upper list, back to 1st list
				html.WriteString(out, "</li></ul>")
			}
			lastlevel = 1
		case MENUIT_NESTEDLINK:
			switch lastlevel {
			case 0: // open 1st list and 2nd one
				html.WriteString(out, "<ul><li><ul>")
			case 1: // open 2nd list
				html.WriteString(out, "<li><ul>")
			case 2:
			}
			lastlevel = 2
		default:
			continue
		}

		mnu.RenderChild(out, item)
	}

	// close the menu
	switch lastlevel {
	case 1:
		html.WriteString(out, "</ul>")
	case 2:
		html.WriteString(out, "</ul></ul>")
	}

	mnutag.RenderClosing(out)

	// add footer
	hasfooter := false
	foottag := html.NewTag("div", html.ParseAttributes(`class="`+mnutag.Classes()+`"`))
	for _, item := range mnu.items {
		if item.Type == MENUIT_FOOTER {
			if !hasfooter {
				foottag.RenderOpening(out)
				html.WriteString(out, `<ul class="menu-list">`)
				hasfooter = true
			}
			mnu.RenderChild(out, item)
		}
	}
	if hasfooter {
		html.WriteString(out, `</ul>`)
		foottag.RenderClosing(out)
	}
	return nil
}
