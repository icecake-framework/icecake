package browser

import (
	"net/url"
	"strings"

	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/event"
	"github.com/sunraylab/icecake/pkg/js"
)

var windows Window

// GetWindow returning attribute 'window' with
// type Window (idl: Window).
func Win() Window {
	if windows.IsDefined() {
		return windows
	}
	jsv := js.Global().Get("window")
	if typ := jsv.Type(); typ == js.TYPE_NULL || typ == js.TYPE_UNDEFINED {
		console.Panicf("Unable to get browser.window")
	}
	windows.JSValue = jsv
	return windows
}

/******************************************************************************
* Window
******************************************************************************/

// Window
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window
type Window struct {
	js.JSValue                  // embedded js.JSValue
	evth       []*event.Handler // eventhandlers added with an listener to this window event-target
}

// CastWindow is casting a js.JSValue into Window.
func CastWindow(_jsv js.JSValueProvider) Window {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting Window failed")
		return Window{}
	}
	win := new(Window)
	win.JSValue = _jsv.Value()
	return *win
}

/******************************************************************************
* Window's properties
******************************************************************************/

// extract the URL object from the js Location
func (_win Window) URL() *url.URL {
	href := _win.Get("location").GetString("href")
	u, _ := url.Parse(href)
	return u
}

// Navigate to the provided URL.  (aka. js Assign)
//
// After the navigation occurs, the user can navigate back by pressing the "back" button.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Location/assign
func (_win Window) Navigate(url url.URL) {
	_win.Get("location").Call("assign", url.String())
}

// Load and Display the provided URL.
//
// The difference from the Navigate() method is that the current page will not be saved in session History,
// meaning the user won't be able to use the back button to navigate to it.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Location/replace
func (_win Window) Display(url url.URL) {
	_win.Get("location").Call("replace", url.String())
}

// Reloads the current URL, like the Refresh button.
func (_win Window) Reload() {
	_win.Get("location").Call("reload")
}

// Loads a specified resource into a new or existing browsing context (that is, a tab, a window, or an iframe) under a specified name.
//
// The special target keywords, _self, _blank, _parent, and _top, can also be used.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/open
func (_win Window) Open(url *url.URL, target string) Window {
	var win js.JSValue
	if url == nil {
		// a blank page is opened into the targeted browsing context.
		win = _win.Call("open")
	} else {
		win = _win.Call("open", url.String(), target)

	}
	return CastWindow(win)
}

// Closed ndicates whether the referenced window is closed or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/closed
func (_win Window) Closed() bool {
	return _win.GetBool("closed")
}

// History returning attribute 'history' with
// type htmlmisc.History (idl: History).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/history
func (_win Window) History() *History {
	value := _win.Get("history")
	return CastHistory(value)
}

// UserAgent returns the user agent string for the current browser.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/userAgent
func (_win Window) UserAgent() string {
	return _win.Get("navigator").GetString("userAgent")
}

// Language returns a string representing the preferred language of the user, usually the language of the browser UI.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/language
func (_win Window) Language() string {
	return _win.Get("navigator").GetString("language")
}

// OnLine Returns the online status of the browser.
//
// The property returns a boolean value, with true meaning online and false meaning offline.
// The property sends updates whenever the browser's ability to connect to the network changes.
// The update occurs when the user follows links or when a script requests a remote page. *
// For example, the property should return false when users click links soon after they lose internet connection.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/onLine
func (_win Window) OnLine() bool {
	return _win.Get("navigator").GetBool("onLine")
}

// CookieEnabled eturns a Boolean value that indicates whether cookies are enabled or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/cookieEnabled
func (_win Window) CookieEnabled() bool {
	return _win.Get("navigator").GetBool("cookieEnabled")
}

// InnerWidth returns the interior width of the window in pixels. This includes the width of the vertical scroll bar, if one is present.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/innerWidth
func (_win Window) InnerSize() (_width int, _height int) {
	_width = _win.GetInt("innerWidth")
	_height = _win.GetInt("innerHeight")
	return _width, _height
}

// OuterWidth returning attribute 'outerWidth' with
func (_win Window) OuterSize() (_width int, _height int) {
	_width = _win.GetInt("outerWidth")
	_height = _win.GetInt("outerHeight")
	return _width, _height
}

// ScrollX returns the number of pixels that the document is currently scrolled horizontally.
// This value is subpixel precise in modern browsers, meaning that it isn't necessarily a whole number.
// You can get the number of pixels the document is scrolled vertically from the scrollY property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollX
func (_win Window) ScrollPos() (_x float64, _y float64) {
	_x = _win.GetFloat("scrollX")
	_y = _win.GetFloat("scrollY")
	return _x, _y
}

// DevicePixelRatio returns the ratio of the resolution in physical pixels to the resolution in CSS pixels for the current display device.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/devicePixelRatio
func (_win Window) DevicePixelRatio() float64 {
	return _win.GetFloat("devicePixelRatio")
}

// accesses a session Storage object for the current origin.
//
// sessionStorage is similar to localStorage; the difference is that while data in localStorage doesn't expire,
// data in sessionStorage is cleared when the page session ends.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/sessionStorage
func (_win Window) SessionStorage() *Storage {
	// TODO: tryget
	rsp := _win.Call("ickSessionStorage")
	return CastStorage(rsp)
}

// allows you to access a Storage object for the Document's origin; the stored data is saved across browser sessions.
//
// returns nil if access is denied to the localstorage
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/localStorage
func (_win Window) LocalStorage() *Storage {
	// TODO: tryget
	//value := _win.jsValue.Get("localStorage")
	jsv := _win.Call("ickLocalStorage")
	return CastStorage(jsv)
}

// instructs the browser to display a dialog with an optional message, and to wait until the user dismisses the dialog.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/alert
func (_win Window) Alert(message string) {
	message = strings.Trim(message, " ")
	if message == "" {
		_win.Call("alert")
	} else {
		_win.Call("alert", message)
	}
}

// Opens the print dialog to print the current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/print
func (_win Window) Print() {
	_win.Call("print")
}

// method closes the current window, or the window on which it was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/close
func (_win Window) Close() {
	_win.Call("close")
}

// stops further resource loading in the current browsing context, equivalent to the stop button in the browser.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/stop
func (_win Window) Stop() {
	_win.Call("stop")
}

// Makes a request to bring the window to the front.
// It may fail due to user settings and the window isn't guaranteed to be frontmost before this method returns.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/focus
func (_win Window) Focus() {
	_win.Call("focus")
}

// Shifts focus away from the window.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/blur_event
func (_win Window) Blur() {
	_win.Call("blur")
}

/******************************************************************************
* Window's GENERIC_EVENT
******************************************************************************/

func (_win *Window) addListener(_evttyp string, _evjsh js.JSFunction) {
	if _win.evth == nil {
		_win.evth = make([]*event.Handler, 0, 1)
	}
	evh := event.NewEventHandler(_evttyp, _evjsh, func() {
		_win.Call("removeEventListener", _evttyp, _evjsh)
		_evjsh.Release()
	})
	_win.evth = append(_win.evth, evh)
	_win.Call("addEventListener", _evttyp, _evjsh)
}

// Release need to be called only when avent handlers have been added to the Event-target.
// Release removes all event listeners and release ressources allocated fot the associated js func
func (_win *Window) RemoveListeners() {
	for _, evh := range _win.evth {
		evh.Release()
	}
}

// event attribute: Event
func makeWindow_Generic_Event(_listener func(evt *event.Event, target *Window)) js.JSFunction {
	fn := func(this js.JSValue, args []js.JSValue) interface{} {
		value := args[0]
		evt := event.CastEvent(value)
		target := CastWindow(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Stackf(r, "Error occurs processing event %q on Window", evt.Type())
			}
		}()
		_listener(evt, &target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_win Window) AddGenericEvent(evttype event.GENERIC_EVENT, _listener func(*event.Event, *Window)) {
	jsevh := makeWindow_Generic_Event(_listener)
	_win.addListener(string(evttype), jsevh)
}

/******************************************************************************
* Window's BeforeUnloadEvent
******************************************************************************/

// event attribute: BeforeUnloadEvent
func makeWindow_BeforeUnload_Event(_listener func(evt *event.BeforeUnloadEvent, target *Window)) js.JSFunction {
	fn := func(this js.JSValue, args []js.JSValue) interface{} {
		value := args[0]
		evt := event.CastBeforeUnloadEvent(value)
		target := CastWindow(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Stackf(r, "Error occurs processing event %q on Window", evt.Type())
			}
		}()
		_listener(evt, &target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_win Window) AddBeforeUnloadEvent(_listener func(*event.BeforeUnloadEvent, *Window)) {
	jsevh := makeWindow_BeforeUnload_Event(_listener)
	_win.addListener("beforeunload", jsevh)
}

/****************************************************************************
* Window's HASCHANGED_EVENT
*****************************************************************************/

// event attribute: HashChangeEvent
func makeWindow_HashChange_Event(_listener func(*event.HashChangeEvent, *Window)) js.JSFunction {
	fn := func(this js.JSValue, args []js.JSValue) interface{} {
		value := args[0]
		evt := event.CastHashChangeEvent(value)
		target := CastWindow(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Stackf(r, "Error occurs processing event %q on Window", evt.Type())
			}
		}()
		_listener(evt, &target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddHashChange is adding doing AddEventListener for 'HashChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_win Window) AddHashChangeEvent(_listener func(*event.HashChangeEvent, *Window)) {
	jsevh := makeWindow_HashChange_Event(_listener)
	_win.addListener("hashchange", jsevh)
}

/****************************************************************************
* Window's PAGETRANSITION_EVENT
*****************************************************************************/

// event attribute: PageTransitionEvent
func makeWindow_PageTransition_Event(_listener func(event *event.PageTransitionEvent, target *Window)) js.JSFunction {
	fn := func(this js.JSValue, args []js.JSValue) interface{} {
		value := args[0]
		evt := event.CastPageTransitionEvent(value)
		target := CastWindow(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Stackf(r, "Error occurs processing event %q on Window", evt.Type())
			}
		}()
		_listener(evt, &target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_win Window) AddPageTransitionEvent(evttype event.PAGETRANSITION_EVENT, _listener func(*event.PageTransitionEvent, *Window)) {
	jsevh := makeWindow_PageTransition_Event(_listener)
	_win.addListener(string(evttype), jsevh)
}

/****************************************************************************
* Window's UI_EVENT
*****************************************************************************/

// event attribute: UIEvent
func makeWindow_UI_Event(_listener func(*event.UIEvent, *Window)) js.JSFunction {
	fn := func(this js.JSValue, args []js.JSValue) interface{} {
		value := args[0]
		evt := event.CastUIEvent(value)
		target := CastWindow(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Stackf(r, "Error occurs processing event %q on Window", evt.Type())
			}
		}()
		_listener(evt, &target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_win Window) AddResizeEvent(_listener func(*event.UIEvent, *Window)) {
	jsevh := makeWindow_UI_Event(_listener)
	_win.addListener("resize", jsevh)
}
