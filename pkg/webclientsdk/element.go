package webclientsdk

import (
	"syscall/js"
)

/****************************************************************************
* Element
*****************************************************************************/

// Element is the most general base class from which all element objects (i.e. objects that represent elements) in a Document inherit.
// It only has methods and properties common to all kinds of elements. More specific classes inherit from Element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element
type Element struct {
	Node
}

// ElementFromJS is casting a js.Value into Element.
func ElementFromJS(value js.Value) *Element {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Element{}
	ret.jsValue = value
	return ret
}

/****************************************************************************
* Element's properties
*****************************************************************************/

// Prefix returning attribute 'prefix' with
// type string (idl: DOMString).
func (_this *Element) Prefix() *string {
	var ret *string
	value := _this.jsValue.Get("prefix")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// LocalName returning attribute 'localName' with
// type string (idl: DOMString).
func (_this *Element) LocalName() string {
	var ret string
	value := _this.jsValue.Get("localName")
	ret = (value).String()
	return ret
}

// TagName returning attribute 'tagName' with
// type string (idl: DOMString).
func (_this *Element) TagName() string {
	var ret string
	value := _this.jsValue.Get("tagName")
	ret = (value).String()
	return ret
}

// Id returning attribute 'id' with
// type string (idl: DOMString).
func (_this *Element) Id() string {
	var ret string
	value := _this.jsValue.Get("id")
	ret = (value).String()
	return ret
}

// SetId setting attribute 'id' with
// type string (idl: DOMString).
func (_this *Element) SetId(value string) {
	input := value
	_this.jsValue.Set("id", input)
}

// ClassName returning attribute 'className' with
// type string (idl: DOMString).
func (_this *Element) ClassName() string {
	var ret string
	value := _this.jsValue.Get("className")
	ret = (value).String()
	return ret
}

// SetClassName setting attribute 'className' with
// type string (idl: DOMString).
func (_this *Element) SetClassName(value string) {
	input := value
	_this.jsValue.Set("className", input)
}

// ClassList returning attribute 'classList' with
// type domcore.DOMTokenList (idl: DOMTokenList).
func (_this *Element) ClassList() *DOMTokenList {
	var ret *DOMTokenList
	value := _this.jsValue.Get("classList")
	ret = DOMTokenListFromJS(value)
	return ret
}

// Slot returning attribute 'slot' with
// type string (idl: DOMString).
func (_this *Element) Slot() string {
	var ret string
	value := _this.jsValue.Get("slot")
	ret = (value).String()
	return ret
}

// SetSlot setting attribute 'slot' with
// type string (idl: DOMString).
func (_this *Element) SetSlot(value string) {
	input := value
	_this.jsValue.Set("slot", input)
}

// Attributes returning attribute 'attributes' with
// type NamedNodeMap (idl: NamedNodeMap).
func (_this *Element) Attributes() *NamedAttrMap {
	var ret *NamedAttrMap
	value := _this.jsValue.Get("attributes")
	ret = NamedNodeMapFromJS(value)
	return ret
}

// InnerHTML returning attribute 'innerHTML' with
// type string (idl: DOMString).
func (_this *Element) InnerHTML() string {
	var ret string
	value := _this.jsValue.Get("innerHTML")
	ret = (value).String()
	return ret
}

// SetInnerHTML setting attribute 'innerHTML' with
// type string (idl: DOMString).
func (_this *Element) SetInnerHTML(value string) {
	input := value
	_this.jsValue.Set("innerHTML", input)
}

// OuterHTML returning attribute 'outerHTML' with
// type string (idl: DOMString).
func (_this *Element) OuterHTML() string {
	var ret string
	value := _this.jsValue.Get("outerHTML")
	ret = (value).String()
	return ret
}

// SetOuterHTML setting attribute 'outerHTML' with
// type string (idl: DOMString).
func (_this *Element) SetOuterHTML(value string) {
	input := value
	_this.jsValue.Set("outerHTML", input)
}

// ScrollTop returning attribute 'scrollTop' with
// type float64 (idl: unrestricted double).
func (_this *Element) ScrollTop() float64 {
	var ret float64
	value := _this.jsValue.Get("scrollTop")
	ret = (value).Float()
	return ret
}

// SetScrollTop setting attribute 'scrollTop' with
// type float64 (idl: unrestricted double).
func (_this *Element) SetScrollTop(value float64) {
	input := value
	_this.jsValue.Set("scrollTop", input)
}

// ScrollLeft returning attribute 'scrollLeft' with
// type float64 (idl: unrestricted double).
func (_this *Element) ScrollLeft() float64 {
	var ret float64
	value := _this.jsValue.Get("scrollLeft")
	ret = (value).Float()
	return ret
}

// SetScrollLeft setting attribute 'scrollLeft' with
// type float64 (idl: unrestricted double).
func (_this *Element) SetScrollLeft(value float64) {
	input := value
	_this.jsValue.Set("scrollLeft", input)
}

// ScrollWidth returning attribute 'scrollWidth' with
// type int (idl: long).
func (_this *Element) ScrollWidth() int {
	var ret int
	value := _this.jsValue.Get("scrollWidth")
	ret = (value).Int()
	return ret
}

// ScrollHeight returning attribute 'scrollHeight' with
// type int (idl: long).
func (_this *Element) ScrollHeight() int {
	var ret int
	value := _this.jsValue.Get("scrollHeight")
	ret = (value).Int()
	return ret
}

// ClientTop returning attribute 'clientTop' with
// type int (idl: long).
func (_this *Element) ClientTop() int {
	var ret int
	value := _this.jsValue.Get("clientTop")
	ret = (value).Int()
	return ret
}

// ClientLeft returning attribute 'clientLeft' with
// type int (idl: long).
func (_this *Element) ClientLeft() int {
	var ret int
	value := _this.jsValue.Get("clientLeft")
	ret = (value).Int()
	return ret
}

// ClientWidth returning attribute 'clientWidth' with
// type int (idl: long).
func (_this *Element) ClientWidth() int {
	var ret int
	value := _this.jsValue.Get("clientWidth")
	ret = (value).Int()
	return ret
}

// ClientHeight returning attribute 'clientHeight' with
// type int (idl: long).
func (_this *Element) ClientHeight() int {
	var ret int
	value := _this.jsValue.Get("clientHeight")
	ret = (value).Int()
	return ret
}

// OnFullscreenChange returning attribute 'onfullscreenchange' with
// type EventHandler (idl: EventHandlerNonNull).
func (_this *Element) OnFullscreenChange() EventHandlerFunc {
	var ret EventHandlerFunc
	value := _this.jsValue.Get("onfullscreenchange")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = EventHandlerFromJS(value)
	}
	return ret
}

// OnFullscreenError returning attribute 'onfullscreenerror' with
// type EventHandler (idl: EventHandlerNonNull).
func (_this *Element) OnFullscreenError() EventHandlerFunc {
	var ret EventHandlerFunc
	value := _this.jsValue.Get("onfullscreenerror")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		ret = EventHandlerFromJS(value)
	}
	return ret
}

// Children returning attribute 'children' with
// type HTMLCollection (idl: HTMLCollection).
func (_this *Element) Children() *HTMLCollection {
	var ret *HTMLCollection
	value := _this.jsValue.Get("children")
	ret = HTMLCollectionFromJS(value)
	return ret
}

// FirstElementChild returning attribute 'firstElementChild' with
// type Element (idl: Element).
func (_this *Element) FirstElementChild() *Element {
	value := _this.jsValue.Get("firstElementChild")
	return ElementFromJS(value)
}

// LastElementChild returning attribute 'lastElementChild' with
// type Element (idl: Element).
func (_this *Element) LastElementChild() *Element {
	value := _this.jsValue.Get("lastElementChild")
	return ElementFromJS(value)
}

// ChildElementCount returning attribute 'childElementCount' with
// type uint (idl: unsigned long).
func (_this *Element) ChildElementCount() uint {
	var ret uint
	value := _this.jsValue.Get("childElementCount")
	ret = (uint)((value).Int())
	return ret
}

// PreviousElementSibling returning attribute 'previousElementSibling' with
// type Element (idl: Element).
func (_this *Element) PreviousElementSibling() *Element {
	value := _this.jsValue.Get("previousElementSibling")
	return ElementFromJS(value)
}

// NextElementSibling returning attribute 'nextElementSibling' with
// type Element (idl: Element).
func (_this *Element) NextElementSibling() *Element {
	value := _this.jsValue.Get("nextElementSibling")
	return ElementFromJS(value)
}

// AssignedSlot returning attribute 'assignedSlot' with
// type js.Value (idl: HTMLSlotElement).
func (_this *Element) AssignedSlot() js.Value {
	var ret js.Value
	value := _this.jsValue.Get("assignedSlot")
	ret = value
	return ret
}

// Role returning attribute 'role' with
// type string (idl: DOMString).
func (_this *Element) Role() *string {
	var ret *string
	value := _this.jsValue.Get("role")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetRole setting attribute 'role' with
// type string (idl: DOMString).
func (_this *Element) SetRole(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("role", input)
}

// AriaActiveDescendant returning attribute 'ariaActiveDescendant' with
// type string (idl: DOMString).
func (_this *Element) AriaActiveDescendant() *string {
	var ret *string
	value := _this.jsValue.Get("ariaActiveDescendant")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaActiveDescendant setting attribute 'ariaActiveDescendant' with
// type string (idl: DOMString).
func (_this *Element) SetAriaActiveDescendant(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaActiveDescendant", input)
}

// AriaAtomic returning attribute 'ariaAtomic' with
// type string (idl: DOMString).
func (_this *Element) AriaAtomic() *string {
	var ret *string
	value := _this.jsValue.Get("ariaAtomic")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaAtomic setting attribute 'ariaAtomic' with
// type string (idl: DOMString).
func (_this *Element) SetAriaAtomic(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaAtomic", input)
}

// AriaAutoComplete returning attribute 'ariaAutoComplete' with
// type string (idl: DOMString).
func (_this *Element) AriaAutoComplete() *string {
	var ret *string
	value := _this.jsValue.Get("ariaAutoComplete")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaAutoComplete setting attribute 'ariaAutoComplete' with
// type string (idl: DOMString).
func (_this *Element) SetAriaAutoComplete(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaAutoComplete", input)
}

// AriaBusy returning attribute 'ariaBusy' with
// type string (idl: DOMString).
func (_this *Element) AriaBusy() *string {
	var ret *string
	value := _this.jsValue.Get("ariaBusy")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaBusy setting attribute 'ariaBusy' with
// type string (idl: DOMString).
func (_this *Element) SetAriaBusy(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaBusy", input)
}

// AriaChecked returning attribute 'ariaChecked' with
// type string (idl: DOMString).
func (_this *Element) AriaChecked() *string {
	var ret *string
	value := _this.jsValue.Get("ariaChecked")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaChecked setting attribute 'ariaChecked' with
// type string (idl: DOMString).
func (_this *Element) SetAriaChecked(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaChecked", input)
}

// AriaColCount returning attribute 'ariaColCount' with
// type string (idl: DOMString).
func (_this *Element) AriaColCount() *string {
	var ret *string
	value := _this.jsValue.Get("ariaColCount")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaColCount setting attribute 'ariaColCount' with
// type string (idl: DOMString).
func (_this *Element) SetAriaColCount(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaColCount", input)
}

// AriaColIndex returning attribute 'ariaColIndex' with
// type string (idl: DOMString).
func (_this *Element) AriaColIndex() *string {
	var ret *string
	value := _this.jsValue.Get("ariaColIndex")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaColIndex setting attribute 'ariaColIndex' with
// type string (idl: DOMString).
func (_this *Element) SetAriaColIndex(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaColIndex", input)
}

// AriaColSpan returning attribute 'ariaColSpan' with
// type string (idl: DOMString).
func (_this *Element) AriaColSpan() *string {
	var ret *string
	value := _this.jsValue.Get("ariaColSpan")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaColSpan setting attribute 'ariaColSpan' with
// type string (idl: DOMString).
func (_this *Element) SetAriaColSpan(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaColSpan", input)
}

// AriaControls returning attribute 'ariaControls' with
// type string (idl: DOMString).
func (_this *Element) AriaControls() *string {
	var ret *string
	value := _this.jsValue.Get("ariaControls")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaControls setting attribute 'ariaControls' with
// type string (idl: DOMString).
func (_this *Element) SetAriaControls(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaControls", input)
}

// AriaCurrent returning attribute 'ariaCurrent' with
// type string (idl: DOMString).
func (_this *Element) AriaCurrent() *string {
	var ret *string
	value := _this.jsValue.Get("ariaCurrent")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaCurrent setting attribute 'ariaCurrent' with
// type string (idl: DOMString).
func (_this *Element) SetAriaCurrent(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaCurrent", input)
}

// AriaDescribedBy returning attribute 'ariaDescribedBy' with
// type string (idl: DOMString).
func (_this *Element) AriaDescribedBy() *string {
	var ret *string
	value := _this.jsValue.Get("ariaDescribedBy")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaDescribedBy setting attribute 'ariaDescribedBy' with
// type string (idl: DOMString).
func (_this *Element) SetAriaDescribedBy(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaDescribedBy", input)
}

// AriaDetails returning attribute 'ariaDetails' with
// type string (idl: DOMString).
func (_this *Element) AriaDetails() *string {
	var ret *string
	value := _this.jsValue.Get("ariaDetails")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaDetails setting attribute 'ariaDetails' with
// type string (idl: DOMString).
func (_this *Element) SetAriaDetails(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaDetails", input)
}

// AriaDisabled returning attribute 'ariaDisabled' with
// type string (idl: DOMString).
func (_this *Element) AriaDisabled() *string {
	var ret *string
	value := _this.jsValue.Get("ariaDisabled")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaDisabled setting attribute 'ariaDisabled' with
// type string (idl: DOMString).
func (_this *Element) SetAriaDisabled(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaDisabled", input)
}

// AriaErrorMessage returning attribute 'ariaErrorMessage' with
// type string (idl: DOMString).
func (_this *Element) AriaErrorMessage() *string {
	var ret *string
	value := _this.jsValue.Get("ariaErrorMessage")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaErrorMessage setting attribute 'ariaErrorMessage' with
// type string (idl: DOMString).
func (_this *Element) SetAriaErrorMessage(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaErrorMessage", input)
}

// AriaExpanded returning attribute 'ariaExpanded' with
// type string (idl: DOMString).
func (_this *Element) AriaExpanded() *string {
	var ret *string
	value := _this.jsValue.Get("ariaExpanded")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaExpanded setting attribute 'ariaExpanded' with
// type string (idl: DOMString).
func (_this *Element) SetAriaExpanded(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaExpanded", input)
}

// AriaFlowTo returning attribute 'ariaFlowTo' with
// type string (idl: DOMString).
func (_this *Element) AriaFlowTo() *string {
	var ret *string
	value := _this.jsValue.Get("ariaFlowTo")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaFlowTo setting attribute 'ariaFlowTo' with
// type string (idl: DOMString).
func (_this *Element) SetAriaFlowTo(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaFlowTo", input)
}

// AriaHasPopup returning attribute 'ariaHasPopup' with
// type string (idl: DOMString).
func (_this *Element) AriaHasPopup() *string {
	var ret *string
	value := _this.jsValue.Get("ariaHasPopup")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaHasPopup setting attribute 'ariaHasPopup' with
// type string (idl: DOMString).
func (_this *Element) SetAriaHasPopup(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaHasPopup", input)
}

// AriaHidden returning attribute 'ariaHidden' with
// type string (idl: DOMString).
func (_this *Element) AriaHidden() *string {
	var ret *string
	value := _this.jsValue.Get("ariaHidden")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaHidden setting attribute 'ariaHidden' with
// type string (idl: DOMString).
func (_this *Element) SetAriaHidden(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaHidden", input)
}

// AriaInvalid returning attribute 'ariaInvalid' with
// type string (idl: DOMString).
func (_this *Element) AriaInvalid() *string {
	var ret *string
	value := _this.jsValue.Get("ariaInvalid")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaInvalid setting attribute 'ariaInvalid' with
// type string (idl: DOMString).
func (_this *Element) SetAriaInvalid(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaInvalid", input)
}

// AriaKeyShortcuts returning attribute 'ariaKeyShortcuts' with
// type string (idl: DOMString).
func (_this *Element) AriaKeyShortcuts() *string {
	var ret *string
	value := _this.jsValue.Get("ariaKeyShortcuts")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaKeyShortcuts setting attribute 'ariaKeyShortcuts' with
// type string (idl: DOMString).
func (_this *Element) SetAriaKeyShortcuts(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaKeyShortcuts", input)
}

// AriaLabel returning attribute 'ariaLabel' with
// type string (idl: DOMString).
func (_this *Element) AriaLabel() *string {
	var ret *string
	value := _this.jsValue.Get("ariaLabel")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaLabel setting attribute 'ariaLabel' with
// type string (idl: DOMString).
func (_this *Element) SetAriaLabel(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaLabel", input)
}

// AriaLabelledBy returning attribute 'ariaLabelledBy' with
// type string (idl: DOMString).
func (_this *Element) AriaLabelledBy() *string {
	var ret *string
	value := _this.jsValue.Get("ariaLabelledBy")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaLabelledBy setting attribute 'ariaLabelledBy' with
// type string (idl: DOMString).
func (_this *Element) SetAriaLabelledBy(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaLabelledBy", input)
}

// AriaLevel returning attribute 'ariaLevel' with
// type string (idl: DOMString).
func (_this *Element) AriaLevel() *string {
	var ret *string
	value := _this.jsValue.Get("ariaLevel")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaLevel setting attribute 'ariaLevel' with
// type string (idl: DOMString).
func (_this *Element) SetAriaLevel(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaLevel", input)
}

// AriaLive returning attribute 'ariaLive' with
// type string (idl: DOMString).
func (_this *Element) AriaLive() *string {
	var ret *string
	value := _this.jsValue.Get("ariaLive")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaLive setting attribute 'ariaLive' with
// type string (idl: DOMString).
func (_this *Element) SetAriaLive(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaLive", input)
}

// AriaModal returning attribute 'ariaModal' with
// type string (idl: DOMString).
func (_this *Element) AriaModal() *string {
	var ret *string
	value := _this.jsValue.Get("ariaModal")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaModal setting attribute 'ariaModal' with
// type string (idl: DOMString).
func (_this *Element) SetAriaModal(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaModal", input)
}

// AriaMultiLine returning attribute 'ariaMultiLine' with
// type string (idl: DOMString).
func (_this *Element) AriaMultiLine() *string {
	var ret *string
	value := _this.jsValue.Get("ariaMultiLine")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaMultiLine setting attribute 'ariaMultiLine' with
// type string (idl: DOMString).
func (_this *Element) SetAriaMultiLine(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaMultiLine", input)
}

// AriaMultiSelectable returning attribute 'ariaMultiSelectable' with
// type string (idl: DOMString).
func (_this *Element) AriaMultiSelectable() *string {
	var ret *string
	value := _this.jsValue.Get("ariaMultiSelectable")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaMultiSelectable setting attribute 'ariaMultiSelectable' with
// type string (idl: DOMString).
func (_this *Element) SetAriaMultiSelectable(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaMultiSelectable", input)
}

// AriaOrientation returning attribute 'ariaOrientation' with
// type string (idl: DOMString).
func (_this *Element) AriaOrientation() *string {
	var ret *string
	value := _this.jsValue.Get("ariaOrientation")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaOrientation setting attribute 'ariaOrientation' with
// type string (idl: DOMString).
func (_this *Element) SetAriaOrientation(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaOrientation", input)
}

// AriaOwns returning attribute 'ariaOwns' with
// type string (idl: DOMString).
func (_this *Element) AriaOwns() *string {
	var ret *string
	value := _this.jsValue.Get("ariaOwns")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaOwns setting attribute 'ariaOwns' with
// type string (idl: DOMString).
func (_this *Element) SetAriaOwns(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaOwns", input)
}

// AriaPlaceholder returning attribute 'ariaPlaceholder' with
// type string (idl: DOMString).
func (_this *Element) AriaPlaceholder() *string {
	var ret *string
	value := _this.jsValue.Get("ariaPlaceholder")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaPlaceholder setting attribute 'ariaPlaceholder' with
// type string (idl: DOMString).
func (_this *Element) SetAriaPlaceholder(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaPlaceholder", input)
}

// AriaPosInSet returning attribute 'ariaPosInSet' with
// type string (idl: DOMString).
func (_this *Element) AriaPosInSet() *string {
	var ret *string
	value := _this.jsValue.Get("ariaPosInSet")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaPosInSet setting attribute 'ariaPosInSet' with
// type string (idl: DOMString).
func (_this *Element) SetAriaPosInSet(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaPosInSet", input)
}

// AriaPressed returning attribute 'ariaPressed' with
// type string (idl: DOMString).
func (_this *Element) AriaPressed() *string {
	var ret *string
	value := _this.jsValue.Get("ariaPressed")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaPressed setting attribute 'ariaPressed' with
// type string (idl: DOMString).
func (_this *Element) SetAriaPressed(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaPressed", input)
}

// AriaReadOnly returning attribute 'ariaReadOnly' with
// type string (idl: DOMString).
func (_this *Element) AriaReadOnly() *string {
	var ret *string
	value := _this.jsValue.Get("ariaReadOnly")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaReadOnly setting attribute 'ariaReadOnly' with
// type string (idl: DOMString).
func (_this *Element) SetAriaReadOnly(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaReadOnly", input)
}

// AriaRelevant returning attribute 'ariaRelevant' with
// type string (idl: DOMString).
func (_this *Element) AriaRelevant() *string {
	var ret *string
	value := _this.jsValue.Get("ariaRelevant")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaRelevant setting attribute 'ariaRelevant' with
// type string (idl: DOMString).
func (_this *Element) SetAriaRelevant(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaRelevant", input)
}

// AriaRequired returning attribute 'ariaRequired' with
// type string (idl: DOMString).
func (_this *Element) AriaRequired() *string {
	var ret *string
	value := _this.jsValue.Get("ariaRequired")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaRequired setting attribute 'ariaRequired' with
// type string (idl: DOMString).
func (_this *Element) SetAriaRequired(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaRequired", input)
}

// AriaRoleDescription returning attribute 'ariaRoleDescription' with
// type string (idl: DOMString).
func (_this *Element) AriaRoleDescription() *string {
	var ret *string
	value := _this.jsValue.Get("ariaRoleDescription")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaRoleDescription setting attribute 'ariaRoleDescription' with
// type string (idl: DOMString).
func (_this *Element) SetAriaRoleDescription(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaRoleDescription", input)
}

// AriaRowCount returning attribute 'ariaRowCount' with
// type string (idl: DOMString).
func (_this *Element) AriaRowCount() *string {
	var ret *string
	value := _this.jsValue.Get("ariaRowCount")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaRowCount setting attribute 'ariaRowCount' with
// type string (idl: DOMString).
func (_this *Element) SetAriaRowCount(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaRowCount", input)
}

// AriaRowIndex returning attribute 'ariaRowIndex' with
// type string (idl: DOMString).
func (_this *Element) AriaRowIndex() *string {
	var ret *string
	value := _this.jsValue.Get("ariaRowIndex")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaRowIndex setting attribute 'ariaRowIndex' with
// type string (idl: DOMString).
func (_this *Element) SetAriaRowIndex(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaRowIndex", input)
}

// AriaRowSpan returning attribute 'ariaRowSpan' with
// type string (idl: DOMString).
func (_this *Element) AriaRowSpan() *string {
	var ret *string
	value := _this.jsValue.Get("ariaRowSpan")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaRowSpan setting attribute 'ariaRowSpan' with
// type string (idl: DOMString).
func (_this *Element) SetAriaRowSpan(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaRowSpan", input)
}

// AriaSelected returning attribute 'ariaSelected' with
// type string (idl: DOMString).
func (_this *Element) AriaSelected() *string {
	var ret *string
	value := _this.jsValue.Get("ariaSelected")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaSelected setting attribute 'ariaSelected' with
// type string (idl: DOMString).
func (_this *Element) SetAriaSelected(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaSelected", input)
}

// AriaSetSize returning attribute 'ariaSetSize' with
// type string (idl: DOMString).
func (_this *Element) AriaSetSize() *string {
	var ret *string
	value := _this.jsValue.Get("ariaSetSize")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaSetSize setting attribute 'ariaSetSize' with
// type string (idl: DOMString).
func (_this *Element) SetAriaSetSize(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaSetSize", input)
}

// AriaSort returning attribute 'ariaSort' with
// type string (idl: DOMString).
func (_this *Element) AriaSort() *string {
	var ret *string
	value := _this.jsValue.Get("ariaSort")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaSort setting attribute 'ariaSort' with
// type string (idl: DOMString).
func (_this *Element) SetAriaSort(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaSort", input)
}

// AriaValueMax returning attribute 'ariaValueMax' with
// type string (idl: DOMString).
func (_this *Element) AriaValueMax() *string {
	var ret *string
	value := _this.jsValue.Get("ariaValueMax")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaValueMax setting attribute 'ariaValueMax' with
// type string (idl: DOMString).
func (_this *Element) SetAriaValueMax(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaValueMax", input)
}

// AriaValueMin returning attribute 'ariaValueMin' with
// type string (idl: DOMString).
func (_this *Element) AriaValueMin() *string {
	var ret *string
	value := _this.jsValue.Get("ariaValueMin")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaValueMin setting attribute 'ariaValueMin' with
// type string (idl: DOMString).
func (_this *Element) SetAriaValueMin(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaValueMin", input)
}

// AriaValueNow returning attribute 'ariaValueNow' with
// type string (idl: DOMString).
func (_this *Element) AriaValueNow() *string {
	var ret *string
	value := _this.jsValue.Get("ariaValueNow")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaValueNow setting attribute 'ariaValueNow' with
// type string (idl: DOMString).
func (_this *Element) SetAriaValueNow(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaValueNow", input)
}

// AriaValueText returning attribute 'ariaValueText' with
// type string (idl: DOMString).
func (_this *Element) AriaValueText() *string {
	var ret *string
	value := _this.jsValue.Get("ariaValueText")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		__tmp := (value).String()
		ret = &__tmp
	}
	return ret
}

// SetAriaValueText setting attribute 'ariaValueText' with
// type string (idl: DOMString).
func (_this *Element) SetAriaValueText(value *string) {
	var input interface{}
	if value != nil {
		input = *(value)
	} else {
		input = nil
	}
	_this.jsValue.Set("ariaValueText", input)
}

// event attribute: Event
func eventFuncElement_Event(listener func(event *Event, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *Event
		value := args[0]
		incoming := value.Get("target")
		ret = EventFromJS(value)
		src := ElementFromJS(incoming)
		listener(ret, src)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddFullscreenChange is adding doing AddEventListener for 'FullscreenChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Element) AddEventFullscreenChange(listener func(event *Event, currentTarget *Element)) js.Func {
	cb := eventFuncElement_Event(listener)
	_this.jsValue.Call("addEventListener", "fullscreenchange", cb)
	return cb
}

// SetOnFullscreenChange is assigning a function to 'onfullscreenchange'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Element) SetOnFullscreenChange(listener func(event *Event, currentTarget *Element)) js.Func {
	cb := eventFuncElement_Event(listener)
	_this.jsValue.Set("onfullscreenchange", cb)
	return cb
}

// AddFullscreenError is adding doing AddEventListener for 'FullscreenError' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Element) AddEventFullscreenError(listener func(event *Event, currentTarget *Element)) js.Func {
	cb := eventFuncElement_Event(listener)
	_this.jsValue.Call("addEventListener", "fullscreenerror", cb)
	return cb
}

// SetOnFullscreenError is assigning a function to 'onfullscreenerror'. This
// This method is returning allocated javascript function that need to be released.
func (_this *Element) SetOnFullscreenError(listener func(event *Event, currentTarget *Element)) js.Func {
	cb := eventFuncElement_Event(listener)
	_this.jsValue.Set("onfullscreenerror", cb)
	return cb
}

/****************************************************************************
* Element's attributes
*****************************************************************************/

func (_this *Element) HasAttributes() (_result bool) {
	var _args [0]interface{}
	_returned := _this.jsValue.Call("hasAttributes", _args[0:0]...)
	return (_returned).Bool()
}

func (_this *Element) HasAttribute(qualifiedName string) (_result bool) {
	var _args [1]interface{}
	_args[0] = qualifiedName
	_returned := _this.jsValue.Call("hasAttribute", _args[0:1]...)
	return (_returned).Bool()
}

func (_this *Element) GetAttributeNames() (_result []string) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("getAttributeNames", _args[0:_end]...)
	var (
		_converted []string // javascript: sequence<DOMString> _what_return_name
	)
	__length0 := _returned.Length()
	__array0 := make([]string, __length0)
	for __idx0 := 0; __idx0 < __length0; __idx0++ {
		var __seq_out0 string
		__seq_in0 := _returned.Index(__idx0)
		__seq_out0 = (__seq_in0).String()
		__array0[__idx0] = __seq_out0
	}
	_converted = __array0
	_result = _converted
	return
}

func (_this *Element) GetAttribute(qualifiedName string) (_result *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := qualifiedName
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("getAttribute", _args[0:_end]...)
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

func (_this *Element) GetAttributeNode(qualifiedName string) (_result *Attr) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := qualifiedName
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("getAttributeNode", _args[0:_end]...)
	var (
		_converted *Attr // javascript: Attr _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = AttrFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *Element) SetAttribute(qualifiedName string, value string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := qualifiedName
	_args[0] = _p0
	_end++
	_p1 := value
	_args[1] = _p1
	_end++
	_this.jsValue.Call("setAttribute", _args[0:_end]...)
	
}

func (_this *Element) RemoveAttribute(qualifiedName string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := qualifiedName
	_args[0] = _p0
	_end++
	_this.jsValue.Call("removeAttribute", _args[0:_end]...)
	
}

func (_this *Element) ToggleAttribute(qualifiedName string, force *bool) (_result bool) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := qualifiedName
	_args[0] = _p0
	_end++
	if force != nil {

		var _p1 interface{}
		if force != nil {
			_p1 = *(force)
		} else {
			_p1 = nil
		}
		_args[1] = _p1
		_end++
	}
	_returned := _this.jsValue.Call("toggleAttribute", _args[0:_end]...)
	return (_returned).Bool()
}

func (_this *Element) SetAttributeNode(attr *Attr) (_result *Attr) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := attr.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("setAttributeNode", _args[0:_end]...)
	var (
		_converted *Attr // javascript: Attr _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = AttrFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *Element) Closest(selectors string) (_result *Element) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := selectors
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("closest", _args[0:_end]...)
	return ElementFromJS(_returned)
}

func (_this *Element) Matches(selectors string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := selectors
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("matches", _args[0:_end]...)
	return (_returned).Bool()
}

func (_this *Element) WebkitMatchesSelector(selectors string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := selectors
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("webkitMatchesSelector", _args[0:_end]...)
	return (_returned).Bool()
}

func (_this *Element) GetElementsByTagName(qualifiedName string) (_result *HTMLCollection) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := qualifiedName
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("getElementsByTagName", _args[0:_end]...)
	return HTMLCollectionFromJS(_returned)
}

func (_this *Element) GetElementsByClassName(classNames string) (_result *HTMLCollection) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := classNames
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("getElementsByClassName", _args[0:_end]...)
	return HTMLCollectionFromJS(_returned)
}

func (_this *Element) InsertAdjacentElement(where string, element *Element) (_result *Element) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := where
	_args[0] = _p0
	_end++
	_p1 := element.JSValue()
	_args[1] = _p1
	_end++
	_returned := _this.jsValue.Call("insertAdjacentElement", _args[0:_end]...)
	var (
		_converted *Element // javascript: Element _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = ElementFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *Element) InsertAdjacentText(where string, data string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := where
	_args[0] = _p0
	_end++
	_p1 := data
	_args[1] = _p1
	_end++
	_this.jsValue.Call("insertAdjacentText", _args[0:_end]...)
}

func (_this *Element) InsertAdjacentHTML(position string, text string) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := position
	_args[0] = _p0
	_end++
	_p1 := text
	_args[1] = _p1
	_end++
	_this.jsValue.Call("insertAdjacentHTML", _args[0:_end]...)
}

func (_this *Element) GetBoundingClientRect() (_result *DOMRect) {
	var (
		_args [0]interface{}
		_end  int
	)
	_returned := _this.jsValue.Call("getBoundingClientRect", _args[0:_end]...)
	return DOMRectFromJS(_returned)
}

func (_this *Element) ScrollIntoView() {
	var (
		_args [1]interface{}
		_end  int
	)
	_this.jsValue.Call("scrollIntoView", _args[0:_end]...)
}

func (_this *Element) ScrollXY(x float64, y float64) {
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
}

func (_this *Element) ScrollToXY(x float64, y float64) {
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
}

func (_this *Element) ScrollByXY(x float64, y float64) {
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
}

func (_this *Element) SetPointerCapture(pointerId int) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := pointerId
	_args[0] = _p0
	_end++
	_this.jsValue.Call("setPointerCapture", _args[0:_end]...)
}

func (_this *Element) ReleasePointerCapture(pointerId int) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := pointerId
	_args[0] = _p0
	_end++
	_this.jsValue.Call("releasePointerCapture", _args[0:_end]...)
}

func (_this *Element) HasPointerCapture(pointerId int) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := pointerId
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("hasPointerCapture", _args[0:_end]...)
	return (_returned).Bool()
}

func (_this *Element) RequestPointerLock() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("requestPointerLock", _args[0:_end]...)
}

func (_this *Element) Prepend(nodes ...*Union) {
	var (
		_args []interface{} = make([]interface{}, 0+len(nodes))
		_end  int
	)
	for _, __in := range nodes {
		__out := __in.JSValue()
		_args[_end] = __out
		_end++
	}
	_this.jsValue.Call("prepend", _args[0:_end]...)
}

func (_this *Element) Append(nodes ...*Union) {
	var (
		_args []interface{} = make([]interface{}, 0+len(nodes))
		_end  int
	)
	for _, __in := range nodes {
		__out := __in.JSValue()
		_args[_end] = __out
		_end++
	}
	_this.jsValue.Call("append", _args[0:_end]...)
}

func (_this *Element) QuerySelector(selectors string) (_result *Element) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := selectors
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("querySelector", _args[0:_end]...)
	return ElementFromJS(_returned)
}

func (_this *Element) QuerySelectorAll(selectors string) (_result *NodeList) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := selectors
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("querySelectorAll", _args[0:_end]...)
	return NodeListFromJS(_returned)
}

func (_this *Element) Before(nodes ...*Union) {
	var (
		_args []interface{} = make([]interface{}, 0+len(nodes))
		_end  int
	)
	for _, __in := range nodes {
		__out := __in.JSValue()
		_args[_end] = __out
		_end++
	}
	_this.jsValue.Call("before", _args[0:_end]...)
}

func (_this *Element) After(nodes ...*Union) {
	var (
		_args []interface{} = make([]interface{}, 0+len(nodes))
		_end  int
	)
	for _, __in := range nodes {
		__out := __in.JSValue()
		_args[_end] = __out
		_end++
	}
	_this.jsValue.Call("after", _args[0:_end]...)
}

func (_this *Element) Remove() {
	var (
		_args [0]interface{}
		_end  int
	)
	_this.jsValue.Call("remove", _args[0:_end]...)
}
