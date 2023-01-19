package webclientsdk

import (
	"syscall/js"
)

// Window
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window
type Window struct {
	EventTarget
}

// WindowFromJS is casting a js.Value into Window.
func WindowFromJS(value js.Value) *Window {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Window{}
	ret.jsValue = value
	return ret
}

// GetWindow returning attribute 'window' with
// type Window (idl: Window).
func GetWindow() *Window {
	var ret *Window
	_klass := js.Global()
	value := _klass.Get("window")
	ret = WindowFromJS(value)
	return ret
}

// Document returning attribute 'document' with
// type Document (idl: Document).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/document
func (_this *Window) Document() *Document {
	var ret *Document
	value := _this.jsValue.Get("document")
	ret = DocumentFromJS(value)
	return ret
}

// Location returning attribute 'location' with
// type htmlmisc.Location (idl: Location).
func (_this *Window) Location() *Location {
	var ret *Location
	value := _this.jsValue.Get("location")
	ret = LocationFromJS(value)
	return ret
}

// History returning attribute 'history' with
// type htmlmisc.History (idl: History).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Window/history
func (_this *Window) History() *History {
	var ret *History
	value := _this.jsValue.Get("history")
	ret = HistoryFromJS(value)
	return ret
}

// Closed returning attribute 'closed' with
// type bool (idl: boolean).
func (_this *Window) Closed() bool {
	var ret bool
	value := _this.jsValue.Get("closed")
	ret = (value).Bool()
	return ret
}

// Top returning attribute 'top' with
// type Window (idl: Window).
func (_this *Window) Top() *Window {
	var ret *Window
	value := _this.jsValue.Get("top")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = WindowFromJS(value)
	}
	return ret
}

// Navigator returning attribute 'navigator' with
// type htmlmisc.Navigator (idl: Navigator).
func (_this *Window) Navigator() *Navigator {
	var ret *Navigator
	value := _this.jsValue.Get("navigator")
	ret = NavigatorFromJS(value)
	return ret
}

// InnerWidth returning attribute 'innerWidth' with
// type int (idl: long).
func (_this *Window) InnerWidth() int {
	var ret int
	value := _this.jsValue.Get("innerWidth")
	ret = (value).Int()
	return ret
}

// InnerHeight returning attribute 'innerHeight' with
// type int (idl: long).
func (_this *Window) InnerHeight() int {
	var ret int
	value := _this.jsValue.Get("innerHeight")
	ret = (value).Int()
	return ret
}

// ScrollX returning attribute 'scrollX' with
// type float64 (idl: double).
func (_this *Window) ScrollX() float64 {
	var ret float64
	value := _this.jsValue.Get("scrollX")
	ret = (value).Float()
	return ret
}

// PageXOffset returning attribute 'pageXOffset' with
// type float64 (idl: double).
func (_this *Window) PageXOffset() float64 {
	var ret float64
	value := _this.jsValue.Get("pageXOffset")
	ret = (value).Float()
	return ret
}

// ScrollY returning attribute 'scrollY' with
// type float64 (idl: double).
func (_this *Window) ScrollY() float64 {
	var ret float64
	value := _this.jsValue.Get("scrollY")
	ret = (value).Float()
	return ret
}

// PageYOffset returning attribute 'pageYOffset' with
// type float64 (idl: double).
func (_this *Window) PageYOffset() float64 {
	var ret float64
	value := _this.jsValue.Get("pageYOffset")
	ret = (value).Float()
	return ret
}

// ScreenX returning attribute 'screenX' with
// type int (idl: long).
func (_this *Window) ScreenX() int {
	var ret int
	value := _this.jsValue.Get("screenX")
	ret = (value).Int()
	return ret
}

// ScreenLeft returning attribute 'screenLeft' with
// type int (idl: long).
func (_this *Window) ScreenLeft() int {
	var ret int
	value := _this.jsValue.Get("screenLeft")
	ret = (value).Int()
	return ret
}

// ScreenY returning attribute 'screenY' with
// type int (idl: long).
func (_this *Window) ScreenY() int {
	var ret int
	value := _this.jsValue.Get("screenY")
	ret = (value).Int()
	return ret
}

// ScreenTop returning attribute 'screenTop' with
// type int (idl: long).
func (_this *Window) ScreenTop() int {
	var ret int
	value := _this.jsValue.Get("screenTop")
	ret = (value).Int()
	return ret
}

// OuterWidth returning attribute 'outerWidth' with
// type int (idl: long).
func (_this *Window) OuterWidth() int {
	var ret int
	value := _this.jsValue.Get("outerWidth")
	ret = (value).Int()
	return ret
}

// OuterHeight returning attribute 'outerHeight' with
// type int (idl: long).
func (_this *Window) OuterHeight() int {
	var ret int
	value := _this.jsValue.Get("outerHeight")
	ret = (value).Int()
	return ret
}

// DevicePixelRatio returning attribute 'devicePixelRatio' with
// type float64 (idl: double).
func (_this *Window) DevicePixelRatio() float64 {
	var ret float64
	value := _this.jsValue.Get("devicePixelRatio")
	ret = (value).Float()
	return ret
}

// Origin returning attribute 'origin' with
// type string (idl: USVString).
func (_this *Window) Origin() string {
	var ret string
	value := _this.jsValue.Get("origin")
	ret = (value).String()
	return ret
}

// SessionStorage returning attribute 'sessionStorage' with
// type Storage (idl: Storage).
func (_this *Window) SessionStorage() *Storage {
	var ret *Storage
	value := _this.jsValue.Get("sessionStorage")
	ret = StorageFromJS(value)
	return ret
}

// LocalStorage returning attribute 'localStorage' with
// type Storage (idl: Storage).
func (_this *Window) LocalStorage() *Storage {
	var ret *Storage
	value := _this.jsValue.Get("localStorage")
	ret = StorageFromJS(value)
	return ret
}

// event attribute: Event
func eventFuncWindow_Event(listener func(event *Event, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *Event
		value := args[0]
		incoming := value.Get("target")
		ret = EventFromJS(value)
		src := WindowFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddAbort is adding doing AddEventListener for 'Abort' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventAbort(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "abort", cb)
	return cb
}

// SetOnAbort is assigning a function to 'onabort'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnAbort(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onabort", cb)
	return cb
}

// AddAfterPrint is adding doing AddEventListener for 'AfterPrint' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventAfterPrint(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "afterprint", cb)
	return cb
}

// SetOnAfterPrint is assigning a function to 'onafterprint'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnAfterPrint(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onafterprint", cb)
	return cb
}

// AddAppInstalled is adding doing AddEventListener for 'AppInstalled' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventAppInstalled(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "appinstalled", cb)
	return cb
}

// SetOnAppInstalled is assigning a function to 'onappinstalled'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnAppInstalled(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onappinstalled", cb)
	return cb
}

// event attribute: MouseEvent
func eventFuncWindow_MouseEvent(listener func(event *MouseEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *MouseEvent
		value := args[0]
		incoming := value.Get("target")
		ret = MouseEventFromJS(value)
		src := WindowFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddAuxclick is adding doing AddEventListener for 'Auxclick' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventAuxclick(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "auxclick", cb)
	return cb
}

// SetOnAuxclick is assigning a function to 'onauxclick'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnAuxclick(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Set("onauxclick", cb)
	return cb
}

// AddBeforePrint is adding doing AddEventListener for 'BeforePrint' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventBeforePrint(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "beforeprint", cb)
	return cb
}

// SetOnBeforePrint is assigning a function to 'onbeforeprint'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnBeforePrint(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onbeforeprint", cb)
	return cb
}

// event attribute: BeforeUnloadEvent
func eventFuncWindow_htmlcommon_BeforeUnloadEvent(listener func(event *BeforeUnloadEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *BeforeUnloadEvent
		value := args[0]
		incoming := value.Get("target")
		ret = BeforeUnloadEventFromJS(value)
		src := WindowFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddBeforeUnload is adding doing AddEventListener for 'BeforeUnload' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventBeforeUnload(listener func(event *BeforeUnloadEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_htmlcommon_BeforeUnloadEvent(listener)
	_this.jsValue.Call("addEventListener", "beforeunload", cb)
	return cb
}

// SetOnBeforeUnload is assigning a function to 'onbeforeunload'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnBeforeUnload(listener func(event *BeforeUnloadEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_htmlcommon_BeforeUnloadEvent(listener)
	_this.jsValue.Set("onbeforeunload", cb)
	return cb
}

// event attribute: FocusEvent
func eventFuncWindow_FocusEvent(listener func(event *FocusEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *FocusEvent
		value := args[0]
		incoming := value.Get("target")
		ret = FocusEventFromJS(value)
		src := WindowFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddBlur is adding doing AddEventListener for 'Blur' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventBlur(listener func(event *FocusEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_FocusEvent(listener)
	_this.jsValue.Call("addEventListener", "blur", cb)
	return cb
}

// SetOnBlur is assigning a function to 'onblur'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnBlur(listener func(event *FocusEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_FocusEvent(listener)
	_this.jsValue.Set("onblur", cb)
	return cb
}

// AddCancel is adding doing AddEventListener for 'Cancel' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventCancel(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "cancel", cb)
	return cb
}

// SetOnCancel is assigning a function to 'oncancel'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnCancel(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("oncancel", cb)
	return cb
}

// AddCanPlay is adding doing AddEventListener for 'CanPlay' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventCanPlay(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "canplay", cb)
	return cb
}

// SetOnCanPlay is assigning a function to 'oncanplay'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnCanPlay(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("oncanplay", cb)
	return cb
}

// AddCanPlayThrough is adding doing AddEventListener for 'CanPlayThrough' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventCanPlayThrough(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "canplaythrough", cb)
	return cb
}

// SetOnCanPlayThrough is assigning a function to 'oncanplaythrough'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnCanPlayThrough(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("oncanplaythrough", cb)
	return cb
}

// AddChange is adding doing AddEventListener for 'Change' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "change", cb)
	return cb
}

// SetOnChange is assigning a function to 'onchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onchange", cb)
	return cb
}

// AddClick is adding doing AddEventListener for 'Click' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventClick(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "click", cb)
	return cb
}

// SetOnClick is assigning a function to 'onclick'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnClick(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Set("onclick", cb)
	return cb
}

// AddClose is adding doing AddEventListener for 'Close' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventClose(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "close", cb)
	return cb
}

// SetOnClose is assigning a function to 'onclose'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnClose(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onclose", cb)
	return cb
}

// AddContextMenu is adding doing AddEventListener for 'ContextMenu' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventContextMenu(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "contextmenu", cb)
	return cb
}

// SetOnContextMenu is assigning a function to 'oncontextmenu'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnContextMenu(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Set("oncontextmenu", cb)
	return cb
}

// AddCueChange is adding doing AddEventListener for 'CueChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventCueChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "cuechange", cb)
	return cb
}

// SetOnCueChange is assigning a function to 'oncuechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnCueChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("oncuechange", cb)
	return cb
}

// AddDblClick is adding doing AddEventListener for 'DblClick' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventDblClick(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "dblclick", cb)
	return cb
}

// SetOnDblClick is assigning a function to 'ondblclick'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnDblClick(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Set("ondblclick", cb)
	return cb
}

// AddDurationChange is adding doing AddEventListener for 'DurationChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventDurationChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "durationchange", cb)
	return cb
}

// SetOnDurationChange is assigning a function to 'ondurationchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnDurationChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("ondurationchange", cb)
	return cb
}

// AddEmptied is adding doing AddEventListener for 'Emptied' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventEmptied(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "emptied", cb)
	return cb
}

// SetOnEmptied is assigning a function to 'onemptied'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnEmptied(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onemptied", cb)
	return cb
}

// AddEnded is adding doing AddEventListener for 'Ended' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventEnded(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "ended", cb)
	return cb
}

// SetOnEnded is assigning a function to 'onended'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnEnded(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onended", cb)
	return cb
}

// AddError is adding doing AddEventListener for 'Error' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventError(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "error", cb)
	return cb
}

// SetOnError is assigning a function to 'onerror'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnError(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onerror", cb)
	return cb
}

// AddFocus is adding doing AddEventListener for 'Focus' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventFocus(listener func(event *FocusEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_FocusEvent(listener)
	_this.jsValue.Call("addEventListener", "focus", cb)
	return cb
}

// SetOnFocus is assigning a function to 'onfocus'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnFocus(listener func(event *FocusEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_FocusEvent(listener)
	_this.jsValue.Set("onfocus", cb)
	return cb
}

// event attribute: PointerEvent
func eventFuncWindow_PointerEvent(listener func(event *PointerEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *PointerEvent
		value := args[0]
		incoming := value.Get("target")
		ret = PointerEventFromJS(value)
		src := WindowFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddGotPointerCapture is adding doing AddEventListener for 'GotPointerCapture' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventGotPointerCapture(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "gotpointercapture", cb)
	return cb
}

// SetOnGotPointerCapture is assigning a function to 'ongotpointercapture'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnGotPointerCapture(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Set("ongotpointercapture", cb)
	return cb
}

// event attribute: HashChangeEvent
func eventFuncWindow_HashChangeEvent(listener func(event *HashChangeEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *HashChangeEvent
		value := args[0]
		incoming := value.Get("target")
		ret = HashChangeEventFromJS(value)
		src := WindowFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddHashChange is adding doing AddEventListener for 'HashChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventHashChange(listener func(event *HashChangeEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_HashChangeEvent(listener)
	_this.jsValue.Call("addEventListener", "hashchange", cb)
	return cb
}

// SetOnHashChange is assigning a function to 'onhashchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnHashChange(listener func(event *HashChangeEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_HashChangeEvent(listener)
	_this.jsValue.Set("onhashchange", cb)
	return cb
}

// event attribute: InputEvent
func eventFuncWindow_InputEvent(listener func(event *InputEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *InputEvent
		value := args[0]
		incoming := value.Get("target")
		ret = InputEventFromJS(value)
		src := WindowFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddInput is adding doing AddEventListener for 'Input' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventInput(listener func(event *InputEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_InputEvent(listener)
	_this.jsValue.Call("addEventListener", "input", cb)
	return cb
}

// SetOnInput is assigning a function to 'oninput'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnInput(listener func(event *InputEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_InputEvent(listener)
	_this.jsValue.Set("oninput", cb)
	return cb
}

// AddInvalid is adding doing AddEventListener for 'Invalid' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventInvalid(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "invalid", cb)
	return cb
}

// SetOnInvalid is assigning a function to 'oninvalid'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnInvalid(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("oninvalid", cb)
	return cb
}

// event attribute: KeyboardEvent
func eventFuncWindow_KeyboardEvent(listener func(event *KeyboardEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *KeyboardEvent
		value := args[0]
		incoming := value.Get("target")
		ret = KeyboardEventFromJS(value)
		src := WindowFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddKeyDown is adding doing AddEventListener for 'KeyDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventKeyDown(listener func(event *KeyboardEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_KeyboardEvent(listener)
	_this.jsValue.Call("addEventListener", "keydown", cb)
	return cb
}

// SetOnKeyDown is assigning a function to 'onkeydown'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnKeyDown(listener func(event *KeyboardEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_KeyboardEvent(listener)
	_this.jsValue.Set("onkeydown", cb)
	return cb
}

// AddKeyPress is adding doing AddEventListener for 'KeyPress' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventKeyPress(listener func(event *KeyboardEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_KeyboardEvent(listener)
	_this.jsValue.Call("addEventListener", "keypress", cb)
	return cb
}

// SetOnKeyPress is assigning a function to 'onkeypress'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnKeyPress(listener func(event *KeyboardEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_KeyboardEvent(listener)
	_this.jsValue.Set("onkeypress", cb)
	return cb
}

// AddKeyUp is adding doing AddEventListener for 'KeyUp' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventKeyUp(listener func(event *KeyboardEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_KeyboardEvent(listener)
	_this.jsValue.Call("addEventListener", "keyup", cb)
	return cb
}

// SetOnKeyUp is assigning a function to 'onkeyup'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnKeyUp(listener func(event *KeyboardEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_KeyboardEvent(listener)
	_this.jsValue.Set("onkeyup", cb)
	return cb
}

// AddLanguageChange is adding doing AddEventListener for 'LanguageChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventLanguageChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "languagechange", cb)
	return cb
}

// SetOnLanguageChange is assigning a function to 'onlanguagechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnLanguageChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onlanguagechange", cb)
	return cb
}

// AddLoad is adding doing AddEventListener for 'Load' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventLoad(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "load", cb)
	return cb
}

// SetOnLoad is assigning a function to 'onload'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnLoad(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onload", cb)
	return cb
}

// AddLoadedData is adding doing AddEventListener for 'LoadedData' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventLoadedData(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "loadeddata", cb)
	return cb
}

// SetOnLoadedData is assigning a function to 'onloadeddata'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnLoadedData(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onloadeddata", cb)
	return cb
}

// AddLoadedMetaData is adding doing AddEventListener for 'LoadedMetaData' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventLoadedMetaData(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "loadedmetadata", cb)
	return cb
}

// SetOnLoadedMetaData is assigning a function to 'onloadedmetadata'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnLoadedMetaData(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onloadedmetadata", cb)
	return cb
}

// AddLoadStart is adding doing AddEventListener for 'LoadStart' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventLoadStart(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "loadstart", cb)
	return cb
}

// SetOnLoadStart is assigning a function to 'onloadstart'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnLoadStart(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onloadstart", cb)
	return cb
}

// AddLostPointerCapture is adding doing AddEventListener for 'LostPointerCapture' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventLostPointerCapture(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "lostpointercapture", cb)
	return cb
}

// SetOnLostPointerCapture is assigning a function to 'onlostpointercapture'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnLostPointerCapture(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Set("onlostpointercapture", cb)
	return cb
}

// AddMouseDown is adding doing AddEventListener for 'MouseDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventMouseDown(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mousedown", cb)
	return cb
}

// SetOnMouseDown is assigning a function to 'onmousedown'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnMouseDown(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Set("onmousedown", cb)
	return cb
}

// AddMouseEnter is adding doing AddEventListener for 'MouseEnter' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventMouseEnter(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseenter", cb)
	return cb
}

// SetOnMouseEnter is assigning a function to 'onmouseenter'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnMouseEnter(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Set("onmouseenter", cb)
	return cb
}

// AddMouseLeave is adding doing AddEventListener for 'MouseLeave' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventMouseLeave(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseleave", cb)
	return cb
}

// SetOnMouseLeave is assigning a function to 'onmouseleave'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnMouseLeave(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Set("onmouseleave", cb)
	return cb
}

// AddMouseMove is adding doing AddEventListener for 'MouseMove' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventMouseMove(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mousemove", cb)
	return cb
}

// SetOnMouseMove is assigning a function to 'onmousemove'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnMouseMove(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Set("onmousemove", cb)
	return cb
}

// AddMouseOut is adding doing AddEventListener for 'MouseOut' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventMouseOut(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseout", cb)
	return cb
}

// SetOnMouseOut is assigning a function to 'onmouseout'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnMouseOut(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Set("onmouseout", cb)
	return cb
}

// AddMouseOver is adding doing AddEventListener for 'MouseOver' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventMouseOver(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseover", cb)
	return cb
}

// SetOnMouseOver is assigning a function to 'onmouseover'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnMouseOver(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Set("onmouseover", cb)
	return cb
}

// AddMouseUp is adding doing AddEventListener for 'MouseUp' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventMouseUp(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseup", cb)
	return cb
}

// SetOnMouseUp is assigning a function to 'onmouseup'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnMouseUp(listener func(event *MouseEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_MouseEvent(listener)
	_this.jsValue.Set("onmouseup", cb)
	return cb
}

// AddOffline is adding doing AddEventListener for 'Offline' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventOffline(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "offline", cb)
	return cb
}

// SetOnOffline is assigning a function to 'onoffline'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnOffline(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onoffline", cb)
	return cb
}

// AddOnline is adding doing AddEventListener for 'Online' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventOnline(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "online", cb)
	return cb
}

// SetOnOnline is assigning a function to 'ononline'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnOnline(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("ononline", cb)
	return cb
}

// AddOrientationchange is adding doing AddEventListener for 'Orientationchange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventOrientationchange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "orientationchange", cb)
	return cb
}

// SetOnOrientationchange is assigning a function to 'onorientationchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnOrientationchange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onorientationchange", cb)
	return cb
}

// event attribute: PageTransitionEvent
func eventFuncWindow_PageTransitionEvent(listener func(event *PageTransitionEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *PageTransitionEvent
		value := args[0]
		incoming := value.Get("target")
		ret = PageTransitionEventFromJS(value)
		src := WindowFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddPageHide is adding doing AddEventListener for 'PageHide' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPageHide(listener func(event *PageTransitionEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PageTransitionEvent(listener)
	_this.jsValue.Call("addEventListener", "pagehide", cb)
	return cb
}

// SetOnPageHide is assigning a function to 'onpagehide'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPageHide(listener func(event *PageTransitionEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PageTransitionEvent(listener)
	_this.jsValue.Set("onpagehide", cb)
	return cb
}

// AddPageShow is adding doing AddEventListener for 'PageShow' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPageShow(listener func(event *PageTransitionEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PageTransitionEvent(listener)
	_this.jsValue.Call("addEventListener", "pageshow", cb)
	return cb
}

// SetOnPageShow is assigning a function to 'onpageshow'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPageShow(listener func(event *PageTransitionEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PageTransitionEvent(listener)
	_this.jsValue.Set("onpageshow", cb)
	return cb
}

// AddPause is adding doing AddEventListener for 'Pause' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPause(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "pause", cb)
	return cb
}

// SetOnPause is assigning a function to 'onpause'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPause(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onpause", cb)
	return cb
}

// AddPlay is adding doing AddEventListener for 'Play' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPlay(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "play", cb)
	return cb
}

// SetOnPlay is assigning a function to 'onplay'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPlay(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onplay", cb)
	return cb
}

// AddPlaying is adding doing AddEventListener for 'Playing' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPlaying(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "playing", cb)
	return cb
}

// SetOnPlaying is assigning a function to 'onplaying'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPlaying(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onplaying", cb)
	return cb
}

// AddPointerCancel is adding doing AddEventListener for 'PointerCancel' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPointerCancel(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointercancel", cb)
	return cb
}

// SetOnPointerCancel is assigning a function to 'onpointercancel'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPointerCancel(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Set("onpointercancel", cb)
	return cb
}

// AddPointerDown is adding doing AddEventListener for 'PointerDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPointerDown(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerdown", cb)
	return cb
}

// SetOnPointerDown is assigning a function to 'onpointerdown'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPointerDown(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Set("onpointerdown", cb)
	return cb
}

// AddPointerEnter is adding doing AddEventListener for 'PointerEnter' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPointerEnter(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerenter", cb)
	return cb
}

// SetOnPointerEnter is assigning a function to 'onpointerenter'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPointerEnter(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Set("onpointerenter", cb)
	return cb
}

// AddPointerLeave is adding doing AddEventListener for 'PointerLeave' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPointerLeave(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerleave", cb)
	return cb
}

// SetOnPointerLeave is assigning a function to 'onpointerleave'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPointerLeave(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Set("onpointerleave", cb)
	return cb
}

// AddPointerMove is adding doing AddEventListener for 'PointerMove' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPointerMove(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointermove", cb)
	return cb
}

// SetOnPointerMove is assigning a function to 'onpointermove'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPointerMove(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Set("onpointermove", cb)
	return cb
}

// AddPointerOut is adding doing AddEventListener for 'PointerOut' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPointerOut(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerout", cb)
	return cb
}

// SetOnPointerOut is assigning a function to 'onpointerout'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPointerOut(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Set("onpointerout", cb)
	return cb
}

// AddPointerOver is adding doing AddEventListener for 'PointerOver' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPointerOver(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerover", cb)
	return cb
}

// SetOnPointerOver is assigning a function to 'onpointerover'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPointerOver(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Set("onpointerover", cb)
	return cb
}

// AddPointerUp is adding doing AddEventListener for 'PointerUp' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventPointerUp(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerup", cb)
	return cb
}

// SetOnPointerUp is assigning a function to 'onpointerup'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnPointerUp(listener func(event *PointerEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_PointerEvent(listener)
	_this.jsValue.Set("onpointerup", cb)
	return cb
}

// AddRateChange is adding doing AddEventListener for 'RateChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventRateChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "ratechange", cb)
	return cb
}

// SetOnRateChange is assigning a function to 'onratechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnRateChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onratechange", cb)
	return cb
}

// AddReset is adding doing AddEventListener for 'Reset' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventReset(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "reset", cb)
	return cb
}

// SetOnReset is assigning a function to 'onreset'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnReset(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onreset", cb)
	return cb
}

// event attribute: UIEvent
func eventFuncWindow_UIEvent(listener func(event *UIEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *UIEvent
		value := args[0]
		incoming := value.Get("target")
		ret = UIEventFromJS(value)
		src := WindowFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddResize is adding doing AddEventListener for 'Resize' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventResize(listener func(event *UIEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_UIEvent(listener)
	_this.jsValue.Call("addEventListener", "resize", cb)
	return cb
}

// SetOnResize is assigning a function to 'onresize'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnResize(listener func(event *UIEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_UIEvent(listener)
	_this.jsValue.Set("onresize", cb)
	return cb
}

// AddScroll is adding doing AddEventListener for 'Scroll' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventScroll(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "scroll", cb)
	return cb
}

// SetOnScroll is assigning a function to 'onscroll'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnScroll(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onscroll", cb)
	return cb
}

// AddSeeked is adding doing AddEventListener for 'Seeked' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventSeeked(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "seeked", cb)
	return cb
}

// SetOnSeeked is assigning a function to 'onseeked'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnSeeked(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onseeked", cb)
	return cb
}

// AddSeeking is adding doing AddEventListener for 'Seeking' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventSeeking(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "seeking", cb)
	return cb
}

// SetOnSeeking is assigning a function to 'onseeking'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnSeeking(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onseeking", cb)
	return cb
}

// AddSelect is adding doing AddEventListener for 'Select' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventSelect(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "select", cb)
	return cb
}

// SetOnSelect is assigning a function to 'onselect'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnSelect(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onselect", cb)
	return cb
}

// AddSelectionChange is adding doing AddEventListener for 'SelectionChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventSelectionChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "selectionchange", cb)
	return cb
}

// SetOnSelectionChange is assigning a function to 'onselectionchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnSelectionChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onselectionchange", cb)
	return cb
}

// AddSelectStart is adding doing AddEventListener for 'SelectStart' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventSelectStart(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "selectstart", cb)
	return cb
}

// SetOnSelectStart is assigning a function to 'onselectstart'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnSelectStart(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onselectstart", cb)
	return cb
}

// AddStalled is adding doing AddEventListener for 'Stalled' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventStalled(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "stalled", cb)
	return cb
}

// SetOnStalled is assigning a function to 'onstalled'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnStalled(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onstalled", cb)
	return cb
}

// AddSubmit is adding doing AddEventListener for 'Submit' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventSubmit(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "submit", cb)
	return cb
}

// SetOnSubmit is assigning a function to 'onsubmit'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnSubmit(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onsubmit", cb)
	return cb
}

// AddSuspend is adding doing AddEventListener for 'Suspend' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventSuspend(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "suspend", cb)
	return cb
}

// SetOnSuspend is assigning a function to 'onsuspend'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnSuspend(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onsuspend", cb)
	return cb
}

// AddTimeUpdate is adding doing AddEventListener for 'TimeUpdate' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventTimeUpdate(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "timeupdate", cb)
	return cb
}

// SetOnTimeUpdate is assigning a function to 'ontimeupdate'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnTimeUpdate(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("ontimeupdate", cb)
	return cb
}

// AddToggle is adding doing AddEventListener for 'Toggle' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventToggle(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "toggle", cb)
	return cb
}

// SetOnToggle is assigning a function to 'ontoggle'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnToggle(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("ontoggle", cb)
	return cb
}

// AddUnload is adding doing AddEventListener for 'Unload' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventUnload(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "unload", cb)
	return cb
}

// SetOnUnload is assigning a function to 'onunload'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnUnload(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onunload", cb)
	return cb
}

// AddVolumeChange is adding doing AddEventListener for 'VolumeChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventVolumeChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "volumechange", cb)
	return cb
}

// SetOnVolumeChange is assigning a function to 'onvolumechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnVolumeChange(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onvolumechange", cb)
	return cb
}

// AddWaiting is adding doing AddEventListener for 'Waiting' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventWaiting(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Call("addEventListener", "waiting", cb)
	return cb
}

// SetOnWaiting is assigning a function to 'onwaiting'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnWaiting(listener func(event *Event, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_Event(listener)
	_this.jsValue.Set("onwaiting", cb)
	return cb
}

// event attribute: WheelEvent
func eventFuncWindow_WheelEvent(listener func(event *WheelEvent, target *Window)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *WheelEvent
		value := args[0]
		incoming := value.Get("target")
		ret = WheelEventFromJS(value)
		src := WindowFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddWheel is adding doing AddEventListener for 'Wheel' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Window) AddEventWheel(listener func(event *WheelEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_WheelEvent(listener)
	_this.jsValue.Call("addEventListener", "wheel", cb)
	return cb
}

// SetOnWheel is assigning a function to 'onwheel'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Window) SetOnWheel(listener func(event *WheelEvent, currentTarget *Window)) js.Func {
	cb := eventFuncWindow_WheelEvent(listener)
	_this.jsValue.Set("onwheel", cb)
	return cb
}

func (_this *Window) Close() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("close", _args[0:_end]...)
	return
}

func (_this *Window) Stop() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("stop", _args[0:_end]...)
	return
}

func (_this *Window) Focus() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("focus", _args[0:_end]...)
	return
}

func (_this *Window) Blur() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("blur", _args[0:_end]...)
	return
}

func (_this *Window) Open(url *string, target *string, features *string) (_result *Window) {
	var (
		_args [3]interface{}
		_end  int
	)
	if url != nil {

		var _p0 interface{}
		if url != nil {
			_p0 = *(url)
		} else {
			_p0 = nil
		}
		_args[0] = _p0
		_end++
	}
	if target != nil {

		var _p1 interface{}
		if target != nil {
			_p1 = *(target)
		} else {
			_p1 = nil
		}
		_args[1] = _p1
		_end++
	}
	if features != nil {

		var _p2 interface{}
		if features != nil {
			_p2 = *(features)
		} else {
			_p2 = nil
		}
		_args[2] = _p2
		_end++
	}
	_returned := _this.jsValue.Call("open", _args[0:_end]...)
	var (
		_converted *Window // javascript: Window _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = WindowFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *Window) Alert() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("alert", _args[0:_end]...)
	return
}

func (_this *Window) AlertMessage(message string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := message
	_args[0] = _p0
	_end++
	_this.jsValue.Call("alert", _args[0:_end]...)
	return
}

func (_this *Window) Confirm(message *string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	if message != nil {

		var _p0 interface{}
		if message != nil {
			_p0 = *(message)
		} else {
			_p0 = nil
		}
		_args[0] = _p0
		_end++
	}
	_returned := _this.jsValue.Call("confirm", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *Window) Prompt(message *string, _default *string) (_result *string) {
	var (
		_args [2]interface{}
		_end  int
	)
	if message != nil {

		var _p0 interface{}
		if message != nil {
			_p0 = *(message)
		} else {
			_p0 = nil
		}
		_args[0] = _p0
		_end++
	}
	if _default != nil {

		var _p1 interface{}
		if _default != nil {
			_p1 = *(_default)
		} else {
			_p1 = nil
		}
		_args[1] = _p1
		_end++
	}
	_returned := _this.jsValue.Call("prompt", _args[0:_end]...)
	var (
		_converted *string // javascript: DOMString _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		__tmp := (_returned).String()
		_converted = &__tmp
	}
	_result = _converted
	return
}

func (_this *Window) Print() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("print", _args[0:_end]...)
	return
}

func (_this *Window) MoveTo(x int, y int) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_this.jsValue.Call("moveTo", _args[0:_end]...)
	return
}

func (_this *Window) MoveBy(x int, y int) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_this.jsValue.Call("moveBy", _args[0:_end]...)
	return
}

func (_this *Window) ResizeTo(x int, y int) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_this.jsValue.Call("resizeTo", _args[0:_end]...)
	return
}

func (_this *Window) ResizeBy(x int, y int) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_this.jsValue.Call("resizeBy", _args[0:_end]...)
	return
}

func (_this *Window) ScrollXY(x float64, y float64) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_this.jsValue.Call("scroll", _args[0:_end]...)
	return
}

func (_this *Window) ScrollToXY(x float64, y float64) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_this.jsValue.Call("scrollTo", _args[0:_end]...)
	return
}

func (_this *Window) ScrollByXY(x float64, y float64) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := x
	_args[0] = _p0
	_end++
	_p1 := y
	_args[1] = _p1
	_end++
	_this.jsValue.Call("scrollBy", _args[0:_end]...)
	return
}
