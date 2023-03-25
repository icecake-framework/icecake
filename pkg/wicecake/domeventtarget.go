package wick

import (
	"syscall/js"

	"github.com/sunraylab/icecake/pkg/console"
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
	JSValue                       // embedded js.Value
	eventHandlers []*eventHandler // eventhandlers added with an listener to this eventtarget
}

// CastEventTarget is casting a js.Value into EventTarget.
func CastEventTarget(_jsv JSValue) *EventTarget {
	if _jsv.Type() != TYPE_OBJECT {
		console.Errorf("casting EventTarget failed")
		return nil
	}
	evttget := new(EventTarget)
	evttget.jsvalue = _jsv.jsvalue
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

// AddListener sets up a function that will be called whenever the specified event is delivered to the target.
//
// Common targets are Element, or its children, Document, and Window, but the target may be any object that supports events (such as XMLHttpRequest).
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (_evttget *EventTarget) AddListener(evh *eventHandler) {
	if _evttget.eventHandlers == nil {
		_evttget.eventHandlers = make([]*eventHandler, 0, 1)
	}
	evh.close = func() {
		_evttget.Call("removeEventListener", evh.eventtype, evh.jsHandler)
		evh.jsHandler.Release()
	}
	_evttget.eventHandlers = append(_evttget.eventHandlers, evh)
	_evttget.Call("addEventListener", evh.eventtype, evh.jsHandler)
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
