package ick

import (
	"net/url"
	"strings"
	"syscall/js"
)

/******************************************************************************
* Window
******************************************************************************/

// Window
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window
type Window struct {
	EventTarget
}

// CastWindow is casting a js.Value into Window.
func CastWindow(value js.Value) Window {
	if value.Type() != js.TypeObject {
		ConsoleErrorf("casting Window failed")
		return Window{}
	}
	win := new(Window)
	win.jsValue = value
	return *win
}

// GetWindow returning attribute 'window' with
// type Window (idl: Window).
func getWindow() Window {
	value := js.Global().Get("window")
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		ConsoleErrorf("Unable to get window")
		panic("Unable to get window")
	}
	win := new(Window)
	win.jsValue = value
	return *win
}

/******************************************************************************
* Window's properties
******************************************************************************/

// extract the URL object from the js Location
func (_win Window) URL() *url.URL {
	href := _win.jsValue.Get("Location").Get("href").String()
	u, _ := url.Parse(href)
	return u
}

// Navigate to the provided URL.  (aka. js Assign)
//
// After the navigation occurs, the user can navigate back by pressing the "back" button.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Location/assign
func (_win Window) Navigate(url url.URL) {
	_win.jsValue.Get("Location").Call("assign", url.String())
}

// Load and Display the provided URL.
//
// The difference from the Navigate() method is that the current page will not be saved in session History,
// meaning the user won't be able to use the back button to navigate to it.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Location/replace
func (_win Window) Display(url url.URL) {
	_win.jsValue.Get("Location").Call("replace", url.String())
}

// Reloads the current URL, like the Refresh button.
func (_win Window) Reload() {
	_win.jsValue.Get("Location").Call("reload")
}

// Loads a specified resource into a new or existing browsing context (that is, a tab, a window, or an iframe) under a specified name.
//
// The special target keywords, _self, _blank, _parent, and _top, can also be used.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/open
func (_win Window) Open(url *url.URL, target string) Window {
	var win js.Value
	if url == nil {
		// a blank page is opened into the targeted browsing context.
		win = _win.jsValue.Call("open")
	} else {
		win = _win.jsValue.Call("open", url.String(), target)

	}
	return CastWindow(win)
}

// Closed ndicates whether the referenced window is closed or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/closed
func (_win Window) Closed() bool {
	return _win.jsValue.Get("closed").Bool()
}

// History returning attribute 'history' with
// type htmlmisc.History (idl: History).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/history
func (_win Window) History() *History {
	value := _win.jsValue.Get("history")
	return CastHistory(value)
}

// Top returns a reference to the topmost window in the window hierarchy.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/top
func (_win Window) Top() Window {
	value := _win.jsValue.Get("top")
	return CastWindow(value)
}

// UserAgent returns the user agent string for the current browser.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/userAgent
func (_win Window) UserAgent() string {
	return _win.jsValue.Get("navigator").Get("userAgent").String()
}

// Language returns a string representing the preferred language of the user, usually the language of the browser UI.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/language
func (_win Window) Language() string {
	return _win.jsValue.Get("navigator").Get("language").String()
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
	return _win.jsValue.Get("navigator").Get("onLine").Bool()
}

// CookieEnabled eturns a Boolean value that indicates whether cookies are enabled or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator/cookieEnabled
func (_win Window) CookieEnabled() bool {
	return _win.jsValue.Get("navigator").Get("cookieEnabled").Bool()
}

// InnerWidth returns the interior width of the window in pixels. This includes the width of the vertical scroll bar, if one is present.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/innerWidth
func (_win Window) InnerSize() (_width int, _height int) {
	_width = _win.jsValue.Get("innerWidth").Int()
	_height = _win.jsValue.Get("innerHeight").Int()
	return _width, _height
}

// OuterWidth returning attribute 'outerWidth' with
func (_win Window) OuterSize() (_width int, _height int) {
	_width = _win.jsValue.Get("outerWidth").Int()
	_height = _win.jsValue.Get("outerHeight").Int()
	return _width, _height
}

// ScrollX returns the number of pixels that the document is currently scrolled horizontally.
// This value is subpixel precise in modern browsers, meaning that it isn't necessarily a whole number.
// You can get the number of pixels the document is scrolled vertically from the scrollY property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollX
func (_win Window) ScrollPos() (_x float64, _y float64) {
	_x = _win.jsValue.Get("scrollX").Float()
	_y = _win.jsValue.Get("scrollY").Float()
	return _x, _y
}

// DevicePixelRatio returns the ratio of the resolution in physical pixels to the resolution in CSS pixels for the current display device.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/devicePixelRatio
func (_win Window) DevicePixelRatio() float64 {
	return _win.jsValue.Get("devicePixelRatio").Float()
}

// accesses a session Storage object for the current origin.
//
// sessionStorage is similar to localStorage; the difference is that while data in localStorage doesn't expire,
// data in sessionStorage is cleared when the page session ends.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/sessionStorage
func (_win Window) SessionStorage() *Storage {
	rsp := _win.jsValue.Call("ickSessionStorage")
	if typ := rsp.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	return CastStorage(rsp)
}

// allows you to access a Storage object for the Document's origin; the stored data is saved across browser sessions.
//
// returns nil if access is denied to the localstorage
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/localStorage
func (_win Window) LocalStorage() *Storage {
	//value := _win.jsValue.Get("localStorage")
	rsp := _win.jsValue.Call("ickLocalStorage")
	if typ := rsp.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	return CastStorage(rsp)
}

// instructs the browser to display a dialog with an optional message, and to wait until the user dismisses the dialog.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/alert
func (_win Window) Alert(message string) {
	message = strings.Trim(message, " ")
	if message == "" {
		_win.jsValue.Call("alert")
	} else {
		_win.jsValue.Call("alert", message)
	}
}

// Opens the print dialog to print the current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/print
func (_win Window) Print() {
	_win.jsValue.Call("print")
}

// method closes the current window, or the window on which it was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/close
func (_win Window) Close() {
	_win.jsValue.Call("close")
}

// stops further resource loading in the current browsing context, equivalent to the stop button in the browser.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/stop
func (_win Window) Stop() {
	_win.jsValue.Call("stop")
}

// Makes a request to bring the window to the front.
// It may fail due to user settings and the window isn't guaranteed to be frontmost before this method returns.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/focus
func (_win Window) Focus() {
	_win.jsValue.Call("focus")
}

// Shifts focus away from the window.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/blur_event
func (_win Window) Blur() {
	_win.jsValue.Call("blur")
}

/******************************************************************************
* Window's GENERIC_EVENT
******************************************************************************/

// event attribute: Event
func makeWindow_Generic_Event(listener func(event *Event, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastEvent(value)
		target := CastWindow(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on Window", evt.Type())
			}
		}()
		listener(evt, &target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_win Window) AddGenericEvent(evttype GENERIC_EVENT, listener func(event *Event, target *Window)) {
	jsevh := makeWindow_Generic_Event(listener)
	_win.AddEventListener(&eventHandler{eventtype: string(evttype), jsHandler: jsevh})
}

/******************************************************************************
* Window's BeforeUnloadEvent
******************************************************************************/

// event attribute: BeforeUnloadEvent
func makeWindow_BeforeUnload_Event(listener func(event *BeforeUnloadEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastBeforeUnloadEvent(value)
		target := CastWindow(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on Window", evt.Type())
			}
		}()
		listener(evt, &target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_win Window) AddBeforeUnloadEvent(listener func(event *BeforeUnloadEvent, target *Window)) {
	jsevh := makeWindow_BeforeUnload_Event(listener)
	_win.AddEventListener(&eventHandler{eventtype: "beforeunload", jsHandler: jsevh})
}

/****************************************************************************
* Window's HASCHANGED_EVENT
*****************************************************************************/

// event attribute: HashChangeEvent
func makeWindow_HashChange_Event(listener func(event *HashChangeEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastHashChangeEvent(value)
		target := CastWindow(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on Window", evt.Type())
			}
		}()
		listener(evt, &target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddHashChange is adding doing AddEventListener for 'HashChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_win Window) AddHashChangeEvent(listener func(event *HashChangeEvent, target *Window)) {
	jsevh := makeWindow_HashChange_Event(listener)
	_win.AddEventListener(&eventHandler{eventtype: "hashchange", jsHandler: jsevh})
}

/****************************************************************************
* Window's PAGETRANSITION_EVENT
*****************************************************************************/

// event attribute: PageTransitionEvent
func makeWindow_PageTransition_Event(listener func(event *PageTransitionEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastPageTransitionEvent(value)
		target := CastWindow(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on Window", evt.Type())
			}
		}()
		listener(evt, &target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_win Window) AddPageTransitionEvent(evttype PAGETRANSITION_EVENT, listener func(event *PageTransitionEvent, target *Window)) {
	jsevh := makeWindow_PageTransition_Event(listener)
	_win.AddEventListener(&eventHandler{eventtype: string(evttype), jsHandler: jsevh})
}

/****************************************************************************
* Window's UI_EVENT
*****************************************************************************/

// event attribute: UIEvent
func makeWindow_UI_Event(listener func(event *UIEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastUIEvent(value)
		target := CastWindow(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on Window", evt.Type())
			}
		}()
		listener(evt, &target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_win Window) AddResizeEvent(listener func(event *UIEvent, target *Window)) {
	jsevh := makeWindow_UI_Event(listener)
	_win.AddEventListener(&eventHandler{eventtype: "resize", jsHandler: jsevh})
}
