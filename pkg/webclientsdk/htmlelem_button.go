package browser

import (
	"syscall/js"
)

/****************************************************************************
* HTMLButtonElement
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLButton
type HTMLButton struct {
	HTMLElement
}

// NewHTMLButtonFromJS is casting a js.Value into HTMLButtonElement.
func NewHTMLButtonFromJS(value js.Value) *HTMLButton {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &HTMLButton{}
	ret.jsValue = value
	return ret
}

/****************************************************************************
* HTMLButtonElement's properties
*****************************************************************************/

// Autofocus returns a boolean value indicating whether or not the control should have input focus when the page loads,
// unless the user overrides it, for example by typing in a different control. Only one form-associated element in a document can have this attribute specified.
func (_this *HTMLButton) Autofocus() bool {
	var ret bool
	value := _this.jsValue.Get("autofocus")
	ret = (value).Bool()
	return ret
}

// SetAutofocus setting attribute 'autofocus' with
// type bool (idl: boolean).
func (_this *HTMLButton) SetAutofocus(value bool) *HTMLButton {
	_this.jsValue.Set("autofocus", value)
	return _this
}

// Disabled returns a boolean value indicating whether or not the control is disabled, meaning that it does not accept any clicks.
func (_this *HTMLButton) Disabled() bool {
	value := _this.jsValue.Get("disabled")
	return (value).Bool()
}

// SetDisabled setting attribute 'disabled' with
// type bool (idl: boolean).
func (_this *HTMLButton) SetDisabled(value bool) *HTMLButton {
	_this.jsValue.Set("disabled", value)
	return _this
}

// Form returning attribute 'form' with
// type HTMLFormElement (idl: HTMLFormElement).
// func (_this *HTMLButtonElement) Form() *HTMLFormElement {
// 	var ret *HTMLFormElement
// 	value := _this.jsValue.Get("form")
// 	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
// 		ret = HTMLFormElementFromJS(value)
// 	}
// 	return ret
// }

// Name returning attribute 'name' with
// type string (idl: DOMString).
func (_this *HTMLButton) Name() string {
	var ret string
	value := _this.jsValue.Get("name")
	ret = (value).String()
	return ret
}

// SetName setting attribute 'name' with
// type string (idl: DOMString).
func (_this *HTMLButton) SetName(value string) {
	input := value
	_this.jsValue.Set("name", input)
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
func (_this *HTMLButton) Type() string {
	var ret string
	value := _this.jsValue.Get("type")
	ret = (value).String()
	return ret
}

// SetType setting attribute 'type' with
// type string (idl: DOMString).
func (_this *HTMLButton) SetType(value string) {
	input := value
	_this.jsValue.Set("type", input)
}

// Value returning attribute 'value' with
// type string (idl: DOMString).
func (_this *HTMLButton) Value() string {
	var ret string
	value := _this.jsValue.Get("value")
	ret = (value).String()
	return ret
}

// SetValue setting attribute 'value' with
// type string (idl: DOMString).
func (_this *HTMLButton) SetValue(value string) {
	input := value
	_this.jsValue.Set("value", input)
}

// WillValidate returning attribute 'willValidate' with
// type bool (idl: boolean).
func (_this *HTMLButton) WillValidate() bool {
	var ret bool
	value := _this.jsValue.Get("willValidate")
	ret = (value).Bool()
	return ret
}

// Validity returning attribute 'validity' with
// type ValidityState (idl: ValidityState).
// func (_this *HTMLButtonElement) Validity() *ValidityState {
// 	var ret *ValidityState
// 	value := _this.jsValue.Get("validity")
// 	ret = ValidityStateFromJS(value)
// 	return ret
// }

// ValidationMessage returning attribute 'validationMessage' with
// type string (idl: DOMString).
func (_this *HTMLButton) ValidationMessage() string {
	var ret string
	value := _this.jsValue.Get("validationMessage")
	ret = (value).String()
	return ret
}

// Labels returning attribute 'labels' with
// type dom.NodeList (idl: NodeList).
// func (_this *HTMLButtonElement) Labels() *dom.NodeList {
// 	var ret *dom.NodeList
// 	value := _this.jsValue.Get("labels")
// 	ret = dom.NodeListFromJS(value)
// 	return ret
// }

func (_this *HTMLButton) CheckValidity() (_result bool) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("checkValidity", _args[0:_end]...)
	return (_returned).Bool()
}

func (_this *HTMLButton) ReportValidity() (_result bool) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("reportValidity", _args[0:_end]...)
	return (_returned).Bool()
}

func (_this *HTMLButton) SetCustomValidity(_error string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := _error
	_args[0] = _p0
	_end++
	_this.jsValue.Call("setCustomValidity", _args[0:_end]...)
}
