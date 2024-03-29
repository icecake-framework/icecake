package ick

import (
	"io"
	"net/url"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

// func init() {
// 	ickcore.RegisterComposer("ick-menu", &Menu{})
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
	ickcore.BareSnippet

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
var _ ickcore.ContentComposer = (*IckMenuItem)(nil)
var _ ickcore.TagBuilder = (*IckMenuItem)(nil)

// ParseHRef tries to parse rawUrl to HRef ignoring error.
func (mnui *IckMenuItem) ParseHRef(rawUrl string) *IckMenuItem {
	mnui.HRef, _ = url.Parse(rawUrl)
	return mnui
}

// Clone clones this navbar and all its items and subitem, keeping their attributes their item index and their key.
func (mnui IckMenuItem) Clone() *IckMenuItem {
	c := new(IckMenuItem)
	c.BareSnippet = *mnui.BareSnippet.Clone()
	c.Key = mnui.Key
	c.Type = mnui.Type
	c.Text = mnui.Text
	if mnui.HRef != nil {
		c.HRef = new(url.URL)
		*c.HRef = *mnui.HRef
	}
	c.IsActive = mnui.IsActive
	return c
}

// BuildTag builds the tag used to render the html element.
func (mnui *IckMenuItem) BuildTag() ickcore.Tag {
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
		ickcore.RenderString(out, mnui.Text)
	case MENUIT_FOOTER:
		ickcore.RenderChild(out, mnui, ickcore.ToHTML(mnui.Text))
	default:
		item := Link(ickcore.ToHTML(mnui.Text)).SetHRef(mnui.HRef)
		item.Tag().SetClassIf(mnui.IsActive, "is-active")
		ickcore.RenderChild(out, mnui, item)
	}
	return nil
}

type MENU_TYPE string

const (
	MENUTYP_MENU  MENU_TYPE = "menu"
	MENUTYP_NAV   MENU_TYPE = "nav"
	MENUTYP_ASIDE MENU_TYPE = "aside"
)

// ICKMenu is an icecake snippet providing the HTML rendering for a [bulma menu].
//
// [bulma menu]: https://bulma.io/documentation/components/menu
type ICKMenu struct {
	ickcore.BareSnippet

	MENU_TYPE
	SIZE

	items []*IckMenuItem // list of Menu items
}

// Ensuring ICKMenu implements the right interface
var _ ickcore.ContentComposer = (*ICKMenu)(nil)
var _ ickcore.TagBuilder = (*ICKMenu)(nil)

func Menu(id string, attrs ...string) *ICKMenu {
	n := new(ICKMenu)
	n.Tag().ParseAttributes(attrs...)
	n.Tag().SetId(id)
	return n
}

// Clone clones this Menu and all its items and subitem, keeping their attributes their item index and their key.
func (src ICKMenu) Clone() *ICKMenu {
	c := new(ICKMenu)
	c.BareSnippet = *src.BareSnippet.Clone()
	c.MENU_TYPE = src.MENU_TYPE
	c.SIZE = src.SIZE
	c.items = make([]*IckMenuItem, len(src.items))
	for i, itm := range src.items {
		c.items[i] = itm.Clone()
	}
	return c
}

func (mnu *ICKMenu) SetType(t MENU_TYPE) *ICKMenu {
	mnu.MENU_TYPE = t
	return mnu
}

func (mnu *ICKMenu) SetSize(sz SIZE) *ICKMenu {
	mnu.SIZE = sz
	return mnu
}

// SetActiveItem look for the key item (or subitem) and sets its IsActive flag.
// warning: does not unset other actve items if any.
func (mnu *ICKMenu) SetActiveItem(key string) *ICKMenu {
	itm := mnu.Item(key)
	if itm != nil {
		itm.IsActive = true
	}
	return mnu
}

// AddItem adds the item to the Menu
func (mnu *ICKMenu) AddItem(key string, itmtyp MENUITEM_TYPE, txt string) *IckMenuItem {
	itm := new(IckMenuItem)
	itm.Key = key
	itm.Type = itmtyp
	itm.Text = txt
	mnu.items = append(mnu.items, itm)
	return itm
}

// At returns the item at a given index.
// returns nil if index is out of range.
func (mnu *ICKMenu) At(index int) *IckMenuItem {
	if index < 0 || index >= len(mnu.items) {
		return nil
	}
	return mnu.items[index]
}

// Item returns the first item found with the given key, walking through all levels.
// returns nil if key is not found
func (mnu *ICKMenu) Item(key string) *IckMenuItem {
	for _, itm := range mnu.items {
		if itm.Key == key {
			return itm
		}
	}
	return nil
}

func (mnu *ICKMenu) NeedRendering() bool {
	return len(mnu.items) > 0
}

// BuildTag builds the tag used to render the html element.
func (mnu *ICKMenu) BuildTag() ickcore.Tag {
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
func (mnu *ICKMenu) RenderContent(out io.Writer) error {
	typ := mnu.MENU_TYPE
	if typ == "" {
		typ = MENUTYP_MENU
	}
	mnutag := ickcore.NewTag(string(typ), `class="menu" role="navigation"`)
	mnutag.AddClass(string(mnu.SIZE))
	mnutag.RenderOpening(out)

	lastlevel := 0
	for _, item := range mnu.items {
		switch item.Type {
		case MENUIT_LABEL:
			switch lastlevel {
			case 0:
			case 1: // close upper list
				ickcore.RenderString(out, "</ul>")
			case 2: // close upper lists
				ickcore.RenderString(out, "</li></ul></ul>")
			}
			lastlevel = 0
		case MENUIT_LINK:
			switch lastlevel {
			case 0: // open 1st list
				ickcore.RenderString(out, `<ul class="menu-list">`)
			case 1:
			case 2: // close upper list, back to 1st list
				ickcore.RenderString(out, "</li></ul>")
			}
			lastlevel = 1
		case MENUIT_NESTEDLINK:
			switch lastlevel {
			case 0: // open 1st list and 2nd one
				ickcore.RenderString(out, "<ul><li><ul>")
			case 1: // open 2nd list
				ickcore.RenderString(out, "<li><ul>")
			case 2:
			}
			lastlevel = 2
		default:
			continue
		}

		ickcore.RenderChild(out, mnu, item)
	}

	// close the menu
	switch lastlevel {
	case 1:
		ickcore.RenderString(out, "</ul>")
	case 2:
		ickcore.RenderString(out, "</ul></ul>")
	}

	mnutag.RenderClosing(out)

	// add footer
	hasfooter := false
	foottag := ickcore.NewTag("div")
	foottag.AddClass(string(mnu.SIZE))
	for _, item := range mnu.items {
		if item.Type == MENUIT_FOOTER {
			if !hasfooter {
				foottag.RenderOpening(out)
				ickcore.RenderString(out, `<ul class="menu-list">`)
				hasfooter = true
			}
			ickcore.RenderChild(out, mnu, item)
		}
	}
	if hasfooter {
		ickcore.RenderString(out, `</ul>`)
		foottag.RenderClosing(out)
	}
	return nil
}
