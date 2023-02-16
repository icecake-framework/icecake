package dom

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
func CastWindow(value js.Value) *Window {
	if value.Type() != js.TypeObject {
		ConsoleError("casting Window failed")
		return nil
	}
	ret := new(Window)
	ret.jsValue = value
	return ret
}

// GetWindow returning attribute 'window' with
// type Window (idl: Window).
func GetWindow() *Window {
	value := js.Global().Get("window")
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		ConsoleError("Unable to get window")
		panic("Unable to get window")
	}
	return CastWindow(value)
}

/******************************************************************************
* Window's properties
******************************************************************************/

// Document returning attribute 'document' with
// type Document (idl: Document).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/document
func (_this *Window) GetDocument() *Document {
	value := _this.jsValue.Get("document")
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		ConsoleError("Unable to get document")
		panic("Unable to get document")
	}
	return CastDocument(value)
}

// Location represents the URL of the current window.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/location
func (_this *Window) Location() *Location {
	value := _this.jsValue.Get("location")
	return CastLocation(value)
}

// History returning attribute 'history' with
// type htmlmisc.History (idl: History).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/history
func (_this *Window) History() *History {
	var ret *History
	value := _this.jsValue.Get("history")
	ret = CastHistory(value)
	return ret
}

// Closed ndicates whether the referenced window is closed or not.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/closed
func (_this *Window) Closed() bool {
	return _this.jsValue.Get("closed").Bool()
}

// Top returns a reference to the topmost window in the window hierarchy.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/top
func (_this *Window) Top() *Window {
	value := _this.jsValue.Get("top")
	return CastWindow(value)
}

// Navigator eturns a reference to the Navigator object, which has methods and properties about the application running the script.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/navigator
func (_this *Window) Navigator() *Navigator {
	var ret *Navigator
	value := _this.jsValue.Get("navigator")
	ret = CastNavigator(value)
	return ret
}

// InnerWidth returns the interior width of the window in pixels. This includes the width of the vertical scroll bar, if one is present.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/innerWidth
func (_this *Window) InnerWidth() int {
	return _this.jsValue.Get("innerWidth").Int()
}

// InnerHeight returns the interior height of the window in pixels, including the height of the horizontal scroll bar, if present.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/innerHeight
func (_this *Window) InnerHeight() int {
	return _this.jsValue.Get("innerHeight").Int()
}

// ScrollX returns the number of pixels that the document is currently scrolled horizontally.
// This value is subpixel precise in modern browsers, meaning that it isn't necessarily a whole number.
// You can get the number of pixels the document is scrolled vertically from the scrollY property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/scrollX
func (_this *Window) ScrollX() float64 {
	return _this.jsValue.Get("scrollX").Float()
}

// PageXOffset returning attribute 'pageXOffset' with
// type float64 (idl: double).
func (_this *Window) PageXOffset() float64 {
	return _this.jsValue.Get("pageXOffset").Float()
}

// ScrollY returning attribute 'scrollY' with
// type float64 (idl: double).
func (_this *Window) ScrollY() float64 {
	return _this.jsValue.Get("scrollY").Float()
}

// PageYOffset returning attribute 'pageYOffset' with
// type float64 (idl: double).
func (_this *Window) PageYOffset() float64 {
	return _this.jsValue.Get("pageYOffset").Float()
}

// ScreenX returning attribute 'screenX' with
func (_this *Window) ScreenX() int {
	return _this.jsValue.Get("screenX").Int()
}

// ScreenLeft returning attribute 'screenLeft' with
func (_this *Window) ScreenLeft() int {
	return _this.jsValue.Get("screenLeft").Int()
}

// ScreenY returning attribute 'screenY' with
func (_this *Window) ScreenY() int {
	return _this.jsValue.Get("screenY").Int()
}

// ScreenTop returning attribute 'screenTop' with
func (_this *Window) ScreenTop() int {
	return _this.jsValue.Get("screenTop").Int()
}

// OuterWidth returning attribute 'outerWidth' with
func (_this *Window) OuterWidth() int {
	return _this.jsValue.Get("outerWidth").Int()
}

// OuterHeight returning attribute 'outerHeight' with
func (_this *Window) OuterHeight() int {
	return _this.jsValue.Get("outerHeight").Int()
}

// DevicePixelRatio returns the ratio of the resolution in physical pixels to the resolution in CSS pixels for the current display device.
// https://developer.mozilla.org/en-US/docs/Web/API/Window/devicePixelRatio
func (_this *Window) DevicePixelRatio() float64 {
	var ret float64
	value := _this.jsValue.Get("devicePixelRatio")
	ret = (value).Float()
	return ret
}

// accesses a session Storage object for the current origin.
//
// sessionStorage is similar to localStorage; the difference is that while data in localStorage doesn't expire,
// data in sessionStorage is cleared when the page session ends.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/sessionStorage
func (_this *Window) SessionStorage() *Storage {
	value := _this.jsValue.Get("sessionStorage")
	return CastStorage(value)
}

// allows you to access a Storage object for the Document's origin; the stored data is saved across browser sessions.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/localStorage
func (_this *Window) LocalStorage() *Storage {
	var ret *Storage
	value := _this.jsValue.Get("localStorage")
	ret = CastStorage(value)
	return ret
}

// Loads a specified resource into a new or existing browsing context (that is, a tab, a window, or an iframe) under a specified name.
//
// The special target keywords, _self, _blank, _parent, and _top, can also be used.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/open
func (_this *Window) Open(url *url.URL, target string) *Window {
	var win js.Value
	if url == nil {
		// a blank page is opened into the targeted browsing context.
		win = _this.jsValue.Call("open")
	} else {
		win = _this.jsValue.Call("open", url.String(), target)

	}
	return CastWindow(win)
}

// instructs the browser to display a dialog with an optional message, and to wait until the user dismisses the dialog.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/alert
func (_this *Window) Alert(message string) {
	message = strings.Trim(message, " ")
	if message == "" {
		_this.jsValue.Call("alert")
	} else {
		_this.jsValue.Call("alert", message)
	}
}

// instructs the browser to display a dialog with an optional message, and to wait until the user either confirms or cancels the dialog.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/confirm
func (_this *Window) Confirm(message string) bool {
	ok := _this.jsValue.Call("confirm", message)
	return (ok).Bool()
}

// instructs the browser to display a dialog with an optional message prompting the user to input some text,
// and to wait until the user either submits the text or cancels the dialog.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/prompt
func (_this *Window) Prompt(message string, _default string) (_rsp string) {
	var rsp js.Value
	if message == "" {
		rsp = _this.jsValue.Call("prompt")
	} else {
		rsp = _this.jsValue.Call("prompt", message)
	}

	if rsp.Type() != js.TypeNull && rsp.Type() != js.TypeUndefined {
		_rsp = rsp.String()
	}
	return _rsp
}

// Opens the print dialog to print the current document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/print
func (_this *Window) Print() {
	_this.jsValue.Call("print")
}

// method closes the current window, or the window on which it was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/close
func (_this *Window) Close() {
	_this.jsValue.Call("close")
}

// stops further resource loading in the current browsing context, equivalent to the stop button in the browser.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/stop
func (_this *Window) Stop() {
	_this.jsValue.Call("stop")
}

// Makes a request to bring the window to the front.
// It may fail due to user settings and the window isn't guaranteed to be frontmost before this method returns.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/focus
func (_this *Window) Focus() {
	_this.jsValue.Call("focus")
}

// Shifts focus away from the window.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/blur_event
func (_this *Window) Blur() {
	_this.jsValue.Call("blur")
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
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_this *Window) AddGenericEvent(evttype GENERIC_EVENT, listener func(event *Event, target *Window)) js.Func {
	callback := makeWindow_Generic_Event(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/******************************************************************************
* Window's MOUSE_EVENT
******************************************************************************/

// event attribute: MouseEvent
func makeWindow_Mouse_Event(listener func(event *MouseEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *MouseEvent
		value := args[0]
		incoming := value.Get("target")
		ret = CastMouseEvent(value)
		src := CastWindow(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_this *Window) AddMouseEvent(evttype MOUSE_EVENT, listener func(event *MouseEvent, target *Window)) js.Func {
	callback := makeWindow_Mouse_Event(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/******************************************************************************
* Window's BeforeUnloadEvent
******************************************************************************/

// event attribute: BeforeUnloadEvent
func makeWindow_BeforeUnload_Event(listener func(event *BeforeUnloadEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *BeforeUnloadEvent
		value := args[0]
		incoming := value.Get("target")
		ret = CastBeforeUnloadEvent(value)
		src := CastWindow(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_this *Window) AddBeforeUnloadEvent(listener func(event *BeforeUnloadEvent, target *Window)) js.Func {
	callback := makeWindow_BeforeUnload_Event(listener)
	_this.jsValue.Call("addEventListener", "beforeunload", callback)
	return callback
}

/******************************************************************************
* Window's FOCUS_EVENT
******************************************************************************/

// event attribute: FocusEvent
func makeWindow_Focus_Event(listener func(event *FocusEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *FocusEvent
		value := args[0]
		incoming := value.Get("target")
		ret = CastFocusEvent(value)
		src := CastWindow(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddBlur is adding doing AddEventListener for 'Blur' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddFocusEvent(evttype FOCUS_EVENT, listener func(event *FocusEvent, target *Window)) js.Func {
	callback := makeWindow_Focus_Event(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/****************************************************************************
* Window's POINTER_EVENT
*****************************************************************************/

// event attribute: PointerEvent
func makeWindow_Pointer_Event(listener func(event *PointerEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *PointerEvent
		value := args[0]
		incoming := value.Get("target")
		ret = CastPointerEvent(value)
		src := CastWindow(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_this *Window) AddPointerEvent(evttype POINTER_EVENT, listener func(event *PointerEvent, target *Window)) js.Func {
	callback := makeWindow_Pointer_Event(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/****************************************************************************
* Window's HASCHANGED_EVENT
*****************************************************************************/

// event attribute: HashChangeEvent
func makeWindow_HashChange_Event(listener func(event *HashChangeEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *HashChangeEvent
		value := args[0]
		incoming := value.Get("target")
		ret = CastHashChangeEvent(value)
		src := CastWindow(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddHashChange is adding doing AddEventListener for 'HashChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddHashChangeEvent(listener func(event *HashChangeEvent, target *Window)) js.Func {
	cb := makeWindow_HashChange_Event(listener)
	_this.jsValue.Call("addEventListener", "hashchange", cb)
	return cb
}

/****************************************************************************
* Window's INPUT_EVENT
*****************************************************************************/

// event attribute: InputEvent
func makeWindow_Input_Event(listener func(event *InputEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *InputEvent
		value := args[0]
		incoming := value.Get("target")
		ret = CastInputEvent(value)
		src := CastWindow(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddInput is adding doing AddEventListener for 'Input' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddInputEvent(evttype INPUT_EVENT, listener func(event *InputEvent, target *Window)) js.Func {
	callback := makeWindow_Input_Event(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/****************************************************************************
* Window's KEYBOARD_EVENT
*****************************************************************************/

// event attribute: KeyboardEvent
func makeWindow_Keyboard_Event(listener func(event *KeyboardEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *KeyboardEvent
		value := args[0]
		incoming := value.Get("target")
		ret = CastKeyboardEvent(value)
		src := CastWindow(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_this *Window) AddKeyboardEvent(evttype KEYBOARD_EVENT, listener func(event *KeyboardEvent, target *Window)) js.Func {
	callback := makeWindow_Keyboard_Event(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/****************************************************************************
* Window's PAGETRANSITION_EVENT
*****************************************************************************/

// event attribute: PageTransitionEvent
func makeWindow_PageTransition_Event(listener func(event *PageTransitionEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *PageTransitionEvent
		value := args[0]
		incoming := value.Get("target")
		ret = CastPageTransitionEvent(value)
		src := CastWindow(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_this *Window) AddPageTransitionEvent(evttype PAGETRANSITION_EVENT, listener func(event *PageTransitionEvent, target *Window)) js.Func {
	callback := makeWindow_PageTransition_Event(listener)
	_this.jsValue.Call("addEventListener", string(evttype), callback)
	return callback
}

/****************************************************************************
* Window's UI_EVENT
*****************************************************************************/

// event attribute: UIEvent
func makeWindow_UI_Event(listener func(event *UIEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *UIEvent
		value := args[0]
		incoming := value.Get("target")
		ret = CastUIEvent(value)
		src := CastWindow(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

func (_this *Window) AddResizeEvent(listener func(event *UIEvent, target *Window)) js.Func {
	callback := makeWindow_UI_Event(listener)
	_this.jsValue.Call("addEventListener", "resize", callback)
	return callback
}

/****************************************************************************
* Window's WHEEL_EVENT
*****************************************************************************/

// event attribute: WheelEvent
func makeWindow_Wheel_Event(listener func(event *WheelEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *WheelEvent
		value := args[0]
		incoming := value.Get("target")
		ret = CastWheelEvent(value)
		src := CastWindow(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddWheel is adding doing AddEventListener for 'Wheel' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddWheelEvent(listener func(event *WheelEvent, target *Window)) js.Func {
	callback := makeWindow_Wheel_Event(listener)
	_this.jsValue.Call("addEventListener", "wheel", callback)
	return callback
}
