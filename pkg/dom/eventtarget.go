package dom

import "syscall/js"

/******************************************************************************
* EventTarget
*******************************************************************************/

// EventTarget is the root of many objetcs: nodes, window...
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type EventTarget struct {
	jsValue js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *EventTarget) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.jsValue
}

func (_this *EventTarget) Wrap(_jsval js.Value) {
	if _this == nil {
		return
	}
	_this.jsValue = _jsval
}

// CastEventTarget is casting a js.Value into EventTarget.
func CastEventTarget(value js.Value) *EventTarget {
	if value.Type() != js.TypeObject {
		ICKError("casting EventTarget failed")
		return nil
	}
	ret := new(EventTarget)
	ret.jsValue = value
	return ret
}

// NewEventTarget create a new EventTarget
// func NewEventTarget() (_result *EventTarget) {
// 	_klass := js.Global().Get("EventTarget")
// 	var _args [0]interface{}
// 	_returned := _klass.New(_args[0:0]...)
// 	return NewEventTargetFromJS(_returned)
// }

/******************************************************************************
* EventTarget's events
*******************************************************************************/

// AddEventListener sets up a function that will be called whenever the specified event is delivered to the target.
//
// Common targets are Element, or its children, Document, and Window, but the target may be any object that supports events (such as XMLHttpRequest).
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
// func (_this *EventTarget) AddEventListener(_type string, callback *EventListener) {
// 	_this.jsValue.Call("addEventListener", _type, callback.JSValue())
// }

// RemoveEventListener removes an event listener previously registered with EventTarget.addEventListener() from the target.
// The event listener to be removed is identified using a combination of the event type, the event listener function itself,
// and various optional options that may affect the matching process;
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
// func (_this *EventTarget) RemoveEventListener(_type string, callback *EventListener) {
// 	_this.jsValue.Call("removeEventListener", _type, callback.JSValue())
// }
