package ui

import (
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/html"
)

// Navbar is an UISnippet registered with the ick-tag `ick-navbar`.
type Navbar struct {
	html.Navbar
	DOM dom.Element
}

func (_nav *Navbar) AddListeners() {
	_nav.DOM.AddMouseEvent(event.MOUSE_ONCLICK, func(*event.MouseEvent, *dom.Element) {
		_nav.Toggle()
	})
}

func (_nav *Navbar) Toggle() {
	menuid := _nav.Id() + `menu`
	if _nav.DOM.HasClass("is-active") {
		_nav.DOM.RemoveClasses("is-active")
		dom.Id(menuid).RemoveClasses("is-active")
	} else {
		_nav.DOM.SetClasses("is-active")
		dom.Id(menuid).SetClasses("is-active")
	}
}
