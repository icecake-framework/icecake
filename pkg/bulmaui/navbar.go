package bulmaui

import (
	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
)

// Navbar is an UISnippet registered with the ick-tag `ick-navbar`.
type Navbar struct {
	bulma.Navbar
	DOM dom.Element
}

// returns nil if the id does not exists or if it's not a Navbar
func WrapNavbar(id string) *Navbar {
	e := dom.RenderedId(id, "ick-navbar")
	if e == nil {
		return nil
	}
	n := new(Navbar)
	n.DOM.Wrap(e)
	n.AddListeners()
	return n
}

func (nav *Navbar) AddListeners() {
	nav.DOM.AddMouseEvent(event.MOUSE_ONCLICK, func(*event.MouseEvent, *dom.Element) {
		nav.Toggle()
	})
}

func (nav *Navbar) Toggle() {
	bs := nav.DOM.ChildrenByClassName("navbar-burger")
	if len(bs) == 0 {
		console.Warnf("unable to toggle navbar Id %q: missing navbar-burger", nav.Id())
		return
	}
	ms := nav.DOM.ChildrenByClassName("navbar-menu")
	if len(ms) == 0 {
		console.Warnf("unable to toggle navbar Id %q: missing navbar-manu", nav.Id())
		return
	}

	if bs[0].HasClass("is-active") {
		bs[0].RemoveClass("is-active")
		ms[0].RemoveClass("is-active")
	} else {
		bs[0].AddClass("is-active")
		ms[0].AddClass("is-active")
	}
}
