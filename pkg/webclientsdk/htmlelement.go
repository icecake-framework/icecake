package browser

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

// NewHTMLElementFromJS is casting a js.Value into HTMLElement.
func NewHTMLElementFromJS(value js.Value) *HTMLElement {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &HTMLElement{}
	ret.jsValue = value
	return ret
}

/****************************************************************************
* HTMLElement's properties & methods
*****************************************************************************/

// InnerText represents the rendered text content of a node and its descendants.
//
// InnerText gets pure text, removing any html or css, while TextContent keeps the representation.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/innerText
func (_this *HTMLElement) InnerText() string {
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
	input := value
	_this.jsValue.Set("innerText", input)
}

// Click simulates a mouse click on an element.
//
// When click() is used with supported elements (such as an <input>), it fires the element's click event.
// This event then bubbles up to elements higher in the document tree (or event chain) and fires their click events.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/click
func (_this *HTMLElement) Click() *HTMLElement {
	_this.jsValue.Call("click")
	return _this
}

// Focus sets focus on the specified element, if it can be focused. The focused element is the element that will receive keyboard and similar events by default.
//
// By default the browser will scroll the element into view after focusing it,
// and it may also provide visible indication of the focused element (typically by displaying a "focus ring" around the element).
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/focus
func (_this *HTMLElement) Focus() *HTMLElement {
	_this.jsValue.Call("focus")
	return _this
}

// Blur removes keyboard focus from the current element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/blur
func (_this *HTMLElement) Blur() *HTMLElement {
	_this.jsValue.Call("blur")
	return _this
}

/****************************************************************************
* HTMLElement's GENERIC_EVENT
*****************************************************************************/

type ListenerHTMLE_Generic func(event *Event, target *HTMLElement)

// event attribute: Event
func makeHTMLElement_domcore_Event(listener ListenerHTMLE_Generic) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewEventFromJS(value)
		target := NewHTMLElementFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddAbort is adding doing AddEventListener for 'Abort' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddGenericEvent(evttype GENERIC_EVENT, listener ListenerHTMLE_Generic) js.Func {
	callback := makeHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", evttype, callback)
	return callback
}

/****************************************************************************
* HTMLElement's MOUSE_EVENT
*****************************************************************************/

type ListenerHTMLE_Mouse func(event *MouseEvent, target *HTMLElement)

// event attribute: MouseEvent
func makeHTMLElement_Mouse_Event(listener ListenerHTMLE_Mouse) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewMouseEventFromJS(value)
		target := NewHTMLElementFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddClick is adding doing AddEventListener for 'Click' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddMouseEvent(evttype MOUSE_EVENT, listener ListenerHTMLE_Mouse) js.Func {
	callback := makeHTMLElement_Mouse_Event(listener)
	_this.jsValue.Call("addEventListener", evttype, callback)
	return callback
}

/****************************************************************************
* HTMLElement's FOCUS_EVENT
*****************************************************************************/

type ListenerHTMLE_Focus func(event *FocusEvent, target *HTMLElement)

// event attribute: FocusEvent
func makeHTMLElement_FocusEvent(listener ListenerHTMLE_Focus) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewFocusEventFromJS(value)
		target := NewHTMLElementFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddBlur is adding doing AddEventListener for 'Blur' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddFocusEvent(evttype FOCUS_EVENT, listener ListenerHTMLE_Focus) js.Func {
	callback := makeHTMLElement_FocusEvent(listener)
	_this.jsValue.Call("addEventListener", evttype, callback)
	return callback
}

/****************************************************************************
* HTMLElement's POINTER_EVENT
*****************************************************************************/

type ListenerHTMLE_Pointer func(event *PointerEvent, target *HTMLElement)

// event attribute: PointerEvent
func makeHTMLElement_PointerEvent(listener ListenerHTMLE_Pointer) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var evt *PointerEvent
		value := args[0]
		evt = NewPointerEventFromJS(value)
		target := NewHTMLElementFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddGotPointerCapture is adding doing AddEventListener for 'GotPointerCapture' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddPointerEvent(evttype POINTER_EVENT, listener ListenerHTMLE_Pointer) js.Func {
	callback := makeHTMLElement_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", evttype, callback)
	return callback
}

/****************************************************************************
* HTMLElement's INPUT_EVENT
*****************************************************************************/

type ListenerHTMLE_Input func(event *InputEvent, target *HTMLElement)

// event attribute: InputEvent
func makeHTMLElement_InputEvent(listener ListenerHTMLE_Input) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewInputEventFromJS(value)
		target := NewHTMLElementFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddInput is adding doing AddEventListener for 'Input' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddInputEvent(evttype INPUT_EVENT, listener ListenerHTMLE_Input) js.Func {
	callback := makeHTMLElement_InputEvent(listener)
	_this.jsValue.Call("addEventListener", evttype, callback)
	return callback
}

/****************************************************************************
* HTMLElement's KEYBOARD_EVENT
*****************************************************************************/

type ListenerHTMLE_Keyboard func(event *KeyboardEvent, target *HTMLElement)

// event attribute: KeyboardEvent
func makeHTMLElement_KeyboardEvent(listener ListenerHTMLE_Keyboard) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewKeyboardEventFromJS(value)
		target := NewHTMLElementFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddKeyDown is adding doing AddEventListener for 'KeyDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddKeyboard(evttype KEYBOARD_EVENT, listener ListenerHTMLE_Keyboard) js.Func {
	callback := makeHTMLElement_KeyboardEvent(listener)
	_this.jsValue.Call("addEventListener", evttype, callback)
	return callback
}

/****************************************************************************
* HTMLElement's UI_EVENT
*****************************************************************************/

type ListenerHTMLE_UI func(event *UIEvent, target *HTMLElement)

// event attribute: UIEvent
func makeHTMLElement_UIEvent(listener ListenerHTMLE_UI) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewUIEventFromJS(value)
		target := NewHTMLElementFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddResize is adding doing AddEventListener for 'Resize' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddResizeEvent(listener ListenerHTMLE_UI) js.Func {
	callback := makeHTMLElement_UIEvent(listener)
	_this.jsValue.Call("addEventListener", "resize", callback)
	return callback
}

/****************************************************************************
* HTMLElement's WHEEL_EVENT
*****************************************************************************/

type ListenerHTMLE_Wheel func(event *WheelEvent, target *HTMLElement)

// event attribute: WheelEvent
func makeHTMLElement_WheelEvent(listener ListenerHTMLE_Wheel) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewWheelEventFromJS(value)
		target := NewHTMLElementFromJS(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// The wheel event fires when the user rotates a wheel button on a pointing device (typically a mouse).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/wheel_event
func (_this *HTMLElement) AddWheelEvent(listener ListenerHTMLE_Wheel) js.Func {
	callback := makeHTMLElement_WheelEvent(listener)
	_this.jsValue.Call("addEventListener", "wheel", callback)
	return callback
}
