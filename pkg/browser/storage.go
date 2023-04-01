package browser

import (
	"fmt"
	"strconv"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/js"
)

/******************************************************************************
* Storage
******************************************************************************/

// https://developer.mozilla.org/fr/docs/Web/API/Web_Storage_API
type Storage struct {
	js.JSValue
}

// CastStorage is casting a js.Value into Storage.
func castStorage(_jsv js.JSValueProvider) *Storage {
	if !_jsv.Value().IsObject() {
		return nil
	}
	ret := new(Storage)
	ret.JSValue = _jsv.Value()
	return ret
}

// accesses a session Storage object for the current origin.
//
// sessionStorage is similar to localStorage; the difference is that while data in localStorage doesn't expire,
// data in sessionStorage is cleared when the page session ends.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/sessionStorage
func SessionStorage() *Storage {
	// TODO: tryget SessionStorage
	rsp := Win().Call("ickSessionStorage")
	return castStorage(rsp)
}

// allows you to access a Storage object for the Document's origin; the stored data is saved across browser sessions.
//
// returns nil if access is denied to the localstorage
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/localStorage
func LocalStorage() *Storage {
	// TODO: tryget LocalStorage
	//value := _win.jsValue.Get("localStorage")
	jsv := Win().Call("ickLocalStorage")
	return castStorage(jsv)
}

// Length returns the number of data items stored in a given Storage object.
//
// returns -1 if the storage object is nil
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/length
func (_store *Storage) Count() int {
	if _store == nil {
		return -1
	}
	return _store.GetInt("length")
}

//	returns the name of the nth key in a given Storage object. The order of keys is user-agent defined, so you should not rely on it.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/key
func (_store *Storage) At(_idx int) string {
	if _store == nil {
		return ""
	}
	key := _store.Call("key", uint(_idx))
	return key.String()
}

// when passed a key name, will return that key's value, or null if the key does not exist, in the given Storage object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/getItem
func (_store *Storage) Get(key string) (_value string) {
	if _store == nil {
		return ""
	}
	val := _store.Call("getItem", key)
	if val.Type() == js.TYPE_STRING {
		_value = val.String()
	}
	return _value
}

// return true if the key exists and the value is not "false" nor "0"
func (_store *Storage) GetBool(key string) (_value bool) {
	if _store == nil {
		return false
	}
	jsval := _store.Call("getItem", key)
	if jsval.Type() == js.TYPE_STRING {
		val := helper.Normalize(jsval.String())
		if val != "false" && val != "0" {
			return true
		}
	}
	return false
}

// convert the returned value in int. returns 0 if the key does not exists
func (_store *Storage) GetInt(key string) (_value int) {
	if _store == nil {
		return 0
	}
	jsval := _store.Call("getItem", key)
	if jsval.IsDefined() {
		i, err := strconv.Atoi(jsval.String())
		if err != nil {
			_value = i
		}
	}
	return _value
}

// when passed a key name, will return that key's value, or null if the key does not exist, in the given Storage object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/getItem
func (_store *Storage) Set(key string, value string) error {
	if _store == nil {
		return nil
	}
	//_store.jsValue.Call("setItem", key, value)

	rsp := js.Global().Call("ickStorageSetItem", _store.Value(), key, value)
	if rsp.Type() == js.TYPE_STRING {
		return fmt.Errorf(rsp.String())
	}
	return nil
}

// when passed a key name, will remove that key from the given Storage object if it exists.
// The Storage interface of the Web Storage API provides access to a particular domain's session or local storage.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/removeItem
func (_store *Storage) Remove(key string) {
	if _store == nil {
		return
	}
	_store.Call("removeItem", key)
}

// clears all keys stored in a given Storage object.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Storage/clear
func (_store *Storage) Clear() {
	if _store == nil {
		return
	}
	_store.Call("clear")
}
