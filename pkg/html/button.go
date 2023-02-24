package html

import (
	"syscall/js"

	"github.com/sunraylab/icecake/internal/helper"
	ick "github.com/sunraylab/icecake/pkg/icecake"
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
	ick.Element
}

// CastButton is casting a js.Value into HTMLButtonElement.
func CastButton(_value js.Value) *Button {
	if _value.Type() != js.TypeObject || _value.Get("tagName").String() != "BUTTON" {
		ick.ConsoleErrorf("casting HTMLButton failed")
		return new(Button)
	}
	cast := new(Button)
	cast.Wrap(_value)
	return cast
}

// GetButtonById returns a Button corresponding to the existing _id into the DOM,
// otherwhise returns an undefined Button.
func GetButtonById(_id string) *Button {
	_id = helper.Normalize(_id)
	jse := ick.GetElementById(_id)
	if etyp := jse.JSValue().Type(); etyp != js.TypeNull && etyp != js.TypeUndefined {
		if jse.JSValue().Get("tagName").String() == "BUTTON" {
			btn := new(Button)
			btn.Wrap(jse.JSValue())
			return btn
		}
	}
	ick.ConsoleWarnf("GetElementById failed: %q not found, or not a <Element>", _id)
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
	return _btn.JSValue().Get("autofocus").Bool()
}

// SetAutofocus setting attribute 'autofocus' with
// type bool (idl: boolean).
func (_btn *Button) SetAutofocus(value bool) *Button {
	if !_btn.IsDefined() {
		return nil
	}
	_btn.JSValue().Set("autofocus", value)
	return _btn
}

// IsDisabled returns a boolean value indicating whether or not the control is disabled, meaning that it does not accept any clicks.
func (_btn *Button) IsDisabled() bool {
	if !_btn.IsDefined() {
		return false
	}
	return _btn.JSValue().Get("disabled").Bool()
}

// SetDisabled setting attribute 'disabled' with
// type bool (idl: boolean).
func (_btn *Button) SetDisabled(value bool) *Button {
	if !_btn.IsDefined() {
		return nil
	}
	_btn.JSValue().Set("disabled", value)
	return _btn
}

// Name of the object when submitted with a form. If specified, it must not be the empty string.
func (_btn *Button) Name() string {
	if !_btn.IsDefined() {
		return ick.UNDEFINED_NODE
	}
	return _btn.JSValue().Get("name").String()
}

// Name of the object when submitted with a form. If specified, it must not be the empty string.
func (_btn *Button) SetName(name string) {
	if !_btn.IsDefined() {
		return
	}
	_btn.JSValue().Set("name", name)
}

// Value A string representing the current form control value of the button.
func (_btn *Button) Value() string {
	if !_btn.IsDefined() {
		return ick.UNDEFINED_NODE
	}
	return _btn.JSValue().Get("value").String()
}

// Value A string representing the current form control value of the button.
func (_btn *Button) SetValue(value string) {
	if !_btn.IsDefined() {
		return
	}
	_btn.JSValue().Set("value", value)
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_btn *Button) Type() BUTTON_TYPE {
	if !_btn.IsDefined() {
		return BTNT_SUBMIT
	}
	typ := _btn.JSValue().Get("type").String()
	return BUTTON_TYPE(typ)
}

// SetType setting attribute 'type' with
// type string (idl: DOMString).
func (_btn *Button) SetType(_typ BUTTON_TYPE) {
	if !_btn.IsDefined() {
		return
	}
	_btn.JSValue().Set("type", string(_typ))
}

// WillValidate returns a boolean value indicating whether the button is a candidate for constraint validation.
// It is false if any conditions bar it from constraint validation, including: its type property is reset or button; *
// it has a <datalist> ancestor; or the disabled property is set to true.
func (_btn *Button) WillValidate() bool {
	if !_btn.IsDefined() {
		return false
	}
	return _btn.JSValue().Get("willValidate").Bool()
}

// ValidationMessage a string representing the localized message that describes the validation constraints
// that the control does not satisfy (if any). This attribute is the empty string if the control is not a
// candidate for constraint validation (willValidate is false), or it satisfies its constraints.
func (_btn *Button) ValidationMessage() string {
	if !_btn.IsDefined() {
		return ick.UNDEFINED_NODE
	}
	return _btn.JSValue().Get("validationMessage").String()
}

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLButtonElement/labels
func (_btn *Button) Labels() []*ick.Node {
	if !_btn.IsDefined() {
		return make([]*ick.Node, 0)
	}
	nodes := _btn.JSValue().Get("labels")
	return ick.MakeNodes(nodes)
}

/*****************************************************************************
* ICButton
******************************************************************************/

// ICButton is an extension of the ick.HTMLElement
// type ICButton struct {
// 	ick.HTMLButton
// }

// func CastICButton(_value js.Value) *ICButton {
// 	if _value.Type() != js.TypeObject || _value.Get("tagName").String() != "BUTTON" {
// 		ick.ConsoleError("casting ICButton failed")
// 		return new(ICButton)
// 	}
// 	ret := new(ICButton)
// 	ret.Wrap(_value)
// 	return ret
// }
