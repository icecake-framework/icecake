package dom

import (
	"syscall/js"
)

/****************************************************************************
* HTMLElement
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement
type HTMLElement struct {
	Element
}

// CastHTMLElement is casting a js.Value into HTMLElement.
func CastHTMLElement(value js.Value) *HTMLElement {
	if value.Type() != js.TypeObject {
		ICKError("casting HTMLElement failed")
		return new(HTMLElement)
	}
	cast := new(HTMLElement)
	cast.jsValue = value
	return cast
}

/****************************************************************************
* HTMLElement's properties & methods
*****************************************************************************/

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (_elem *HTMLElement) AccessKey() string {
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return _elem.jsValue.Get("accessKey").String()
}

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (_htmle *HTMLElement) SetAccessKey(key bool) *HTMLElement {
	if !_htmle.IsDefined() {
		return nil
	}
	_htmle.jsValue.Set("accessKey", key)
	return _htmle
}

// InnerText represents the rendered text content of a node and its descendants.
//
// InnerText gets pure text, removing any html or css, while TextContent keeps the representation.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (_htmle *HTMLElement) InnerText() string {
	if !_htmle.IsDefined() {
		return UNDEFINED_NODE
	}
	var ret string
	value := _htmle.jsValue.Get("innerText")
	ret = (value).String()
	return ret
}

// InnerText represents the rendered text content of a node and its descendants.
//
// InnerText gets pure text, removing any html or css, while TextContent keeps the representation.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (_htmle *HTMLElement) SetInnerText(value string) {
	if !_htmle.IsDefined() {
		return
	}
	input := value
	_htmle.jsValue.Set("innerText", input)
}

// Focus sets focus on the specified element, if it can be focused. The focused element is the element that will receive keyboard and similar events by default.
//
// By default the browser will scroll the element into view after focusing it,
// and it may also provide visible indication of the focused element (typically by displaying a "focus ring" around the element).
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus
func (_htmle *HTMLElement) Focus() {
	if !_htmle.IsDefined() {
		return
	}
	_htmle.jsValue.Call("focus")
}

// Blur removes keyboard focus from the current element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/blur
func (_htmle *HTMLElement) Blur() {
	if !_htmle.IsDefined() {
		return
	}
	_htmle.jsValue.Call("blur")
}

/****************************************************************************
* HTMLElement's GENERIC_EVENT
*****************************************************************************/

// event attribute: Event
func makeHTMLElement_domcore_Event(listener func(event *Event, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastEvent(value)
		target := CastHTMLElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddAbort is adding doing AddEventListener for 'Abort' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *HTMLElement) AddGenericEvent(evttype GENERIC_EVENT, listener func(event *Event, target *HTMLElement)) js.Func {
	if !_htmle.IsDefined() {
		ICKWarn("AddGenericEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_domcore_Event(listener)
	_htmle.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/****************************************************************************
* HTMLElement's MOUSE_EVENT
*****************************************************************************/

// event attribute: MouseEvent
func makeHTMLElement_Mouse_Event(listener func(event *MouseEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastMouseEvent(value)
		target := CastHTMLElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddClick is adding doing AddEventListener for 'Click' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *HTMLElement) AddMouseEvent(evttype MOUSE_EVENT, listener func(event *MouseEvent, target *HTMLElement)) js.Func {
	if !_htmle.IsDefined() {
		ICKWarn("AddMouseEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_Mouse_Event(listener)
	_htmle.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/****************************************************************************
* HTMLElement's FOCUS_EVENT
*****************************************************************************/

// event attribute: FocusEvent
func makeHTMLElement_FocusEvent(listener func(event *FocusEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastFocusEvent(value)
		target := CastHTMLElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddBlur is adding doing AddEventListener for 'Blur' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *HTMLElement) AddFocusEvent(evttype FOCUS_EVENT, listener func(event *FocusEvent, target *HTMLElement)) js.Func {
	if !_htmle.IsDefined() {
		ICKWarn("AddFocusEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_FocusEvent(listener)
	_htmle.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/****************************************************************************
* HTMLElement's POINTER_EVENT
*****************************************************************************/

// event attribute: PointerEvent
func makeHTMLElement_PointerEvent(listener func(event *PointerEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var evt *PointerEvent
		value := args[0]
		evt = CastPointerEvent(value)
		target := CastHTMLElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddGotPointerCapture is adding doing AddEventListener for 'GotPointerCapture' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *HTMLElement) AddPointerEvent(evttype POINTER_EVENT, listener func(event *PointerEvent, target *HTMLElement)) js.Func {
	if !_htmle.IsDefined() {
		ICKWarn("AddPointerEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_PointerEvent(listener)
	_htmle.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/****************************************************************************
* HTMLElement's INPUT_EVENT
*****************************************************************************/

// event attribute: InputEvent
func makeHTMLElement_InputEvent(listener func(event *InputEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastInputEvent(value)
		target := CastHTMLElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddInput is adding doing AddEventListener for 'Input' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *HTMLElement) AddInputEvent(evttype INPUT_EVENT, listener func(event *InputEvent, target *HTMLElement)) js.Func {
	if !_htmle.IsDefined() {
		ICKWarn("AddInputEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_InputEvent(listener)
	_htmle.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/****************************************************************************
* HTMLElement's KEYBOARD_EVENT
*****************************************************************************/

// event attribute: KeyboardEvent
func makeHTMLElement_KeyboardEvent(listener func(event *KeyboardEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastKeyboardEvent(value)
		target := CastHTMLElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddKeyDown is adding doing AddEventListener for 'KeyDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *HTMLElement) AddKeyboard(evttype KEYBOARD_EVENT, listener func(event *KeyboardEvent, target *HTMLElement)) js.Func {
	if !_htmle.IsDefined() {
		ICKWarn("AddKeyboard not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_KeyboardEvent(listener)
	_htmle.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/****************************************************************************
* HTMLElement's UI_EVENT
*****************************************************************************/

// event attribute: UIEvent
func makeHTMLElement_UIEvent(listener func(event *UIEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastUIEvent(value)
		target := CastHTMLElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddResize is adding doing AddEventListener for 'Resize' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *HTMLElement) AddResizeEvent(listener func(event *UIEvent, target *HTMLElement)) js.Func {
	if !_htmle.IsDefined() {
		ICKWarn("AddResizeEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_UIEvent(listener)
	_htmle.jsValue.Call("addEventListener", "resize", callback)
	return callback
}

/****************************************************************************
* HTMLElement's WHEEL_EVENT
*****************************************************************************/

// event attribute: WheelEvent
func makeHTMLElement_WheelEvent(listener func(event *WheelEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastWheelEvent(value)
		target := CastHTMLElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// The wheel event fires when the user rotates a wheel button on a pointing device (typically a mouse).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/wheel_event
func (_htmle *HTMLElement) AddWheelEvent(listener func(event *WheelEvent, target *HTMLElement)) js.Func {
	if !_htmle.IsDefined() {
		ICKWarn("AddWheelEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_WheelEvent(listener)
	_htmle.jsValue.Call("addEventListener", "wheel", callback)
	return callback
}
