package ick

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"syscall/js"

	"github.com/sunraylab/icecake/internal/helper"
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

// CastElement is casting a js.Value into Element.
func CastElement(value js.Value) *Element {
	if value.Type() != js.TypeObject {
		ConsoleErrorf("casting Element failed")
		return new(Element)
	}
	cast := new(Element)
	cast.Value = value
	return cast
}

// GetICElementById returns an Element corresponding to the _id if it exists into the DOM,
// otherwhise returns an undefined Element.
// The _id must be normalized before call
func GetElementById(_id string) *Element {
	jse := App().Call("getElementById", _id)
	if jse.Truthy() && CastNode(jse).NodeType() == NT_ELEMENT {
		elem := new(Element)
		elem.Wrap(jse)
		return elem
	}
	ConsoleWarnf("GetElementById failed: %q not found, or not a <Element>", _id)
	return new(Element)
}

// IsDefined returns true if the Element is not nil AND it's type is not TypeNull and not TypeUndefined
func (_node *Element) IsDefined() bool {
	if _node == nil || _node.Value.Type() == js.TypeNull || _node.Value.Type() == js.TypeUndefined {
		return false
	}
	return true
}

// Remove removes the element from the DOM.
// d
// https://developer.mozilla.org/en-US/docs/Web/API/Element/remove
func (_elem *Element) Remove() {
	if !_elem.IsDefined() {
		return
	}
	_elem.Call("remove")
}

/****************************************************************************
* Element's Properties & Methods
*****************************************************************************/

// TagName returns the tag name of the element on which it's called.
// Always in Uppercase for HTML element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/tagName
func (_elem *Element) TagName() string {
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return _elem.Get("tagName").String()
}

// Id rrepresents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (_elem *Element) Id() string {
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return _elem.Get("id").String()
}

// Id rrepresents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (_elem *Element) SetId(_id string) *Element {
	if !_elem.IsDefined() {
		return _elem
	}
	_elem.Set("id", _id)
	return _elem
}

// SetClassName gets and sets the value of the class attribute of the specified element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (_elem *Element) ClassName() string {
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	value := _elem.Get("className")
	return (value).String()
}

// SetClassName gets and sets the value of the class attribute of the specified element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (_elem *Element) SetClassName(_name string) *Element {
	if !_elem.IsDefined() {
		return _elem
	}
	_elem.Set("className", _name)
	return _elem
}

// ClassList returns a live DOMTokenList collection of the class attributes of the element.
// This can then be used to manipulate the class list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/classList
func (_elem *Element) ClassList() *TokenList {
	if !_elem.IsDefined() {
		return new(TokenList)
	}
	value := _elem.Get("classList")
	return CastTokenList(value)
}

// Attributes returns a live collection of all attribute nodes registered to the specified node.
// It is a NamedNodeMap, not an Array, so it has no Array methods and the Attr nodes' indexes may differ among browsers.
// To be more specific, attributes is a key/value pair of strings that represents any information regarding that attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
func (_elem *Element) Attributes() *Attributes {
	if !_elem.IsDefined() {
		return NewAttributes(_elem)
	}
	attrs := NewAttributes(_elem)
	namedNodeMap := _elem.Get("attributes")
	if typ := namedNodeMap.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		ConsoleWarnf("No attributes found")
		return attrs
	}

	data := namedNodeMap.Get("length")
	len := data.Int()
	for i := 0; i < len; i++ {
		item := namedNodeMap.Call("item", uint(i))
		attr := CastAttribute(item)
		attrs.attributes = append(attrs.attributes, attr)
	}
	return attrs
}

// InnerHTML gets or sets the HTML or XML markup contained within the element.
//
// To insert the HTML into the document rather than replace the contents of an element, use the method insertAdjacentHTML().
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (_elem *Element) InnerHTML() string {
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	value := _elem.Get("innerHTML")
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
func (_elem *Element) SetInnerHTML(_unsafeHtml string) *Element {
	if !_elem.IsDefined() {
		return _elem
	}
	_elem.Set("innerHTML", _unsafeHtml)
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
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return _elem.Get("outerHTML").String()
}

// OuterHTML gets the serialized HTML fragment describing the element including its descendants.
// It can also be set to replace the element with nodes parsed from the given string.
//
// To only obtain the HTML representation of the contents of an element,
// or to replace the contents of an element, use the innerHTML property instead.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
func (_elem *Element) SetOuterHTML(_unsafeHtml string) *Element {
	if !_elem.IsDefined() {
		return _elem
	}
	_elem.Set("outerHTML", _unsafeHtml)
	return _elem
}

// ChildElementCount returns the number of child elements of this element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/childElementCount
func (_elem *Element) ChildrenCount() int {
	if !_elem.IsDefined() {
		return 0
	}
	count := _elem.Get("childElementCount")
	return count.Int()
}

// Children returns a live HTMLCollection which contains all of the child elements of the element upon which it was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/children
func (_elem *Element) Children() []*Node {
	if !_elem.IsDefined() {
		return make([]*Node, 0)
	}
	nodes := _elem.Get("children")
	return MakeNodes(nodes)
}

// FirstElementChild  returns an element's first child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/firstElementChild
func (_elem *Element) ChildFirst() *Element {
	if !_elem.IsDefined() {
		return new(Element)
	}
	child := _elem.Get("firstElementChild")
	if typ := child.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return new(Element)
	}
	return CastElement(child)
}

// LastElementChild returns an element's last child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/lastElementChild
func (_elem *Element) ChildLast() *Element {
	if !_elem.IsDefined() {
		return new(Element)
	}
	child := _elem.Get("lastElementChild")
	if typ := child.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return new(Element)
	}
	return CastElement(child)
}

// GetElementsByTagName returns a live HTMLCollection of elements with the given tag name.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName
func (_elem *Element) ChildrenByTagName(_tagName string) []*Node {
	if !_elem.IsDefined() {
		return make([]*Node, 0)
	}
	nodes := _elem.Call("getElementsByTagName", _tagName)
	return MakeNodes(nodes)
}

// GetElementsByClassName returns a live HTMLCollection which contains every descendant element which has the specified class name or names.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByClassName
func (_elem *Element) ChildrenByClassName(_classNames string) []*Node {
	if !_elem.IsDefined() {
		return make([]*Node, 0)
	}
	nodes := _elem.Call("getElementsByClassName", _classNames)
	return MakeNodes(nodes)
}

// PreviousElementSibling returns the Element immediately prior to the specified one in its parent's children list,
// or null if the specified element is the first one in the list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/previousElementSibling
func (_elem *Element) SiblingPrevious() *Element {
	if !_elem.IsDefined() {
		return new(Element)
	}
	sibling := _elem.Get("previousElementSibling")
	if typ := sibling.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return new(Element)
	}
	return CastElement(sibling)
}

// NextElementSibling returns the element immediately following the specified one in its parent's children list, or null if the specified element is the last one in the list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/nextElementSibling
func (_elem *Element) SiblingNext() *Element {
	if !_elem.IsDefined() {
		return new(Element)
	}
	sibling := _elem.Get("nextElementSibling")
	if typ := sibling.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return new(Element)
	}
	return CastElement(sibling)
}

// Traverses the element and its parents (heading toward the document root) until
// it finds a node that matches the specified CSS selector.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/closest
func (_elem *Element) SelectorClosest(_selectors string) *Element {
	if !_elem.IsDefined() {
		return new(Element)
	}
	elem := _elem.Call("closest", _selectors)
	if typ := elem.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return new(Element)
	}
	return CastElement(elem)
}

// Matches tests whether the element would be selected by the specified CSS selector.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/matches
func (_elem *Element) SelectorMatches(_selectors string) bool {
	ok := _elem.Call("matches", _selectors)
	return ok.Bool()
}

// QuerySelector returns the first element that is a descendant of the element on which it is invoked
// that matches the specified group of selectors.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
func (_elem *Element) SelectorQueryFirst(_selectors string) *Element {
	if !_elem.IsDefined() {
		return new(Element)
	}
	elem := _elem.Call("querySelector", _selectors)
	if typ := elem.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return new(Element)
	}
	return CastElement(elem)
}

// QuerySelectorAll returns a static (not live) NodeList representing a list of elements matching
// the specified group of selectors which are descendants of the element on which the method was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
func (_elem *Element) SelectorQueryAll(_selectors string) []*Node {
	if !_elem.IsDefined() {
		return make([]*Node, 0)
	}
	nodes := _elem.Call("querySelectorAll", _selectors)
	return MakeNodes(nodes)
}

// InsertAdjacentElement inserts a given element node at a given position relative to the element it is invoked upon.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentElement
func (_elem *Element) InsertAdjacentElement(_where WhereInsert, _element *Element) *Element {
	if !_elem.IsDefined() {
		return new(Element)
	}
	elem := _elem.Call("insertAdjacentElement", string(_where), _element.JSValue())
	return CastElement(elem)
}

// InsertAdjacentText given a relative position and a string, inserts a new text node at the given position relative to the element it is called from.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentText
func (_elem *Element) InsertAdjacentText(_where WhereInsert, _text string) {
	if !_elem.IsDefined() {
		return
	}
	_elem.Call("insertAdjacentText", string(_where), _text)
}

// InsertAdjacentHTML parses the specified text as HTML or XML and inserts the resulting nodes into the DOM tree at a specified position.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentHTML
func (_elem *Element) InsertAdjacentHTML(_where WhereInsert, _text string) {
	if !_elem.IsDefined() {
		return
	}
	_elem.Call("insertAdjacentHTML", string(_where), _text)
}

// Prepend inserts a set of Node objects or string objects before the first child of the Element.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/prepend
func (_elem *Element) PrependNodes(_nodes ...*Node) {
	if !_elem.IsDefined() {
		return
	}
	var args []interface{} = make([]interface{}, len(_nodes))
	var end int
	for _, n := range _nodes {
		if n != nil {
			jsn := n.JSValue()
			args[end] = jsn
			end++
		}
	}
	_elem.Call("prepend", args[0:end]...)
}

// Prepend inserts a set of Node objects or string objects before the first child of the Element.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/prepend
func (_elem *Element) PrependStrings(_strs []string) {
	if !_elem.IsDefined() {
		return
	}
	var args []interface{} = make([]interface{}, len(_strs))
	var end int
	for _, n := range _strs {
		args[end] = n
		end++
	}
	_elem.Call("prepend", args[0:end]...)
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
func (_elem *Element) AppendNodes(_nodes []*Node) {
	if !_elem.IsDefined() {
		return
	}
	var args []interface{} = make([]interface{}, len(_nodes))
	var end int
	for _, n := range _nodes {
		if n != nil {
			jsn := n.JSValue()
			args[end] = jsn
			end++
		}
	}
	_elem.Call("append", args[0:end]...)
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
func (_elem *Element) AppendStrings(_strs ...string) {
	if !_elem.IsDefined() {
		return
	}
	var args []interface{} = make([]interface{}, len(_strs))
	var end int
	for _, n := range _strs {
		args[end] = n
		end++
	}
	_elem.Call("append", args[0:end]...)
}

func (_elem *Element) InsertNodesBefore(_nodes []*Node) {
	if !_elem.IsDefined() {
		return
	}
	var _args []interface{} = make([]interface{}, len(_nodes))
	var _end int
	for _, n := range _nodes {
		if n != nil {
			jsn := n.JSValue()
			_args[_end] = jsn
			_end++
		}
	}
	_elem.Call("before", _args[0:_end]...)
}

func (_elem *Element) InsertNodesAfter(_nodes []*Node) {
	if !_elem.IsDefined() {
		return
	}
	var _args []interface{} = make([]interface{}, len(_nodes))
	var _end int
	for _, n := range _nodes {
		if n != nil {
			jsn := n.JSValue()
			_args[_end] = jsn
			_end++
		}
	}
	_elem.Call("after", _args[0:_end]...)
}

// ScrollTop gets or sets the number of pixels that an element's content is scrolled vertically.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollTop
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollWidth
func (_elem *Element) ScrollRect() (_rect lib.Rect) {
	if !_elem.IsDefined() {
		return
	}
	_rect.X = _elem.Get("scrollLeft").Float()
	_rect.Y = _elem.Get("scrollTop").Float()
	_rect.Width = _elem.Get("scrollWidth").Float()
	_rect.Height = _elem.Get("scrollHeight").Float()
	return _rect
}

// ClientRect returns border coordinates of an element in pixels.
//
// Note: This property will round the value to an integer. If you need a fractional value, use element.getBoundingClientRect().
//
//   - https://developer.mozilla.org/en-US/docs/Web/API/Element/clientTop
//   - https://developer.mozilla.org/en-US/docs/Web/API/Element/clientLeft
func (_elem *Element) ClientRect() (_rect lib.Rect) {
	if !_elem.IsDefined() {
		return lib.Rect{}
	}
	_rect.X = float64(_elem.Get("clientLeft").Int())
	_rect.Y = float64(_elem.Get("clientTop").Int())
	_rect.Width = float64(_elem.Get("clientWidth").Int())
	_rect.Height = float64(_elem.Get("clientHeight").Int())
	return _rect
}

// GetBoundingClientRect eturns a DOMRect object providing information about the size of an element and its position relative to the viewport.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
func (_elem *Element) BoundingClientRect() lib.Rect {
	if !_elem.IsDefined() {
		return lib.Rect{}
	}
	jsrect := _elem.Call("getBoundingClientRect")

	rect := new(lib.Rect)
	rect.X = jsrect.Get("x").Float()
	rect.Y = jsrect.Get("y").Float()
	rect.Width = jsrect.Get("width").Float()
	rect.Height = jsrect.Get("height").Float()
	return *rect
}

// ScrollIntoView scrolls the element's ancestor containers such that the element on which scrollIntoView() is called is visible to the user.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollIntoView
func (_elem *Element) ScrollIntoView() {
	if !_elem.IsDefined() {
		return
	}
	_elem.Call("scrollIntoView")
}

/****************************************************************************
* HTMLElement's properties & methods
*****************************************************************************/

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (_elem *Element) AccessKey() string {
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return _elem.Get("accessKey").String()
}

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (_htmle *Element) SetAccessKey(key bool) *Element {
	if !_htmle.IsDefined() {
		return nil
	}
	_htmle.Set("accessKey", key)
	return _htmle
}

// InnerText represents the rendered text content of a node and its descendants.
//
// InnerText gets pure text, removing any html or css, while TextContent keeps the representation.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerText
func (_htmle *Element) InnerText() string {
	if !_htmle.IsDefined() {
		return UNDEFINED_NODE
	}
	var ret string
	value := _htmle.Get("innerText")
	ret = (value).String()
	return ret
}

// InnerText represents the rendered text content of a node and its descendants.
//
// InnerText gets pure text, removing any html or css, while TextContent keeps the representation.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerText
func (_htmle *Element) SetInnerText(value string) {
	if !_htmle.IsDefined() {
		return
	}
	input := value
	_htmle.Set("innerText", input)
}

// Focus sets focus on the specified element, if it can be focused. The focused element is the element that will receive keyboard and similar events by default.
//
// By default the browser will scroll the element into view after focusing it,
// and it may also provide visible indication of the focused element (typically by displaying a "focus ring" around the element).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/focus
func (_htmle *Element) Focus() {
	if !_htmle.IsDefined() {
		return
	}
	_htmle.Call("focus")
}

// Blur removes keyboard focus from the current element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/blur
func (_htmle *Element) Blur() {
	if !_htmle.IsDefined() {
		return
	}
	_htmle.Call("blur")
}

/****************************************************************************
* Element's events
*****************************************************************************/

// event attribute: Event
func eventFuncElement_Event(listener func(event *Event, target *Element)) js.Func {

	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastEvent(value)
		targ := CastElement(value.Get("target"))
		listener(evt, targ)
		return js.Undefined()
	}

	return js.FuncOf(fn)
}

// AddFullscreenChange is adding doing AddEventListener for 'FullscreenChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_elem *Element) AddFullscreenEvent(evttype FULLSCREEN_EVENT, listener func(event *Event, target *Element)) js.Func {
	if !_elem.IsDefined() {
		ConsoleWarnf("AddFullscreenEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	cb := eventFuncElement_Event(listener)
	_elem.Call("addEventListener", string(evttype), cb)
	return cb
}

/****************************************************************************
* HTMLElement's events
*****************************************************************************/

// event attribute: Event
func makeHTMLElement_domcore_Event(listener func(event *Event, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastEvent(value)
		target := CastElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddAbort is adding doing AddEventListener for 'Abort' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddGenericEvent(evttype GENERIC_EVENT, listener func(event *Event, target *Element)) js.Func {
	if !_htmle.IsDefined() {
		ConsoleWarnf("AddGenericEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_domcore_Event(listener)
	_htmle.Call("addEventListener", string(evttype), callback)
	return callback
}

// event attribute: MouseEvent
func makeHTMLElement_Mouse_Event(listener func(event *MouseEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastMouseEvent(value)
		target := CastElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddClick is adding doing AddEventListener for 'Click' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddMouseEvent(evttype MOUSE_EVENT, listener func(event *MouseEvent, target *Element)) js.Func {
	if !_htmle.IsDefined() {
		ConsoleWarnf("AddMouseEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_Mouse_Event(listener)
	_htmle.Call("addEventListener", string(evttype), callback)
	return callback
}

// event attribute: FocusEvent
func makeHTMLElement_FocusEvent(listener func(event *FocusEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastFocusEvent(value)
		target := CastElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddBlur is adding doing AddEventListener for 'Blur' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddFocusEvent(evttype FOCUS_EVENT, listener func(event *FocusEvent, target *Element)) js.Func {
	if !_htmle.IsDefined() {
		ConsoleWarnf("AddFocusEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_FocusEvent(listener)
	_htmle.Call("addEventListener", string(evttype), callback)
	return callback
}

// event attribute: PointerEvent
func makeHTMLElement_PointerEvent(listener func(event *PointerEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var evt *PointerEvent
		value := args[0]
		evt = CastPointerEvent(value)
		target := CastElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddGotPointerCapture is adding doing AddEventListener for 'GotPointerCapture' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddPointerEvent(evttype POINTER_EVENT, listener func(event *PointerEvent, target *Element)) js.Func {
	if !_htmle.IsDefined() {
		ConsoleWarnf("AddPointerEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_PointerEvent(listener)
	_htmle.Call("addEventListener", string(evttype), callback)
	return callback
}

// event attribute: InputEvent
func makeHTMLElement_InputEvent(listener func(event *InputEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastInputEvent(value)
		target := CastElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddInput is adding doing AddEventListener for 'Input' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddInputEvent(evttype INPUT_EVENT, listener func(event *InputEvent, target *Element)) js.Func {
	if !_htmle.IsDefined() {
		ConsoleWarnf("AddInputEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_InputEvent(listener)
	_htmle.Call("addEventListener", string(evttype), callback)
	return callback
}

// event attribute: KeyboardEvent
func makeHTMLElement_KeyboardEvent(listener func(event *KeyboardEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastKeyboardEvent(value)
		target := CastElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddKeyDown is adding doing AddEventListener for 'KeyDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddKeyboard(evttype KEYBOARD_EVENT, listener func(event *KeyboardEvent, target *Element)) js.Func {
	if !_htmle.IsDefined() {
		ConsoleWarnf("AddKeyboard not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_KeyboardEvent(listener)
	_htmle.Call("addEventListener", string(evttype), callback)
	return callback
}

// event attribute: UIEvent
func makeHTMLElement_UIEvent(listener func(event *UIEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastUIEvent(value)
		target := CastElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddResize is adding doing AddEventListener for 'Resize' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddResizeEvent(listener func(event *UIEvent, target *Element)) js.Func {
	if !_htmle.IsDefined() {
		ConsoleWarnf("AddResizeEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_UIEvent(listener)
	_htmle.Call("addEventListener", "resize", callback)
	return callback
}

// event attribute: WheelEvent
func makeHTMLElement_WheelEvent(listener func(event *WheelEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastWheelEvent(value)
		target := CastElement(value.Get("target"))
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// The wheel event fires when the user rotates a wheel button on a pointing device (typically a mouse).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/wheel_event
func (_htmle *Element) AddWheelEvent(listener func(event *WheelEvent, target *Element)) js.Func {
	if !_htmle.IsDefined() {
		ConsoleWarnf("AddWheelEvent not listening on nil Element")
		return js.FuncOf(func(this js.Value, args []js.Value) interface{} { return js.Undefined() })
	}
	callback := makeHTMLElement_WheelEvent(listener)
	_htmle.Call("addEventListener", "wheel", callback)
	return callback
}

/****************************************************************************
* Extra rendering
*****************************************************************************/

// SetInnerValue set the innext text of the element with a formated value.
// The format string follow the fmt rules: https://pkg.go.dev/fmt#hdr-Printing
func (_elem *Element) RenderValue(format string, _value ...any) {
	if !_elem.IsDefined() {
		return
	}
	text := fmt.Sprintf(format, _value...)
	_elem.SetInnerText(text)
}

// RenderHtml set inner HTML with the htmlTemplate executed with the _data and unfolding components if any
func (_elem *Element) RenderHtml(_unsafeHtmlTemplate string, _data any) (_err error) {
	if !_elem.IsDefined() {
		return
	}
	name := _elem.TagName() + "/" + _elem.Id()
	var html string
	html, _err = unfoldComponents(name, _unsafeHtmlTemplate, _data, 0)
	if _err == nil {
		_elem.SetInnerHTML(html)
	}
	return _err
}

// RenderNamedValue look recursively for any _elem children having the "data-ic-namedvalue" token matching _name
// and render inner text with the _value
func (_elem *Element) RenderChildrenValue(_name string, _format string, _value ...any) {
	if !_elem.IsDefined() {
		return
	}
	_name = helper.Normalize(_name)
	text := fmt.Sprintf(_format, _value...)

	children := _elem.FilteredChildren(NT_ELEMENT, 99, func(_node *Node) bool {
		dataset := CastElement(_node.JSValue()).Attributes().Dataset()
		for i := 0; i < dataset.Count(); i++ {
			if dataset.At(i).Name() == "data-ic-namedvalue" && dataset.At(i).Value() == _name {
				return true
			}
		}
		return false
	})

	for _, node := range children {
		CastElement(node.JSValue()).RenderValue(text)
	}
}

// InsertComponent
func (_elem *Element) InsertNewComponent(_newcmp any) (_newcmpid string, _err error) {
	if !_elem.IsDefined() {
		_err = errors.New("InsertComponent failed on nil element")
		ConsoleWarnf(_err.Error())
		return "", _err
	}

	// check if _newcmp is a registered component and get its tagname
	cmptype := LookupComponent(reflect.TypeOf(_newcmp))
	if cmptype == "" {
		ConsoleWarnf("InsertNewComponent: Inserting a non registered component %q...", reflect.TypeOf(_newcmp).String())
	}

	var cmpelem *Element
	switch compounder := _newcmp.(type) {
	case HtmlCompounder:

		// create the HTML component into the DOM
		cmpelem, _err = CreateCompoundElement(compounder)
		if _err == nil {
			nc := GetNextComponentIndex()
			_newcmpid = "c" + strconv.Itoa(nc)
			if cmptype != "" {
				_newcmpid = cmptype + "-" + _newcmpid
			}
			cmpelem.SetId(_newcmpid)

			// name the component
			name := cmpelem.TagName() + "/" + _newcmpid

			// unfold and render html for a compounder
			data := TemplateData{
				Me:     _newcmp,
				Global: &GData,
			}
			html, _ := unfoldComponents(name, compounder.Template(), data, 0)
			cmpelem.SetInnerHTML(html)

			//elem.InsertAdjacentHTML(WI_INSIDEFIRST, html)
			_elem.PrependNodes(&cmpelem.Node)

			// wrap this new html element to th _cmp
			switch wrapper := _newcmp.(type) {
			case JSWrapper:
				if typ := wrapper.JSValue().Type(); typ == js.TypeNull || typ == js.TypeUndefined {
					// fmt.Println("_newcmp is a Element")
					wrapper.Wrap(cmpelem.JSValue())
				} else {
					return "", fmt.Errorf("component %q has already been inserted", reflect.TypeOf(_newcmp).String())
				}
			default:
				return "", fmt.Errorf("component %q is not an Element", reflect.TypeOf(_newcmp).String())
			}

			// TODO: add style

			// addlisteners
			switch listener := _newcmp.(type) {
			case HtmlListener:
				// fmt.Println("_newcmp is a listener")
				listener.AddListeners()
			}

		} else {
			ConsoleWarnf(_err.Error())
			return "", _err
		}
	default:
		return "", errors.New("InsertComponent failed: _newcmp is not a compounder")
	}

	return _newcmpid, nil
}
