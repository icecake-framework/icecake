package webclientsdk

import "syscall/js"

/******************************************************************************************
* Event
 */

type Event struct {
	jsValue js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *Event) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.jsValue
}

// EventFromJS is casting a js.Value into Event.
func EventFromJS(value js.Value) *Event {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Event{}
	ret.jsValue = value
	return ret
}

// Type returning attribute 'type' with
// type string (idl: DOMString).
//
// commonly used to refer to the specific event, such as click, load, or error
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/type
func (_this *Event) Type() string {
	var ret string
	value := _this.jsValue.Get("type")
	ret = (value).String()
	return ret
}

// Target returning attribute 'target' with
// type EventTarget (idl: EventTarget).
//
// a reference to the object onto which the event was dispatched.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/target
func (_this *Event) Target() *EventTarget {
	var ret *EventTarget
	value := _this.jsValue.Get("target")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = EventTargetFromJS(value)
	}
	return ret
}

// CurrentTarget returning attribute 'currentTarget' with
// type EventTarget (idl: EventTarget).
//
// It always refers to the element to which the event handler has been attached, as opposed to Event.target, which identifies the element on which the event occurred and which may be its descendant.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/currentTarget
func (_this *Event) CurrentTarget() *EventTarget {
	var ret *EventTarget
	value := _this.jsValue.Get("currentTarget")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = EventTargetFromJS(value)
	}
	return ret
}

/******************************************************************************************
* Event Handler
 */

// callback: EventHandlerNonNull
type EventHandlerFunc func(event *Event) interface{}

// EventHandler is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type EventHandler js.Func

func EventHandlerToJS(callback EventHandlerFunc) *EventHandler {
	if callback == nil {
		return nil
	}
	ret := EventHandler(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Event // javascript: Event event
		)
		_p0 = EventFromJS(args[0])
		_returned := callback(_p0)
		_converted := _returned
		return _converted
	}))
	return &ret
}

func EventHandlerFromJS(_value js.Value) EventHandlerFunc {
	return func(event *Event) (_result interface{}) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := event.jsValue
		_args[0] = _p0
		_end++
		_returned := _value.Invoke(_args[0:_end]...)
		var (
			_converted js.Value // javascript: any
		)
		_converted = _returned
		_result = _converted
		return
	}
}

/******************************************************************************************
* EventListener
 */

// EventListener is a callback interface.
type EventListener interface {
	HandleEvent(event *Event)
}

// EventListenerValue is javascript reference value for callback interface EventListener.
// This is holding the underlying javascript object.
type EventListenerValue struct {
	// Value is the underlying javascript object or function.
	jsValue js.Value

	// Functions is the underlying function objects that is allocated for the interface callback
	Functions [1]js.Func

	// Go interface to invoke
	impl      EventListener
	function  func(event *Event)
	useInvoke bool
}

// JSValue is returning the javascript object that implements this callback interface
func (t *EventListenerValue) JSValue() js.Value {
	return t.JSValue()
}

// Release is releasing all resources that is allocated.
func (t *EventListenerValue) Release() {
	for i := range t.Functions {
		if t.Functions[i].Type() != js.TypeUndefined {
			t.Functions[i].Release()
		}
	}
}

// NewEventListener is allocating a new javascript object that
// implements EventListener.
func NewEventListener(callback EventListener) *EventListenerValue {
	ret := &EventListenerValue{impl: callback}
	ret.jsValue = js.Global().Get("Object").New()
	ret.Functions[0] = ret.allocateHandleEvent()
	ret.jsValue.Set("handleEvent", ret.Functions[0])
	return ret
}

// NewEventListenerFunc is allocating a new javascript
// function is implements
// EventListener interface.
func NewEventListenerFunc(f func(event *Event)) *EventListenerValue {
	// single function will result in javascript function type, not an object
	ret := &EventListenerValue{function: f}
	ret.Functions[0] = ret.allocateHandleEvent()
	ret.jsValue = ret.Functions[0].Value
	return ret
}

// EventListenerFromJS is taking an javascript object that reference to a
// callback interface and return a corresponding interface that can be used
// to invoke on that element.
func EventListenerFromJS(value js.Value) *EventListenerValue {
	if value.Type() == js.TypeObject {
		return &EventListenerValue{jsValue: value}
	}
	if value.Type() == js.TypeFunction {
		return &EventListenerValue{jsValue: value, useInvoke: true}
	}
	panic("unsupported type")
}

func (t *EventListenerValue) allocateHandleEvent() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var (
			_p0 *Event // javascript: Event event
		)
		_p0 = EventFromJS(args[0])
		if t.function != nil {
			t.function(_p0)
		} else {
			t.impl.HandleEvent(_p0)
		}

		// returning no return value
		return nil
	})
}

func (_this *EventListenerValue) HandleEvent(event *Event) {
	if _this.function != nil {
		_this.function(event)
	}
	if _this.impl != nil {
		_this.impl.HandleEvent(event)
	}
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := event.jsValue
	_args[0] = _p0
	_end++
	if _this.useInvoke {

		// invoke a javascript function
		_this.jsValue.Invoke(_args[0:_end]...)
	} else {
		_this.jsValue.Call("handleEvent", _args[0:_end]...)
	}
	return
}

/******************************************************************************************
* EventTarget
 */

// EventTarget
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

// EventTargetFromJS is casting a js.Value into EventTarget.
func EventTargetFromJS(value js.Value) *EventTarget {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &EventTarget{}
	ret.jsValue = value
	return ret
}

func NewEventTarget() (_result *EventTarget) {
	_klass := js.Global().Get("EventTarget")
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _klass.New(_args[0:_end]...)
	var (
		_converted *EventTarget // javascript: EventTarget _what_return_name
	)
	_converted = EventTargetFromJS(_returned)
	_result = _converted
	return
}

func (_this *EventTarget) AddEventListener(_type string, callback *EventListenerValue) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_p1 := callback.JSValue()
	_args[1] = _p1
	_end++
	_this.jsValue.Call("addEventListener", _args[0:_end]...)
	return
}

func (_this *EventTarget) RemoveEventListener(_type string, callback *EventListenerValue) {
	var (
		_args [3]interface{}
		_end  int
	)
	_p0 := _type
	_args[0] = _p0
	_end++
	_p1 := callback.JSValue()
	_args[1] = _p1
	_end++
	_this.jsValue.Call("removeEventListener", _args[0:_end]...)
	return
}
