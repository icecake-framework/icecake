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

// MakeElementFromJS is casting a js.Value into Element.
func MakeElementFromJS(value js.Value) *Element {
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

// Returns the namespace prefix of the specified element, or null if no prefix is specified.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/prefix
func (_this *Element) Prefix() string {
	value := _this.jsValue.Get("prefix")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		return ""
	}
	return (value).String()
}

// LocalName returns the local part of the qualified name of an element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/localName
func (_this *Element) LocalName() string {
	value := _this.jsValue.Get("localName")
	return (value).String()
}

// TagName returns the tag name of the element on which it's called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/tagName
func (_this *Element) TagName() string {
	value := _this.jsValue.Get("tagName")
	return (value).String()
}

// Id rrepresents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (_this *Element) Id() string {
	value := _this.jsValue.Get("id")
	return (value).String()
}

// Id rrepresents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (_this *Element) SetId(value string) {
	_this.jsValue.Set("id", value)
}

// SetClassName gets and sets the value of the class attribute of the specified element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (_this *Element) ClassName() string {
	value := _this.jsValue.Get("className")
	return (value).String()
}

// SetClassName gets and sets the value of the class attribute of the specified element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (_this *Element) SetClassName(value string) {
	_this.jsValue.Set("className", value)
}

// ClassList returns a live DOMTokenList collection of the class attributes of the element.
// This can then be used to manipulate the class list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/classList
func (_this *Element) ClassList() *TokenList {
	value := _this.jsValue.Get("classList")
	return MakeTokenListFromJS(value)
}

// Attributes returns a live collection of all attribute nodes registered to the specified node.
// It is a NamedNodeMap, not an Array, so it has no Array methods and the Attr nodes' indexes may differ among browsers.
// To be more specific, attributes is a key/value pair of strings that represents any information regarding that attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
func (_this *Element) Attributes() *Attributes {
	value := _this.jsValue.Get("attributes")
	return MakeAttributesFromNamedNodeMapJS(value)
}

// InnerHTML ets or sets the HTML or XML markup contained within the element.
//
// To insert the HTML into the document rather than replace the contents of an element, use the method insertAdjacentHTML().
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (_this *Element) InnerHTML() string {
	value := _this.jsValue.Get("innerHTML")
	return (value).String()
}

// InnerHTML ets or sets the HTML or XML markup contained within the element.
//
// To insert the HTML into the document rather than replace the contents of an element, use the method insertAdjacentHTML().
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (_this *Element) SetInnerHTML(value string) {
	input := value
	_this.jsValue.Set("innerHTML", input)
}

// OuterHTML gets the serialized HTML fragment describing the element including its descendants.
// It can also be set to replace the element with nodes parsed from the given string.
//
// To only obtain the HTML representation of the contents of an element,
// or to replace the contents of an element, use the innerHTML property instead.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
func (_this *Element) OuterHTML() string {
	value := _this.jsValue.Get("outerHTML")
	return (value).String()
}

// OuterHTML gets the serialized HTML fragment describing the element including its descendants.
// It can also be set to replace the element with nodes parsed from the given string.
//
// To only obtain the HTML representation of the contents of an element,
// or to replace the contents of an element, use the innerHTML property instead.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
func (_this *Element) SetOuterHTML(value string) {
	_this.jsValue.Set("outerHTML", value)
}

// ScrollTop gets or sets the number of pixels that an element's content is scrolled vertically.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollTop
func (_this *Element) ScrollTop() float64 {
	value := _this.jsValue.Get("scrollTop")
	return (value).Float()
}

// ScrollTop gets or sets the number of pixels that an element's content is scrolled vertically.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollTop
func (_this *Element) SetScrollTop(value float64) {
	input := value
	_this.jsValue.Set("scrollTop", input)
}

// ScrollLeft  gets or sets the number of pixels that an element's content is scrolled from its left edge.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
func (_this *Element) ScrollLeft() float64 {
	value := _this.jsValue.Get("scrollLeft")
	return (value).Float()
}

// ScrollLeft  gets or sets the number of pixels that an element's content is scrolled from its left edge.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
func (_this *Element) SetScrollLeft(value float64) {
	_this.jsValue.Set("scrollLeft", value)
}

// ScrollWidth s a measurement of the width of an element's content, including content not visible on the screen due to overflow.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollWidth
func (_this *Element) ScrollWidth() int {
	var ret int
	value := _this.jsValue.Get("scrollWidth")
	ret = (value).Int()
	return ret
}

// ScrollHeight  is a measurement of the height of an element's content, including content not visible on the screen due to overflow.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollHeight
func (_this *Element) ScrollHeight() int {
	var ret int
	value := _this.jsValue.Get("scrollHeight")
	ret = (value).Int()
	return ret
}

// ClientTop The width of the top border of an element in pixels.
//
// Note: This property will round the value to an integer. If you need a fractional value, use element.getBoundingClientRect().
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/clientTop
func (_this *Element) ClientTop() int {
	value := _this.jsValue.Get("clientTop")
	return (value).Int()
}

// ClientLeft the width of the left border of an element in pixels.
// It includes the width of the vertical scrollbar if the text direction of the element is right-to-left and if there is an overflow causing a left vertical scrollbar to be rendered.
// clientLeft does not include the left margin or the left padding. clientLeft is read-only.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/clientLeft
func (_this *Element) ClientLeft() int {
	var ret int
	value := _this.jsValue.Get("clientLeft")
	ret = (value).Int()
	return ret
}

// ClientWidth The Element.clientWidth property is zero for inline elements and elements with no CSS;
// otherwise, it's the inner width of an element in pixels. It includes padding but excludes borders, margins, and vertical scrollbars (if present).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/clientWidth
func (_this *Element) ClientWidth() int {
	var ret int
	value := _this.jsValue.Get("clientWidth")
	ret = (value).Int()
	return ret
}

// ClientHeight is zero for elements with no CSS or inline layout boxes;
// otherwise, it's the inner height of an element in pixels. It includes padding but excludes borders, margins, and horizontal scrollbars (if present).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/clientHeight
func (_this *Element) ClientHeight() int {
	var ret int
	value := _this.jsValue.Get("clientHeight")
	ret = (value).Int()
	return ret
}

// Children returns a live HTMLCollection which contains all of the child elements of the element upon which it was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/children
func (_this *Element) Children() *HTMLCollection {
	var ret *HTMLCollection
	value := _this.jsValue.Get("children")
	ret = HTMLCollectionFromJS(value)
	return ret
}

// FirstElementChild  returns an element's first child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/firstElementChild
func (_this *Element) FirstElementChild() *Element {
	value := _this.jsValue.Get("firstElementChild")
	return MakeElementFromJS(value)
}

// LastElementChild returns an element's last child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/lastElementChild
func (_this *Element) LastElementChild() *Element {
	value := _this.jsValue.Get("lastElementChild")
	return MakeElementFromJS(value)
}

// ChildElementCount returns the number of child elements of this element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/childElementCount
func (_this *Element) ChildElementCount() uint {
	var ret uint
	value := _this.jsValue.Get("childElementCount")
	ret = (uint)((value).Int())
	return ret
}

// PreviousElementSibling returns the Element immediately prior to the specified one in its parent's children list,
// or null if the specified element is the first one in the list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/previousElementSibling
func (_this *Element) PreviousElementSibling() *Element {
	value := _this.jsValue.Get("previousElementSibling")
	return MakeElementFromJS(value)
}

// NextElementSibling returns the element immediately following the specified one in its parent's children list, or null if the specified element is the last one in the list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/nextElementSibling
func (_this *Element) NextElementSibling() *Element {
	value := _this.jsValue.Get("nextElementSibling")
	return MakeElementFromJS(value)
}

// Role get a,d set the attribute 'role'
func (_this *Element) Role() string {
	value := _this.jsValue.Get("role")
	if value.Type() != js.TypeNull && value.Type() != js.TypeUndefined {
		return ""
	}
	return (value).String()
}

// Role get a,d set the attribute 'role'
func (_this *Element) SetRole(value string) {
	var input interface{}
	if value != "" {
		input = value
	} else {
		input = ""
	}
	_this.jsValue.Set("role", input)
}

/****************************************************************************
* Element's events
*****************************************************************************/

// event attribute: Event
func eventFuncElement_Event(listener func(event *Event, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var ret *Event
		value := args[0]
		incoming := value.Get("target")
		ret = MakeEventFromJS(value)
		src := MakeElementFromJS(incoming)
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

// AddFullscreenError is adding doing AddEventListener for 'FullscreenError' on target.
// This method is returning allocated javascript function that need to be released.
func (_this *Element) AddEventFullscreenError(listener func(event *Event, currentTarget *Element)) js.Func {
	cb := eventFuncElement_Event(listener)
	_this.jsValue.Call("addEventListener", "fullscreenerror", cb)
	return cb
}

/****************************************************************************
* Element's attributes
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Element/hasAttributes
func (_this *Element) HasAttributes() bool {
	var _args [0]interface{}
	_returned := _this.jsValue.Call("hasAttributes", _args[0:0]...)
	return (_returned).Bool()
}

// https://developer.mozilla.org/en-US/docs/Web/API/Element/hasAttribute
func (_this *Element) HasAttribute(qualifiedName string) (_result bool) {
	var _args [1]interface{}
	_args[0] = qualifiedName
	_returned := _this.jsValue.Call("hasAttribute", _args[0:1]...)
	return (_returned).Bool()
}

func (_this *Element) GetAttributeNode(qualifiedName string) (_result *Attribute) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := qualifiedName
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("getAttributeNode", _args[0:_end]...)
	var (
		_converted *Attribute // javascript: Attr _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = MakeAttributeFromJS(_returned)
	}
	_result = _converted
	return
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

// Traverses the element and its parents (heading toward the document root) until it finds a node that matches the specified CSS selector.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/closest
func (_this *Element) Closest(selectors string) (_result *Element) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := selectors
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("closest", _args[0:_end]...)
	return MakeElementFromJS(_returned)
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
		_converted = MakeElementFromJS(_returned)
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

func (_this *Element) GetBoundingClientRect() (_result *Rect) {
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
	return MakeElementFromJS(_returned)
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
	return MakeNodeListFromJS(_returned)
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

// IsDefined returns true if the Element is not nil AND it's type is not TypeNull and not TypeUndefined
func (_this *Element) IsDefined() bool {
	if _this == nil || _this.jsValue.Type() != js.TypeNull && _this.jsValue.Type() != js.TypeUndefined {
		return false
	}
	return true
}
