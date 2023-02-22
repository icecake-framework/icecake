package ick

import "syscall/js"

// represents the state and the identity of the user agent. It allows scripts to query it and to register themselves to carry on some activities.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator
type Navigator struct {
	jsValue js.Value
}

// CastNavigator is casting a js.Value into Navigator.
func CastNavigator(value js.Value) *Navigator {
	if value.Type() != js.TypeObject {
		ConsoleErrorf("casting Navigator failed")
		return nil
	}
	cast := new(Navigator)
	cast.jsValue = value
	return cast
}

// UserAgent returns the user agent string for the current browser.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/userAgent
func (_navigator *Navigator) UserAgent() string {
	return _navigator.jsValue.Get("userAgent").String()
}

// Language returns a string representing the preferred language of the user, usually the language of the browser UI.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/language
func (_navigator *Navigator) Language() string {
	return _navigator.jsValue.Get("language").String()
}

// OnLine Returns the online status of the browser.
//
// The property returns a boolean value, with true meaning online and false meaning offline.
// The property sends updates whenever the browser's ability to connect to the network changes.
// The update occurs when the user follows links or when a script requests a remote page. *
// For example, the property should return false when users click links soon after they lose internet connection.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/onLine
func (_navigator *Navigator) OnLine() bool {
	return _navigator.jsValue.Get("onLine").Bool()
}

// CookieEnabled eturns a Boolean value that indicates whether cookies are enabled or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/cookieEnabled
func (_navigator *Navigator) CookieEnabled() bool {
	return _navigator.jsValue.Get("cookieEnabled").Bool()
}
