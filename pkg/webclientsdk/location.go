package browser

import (
	"net/url"
	"syscall/js"
)

// https://developer.mozilla.org/en-US/docs/Web/API/Location
type Location struct {
	jsValue js.Value
}

// NewLocationFromJS is casting a js.Value into Location.
func NewLocationFromJS(value js.Value) *Location {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := new(Location)
	ret.jsValue = value
	return ret
}

// extract the URL object from the js Location
func (_this *Location) URL() (_ret *url.URL) {
	_ret, _ = url.Parse(_this.jsValue.Get("href").String())
	return _ret
}

// Navigate to the provided URL.  (aka. js Assign)
//
// After the navigation occurs, the user can navigate back by pressing the "back" button.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Location/assign
func (_this *Location) Navigate(url url.URL) {
	_this.jsValue.Call("assign", url.String())
}

// Load and Display the provided URL.
//
// The difference from the Navigate() method is that the current page will not be saved in session History,
// meaning the user won't be able to use the back button to navigate to it.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Location/replace
func (_this *Location) Display(url url.URL) {
	_this.jsValue.Call("replace", url.String())
}

// Reloads the current URL, like the Refresh button.
func (_this *Location) Reload() {
	_this.jsValue.Call("reload")
}
