package browser

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
	return MakeAttributesFromNamedNodeMapJS(value, _this)
}

// InnerHTML gets or sets the HTML or XML markup contained within the element.
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
// When writing to innerHTML, it will overwrite the content of the source element.
// That means the HTML has to be loaded and re-parsed. This is not very efficient especially when using inside loops.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (_this *Element) SetInnerHTML(value string) (_ret *Element) {
	input := value
	_this.jsValue.Set("innerHTML", input)
	return _this
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
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollWidth
func (_this *Element) ScrollRect() (_ret Rect) {
	_ret.X = _this.jsValue.Get("scrollLeft").Float()
	_ret.Y = _this.jsValue.Get("scrollTop").Float()
	_ret.Width = _this.jsValue.Get("scrollWidth").Float()
	_ret.Height = _this.jsValue.Get("scrollHeight").Float()
	return _ret
}

// ClientRect returns border coordinates of an element in pixels.
//
// Note: This property will round the value to an integer. If you need a fractional value, use element.getBoundingClientRect().
//
//   - https://developer.mozilla.org/en-US/docs/Web/API/Element/clientTop
//   - https://developer.mozilla.org/en-US/docs/Web/API/Element/clientLeft
func (_this *Element) ClientRect() (_ret Rect) {
	_ret = Rect{}
	_ret.X = float64(_this.jsValue.Get("clientLeft").Int())
	_ret.Y = float64(_this.jsValue.Get("clientTop").Int())
	_ret.Width = float64(_this.jsValue.Get("clientWidth").Int())
	_ret.Height = float64(_this.jsValue.Get("clientHeight").Int())
	return _ret
}

// GetBoundingClientRect eturns a DOMRect object providing information about the size of an element and its position relative to the viewport.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
func (_this *Element) BoundingClientRect() (_result *Rect) {
	_returned := _this.jsValue.Call("getBoundingClientRect")
	return DOMRectFromJS(_returned)
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
* Element's methods
*****************************************************************************/

// Traverses the element and its parents (heading toward the document root) until
// it finds a node that matches the specified CSS selector.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/closest
func (_this *Element) Closest(selectors string) (_result *Element) {
	_returned := _this.jsValue.Call("closest", selectors)
	return MakeElementFromJS(_returned)
}

// Matches tests whether the element would be selected by the specified CSS selector.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/matches
func (_this *Element) Matches(selectors string) (_result bool) {
	_returned := _this.jsValue.Call("matches", selectors)
	return (_returned).Bool()
}

// QuerySelector returns the first element that is a descendant of the element on which it is invoked
// that matches the specified group of selectors.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
func (_this *Element) QuerySelector(selectors string) (_result *Element) {
	_returned := _this.jsValue.Call("querySelector", selectors)
	return MakeElementFromJS(_returned)
}

// QuerySelectorAll eturns a static (not live) NodeList representing a list of elements matching
// the specified group of selectors which are descendants of the element on which the method was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
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

// GetElementsByTagName returns a live HTMLCollection of elements with the given tag name.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName
func (_this *Element) GetElementsByTagName(qualifiedName string) (_result *HTMLCollection) {
	_returned := _this.jsValue.Call("getElementsByTagName", qualifiedName)
	return HTMLCollectionFromJS(_returned)
}

// GetElementsByClassName returns a live HTMLCollection which contains every descendant element which has the specified class name or names.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByClassName
func (_this *Element) GetElementsByClassName(classNames string) (_result *HTMLCollection) {
	_returned := _this.jsValue.Call("getElementsByClassName", classNames)
	return HTMLCollectionFromJS(_returned)
}

// ScrollIntoView scrolls the element's ancestor containers such that the element on which scrollIntoView() is called is visible to the user.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollIntoView
func (_this *Element) ScrollIntoView() {
	_this.jsValue.Call("scrollIntoView")
}

type WhereInsert string

const (
	WI_BEFOREBEGIN WhereInsert = "beforebegin" //  Before the Element itself.
	WI_AFTERBEGIN  WhereInsert = "afterbegin"  // Just inside the element, before its first child.
	WI_BEFOREEND   WhereInsert = "beforeend"   // Just inside the element, after its last child.
	WI_AFTEREND    WhereInsert = "afterend"    // After the element itself.
)

// InsertAdjacentElement inserts a given element node at a given position relative to the element it is invoked upon.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentElement
func (_this *Element) InsertAdjacentElement(where WhereInsert, element *Element) (_result *Element) {
	_returned := _this.jsValue.Call("insertAdjacentElement", where, element.JSValue())
	return MakeElementFromJS(_returned)
}

// InsertAdjacentText given a relative position and a string, inserts a new text node at the given position relative to the element it is called from.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentText
func (_this *Element) InsertAdjacentText(where WhereInsert, text string) {
	_this.jsValue.Call("insertAdjacentText", where, text)
}

// InsertAdjacentHTML parses the specified text as HTML or XML and inserts the resulting nodes into the DOM tree at a specified position.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentHTML
func (_this *Element) InsertAdjacentHTML(where WhereInsert, text string) {
	_this.jsValue.Call("insertAdjacentHTML", where, text)
}

// Prepend inserts a set of Node objects or string objects before the first child of the Element.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/prepend
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

// Append inserts a set of Node objects or string objects after the last child of the Element.
// String objects are inserted as equivalent Text nodes.
//
// This method is supported by all browsers and is a much cleaner way of inserting nodes, text, data, etc. into the DOM.
//
// Allows you to also append string objects, whereas Node.appendChild() only accepts Node objects.
//
// Has no return value, whereas Node.appendChild() returns the appended Node object.
//
// Can append several nodes and strings, whereas Node.appendChild() can only append one node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/append
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
	if _this == nil || _this.jsValue.Type() == js.TypeNull || _this.jsValue.Type() == js.TypeUndefined {
		return false
	}
	return true
}
