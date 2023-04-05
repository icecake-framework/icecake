package dom

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	syscalljs "syscall/js"

	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/event"
	"github.com/sunraylab/icecake/pkg/html"
	"github.com/sunraylab/icecake/pkg/ick"
	"github.com/sunraylab/icecake/pkg/js"
)

/****************************************************************************
* Enum
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

/******************************************************************************
* Element's Properties & Methods
*******************************************************************************/

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

// TabIndex represents the tab order of the current element.
//
// Tab order is as follows:
//  1. Elements with a positive tabIndex. Elements that have identical tabIndex values should be navigated in the order they appear. Navigation proceeds from the lowest tabIndex to the highest tabIndex.
//  1. Elements that do not support the tabIndex attribute or support it and assign tabIndex to 0, in the order they appear.
//  1. Elements that are disabled do not participate in the tabbing order.
//
// Values don't need to be sequential, nor must they begin with any particular value.
// They may even be negative, though each browser trims very large values.
//
// https://developer.mozilla.org/fr/docs/Web/HTML/Global_attributes/tabindex
func (_elem *Element) TabIndex() (_idx int) {
	found := _elem.Call("hasAttribute", "tabIndex").Bool()
	if found {
		sidx := _elem.Call("getAttribute", "tabIndex").String()
		_idx, _ = strconv.Atoi(string(sidx))
	}
	return _idx
}

// Classes returns the class object related to _elem.
// If _elem is defined, the class object is wrapped with the DOMTokenList collection of the class attribute of _elem.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/classList
func (_elem *Element) Classes() string {
	return _elem.Get("className").String()
}

// Style returns the style attribute
func (_elem *Element) Style() (_style string) {
	found := _elem.Call("hasAttribute", "style").Bool()
	if found {
		_style = _elem.Call("getAttribute", "style").String()
	}
	return _style
}

// Style returns the style attribute
func (_elem *Element) IsDisabled() bool {
	found := _elem.Call("hasAttribute", "disabled").Bool()
	if found {
		str := _elem.Call("getAttribute", "disabled").String()
		if strings.ToLower(str) != "false" && str != "0" {
			return true
		}
	}
	return false
}

func (_elem *Element) Attributes() string {
	str := ""
	jsa := _elem.Get("attributes")
	len := jsa.GetInt("length")
	for i := 0; i < len; i++ {
		jsi := jsa.Call("item", i)
		str += jsi.GetString("name")
		value := jsi.GetString("value")
		if value != "" {
			delim := "'"
			if strings.ContainsRune(value, rune('\'')) {
				delim = "\""
			}
			str += `=` + delim + value + delim
		}
		str += " "
	}
	return strings.TrimRight(str, " ")
}

// Id rrepresents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (_elem *Element) SetId(_id string) *Element {
	_elem.Set("id", _id)
	return _elem
}

func (_elem *Element) SetDisabled(_f bool) {
	if _f {
		_elem.Call("setAttribute", "disabled", "")
	} else {
		_elem.Call("removeAttribute", "disabled")
	}
}

func (_elem *Element) SetStyle(_style html.String) *Element {
	_elem.Call("setAttribute", "style", string(_style))
	return _elem
}

// TabIndex represents the tab order of the current element.
//
// Tab order is as follows:
//  1. Elements with a positive tabIndex. Elements that have identical tabIndex values should be navigated in the order they appear. Navigation proceeds from the lowest tabIndex to the highest tabIndex.
//  1. Elements that do not support the tabIndex attribute or support it and assign tabIndex to 0, in the order they appear.
//  1. Elements that are disabled do not participate in the tabbing order.
//
// Values don't need to be sequential, nor must they begin with any particular value.
// They may even be negative, though each browser trims very large values.
//
// https://developer.mozilla.org/fr/docs/Web/HTML/Global_attributes/tabindex
func (_elem *Element) SetTabIndex(_index int) *Element {
	_elem.Call("setAttribute", "tabIndex", strconv.Itoa(_index))
	return _elem
}

func (_elem *Element) ResetClasses(_list html.String) *Element {
	_elem.Set("className", string(_list))
	return _elem
}

func (_elem *Element) SetClasses(_list html.String) *Element {
	listf := strings.Fields(string(_list))
	callp := make([]any, len(listf))
	for i, listc := range listf {
		if listc != "" {
			callp[i] = listc
		}
	}
	if len(callp) > 0 {
		_elem.Get("classList").Call("add", callp...)
	}
	return _elem
}

func (_elem *Element) RemoveClasses(_list string) *Element {
	listf := strings.Fields(string(_list))
	callp := make([]any, len(listf))
	for i, listc := range listf {
		if listc != "" {
			callp[i] = listc
		}
	}
	if len(callp) > 0 {
		_elem.Get("classList").Call("remove", callp...)
	}
	return _elem
}

func (_elem *Element) SwitchClasses(_remove string, _new html.String) *Element {
	_elem.RemoveClasses(_remove)
	_elem.SetClasses(_new)
	return _elem
}

func (_elem *Element) HasClass(_class string) bool {
	_class = strings.Trim(_class, " ")
	if _class == "" {
		return false
	}
	return _elem.Get("classList").Call("contains", _class).Bool()
}

func (_elem *Element) Attribute(_key string) (string, bool) {
	_key = strings.Trim(_key, " ")
	attr := _elem.Call("getAttribute", _key)
	if attr.IsDefined() && attr.String() != "" {
		return attr.String(), true
	}
	return "", false
}

func (_elem *Element) CreateAttribute(_key string, _value any) *Element {
	_elem.setAttribute(_key, _value, false)
	return _elem
}

func (_elem *Element) SetAttribute(_key string, _value any) *Element {
	_elem.setAttribute(_key, _value, true)
	return _elem
}

func (_elem *Element) setAttribute(_key string, _value any, overwrite bool) error {
	_key = strings.Trim(_key, " ")
	switch strings.ToLower(_key) {
	case "id":
		found := _elem.Get("id").Type() == js.TYPE_STRING
		if !found || overwrite {
			switch v := _value.(type) {
			case string:
				_elem.SetId(v)
			case html.String:
				_elem.SetId(string(v))
			default:
				return errors.New("wrong value type for id")
			}
		}
	case "tabindex":
		found := _elem.Get("tabIndex").Type() == js.TYPE_STRING
		if !found || overwrite {
			switch v := _value.(type) {
			case string:
				idx, _ := strconv.Atoi(v)
				_elem.SetTabIndex(idx)
			case html.String:
				idx, _ := strconv.Atoi(string(v))
				_elem.SetTabIndex(idx)
			case int:
				_elem.SetTabIndex(v)
			case uint:
				_elem.SetTabIndex(int(v))
			case float32:
				_elem.SetTabIndex(int(v))
			case float64:
				_elem.SetTabIndex(int(v))
			default:
				return errors.New("wrong value type for tabindex")
			}
		}
	case "class":
		var lst html.String
		switch v := _value.(type) {
		case string:
			lst = html.String(v)
		case html.String:
			lst = v
		default:
			return errors.New("wrong value type for class")
		}
		if overwrite {
			_elem.ResetClasses(lst)
		} else if _value != "" {
			_elem.SetClasses(lst)
		}
	case "style":
		// TODO: handle style update to not overwrite
		found := _elem.Get("style").Type() == js.TYPE_STRING
		if !found || overwrite {
			var style html.String
			switch v := _value.(type) {
			case string:
				style = html.String(v)
			case html.String:
				style = v
			default:
				return errors.New("wrong value type for class")
			}
			_elem.SetStyle(style)
		}
	default:
		_, err := _elem.Check(_key)
		if err != nil || overwrite {
			var strv html.String
			switch v := _value.(type) {
			case string:
				strv = html.String(v)
			case html.String:
				strv = v
			case bool:
				if v {
					strv = ""
				} else {
					_elem.Call("removeAttribute", _key)
					break
				}
			case int, uint, float32, float64:
				strv = html.String(fmt.Sprintf("%v", v))
			default:
				return errors.New("wrong value type for " + _key)
			}
			_elem.Call("setAttribute", _key, string(strv))
		}
	}
	return nil
}

func (_elem *Element) RemoveAttribute(_key string) *Element {
	_key = strings.Trim(_key, " ")
	_elem.Call("removeAttribute", _key)
	return _elem
}

func (_elem *Element) ToggleAttribute(_key string) *Element {
	_key = strings.Trim(_key, " ")
	found := _elem.Call("hasAttribute", _key).Bool()
	if found {
		_elem.Call("removeAttribute", _key)
	} else {
		_elem.Call("setAttribute", _key, string(""))
	}
	return _elem
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
		v, found := e.Attribute(_data)
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

// TODO: handle exceptions InsertRawHTML
func (_me *Element) InsertRawHTML(_where INSERT_WHERE, _unsafeHtml html.String) {
	if !_me.IsDefined() {
		return
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
}

// InsertHTML unfolds and renders the html of the _html and write it into the DOM.
// All embedded components are wrapped with their DOM element and their listeners are added to the DOM.
// Returns an error if _elem in not the DOM or if an error occurs during UnfoldHtml or mounting process.
// HACK: better rendering with a reader ?
func (_elem *Element) InsertHTML(_where INSERT_WHERE, _html html.String, _data *html.DataState) (_err error) {
	if !_elem.IsDefined() || !_elem.IsInDOM() {
		return fmt.Errorf("unable to render Html on nil element or for an element not into the DOM")
	}

	var embedded map[string]any
	out := new(bytes.Buffer)
	embedded, _err = html.UnfoldHtml(out, _html, _data)
	if _err == nil {
		// insert the html element into the dom and wrapit
		_elem.InsertRawHTML(_where, html.String(out.String()))
		// mount every embedded components
		if embedded != nil {
			// DEBUG: console.Warnf("scanning %+v", embedded)
			for subid, sub := range embedded {
				// look everywhere in the DOM
				if sube := Id(subid); sube != nil {
					if cmp, ok := sub.(UIComposer); ok {
						// DEBUG: console.Warnf("wrapping %+v", w)
						_err = mountDeepSnippet(cmp, sube)
					}
				}
			}
		} else {
			_err = console.Errorf("html string does not have any embedded components")
		}
	} else {
		console.Errorf(_err.Error())
	}
	return _err
}

// InsertSnippet unfolds and renders the html of the _snippet and write it into the DOM.
// The _snippet and all its embedded components are wrapped with their DOM element and their listeners are added to the DOM.
// _snippet can be either an HTMLComposer or an UIComposer.
// Returns an error if _elem in not in the DOM or the _snippet has an Id and it's already in the DOM.
// Returns an error if WriteHTMLSnippet or mounting process fail.
func (_elem *Element) InsertSnippet(_where INSERT_WHERE, _snippet any, _data *html.DataState) (_id string, _err error) {
	if !_elem.IsDefined() {
		return "", console.Errorf("Element:InsertSnippet failed on undefined element")
	}

	snippet, ok := _snippet.(html.HTMLComposer)
	if !ok {
		return "", console.Errorf("Element:InsertSnippet failed. snippet must implement HTMLComposer interface or UIComposer interface")
	}
	snippetid := snippet.Id()
	if snippetid != "" {
		if _, err := Doc().CheckId(snippet.Id()); err == nil {
			return "", console.Errorf("Element:InsertSnippet failed. snippet's ID %q is already in the DOM.", snippetid)
		}
	}

	out := new(bytes.Buffer)
	_id, _err = html.WriteHTMLSnippet(out, snippet, _data)
	if _err == nil {
		// insert the html element into the dom and wrapit
		_elem.InsertRawHTML(_where, html.String(out.String()))
		if newe := Id(_id); newe != nil {
			// wrap the snippet with the fresh new Element and wrap every embedded components with their dom element
			if snippet, ok := _snippet.(UIComposer); ok {
				_err = mountDeepSnippet(snippet, newe)
			} else {
				_err = console.Warnf("snippet %q(%v) not mounted, it's not an UIComposer", _id, reflect.TypeOf(_snippet).String())
			}
		} else {
			_err = console.Warnf("snippet %q(%v) not mounted: id not found in the DOM", _id, reflect.TypeOf(_snippet).String())
		}
	} else {
		console.Errorf(_err.Error())
	}
	return _id, nil
}

// InsertText insert the formated _value as a simple text (not an HTML string) at the _where position.
// The format string follow the fmt rules: https://pkg.go.dev/fmt#hdr-Printing
func (_me *Element) InsertText(_where INSERT_WHERE, _format string, _value ...any) {
	if !_me.IsDefined() {
		return
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
}

func (_me *Element) InsertElement(_where INSERT_WHERE, _elem *Element) {
	if !_me.IsDefined() {
		return
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
}

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

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (_elem *Element) AccessKey() string {
	if !_elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return _elem.GetString("accessKey")
}

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (_htmle *Element) SetAccessKey(key string) *Element {
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
* Event handling
*****************************************************************************/

// AddFullscreenChange is adding doing AddEventListener for 'FullscreenChange' on target.
// This method is returning allocated javascript function that need to be released.
func (_elem *Element) AddFullscreenEvent(_evttyp event.FULLSCREEN_EVENT, _listener func(*event.Event, *Element)) {
	if !_elem.IsDefined() || !_elem.IsInDOM() {
		console.Warnf("AddFullscreenEvent not listening on nil Element")
		return
	}
	evh := makehandler_Element_Event(_listener)
	_elem.addListener(string(_evttyp), evh)
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

/******************************************************************************
* Private Area
*******************************************************************************/

// Event generic
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

// MouseEvent
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

// FocusEvent
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

// PointerEvent
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

// InputEvent
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

// KeyboardEvent
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

// UIEvent
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

// WheelEvent
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
