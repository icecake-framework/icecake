package dom

import "syscall/js"

// represents the state and the identity of the user agent. It allows scripts to query it and to register themselves to carry on some activities.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator
type Navigator struct {
	jsValue js.Value
}

// NewNavigatorFromJS is casting a js.Value into Navigator.
func NewNavigatorFromJS(value js.Value) *Navigator {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := new(Navigator)
	ret.jsValue = value
	return ret
}

// UserAgent returns the user agent string for the current browser.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/userAgent
func (_this *Navigator) UserAgent() string {
	var ret string
	value := _this.jsValue.Get("userAgent")
	ret = (value).String()
	return ret
}

// Language returns a string representing the preferred language of the user, usually the language of the browser UI.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/language
func (_this *Navigator) Language() string {
	var ret string
	value := _this.jsValue.Get("language")
	ret = (value).String()
	return ret
}

// OnLine Returns the online status of the browser.
//
// The property returns a boolean value, with true meaning online and false meaning offline.
// The property sends updates whenever the browser's ability to connect to the network changes.
// The update occurs when the user follows links or when a script requests a remote page. *
// For example, the property should return false when users click links soon after they lose internet connection.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/onLine
func (_this *Navigator) OnLine() bool {
	var ret bool
	value := _this.jsValue.Get("onLine")
	ret = (value).Bool()
	return ret
}

// CookieEnabled eturns a Boolean value that indicates whether cookies are enabled or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/cookieEnabled
func (_this *Navigator) CookieEnabled() bool {
	return _this.jsValue.Get("cookieEnabled").Bool()
}
