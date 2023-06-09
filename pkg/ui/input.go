package ui

import (
	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/html"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	html.RegisterComposer("ick-input", &InputField{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type INPUT_STATE string

const (
	INPUT_NORMAL   INPUT_STATE = "normal" // the default input state
	INPUT_SUCCESS  INPUT_STATE = "success"
	INPUT_WARNING  INPUT_STATE = "warning"
	INPUT_ERROR    INPUT_STATE = "error"
	INPUT_LOADING  INPUT_STATE = "loading"
	INPUT_READONLY INPUT_STATE = "readlonly"
)

type InputField struct {
	dom.UISnippet

	Label       html.String // Optional
	PlaceHolder html.String // Optional
	Help        html.String // Optional

	// The input value
	Value string

	// icon left
	IsRounded bool // Rounded style

	State INPUT_STATE

	uilabel   dom.UISnippet
	uicontrol dom.UISnippet
	uiinput   dom.UISnippet
	uihelp    dom.UISnippet
}

func (_in *InputField) Template(_data *html.DataState) (_t html.SnippetTemplate) {
	_t.TagName = "div"
	_t.Attributes = `class="field"`

	// <label>
	_in.uilabel.TagName = "label"
	_in.uilabel.SetClasses("label")
	if _in.Label != "" {
		_in.uilabel.Body = _in.Label
		_t.Body += _in.RenderHTML(&_in.uilabel)
	}

	// <input>
	_in.uiinput.Tag = html.Tag{"input", true}
	_in.uiinput.SetClasses("input").SetClassesIf(_in.IsRounded, "is-rounded")
	_in.uiinput.SetAttribute("type", "text")
	_in.uiinput.SetAttributeIf(_in.PlaceHolder != "", "placeholder", _in.PlaceHolder)
	switch _in.State {
	case "success":
		_in.uiinput.SetClasses("is-success")
	case "warning":
		_in.uiinput.SetClasses("is-warning")
	case "error":
		_in.uiinput.SetClasses("is-danger")
	case "readonly":
		_in.uiinput.SetAttribute("readonly", "")
	}
	_in.uiinput.SetAttributeIf(_in.Value != "", "value", _in.Value)

	// <div control>
	_in.uicontrol.TagName = "div"
	_in.uicontrol.SetClasses("control")
	_in.uicontrol.SetClassesIf(_in.State == INPUT_LOADING, "is-loading")
	_in.uicontrol.Body = _in.RenderHTML(&_in.uiinput)

	_t.Body += _in.RenderHTML(&_in.uicontrol)

	// <p help>
	_in.uihelp.TagName = "p"
	_in.uihelp.SetClasses("help")
	if _in.Help != "" {
		switch _in.State {
		case "success":
			_in.uihelp.SetClasses("is-success")
		case "warning":
			_in.uihelp.SetClasses("is-warning")
		case "error":
			_in.uihelp.SetClasses("is-danger")
		}
		_in.uihelp.Body = _in.Help
		_t.Body += _in.RenderHTML(&_in.uihelp)
	}

	return
}

func (_in *InputField) AddListeners() {
	_in.uiinput.DOM.AddInputEvent(event.INPUT_ONINPUT, _in.OnInput)

}

func (_in *InputField) OnInput(_event *event.InputEvent, _target *dom.Element) {
	_in.Value = _target.JSValue.String()
	console.Warnf("InputField: %s", _in.Value)
}

func (_in *InputField) SetState(_state INPUT_STATE) *InputField {
	_in.State = _state
	if _in.DOM.IsInDOM() {
		switch _in.State {
		case "success":
			_in.uicontrol.DOM.RemoveClasses("is-loading")
			_in.uiinput.DOM.SwitchClasses("is-warning is-danger", "is-success")
			_in.uiinput.DOM.RemoveAttribute("readonly")
			_in.uihelp.DOM.SwitchClasses("is-warning is-danger", "is-success")
		case "warning":
			_in.uicontrol.DOM.RemoveClasses("is-loading")
			_in.uiinput.DOM.SwitchClasses("is-success is-danger", "is-warning")
			_in.uiinput.DOM.RemoveAttribute("readonly")
			_in.uihelp.DOM.SwitchClasses("is-success is-danger", "is-warning")
		case "error":
			_in.uicontrol.DOM.RemoveClasses("is-loading")
			_in.uiinput.DOM.SwitchClasses("is-warning is-success", "is-danger")
			_in.uiinput.DOM.RemoveAttribute("readonly")
			_in.uihelp.DOM.SwitchClasses("is-warning is-success", "is-danger")
		case "loading":
			_in.uicontrol.DOM.SetClasses("is-loading")
			_in.uiinput.DOM.RemoveAttribute("readonly")
			_in.uiinput.DOM.RemoveClasses("is-success is-warning is-danger")
			_in.uihelp.DOM.RemoveClasses("is-success is-warning is-danger")
		case "readonly":
			_in.uicontrol.DOM.RemoveClasses("is-loading")
			_in.uiinput.DOM.SetAttribute("readonly", "")
			_in.uiinput.DOM.RemoveClasses("is-success is-warning is-danger")
			_in.uihelp.DOM.RemoveClasses("is-success is-warning is-danger")
		default:
			_in.uicontrol.DOM.RemoveClasses("is-loading")
			_in.uiinput.DOM.RemoveAttribute("readonly")
			_in.uiinput.DOM.RemoveClasses("is-success is-warning is-danger")
			_in.uihelp.DOM.RemoveClasses("is-success is-warning is-danger")
		}
	}
	return _in
}

func (_in *InputField) SetHelp(_help html.String, _state INPUT_STATE) *InputField {
	_in.Help = _help
	_in.uihelp.DOM.InsertHTML(dom.INSERT_BODY, _help, nil)
	_in.SetState(_state)
	return _in
}

func (_in *InputField) SetRounded(_f bool) *InputField {
	_in.IsRounded = _f
	if _f {
		_in.uiinput.DOM.SetClasses("is-rounded")
	} else {
		_in.uiinput.DOM.RemoveClasses("is-rounded")
	}
	return _in
}
