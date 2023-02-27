package ick

import (
	"errors"
	"fmt"
	"reflect"
	"syscall/js"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/lib"
)

/****************************************************************************
* enum used by Element
*****************************************************************************/

type WHERE_INSERT string

const (
	WI_BEFOREBEGIN WHERE_INSERT = "beforebegin" // Before the Element itself.
	WI_INSIDEFIRST WHERE_INSERT = "afterbegin"  // Just inside the element, before its first child.
	WI_INSIDELAST  WHERE_INSERT = "beforeend"   // Just inside the element, after its last child.
	WI_AFTEREND    WHERE_INSERT = "afterend"    // After the element itself.
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
	cast.jsValue = value
	return cast
}

// GetICElementById returns an Element corresponding to the _id if it exists into the DOM,
// otherwhise returns an undefined Element.
// The _id must be normalized before call
func GetElementById(_id string) *Element {
	jse := App().jsValue.Call("getElementById", _id)
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
	if _node == nil || _node.jsValue.Type() == js.TypeNull || _node.jsValue.Type() == js.TypeUndefined {
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
	_elem.jsValue.Call("remove")
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
	return _elem.jsValue.Get("tagName").String()
}

// Id rrepresents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (_elem *Element) Id() string {
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return _elem.jsValue.Get("id").String()
}

// Id rrepresents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (_elem *Element) SetId(_id string) *Element {
	if !_elem.IsDefined() {
		return _elem
	}
	_elem.jsValue.Set("id", _id)
	return _elem
}

// SetClassName gets and sets the value of the class attribute of the specified element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (_elem *Element) ClassName() string {
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	value := _elem.jsValue.Get("className")
	return (value).String()
}

// SetClassName gets and sets the value of the class attribute of the specified element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (_elem *Element) SetClassName(_name string) *Element {
	if !_elem.IsDefined() {
		return _elem
	}
	_elem.jsValue.Set("className", _name)
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
	value := _elem.jsValue.Get("classList")
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
	namedNodeMap := _elem.jsValue.Get("attributes")
	if typ := namedNodeMap.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		ConsoleLogf("No attributes found")
		return attrs
	}

	var value string
	data := namedNodeMap.Get("length")
	len := data.Int()
	for i := 0; i < len; i++ {
		item := namedNodeMap.Call("item", uint(i))
		name := item.Get("name").String()
		jsvalue := item.Get("value")
		if jsvalue.Truthy() {
			value = jsvalue.String()
		}
		attrs.Set(name, value)
	}
	return attrs
}

func (_elem *Element) SetAttribute(_Name string, _Value string) {
	_elem.jsValue.Call("setAttribute", _Name, _Value)
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
func (_elem *Element) SetInnerHTML(_unsafeHtml string) *Element {
	if !_elem.IsDefined() {
		return _elem
	}
	_elem.jsValue.Set("innerHTML", _unsafeHtml)
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
	return _elem.jsValue.Get("outerHTML").String()
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
	_elem.jsValue.Set("outerHTML", _unsafeHtml)
	return _elem
}

// ChildElementCount returns the number of child elements of this element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/childElementCount
func (_elem *Element) ChildrenCount() int {
	if !_elem.IsDefined() {
		return 0
	}
	count := _elem.jsValue.Get("childElementCount")
	return count.Int()
}

// Children returns a live HTMLCollection which contains all of the child elements of the element upon which it was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/children
func (_elem *Element) Children() []*Node {
	if !_elem.IsDefined() {
		return make([]*Node, 0)
	}
	nodes := _elem.jsValue.Get("children")
	return MakeNodes(nodes)
}

// FirstElementChild  returns an element's first child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/firstElementChild
func (_elem *Element) ChildFirst() *Element {
	if !_elem.IsDefined() {
		return new(Element)
	}
	child := _elem.jsValue.Get("firstElementChild")
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
	child := _elem.jsValue.Get("lastElementChild")
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
	nodes := _elem.jsValue.Call("getElementsByTagName", _tagName)
	return MakeNodes(nodes)
}

// GetElementsByClassName returns a live HTMLCollection which contains every descendant element which has the specified class name or names.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByClassName
func (_elem *Element) ChildrenByClassName(_classNames string) []*Node {
	if !_elem.IsDefined() {
		return make([]*Node, 0)
	}
	nodes := _elem.jsValue.Call("getElementsByClassName", _classNames)
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
	sibling := _elem.jsValue.Get("previousElementSibling")
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
	sibling := _elem.jsValue.Get("nextElementSibling")
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
	elem := _elem.jsValue.Call("closest", _selectors)
	if typ := elem.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return new(Element)
	}
	return CastElement(elem)
}

// Matches tests whether the element would be selected by the specified CSS selector.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/matches
func (_elem *Element) SelectorMatches(_selectors string) bool {
	ok := _elem.jsValue.Call("matches", _selectors)
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
	elem := _elem.jsValue.Call("querySelector", _selectors)
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
	nodes := _elem.jsValue.Call("querySelectorAll", _selectors)
	return MakeNodes(nodes)
}

// InsertAdjacentElement inserts a given element node at a given position relative to the element it is invoked upon.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentElement
func (_elem *Element) InsertAdjacentElement(_where WHERE_INSERT, _element *Element) *Element {
	if !_elem.IsDefined() {
		return new(Element)
	}
	elem := _elem.jsValue.Call("insertAdjacentElement", string(_where), _element.JSValue())
	return CastElement(elem)
}

// InsertAdjacentText given a relative position and a string, inserts a new text node at the given position relative to the element it is called from.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentText
func (_elem *Element) InsertAdjacentText(_where WHERE_INSERT, _text string) {
	if !_elem.IsDefined() {
		return
	}
	_elem.jsValue.Call("insertAdjacentText", string(_where), _text)
}

// InsertAdjacentHTML parses the specified text as HTML or XML and inserts the resulting nodes into the DOM tree at a specified position.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentHTML
func (_elem *Element) InsertAdjacentHTML(_where WHERE_INSERT, _text string) {
	if !_elem.IsDefined() {
		return
	}
	_elem.jsValue.Call("insertAdjacentHTML", string(_where), _text)
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
	_elem.jsValue.Call("prepend", args[0:end]...)
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
	_elem.jsValue.Call("prepend", args[0:end]...)
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
	_elem.jsValue.Call("append", args[0:end]...)
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
	_elem.jsValue.Call("append", args[0:end]...)
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
	_elem.jsValue.Call("before", _args[0:_end]...)
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
	_elem.jsValue.Call("after", _args[0:_end]...)
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
	_rect.X = _elem.jsValue.Get("scrollLeft").Float()
	_rect.Y = _elem.jsValue.Get("scrollTop").Float()
	_rect.Width = _elem.jsValue.Get("scrollWidth").Float()
	_rect.Height = _elem.jsValue.Get("scrollHeight").Float()
	return _rect
}

// ClientRect returns border coordinates of an element in pixels.
//
// Note: This property will round the value to an integer. If you need a fractional value, use element.jsValue.getBoundingClientRect().
//
//   - https://developer.mozilla.org/en-US/docs/Web/API/Element/clientTop
//   - https://developer.mozilla.org/en-US/docs/Web/API/Element/clientLeft
func (_elem *Element) ClientRect() (_rect lib.Rect) {
	if !_elem.IsDefined() {
		return lib.Rect{}
	}
	_rect.X = float64(_elem.jsValue.Get("clientLeft").Int())
	_rect.Y = float64(_elem.jsValue.Get("clientTop").Int())
	_rect.Width = float64(_elem.jsValue.Get("clientWidth").Int())
	_rect.Height = float64(_elem.jsValue.Get("clientHeight").Int())
	return _rect
}

// GetBoundingClientRect eturns a DOMRect object providing information about the size of an element and its position relative to the viewport.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
func (_elem *Element) BoundingClientRect() lib.Rect {
	if !_elem.IsDefined() {
		return lib.Rect{}
	}
	jsrect := _elem.jsValue.Call("getBoundingClientRect")

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
	_elem.jsValue.Call("scrollIntoView")
}

/****************************************************************************
* HTMLElement's properties & methods
*****************************************************************************/

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (_elem *Element) AccessKey() string {
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return _elem.jsValue.Get("accessKey").String()
}

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (_htmle *Element) SetAccessKey(key bool) *Element {
	if !_htmle.IsDefined() {
		return nil
	}
	_htmle.jsValue.Set("accessKey", key)
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
	value := _htmle.jsValue.Get("innerText")
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
	_htmle.jsValue.Set("innerText", input)
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
	_htmle.jsValue.Call("focus")
}

// Blur removes keyboard focus from the current element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/blur
func (_htmle *Element) Blur() {
	if !_htmle.IsDefined() {
		return
	}
	_htmle.jsValue.Call("blur")
}

/****************************************************************************
* Element's events
*****************************************************************************/

// event attribute: Event
func makelistenerElement_Event(listener func(event *Event, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on %q id=%q", evt.Type(), target.TagName(), target.Id())
			}
		}()
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddFullscreenChange is adding doing AddEventListener for 'FullscreenChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_elem *Element) AddFullscreenEvent(evttype FULLSCREEN_EVENT, listener func(event *Event, target *Element)) {
	if !_elem.IsDefined() || !_elem.IsInDOM() {
		ConsoleWarnf("AddFullscreenEvent not listening on nil Element")
		return
	}
	evh := makelistenerElement_Event(listener)
	_elem.jsValue.Call("addEventListener", string(evttype), evh)
	_elem.AddEventListener(&eventHandler{eventtype: string(evttype), jsHandler: evh})
}

/****************************************************************************
* HTMLElement's events
*****************************************************************************/

// event attribute: Event
func makehandler_Element_Event(listener func(event *Event, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on %q id=%q", evt.Type(), target.TagName(), target.Id())
			}
		}()
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddGenericEvent adds Event Listener for a GENERIC_EVENT  on target.
// Returns the function to call to remove and release the listener
func (_htmle *Element) AddGenericEvent(evttype GENERIC_EVENT, listener func(event *Event, target *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		ConsoleWarnf("AddGenericEvent failed: nil Element or not in DOM")
		return
	}
	jsevh := makehandler_Element_Event(listener)
	_htmle.AddEventListener(&eventHandler{eventtype: string(evttype), jsHandler: jsevh})
}

// event attribute: MouseEvent
func makehandler_Element_MouseEvent(listener func(event *MouseEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastMouseEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on %q id=%q", evt.Type(), target.TagName(), target.Id())
			}
		}()
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddClick adds Event Listener for MOUSE_EVENT on target.
// Returns the function to call to remove and release the listener
func (_htmle *Element) AddMouseEvent(evttype MOUSE_EVENT, listener func(event *MouseEvent, target *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		ConsoleWarnf("AddMouseEvent failed: nil Element or not in DOM")
		return
	}
	evh := makehandler_Element_MouseEvent(listener)
	_htmle.AddEventListener(&eventHandler{eventtype: string(evttype), jsHandler: evh})
}

// event attribute: FocusEvent
func makelistenerElement_FocusEvent(listener func(event *FocusEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastFocusEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on %q id=%q", evt.Type(), target.TagName(), target.Id())
			}
		}()
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddBlur is adding doing AddEventListener for 'Blur' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddFocusEvent(evttype FOCUS_EVENT, listener func(event *FocusEvent, target *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		ConsoleWarnf("AddFocusEvent failed: nil Element or not in DOM")
		return
	}
	evh := makelistenerElement_FocusEvent(listener)
	_htmle.jsValue.Call("addEventListener", string(evttype), evh)
	_htmle.AddEventListener(&eventHandler{eventtype: string(evttype), jsHandler: evh})
}

// event attribute: PointerEvent
func makelistenerElement_PointerEvent(listener func(event *PointerEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		var evt *PointerEvent
		value := args[0]
		evt = CastPointerEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on %q id=%q", evt.Type(), target.TagName(), target.Id())
			}
		}()
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddGotPointerCapture is adding doing AddEventListener for 'GotPointerCapture' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddPointerEvent(evttype POINTER_EVENT, listener func(event *PointerEvent, target *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		ConsoleWarnf("AddPointerEvent failed: nil Element or not in DOM")
		return
	}
	evh := makelistenerElement_PointerEvent(listener)
	_htmle.jsValue.Call("addEventListener", string(evttype), evh)
	_htmle.AddEventListener(&eventHandler{eventtype: string(evttype), jsHandler: evh})
}

// event attribute: InputEvent
func makelistenerElement_InputEvent(listener func(event *InputEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastInputEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on %q id=%q", evt.Type(), target.TagName(), target.Id())
			}
		}()
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddInput is adding doing AddEventListener for 'Input' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddInputEvent(evttype INPUT_EVENT, listener func(event *InputEvent, target *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		ConsoleWarnf("AddInputEvent failed: nil Element or not in DOM")
		return
	}
	evh := makelistenerElement_InputEvent(listener)
	_htmle.AddEventListener(&eventHandler{eventtype: string(evttype), jsHandler: evh})
}

// event attribute: KeyboardEvent
func makelistenerElement_KeyboardEvent(listener func(event *KeyboardEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastKeyboardEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on %q id=%q", evt.Type(), target.TagName(), target.Id())
			}
		}()
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddKeyDown is adding doing AddEventListener for 'KeyDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddKeyboard(evttype KEYBOARD_EVENT, listener func(event *KeyboardEvent, target *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		ConsoleWarnf("AddKeyboard failed: nil Element or not in DOM")
		return
	}
	evh := makelistenerElement_KeyboardEvent(listener)
	_htmle.jsValue.Call("addEventListener", string(evttype), evh)
	_htmle.AddEventListener(&eventHandler{eventtype: string(evttype), jsHandler: evh})
}

// event attribute: UIEvent
func makelistenerElement_UIEvent(listener func(event *UIEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastUIEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on %q id=%q", evt.Type(), target.TagName(), target.Id())
			}
		}()
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// AddResize is adding doing AddEventListener for 'Resize' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddResizeEvent(listener func(event *UIEvent, target *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		ConsoleWarnf("AddResizeEvent failed: nil Element or not in DOM")
		return
	}
	evh := makelistenerElement_UIEvent(listener)
	_htmle.jsValue.Call("addEventListener", "resize", evh)
	_htmle.AddEventListener(&eventHandler{eventtype: "resize", jsHandler: evh})
}

// event attribute: WheelEvent
func makeHTMLElement_WheelEvent(listener func(event *WheelEvent, target *Element)) js.Func {
	fn := func(this js.Value, args []js.Value) interface{} {
		value := args[0]
		evt := CastWheelEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				ConsolePanicf(r, "Error occurs processing event %q on %q id=%q", evt.Type(), target.TagName(), target.Id())
			}
		}()
		listener(evt, target)
		return js.Undefined()
	}
	return js.FuncOf(fn)
}

// The wheel event fires when the user rotates a wheel button on a pointing device (typically a mouse).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/wheel_event
func (_htmle *Element) AddWheelEvent(listener func(event *WheelEvent, target *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		ConsoleWarnf("AddWheelEvent failed: nil Element or not in DOM")
		return
	}
	evh := makeHTMLElement_WheelEvent(listener)
	_htmle.jsValue.Call("addEventListener", "wheel", evh)
	_htmle.AddEventListener(&eventHandler{eventtype: "wheel", jsHandler: evh})
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
// The element must be in the DOM to
func (_elem *Element) RenderHtml(_unsafeHtmlTemplate string, _data any) (_err error) {
	if !_elem.IsDefined() || !_elem.IsInDOM() {
		ConsoleWarnf("Unable to render Html on nil element or for an element not into the DOM")
		return
	}
	name := _elem.TagName() + "/" + _elem.Id()
	var html string
	unfoldedCmps := make(map[string]HtmlListener, 0)
	html, _err = unfoldComponents(unfoldedCmps, name, _unsafeHtmlTemplate, _data, 0)
	if _err == nil {
		_elem.SetInnerHTML(html)
		addComponentslisteners(unfoldedCmps)
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
		// BUG:
		namedvalue := CastElement(_node.JSValue()).Attributes().Get("data-ick-namedvalue")
		return _name == namedvalue
	})

	for _, node := range children {
		CastElement(node.JSValue()).RenderValue(text)
	}
}

// InsertComponent
func (_elem *Element) InsertNewComponent(_newcmp Composer) (_newcmpid string, _err error) {
	if !_elem.IsDefined() {
		_err = errors.New("InsertNewComponent failed: nil element")
		ConsoleWarnf(_err.Error())
		return "", _err
	}

	// check if _newcmp is a registered component and get its tagname
	cmptype := LookupComponent(reflect.TypeOf(_newcmp))
	if cmptype == "" {
		ConsoleWarnf("InsertNewComponent: inserting a non registered component %q...", reflect.TypeOf(_newcmp).String())
	}

	// create the HTML component into the DOM
	if newcmpelem, err := CreateComponentElement(_newcmp); err == nil {
		_newcmpid = GetNextComponentId(cmptype)
		newcmpelem.SetId(_newcmpid)

		// name the component
		name := newcmpelem.TagName() + "/" + _newcmpid

		// unfold and render html for a composer
		unfoldedCmps := make(map[string]HtmlListener, 0)
		data := TemplateData{
			Me:     _newcmp,
			Global: &GData,
		}
		html, _ := unfoldComponents(unfoldedCmps, name, _newcmp.Template(), data, 0)
		newcmpelem.SetInnerHTML(html)

		// Insert the component element into the DOM
		_elem.PrependNodes(&newcmpelem.Node) //elem.InsertAdjacentHTML(WI_INSIDEFIRST, html)

		// TODO: add style

		// addlisteners
		addComponentslisteners(unfoldedCmps)
		_newcmp.AddListeners()

	} else {
		ConsoleWarnf(err.Error())
		return "", err
	}

	return _newcmpid, nil
}
