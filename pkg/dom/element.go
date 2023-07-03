package dom

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	syscalljs "syscall/js"

	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/js"
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
func (elem *Element) TagName() string {
	if !elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return elem.GetString("tagName")
}

// AttributeString returns the formated list of attributes, ready to use to generate the tag element.
func (elem *Element) AttributeString() string {
	str := ""
	jsa := elem.Get("attributes")
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

// Attribute returns the value of the attribute identified by its name.
// Returns false if the attribute does not exist.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive
func (elem *Element) Attribute(aname string) (string, bool) {
	aname = strings.Trim(aname, " ")
	attr := elem.Call("getAttribute", aname)
	if attr.IsDefined() && attr.String() != "" {
		return attr.String(), true
	}
	return "", false
}

// func (_elem *Element) CreateAttribute(key string, _value any) *Element {
// 	_elem.setAttribute(key, _value, false)
// 	return _elem
// }

// SetAttribute creates an attribute in the map and set its value.
// If the attribute already exists in the map it is updated.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive
func (elem *Element) SetAttribute(key string, value string) *Element {
	elem.setAttribute(key, value, true)
	return elem
}

func (elem *Element) setAttribute(key string, value string, update bool) error {
	key = strings.Trim(key, " ")
	switch strings.ToLower(key) {
	case "id":
		found := elem.Get("id").Type() == js.TYPE_STRING
		if !found || update {
			elem.SetId(value)
		}
	case "class":
		if update {
			elem.ResetClass()
		}
		if value != "" {
			elem.AddClass(value)
		}
	case "tabindex":
		found := elem.Get("tabindex").Type() == js.TYPE_STRING
		if !found || update {
			idx, _ := strconv.Atoi(value)
			elem.SetTabIndex(idx)
		}
	case "style":
		found := elem.Get("style").Type() == js.TYPE_STRING
		if !found || update {
			elem.SetStyle(value)
		}
	default:
		_, err := elem.Check(key)
		if err != nil || update {
			elem.Call("setAttribute", key, value)
		}
	}
	return nil
}

// SetAttributeIf SetAttribute if the condition is true, otherwise remove the attribute.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive.
func (elem *Element) SetAttributeIf(condition bool, aname string, value string) *Element {
	if condition {
		elem.SetAttribute(aname, value)
	} else {
		elem.RemoveAttribute(aname)
	}
	return elem
}

// RemoveAttribute removes the attribute identified by its name.
// Does nothing if the name is not in the map.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive
func (elem *Element) RemoveAttribute(aname string) *Element {
	aname = strings.Trim(aname, " ")
	elem.Call("removeAttribute", aname)
	return elem
}

// ToggleAttribute toggles an attribute like a boolean.
// If the attribute exists it's removed, if it does not exists it's created without value.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive
func (elem *Element) ToggleAttribute(aname string) *Element {
	aname = strings.Trim(aname, " ")
	found := elem.Call("hasAttribute", aname).Bool()
	if found {
		elem.Call("removeAttribute", aname)
	} else {
		elem.Call("setAttribute", aname, string(""))
	}
	return elem
}

// Id rrepresents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (elem *Element) Id() string {
	if !elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return elem.GetString("id")
}

// Id represents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (elem *Element) SetId(id string) *Element {
	id = strings.Trim(id, " ")
	elem.Set("id", id)
	return elem
}

// Name returns the value of the name attribute
func (elem *Element) Name() string {
	if !elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return elem.GetString("name")
}

// SetName sets or overwrites the name attribute. In HTML5 name is case sensitive.
// blanks at the ends of the id are automatically trimmed.
// if name is empty, the name attribute is removed.
func (elem *Element) SetName(name string) *Element {
	name = strings.Trim(name, " ")
	if name == "" {
		elem.Call("removeAttribute", name)
	} else {
		elem.Set("name", name)
	}
	return elem
}

// Classes returns the class object related to _elem.
// If _elem is defined, the class object is wrapped with the DOMTokenList collection of the class attribute of _elem.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/classList
func (_elem *Element) Classes() string {
	return _elem.Get("className").String()
}

// HasClass returns if the class exists in the list of classes.
func (elem *Element) HasClass(class string) bool {
	class = strings.Trim(class, " ")
	if class == "" {
		return false
	}
	return elem.Get("classList").Call("contains", class).Bool()
}

// AddClass adds each class in lists strings to the class attribute.
// Each lists class string can be either a single class or a string of multiple classes separated by spaces.
// Duplicates are not inserted twice.
func (elem *Element) AddClass(lists ...string) *Element {
	for _, list := range lists {
		listf := strings.Fields(list)
		callp := make([]any, len(listf))
		for i, listc := range listf {
			if listc != "" {
				callp[i] = listc
			}
		}
		if len(callp) > 0 {
			elem.Get("classList").Call("add", callp...)
		}
	}
	return elem
}

// AddClassIf adds each c classe to the class attribute if the condition is true.
// Does nothing if the condition is false. See SetClassIf to also remove the class if the condition is false.
func (elem *Element) AddClassIf(condition bool, addlist ...string) *Element {
	if condition {
		elem.AddClass(addlist...)
	}
	return elem
}

// SetClassIf adds each c classe to the class attribute if the condition is true, otherwise remove them.
func (elem *Element) SetClassIf(condition bool, lists ...string) *Element {
	if condition {
		elem.AddClass(lists...)
	} else {
		elem.RemoveClass(lists...)
	}
	return elem
}

// PickClass set the picked class and only that one, removing all others in the classlist.
// If picked is empty or not in the classlist then it's not added.
func (elem *Element) PickClass(classlist string, picked string) *Element {
	elem.RemoveClass(classlist)
	if strings.Contains(classlist, picked) {
		elem.AddClass(picked)
	}
	return elem
}

// ResetClass removes all classes by removing the class attribute
func (elem *Element) ResetClass() *Element {
	elem.Call("removeAttribute", "classList")
	return elem
}

// RemoveClass removes each class in lists strings from the class attribute.
// Each lists class string can be either a single class or a string of multiple classes separated by spaces.
func (elem *Element) RemoveClass(lists ...string) *Element {
	for _, list := range lists {
		listf := strings.Fields(string(list))
		callp := make([]any, len(listf))
		for i, listc := range listf {
			if listc != "" {
				callp[i] = listc
			}
		}
		if len(callp) > 0 {
			elem.Get("classList").Call("remove", callp...)
		}
	}
	return elem
}

// SwitchClass removes one class and set a new one
func (elem *Element) SwitchClasses(remove string, new string) *Element {
	elem.RemoveClass(remove)
	elem.AddClass(new)
	return elem
}

// TabIndex represents the tab order of the current element.
//
// Tab order is as follows:
//  1. Elements with a positive tabindex. Elements that have identical tabindex values should be navigated in the order they appear. Navigation proceeds from the lowest tabindex to the highest tabindex.
//  1. Elements that do not support the tabindex attribute or support it and assign tabindex to 0, in the order they appear.
//  1. Elements that are disabled do not participate in the tabbing order.
//
// Values don't need to be sequential, nor must they begin with any particular value.
// They may even be negative, though each browser trims very large values.
//
// https://developer.mozilla.org/fr/docs/Web/HTML/Global_attributes/tabindex
func (elem *Element) TabIndex() (idx int) {
	found := elem.Call("hasAttribute", "tabindex").Bool()
	if found {
		sidx := elem.Call("getAttribute", "tabindex").String()
		idx, _ = strconv.Atoi(string(sidx))
	}
	return idx
}

// TabIndex represents the tab order of the current element.
//
// Tab order is as follows:
//  1. Elements with a positive tabindex. Elements that have identical tabindex values should be navigated in the order they appear. Navigation proceeds from the lowest tabindex to the highest tabindex.
//  1. Elements that do not support the tabindex attribute or support it and assign tabindex to 0, in the order they appear.
//  1. Elements that are disabled do not participate in the tabbing order.
//
// Values don't need to be sequential, nor must they begin with any particular value.
// They may even be negative, though each browser trims very large values.
//
// https://developer.mozilla.org/fr/docs/Web/HTML/Global_attributes/tabindex
func (elem *Element) SetTabIndex(index int) *Element {
	elem.Call("setAttribute", "tabindex", strconv.Itoa(index))
	return elem
}

// Style returns the style attribute
func (elem *Element) Style() (style string) {
	found := elem.Call("hasAttribute", "style").Bool()
	if found {
		style = elem.Call("getAttribute", "style").String()
	}
	return style
}

// SetStyle sets or overwrites the style attribute.
func (elem *Element) SetStyle(style string) *Element {
	elem.Call("setAttribute", "style", style)
	return elem
}

// AttributeIsTrue returns true if the attribute exists and if it's value is not false nor 0.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive.
func (elem *Element) AttributeIsTrue(aname string) bool {
	aname = strings.Trim(aname, " ")
	found := elem.Call("hasAttribute", aname).Bool()
	if found {
		str := elem.Call("getAttribute", aname).String()
		if strings.ToLower(str) != "false" && str != "0" {
			return true
		}
	}
	return false
}

// SetBoolAttribute sets or overwrites a boolean attribute and returns the map to allow chainning.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name is case sensitive.
func (elem *Element) SetBool(aname string, f bool) *Element {
	aname = strings.Trim(aname, " ")
	if elem.IsDefined() && elem.IsInDOM() {
		if f {
			elem.Call("setAttribute", aname, "")
		} else {
			elem.Call("removeAttribute", aname)
		}
	}
	return elem
}

// Style returns the style attribute
func (elem *Element) IsDisabled() bool {
	return elem.AttributeIsTrue("disabled")
}

// SetDisabled sets the boolean disabled attribute
func (elem *Element) SetDisabled(f bool) *Element {
	elem.SetBool("disabled", f)
	return elem
}

// AttributeString r
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

// TODO: handle Element.InsertRawHTML exceptions
func (_me *Element) InsertRawHTML(_where INSERT_WHERE, _unsafeHtml string) {
	if !_me.IsDefined() {
		return
	}
	switch _where {
	case INSERT_BEFORE_ME:
		_me.Call("insertAdjacentHTML", "beforebegin", _unsafeHtml)
	case INSERT_FIRST_CHILD:
		_me.Call("insertAdjacentHTML", "afterbegin", _unsafeHtml)
	case INSERT_LAST_CHILD:
		_me.Call("insertAdjacentHTML", "beforeend", _unsafeHtml)
	case INSERT_AFTER_ME:
		_me.Call("insertAdjacentHTML", "afterend", _unsafeHtml)
	case INSERT_OUTER:
		_me.Set("outerHTML", _unsafeHtml)
	case INSERT_BODY:
		_me.Set("innerHTML", _unsafeHtml)
	}
}

// InsertHTML unfolds and renders the html of the _html and write it into the DOM.
// All embedded components are wrapped with their DOM element and their listeners are added to the DOM.
// Returns an error if _elem in not the DOM or if an error occurs during UnfoldHtml or mounting process.
// FIXME: InsertHTML
// func (_elem *Element) InsertHTML(_where INSERT_WHERE, htmltemplate html.HTMLString, ds *html.DataState) (_err error) {
// 	if !_elem.IsDefined() || !_elem.IsInDOM() {
// 		return fmt.Errorf("unable to render Html on nil element or for an element not into the DOM")
// 	}

// 	// temparent := &html.HTMLSnippet{}
// 	out := new(bytes.Buffer)
// 	_err = html.RenderHTML(out, nil, htmltemplate)
// 	// _err = html.RenderHTML(out, temparent, htmltemplate)
// 	if _err == nil {
// 		// insert the html element into the dom and wrapit
// 		_elem.InsertRawHTML(_where, out.String())
// 		// mount every embedded components
// 		embedded := html.ComposerMap{}
// 		// embedded := temparent.Embedded()
// 		if embedded != nil {
// 			// DEBUG: console.Warnf("scanning %+v", embedded)
// 			for subid, sub := range embedded {
// 				// look everywhere in the DOM
// 				if sube := Id(subid); sube != nil {
// 					if cmp, ok := sub.(UIComposer); ok {
// 						// DEBUG: console.Warnf("wrapping %+v", w)
// 						_err = mountDeepSnippet(cmp, sube)
// 					}
// 				}
// 			}
// 		} else {
// 			_err = console.Errorf("html string does not have any embedded components")
// 		}
// 	} else {
// 		console.Errorf(_err.Error())
// 	}
// 	return _err
// }

// InsertSnippet unfolds and renders the html of the _snippet and write it into the DOM.
// The _snippet and all its embedded components are wrapped with their DOM element and their listeners are added to the DOM.
// _snippet can be either an HTMLComposer or an UIComposer.
// Returns an error if _elem in not in the DOM or the _snippet has an Id and it's already in the DOM.
// Returns an error if WriteSnippet or mounting process fail.
func (_elem *Element) InsertSnippet(where INSERT_WHERE, cmps ...html.HTMLComposer) (snippetid string, err error) {
	if !_elem.IsDefined() {
		return "", console.Errorf("Element:InsertSnippet failed on undefined element")
	}

	// snippet, ok := composer.(html.HTMLComposer)
	// if !ok {
	// 	return "", console.Errorf("Element:InsertSnippet failed. snippet must implement HTMLComposer interface or UIComposer interface")
	// }

	out := new(bytes.Buffer)
	err = html.Render(out, nil, cmps...)
	if err != nil {
		console.Errorf(err.Error())
		return "", err
	}

	// insert the html element into the dom and wrapit
	_elem.InsertRawHTML(where, out.String())
	if newe := Id(snippetid); newe != nil {
		// wrap the snippet with the fresh new Element and wrap every embedded components with their dom element
		for _, cmp := range cmps {
			if snippet, ok := cmp.(UIComposer); ok {
				err = mountDeepSnippet(snippet, newe)
			} else {
				err = console.Warnf("snippet %q(%v) not mounted, it's not an UIComposer", snippetid, reflect.TypeOf(cmps).String())
			}
		}
	} else {
		err = console.Warnf("snippet %q(%v) not mounted: id not found in the DOM", snippetid, reflect.TypeOf(cmps).String())
	}
	return snippetid, nil
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
