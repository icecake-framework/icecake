package ui

import (
	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-navbar", &Navbar{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type NavbarBurger struct {
	dom.UISnippet

	// the element id corresponding to the navbar menu associated with this burger button
	NavbarMenuID string
}

func (_burger *NavbarBurger) Template(*html.DataState) (_t html.SnippetTemplate) {
	_t.TagName = "a"
	_t.Attributes = `class="navbar-burger" role="button"`
	_t.Body = `<span></span><span></span><span></span>`
	return _t
}

func (_burger *NavbarBurger) AddListeners() {
	if _burger.NavbarMenuID != "" {
		_burger.DOM.AddMouseEvent(event.MOUSE_ONCLICK, func(*event.MouseEvent, *dom.Element) {
			_burger.Toggle()
		})

	} else {
		console.Warnf("missing NavbarMenuID for the NavbarBurger snippet")
	}
}

func (_burger *NavbarBurger) Toggle() {
	if _burger.DOM.HasClass("is-active") {
		_burger.DOM.RemoveClasses("is-active")
		dom.Id(_burger.NavbarMenuID).RemoveClasses("is-active")
	} else {
		_burger.DOM.SetClasses("is-active")
		dom.Id(_burger.NavbarMenuID).SetClasses("is-active")
	}
}

type NavbarMenu struct {
	dom.UISnippet
}

func (_burger *NavbarMenu) Template(*html.DataState) (_t html.SnippetTemplate) {
	_t.TagName = "div"
	_t.Attributes = `class="navbar-menu"`
	_t.Body = `<div class="navbar-start"></div>`
	_t.Body += `<div class="navbar-end"></div>`
	return _t
}

func (_burger *NavbarMenu) AddListeners() {
}

// Navbar is an UISnippet registered with the ick-tag `ick-navbar`.
type Navbar struct {
	dom.UISnippet
}

func (_nav *Navbar) Template(*html.DataState) (_t html.SnippetTemplate) {
	_t.TagName = "nav"
	_t.Attributes = `class="navbar" role="navigation"`

	menu := &NavbarMenu{}
	menuid := "menu" + _nav.Id()
	menu.SetId(html.String(menuid))

	burger := &NavbarBurger{NavbarMenuID: menuid}

	_t.Body = `<div class="navbar-brand">`
	_t.Body += `<a class="navbar-item" href="https://bulma.io">
				<img src="https://bulma.io/images/bulma-logo.png" width="112" height="28">
				</a>`
	_t.Body += _nav.RenderChildHTML(burger)
	_t.Body += _nav.RenderChildHTML(menu)
	_t.Body += `</div>`
	return _t
}

func (_nav *Navbar) AddListeners() {

}
