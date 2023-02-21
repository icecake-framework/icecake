package dom

import "syscall/js"

/*
*********************************************************************
// Promise
*/

type PromiseOnFulfilledFunc func(value js.Value)
type PromiseOnFulfilled js.Func

type PromiseOnRejectedFunc func(reason js.Value)
type PromiseOnRejected js.Func

type PromiseFinallyFunc func()
type PromiseFinally js.Func

type Promise struct {
	Value_JS js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *Promise) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.Value_JS
}

// PromiseFromJS is casting a js.Value into Promise.
func CastPromise(value js.Value) *Promise {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Promise{}
	ret.Value_JS = value
	return ret
}

func (_this *Promise) Then(onFulfilled *PromiseOnFulfilled, onRejected *PromiseOnRejected) *Promise {

	var __callback0, __callback1 js.Value

	if onFulfilled != nil {
		__callback0 = (*onFulfilled).Value
	} else {
		__callback0 = js.Null()
	}

	if onRejected != nil {
		__callback1 = (*onRejected).Value
	} else {
		__callback1 = js.Null()
	}

	promise := _this.Value_JS.Call("then", __callback0, __callback1)
	return CastPromise(promise)
}

func (_this *Promise) Catch(onRejected *PromiseOnRejected) *Promise {

	var __callback0 js.Value
	if onRejected != nil {
		__callback0 = (*onRejected).Value
	} else {
		__callback0 = js.Null()
	}
	promise := _this.Value_JS.Call("catch", __callback0)
	return CastPromise(promise)
}

func (_this *Promise) Finally(onFinally *PromiseFinally) *Promise {

	var __callback0 js.Value
	if onFinally != nil {
		__callback0 = (*onFinally).Value
	} else {
		__callback0 = js.Null()
	}
	promise := _this.Value_JS.Call("finally", __callback0)
	return CastPromise(promise)
}
