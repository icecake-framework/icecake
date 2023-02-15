package dom

import "syscall/js"

/******************************************************************************
* Storage
******************************************************************************/

// https://developer.mozilla.org/fr/docs/Web/API/Web_Storage_API
type Storage struct {
	jsValue js.Value
}

// CastStorage is casting a js.Value into Storage.
func CastStorage(value js.Value) *Storage {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := new(Storage)
	ret.jsValue = value
	return ret
}

// Length returns the number of data items stored in a given Storage object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/length
func (_store *Storage) Count() int {
	return _store.jsValue.Get("length").Int()
}

//	returns the name of the nth key in a given Storage object. The order of keys is user-agent defined, so you should not rely on it.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/key
func (_store *Storage) At(_idx int) (_key string) {
	key := _store.jsValue.Call("key", uint(_idx))
	if key.Type() != js.TypeNull && key.Type() != js.TypeUndefined {
		_key = key.String()
	}
	return _key
}

// when passed a key name, will return that key's value, or null if the key does not exist, in the given Storage object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/getItem
func (_store *Storage) Item(key string) (_item string) {
	item := _store.jsValue.Call("getItem", key)
	if item.Type() != js.TypeNull && item.Type() != js.TypeUndefined {
		_item = item.String()
	}
	return _item
}

// when passed a key name, will return that key's value, or null if the key does not exist, in the given Storage object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/getItem
func (_store *Storage) SetItem(key string, value string) {
	_store.jsValue.Call("setItem", key, value)
}

// when passed a key name, will remove that key from the given Storage object if it exists.
// The Storage interface of the Web Storage API provides access to a particular domain's session or local storage.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/removeItem
func (_store *Storage) RemoveItem(key string) {
	_store.jsValue.Call("removeItem", key)
}

// clears all keys stored in a given Storage object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/clear
func (_store *Storage) Clear() {
	_store.jsValue.Call("clear")
}
