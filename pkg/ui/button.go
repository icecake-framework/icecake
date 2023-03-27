package ui

import (
	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/js"
)

// A string indicating the behavior of the button
type BUTTON_TYPE string

const (
	BTNT_SUBMIT BUTTON_TYPE = "submit" // The button submits the form. This is the default value if the attribute is not specified, or if it is dynamically changed to an empty or invalid value.
	BTNT_RESET  BUTTON_TYPE = "reset"  // The button resets the form.
	BTNT_BUTTON BUTTON_TYPE = "button" // The button does nothing with the form.
	BTNT_MENU   BUTTON_TYPE = "menu"   // The button displays a menu.
)

/****************************************************************************
* HTMLButtonElement
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Button
type Button struct {
	dom.Element
}

// CastButton is casting a js.Value into HTMLButtonElement.
func CastButton(_jsvp js.JSValueProvider) *Button {
	if _jsvp.Value().Type() != js.TYPE_OBJECT || _jsvp.Value().GetString("tagName") != "BUTTON" {
		console.Errorf("casting HTMLButton failed")
		return &Button{}
	}
	cast := new(Button)
	cast.JSValue = _jsvp.Value()
	return cast
}

// GetButtonById returns a Button corresponding to the existing _id into the DOM,
// otherwhise returns an undefined Button.
func ButtonById(_id string) *Button {
	_id = helper.Normalize(_id)
	jse := dom.Doc().ChildById(_id)
	if jse.IsObject() && jse.GetString("tagName") == "BUTTON" {
		btn := new(Button)
		btn.JSValue = jse.JSValue
		return btn
	}

	console.Warnf("GetElementById failed: %q not found, or not a <Element>", _id)
	return new(Button)
}

/****************************************************************************
* HTMLButtonElement's properties
*****************************************************************************/

// IsAutofocus returns a boolean value indicating whether or not the control should have input focus when the page loads,
// unless the user overrides it, for example by typing in a different control. Only one form-associated element in a document can have this attribute specified.
func (_btn *Button) IsAutofocus() bool {
	if !_btn.IsDefined() {
		return false
	}
	return _btn.GetBool("autofocus")
}

// SetAutofocus setting attribute 'autofocus' with
// type bool (idl: boolean).
func (_btn *Button) SetAutofocus(value bool) *Button {
	if !_btn.IsDefined() {
		return nil
	}
	_btn.Set("autofocus", value)
	return _btn
}

// IsDisabled returns a boolean value indicating whether or not the control is disabled, meaning that it does not accept any clicks.
func (_btn *Button) IsDisabled() bool {
	if !_btn.IsDefined() {
		return false
	}
	return _btn.GetBool("disabled")
}

// SetDisabled setting attribute 'disabled' with
// type bool (idl: boolean).
func (_btn *Button) SetDisabled(value bool) *Button {
	if !_btn.IsDefined() {
		return nil
	}
	_btn.Set("disabled", value)
	return _btn
}

// Name of the object when submitted with a form. If specified, it must not be the empty string.
func (_btn *Button) Name() string {
	if !_btn.IsDefined() {
		return dom.UNDEFINED_NODE
	}
	return _btn.GetString("name")
}

// Name of the object when submitted with a form. If specified, it must not be the empty string.
func (_btn *Button) SetName(name string) {
	if !_btn.IsDefined() {
		return
	}
	_btn.Set("name", name)
}

// Value A string representing the current form control value of the button.
func (_btn *Button) Value() string {
	if !_btn.IsDefined() {
		return dom.UNDEFINED_NODE
	}
	return _btn.GetString("value")
}

// Value A string representing the current form control value of the button.
func (_btn *Button) SetValue(value string) {
	if !_btn.IsDefined() {
		return
	}
	_btn.Set("value", value)
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_btn *Button) Type() BUTTON_TYPE {
	if !_btn.IsDefined() {
		return BTNT_SUBMIT
	}
	typ := _btn.GetString("type")
	return BUTTON_TYPE(typ)
}

// SetType setting attribute 'type' with
// type string (idl: DOMString).
func (_btn *Button) SetType(_typ BUTTON_TYPE) {
	if !_btn.IsDefined() {
		return
	}
	_btn.Set("type", string(_typ))
}

// WillValidate returns a boolean value indicating whether the button is a candidate for constraint validation.
// It is false if any conditions bar it from constraint validation, including: its type property is reset or button; *
// it has a <datalist> ancestor; or the disabled property is set to true.
func (_btn *Button) WillValidate() bool {
	if !_btn.IsDefined() {
		return false
	}
	return _btn.GetBool("willValidate")
}

// ValidationMessage a string representing the localized message that describes the validation constraints
// that the control does not satisfy (if any). This attribute is the empty string if the control is not a
// candidate for constraint validation (willValidate is false), or it satisfies its constraints.
func (_btn *Button) ValidationMessage() string {
	if !_btn.IsDefined() {
		return dom.UNDEFINED_NODE
	}
	return _btn.GetString("validationMessage")
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLButtonElement/labels
func (_btn *Button) Labels() []*dom.Element {
	if !_btn.IsDefined() {
		return make([]*dom.Element, 0)
	}
	elems := _btn.Get("labels")
	return dom.CastElements(elems)
}
