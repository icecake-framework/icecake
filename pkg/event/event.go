package event

import (
	"net/url"

	syscalljs "syscall/js"

	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/js"
)

/*********************************************************************************
 * HashChangeEvent
 */

type GENERIC_EVENT string

const (
	GENERIC_ONCANCEL            GENERIC_EVENT = "cancel"
	GENERIC_ONCHANGE            GENERIC_EVENT = "change"
	GENERIC_ONCLOSE             GENERIC_EVENT = "close"
	GENERIC_ONCUECHANGE         GENERIC_EVENT = "cuechange"
	GENERIC_ONINVALID           GENERIC_EVENT = "invalid"
	GENERIC_ONLOAD              GENERIC_EVENT = "load"
	GENERIC_ONRESET             GENERIC_EVENT = "reset"
	GENERIC_ONSCROLL            GENERIC_EVENT = "scroll"
	GENERIC_ONSELECT            GENERIC_EVENT = "select"
	GENERIC_ONSELECTIONCHANGE   GENERIC_EVENT = "selectionchange"
	GENERIC_ONSELECTSTART       GENERIC_EVENT = "selectstart"
	GENERIC_ONSUBMIT            GENERIC_EVENT = "submit"
	GENERIC_ONTOGGLE            GENERIC_EVENT = "toggle"
	GENERIC_ONPOINTERLOCKCHANGE GENERIC_EVENT = "pointerlockchange"
	GENERIC_ONPOINTERLOCKERROR  GENERIC_EVENT = "pointerlockerror"
)

const (
	GENERIC_WIN_ONAFTERPRINT      GENERIC_EVENT = "afterprint"
	GENERIC_WIN_ONAPPINSTALLED    GENERIC_EVENT = "appinstalled"
	GENERIC_WIN_BEFOREPRINT       GENERIC_EVENT = "beforeprint"
	GENERIC_WIN_LANGUAGECHANGE    GENERIC_EVENT = "languagechange"
	GENERIC_WIN_OFFLINE           GENERIC_EVENT = "offline"
	GENERIC_WIN_ONLINE            GENERIC_EVENT = "online"
	GENERIC_WIN_ORIENTATIONCHANGE GENERIC_EVENT = "orientationchange"
)

const (
	GENERIC_MEDIA_ONABORT          GENERIC_EVENT = "abort"          // Fired when the resource was not fully loaded, but not as the result of an error.
	GENERIC_MEDIA_ONCANPLAY        GENERIC_EVENT = "canplay"        // Fired when the user agent can play the media, but estimates that not enough data has been loaded to play the media up to its end without having to stop for further buffering of content.
	GENERIC_MEDIA_ONCANPLAYTHROUGH GENERIC_EVENT = "canplaythrough" // Fired when the user agent can play the media, and estimates that enough data has been loaded to play the media up to its end without having to stop for further buffering of content.
	GENERIC_MEDIA_ONDURATIONCHANGE GENERIC_EVENT = "durationchange" // Fired when the duration property has been updated.
	GENERIC_MEDIA_ONEMPTIED        GENERIC_EVENT = "emptied"        // Fired when the media has become empty; for example, when the media has already been loaded (or partially loaded), and the HTMLMediaElement.load() method is called to reload it.
	GENERIC_MEDIA_ONENDED          GENERIC_EVENT = "ended"          // Fired when playback stops when end of the media (<audio> or <video>) is reached or because no further data is available.
	GENERIC_MEDIA_ONERROR          GENERIC_EVENT = "error"          // Fired when the resource could not be loaded due to an error.
	GENERIC_MEDIA_ONLOADEDDATA     GENERIC_EVENT = "loadeddata"     // Fired when the first frame of the media has finished loading.
	GENERIC_MEDIA_ONLOADEDMETADATA GENERIC_EVENT = "loadedmetadata" // Fired when the metadata has been loaded.
	GENERIC_MEDIA_ONLOADSTART      GENERIC_EVENT = "loadstart"      // Fired when the browser has started to load a resource.
	GENERIC_MEDIA_ONPAUSE          GENERIC_EVENT = "pause"          // Fired when a request to pause play is handled and the activity has entered its paused state, most commonly occurring when the media's HTMLMediaElement.pause() method is called.
	GENERIC_MEDIA_ONPLAY           GENERIC_EVENT = "play"           // Fired when the paused property is changed from true to false, as a result of the HTMLMediaElement.play() method, or the autoplay attribute.
	GENERIC_MEDIA_ONPLAYING        GENERIC_EVENT = "playing"        // Fired when playback is ready to start after having been paused or delayed due to lack of data.
	GENERIC_MEDIA_ONRATECHANGE     GENERIC_EVENT = "ratechange"     // Fired when the playback rate has changed.
	GENERIC_MEDIA_ONSEEKED         GENERIC_EVENT = "seeked"         // Fired when a seek operation completes.
	GENERIC_MEDIA_ONSEEKING        GENERIC_EVENT = "seeking"        // Fired when a seek operation begins.
	GENERIC_MEDIA_ONSTALLED        GENERIC_EVENT = "stalled"        // Fired when the user agent is trying to fetch media data, but data is unexpectedly not forthcoming.
	GENERIC_MEDIA_ONSUSPEND        GENERIC_EVENT = "suspend"        // Fired when the media data loading has been suspended.
	GENERIC_MEDIA_ONTIMEUPDATE     GENERIC_EVENT = "timeupdate"     // Fired when the time indicated by the currentTime property has been updated.
	GENERIC_MEDIA_ONVOLUMECHANGE   GENERIC_EVENT = "volumechange"   // Fired when the volume has changed.
	GENERIC_MEDIA_ONWAITING        GENERIC_EVENT = "waiting"        // Fired when playback has stopped because of a temporary lack of data.
)

type FULLSCREEN_EVENT string

const (
	FULLSCREEN_ONCHANGE FULLSCREEN_EVENT = "fullscreenchange"
	FULLSCREEN_ONERROR  FULLSCREEN_EVENT = "fullscreenerror"
)

/******************************************************************************
* event Handler
*******************************************************************************/

type Handler struct {
	eventtype string // 'onclick'...
	jsHandler syscalljs.Func
	release   func()
}

func NewEventHandler(_eventtype string, _jsHandler syscalljs.Func, _release func()) *Handler {
	evt := &Handler{eventtype: _eventtype, jsHandler: _jsHandler, release: _release}
	return evt
}

func (h Handler) Release() {
	if h.release != nil {
		h.release()
	}
}

/******************************************************************************
* Event
******************************************************************************/

// Event represents an event which takes place in the DOM.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event
type Event struct {
	js.JSValue
}

// CastEvent is casting a js.Value into Event.
func CastEvent(_jsv js.JSValueProvider) *Event {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting Event failed")
		return nil
	}
	ret := new(Event)
	ret.JSValue = _jsv.Value()
	return ret
}

/******************************************************************************
* Event's properties
******************************************************************************/

// Type: there are many types of events, some of which use other interfaces based on the main Event interface.
// Event itself contains the properties and methods which are common to all events.
//
// It is set when the event is constructed.
// Commonly used to refer to the specific event, such as click, load, or error
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/type
func (_evt *Event) Type() string {
	return _evt.Get("type").String()
}

// Target: a reference to the object onto which the event was dispatched.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/target
func (_evt *Event) Target() js.JSValue {
	target := _evt.Get("target")
	return target
}

// CurrentTarget always refers to the element to which the event handler has been attached,
// as opposed to Event.target, which identifies the element on which the event occurred and which may be its descendant.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/currentTarget
func (_evt *Event) CurrentTarget() js.JSValue {
	target := _evt.Get("currentTarget")
	return target //CastEventTarget(target)
}

/*********************************************************************************
 * HashChangeEvent
 */

// fire when the fragment identifier of the URL has changed.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HashChangeEvent
type HashChangeEvent struct {
	Event
}

// CastHashChangeEvent is casting a js.Value into HashChangeEvent.
func CastHashChangeEvent(_jsv js.JSValueProvider) *HashChangeEvent {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting HashChangeEvent failed")
		return nil
	}
	ret := new(HashChangeEvent)
	ret.JSValue = _jsv.Value()
	return ret
}

// OldURL returning attribute 'oldURL' with
// type string (idl: USVString).
func (_evt *HashChangeEvent) OldURL() *url.URL {
	ref := _evt.Get("oldURL").String()
	u, _ := url.Parse(ref)
	return u
}

// NewURL returning attribute 'newURL' with
// type string (idl: USVString).
func (_evt *HashChangeEvent) NewURL() *url.URL {
	ref := _evt.Get("newURL").String()
	u, _ := url.Parse(ref)
	return u
}

/*********************************************************************************
 * PageTransitionEvent
 */

type PAGETRANSITION_EVENT string

const (
	PAGETRANSITION_ONPAGEHIDE PAGETRANSITION_EVENT = "pagehide"
	PAGETRANSITION_ONPAGESHOW PAGETRANSITION_EVENT = "pageshow"
)

// The PageTransitionEvent event object is available inside handler functions for the pageshow and pagehide events,
// fired when a document is being loaded or unloaded.
//
// https://developer.mozilla.org/en-US/docs/Web/API/PageTransitionEvent
type PageTransitionEvent struct {
	Event
}

// CastPageTransitionEvent is casting a js.Value into PageTransitionEvent.
func CastPageTransitionEvent(_jsv js.JSValueProvider) *PageTransitionEvent {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting PageTransitionEvent failed")
		return nil
	}
	ret := new(PageTransitionEvent)
	ret.JSValue = _jsv.Value()
	return ret
}

// Persisted returning attribute 'persisted' with
func (_evt *PageTransitionEvent) Persisted() bool {
	return _evt.Get("persisted").Bool()
}

/*********************************************************************************
 * BeforeUnloadEvent
 */

// fired when the window, the document and its resources are about to be unloaded.
//
// https://developer.mozilla.org/en-US/docs/Web/API/BeforeUnloadEvent
type BeforeUnloadEvent struct {
	Event
}

// CastBeforeUnloadEvent is casting a js.Value into BeforeUnloadEvent.
func CastBeforeUnloadEvent(_jsv js.JSValueProvider) *BeforeUnloadEvent {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting BeforeUnloadEvent failed")
		return nil
	}
	ret := new(BeforeUnloadEvent)
	ret.JSValue = _jsv.Value()
	return ret
}

// ReturnValue returning attribute 'returnValue' with
func (_evt *BeforeUnloadEvent) ReturnValue() string {
	return _evt.Get("returnValue").String()
}

// SetReturnValue setting attribute 'returnValue' with
func (_evt *BeforeUnloadEvent) SetReturnValue(value string) {
	_evt.Set("returnValue", value)
}

/*********************************************************************************
 * UIEvent
 */

// UIEvent represents simple user interface events.
//
// https://developer.mozilla.org/en-US/docs/Web/API/UIEvent
type UIEvent struct {
	Event
}

// CastUIEvent is casting a js.Value into UIEvent.
func CastUIEvent(_jsv js.JSValueProvider) *UIEvent {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting UIEvent failed")
		return nil
	}
	ret := new(UIEvent)
	ret.JSValue = _jsv.Value()
	return ret
}

// // View Returns a WindowProxy that contains the view that generated the event.
// func (_evt *UIEvent) View() Window {
// 	win := _evt.Get("view")
// 	return CastWindow(win)
// }

// Detail when non-zero, provides the current (or next, depending on the event) click count.
//
// https://developer.mozilla.org/en-US/docs/Web/API/UIEvent/detail
func (_evt *UIEvent) Detail() int {
	return _evt.Get("detail").Int()
}

/*********************************************************************************
 * MouseEvent
 */

type MOUSE_EVENT string

const (
	MOUSE_ONCLICK       MOUSE_EVENT = "click"
	MOUSE_ONAUXCLICK    MOUSE_EVENT = "auxclick"
	MOUSE_ONDBLCLICK    MOUSE_EVENT = "dblclick"
	MOUSE_ONMOUSEDOWN   MOUSE_EVENT = "mousedown"
	MOUSE_ONMOUSEENTER  MOUSE_EVENT = "mouseenter"
	MOUSE_ONMOUSELEAVE  MOUSE_EVENT = "mouseleave"
	MOUSE_ONMOUSEMOVE   MOUSE_EVENT = "mousemove"
	MOUSE_ONMOUSEOUT    MOUSE_EVENT = "mouseout"
	MOUSE_ONMOUSEUP     MOUSE_EVENT = "mouseup"
	MOUSE_ONMOUSEOVER   MOUSE_EVENT = "mouseover"
	MOUSE_ONCONTEXTMENU MOUSE_EVENT = "contextmenu"
)

type MouseEvent struct {
	UIEvent
}

// CastMouseEvent is casting a js.Value into MouseEvent.
func CastMouseEvent(_jsv js.JSValueProvider) *MouseEvent {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting MouseEvent failed")
		return nil
	}
	ret := new(MouseEvent)
	ret.JSValue = _jsv.Value()
	return ret
}

// CtrlKey returning attribute 'ctrlKey' with
func (_this *MouseEvent) CtrlKey() bool {
	return _this.Get("ctrlKey").Bool()
}

// ShiftKey returning attribute 'shiftKey' with
func (_this *MouseEvent) ShiftKey() bool {
	return _this.Get("shiftKey").Bool()
}

// AltKey returning attribute 'altKey' with
func (_this *MouseEvent) AltKey() bool {
	return _this.Get("altKey").Bool()
}

// MetaKey returning attribute 'metaKey' with
func (_this *MouseEvent) MetaKey() bool {
	return _this.Get("metaKey").Bool()
}

// Button returning attribute 'button' with
func (_this *MouseEvent) Button() int {
	return _this.Get("button").Int()
}

// Buttons returning attribute 'buttons' with
func (_this *MouseEvent) Buttons() int {
	return _this.Get("buttons").Int()
}

// RelatedTarget returning attribute 'relatedTarget' with
func (_this *MouseEvent) RelatedTarget() js.JSValue {
	value := _this.Get("relatedTarget")
	return value // CastEventTarget(value)
}

// ScreenX returning attribute 'screenX' with
func (_this *MouseEvent) ScreenX() float64 {
	return _this.Get("screenX").Float()
}

// ScreenY returning attribute 'screenY' with
func (_this *MouseEvent) ScreenY() float64 {
	return _this.Get("screenY").Float()
}

// PageX returning attribute 'pageX' with
func (_this *MouseEvent) PageX() float64 {
	return _this.Get("pageX").Float()
}

// PageY returning attribute 'pageY' with
func (_this *MouseEvent) PageY() float64 {
	return _this.Get("pageY").Float()
}

// ClientX returning attribute 'clientX' with
func (_this *MouseEvent) ClientX() float64 {
	return _this.Get("clientX").Float()
}

// ClientY returning attribute 'clientY' with
func (_this *MouseEvent) ClientY() float64 {
	return _this.Get("clientY").Float()
}

// X returning attribute 'x' with
func (_this *MouseEvent) X() float64 {
	return _this.Get("x").Float()
}

// Y returning attribute 'y' with
func (_this *MouseEvent) Y() float64 {
	return _this.Get("y").Float()
}

// OffsetX returning attribute 'offsetX' with
func (_this *MouseEvent) OffsetX() float64 {
	return _this.Get("offsetX").Float()
}

// OffsetY returning attribute 'offsetY' with
func (_this *MouseEvent) OffsetY() float64 {
	return _this.Get("offsetY").Float()
}

// MovementX returning attribute 'movementX' with
func (_this *MouseEvent) MovementX() int {
	return _this.Get("movementX").Int()
}

// MovementY returning attribute 'movementY' with
func (_this *MouseEvent) MovementY() int {
	return _this.Get("movementY").Int()
}

func (_this *MouseEvent) GetModifierState(keyArg string) bool {
	return _this.Call("getModifierState", keyArg).Bool()
}

/*********************************************************************************
* WheelEvent
 */

type WHEEL_DELTA_MODE uint

const (
	WHEEL_DELTA_PIXEL WHEEL_DELTA_MODE = 0x00
	WHEEL_DELTA_LINE  WHEEL_DELTA_MODE = 0x01
	WHEEL_DELTA_PAGE  WHEEL_DELTA_MODE = 0x02
)

type WheelEvent struct {
	MouseEvent
}

// CastWheelEvent is casting a js.Value into WheelEvent.
func CastWheelEvent(_jsv js.JSValueProvider) *WheelEvent {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting WheelEvent failed")
		return nil
	}
	ret := &WheelEvent{}
	ret.JSValue = _jsv.Value()
	return ret
}

// DeltaX returning attribute 'deltaX' with
func (_this *WheelEvent) DeltaX() float64 {
	return _this.Get("deltaX").Float()
}

// DeltaY returning attribute 'deltaY' with
func (_this *WheelEvent) DeltaY() float64 {
	return _this.Get("deltaY").Float()
}

// DeltaZ returning attribute 'deltaZ' with
func (_this *WheelEvent) DeltaZ() float64 {
	return _this.Get("deltaZ").Float()
}

// DeltaMode returning attribute 'deltaMode' with
func (_this *WheelEvent) DeltaMode() WHEEL_DELTA_MODE {
	return WHEEL_DELTA_MODE(_this.Get("deltaMode").Int())
}

/**********************************************************************************
 * class: FocusEvent
 */

type FOCUS_EVENT string

const (
	FOCUS_ONBLUR  FOCUS_EVENT = "blur"
	FOCUS_ONFOCUS FOCUS_EVENT = "focus"
)

type FocusEvent struct {
	UIEvent
}

// NewFocusEventFromJS is casting a js.Value into FocusEvent.
func CastFocusEvent(_jsv js.JSValueProvider) *FocusEvent {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting FocusEvent failed")
		return nil
	}
	ret := new(FocusEvent)
	ret.JSValue = _jsv.Value()
	return ret
}

// RelatedTarget returning attribute 'relatedTarget'
func (_evt *FocusEvent) RelatedTarget() js.JSValue {
	value := _evt.Get("relatedTarget")
	return value // CastEventTarget(value)
}

/**********************************************************************************
 * PointerEvent
 */

type POINTER_EVENT string

const (
	POINTER_ONGOTPOINTERCAPTURE  POINTER_EVENT = "gotpointercapture"
	POINTER_ONLOSTPOINTERCAPTURE POINTER_EVENT = "lostpointercapture"
	POINTER_ONPOINTERCANCEL      POINTER_EVENT = "pointercancel"
	POINTER_ONPOINTERDOWN        POINTER_EVENT = "pointerdown"
	POINTER_ONPOINTERCENTER      POINTER_EVENT = "pointerenter"
	POINTER_ONPOINTERLEAVE       POINTER_EVENT = "pointerleave"
	POINTER_ONPOINTERMOVE        POINTER_EVENT = "pointermove"
	POINTER_ONPOINTEROUT         POINTER_EVENT = "pointerout"
	POINTER_ONPOINTEROVER        POINTER_EVENT = "pointerover"
	POINTER_ONPOINTERUP          POINTER_EVENT = "pointerup"
)

type PointerEvent struct {
	MouseEvent
}

// CastPointerEvent is casting a js.Value into PointerEvent.
func CastPointerEvent(_jsv js.JSValueProvider) *PointerEvent {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting PointerEvent failed")
		return nil
	}
	ret := new(PointerEvent)
	ret.JSValue = _jsv.Value()
	return ret
}

// PointerId returning attribute 'pointerId' with
func (_this *PointerEvent) PointerId() int {
	return _this.Get("pointerId").Int()
}

// Width returning attribute 'width' with
func (_this *PointerEvent) Width() float64 {
	return _this.Get("width").Float()
}

// Height returning attribute 'height' with
func (_this *PointerEvent) Height() float64 {
	return _this.Get("height").Float()
}

// Pressure returning attribute 'pressure' with
func (_this *PointerEvent) Pressure() float64 {
	return _this.Get("pressure").Float()
}

// TangentialPressure returning attribute 'tangentialPressure' with
func (_this *PointerEvent) TangentialPressure() float64 {
	return _this.Get("tangentialPressure").Float()
}

// TiltX returning attribute 'tiltX' with
func (_this *PointerEvent) TiltX() int {
	return _this.Get("tiltX").Int()
}

// TiltY returning attribute 'tiltY' with
func (_this *PointerEvent) TiltY() int {
	return _this.Get("tiltY").Int()
}

// Twist returning attribute 'twist' with
func (_this *PointerEvent) Twist() int {
	return _this.Get("twist").Int()
}

// PointerType returning attribute 'pointerType' with
func (_this *PointerEvent) PointerType() string {
	return _this.Get("pointerType").String()
}

// IsPrimary returning attribute 'isPrimary' with
func (_this *PointerEvent) IsPrimary() bool {
	return _this.Get("isPrimary").Bool()
}

/**********************************************************************************
 * InputEvent
 */

// represents an event notifying the user of editable content changes.
type INPUT_EVENT string

const (
	INPUT_ONINPUT       INPUT_EVENT = "input"       // Fired when the value of an <input>, <select>, or <textarea> element has been changed.
	INPUT_ONCHANGE      INPUT_EVENT = "change"      // Fired when the value of an <input>, <select>, or <textarea> element has been changed and committed by the user.
	INPUT_ONBEFOREINPUT INPUT_EVENT = "beforeinput" // Fired when the value of an <input>, <select>, or <textarea> element is about to be modified.
)

type InputEvent struct {
	UIEvent
}

// CastInputEvent is casting a js.Value into InputEvent.
func CastInputEvent(_jsv js.JSValueProvider) *InputEvent {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting InputEvent failed")
		return nil
	}
	ret := new(InputEvent)
	ret.JSValue = _jsv.Value()
	return ret
}

// Data returning attribute 'data' with
func (_this *InputEvent) Data() string {
	return _this.Get("data").String()
}

// IsComposing returning attribute 'isComposing' with
func (_this *InputEvent) IsComposing() bool {
	return _this.Get("isComposing").Bool()
}

// InputType returning attribute 'inputType' with
func (_this *InputEvent) InputType() string {
	return _this.Get("inputType").String()
}

/**********************************************************************************
 * KeyboardEvent
 */

type KEYBOARD_EVENT string

const (
	KEYBOARD_ONKEYDOWN  KEYBOARD_EVENT = "keydown"
	KEYBOARD_ONKEYPRESS KEYBOARD_EVENT = "keypress"
	KEYBOARD_ONKEYUP    KEYBOARD_EVENT = "keyup"
)

type KEYBOARD_LOCATION uint

const (
	KEYBOARD_LOC_STANDARD KEYBOARD_LOCATION = 0x00
	KEYBOARD_LOC_LEFT     KEYBOARD_LOCATION = 0x01
	KEYBOARD_LOC_RIGHT    KEYBOARD_LOCATION = 0x02
	KEYBOARD_LOC_NUMPAD   KEYBOARD_LOCATION = 0x03
)

type KeyboardEvent struct {
	UIEvent
}

// CastKeyboardEvent is casting a js.Value into KeyboardEvent.
func CastKeyboardEvent(_jsv js.JSValueProvider) *KeyboardEvent {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting KeyboardEvent failed")
		return nil
	}
	ret := new(KeyboardEvent)
	ret.JSValue = _jsv.Value()
	return ret
}

// Key returning attribute 'key' with
func (_this *KeyboardEvent) Key() string {
	return _this.Get("key").String()
}

// Code returning attribute 'code' with
func (_this *KeyboardEvent) Code() string {
	return _this.Get("code").String()
}

// Location returning attribute 'location' with
func (_this *KeyboardEvent) Location() KEYBOARD_LOCATION {
	return KEYBOARD_LOCATION(_this.Get("location").Int())
}

// CtrlKey returning attribute 'ctrlKey' with
func (_this *KeyboardEvent) CtrlKey() bool {
	return _this.Get("ctrlKey").Bool()
}

// ShiftKey returning attribute 'shiftKey' with
func (_this *KeyboardEvent) ShiftKey() bool {
	return _this.Get("shiftKey").Bool()
}

// AltKey returning attribute 'altKey' with
func (_this *KeyboardEvent) AltKey() bool {
	return _this.Get("altKey").Bool()
}

// MetaKey returning attribute 'metaKey' with
func (_this *KeyboardEvent) MetaKey() bool {
	return _this.Get("metaKey").Bool()
}

// Repeat returning attribute 'repeat' with
func (_this *KeyboardEvent) Repeat() bool {
	return _this.Get("repeat").Bool()
}

// IsComposing returning attribute 'isComposing' with
func (_this *KeyboardEvent) IsComposing() bool {
	return _this.Get("isComposing").Bool()
}

// CharCode returning attribute 'charCode' with
func (_this *KeyboardEvent) CharCode() uint {
	return uint(_this.Get("charCode").Int())
}

// KeyCode returning attribute 'keyCode' with
func (_this *KeyboardEvent) KeyCode() uint {
	return uint(_this.Get("keyCode").Int())
}

// returns the current state of the specified modifier key: true if the modifier is active (that is the modifier key is pressed or locked), otherwise, false.
//
// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/getModifierState
func (_this *KeyboardEvent) GetModifierState(keyArg string) bool {
	return _this.Call("getModifierState", keyArg).Bool()
}
