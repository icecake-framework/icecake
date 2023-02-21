package dom

import (
	"net/url"
	"syscall/js"
)

// https://developer.mozilla.org/en-US/docs/Web/API/Location
type Location struct {
	jsValue js.Value
}

// CastLocation is casting a js.Value into Location.
func CastLocation(value js.Value) *Location {
	if value.Type() != js.TypeObject {
		ICKError("casting Location failed")
		return nil
	}
	ret := new(Location)
	ret.jsValue = value
	return ret
}

// extract the URL object from the js Location
func (_loc *Location) URL() *url.URL {
	href := _loc.jsValue.Get("href").String()
	u, _ := url.Parse(href)
	return u
}

// Navigate to the provided URL.  (aka. js Assign)
//
// After the navigation occurs, the user can navigate back by pressing the "back" button.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Location/assign
func (_loc *Location) Navigate(url url.URL) {
	_loc.jsValue.Call("assign", url.String())
}

// Load and Display the provided URL.
//
// The difference from the Navigate() method is that the current page will not be saved in session History,
// meaning the user won't be able to use the back button to navigate to it.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Location/replace
func (_loc *Location) Display(url url.URL) {
	_loc.jsValue.Call("replace", url.String())
}

// Reloads the current URL, like the Refresh button.
func (_loc *Location) Reload() {
	_loc.jsValue.Call("reload")
}
