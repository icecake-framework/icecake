package browser

import (
	"syscall/js"
)

/****************************************************************************
* HTMLElement
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement
type HTMLElement struct {
	Element
}

// HTMLElementFromJS is casting a js.Value into HTMLElement.
func HTMLElementFromJS(value js.Value) *HTMLElement {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &HTMLElement{}
	ret.jsValue = value
	return ret
}

// Title returning attribute 'title' with
// type string (idl: DOMString).
func (_this *HTMLElement) Title() string {
	var ret string
	value := _this.jsValue.Get("title")
	ret = (value).String()
	return ret
}

// SetTitle setting attribute 'title' with
// type string (idl: DOMString).
func (_this *HTMLElement) SetTitle(value string) {
	input := value
	_this.jsValue.Set("title", input)
}

// Lang returning attribute 'lang' with
// type string (idl: DOMString).
func (_this *HTMLElement) Lang() string {
	var ret string
	value := _this.jsValue.Get("lang")
	ret = (value).String()
	return ret
}

// SetLang setting attribute 'lang' with
// type string (idl: DOMString).
func (_this *HTMLElement) SetLang(value string) {
	input := value
	_this.jsValue.Set("lang", input)
}

// Translate returning attribute 'translate' with
// type bool (idl: boolean).
func (_this *HTMLElement) Translate() bool {
	var ret bool
	value := _this.jsValue.Get("translate")
	ret = (value).Bool()
	return ret
}

// SetTranslate setting attribute 'translate' with
// type bool (idl: boolean).
func (_this *HTMLElement) SetTranslate(value bool) {
	input := value
	_this.jsValue.Set("translate", input)
}

// Dir returning attribute 'dir' with
// type string (idl: DOMString).
func (_this *HTMLElement) Dir() string {
	var ret string
	value := _this.jsValue.Get("dir")
	ret = (value).String()
	return ret
}

// SetDir setting attribute 'dir' with
// type string (idl: DOMString).
func (_this *HTMLElement) SetDir(value string) {
	input := value
	_this.jsValue.Set("dir", input)
}

// Hidden returning attribute 'hidden' with
// type bool (idl: boolean).
func (_this *HTMLElement) Hidden() bool {
	var ret bool
	value := _this.jsValue.Get("hidden")
	ret = (value).Bool()
	return ret
}

// SetHidden setting attribute 'hidden' with
// type bool (idl: boolean).
func (_this *HTMLElement) SetHidden(value bool) {
	input := value
	_this.jsValue.Set("hidden", input)
}

// AccessKey returning attribute 'accessKey' with
// type string (idl: DOMString).
func (_this *HTMLElement) AccessKey() string {
	var ret string
	value := _this.jsValue.Get("accessKey")
	ret = (value).String()
	return ret
}

// SetAccessKey setting attribute 'accessKey' with
// type string (idl: DOMString).
func (_this *HTMLElement) SetAccessKey(value string) {
	input := value
	_this.jsValue.Set("accessKey", input)
}

// AccessKeyLabel returning attribute 'accessKeyLabel' with
// type string (idl: DOMString).
func (_this *HTMLElement) AccessKeyLabel() string {
	var ret string
	value := _this.jsValue.Get("accessKeyLabel")
	ret = (value).String()
	return ret
}

// Draggable returning attribute 'draggable' with
// type bool (idl: boolean).
func (_this *HTMLElement) Draggable() bool {
	var ret bool
	value := _this.jsValue.Get("draggable")
	ret = (value).Bool()
	return ret
}

// SetDraggable setting attribute 'draggable' with
// type bool (idl: boolean).
func (_this *HTMLElement) SetDraggable(value bool) {
	input := value
	_this.jsValue.Set("draggable", input)
}

// Spellcheck returning attribute 'spellcheck' with
// type bool (idl: boolean).
func (_this *HTMLElement) Spellcheck() bool {
	var ret bool
	value := _this.jsValue.Get("spellcheck")
	ret = (value).Bool()
	return ret
}

// SetSpellcheck setting attribute 'spellcheck' with
// type bool (idl: boolean).
func (_this *HTMLElement) SetSpellcheck(value bool) {
	input := value
	_this.jsValue.Set("spellcheck", input)
}

// Autocapitalize returning attribute 'autocapitalize' with
// type string (idl: DOMString).
func (_this *HTMLElement) Autocapitalize() string {
	var ret string
	value := _this.jsValue.Get("autocapitalize")
	ret = (value).String()
	return ret
}

// SetAutocapitalize setting attribute 'autocapitalize' with
// type string (idl: DOMString).
func (_this *HTMLElement) SetAutocapitalize(value string) {
	input := value
	_this.jsValue.Set("autocapitalize", input)
}

// InnerText returning attribute 'innerText' with
// type string (idl: DOMString).
func (_this *HTMLElement) InnerText() string {
	var ret string
	value := _this.jsValue.Get("innerText")
	ret = (value).String()
	return ret
}

// SetInnerText setting attribute 'innerText' with
// type string (idl: DOMString).
func (_this *HTMLElement) SetInnerText(value string) {
	input := value
	_this.jsValue.Set("innerText", input)
}

// OffsetParent returning attribute 'offsetParent' with
// type Element (idl: Element).
func (_this *HTMLElement) OffsetParent() *Element {
	var ret *Element
	value := _this.jsValue.Get("offsetParent")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = MakeElementFromJS(value)
	}
	return ret
}

// OffsetTop returning attribute 'offsetTop' with
// type int (idl: long).
func (_this *HTMLElement) OffsetTop() int {
	var ret int
	value := _this.jsValue.Get("offsetTop")
	ret = (value).Int()
	return ret
}

// OffsetLeft returning attribute 'offsetLeft' with
// type int (idl: long).
func (_this *HTMLElement) OffsetLeft() int {
	var ret int
	value := _this.jsValue.Get("offsetLeft")
	ret = (value).Int()
	return ret
}

// OffsetWidth returning attribute 'offsetWidth' with
// type int (idl: long).
func (_this *HTMLElement) OffsetWidth() int {
	var ret int
	value := _this.jsValue.Get("offsetWidth")
	ret = (value).Int()
	return ret
}

// OffsetHeight returning attribute 'offsetHeight' with
// type int (idl: long).
func (_this *HTMLElement) OffsetHeight() int {
	var ret int
	value := _this.jsValue.Get("offsetHeight")
	ret = (value).Int()
	return ret
}

// ContentEditable returning attribute 'contentEditable' with
// type string (idl: DOMString).
func (_this *HTMLElement) ContentEditable() string {
	var ret string
	value := _this.jsValue.Get("contentEditable")
	ret = (value).String()
	return ret
}

// SetContentEditable setting attribute 'contentEditable' with
// type string (idl: DOMString).
func (_this *HTMLElement) SetContentEditable(value string) {
	input := value
	_this.jsValue.Set("contentEditable", input)
}

// EnterKeyHint returning attribute 'enterKeyHint' with
// type string (idl: DOMString).
func (_this *HTMLElement) EnterKeyHint() string {
	var ret string
	value := _this.jsValue.Get("enterKeyHint")
	ret = (value).String()
	return ret
}

// SetEnterKeyHint setting attribute 'enterKeyHint' with
// type string (idl: DOMString).
func (_this *HTMLElement) SetEnterKeyHint(value string) {
	input := value
	_this.jsValue.Set("enterKeyHint", input)
}

// IsContentEditable returning attribute 'isContentEditable' with
// type bool (idl: boolean).
func (_this *HTMLElement) IsContentEditable() bool {
	var ret bool
	value := _this.jsValue.Get("isContentEditable")
	ret = (value).Bool()
	return ret
}

// InputMode returning attribute 'inputMode' with
// type string (idl: DOMString).
func (_this *HTMLElement) InputMode() string {
	var ret string
	value := _this.jsValue.Get("inputMode")
	ret = (value).String()
	return ret
}

// SetInputMode setting attribute 'inputMode' with
// type string (idl: DOMString).
func (_this *HTMLElement) SetInputMode(value string) {
	input := value
	_this.jsValue.Set("inputMode", input)
}

// Nonce returning attribute 'nonce' with
// type string (idl: DOMString).
func (_this *HTMLElement) Nonce() string {
	var ret string
	value := _this.jsValue.Get("nonce")
	ret = (value).String()
	return ret
}

// SetNonce setting attribute 'nonce' with
// type string (idl: DOMString).
func (_this *HTMLElement) SetNonce(value string) {
	input := value
	_this.jsValue.Set("nonce", input)
}

// TabIndex returning attribute 'tabIndex' with
// type int (idl: long).
func (_this *HTMLElement) TabIndex() int {
	var ret int
	value := _this.jsValue.Get("tabIndex")
	ret = (value).Int()
	return ret
}

// SetTabIndex setting attribute 'tabIndex' with
// type int (idl: long).
func (_this *HTMLElement) SetTabIndex(value int) {
	input := value
	_this.jsValue.Set("tabIndex", input)
}

// event attribute: Event
func eventFuncHTMLElement_domcore_Event(listener func(event *Event, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *Event
		value := args[0]
		incoming := value.Get("target")
		ret = MakeEventFromJS(value)
		src := HTMLElementFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddAbort is adding doing AddEventListener for 'Abort' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventAbort(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "abort", cb)
	return cb
}

// SetOnAbort is assigning a function to 'onabort'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnAbort(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onabort", cb)
	return cb
}

// event attribute: MouseEvent
func eventFuncHTMLElement_MouseEvent(listener func(event *MouseEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *MouseEvent
		value := args[0]
		incoming := value.Get("target")
		ret = MouseEventFromJS(value)
		src := HTMLElementFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddAuxclick is adding doing AddEventListener for 'Auxclick' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventAuxclick(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "auxclick", cb)
	return cb
}

// SetOnAuxclick is assigning a function to 'onauxclick'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnAuxclick(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Set("onauxclick", cb)
	return cb
}

// event attribute: FocusEvent
func eventFuncHTMLElement_htmlevent_FocusEvent(listener func(event *FocusEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *FocusEvent
		value := args[0]
		incoming := value.Get("target")
		ret = FocusEventFromJS(value)
		src := HTMLElementFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddBlur is adding doing AddEventListener for 'Blur' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventBlur(listener func(event *FocusEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_htmlevent_FocusEvent(listener)
	_this.jsValue.Call("addEventListener", "blur", cb)
	return cb
}

// SetOnBlur is assigning a function to 'onblur'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnBlur(listener func(event *FocusEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_htmlevent_FocusEvent(listener)
	_this.jsValue.Set("onblur", cb)
	return cb
}

// AddCancel is adding doing AddEventListener for 'Cancel' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventCancel(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "cancel", cb)
	return cb
}

// SetOnCancel is assigning a function to 'oncancel'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnCancel(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("oncancel", cb)
	return cb
}

// AddCanPlay is adding doing AddEventListener for 'CanPlay' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventCanPlay(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "canplay", cb)
	return cb
}

// SetOnCanPlay is assigning a function to 'oncanplay'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnCanPlay(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("oncanplay", cb)
	return cb
}

// AddCanPlayThrough is adding doing AddEventListener for 'CanPlayThrough' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventCanPlayThrough(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "canplaythrough", cb)
	return cb
}

// SetOnCanPlayThrough is assigning a function to 'oncanplaythrough'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnCanPlayThrough(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("oncanplaythrough", cb)
	return cb
}

// AddChange is adding doing AddEventListener for 'Change' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventChange(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "change", cb)
	return cb
}

// SetOnChange is assigning a function to 'onchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnChange(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onchange", cb)
	return cb
}

// AddClick is adding doing AddEventListener for 'Click' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventClick(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "click", cb)
	return cb
}

// SetOnClick is assigning a function to 'onclick'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnClick(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Set("onclick", cb)
	return cb
}

// AddClose is adding doing AddEventListener for 'Close' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventClose(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "close", cb)
	return cb
}

// SetOnClose is assigning a function to 'onclose'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnClose(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onclose", cb)
	return cb
}

// AddCueChange is adding doing AddEventListener for 'CueChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventCueChange(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "cuechange", cb)
	return cb
}

// SetOnCueChange is assigning a function to 'oncuechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnCueChange(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("oncuechange", cb)
	return cb
}

// AddDblClick is adding doing AddEventListener for 'DblClick' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventDblClick(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "dblclick", cb)
	return cb
}

// SetOnDblClick is assigning a function to 'ondblclick'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnDblClick(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Set("ondblclick", cb)
	return cb
}

// AddDurationChange is adding doing AddEventListener for 'DurationChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventDurationChange(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "durationchange", cb)
	return cb
}

// SetOnDurationChange is assigning a function to 'ondurationchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnDurationChange(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("ondurationchange", cb)
	return cb
}

// AddEmptied is adding doing AddEventListener for 'Emptied' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventEmptied(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "emptied", cb)
	return cb
}

// SetOnEmptied is assigning a function to 'onemptied'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnEmptied(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onemptied", cb)
	return cb
}

// AddEnded is adding doing AddEventListener for 'Ended' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventEnded(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "ended", cb)
	return cb
}

// SetOnEnded is assigning a function to 'onended'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnEnded(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onended", cb)
	return cb
}

// AddError is adding doing AddEventListener for 'Error' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventError(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "error", cb)
	return cb
}

// SetOnError is assigning a function to 'onerror'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnError(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onerror", cb)
	return cb
}

// AddFocus is adding doing AddEventListener for 'Focus' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventFocus(listener func(event *FocusEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_htmlevent_FocusEvent(listener)
	_this.jsValue.Call("addEventListener", "focus", cb)
	return cb
}

// SetOnFocus is assigning a function to 'onfocus'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnFocus(listener func(event *FocusEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_htmlevent_FocusEvent(listener)
	_this.jsValue.Set("onfocus", cb)
	return cb
}

// event attribute: PointerEvent
func eventFuncHTMLElement_PointerEvent(listener func(event *PointerEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *PointerEvent
		value := args[0]
		incoming := value.Get("target")
		ret = PointerEventFromJS(value)
		src := HTMLElementFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddGotPointerCapture is adding doing AddEventListener for 'GotPointerCapture' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventGotPointerCapture(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "gotpointercapture", cb)
	return cb
}

// SetOnGotPointerCapture is assigning a function to 'ongotpointercapture'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnGotPointerCapture(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Set("ongotpointercapture", cb)
	return cb
}

// event attribute: InputEvent
func eventFuncHTMLElement_htmlevent_InputEvent(listener func(event *InputEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *InputEvent
		value := args[0]
		incoming := value.Get("target")
		ret = InputEventFromJS(value)
		src := HTMLElementFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddInput is adding doing AddEventListener for 'Input' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventInput(listener func(event *InputEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_htmlevent_InputEvent(listener)
	_this.jsValue.Call("addEventListener", "input", cb)
	return cb
}

// SetOnInput is assigning a function to 'oninput'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnInput(listener func(event *InputEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_htmlevent_InputEvent(listener)
	_this.jsValue.Set("oninput", cb)
	return cb
}

// AddInvalid is adding doing AddEventListener for 'Invalid' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventInvalid(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "invalid", cb)
	return cb
}

// SetOnInvalid is assigning a function to 'oninvalid'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnInvalid(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("oninvalid", cb)
	return cb
}

// event attribute: KeyboardEvent
func eventFuncHTMLElement_KeyboardEvent(listener func(event *KeyboardEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *KeyboardEvent
		value := args[0]
		incoming := value.Get("target")
		ret = KeyboardEventFromJS(value)
		src := HTMLElementFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddKeyDown is adding doing AddEventListener for 'KeyDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventKeyDown(listener func(event *KeyboardEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_KeyboardEvent(listener)
	_this.jsValue.Call("addEventListener", "keydown", cb)
	return cb
}

// SetOnKeyDown is assigning a function to 'onkeydown'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnKeyDown(listener func(event *KeyboardEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_KeyboardEvent(listener)
	_this.jsValue.Set("onkeydown", cb)
	return cb
}

// AddKeyPress is adding doing AddEventListener for 'KeyPress' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventKeyPress(listener func(event *KeyboardEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_KeyboardEvent(listener)
	_this.jsValue.Call("addEventListener", "keypress", cb)
	return cb
}

// SetOnKeyPress is assigning a function to 'onkeypress'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnKeyPress(listener func(event *KeyboardEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_KeyboardEvent(listener)
	_this.jsValue.Set("onkeypress", cb)
	return cb
}

// AddKeyUp is adding doing AddEventListener for 'KeyUp' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventKeyUp(listener func(event *KeyboardEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_KeyboardEvent(listener)
	_this.jsValue.Call("addEventListener", "keyup", cb)
	return cb
}

// SetOnKeyUp is assigning a function to 'onkeyup'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnKeyUp(listener func(event *KeyboardEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_KeyboardEvent(listener)
	_this.jsValue.Set("onkeyup", cb)
	return cb
}

// AddLoad is adding doing AddEventListener for 'Load' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventLoad(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "load", cb)
	return cb
}

// SetOnLoad is assigning a function to 'onload'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnLoad(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onload", cb)
	return cb
}

// AddLoadedData is adding doing AddEventListener for 'LoadedData' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventLoadedData(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "loadeddata", cb)
	return cb
}

// SetOnLoadedData is assigning a function to 'onloadeddata'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnLoadedData(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onloadeddata", cb)
	return cb
}

// AddLoadedMetaData is adding doing AddEventListener for 'LoadedMetaData' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventLoadedMetaData(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "loadedmetadata", cb)
	return cb
}

// SetOnLoadedMetaData is assigning a function to 'onloadedmetadata'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnLoadedMetaData(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onloadedmetadata", cb)
	return cb
}

// AddLoadStart is adding doing AddEventListener for 'LoadStart' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventLoadStart(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "loadstart", cb)
	return cb
}

// SetOnLoadStart is assigning a function to 'onloadstart'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnLoadStart(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onloadstart", cb)
	return cb
}

// AddLostPointerCapture is adding doing AddEventListener for 'LostPointerCapture' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventLostPointerCapture(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "lostpointercapture", cb)
	return cb
}

// SetOnLostPointerCapture is assigning a function to 'onlostpointercapture'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnLostPointerCapture(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Set("onlostpointercapture", cb)
	return cb
}

// AddMouseDown is adding doing AddEventListener for 'MouseDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventMouseDown(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mousedown", cb)
	return cb
}

// SetOnMouseDown is assigning a function to 'onmousedown'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnMouseDown(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Set("onmousedown", cb)
	return cb
}

// AddMouseEnter is adding doing AddEventListener for 'MouseEnter' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventMouseEnter(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseenter", cb)
	return cb
}

// SetOnMouseEnter is assigning a function to 'onmouseenter'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnMouseEnter(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Set("onmouseenter", cb)
	return cb
}

// AddMouseLeave is adding doing AddEventListener for 'MouseLeave' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventMouseLeave(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseleave", cb)
	return cb
}

// SetOnMouseLeave is assigning a function to 'onmouseleave'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnMouseLeave(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Set("onmouseleave", cb)
	return cb
}

// AddMouseMove is adding doing AddEventListener for 'MouseMove' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventMouseMove(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mousemove", cb)
	return cb
}

// SetOnMouseMove is assigning a function to 'onmousemove'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnMouseMove(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Set("onmousemove", cb)
	return cb
}

// AddMouseOut is adding doing AddEventListener for 'MouseOut' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventMouseOut(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseout", cb)
	return cb
}

// SetOnMouseOut is assigning a function to 'onmouseout'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnMouseOut(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Set("onmouseout", cb)
	return cb
}

// AddMouseOver is adding doing AddEventListener for 'MouseOver' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventMouseOver(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseover", cb)
	return cb
}

// SetOnMouseOver is assigning a function to 'onmouseover'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnMouseOver(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Set("onmouseover", cb)
	return cb
}

// AddMouseUp is adding doing AddEventListener for 'MouseUp' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventMouseUp(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Call("addEventListener", "mouseup", cb)
	return cb
}

// SetOnMouseUp is assigning a function to 'onmouseup'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnMouseUp(listener func(event *MouseEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_MouseEvent(listener)
	_this.jsValue.Set("onmouseup", cb)
	return cb
}

// AddPause is adding doing AddEventListener for 'Pause' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventPause(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "pause", cb)
	return cb
}

// SetOnPause is assigning a function to 'onpause'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnPause(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onpause", cb)
	return cb
}

// AddPlay is adding doing AddEventListener for 'Play' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventPlay(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "play", cb)
	return cb
}

// SetOnPlay is assigning a function to 'onplay'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnPlay(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onplay", cb)
	return cb
}

// AddPlaying is adding doing AddEventListener for 'Playing' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventPlaying(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "playing", cb)
	return cb
}

// SetOnPlaying is assigning a function to 'onplaying'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnPlaying(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onplaying", cb)
	return cb
}

// AddPointerCancel is adding doing AddEventListener for 'PointerCancel' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventPointerCancel(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointercancel", cb)
	return cb
}

// SetOnPointerCancel is assigning a function to 'onpointercancel'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnPointerCancel(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Set("onpointercancel", cb)
	return cb
}

// AddPointerDown is adding doing AddEventListener for 'PointerDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventPointerDown(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerdown", cb)
	return cb
}

// SetOnPointerDown is assigning a function to 'onpointerdown'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnPointerDown(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Set("onpointerdown", cb)
	return cb
}

// AddPointerEnter is adding doing AddEventListener for 'PointerEnter' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventPointerEnter(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerenter", cb)
	return cb
}

// SetOnPointerEnter is assigning a function to 'onpointerenter'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnPointerEnter(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Set("onpointerenter", cb)
	return cb
}

// AddPointerLeave is adding doing AddEventListener for 'PointerLeave' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventPointerLeave(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerleave", cb)
	return cb
}

// SetOnPointerLeave is assigning a function to 'onpointerleave'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnPointerLeave(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Set("onpointerleave", cb)
	return cb
}

// AddPointerMove is adding doing AddEventListener for 'PointerMove' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventPointerMove(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointermove", cb)
	return cb
}

// SetOnPointerMove is assigning a function to 'onpointermove'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnPointerMove(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Set("onpointermove", cb)
	return cb
}

// AddPointerOut is adding doing AddEventListener for 'PointerOut' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventPointerOut(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerout", cb)
	return cb
}

// SetOnPointerOut is assigning a function to 'onpointerout'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnPointerOut(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Set("onpointerout", cb)
	return cb
}

// AddPointerOver is adding doing AddEventListener for 'PointerOver' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventPointerOver(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerover", cb)
	return cb
}

// SetOnPointerOver is assigning a function to 'onpointerover'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnPointerOver(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Set("onpointerover", cb)
	return cb
}

// AddPointerUp is adding doing AddEventListener for 'PointerUp' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventPointerUp(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Call("addEventListener", "pointerup", cb)
	return cb
}

// SetOnPointerUp is assigning a function to 'onpointerup'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnPointerUp(listener func(event *PointerEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_PointerEvent(listener)
	_this.jsValue.Set("onpointerup", cb)
	return cb
}

// AddRateChange is adding doing AddEventListener for 'RateChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventRateChange(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "ratechange", cb)
	return cb
}

// SetOnRateChange is assigning a function to 'onratechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnRateChange(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onratechange", cb)
	return cb
}

// AddReset is adding doing AddEventListener for 'Reset' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventReset(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "reset", cb)
	return cb
}

// SetOnReset is assigning a function to 'onreset'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnReset(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onreset", cb)
	return cb
}

// event attribute: UIEvent
func eventFuncHTMLElement_UIEvent(listener func(event *UIEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *UIEvent
		value := args[0]
		incoming := value.Get("target")
		ret = UIEventFromJS(value)
		src := HTMLElementFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddResize is adding doing AddEventListener for 'Resize' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventResize(listener func(event *UIEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_UIEvent(listener)
	_this.jsValue.Call("addEventListener", "resize", cb)
	return cb
}

// SetOnResize is assigning a function to 'onresize'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnResize(listener func(event *UIEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_UIEvent(listener)
	_this.jsValue.Set("onresize", cb)
	return cb
}

// AddScroll is adding doing AddEventListener for 'Scroll' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventScroll(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "scroll", cb)
	return cb
}

// SetOnScroll is assigning a function to 'onscroll'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnScroll(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onscroll", cb)
	return cb
}

// AddSeeked is adding doing AddEventListener for 'Seeked' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventSeeked(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "seeked", cb)
	return cb
}

// SetOnSeeked is assigning a function to 'onseeked'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnSeeked(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onseeked", cb)
	return cb
}

// AddSeeking is adding doing AddEventListener for 'Seeking' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventSeeking(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "seeking", cb)
	return cb
}

// SetOnSeeking is assigning a function to 'onseeking'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnSeeking(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onseeking", cb)
	return cb
}

// AddSelect is adding doing AddEventListener for 'Select' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventSelect(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "select", cb)
	return cb
}

// SetOnSelect is assigning a function to 'onselect'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnSelect(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onselect", cb)
	return cb
}

// AddSelectionChange is adding doing AddEventListener for 'SelectionChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventSelectionChange(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "selectionchange", cb)
	return cb
}

// SetOnSelectionChange is assigning a function to 'onselectionchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnSelectionChange(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onselectionchange", cb)
	return cb
}

// AddSelectStart is adding doing AddEventListener for 'SelectStart' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventSelectStart(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "selectstart", cb)
	return cb
}

// SetOnSelectStart is assigning a function to 'onselectstart'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnSelectStart(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onselectstart", cb)
	return cb
}

// AddStalled is adding doing AddEventListener for 'Stalled' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventStalled(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "stalled", cb)
	return cb
}

// SetOnStalled is assigning a function to 'onstalled'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnStalled(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onstalled", cb)
	return cb
}

// AddSubmit is adding doing AddEventListener for 'Submit' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventSubmit(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "submit", cb)
	return cb
}

// SetOnSubmit is assigning a function to 'onsubmit'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnSubmit(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onsubmit", cb)
	return cb
}

// AddSuspend is adding doing AddEventListener for 'Suspend' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventSuspend(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "suspend", cb)
	return cb
}

// SetOnSuspend is assigning a function to 'onsuspend'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnSuspend(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onsuspend", cb)
	return cb
}

// AddTimeUpdate is adding doing AddEventListener for 'TimeUpdate' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventTimeUpdate(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "timeupdate", cb)
	return cb
}

// SetOnTimeUpdate is assigning a function to 'ontimeupdate'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnTimeUpdate(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("ontimeupdate", cb)
	return cb
}

// AddToggle is adding doing AddEventListener for 'Toggle' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventToggle(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "toggle", cb)
	return cb
}

// SetOnToggle is assigning a function to 'ontoggle'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnToggle(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("ontoggle", cb)
	return cb
}

// AddVolumeChange is adding doing AddEventListener for 'VolumeChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventVolumeChange(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "volumechange", cb)
	return cb
}

// SetOnVolumeChange is assigning a function to 'onvolumechange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnVolumeChange(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onvolumechange", cb)
	return cb
}

// AddWaiting is adding doing AddEventListener for 'Waiting' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventWaiting(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Call("addEventListener", "waiting", cb)
	return cb
}

// SetOnWaiting is assigning a function to 'onwaiting'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnWaiting(listener func(event *Event, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_domcore_Event(listener)
	_this.jsValue.Set("onwaiting", cb)
	return cb
}

// event attribute: WheelEvent
func eventFuncHTMLElement_htmlevent_WheelEvent(listener func(event *WheelEvent, target *HTMLElement)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *WheelEvent
		value := args[0]
		incoming := value.Get("target")
		ret = WheelEventFromJS(value)
		src := HTMLElementFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddWheel is adding doing AddEventListener for 'Wheel' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) AddEventWheel(listener func(event *WheelEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_htmlevent_WheelEvent(listener)
	_this.jsValue.Call("addEventListener", "wheel", cb)
	return cb
}

// SetOnWheel is assigning a function to 'onwheel'. This
// This method is returning allocated javascript function that need to be released.
func (_this *HTMLElement) SetOnWheel(listener func(event *WheelEvent, currentTarget *HTMLElement)) js.Func {
	cb := eventFuncHTMLElement_htmlevent_WheelEvent(listener)
	_this.jsValue.Set("onwheel", cb)
	return cb
}

func (_this *HTMLElement) Click() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("click", _args[0:_end]...)
}

func (_this *HTMLElement) Focus() {
	var (
		_args [1]interface{}
		_end  int
	)
	_this.jsValue.Call("focus", _args[0:_end]...)
}

func (_this *HTMLElement) Blur() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("blur", _args[0:_end]...)
}

/****************************************************************************
* HTMLHeadElement
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLHeadElement
type HTMLHeadElement struct {
	HTMLElement
}

// HTMLHeadElementFromJS is casting a js.Value into HTMLHeadElement.
func HTMLHeadElementFromJS(value js.Value) *HTMLHeadElement {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &HTMLHeadElement{}
	ret.jsValue = value
	return ret
}
