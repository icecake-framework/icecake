package webclientsdk

import "syscall/js"

/***************************************************************************************
* Storage
 */

// https://developer.mozilla.org/fr/docs/Web/API/Web_Storage_API
type Storage struct {
	jsValue js.Value
}

// StorageFromJS is casting a js.Value into Storage.
func StorageFromJS(value js.Value) *Storage {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Storage{}
	ret.jsValue = value
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *Storage) Length() uint {
	var ret uint
	value := _this.jsValue.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *Storage) Key(index uint) (_result *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("key", _args[0:_end]...)
	var (
		_converted *string // javascript: DOMString _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		__tmp := (_returned).String()
		_converted = &__tmp
	}
	_result = _converted
	return
}

func (_this *Storage) GetItem(key string) (_result *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := key
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("getItem", _args[0:_end]...)
	var (
		_converted *string // javascript: DOMString _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		__tmp := (_returned).String()
		_converted = &__tmp
	}
	_result = _converted
	return
}

func (_this *Storage) SetItem(key string, value string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := key
	_args[0] = _p0
	_end++
	_p1 := value
	_args[1] = _p1
	_end++
	_this.jsValue.Call("setItem", _args[0:_end]...)
}

func (_this *Storage) RemoveItem(key string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := key
	_args[0] = _p0
	_end++
	_this.jsValue.Call("removeItem", _args[0:_end]...)
}

func (_this *Storage) Clear() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("clear", _args[0:_end]...)
}

/**********************************************************************************
* StorageEstimate
 */

type StorageEstimate struct {
	Usage int
	Quota int
}

// StorageEstimateFromJS is allocating a new
// StorageEstimate object and copy all values in the value javascript object.
func StorageEstimateFromJS(value js.Value) *StorageEstimate {
	var out StorageEstimate
	var (
		value0 int // javascript: unsigned long long {usage Usage usage}
		value1 int // javascript: unsigned long long {quota Quota quota}
	)
	value0 = (value.Get("usage")).Int()
	out.Usage = value0
	value1 = (value.Get("quota")).Int()
	out.Quota = value1
	return &out
}

// class: Promise
type PromiseStorageEstimate struct {
	jsValue js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *PromiseStorageEstimate) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.jsValue
}

// PromiseStorageEstimateFromJS is casting a js.Value into PromiseStorageEstimate.
func PromiseStorageEstimateFromJS(value js.Value) *PromiseStorageEstimate {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PromiseStorageEstimate{}
	ret.jsValue = value
	return ret
}

func (_this *PromiseStorageEstimate) Finally(onFinally *PromiseFinally) (_result *PromiseStorageEstimate) {
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
	_returned := _this.jsValue.Call("finally", _args[0:_end]...)
	return PromiseStorageEstimateFromJS(_returned)
}

/*************************************************************************************************
* StorageManager
 */

// https://developer.mozilla.org/en-US/docs/Web/API/StorageManager
type StorageManager struct {
	jsValue js.Value
}

// StorageManagerFromJS is casting a js.Value into StorageManager.
func StorageManagerFromJS(value js.Value) *StorageManager {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &StorageManager{}
	ret.jsValue = value
	return ret
}

func (_this *StorageManager) Persisted() (_result *PromiseBool) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("persisted", _args[0:_end]...)
	return PromiseBoolFromJS(_returned)
}

func (_this *StorageManager) Persist() (_result *PromiseBool) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("persist", _args[0:_end]...)
	return PromiseBoolFromJS(_returned)
}

func (_this *StorageManager) Estimate() (_result *PromiseStorageEstimate) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("estimate", _args[0:_end]...)
	return PromiseStorageEstimateFromJS(_returned)
}
