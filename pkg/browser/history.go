package browser

import (
	"net/url"

	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/js"
)

/******************************************************************************
* History
******************************************************************************/

type HISTORY_SCROLL_RESTORATION string

const (
	HISTORY_SR_AUTO   HISTORY_SCROLL_RESTORATION = "auto"
	HISTORY_SR_MANUAL HISTORY_SCROLL_RESTORATION = "manual"
)

// The DOM Window object provides access to the browser's session history (not to be confused for WebExtensions history)
// through the history object.
// It exposes useful methods and properties that let you navigate back and forth through the user's history,
// and manipulate the contents of the history stack.
//
// https://developer.mozilla.org/en-US/docs/Web/API/History
type History struct {
	js.JSValue
}

// CastHistory is casting a js.Value into History.
func castHistory(_jsvp js.JSValueProvider) *History {
	if _jsvp.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting History failed")
		return nil
	}
	cast := new(History)
	cast.JSValue = _jsvp.Value()
	return cast
}

// History returning attribute 'history' with
// type htmlmisc.History (idl: History).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/history
func SessionHistory() *History {
	value := Win().Get("history")
	return castHistory(value)
}

// Length eturns an integer representing the number of elements in the session history, including the currently loaded page.
//
// https://developer.mozilla.org/en-US/docs/Web/API/History/length
func (_this *History) Count() int {
	return _this.GetInt("length")
}

// ScrollRestoration  allows web applications to explicitly set default scroll restoration behavior on history navigation.
//
// https://developer.mozilla.org/en-US/docs/Web/API/History/scrollRestoration
func (_this *History) ScrollRestoration() HISTORY_SCROLL_RESTORATION {
	return HISTORY_SCROLL_RESTORATION(_this.GetString("scrollRestoration"))
}

// ScrollRestoration  allows web applications to explicitly set default scroll restoration behavior on history navigation.
//
// https://developer.mozilla.org/en-US/docs/Web/API/History/scrollRestoration
func (_this *History) SetScrollRestoration(value HISTORY_SCROLL_RESTORATION) {
	_this.Set("scrollRestoration", value)
}

// State returns a value representing the state at the top of the history stack.
// This is a way to look at the state without having to wait for a popstate event.
//
// The value is null until the pushState() or replaceState() method is used.
//
// TODO: browser - handle History.State returned value
//
// https://developer.mozilla.org/en-US/docs/Web/API/History/state
func (_this *History) State() js.JSValue {
	return _this.Get("state")
}

// Go Loads a specific page from the session history.
// You can use it to move forwards and backwards through the history depending on the value of a parameter.
//
// This method is asynchronous. Add a listener for the popstate event in order to determine when the navigation has completed.
//
// https://developer.mozilla.org/en-US/docs/Web/API/History/go
func (_this *History) Go(delta int) {
	_this.Call("go", delta)
}

// causes the browser to move back one page in the session history.
//
// https://developer.mozilla.org/en-US/docs/Web/API/History/back
func (_this *History) Back() {
	_this.Call("back")
}

// causes the browser to move forward one page in the session history. It has the same effect as calling history.go(1).
//
// https://developer.mozilla.org/en-US/docs/Web/API/History/forward
func (_this *History) Forward() {
	_this.Call("forward")
}

// PushState adds an entry to the browser's session history stack.
//
// https://developer.mozilla.org/en-US/docs/Web/API/History/pushState
func (_this *History) PushState(data interface{}, url *url.URL) {
	if url == nil {
		_this.Call("pushState", data)
	} else {
		_this.Call("pushState", data, url.String())
	}
}

// modifies the current history entry, replacing it with the state object and URL passed in the method parameters.
// This method is particularly useful when you want to update the state object or URL
// of the current history entry in response to some user action.
//
// https://developer.mozilla.org/en-US/docs/Web/API/History/replaceState
func (_this *History) ReplaceState(data interface{}, url *url.URL) {
	if url == nil {
		_this.Call("replaceState", data)
	} else {
		_this.Call("replaceState", data, url.String())
	}
}
