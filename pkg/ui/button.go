package ui

import (
	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
)

// Button is an UISnippet registered with the ick-tag `ick-button`.
//
// According to the ButtonType property, Button can be used either as a standard <button> element but also as an anchor link or a submit or reset form input.
// The core text is handle with the Title html property, allowing simple text or complex rendering.
//
// The IsDisabled property is directly handled by the embedded UISnippet.
type Button struct {
	*bulma.Button
	DOM dom.Element
}

func NewButton(title html.HTMLString, id string, rawURL string, attrs ...string) *Button {
	btn := new(Button)
	btn.Button = bulma.NewButton(title, id, rawURL, attrs...)
	return btn
}

func (btn *Button) SetLoading(_f bool) *Button {
	btn.IsLoading = _f
	if _f {
		btn.DOM.SetClasses("is-loading")
	} else {
		btn.DOM.RemoveClasses("is-loading")
	}
	return btn
}

func (btn *Button) SetRounded(_f bool) *Button {
	btn.IsRounded = _f
	if _f {
		btn.DOM.SetClasses("is-rounded")
	} else {
		btn.DOM.RemoveClasses("is-rounded")
	}
	return btn
}

func (btn *Button) SetOutlined(_f bool) *Button {
	btn.IsOutlined = _f
	if _f {
		btn.DOM.SetClasses("is-outlined")
	} else {
		btn.DOM.RemoveClasses("is-outlined")
	}
	return btn
}

func (btn *Button) SetDisabled(disabled bool) {
	btn.IsDisabled = disabled
	if btn.DOM.IsDefined() && btn.DOM.IsInDOM() {
		btn.DOM.SetDisabled(disabled)
	}
}
