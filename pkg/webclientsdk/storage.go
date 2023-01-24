package browser

import "syscall/js"

/******************************************************************************
* Storage
******************************************************************************/

// https://developer.mozilla.org/fr/docs/Web/API/Web_Storage_API
type Storage struct {
	jsValue js.Value
}

// NewStorageFromJS is casting a js.Value into Storage.
func NewStorageFromJS(value js.Value) *Storage {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := new(Storage)
	ret.jsValue = value
	return ret
}

// Length eturns the number of data items stored in a given Storage object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/length
func (_this *Storage) Length() uint {
	return uint(_this.jsValue.Get("length").Int())
}

//	returns the name of the nth key in a given Storage object. The order of keys is user-agent defined, so you should not rely on it.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/key
func (_this *Storage) Index(index uint) (_result string) {
	_returned := _this.jsValue.Call("key", index)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_result = _returned.String()
	}
	return _result
}

// when passed a key name, will return that key's value, or null if the key does not exist, in the given Storage object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/getItem
func (_this *Storage) Item(key string) (_result string) {
	_returned := _this.jsValue.Call("getItem", key)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_result = _returned.String()
	}
	return _result
}

// when passed a key name, will return that key's value, or null if the key does not exist, in the given Storage object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/getItem
func (_this *Storage) SetItem(key string, value string) {
	_this.jsValue.Call("setItem", key, value)
}

// when passed a key name, will remove that key from the given Storage object if it exists.
// The Storage interface of the Web Storage API provides access to a particular domain's session or local storage.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/removeItem
func (_this *Storage) RemoveItem(key string) {
	_this.jsValue.Call("removeItem", key)
}

// clears all keys stored in a given Storage object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/clear
func (_this *Storage) Clear() {
	_this.jsValue.Call("clear")
}
