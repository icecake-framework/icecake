package ickui

import (
	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/ick"
)

// ICKNavbar is an UISnippet registered with the ick-tag `ick-navbar`.
type ICKNavbar struct {
	ick.ICKNavbar
	dom.UI
}

func NavBar() *ICKNavbar {
	return new(ICKNavbar)
}

func (nav *ICKNavbar) AddListeners() {
	nav.DOM.AddMouseEvent(event.MOUSE_ONCLICK, func(*event.MouseEvent, *dom.Element) {
		nav.Toggle()
	})
}

func (nav *ICKNavbar) RemoveListeners() {
	nav.RemoveListeners()
}

func (nav *ICKNavbar) Toggle() {
	bs := nav.DOM.ChildrenByClassName("navbar-burger")
	if len(bs) == 0 {
		console.Warnf("unable to toggle navbar Id %q: missing navbar-burger", nav.DOM.Id())
		return
	}
	ms := nav.DOM.ChildrenByClassName("navbar-menu")
	if len(ms) == 0 {
		console.Warnf("unable to toggle navbar Id %q: missing navbar-manu", nav.DOM.Id())
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
