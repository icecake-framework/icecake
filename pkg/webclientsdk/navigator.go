package browser

import "syscall/js"

// https://developer.mozilla.org/en-US/docs/Web/API/Navigator
type Navigator struct {
	jsValue js.Value
}

// NavigatorFromJS is casting a js.Value into Navigator.
func NavigatorFromJS(value js.Value) *Navigator {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Navigator{}
	ret.jsValue = value
	return ret
}

// UserAgent returning attribute 'userAgent' with
// type string (idl: DOMString).
func (_this *Navigator) UserAgent() string {
	var ret string
	value := _this.jsValue.Get("userAgent")
	ret = (value).String()
	return ret
}

// Language returning attribute 'language' with
// type string (idl: DOMString).
func (_this *Navigator) Language() string {
	var ret string
	value := _this.jsValue.Get("language")
	ret = (value).String()
	return ret
}

// OnLine returning attribute 'onLine' with
// type bool (idl: boolean).
func (_this *Navigator) OnLine() bool {
	var ret bool
	value := _this.jsValue.Get("onLine")
	ret = (value).Bool()
	return ret
}

// CookieEnabled returning attribute 'cookieEnabled' with
// type bool (idl: boolean).
func (_this *Navigator) CookieEnabled() bool {
	var ret bool
	value := _this.jsValue.Get("cookieEnabled")
	ret = (value).Bool()
	return ret
}

// Storage returning attribute 'storage' with
// type storage.StorageManager (idl: StorageManager).
func (_this *Navigator) Storage() *StorageManager {
	var ret *StorageManager
	value := _this.jsValue.Get("storage")
	ret = StorageManagerFromJS(value)
	return ret
}
