package ui

import (
	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/html"
	"github.com/sunraylab/icecake/pkg/ick"
)

func init() {
	ick.RegisterComposer("ick-button", &Button{})
}

type BUTTON_TYPE int

const (
	BTN_TYPE_BUTTON BUTTON_TYPE = iota // <button> form buttons. The default value
	BTN_TYPE_A                         // <a> anchor links
	BTN_TYPE_SUBMIT                    // <input type="submit"> submit inputs
	BTN_TYPE_RESET                     // <input type="reset"> reset inputs
)

type Button struct {
	dom.UISnippet

	ButtonType BUTTON_TYPE

	Title html.String

	IsOutlined bool
	IsRounded  bool
	IsLoading  bool
	IsDisabled bool
}

func (_btn *Button) Template(*html.DataState) (_t html.SnippetTemplate) {
	switch _btn.ButtonType {
	case BTN_TYPE_A:
		_t.TagName = "a"
		_t.Attributes = `class="button"`
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
	if _btn.IsDisabled {
		_btn.SetAttribute("disabled", "", true)
	}

	// TODO: finalize button

	return _t
}

func (_btn *Button) Loading(_f bool) {
	_btn.IsLoading = _f
	if _f {
		_btn.DOM.Classes().AddTokens("is-loading")
	} else {
		_btn.DOM.Classes().RemoveTokens("is-loading")
	}
}

func (_btn *Button) Rounded(_f bool) {
	_btn.IsRounded = _f
	if _f {
		_btn.DOM.Classes().AddTokens("is-rounded")
	} else {
		_btn.DOM.Classes().RemoveTokens("is-rounded")
	}
}

func (_btn *Button) Outlined(_f bool) {
	_btn.IsOutlined = _f
	if _f {
		_btn.DOM.Classes().AddTokens("is-outlined")
	} else {
		_btn.DOM.Classes().RemoveTokens("is-outlined")
	}
}
