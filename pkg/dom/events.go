package dom

import (
	"net/url"
	"syscall/js"
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
* Event
******************************************************************************/

// Event represents an event which takes place in the DOM.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event
type Event struct {
	jsValue js.Value
}

// CastEvent is casting a js.Value into Event.
func CastEvent(value js.Value) *Event {
	if value.Type() != js.TypeObject {
		ICKError("casting Event failed")
		return nil
	}
	ret := new(Event)
	ret.jsValue = value
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
	return _evt.jsValue.Get("type").String()
}

// Target: a reference to the object onto which the event was dispatched.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/target
func (_evt *Event) Target() *EventTarget {
	target := _evt.jsValue.Get("target")
	return CastEventTarget(target)
}

// CurrentTarget always refers to the element to which the event handler has been attached,
// as opposed to Event.target, which identifies the element on which the event occurred and which may be its descendant.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/currentTarget
func (_evt *Event) CurrentTarget() *EventTarget {
	target := _evt.jsValue.Get("currentTarget")
	return CastEventTarget(target)
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
func CastHashChangeEvent(value js.Value) *HashChangeEvent {
	if value.Type() != js.TypeObject {
		ICKError("casting HashChangeEvent failed")
		return nil
	}
	ret := new(HashChangeEvent)
	ret.jsValue = value
	return ret
}

// OldURL returning attribute 'oldURL' with
// type string (idl: USVString).
func (_evt *HashChangeEvent) OldURL() *url.URL {
	ref := _evt.jsValue.Get("oldURL").String()
	u, _ := url.Parse(ref)
	return u
}

// NewURL returning attribute 'newURL' with
// type string (idl: USVString).
func (_evt *HashChangeEvent) NewURL() *url.URL {
	ref := _evt.jsValue.Get("newURL").String()
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
func CastPageTransitionEvent(value js.Value) *PageTransitionEvent {
	if value.Type() != js.TypeObject {
		ICKError("casting PageTransitionEvent failed")
		return nil
	}
	ret := new(PageTransitionEvent)
	ret.jsValue = value
	return ret
}

// Persisted returning attribute 'persisted' with
func (_evt *PageTransitionEvent) Persisted() bool {
	return _evt.jsValue.Get("persisted").Bool()
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
func CastBeforeUnloadEvent(value js.Value) *BeforeUnloadEvent {
	if value.Type() != js.TypeObject {
		ICKError("casting BeforeUnloadEvent failed")
		return nil
	}
	ret := new(BeforeUnloadEvent)
	ret.jsValue = value
	return ret
}

// ReturnValue returning attribute 'returnValue' with
func (_evt *BeforeUnloadEvent) ReturnValue() string {
	return _evt.jsValue.Get("returnValue").String()
}

// SetReturnValue setting attribute 'returnValue' with
func (_evt *BeforeUnloadEvent) SetReturnValue(value string) {
	_evt.jsValue.Set("returnValue", value)
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
func CastUIEvent(value js.Value) *UIEvent {
	if value.Type() != js.TypeObject {
		ICKError("casting UIEvent failed")
		return nil
	}
	ret := new(UIEvent)
	ret.jsValue = value
	return ret
}

// View Returns a WindowProxy that contains the view that generated the event.
func (_evt *UIEvent) View() *Window {
	win := _evt.jsValue.Get("view")
	return CastWindow(win)
}

// Detail when non-zero, provides the current (or next, depending on the event) click count.
//
// https://developer.mozilla.org/en-US/docs/Web/API/UIEvent/detail
func (_evt *UIEvent) Detail() int {
	return _evt.jsValue.Get("detail").Int()
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
func CastMouseEvent(value js.Value) *MouseEvent {
	if value.Type() != js.TypeObject {
		ICKError("casting MouseEvent failed")
		return nil
	}
	ret := new(MouseEvent)
	ret.jsValue = value
	return ret
}

// CtrlKey returning attribute 'ctrlKey' with
func (_this *MouseEvent) CtrlKey() bool {
	return _this.jsValue.Get("ctrlKey").Bool()
}

// ShiftKey returning attribute 'shiftKey' with
func (_this *MouseEvent) ShiftKey() bool {
	return _this.jsValue.Get("shiftKey").Bool()
}

// AltKey returning attribute 'altKey' with
func (_this *MouseEvent) AltKey() bool {
	return _this.jsValue.Get("altKey").Bool()
}

// MetaKey returning attribute 'metaKey' with
func (_this *MouseEvent) MetaKey() bool {
	return _this.jsValue.Get("metaKey").Bool()
}

// Button returning attribute 'button' with
func (_this *MouseEvent) Button() int {
	return _this.jsValue.Get("button").Int()
}

// Buttons returning attribute 'buttons' with
func (_this *MouseEvent) Buttons() int {
	return _this.jsValue.Get("buttons").Int()
}

// RelatedTarget returning attribute 'relatedTarget' with
func (_this *MouseEvent) RelatedTarget() *EventTarget {
	value := _this.jsValue.Get("relatedTarget")
	return CastEventTarget(value)
}

// ScreenX returning attribute 'screenX' with
func (_this *MouseEvent) ScreenX() float64 {
	return _this.jsValue.Get("screenX").Float()
}

// ScreenY returning attribute 'screenY' with
func (_this *MouseEvent) ScreenY() float64 {
	return _this.jsValue.Get("screenY").Float()
}

// PageX returning attribute 'pageX' with
func (_this *MouseEvent) PageX() float64 {
	return _this.jsValue.Get("pageX").Float()
}

// PageY returning attribute 'pageY' with
func (_this *MouseEvent) PageY() float64 {
	return _this.jsValue.Get("pageY").Float()
}

// ClientX returning attribute 'clientX' with
func (_this *MouseEvent) ClientX() float64 {
	return _this.jsValue.Get("clientX").Float()
}

// ClientY returning attribute 'clientY' with
func (_this *MouseEvent) ClientY() float64 {
	return _this.jsValue.Get("clientY").Float()
}

// X returning attribute 'x' with
func (_this *MouseEvent) X() float64 {
	return _this.jsValue.Get("x").Float()
}

// Y returning attribute 'y' with
func (_this *MouseEvent) Y() float64 {
	return _this.jsValue.Get("y").Float()
}

// OffsetX returning attribute 'offsetX' with
func (_this *MouseEvent) OffsetX() float64 {
	return _this.jsValue.Get("offsetX").Float()
}

// OffsetY returning attribute 'offsetY' with
func (_this *MouseEvent) OffsetY() float64 {
	return _this.jsValue.Get("offsetY").Float()
}

// MovementX returning attribute 'movementX' with
func (_this *MouseEvent) MovementX() int {
	return _this.jsValue.Get("movementX").Int()
}

// MovementY returning attribute 'movementY' with
func (_this *MouseEvent) MovementY() int {
	return _this.jsValue.Get("movementY").Int()
}

func (_this *MouseEvent) GetModifierState(keyArg string) bool {
	return _this.jsValue.Call("getModifierState", keyArg).Bool()
}

/*********************************************************************************
* WheelEvent
 */

type DOM_DELTA uint

const (
	DOM_DELTA_PIXEL DOM_DELTA = 0x00
	DOM_DELTA_LINE  DOM_DELTA = 0x01
	DOM_DELTA_PAGE  DOM_DELTA = 0x02
)

type WheelEvent struct {
	MouseEvent
}

// CastWheelEvent is casting a js.Value into WheelEvent.
func CastWheelEvent(value js.Value) *WheelEvent {
	if value.Type() != js.TypeObject {
		ICKError("casting MouseEvent failed")
		return nil
	}
	ret := &WheelEvent{}
	ret.jsValue = value
	return ret
}

// DeltaX returning attribute 'deltaX' with
func (_this *WheelEvent) DeltaX() float64 {
	return _this.jsValue.Get("deltaX").Float()
}

// DeltaY returning attribute 'deltaY' with
func (_this *WheelEvent) DeltaY() float64 {
	return _this.jsValue.Get("deltaY").Float()
}

// DeltaZ returning attribute 'deltaZ' with
func (_this *WheelEvent) DeltaZ() float64 {
	return _this.jsValue.Get("deltaZ").Float()
}

// DeltaMode returning attribute 'deltaMode' with
func (_this *WheelEvent) DeltaMode() DOM_DELTA {
	return DOM_DELTA(_this.jsValue.Get("deltaMode").Int())
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
func CastFocusEvent(value js.Value) *FocusEvent {
	if value.Type() != js.TypeObject {
		ICKError("casting FocusEvent failed")
		return nil
	}
	ret := new(FocusEvent)
	ret.jsValue = value
	return ret
}

// RelatedTarget returning attribute 'relatedTarget'
func (_evt *FocusEvent) RelatedTarget() *EventTarget {
	value := _evt.jsValue.Get("relatedTarget")
	return CastEventTarget(value)
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
func CastPointerEvent(value js.Value) *PointerEvent {
	if value.Type() != js.TypeObject {
		ICKError("casting FocusEvent failed")
		return nil
	}
	ret := new(PointerEvent)
	ret.jsValue = value
	return ret
}

// PointerId returning attribute 'pointerId' with
func (_this *PointerEvent) PointerId() int {
	return _this.jsValue.Get("pointerId").Int()
}

// Width returning attribute 'width' with
func (_this *PointerEvent) Width() float64 {
	return _this.jsValue.Get("width").Float()
}

// Height returning attribute 'height' with
func (_this *PointerEvent) Height() float64 {
	return _this.jsValue.Get("height").Float()
}

// Pressure returning attribute 'pressure' with
func (_this *PointerEvent) Pressure() float64 {
	return _this.jsValue.Get("pressure").Float()
}

// TangentialPressure returning attribute 'tangentialPressure' with
func (_this *PointerEvent) TangentialPressure() float64 {
	return _this.jsValue.Get("tangentialPressure").Float()
}

// TiltX returning attribute 'tiltX' with
func (_this *PointerEvent) TiltX() int {
	return _this.jsValue.Get("tiltX").Int()
}

// TiltY returning attribute 'tiltY' with
func (_this *PointerEvent) TiltY() int {
	return _this.jsValue.Get("tiltY").Int()
}

// Twist returning attribute 'twist' with
func (_this *PointerEvent) Twist() int {
	return _this.jsValue.Get("twist").Int()
}

// PointerType returning attribute 'pointerType' with
func (_this *PointerEvent) PointerType() string {
	return _this.jsValue.Get("pointerType").String()
}

// IsPrimary returning attribute 'isPrimary' with
func (_this *PointerEvent) IsPrimary() bool {
	return _this.jsValue.Get("isPrimary").Bool()
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
func CastInputEvent(value js.Value) *InputEvent {
	if value.Type() != js.TypeObject {
		ICKError("casting InputEvent failed")
		return nil
	}
	ret := new(InputEvent)
	ret.jsValue = value
	return ret
}

// Data returning attribute 'data' with
func (_this *InputEvent) Data() string {
	return _this.jsValue.Get("data").String()
}

// IsComposing returning attribute 'isComposing' with
func (_this *InputEvent) IsComposing() bool {
	return _this.jsValue.Get("isComposing").Bool()
}

// InputType returning attribute 'inputType' with
func (_this *InputEvent) InputType() string {
	return _this.jsValue.Get("inputType").String()
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

type DOM_KEY_LOCATION uint

const (
	DOM_KEY_LOCATION_STANDARD DOM_KEY_LOCATION = 0x00
	DOM_KEY_LOCATION_LEFT     DOM_KEY_LOCATION = 0x01
	DOM_KEY_LOCATION_RIGHT    DOM_KEY_LOCATION = 0x02
	DOM_KEY_LOCATION_NUMPAD   DOM_KEY_LOCATION = 0x03
)

type KeyboardEvent struct {
	UIEvent
}

// CastKeyboardEvent is casting a js.Value into KeyboardEvent.
func CastKeyboardEvent(value js.Value) *KeyboardEvent {
	if value.Type() != js.TypeObject {
		ICKError("casting KeyboardEvent failed")
		return nil
	}
	ret := new(KeyboardEvent)
	ret.jsValue = value
	return ret
}

// Key returning attribute 'key' with
func (_this *KeyboardEvent) Key() string {
	return _this.jsValue.Get("key").String()
}

// Code returning attribute 'code' with
func (_this *KeyboardEvent) Code() string {
	return _this.jsValue.Get("code").String()
}

// Location returning attribute 'location' with
func (_this *KeyboardEvent) Location() DOM_KEY_LOCATION {
	return DOM_KEY_LOCATION(_this.jsValue.Get("location").Int())
}

// CtrlKey returning attribute 'ctrlKey' with
func (_this *KeyboardEvent) CtrlKey() bool {
	return _this.jsValue.Get("ctrlKey").Bool()
}

// ShiftKey returning attribute 'shiftKey' with
func (_this *KeyboardEvent) ShiftKey() bool {
	return _this.jsValue.Get("shiftKey").Bool()
}

// AltKey returning attribute 'altKey' with
func (_this *KeyboardEvent) AltKey() bool {
	return _this.jsValue.Get("altKey").Bool()
}

// MetaKey returning attribute 'metaKey' with
func (_this *KeyboardEvent) MetaKey() bool {
	return _this.jsValue.Get("metaKey").Bool()
}

// Repeat returning attribute 'repeat' with
func (_this *KeyboardEvent) Repeat() bool {
	return _this.jsValue.Get("repeat").Bool()
}

// IsComposing returning attribute 'isComposing' with
func (_this *KeyboardEvent) IsComposing() bool {
	return _this.jsValue.Get("isComposing").Bool()
}

// CharCode returning attribute 'charCode' with
func (_this *KeyboardEvent) CharCode() uint {
	return uint(_this.jsValue.Get("charCode").Int())
}

// KeyCode returning attribute 'keyCode' with
func (_this *KeyboardEvent) KeyCode() uint {
	return uint(_this.jsValue.Get("keyCode").Int())
}

// returns the current state of the specified modifier key: true if the modifier is active (that is the modifier key is pressed or locked), otherwise, false.
//
// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/getModifierState
func (_this *KeyboardEvent) GetModifierState(keyArg string) bool {
	return _this.jsValue.Call("getModifierState", keyArg).Bool()
}
