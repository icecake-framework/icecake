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
		ConsoleError("casting HTMLElement failed")
		return nil
	}
	ret := new(HTMLElement)
	ret.jsValue = value
	return ret
}

/****************************************************************************
* HTMLElement's properties & methods
*****************************************************************************/

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (_elem *HTMLElement) AccessKey() string {
	return _elem.jsValue.Get("accessKey").String()
}

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (_elem *HTMLElement) SetAccessKey(key bool) *HTMLElement {
	_elem.jsValue.Set("accessKey", key)
	return _elem
}

// InnerText represents the rendered text content of a node and its descendants.
//
// InnerText gets pure text, removing any html or css, while TextContent keeps the representation.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (_this *HTMLElement) InnerText() string {
	if _this == nil {
		return ""
	}
	var ret string
	value := _this.jsValue.Get("innerText")
	ret = (value).String()
	return ret
}

// InnerText represents the rendered text content of a node and its descendants.
//
// InnerText gets pure text, removing any html or css, while TextContent keeps the representation.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (_this *HTMLElement) SetInnerText(value string) {
	if _this == nil {
		ConsoleWarn("SetInnerText fail: nil HTMLelement")
		return
	}
	input := value
	_this.jsValue.Set("innerText", input)
}

// Focus sets focus on the specified element, if it can be focused. The focused element is the element that will receive keyboard and similar events by default.
//
// By default the browser will scroll the element into view after focusing it,
// and it may also provide visible indication of the focused element (typically by displaying a "focus ring" around the element).
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus
func (_this *HTMLElement) Focus() {
	if _this == nil {
		ConsoleWarn("Focus fail: nil HTMLelement")
		return
	}
	_this.jsValue.Call("focus")
}

// Blur removes keyboard focus from the current element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/blur
func (_this *HTMLElement) Blur() {
	if _this == nil {
		ConsoleWarn("Blur fail: nil HTMLelement")
		return
	}
	_this.jsValue.Call("blur")
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
func (_this *HTMLElement) AddGenericEvent(evttype GENERIC_EVENT, listener func(event *Event, target *HTMLElement)) js.Func {
	if _this == nil {
		ConsoleWarn("AddGenericEvent fail: nil HTMLelement")
		return js.Func{}
	}
	callback := makeHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
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
func (_this *HTMLElement) AddMouseEvent(evttype MOUSE_EVENT, listener func(event *MouseEvent, target *HTMLElement)) js.Func {
	if _this == nil {
		ConsoleWarn("AddMouseEvent fail: nil HTMLelement")
		return js.Func{}
	}
	callback := makeHTMLElement_Mouse_Event(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
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
func (_this *HTMLElement) AddFocusEvent(evttype FOCUS_EVENT, listener func(event *FocusEvent, target *HTMLElement)) js.Func {
	if _this == nil {
		ConsoleWarn("AddFocusEvent fail: nil HTMLelement")
		return js.Func{}
	}
	callback := makeHTMLElement_FocusEvent(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
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
func (_this *HTMLElement) AddPointerEvent(evttype POINTER_EVENT, listener func(event *PointerEvent, target *HTMLElement)) js.Func {
	callback := makeHTMLElement_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
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
func (_this *HTMLElement) AddInputEvent(evttype INPUT_EVENT, listener func(event *InputEvent, target *HTMLElement)) js.Func {
	if _this == nil {
		ConsoleWarn("AddInputEvent fail: nil HTMLelement")
		return js.Func{}
	}
	callback := makeHTMLElement_InputEvent(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
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
func (_this *HTMLElement) AddKeyboard(evttype KEYBOARD_EVENT, listener func(event *KeyboardEvent, target *HTMLElement)) js.Func {
	if _this == nil {
		ConsoleWarn("AddKeyboard fail: nil HTMLelement")
		return js.Func{}
	}
	callback := makeHTMLElement_KeyboardEvent(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
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
func (_this *HTMLElement) AddResizeEvent(listener func(event *UIEvent, target *HTMLElement)) js.Func {
	if _this == nil {
		ConsoleWarn("AddResizeEvent fail: nil HTMLelement")
		return js.Func{}
	}
	callback := makeHTMLElement_UIEvent(listener)
	_this.jsValue.Call("addEventListener", "resize", callback)
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
func (_this *HTMLElement) AddWheelEvent(listener func(event *WheelEvent, target *HTMLElement)) js.Func {
	if _this == nil {
		ConsoleWarn("AddWheelEvent fail: nil HTMLelement")
		return js.Func{}
	}
	callback := makeHTMLElement_WheelEvent(listener)
	_this.jsValue.Call("addEventListener", "wheel", callback)
	return callback
}
