package browser

import "syscall/js"

/**********************************************************************
* PromiseBoolOnFulfilled
 */

type PromiseBoolOnFulfilledFunc func(value bool)

// PromiseBoolOnFulfilled is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseBoolOnFulfilled js.Func

func PromiseBoolOnFulfilledToJS(callback PromiseBoolOnFulfilledFunc) *PromiseBoolOnFulfilled {
	if callback == nil {
		return nil
	}
	ret := PromiseBoolOnFulfilled(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		_p0 := (args[0]).Bool()
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseBoolOnFulfilledFromJS(_value js.Value) PromiseBoolOnFulfilledFunc {
	return func(value bool) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := value
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
	}
}

/**********************************************************************
* PromiseBoolOnRejected
 */

type PromiseBoolOnRejectedFunc func(reason js.Value)

// PromiseBoolOnRejected is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseBoolOnRejected js.Func

func PromiseBoolOnRejectedToJS(callback PromiseBoolOnRejectedFunc) *PromiseBoolOnRejected {
	if callback == nil {
		return nil
	}
	ret := PromiseBoolOnRejected(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		_p0 := args[0]
		callback(_p0)

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseBoolOnRejectedFromJS(_value js.Value) PromiseBoolOnRejectedFunc {
	return func(reason js.Value) {
		var (
			_args [1]interface{}
			_end  int
		)
		_p0 := reason
		_args[0] = _p0
		_end++
		_value.Invoke(_args[0:_end]...)
	}
}

/**********************************************************************
* PromiseFinally
 */

// callback: PromiseFinally
type PromiseFinallyFunc func()

// PromiseFinally is a javascript function type.
//
// Call Release() when done to release resouces
// allocated to this type.
type PromiseFinally js.Func

func PromiseFinallyToJS(callback PromiseFinallyFunc) *PromiseFinally {
	if callback == nil {
		return nil
	}
	ret := PromiseFinally(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var ()
		callback()

		// returning no return value
		return nil
	}))
	return &ret
}

func PromiseFinallyFromJS(_value js.Value) PromiseFinallyFunc {
	return func() {
		var (
			_args [0]interface{}
			_end  int
		)
		_value.Invoke(_args[0:_end]...)
	}
}

/**********************************************************************
* PromiseBool
 */

// class: Promise
type PromiseBool struct {
	Value js.Value
}

// PromiseBoolFromJS is casting a js.Value into PromiseBool.
func PromiseBoolFromJS(value js.Value) *PromiseBool {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromiseBool{}
	ret.Value = value
	return ret
}

func (_this *PromiseBool) Then(onFulfilled *PromiseBoolOnFulfilled, onRejected *PromiseBoolOnRejected) (_result *PromiseBool) {
	var (
		_args [2]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFulfilled != nil {
		__callback0 = (*onFulfilled).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	if onRejected != nil {

		var __callback1 js.Value
		if onRejected != nil {
			__callback1 = (*onRejected).Value
		} else {
			__callback1 = js.Null()
		}
		_p1 := __callback1
		_args[1] = _p1
		_end++
	}
	_returned := _this.Value.Call("then", _args[0:_end]...)
	return PromiseBoolFromJS(_returned)
}

func (_this *PromiseBool) Catch(onRejected *PromiseBoolOnRejected) (_result *PromiseBool) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onRejected != nil {
		__callback0 = (*onRejected).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value.Call("catch", _args[0:_end]...)
	return PromiseBoolFromJS(_returned)
}

func (_this *PromiseBool) Finally(onFinally *PromiseFinally) (_result *PromiseBool) {
	var (
		_args [1]interface{}
		_end  int
	)

	var __callback0 js.Value
	if onFinally != nil {
		__callback0 = (*onFinally).Value
	} else {
		__callback0 = js.Null()
	}
	_p0 := __callback0
	_args[0] = _p0
	_end++
	_returned := _this.Value.Call("finally", _args[0:_end]...)
	return PromiseBoolFromJS(_returned)
}
