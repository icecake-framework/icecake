package ui

import (
	"net/url"

	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-button", &Button{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

// The button type can be one of the following:
//
//	BTN_TYPE_BUTTON  // <button> form buttons.
//	BTN_TYPE_A       // <a> anchor links
//	BTN_TYPE_SUBMIT  // <input type="submit"> submit inputs
//	BTN_TYPE_RESET   // <input type="reset"> reset inputs
type BUTTON_TYPE int

const (
	BTN_TYPE_BUTTON BUTTON_TYPE = iota // <button> form buttons.
	BTN_TYPE_A                         // <a> anchor links
	BTN_TYPE_SUBMIT                    // <input type="submit"> submit inputs
	BTN_TYPE_RESET                     // <input type="reset"> reset inputs
)

// Button is an UISnippet registered with the ick-tag `ick-button`.
//
// According to the ButtonType property, Button can be used either as a standard <button> element but also as an anchor link or a submit or reset form input.
// The core text is handle with the Title html property, allowing simple text or complex rendering.
//
// The IsDisabled property is directly handled by the embedded UISnippet.
type Button struct {
	dom.UISnippet

	// The button type.
	// If nothing is specified, the default ButtonType is BTN_TYPE_BUTTON.
	ButtonType BUTTON_TYPE

	// If the ButtonType is BTN_TYPE_A then HRef defines the associated url link.
	// HRef has no effect on other ButtonType. If HRef is defined then the button type is automatically set to BTN_TYPE_A.
	HRef *url.URL

	// The title of the Button. Can be a simple text or a more complex html string.
	Title html.String

	IsOutlined bool // Outlined button style
	IsRounded  bool // Rounded button style

	IsLoading bool // Loading button state

	// TODO: handles buttons properties for color, size, display
}

func (_btn *Button) Template(*html.DataState) (_t html.SnippetTemplate) {
	href := ""
	if _btn.HRef != nil {
		href = _btn.HRef.String()
	}
	if href != "" {
		_btn.ButtonType = BTN_TYPE_A
	}
	switch _btn.ButtonType {
	case BTN_TYPE_A:
		_t.TagName = "a"
		_t.Attributes = `class="button"`
		_t.Attributes += html.String(` href="` + href + `"`)
	case BTN_TYPE_SUBMIT:
		_t.TagName = "input"
		_t.Attributes = `class="button" type="submit"`
	case BTN_TYPE_RESET:
		_t.TagName = "input"
		_t.Attributes = `class="button" type="reset"`
	default:
		_t.TagName = "button"
		_t.Attributes = `class="button"`
	}
	_t.Body = _btn.Title

	if _btn.IsOutlined {
		_btn.SetClasses("is-outlined")
	}
	if _btn.IsRounded {
		_btn.SetClasses("is-rounded")
	}
	if _btn.IsLoading {
		_btn.SetClasses("is-loading")
	}
	_btn.SetDisabled(_btn.IsDisabled())

	// TODO: finalize button

	return _t
}

func (_btn *Button) SetLoading(_f bool) {
	_btn.IsLoading = _f
	if _f {
		_btn.DOM.SetClasses("is-loading")
	} else {
		_btn.DOM.RemoveClasses("is-loading")
	}
}

func (_btn *Button) SetRounded(_f bool) {
	_btn.IsRounded = _f
	if _f {
		_btn.DOM.SetClasses("is-rounded")
	} else {
		_btn.DOM.RemoveClasses("is-rounded")
	}
}

func (_btn *Button) SetOutlined(_f bool) {
	_btn.IsOutlined = _f
	if _f {
		_btn.DOM.SetClasses("is-outlined")
	} else {
		_btn.DOM.RemoveClasses("is-outlined")
	}
}
