package dom

import (
	"syscall/js"

	"github.com/sunraylab/icecake/pkg/lib"
)

/****************************************************************************
* enum used by Element
*****************************************************************************/

type WhereInsert string

const (
	WI_BEFOREBEGIN WhereInsert = "beforebegin" // Before the Element itself.
	WI_INSIDEFIRST WhereInsert = "afterbegin"  // Just inside the element, before its first child.
	WI_INSIDELAST  WhereInsert = "beforeend"   // Just inside the element, after its last child.
	WI_AFTEREND    WhereInsert = "afterend"    // After the element itself.
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

// NewElementFromJS is casting a js.Value into Element.
func NewElementFromJS(value js.Value) *Element {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Element{}
	ret.jsValue = value
	return ret
}

// Remove removes the element from the DOM.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/remove
func (_elem *Element) Remove() {
	_elem.jsValue.Call("remove")
}

/****************************************************************************
* Element's Properties & Methods
*****************************************************************************/

// TagName returns the tag name of the element on which it's called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/tagName
func (_elem *Element) TagName() string {
	return _elem.jsValue.Get("tagName").String()
}

// Id rrepresents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (_elem *Element) Id() string {
	return _elem.jsValue.Get("id").String()
}

// Id rrepresents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (_elem *Element) SetId(value string) *Element {
	_elem.jsValue.Set("id", value)
	return _elem
}

// SetClassName gets and sets the value of the class attribute of the specified element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (_elem *Element) ClassName() string {
	value := _elem.jsValue.Get("className")
	return (value).String()
}

// SetClassName gets and sets the value of the class attribute of the specified element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (_elem *Element) SetClassName(value string) *Element {
	_elem.jsValue.Set("className", value)
	return _elem
}

// ClassList returns a live DOMTokenList collection of the class attributes of the element.
// This can then be used to manipulate the class list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/classList
func (_elem *Element) ClassList() *TokenList {
	value := _elem.jsValue.Get("classList")
	return NewTokenListFromJS(value)
}

// Attributes returns a live collection of all attribute nodes registered to the specified node.
// It is a NamedNodeMap, not an Array, so it has no Array methods and the Attr nodes' indexes may differ among browsers.
// To be more specific, attributes is a key/value pair of strings that represents any information regarding that attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
func (_elem *Element) Attributes() *Attributes {
	value := _elem.jsValue.Get("attributes")
	return NewAttributesFromJSNodeMap(value, _elem)
}

// InnerHTML gets or sets the HTML or XML markup contained within the element.
//
// To insert the HTML into the document rather than replace the contents of an element, use the method insertAdjacentHTML().
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (_elem *Element) InnerHTML() string {
	value := _elem.jsValue.Get("innerHTML")
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
func (_elem *Element) SetInnerHTML(value string) *Element {
	_elem.jsValue.Set("innerHTML", value)
	return _elem
}

// OuterHTML gets the serialized HTML fragment describing the element including its descendants.
// It can also be set to replace the element with nodes parsed from the given string.
//
// To only obtain the HTML representation of the contents of an element,
// or to replace the contents of an element, use the innerHTML property instead.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
func (_elem *Element) OuterHTML() string {
	return _elem.jsValue.Get("outerHTML").String()
}

// OuterHTML gets the serialized HTML fragment describing the element including its descendants.
// It can also be set to replace the element with nodes parsed from the given string.
//
// To only obtain the HTML representation of the contents of an element,
// or to replace the contents of an element, use the innerHTML property instead.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
func (_elem *Element) SetOuterHTML(value string) *Element {
	_elem.jsValue.Set("outerHTML", value)
	return _elem
}

// ScrollTop gets or sets the number of pixels that an element's content is scrolled vertically.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollTop
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollWidth
func (_elem *Element) ScrollRect() (_ret lib.Rect) {
	_ret.X = _elem.jsValue.Get("scrollLeft").Float()
	_ret.Y = _elem.jsValue.Get("scrollTop").Float()
	_ret.Width = _elem.jsValue.Get("scrollWidth").Float()
	_ret.Height = _elem.jsValue.Get("scrollHeight").Float()
	return _ret
}

// ClientRect returns border coordinates of an element in pixels.
//
// Note: This property will round the value to an integer. If you need a fractional value, use element.getBoundingClientRect().
//
//   - https://developer.mozilla.org/en-US/docs/Web/API/Element/clientTop
//   - https://developer.mozilla.org/en-US/docs/Web/API/Element/clientLeft
func (_elem *Element) ClientRect() (_ret lib.Rect) {
	_ret.X = float64(_elem.jsValue.Get("clientLeft").Int())
	_ret.Y = float64(_elem.jsValue.Get("clientTop").Int())
	_ret.Width = float64(_elem.jsValue.Get("clientWidth").Int())
	_ret.Height = float64(_elem.jsValue.Get("clientHeight").Int())
	return _ret
}

// GetBoundingClientRect eturns a DOMRect object providing information about the size of an element and its position relative to the viewport.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
func (_elem *Element) BoundingClientRect() (_result *lib.Rect) {
	_returned := _elem.jsValue.Call("getBoundingClientRect")
	return RectFromJS(_returned)
}

// Children returns a live HTMLCollection which contains all of the child elements of the element upon which it was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/children
func (_elem *Element) Children() Nodes {
	value := _elem.jsValue.Get("children")
	//	ret = newHTMLCollectionFromJS(value)
	ret := MakeNodesFromJSNodeList(value)
	return ret
}

// ChildElementCount returns the number of child elements of this element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/childElementCount
func (_elem *Element) ChildrenCount() uint {
	var ret uint
	value := _elem.jsValue.Get("childElementCount")
	ret = (uint)((value).Int())
	return ret
}

// FirstElementChild  returns an element's first child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/firstElementChild
func (_elem *Element) ChildFirst() *Element {
	value := _elem.jsValue.Get("firstElementChild")
	return NewElementFromJS(value)
}

// LastElementChild returns an element's last child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/lastElementChild
func (_elem *Element) ChildLast() *Element {
	value := _elem.jsValue.Get("lastElementChild")
	return NewElementFromJS(value)
}

// GetElementsByTagName returns a live HTMLCollection of elements with the given tag name.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName
func (_elem *Element) ChildrenByTagName(_tagName string) (_result Nodes) {
	returned := _elem.jsValue.Call("getElementsByTagName", _tagName)
	return MakeNodesFromJSNodeList(returned)
	// return newHTMLCollectionFromJS(_returned)
}

// GetElementsByClassName returns a live HTMLCollection which contains every descendant element which has the specified class name or names.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByClassName
func (_elem *Element) ChildrenByClassName(classNames string) (_result Nodes) {
	_returned := _elem.jsValue.Call("getElementsByClassName", classNames)
	return MakeNodesFromJSNodeList(_returned)
	//	return newHTMLCollectionFromJS(_returned)
}

// PreviousElementSibling returns the Element immediately prior to the specified one in its parent's children list,
// or null if the specified element is the first one in the list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/previousElementSibling
func (_elem *Element) SiblingPrevious() *Element {
	value := _elem.jsValue.Get("previousElementSibling")
	return NewElementFromJS(value)
}

// NextElementSibling returns the element immediately following the specified one in its parent's children list, or null if the specified element is the last one in the list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/nextElementSibling
func (_elem *Element) SiblingNext() *Element {
	value := _elem.jsValue.Get("nextElementSibling")
	return NewElementFromJS(value)
}

// Traverses the element and its parents (heading toward the document root) until
// it finds a node that matches the specified CSS selector.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/closest
func (_elem *Element) SelectorClosest(selectors string) (_result *Element) {
	_returned := _elem.jsValue.Call("closest", selectors)
	return NewElementFromJS(_returned)
}

// Matches tests whether the element would be selected by the specified CSS selector.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/matches
func (_elem *Element) SelectorMatches(selectors string) (_result bool) {
	_returned := _elem.jsValue.Call("matches", selectors)
	return (_returned).Bool()
}

// QuerySelector returns the first element that is a descendant of the element on which it is invoked
// that matches the specified group of selectors.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
func (_elem *Element) SelectorQueryFirst(selectors string) (_result *Element) {
	_returned := _elem.jsValue.Call("querySelector", selectors)
	return NewElementFromJS(_returned)
}

// QuerySelectorAll returns a static (not live) NodeList representing a list of elements matching
// the specified group of selectors which are descendants of the element on which the method was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
func (_elem *Element) SelectorQueryAll(selectors string) (_result Nodes) {
	_returned := _elem.jsValue.Call("querySelectorAll", selectors)
	return MakeNodesFromJSNodeList(_returned)
}

// ScrollIntoView scrolls the element's ancestor containers such that the element on which scrollIntoView() is called is visible to the user.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollIntoView
func (_elem *Element) ScrollIntoView() {
	_elem.jsValue.Call("scrollIntoView")
}

// InsertAdjacentElement inserts a given element node at a given position relative to the element it is invoked upon.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentElement
func (_elem *Element) InsertAdjacentElement(where WhereInsert, element *Element) (_result *Element) {
	_returned := _elem.jsValue.Call("insertAdjacentElement", where, element.JSValue())
	return NewElementFromJS(_returned)
}

// InsertAdjacentText given a relative position and a string, inserts a new text node at the given position relative to the element it is called from.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentText
func (_elem *Element) InsertAdjacentText(where WhereInsert, text string) {
	_elem.jsValue.Call("insertAdjacentText", where, text)
}

// InsertAdjacentHTML parses the specified text as HTML or XML and inserts the resulting nodes into the DOM tree at a specified position.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentHTML
func (_elem *Element) InsertAdjacentHTML(where WhereInsert, text string) {
	_elem.jsValue.Call("insertAdjacentHTML", where, text)
}

// Prepend inserts a set of Node objects or string objects before the first child of the Element.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/prepend
func (_elem *Element) Prepend(nodes ...*Union) {
	var _args []interface{} = make([]interface{}, len(nodes))
	var _end int
	for _, n := range nodes {
		jsn := n.JSValue()
		_args[_end] = jsn
		_end++
	}
	_elem.jsValue.Call("prepend", _args[0:_end]...)
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
func (_elem *Element) Append(nodes ...*Union) {
	var _args []interface{} = make([]interface{}, len(nodes))
	var _end int
	for _, n := range nodes {
		jsn := n.JSValue()
		_args[_end] = jsn
		_end++
	}
	_elem.jsValue.Call("append", _args[0:_end]...)
}

func (_elem *Element) Before(nodes ...*Union) {
	var _args []interface{} = make([]interface{}, len(nodes))
	var _end int
	for _, n := range nodes {
		jsn := n.JSValue()
		_args[_end] = jsn
		_end++
	}
	_elem.jsValue.Call("before", _args[0:_end]...)
}

func (_elem *Element) After(nodes ...*Union) {
	var _args []interface{} = make([]interface{}, len(nodes))
	var _end int
	for _, n := range nodes {
		jsn := n.JSValue()
		_args[_end] = jsn
		_end++
	}
	_elem.jsValue.Call("after", _args[0:_end]...)
}

/****************************************************************************
* Element's events
*****************************************************************************/

type ListenerElem_Generic func(event *Event, target *Element)

// event attribute: Event
func eventFuncElement_Event(listener ListenerElem_Generic) js.Func {

	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := NewEventFromJS(value)
		targ := NewElementFromJS(value.Get("target"))
		listener(evt, targ)
		return js.Undefined()
	}

	return js.FuncOf(fn)
}

// AddFullscreenChange is adding doing AddEventListener for 'FullscreenChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_elem *Element) AddFullscreenEvent(evttype FULLSCREEN_EVENT, listener ListenerElem_Generic) js.Func {
	cb := eventFuncElement_Event(listener)
	_elem.jsValue.Call("addEventListener", evttype, cb)
	return cb
}
