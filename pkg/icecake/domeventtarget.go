package ick

import (
	"syscall/js"
)

type eventHandler struct {
	eventtype string // 'onclick'...
	jsHandler js.Func
	close     func()
}

/******************************************************************************
* EventTarget
*******************************************************************************/

// EventTarget is the root of many objetcs: nodes, window...
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
type EventTarget struct {
	jsValue       js.Value // update with Wrap(), get with JSValue()
	eventHandlers []*eventHandler
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_evttget EventTarget) JSValue() js.Value {
	return _evttget.jsValue
}

func (_evttget *EventTarget) Wrap(_jsval js.Value) {
	if _evttget == nil {
		ConsoleErrorf("unable to wrap a nil element, abort")
		return
	}
	if _evttget.jsValue.Truthy() {
		ConsoleWarnf("wrapping an already wrapped element")
		_evttget.RemoveListeners()
	}
	_evttget.jsValue = _jsval
}

// CastEventTarget is casting a js.Value into EventTarget.
func CastEventTarget(value js.Value) *EventTarget {
	if value.Type() != js.TypeObject {
		ConsoleErrorf("casting EventTarget failed")
		return nil
	}
	evttget := new(EventTarget)
	evttget.jsValue = value
	return evttget
}

// NewEventTarget create a new EventTarget
// func NewEventTarget() (_result *EventTarget) {
// 	_klass := js.Global().Get("EventTarget")
// 	var _args [0]interface{}
// 	_returned := _klass.New(_args[0:0]...)
// 	return NewEventTargetFromJS(_returned)
// }

/******************************************************************************
* EventTarget's methods
*******************************************************************************/

// AddEventListener sets up a function that will be called whenever the specified event is delivered to the target.
//
// Common targets are Element, or its children, Document, and Window, but the target may be any object that supports events (such as XMLHttpRequest).
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (_evttget *EventTarget) AddEventListener(evh *eventHandler) {
	if _evttget.eventHandlers == nil {
		_evttget.eventHandlers = make([]*eventHandler, 0, 1)
	}
	evh.close = func() {
		_evttget.jsValue.Call("removeEventListener", evh.eventtype, evh.jsHandler)
		evh.jsHandler.Release()
	}
	_evttget.eventHandlers = append(_evttget.eventHandlers, evh)
	_evttget.jsValue.Call("addEventListener", evh.eventtype, evh.jsHandler)
}

// Release need to be called only when avent handlers have been added to the Event-target.
// Release removes all event listeners and release ressources allocated fot the associated js func
func (_evttget *EventTarget) RemoveListeners() {
	if len(_evttget.eventHandlers) > 0 {
		for _, evh := range _evttget.eventHandlers {
			evh.close()
		}
	}
}

// RemoveEventListener removes an event listener previously registered with EventTarget.addEventListener() from the target.
// The event listener to be removed is identified using a combination of the event type, the event listener function itself,
// and various optional options that may affect the matching process;
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
// func (_this *EventTarget) RemoveEventListener(_type string, callback *EventListener) {
// 	_this.jsValue.Call("removeEventListener", _type, callback.JSValue())
// }
