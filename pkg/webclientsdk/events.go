package webclientsdk

import (
	"syscall/js"
)

/*********************************************************************************
 * HashChangeEvent
 */

type HashChangeEvent struct {
	Event
}

// HashChangeEventFromJS is casting a js.Value into HashChangeEvent.
func HashChangeEventFromJS(value js.Value) *HashChangeEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &HashChangeEvent{}
	ret.jsValue = value
	return ret
}

// OldURL returning attribute 'oldURL' with
// type string (idl: USVString).
func (_this *HashChangeEvent) OldURL() string {
	var ret string
	value := _this.jsValue.Get("oldURL")
	ret = (value).String()
	return ret
}

// NewURL returning attribute 'newURL' with
// type string (idl: USVString).
func (_this *HashChangeEvent) NewURL() string {
	var ret string
	value := _this.jsValue.Get("newURL")
	ret = (value).String()
	return ret
}

/*********************************************************************************
 * PageTransitionEvent
 */

type PageTransitionEvent struct {
	Event
}

// PageTransitionEventFromJS is casting a js.Value into PageTransitionEvent.
func PageTransitionEventFromJS(value js.Value) *PageTransitionEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PageTransitionEvent{}
	ret.jsValue = value
	return ret
}

// Persisted returning attribute 'persisted' with
// type bool (idl: boolean).
func (_this *PageTransitionEvent) Persisted() bool {
	var ret bool
	value := _this.jsValue.Get("persisted")
	ret = (value).Bool()
	return ret
}

/*********************************************************************************
 * BeforeUnloadEvent
 */

type BeforeUnloadEvent struct {
	Event
}

// BeforeUnloadEventFromJS is casting a js.Value into BeforeUnloadEvent.
func BeforeUnloadEventFromJS(value js.Value) *BeforeUnloadEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &BeforeUnloadEvent{}
	ret.jsValue = value
	return ret
}

// ReturnValue returning attribute 'returnValue' with
// type string (idl: DOMString).
func (_this *BeforeUnloadEvent) ReturnValue() string {
	var ret string
	value := _this.jsValue.Get("returnValue")
	ret = (value).String()
	return ret
}

// SetReturnValue setting attribute 'returnValue' with
// type string (idl: DOMString).
func (_this *BeforeUnloadEvent) SetReturnValue(value string) {
	input := value
	_this.jsValue.Set("returnValue", input)
}

/*********************************************************************************
 * UIEvent
 */

type UIEvent struct {
	Event
}

// UIEventFromJS is casting a js.Value into UIEvent.
func UIEventFromJS(value js.Value) *UIEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &UIEvent{}
	ret.jsValue = value
	return ret
}

// View returning attribute 'view' with
// type js.Value (idl: Window).
func (_this *UIEvent) View() js.Value {
	var ret js.Value
	value := _this.jsValue.Get("view")
	ret = value
	return ret
}

// Detail returning attribute 'detail' with
// type int (idl: long).
func (_this *UIEvent) Detail() int {
	var ret int
	value := _this.jsValue.Get("detail")
	ret = (value).Int()
	return ret
}

/*********************************************************************************
 * MouseEvent
 */

type MouseEvent struct {
	UIEvent
}

// MouseEventFromJS is casting a js.Value into MouseEvent.
func MouseEventFromJS(value js.Value) *MouseEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &MouseEvent{}
	ret.jsValue = value
	return ret
}

// CtrlKey returning attribute 'ctrlKey' with
// type bool (idl: boolean).
func (_this *MouseEvent) CtrlKey() bool {
	var ret bool
	value := _this.jsValue.Get("ctrlKey")
	ret = (value).Bool()
	return ret
}

// ShiftKey returning attribute 'shiftKey' with
// type bool (idl: boolean).
func (_this *MouseEvent) ShiftKey() bool {
	var ret bool
	value := _this.jsValue.Get("shiftKey")
	ret = (value).Bool()
	return ret
}

// AltKey returning attribute 'altKey' with
// type bool (idl: boolean).
func (_this *MouseEvent) AltKey() bool {
	var ret bool
	value := _this.jsValue.Get("altKey")
	ret = (value).Bool()
	return ret
}

// MetaKey returning attribute 'metaKey' with
// type bool (idl: boolean).
func (_this *MouseEvent) MetaKey() bool {
	var ret bool
	value := _this.jsValue.Get("metaKey")
	ret = (value).Bool()
	return ret
}

// Button returning attribute 'button' with
// type int (idl: short).
func (_this *MouseEvent) Button() int {
	var ret int
	value := _this.jsValue.Get("button")
	ret = (value).Int()
	return ret
}

// Buttons returning attribute 'buttons' with
// type int (idl: unsigned short).
func (_this *MouseEvent) Buttons() int {
	var ret int
	value := _this.jsValue.Get("buttons")
	ret = (value).Int()
	return ret
}

// RelatedTarget returning attribute 'relatedTarget' with
// type domcore.EventTarget (idl: EventTarget).
func (_this *MouseEvent) RelatedTarget() *EventTarget {
	var ret *EventTarget
	value := _this.jsValue.Get("relatedTarget")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = MakeEventTargetFromJS(value)
	}
	return ret
}

// ScreenX returning attribute 'screenX' with
// type float64 (idl: double).
func (_this *MouseEvent) ScreenX() float64 {
	var ret float64
	value := _this.jsValue.Get("screenX")
	ret = (value).Float()
	return ret
}

// ScreenY returning attribute 'screenY' with
// type float64 (idl: double).
func (_this *MouseEvent) ScreenY() float64 {
	var ret float64
	value := _this.jsValue.Get("screenY")
	ret = (value).Float()
	return ret
}

// PageX returning attribute 'pageX' with
// type float64 (idl: double).
func (_this *MouseEvent) PageX() float64 {
	var ret float64
	value := _this.jsValue.Get("pageX")
	ret = (value).Float()
	return ret
}

// PageY returning attribute 'pageY' with
// type float64 (idl: double).
func (_this *MouseEvent) PageY() float64 {
	var ret float64
	value := _this.jsValue.Get("pageY")
	ret = (value).Float()
	return ret
}

// ClientX returning attribute 'clientX' with
// type float64 (idl: double).
func (_this *MouseEvent) ClientX() float64 {
	var ret float64
	value := _this.jsValue.Get("clientX")
	ret = (value).Float()
	return ret
}

// ClientY returning attribute 'clientY' with
// type float64 (idl: double).
func (_this *MouseEvent) ClientY() float64 {
	var ret float64
	value := _this.jsValue.Get("clientY")
	ret = (value).Float()
	return ret
}

// X returning attribute 'x' with
// type float64 (idl: double).
func (_this *MouseEvent) X() float64 {
	var ret float64
	value := _this.jsValue.Get("x")
	ret = (value).Float()
	return ret
}

// Y returning attribute 'y' with
// type float64 (idl: double).
func (_this *MouseEvent) Y() float64 {
	var ret float64
	value := _this.jsValue.Get("y")
	ret = (value).Float()
	return ret
}

// OffsetX returning attribute 'offsetX' with
// type float64 (idl: double).
func (_this *MouseEvent) OffsetX() float64 {
	var ret float64
	value := _this.jsValue.Get("offsetX")
	ret = (value).Float()
	return ret
}

// OffsetY returning attribute 'offsetY' with
// type float64 (idl: double).
func (_this *MouseEvent) OffsetY() float64 {
	var ret float64
	value := _this.jsValue.Get("offsetY")
	ret = (value).Float()
	return ret
}

// MovementX returning attribute 'movementX' with
// type int (idl: long).
func (_this *MouseEvent) MovementX() int {
	var ret int
	value := _this.jsValue.Get("movementX")
	ret = (value).Int()
	return ret
}

// MovementY returning attribute 'movementY' with
// type int (idl: long).
func (_this *MouseEvent) MovementY() int {
	var ret int
	value := _this.jsValue.Get("movementY")
	ret = (value).Int()
	return ret
}

func (_this *MouseEvent) GetModifierState(keyArg string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := keyArg
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("getModifierState", _args[0:_end]...)
	return (_returned).Bool()
}

/*********************************************************************************
* WheelEvent
 */

type WheelEvent struct {
	MouseEvent
}

// WheelEventFromJS is casting a js.Value into WheelEvent.
func WheelEventFromJS(value js.Value) *WheelEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &WheelEvent{}
	ret.jsValue = value
	return ret
}

const (
	DOM_DELTA_PIXEL uint = 0x00
	DOM_DELTA_LINE  uint = 0x01
	DOM_DELTA_PAGE  uint = 0x02
)

// DeltaX returning attribute 'deltaX' with
// type float64 (idl: double).
func (_this *WheelEvent) DeltaX() float64 {
	var ret float64
	value := _this.jsValue.Get("deltaX")
	ret = (value).Float()
	return ret
}

// DeltaY returning attribute 'deltaY' with
// type float64 (idl: double).
func (_this *WheelEvent) DeltaY() float64 {
	var ret float64
	value := _this.jsValue.Get("deltaY")
	ret = (value).Float()
	return ret
}

// DeltaZ returning attribute 'deltaZ' with
// type float64 (idl: double).
func (_this *WheelEvent) DeltaZ() float64 {
	var ret float64
	value := _this.jsValue.Get("deltaZ")
	ret = (value).Float()
	return ret
}

// DeltaMode returning attribute 'deltaMode' with
// type uint (idl: unsigned long).
func (_this *WheelEvent) DeltaMode() uint {
	var ret uint
	value := _this.jsValue.Get("deltaMode")
	ret = (uint)((value).Int())
	return ret
}

/**********************************************************************************
 * class: FocusEvent
 */

type FocusEvent struct {
	UIEvent
}

// FocusEventFromJS is casting a js.Value into FocusEvent.
func FocusEventFromJS(value js.Value) *FocusEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &FocusEvent{}
	ret.jsValue = value
	return ret
}

// RelatedTarget returning attribute 'relatedTarget' with
// type domcore.EventTarget (idl: EventTarget).
func (_this *FocusEvent) RelatedTarget() *EventTarget {
	var ret *EventTarget
	value := _this.jsValue.Get("relatedTarget")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = MakeEventTargetFromJS(value)
	}
	return ret
}

/**********************************************************************************
 * PointerEvent
 */

type PointerEvent struct {
	MouseEvent
}

// PointerEventFromJS is casting a js.Value into PointerEvent.
func PointerEventFromJS(value js.Value) *PointerEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &PointerEvent{}
	ret.jsValue = value
	return ret
}

// PointerId returning attribute 'pointerId' with
// type int (idl: long).
func (_this *PointerEvent) PointerId() int {
	var ret int
	value := _this.jsValue.Get("pointerId")
	ret = (value).Int()
	return ret
}

// Width returning attribute 'width' with
// type float64 (idl: double).
func (_this *PointerEvent) Width() float64 {
	var ret float64
	value := _this.jsValue.Get("width")
	ret = (value).Float()
	return ret
}

// Height returning attribute 'height' with
// type float64 (idl: double).
func (_this *PointerEvent) Height() float64 {
	var ret float64
	value := _this.jsValue.Get("height")
	ret = (value).Float()
	return ret
}

// Pressure returning attribute 'pressure' with
// type float32 (idl: float).
func (_this *PointerEvent) Pressure() float32 {
	var ret float32
	value := _this.jsValue.Get("pressure")
	ret = (float32)((value).Float())
	return ret
}

// TangentialPressure returning attribute 'tangentialPressure' with
// type float32 (idl: float).
func (_this *PointerEvent) TangentialPressure() float32 {
	var ret float32
	value := _this.jsValue.Get("tangentialPressure")
	ret = (float32)((value).Float())
	return ret
}

// TiltX returning attribute 'tiltX' with
// type int (idl: long).
func (_this *PointerEvent) TiltX() int {
	var ret int
	value := _this.jsValue.Get("tiltX")
	ret = (value).Int()
	return ret
}

// TiltY returning attribute 'tiltY' with
// type int (idl: long).
func (_this *PointerEvent) TiltY() int {
	var ret int
	value := _this.jsValue.Get("tiltY")
	ret = (value).Int()
	return ret
}

// Twist returning attribute 'twist' with
// type int (idl: long).
func (_this *PointerEvent) Twist() int {
	var ret int
	value := _this.jsValue.Get("twist")
	ret = (value).Int()
	return ret
}

// PointerType returning attribute 'pointerType' with
// type string (idl: DOMString).
func (_this *PointerEvent) PointerType() string {
	var ret string
	value := _this.jsValue.Get("pointerType")
	ret = (value).String()
	return ret
}

// IsPrimary returning attribute 'isPrimary' with
// type bool (idl: boolean).
func (_this *PointerEvent) IsPrimary() bool {
	var ret bool
	value := _this.jsValue.Get("isPrimary")
	ret = (value).Bool()
	return ret
}

func (_this *PointerEvent) GetCoalescedEvents() (_result []*PointerEvent) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("getCoalescedEvents", _args[0:_end]...)
	var (
		_converted []*PointerEvent // javascript: sequence<PointerEvent> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*PointerEvent, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *PointerEvent
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = PointerEventFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *PointerEvent) GetPredictedEvents() (_result []*PointerEvent) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("getPredictedEvents", _args[0:_end]...)
	var (
		_converted []*PointerEvent // javascript: sequence<PointerEvent> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]*PointerEvent, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 *PointerEvent
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = PointerEventFromJS(__seq_in0)
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

/**********************************************************************************
 * InputEvent
 */

type InputEvent struct {
	UIEvent
}

// InputEventFromJS is casting a js.Value into InputEvent.
func InputEventFromJS(value js.Value) *InputEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &InputEvent{}
	ret.jsValue = value
	return ret
}

// Data returning attribute 'data' with
// type string (idl: DOMString).
func (_this *InputEvent) Data() *string {
	var ret *string
	value := _this.jsValue.Get("data")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// IsComposing returning attribute 'isComposing' with
// type bool (idl: boolean).
func (_this *InputEvent) IsComposing() bool {
	var ret bool
	value := _this.jsValue.Get("isComposing")
	ret = (value).Bool()
	return ret
}

// InputType returning attribute 'inputType' with
// type string (idl: DOMString).
func (_this *InputEvent) InputType() string {
	var ret string
	value := _this.jsValue.Get("inputType")
	ret = (value).String()
	return ret
}

/**********************************************************************************
 * KeyboardEvent
 */

type KeyboardEvent struct {
	UIEvent
}

// KeyboardEventFromJS is casting a js.Value into KeyboardEvent.
func KeyboardEventFromJS(value js.Value) *KeyboardEvent {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &KeyboardEvent{}
	ret.jsValue = value
	return ret
}

const (
	DOM_KEY_LOCATION_STANDARD uint = 0x00
	DOM_KEY_LOCATION_LEFT     uint = 0x01
	DOM_KEY_LOCATION_RIGHT    uint = 0x02
	DOM_KEY_LOCATION_NUMPAD   uint = 0x03
)

// Key returning attribute 'key' with
// type string (idl: DOMString).
func (_this *KeyboardEvent) Key() string {
	var ret string
	value := _this.jsValue.Get("key")
	ret = (value).String()
	return ret
}

// Code returning attribute 'code' with
// type string (idl: DOMString).
func (_this *KeyboardEvent) Code() string {
	var ret string
	value := _this.jsValue.Get("code")
	ret = (value).String()
	return ret
}

// Location returning attribute 'location' with
// type uint (idl: unsigned long).
func (_this *KeyboardEvent) Location() uint {
	var ret uint
	value := _this.jsValue.Get("location")
	ret = (uint)((value).Int())
	return ret
}

// CtrlKey returning attribute 'ctrlKey' with
// type bool (idl: boolean).
func (_this *KeyboardEvent) CtrlKey() bool {
	var ret bool
	value := _this.jsValue.Get("ctrlKey")
	ret = (value).Bool()
	return ret
}

// ShiftKey returning attribute 'shiftKey' with
// type bool (idl: boolean).
func (_this *KeyboardEvent) ShiftKey() bool {
	var ret bool
	value := _this.jsValue.Get("shiftKey")
	ret = (value).Bool()
	return ret
}

// AltKey returning attribute 'altKey' with
// type bool (idl: boolean).
func (_this *KeyboardEvent) AltKey() bool {
	var ret bool
	value := _this.jsValue.Get("altKey")
	ret = (value).Bool()
	return ret
}

// MetaKey returning attribute 'metaKey' with
// type bool (idl: boolean).
func (_this *KeyboardEvent) MetaKey() bool {
	var ret bool
	value := _this.jsValue.Get("metaKey")
	ret = (value).Bool()
	return ret
}

// Repeat returning attribute 'repeat' with
// type bool (idl: boolean).
func (_this *KeyboardEvent) Repeat() bool {
	var ret bool
	value := _this.jsValue.Get("repeat")
	ret = (value).Bool()
	return ret
}

// IsComposing returning attribute 'isComposing' with
// type bool (idl: boolean).
func (_this *KeyboardEvent) IsComposing() bool {
	var ret bool
	value := _this.jsValue.Get("isComposing")
	ret = (value).Bool()
	return ret
}

// CharCode returning attribute 'charCode' with
// type uint (idl: unsigned long).
func (_this *KeyboardEvent) CharCode() uint {
	var ret uint
	value := _this.jsValue.Get("charCode")
	ret = (uint)((value).Int())
	return ret
}

// KeyCode returning attribute 'keyCode' with
// type uint (idl: unsigned long).
func (_this *KeyboardEvent) KeyCode() uint {
	var ret uint
	value := _this.jsValue.Get("keyCode")
	ret = (uint)((value).Int())
	return ret
}

func (_this *KeyboardEvent) GetModifierState(keyArg string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := keyArg
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("getModifierState", _args[0:_end]...)
	return (_returned).Bool()
}
