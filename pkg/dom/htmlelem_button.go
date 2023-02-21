package dom

import (
	"syscall/js"
)

// A string indicating the behavior of the button
type BUTTON_TYPE string

const (
	BT_SUBMIT BUTTON_TYPE = "submit" // The button submits the form. This is the default value if the attribute is not specified, or if it is dynamically changed to an empty or invalid value.
	BT_RESET  BUTTON_TYPE = "reset"  // The button resets the form.
	BT_BUTTON BUTTON_TYPE = "button" // The button does nothing with the form.
	BT_MENU   BUTTON_TYPE = "menu"   // The button displays a menu.
)

/****************************************************************************
* HTMLButtonElement
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLButton
type HTMLButton struct {
	Element
}

// CastHTMLButton is casting a js.Value into HTMLButtonElement.
func CastHTMLButton(_value js.Value) *HTMLButton {
	if _value.Type() != js.TypeObject || _value.Get("tagName").String() != "BUTTON" {
		ICKError("casting HTMLButton failed")
		return new(HTMLButton)
	}
	cast := new(HTMLButton)
	cast.jsValue = _value
	return cast
}

/****************************************************************************
* HTMLButtonElement's properties
*****************************************************************************/

// IsAutofocus returns a boolean value indicating whether or not the control should have input focus when the page loads,
// unless the user overrides it, for example by typing in a different control. Only one form-associated element in a document can have this attribute specified.
func (_btn *HTMLButton) IsAutofocus() bool {
	if !_btn.IsDefined() {
		return false
	}
	return _btn.jsValue.Get("autofocus").Bool()
}

// SetAutofocus setting attribute 'autofocus' with
// type bool (idl: boolean).
func (_btn *HTMLButton) SetAutofocus(value bool) *HTMLButton {
	if !_btn.IsDefined() {
		return nil
	}
	_btn.jsValue.Set("autofocus", value)
	return _btn
}

// IsDisabled returns a boolean value indicating whether or not the control is disabled, meaning that it does not accept any clicks.
func (_btn *HTMLButton) IsDisabled() bool {
	if !_btn.IsDefined() {
		return false
	}
	return _btn.jsValue.Get("disabled").Bool()
}

// SetDisabled setting attribute 'disabled' with
// type bool (idl: boolean).
func (_btn *HTMLButton) SetDisabled(value bool) *HTMLButton {
	if !_btn.IsDefined() {
		return nil
	}
	_btn.jsValue.Set("disabled", value)
	return _btn
}

// Name of the object when submitted with a form. If specified, it must not be the empty string.
func (_btn *HTMLButton) Name() string {
	if !_btn.IsDefined() {
		return UNDEFINED_NODE
	}
	return _btn.jsValue.Get("name").String()
}

// Name of the object when submitted with a form. If specified, it must not be the empty string.
func (_btn *HTMLButton) SetName(name string) {
	if !_btn.IsDefined() {
		return
	}
	_btn.jsValue.Set("name", name)
}

// Value A string representing the current form control value of the button.
func (_btn *HTMLButton) Value() string {
	if !_btn.IsDefined() {
		return UNDEFINED_NODE
	}
	return _btn.jsValue.Get("value").String()
}

// Value A string representing the current form control value of the button.
func (_btn *HTMLButton) SetValue(value string) {
	if !_btn.IsDefined() {
		return
	}
	_btn.jsValue.Set("value", value)
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_btn *HTMLButton) Type() BUTTON_TYPE {
	if !_btn.IsDefined() {
		return BT_SUBMIT
	}
	typ := _btn.jsValue.Get("type").String()
	return BUTTON_TYPE(typ)
}

// SetType setting attribute 'type' with
// type string (idl: DOMString).
func (_btn *HTMLButton) SetType(_typ BUTTON_TYPE) {
	if !_btn.IsDefined() {
		return
	}
	_btn.jsValue.Set("type", string(_typ))
}

// WillValidate returns a boolean value indicating whether the button is a candidate for constraint validation.
// It is false if any conditions bar it from constraint validation, including: its type property is reset or button; *
// it has a <datalist> ancestor; or the disabled property is set to true.
func (_btn *HTMLButton) WillValidate() bool {
	if !_btn.IsDefined() {
		return false
	}
	return _btn.jsValue.Get("willValidate").Bool()
}

// ValidationMessage a string representing the localized message that describes the validation constraints
// that the control does not satisfy (if any). This attribute is the empty string if the control is not a
// candidate for constraint validation (willValidate is false), or it satisfies its constraints.
func (_btn *HTMLButton) ValidationMessage() string {
	if !_btn.IsDefined() {
		return UNDEFINED_NODE
	}
	return _btn.jsValue.Get("validationMessage").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLButtonElement/labels
func (_btn *HTMLButton) Labels() []*Node {
	if !_btn.IsDefined() {
		return make([]*Node, 0)
	}
	nodes := _btn.jsValue.Get("labels")
	return MakeNodes(nodes)
}
