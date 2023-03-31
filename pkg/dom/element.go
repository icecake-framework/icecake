package dom

import (
	"bytes"
	"fmt"
	"reflect"

	syscalljs "syscall/js"

	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/event"
	"github.com/sunraylab/icecake/pkg/html"
	"github.com/sunraylab/icecake/pkg/ick"
	"github.com/sunraylab/icecake/pkg/js"
)

/****************************************************************************
* enum used by Element
*****************************************************************************/

type INSERT_WHERE int

const (
	INSERT_BEFORE_ME   INSERT_WHERE = iota // Before the Element itself.
	INSERT_FIRST_CHILD                     // Just inside the element, before its first child.
	INSERT_LAST_CHILD                      // Just inside the element, after its last child.
	INSERT_AFTER_ME                        // After the element itself.
	INSERT_OUTER                           // like outerhtml
	INSERT_BODY                            // like inner html
)

// const (
// 	WI_BEFOREBEGIN WHERE_INSERT = "beforebegin" // Before the Element itself.
// 	WI_INSIDEFIRST WHERE_INSERT = "afterbegin"  // Just inside the element, before its first child.
// 	WI_INSIDELAST  WHERE_INSERT = "beforeend"   // Just inside the element, after its last child.
// 	WI_AFTEREND    WHERE_INSERT = "afterend"    // After the element itself.
// )

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
func CastElement(_jsv js.JSValueProvider) *Element {
	if _jsv.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting Element failed")
		return new(Element)
	}
	cast := new(Element)
	cast.JSValue = _jsv.Value()
	return cast
}

func CastElements(_jsvp js.JSValueProvider) []*Element {
	elems := make([]*Element, 0)
	if _jsvp.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting Elements failed")
		return elems
	}
	len := _jsvp.Value().GetInt("length")
	for i := 0; i < len; i++ {
		_returned := _jsvp.Value().Call("item", uint(i))
		elem := CastElement(_returned)
		elems = append(elems, elem)
	}
	return elems
}

// IsDefined returns true if the Element is not nil AND it's type is not TypeNull and not TypeUndefined
func (_elem *Element) IsDefined() bool {
	if _elem == nil {
		return false
	}
	return _elem.JSValue.IsDefined()
}

// Remove removes the element from the DOM.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/remove
func (_elem *Element) Remove() {
	if !_elem.IsDefined() {
		return
	}
	_elem.RemoveListeners()
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
	return _elem.GetString("tagName")
}

// Id rrepresents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (_elem *Element) Id() string {
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return _elem.GetString("id")
}

// Id rrepresents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (_elem *Element) SetId(_id string) *Element {
	_elem.Set("id", _id)
	return _elem
}

// Classes returns the class object related to _elem.
// If _elem is defined, the class object is wrapped with the DOMTokenList collection of the class attribute of _elem.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/classList
func (_elem *Element) Classes() *JSClasses {
	jsclass := new(JSClasses)
	if _elem.IsDefined() {
		jsclass.owner = _elem
		jslist := _elem.Get("classList")
		jsclass.jslist = &jslist
	}
	return jsclass
}

// Classes returns the class object related to _elem.
// If _elem is defined, the class object is wrapped with the DOMTokenList collection of the class attribute of _elem.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/classList
func (_elem *Element) Attributes() *JSAttributes {
	jsattr := new(JSAttributes)
	if _elem.IsDefined() {
		//jslist := _elem.Get("classList")
		jsattr.owner = _elem
		//_elem.classes.jslist = &jslist
	}
	return jsattr
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
	return _elem.GetString("innerHTML")
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
	return _elem.GetString("outerHTML")
}

// ChildElementCount returns the number of child elements of this element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/childElementCount
func (_elem *Element) ChildrenCount() int {
	if !_elem.IsDefined() {
		return 0
	}
	return _elem.GetInt("childElementCount")
}

// Children returns a live HTMLCollection which contains all of the child elements of the element upon which it was called.
//
// Includes only element nodes. To get all child nodes, including non-element nodes like text and comment nodes, use Node.Children
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/children
func (_elem *Element) Children() []*Element {
	if !_elem.IsDefined() {
		return make([]*Element, 0)
	}
	elems := _elem.Get("children")
	return CastElements(elems)
}

// GetElementsByTagName returns a live HTMLCollection of elements with the given tag name.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName
func (_elem *Element) ChildrenByTagName(_tagName string) []*Element {
	if !_elem.IsDefined() {
		return make([]*Element, 0)
	}
	elems := _elem.Call("getElementsByTagName", _tagName)
	return CastElements(elems)
}

// GetElementsByClassName returns a live HTMLCollection which contains every descendant element which has the specified class name or names.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByClassName
func (_elem *Element) ChildrenByClassName(_classNames string) []*Element {
	if !_elem.IsDefined() {
		return make([]*Element, 0)
	}
	elems := _elem.Call("getElementsByClassName", _classNames)
	return CastElements(elems)
}

// ChildrenMatching make a slice of Element, scaning recursively every children of this Element.
// Only nodes having data attribites matching _data like "data-myvalue"
func (_root *Element) ChildrenByData(_data string, _value string) []*Element {
	if !_root.IsDefined() || !_root.HasChildren() {
		return make([]*Element, 0)
	}
	return _root.childrenMatching(999, func(e *Element) bool {
		v, found := e.Attributes().Attribute(_data)
		if found {
			if v != _value {
				found = false
			}
		}
		return found
	})
}

// ChildrenMatching make a slice of Element, scaning recursively every children of this Element.
// Only nodes matching the optional _match function are included.
func (_root *Element) ChildrenMatching(_match func(*Element) bool) []*Element {
	if !_root.IsDefined() || !_root.HasChildren() {
		return make([]*Element, 0)
	}
	return _root.childrenMatching(999, _match)
}

func (_root *Element) childrenMatching(_deepmax int, _match func(*Element) bool) []*Element {
	nodes := make([]*Element, 0)

	for _, scan := range _root.Children() {
		// apply the filter to children if not too deep and the type node is selected
		if scan.HasChildren() {
			if _deepmax > 0 {
				sub := scan.childrenMatching(_deepmax-1, _match)
				nodes = append(nodes, sub...)
			} else {
				console.Warnf("ChildrenMatching reached max level")
			}
		}

		// apply matching function
		match := false
		if _match != nil {
			match = _match(scan)
		}

		if match {
			nodes = append(nodes, scan)
		}
	}
	return nodes
}

// FirstElementChild  returns an element's first child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/firstElementChild
func (_elem *Element) ChildFirst() *Element {
	if !_elem.IsDefined() {
		return new(Element)
	}
	child := _elem.Get("firstElementChild")
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
	return CastElement(child)
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
	return CastElement(elem)
}

// QuerySelectorAll returns a static (not live) NodeList representing a list of elements matching
// the specified group of selectors which are descendants of the element on which the method was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
func (_elem *Element) SelectorQueryAll(_selectors string) []*Element {
	if !_elem.IsDefined() {
		return make([]*Element, 0)
	}
	elems := _elem.Call("querySelectorAll", _selectors)
	return CastElements(elems)
}

// 	WI_BEFOREBEGIN WHERE_INSERT = "beforebegin" // Before the Element itself.
// 	WI_INSIDEFIRST WHERE_INSERT = "afterbegin"  // Just inside the element, before its first child.
// 	WI_INSIDELAST  WHERE_INSERT = "beforeend"   // Just inside the element, after its last child.
// 	WI_AFTEREND    WHERE_INSERT = "afterend"    // After the element itself.

func (_me *Element) InsertHTML(_where INSERT_WHERE, _unsafeHtml html.HTMLstring) *Element {
	if !_me.IsDefined() {
		return _me
	}
	switch _where {
	case INSERT_BEFORE_ME:
		_me.Call("insertAdjacentHTML", "beforebegin", string(_unsafeHtml))
	case INSERT_FIRST_CHILD:
		_me.Call("insertAdjacentHTML", "afterbegin", string(_unsafeHtml))
	case INSERT_LAST_CHILD:
		_me.Call("insertAdjacentHTML", "beforeend", string(_unsafeHtml))
	case INSERT_AFTER_ME:
		_me.Call("insertAdjacentHTML", "afterend", string(_unsafeHtml))
	case INSERT_OUTER:
		_me.Set("outerHTML", string(_unsafeHtml))
	case INSERT_BODY:
		_me.Set("innerHTML", string(_unsafeHtml))
	}
	return _me
}

// InsertText insert the formated _value as a simple text (not an HTML string) at the _where position.
// The format string follow the fmt rules: https://pkg.go.dev/fmt#hdr-Printing
func (_me *Element) InsertText(_where INSERT_WHERE, _format string, _value ...any) *Element {
	if !_me.IsDefined() {
		return _me
	}
	text := fmt.Sprintf(_format, _value...)
	switch _where {
	case INSERT_BEFORE_ME:
		_me.Call("insertAdjacentText", "beforebegin", text)
	case INSERT_FIRST_CHILD:
		_me.Call("prepend", text)
	case INSERT_LAST_CHILD:
		_me.Call("append", text)
	case INSERT_AFTER_ME:
		_me.Call("insertAdjacentText", "afterend", text)
	case INSERT_OUTER:
		_me.Set("outerHTML", text)
	case INSERT_BODY:
		_me.Set("innerText", text)
	}
	return _me
}

func (_me *Element) InsertElement(_where INSERT_WHERE, _elem *Element) *Element {
	if !_me.IsDefined() {
		return _me
	}
	switch _where {
	case INSERT_BEFORE_ME:
		_me.Call("insertAdjacentElement", "beforebegin", _elem.Value())
	case INSERT_FIRST_CHILD:
		_me.Call("prepend", _elem.Value())
	case INSERT_LAST_CHILD:
		_me.Call("append", _elem.Value())
	case INSERT_AFTER_ME:
		_me.Call("insertAdjacentElement", "afterend", _elem.Value())
	case INSERT_OUTER:
		_me.Set("replaceWith", _elem.Value())
	case INSERT_BODY:
		_me.Set("innerHTML", "")
		_me.Call("append", _elem.Value())
	}
	return _me
}

// InnerHTML ets or sets the HTML or XML markup contained within the element.
//
// To insert the HTML into the document rather than replace the contents of an element, use the method insertAdjacentHTML().
//
// When writing to innerHTML, it will overwrite the content of the source element.
// That means the HTML has to be loaded and re-parsed. This is not very efficient especially when using inside loops.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
// func (_elem *Element) SetInnerHTML(_unsafeHtml string) *Element {
// 	if !_elem.IsDefined() {
// 		return _elem
// 	}
// 	_elem.Set("innerHTML", _unsafeHtml)
// 	return _elem
// }

// OuterHTML gets the serialized HTML fragment describing the element including its descendants.
// It can also be set to replace the element with nodes parsed from the given string.
//
// To only obtain the HTML representation of the contents of an element,
// or to replace the contents of an element, use the innerHTML property instead.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
// func (_elem *Element) SetOuterHTML(_unsafeHtml string) *Element {
// 	if !_elem.IsDefined() {
// 		return _elem
// 	}
// 	_elem.Set("outerHTML", _unsafeHtml)
// 	return _elem
// }

// InnerText represents the rendered text content of a node and its descendants.
//
// InnerText gets pure text, removing any html or css, while TextContent keeps the representation.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerText
// func (_htmle *Element) SetInnerText(value string) *Element {
// 	if !_htmle.IsDefined() {
// 		return &Element{}
// 	}
// 	input := value
// 	_htmle.Set("innerText", input)
// 	return _htmle
// }

// InsertAdjacentElement inserts a given element node at a given position relative to the element it is invoked upon.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentElement
// func (_elem *Element) InsertAdjacentElement(_where WHERE_INSERT, _element *Element) *Element {
// 	if !_elem.IsDefined() {
// 		return new(Element)
// 	}
// 	_elem.Call("insertAdjacentElement", string(_where), _element.JSValue)
// 	return _elem
// }

// InsertAdjacentText given a relative position and a string, inserts a new text node at the given position relative to the element it is called from.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentText
// func (_elem *Element) InsertAdjacentText(_where WHERE_INSERT, _text string) *Element {
// 	if !_elem.IsDefined() {
// 		return &Element{}
// 	}
// 	_elem.Call("insertAdjacentText", string(_where), _text)
// 	return _elem
// }

// InsertAdjacentHTML parses the specified text as HTML or XML and inserts the resulting nodes into the DOM tree at a specified position.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/insertAdjacentHTML
// func (_elem *Element) InsertAdjacentHTML(_where WHERE_INSERT, _text string) *Element {
// 	if !_elem.IsDefined() {
// 		return &Element{}
// 	}
// 	_elem.Call("insertAdjacentHTML", string(_where), _text)
// 	return _elem
// }

// Prepend inserts a set of Node objects or string objects before the first child of the Element.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/prepend
// func (_elem *Element) PrependNodes(_nodes ...*Node) *Element {
// 	if !_elem.IsDefined() {
// 		return &Element{}
// 	}
// 	var args []any = make([]any, len(_nodes))
// 	var end int
// 	for _, n := range _nodes {
// 		if n != nil {
// 			jsn := n.JSValue
// 			args[end] = jsn
// 			end++
// 		}
// 	}
// 	_elem.Call("prepend", args[0:end]...)
// 	return _elem
// }

// Prepend inserts a set of Node objects or string objects before the first child of the Element.
// String objects are inserted as equivalent Text nodes.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/prepend
// func (_elem *Element) PrependStrings(_strs []string) *Element {
// 	if !_elem.IsDefined() {
// 		return &Element{}
// 	}
// 	var args []any = make([]any, len(_strs))
// 	var end int
// 	for _, n := range _strs {
// 		args[end] = n
// 		end++
// 	}
// 	_elem.Call("prepend", args[0:end]...)
// 	return _elem
// }

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
// func (_elem *Element) AppendNodes(_nodes []*Node) *Element {
// 	if !_elem.IsDefined() {
// 		return &Element{}
// 	}
// 	var args []any = make([]any, len(_nodes))
// 	var end int
// 	for _, n := range _nodes {
// 		if n != nil {
// 			jsn := n.JSValue
// 			args[end] = jsn
// 			end++
// 		}
// 	}
// 	_elem.Call("append", args[0:end]...)
// 	return _elem
// }

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
// func (_elem *Element) AppendStrings(_strs ...string) *Element {
// 	if !_elem.IsDefined() {
// 		return &Element{}
// 	}
// 	var args []any = make([]any, len(_strs))
// 	var end int
// 	for _, n := range _strs {
// 		args[end] = n
// 		end++
// 	}
// 	_elem.Call("append", args[0:end]...)
// 	return _elem
// }

// func (_elem *Element) InsertNodesBefore(_nodes []*Node) *Element {
// 	if !_elem.IsDefined() {
// 		return &Element{}
// 	}
// 	var _args []any = make([]any, len(_nodes))
// 	var _end int
// 	for _, n := range _nodes {
// 		if n != nil {
// 			jsn := n.JSValue
// 			_args[_end] = jsn
// 			_end++
// 		}
// 	}
// 	_elem.Call("before", _args[0:_end]...)
// 	return _elem
// }

// func (_elem *Element) InsertNodesAfter(_nodes []*Node) *Element {
// 	if !_elem.IsDefined() {
// 		return &Element{}
// 	}
// 	var _args []any = make([]any, len(_nodes))
// 	var _end int
// 	for _, n := range _nodes {
// 		if n != nil {
// 			jsn := n.JSValue
// 			_args[_end] = jsn
// 			_end++
// 		}
// 	}
// 	_elem.Call("after", _args[0:_end]...)
// 	return _elem
// }

// ScrollTop gets or sets the number of pixels that an element's content is scrolled vertically.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollTop
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollWidth
func (_elem *Element) ScrollRect() (_rect ick.Rect) {
	if !_elem.IsDefined() {
		return
	}
	_rect.X = _elem.GetFloat("scrollLeft")
	_rect.Y = _elem.GetFloat("scrollTop")
	_rect.Width = _elem.GetFloat("scrollWidth")
	_rect.Height = _elem.GetFloat("scrollHeight")
	return _rect
}

// ClientRect returns border coordinates of an element in pixels.
//
// Note: This property will round the value to an integer. If you need a fractional value, use element.jsValue.getBoundingClientRect().
//
//   - https://developer.mozilla.org/en-US/docs/Web/API/Element/clientTop
//   - https://developer.mozilla.org/en-US/docs/Web/API/Element/clientLeft
func (_elem *Element) ClientRect() (_rect ick.Rect) {
	if !_elem.IsDefined() {
		return
	}
	_rect.X = _elem.GetFloat("clientLeft")
	_rect.Y = _elem.GetFloat("clientTop")
	_rect.Width = _elem.GetFloat("clientWidth")
	_rect.Height = _elem.GetFloat("clientHeight")
	return _rect
}

// GetBoundingClientRect eturns a DOMRect object providing information about the size of an element and its position relative to the viewport.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
func (_elem *Element) BoundingClientRect() (_rect ick.Rect) {
	if !_elem.IsDefined() {
		return
	}
	jsrect := _elem.Call("getBoundingClientRect")

	_rect.X = jsrect.GetFloat("x")
	_rect.Y = jsrect.GetFloat("y")
	_rect.Width = jsrect.GetFloat("width")
	_rect.Height = jsrect.GetFloat("height")
	return _rect
}

// ScrollIntoView scrolls the element's ancestor containers such that the element on which scrollIntoView() is called is visible to the user.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollIntoView
func (_elem *Element) ScrollIntoView() *Element {
	if !_elem.IsDefined() {
		return &Element{}
	}
	_elem.Call("scrollIntoView")
	return _elem
}

/****************************************************************************
* HTMLElement's properties & methods
*****************************************************************************/

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (_elem *Element) AccessKey() string {
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return _elem.GetString("accessKey")
}

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (_htmle *Element) SetAccessKey(key bool) *Element {
	if !_htmle.IsDefined() {
		return &Element{}
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
	return _htmle.GetString("innerText")
}

// Focus sets focus on the specified element, if it can be focused. The focused element is the element that will receive keyboard and similar events by default.
//
// By default the browser will scroll the element into view after focusing it,
// and it may also provide visible indication of the focused element (typically by displaying a "focus ring" around the element).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/focus
func (_htmle *Element) Focus() *Element {
	if !_htmle.IsDefined() {
		return &Element{}
	}
	_htmle.Call("focus")
	return _htmle
}

// Blur removes keyboard focus from the current element.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/blur
func (_htmle *Element) Blur() *Element {
	if !_htmle.IsDefined() {
		return &Element{}
	}
	_htmle.Call("blur")
	return _htmle
}

/****************************************************************************
* Element's events
*****************************************************************************/

// event attribute: Event
func makelistenerElement_Event(_listener func(*event.Event, *Element)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Errorf("Error processing event %q on %q(id=%s): %s", evt.Type(), target.TagName(), target.Id(), r)
				console.Stackf()
			}
		}()
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

// AddFullscreenChange is adding doing AddEventListener for 'FullscreenChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_elem *Element) AddFullscreenEvent(_evttyp event.FULLSCREEN_EVENT, _listener func(*event.Event, *Element)) {
	if !_elem.IsDefined() || !_elem.IsInDOM() {
		console.Warnf("AddFullscreenEvent not listening on nil Element")
		return
	}
	evh := makelistenerElement_Event(_listener)
	_elem.addListener(string(_evttyp), evh)
}

/****************************************************************************
* HTMLElement's events
*****************************************************************************/

// event attribute: Event
func makehandler_Element_Event(_listener func(*event.Event, *Element)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Errorf("Error processing event %q on %q(id=%s): %s", evt.Type(), target.TagName(), target.Id(), r)
				console.Stackf()
			}
		}()
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

// AddGenericEvent adds Event Listener for a GENERIC_EVENT  on target.
// Returns the function to call to remove and release the listener
func (_htmle *Element) AddGenericEvent(_evttyp event.GENERIC_EVENT, _listener func(*event.Event, *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		console.Warnf("AddGenericEvent failed: nil Element or not in DOM")
		return
	}
	jsevh := makehandler_Element_Event(_listener)
	_htmle.addListener(string(_evttyp), jsevh)
}

// event attribute: MouseEvent
func makehandler_Element_MouseEvent(listener func(*event.MouseEvent, *Element)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastMouseEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Errorf("Error processing event %q on %q(id=%s): %s", evt.Type(), target.TagName(), target.Id(), r)
				console.Stackf()
			}
		}()
		listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

// AddClick adds Event Listener for MOUSE_EVENT on target.
// Returns the function to call to remove and release the listener
func (_htmle *Element) AddMouseEvent(_evttyp event.MOUSE_EVENT, listener func(*event.MouseEvent, *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		console.Warnf("AddMouseEvent failed: nil Element or not in DOM")
		return
	}
	evh := makehandler_Element_MouseEvent(listener)
	_htmle.addListener(string(_evttyp), evh)
}

// event attribute: FocusEvent
func makelistenerElement_FocusEvent(_listener func(*event.FocusEvent, *Element)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastFocusEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Errorf("Error processing event %q on %q(id=%s): %s", evt.Type(), target.TagName(), target.Id(), r)
				console.Stackf()
			}
		}()
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

// AddBlur is adding doing AddEventListener for 'Blur' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddFocusEvent(_evttyp event.FOCUS_EVENT, listener func(*event.FocusEvent, *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		console.Warnf("AddFocusEvent failed: nil Element or not in DOM")
		return
	}
	evh := makelistenerElement_FocusEvent(listener)
	_htmle.addListener(string(_evttyp), evh)
}

// event attribute: PointerEvent
func makelistenerElement_PointerEvent(_listener func(*event.PointerEvent, *Element)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastPointerEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Errorf("Error processing event %q on %q(id=%s): %s", evt.Type(), target.TagName(), target.Id(), r)
				console.Stackf()
			}
		}()
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

// AddGotPointerCapture is adding doing AddEventListener for 'GotPointerCapture' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddPointerEvent(_evttyp event.POINTER_EVENT, _listener func(*event.PointerEvent, *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		console.Warnf("AddPointerEvent failed: nil Element or not in DOM")
		return
	}
	evh := makelistenerElement_PointerEvent(_listener)
	_htmle.addListener(string(_evttyp), evh)
}

// event attribute: InputEvent
func makelistenerElement_InputEvent(_listener func(*event.InputEvent, *Element)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastInputEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Errorf("Error processing event %q on %q(id=%s): %s", evt.Type(), target.TagName(), target.Id(), r)
				console.Stackf()
			}
		}()
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

// AddInput is adding doing AddEventListener for 'Input' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddInputEvent(_evttyp event.INPUT_EVENT, listener func(*event.InputEvent, *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		console.Warnf("AddInputEvent failed: nil Element or not in DOM")
		return
	}
	evh := makelistenerElement_InputEvent(listener)
	_htmle.addListener(string(_evttyp), evh)
}

// event attribute: KeyboardEvent
func makelistenerElement_KeyboardEvent(_listener func(*event.KeyboardEvent, *Element)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastKeyboardEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Errorf("Error processing event %q on %q(id=%s): %s", evt.Type(), target.TagName(), target.Id(), r)
				console.Stackf()
			}
		}()
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

// AddKeyDown is adding doing AddEventListener for 'KeyDown' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddKeyboard(_evttyp event.KEYBOARD_EVENT, listener func(*event.KeyboardEvent, *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		console.Warnf("AddKeyboard failed: nil Element or not in DOM")
		return
	}
	evh := makelistenerElement_KeyboardEvent(listener)
	_htmle.addListener(string(_evttyp), evh)
}

// event attribute: UIEvent
func makelistenerElement_UIEvent(_listener func(*event.UIEvent, *Element)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastUIEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Errorf("Error processing event %q on %q(id=%s): %s", evt.Type(), target.TagName(), target.Id(), r)
				console.Stackf()
			}
		}()
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

// AddResize is adding doing AddEventListener for 'Resize' on target.
// This method is returning allocated javascript function that need to be released.
func (_htmle *Element) AddResizeEvent(_listener func(*event.UIEvent, *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		console.Warnf("AddResizeEvent failed: nil Element or not in DOM")
		return
	}
	evh := makelistenerElement_UIEvent(_listener)
	_htmle.addListener("resize", evh)
}

// event attribute: WheelEvent
func makeHTMLElement_WheelEvent(_listener func(*event.WheelEvent, *Element)) syscalljs.Func {
	fn := func(this js.JSValue, args []js.JSValue) any {
		value := args[0]
		evt := event.CastWheelEvent(value)
		target := CastElement(value.Get("target"))
		defer func() {
			if r := recover(); r != nil {
				console.Errorf("Error processing event %q on %q(id=%s): %s", evt.Type(), target.TagName(), target.Id(), r)
				console.Stackf()
			}
		}()
		_listener(evt, target)
		return nil
	}
	return js.FuncOf(fn)
}

// The wheel event fires when the user rotates a wheel button on a pointing device (typically a mouse).
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/wheel_event
func (_htmle *Element) AddWheelEvent(_listener func(*event.WheelEvent, *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		console.Warnf("AddWheelEvent failed: nil Element or not in DOM")
		return
	}
	evh := makeHTMLElement_WheelEvent(_listener)
	_htmle.addListener("wheel", evh)
}

/****************************************************************************
* Extra rendering
*****************************************************************************/

// // RenderNamedValue look recursively for any _elem children having the "data-ick-namedvalue" token matching _name
// // and render inner text with the _value
// func (_elem *Element) RenderChildrenValue(_name string, _format string, _value ...any) {
// 	if !_elem.IsDefined() {
// 		return
// 	}
// 	_name = helper.Normalize(_name)
// 	text := fmt.Sprintf(_format, _value...)

// 	children := _elem.FilteredChildren(NT_ELEMENT, func(_node *Node) bool {
// 		// BUG:
// 		namedvalue, _ := CastElement(_node).Attributes().Attribute("data-ick-namedvalue")
// 		return _name == namedvalue
// 	})

// 	for _, node := range children {
// 		CastElement(node).RenderValue(text)
// 	}
// }

// RenderHtml unfolding components if any
// _elem must be in the DOM
func (_elem *Element) RenderHtml(_where INSERT_WHERE, _body html.HTMLstring, _data *html.DataState) error {
	if !_elem.IsDefined() || !_elem.IsInDOM() {
		return fmt.Errorf("unable to render Html on nil element or for an element not into the DOM")
	}

	out := new(bytes.Buffer)
	embedded, err := html.UnfoldHtml(out, _body, _data)
	if err == nil {
		_elem.InsertHTML(_where, html.HTMLstring(out.String()))
		if embedded != nil {
			for _, e := range embedded {
				if l, ok := e.(event.Listener); ok {
					l.AddListeners()
				}
			}
		}
	}
	return err
}

// RenderComponent
func (_elem *Element) RenderSnippet(_where INSERT_WHERE, _snippet any, _data *html.DataState) (_id string, _err error) {
	if !_elem.IsDefined() {
		return "", console.Errorf("RenderComponent: failed on undefined element")
	}

	// render the html element and body, unfolding sub components
	out := new(bytes.Buffer)
	_id, _err = html.WriteHtmlSnippet(out, _snippet, _data)
	if _err == nil {
		// insert the html element into the dom
		_elem.InsertHTML(_where, html.HTMLstring(out.String()))
		if c, ok := _snippet.(html.HtmlComposer); ok {
			embedded := c.Embedded()
			// proceed with embedded components
			if embedded != nil {
				for id, sub := range embedded {
					// look if the id is in the DOM and wrap it to the component
					if w, ok := sub.(js.JSValueWrapper); ok {
						cmpe := Id(id) // look everywhere in the DOM
						if cmpe != nil {
							w.Wrap(cmpe)
						}
					} else {
						console.Errorf("sub component of HtmlComposer %q is not a js wrapper: %s", _id, reflect.TypeOf(sub).String())
					}
				}
			} else {
				console.Errorf("HtmlComposer %q do not have any embedded components", _id)
			}
		} else {
			// can be a simple html component without event handling
			console.Errorf("_snippet %q not an HtmlComposer %v", _id, reflect.TypeOf(_snippet).String())
		}

		if l, ok := _snippet.(event.Listener); ok {
			l.AddListeners()
		}
	}
	return _id, nil
}

// func Render(c HtmlComposer) (_id string, _html HTMLstring) {
// 	out := new(bytes.Buffer)
// 	id, err := RenderHtmlSnippet(out, c, nil)
// 	if err != nil {
// 		log.Printf("error rendering html snippet: %s\n", err.Error())
// 	}
// 	return id, HTMLstring(out.String())
// }
