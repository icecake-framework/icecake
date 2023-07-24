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
	"github.com/icecake-framework/icecake/pkg/ickcore"
	"github.com/icecake-framework/icecake/pkg/js"
	"github.com/lolorenzo777/verbose"
)

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

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
func (e *Element) IsDefined() bool {
	if e == nil {
		return false
	}
	return e.JSValue.IsDefined()
}

// Remove removes the element and its listeners from the DOM. Does nothing if e is not defined.
//
// See [Web Api Element.Remove].
//
// [Web Api Element.Remove]: https://developer.mozilla.org/en-US/docs/Web/API/Element/remove
func (e *Element) Remove() {
	if !e.IsDefined() {
		return
	}
	e.RemoveListeners()
	e.Call("remove")
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
	if !elem.IsDefined() {
		return ""
	}
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
	if !elem.IsDefined() {
		return "", false
	}
	aname = strings.Trim(aname, " ")
	attr := elem.Call("getAttribute", aname)
	if attr.IsDefined() && attr.String() != "" {
		return attr.String(), true
	}
	return "", false
}

// func (elem *Element) CreateAttribute(key string, _value any) *Element {
// 	elem.setAttribute(key, _value, false)
// 	return elem
// }

// SetAttribute creates an attribute in the map and set its value.
// If the attribute already exists in the map it is updated.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive
func (elem *Element) SetAttribute(key string, value string) *Element {
	err := elem.setAttribute(key, value, true)
	if err != nil {
		console.Errorf("SetAttribute: %s", err.Error())
	}
	return elem
}

func (elem *Element) setAttribute(key string, value string, update bool) error {
	if !elem.IsDefined() {
		return nil
	}
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
func (elem *Element) SetAttributeIf(condition bool, aname string, valueiftrue string, valueiffalse ...string) *Element {
	if condition {
		elem.SetAttribute(aname, valueiftrue)
	} else if len(valueiffalse) > 0 {
		elem.SetAttribute(aname, valueiffalse[0])
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
	if !elem.IsDefined() {
		return elem
	}
	aname = strings.Trim(aname, " ")
	elem.Call("removeAttribute", aname)
	return elem
}

// ToggleAttribute toggles an attribute like a boolean.
// If the attribute exists it's removed, if it does not exists it's created without value.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive
func (elem *Element) ToggleAttribute(aname string) *Element {
	if !elem.IsDefined() {
		return elem
	}
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
// Returns the UNDEFINED_NODE string if elem is nil or not a defined js value.
// Returns the UNDEFINED_ATTR string if the returned string is empty.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (elem *Element) Id() string {
	if !elem.IsDefined() {
		return UNDEFINED_NODE
	}
	v := elem.GetString("id")
	if v == "" {
		v = UNDEFINED_ATTR
	}
	return v
}

// Id represents the element's identifier, reflecting the id global attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (elem *Element) SetId(id string) *Element {
	if !elem.IsDefined() {
		return elem
	}
	id = strings.Trim(id, " ")
	elem.Set("id", id)
	return elem
}

// Name returns the value of the name attribute.
// Returns the UNDEFINED_NODE string if elem is nil or not a defined js value.
// Returns the UNDEFINED_ATTR string if the returned string is empty.
func (elem *Element) Name() string {
	if !elem.IsDefined() {
		return UNDEFINED_NODE
	}
	v := elem.GetString("name")
	if v == "" {
		v = UNDEFINED_ATTR
	}
	return v
}

// SetName sets or overwrites the name attribute. In HTML5 name is case sensitive.
// blanks at the ends of the id are automatically trimmed.
// if name is empty, the name attribute is removed.
func (elem *Element) SetName(name string) *Element {
	if !elem.IsDefined() {
		return elem
	}
	name = strings.Trim(name, " ")
	if name == "" {
		elem.Call("removeAttribute", name)
	} else {
		elem.Set("name", name)
	}
	return elem
}

// Classes returns the class object related to elem.
// If elem is defined, the class object is wrapped with the DOMTokenList collection of the class attribute of elem.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/classList
func (elem *Element) Classes() string {
	if !elem.IsDefined() {
		return ""
	}
	return elem.Get("className").String()
}

// HasClass returns if the class exists in the list of classes.
func (elem *Element) HasClass(class string) bool {
	if !elem.IsDefined() {
		return false
	}
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
	if !elem.IsDefined() {
		return elem
	}
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
func (elem *Element) SetClassIf(condition bool, classiftrue string, classiffalse ...string) *Element {
	if condition {
		elem.AddClass(classiftrue)
		if len(classiffalse) > 0 {
			elem.RemoveClass(classiffalse[0])
		}
	} else {
		elem.RemoveClass(classiftrue)
		if len(classiffalse) > 0 {
			elem.AddClass(classiffalse[0])
		}
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
	if !elem.IsDefined() {
		return elem
	}
	elem.Call("removeAttribute", "classList")
	return elem
}

// RemoveClass removes each class in lists strings from the class attribute.
// Each lists class string can be either a single class or a string of multiple classes separated by spaces.
func (elem *Element) RemoveClass(lists ...string) *Element {
	if !elem.IsDefined() {
		return elem
	}
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
	if !elem.IsDefined() {
		return 0
	}
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
	if !elem.IsDefined() {
		return elem
	}
	elem.Call("setAttribute", "tabindex", strconv.Itoa(index))
	return elem
}

// Style returns the style attribute
func (elem *Element) Style() (style string) {
	if !elem.IsDefined() {
		return ""
	}
	found := elem.Call("hasAttribute", "style").Bool()
	if found {
		style = elem.Call("getAttribute", "style").String()
	}
	return style
}

// SetStyle sets or overwrites the style attribute.
func (elem *Element) SetStyle(style string) *Element {
	if !elem.IsDefined() {
		return elem
	}
	elem.Call("setAttribute", "style", style)
	return elem
}

// AttributeIsTrue returns true if the attribute exists and if it's value is not false nor 0.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive.
func (elem *Element) AttributeIsTrue(aname string) bool {
	if !elem.IsDefined() {
		return false
	}
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
	if elem.IsDefined() && elem.IsInDOM() {
		aname = strings.Trim(aname, " ")
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
func (elem *Element) InnerHTML() string {
	if !elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return elem.GetString("innerHTML")
}

// OuterHTML gets the serialized HTML fragment describing the element including its descendants.
// It can also be set to replace the element with nodes parsed from the given string.
//
// To only obtain the HTML representation of the contents of an element,
// or to replace the contents of an element, use the innerHTML property instead.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
func (elem *Element) OuterHTML() string {
	if !elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return elem.GetString("outerHTML")
}

// ChildElementCount returns the number of child elements of this element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/childElementCount
func (elem *Element) ChildrenCount() int {
	if !elem.IsDefined() {
		return 0
	}
	return elem.GetInt("childElementCount")
}

// Children returns a live HTMLCollection which contains all of the child elements of the element upon which it was called.
//
// Includes only element nodes. To get all child nodes, including non-element nodes like text and comment nodes, use Node.Children
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/children
func (elem *Element) Children() []*Element {
	if !elem.IsDefined() {
		return make([]*Element, 0)
	}
	elems := elem.Get("children")
	return CastElements(elems)
}

// GetElementsByTagName returns a live HTMLCollection of elements with the given tag name.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName
func (elem *Element) ChildrenByTagName(_tagName string) []*Element {
	if !elem.IsDefined() {
		return make([]*Element, 0)
	}
	elems := elem.Call("getElementsByTagName", _tagName)
	return CastElements(elems)
}

// GetElementsByClassName returns a live HTMLCollection which contains every descendant element which has the specified class name or names.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByClassName
func (elem *Element) ChildrenByClassName(_classNames string) []*Element {
	if !elem.IsDefined() {
		return make([]*Element, 0)
	}
	elems := elem.Call("getElementsByClassName", _classNames)
	return CastElements(elems)
}

// ChildrenMatching make a slice of Element, scaning recursively every children of this Element.
// Only nodes having data attribites matching _data like "data-myvalue"
func (root *Element) ChildrenByData(_data string, _value string) []*Element {
	if !root.IsDefined() || !root.HasChildren() {
		return make([]*Element, 0)
	}
	return root.childrenMatching(999, func(e *Element) bool {
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
func (root *Element) ChildrenMatching(_match func(*Element) bool) []*Element {
	if !root.IsDefined() || !root.HasChildren() {
		return make([]*Element, 0)
	}
	return root.childrenMatching(999, _match)
}

func (root *Element) childrenMatching(_deepmax int, _match func(*Element) bool) []*Element {
	nodes := make([]*Element, 0)

	for _, scan := range root.Children() {
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
func (elem *Element) ChildFirst() *Element {
	if !elem.IsDefined() {
		return new(Element)
	}
	child := elem.Get("firstElementChild")
	return CastElement(child)
}

// LastElementChild returns an element's last child Element, or null if there are no child elements.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/lastElementChild
func (elem *Element) ChildLast() *Element {
	if !elem.IsDefined() {
		return new(Element)
	}
	child := elem.Get("lastElementChild")
	return CastElement(child)
}

// PreviousElementSibling returns the Element immediately prior to the specified one in its parent's children list,
// or null if the specified element is the first one in the list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/previousElementSibling
func (elem *Element) SiblingPrevious() *Element {
	if !elem.IsDefined() {
		return new(Element)
	}
	sibling := elem.Get("previousElementSibling")
	return CastElement(sibling)
}

// NextElementSibling returns the element immediately following the specified one in its parent's children list, or null if the specified element is the last one in the list.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/nextElementSibling
func (elem *Element) SiblingNext() *Element {
	if !elem.IsDefined() {
		return new(Element)
	}
	sibling := elem.Get("nextElementSibling")
	return CastElement(sibling)
}

// Traverses the element and its parents (heading toward the document root) until
// it finds a node that matches the specified CSS selector.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/closest
func (elem *Element) SelectorClosest(_selectors string) *Element {
	if !elem.IsDefined() {
		return new(Element)
	}
	e := elem.Call("closest", _selectors)
	return CastElement(e)
}

// Matches tests whether the element would be selected by the specified CSS selector.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/matches
func (elem *Element) SelectorMatches(_selectors string) bool {
	ok := elem.Call("matches", _selectors)
	return ok.Bool()
}

// QuerySelector returns the first element that is a descendant of the element on which it is invoked
// that matches the specified group of selectors.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
func (elem *Element) SelectorQueryFirst(_selectors string) *Element {
	if !elem.IsDefined() {
		return new(Element)
	}
	e := elem.Call("querySelector", _selectors)
	return CastElement(e)
}

// QuerySelectorAll returns a static (not live) NodeList representing a list of elements matching
// the specified group of selectors which are descendants of the element on which the method was called.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
func (elem *Element) SelectorQueryAll(_selectors string) []*Element {
	if !elem.IsDefined() {
		return make([]*Element, 0)
	}
	elems := elem.Call("querySelectorAll", _selectors)
	return CastElements(elems)
}

// InsertSnippet unfolds and renders the html of the composer and write it into the DOM.
// The composer and all its embedded components are wrapped with their DOM element and their listeners are added to the DOM.
// Composer can be either a simple html ElementComposer or an UIComposer.
//
// Returns an error if elem in not in the DOM or the _snippet has an Id and it's already in the DOM.
// Returns an error if WriteSnippet or mounting process fail.
func (elem *Element) InsertSnippet(where INSERT_WHERE, cmp ickcore.Composer) (errx error) {
	if !elem.IsDefined() {
		return console.Errorf("Element:InsertSnippet failed on undefined element")
	}

	// nothing to render ?
	cmptyp := ""
	if cmp != nil {
		cmptyp = reflect.TypeOf(cmp).String()
	}
	if cmp == nil || reflect.TypeOf(cmp).Kind() != reflect.Ptr || reflect.ValueOf(cmp).IsNil() {
		console.Warnf("InsertSnippet: empty composer %s\n", cmptyp)
		return nil
	}

	if !cmp.NeedRendering() {
		verbose.Printf(verbose.WARNING, "InsertSnippet: composer %s does not need rendering\n", cmptyp)
		return nil
	}

	// rendering html of the composers. custodian embeds all direct childrens.
	rendering := false
	out := new(bytes.Buffer)

	var tag ickcore.Tag
	tb, isnetb := cmp.(ickcore.TagBuilder)
	if isnetb {
		tag = ickcore.BuildTag(tb)
		isnetb = !tag.IsEmpty()
	}
	if isnetb {
		// if the composer is a tag builder then create an element from the composer and insert it into the dom
		// then wrap it. By this way interacting with the DOM of the UIcomposer is possible event if it has no id.
		verbose.Debug("InsertSnippet: rendering L.0 composer %s\n", cmptyp)
		rendering = true
		tagn, _ := tag.TagName()
		newe := CreateElement(tagn)
		for attrn, attrv := range tag.AttributeMap {
			newe.SetAttribute(attrn, attrv)
		}

		// Render the html content
		if cc, iscc := cmp.(ickcore.ContentComposer); iscc {
			errx = cc.RenderContent(out)
			if errx != nil {
				cmp.RMeta().RError = errx
			} else {
				cc.RMeta().IsRender = true
				newe.Set("innerHTML", out.String())
				elem.InsertElement(where, newe)
				_, errx = mountSnippet(cmp, newe)
			}

			// DEBUG: verbose.Debug("InsertSnippet: %+v", cc)
		}
	} else if cc, iscc := cmp.(ickcore.ContentComposer); iscc && cc.NeedRendering() {
		// otherwise insert the rendered snippet html into the dom
		verbose.Printf(verbose.INFO, "InsertSnippet: inserting content")
		rendering = true
		custodian := &ickcore.RMetaData{}
		errx = ickcore.RenderChild(out, custodian, cc)
		if errx == nil {
			elem.InsertRawHTML(where, out.String())
		}

		// mount every embedded components with an ID
		_, errx = mountSnippet(custodian, nil)
	}

	// nor a tag builder, nor a simple contentcomposer, nothing to render
	if !rendering {
		console.Logf("InsertSnippet: composer %s does not need rendering\n", reflect.TypeOf(cmp).String())
	}

	if errx != nil {
		console.Errorf(errx.Error())
	}
	return errx
}

// TODO: dom - handle Element.InsertRawHTML exceptions
func (elem *Element) InsertRawHTML(where INSERT_WHERE, unsafeHtml string) {
	if !elem.IsDefined() {
		return
	}
	switch where {
	case INSERT_BEFORE_ME:
		elem.Call("insertAdjacentHTML", "beforebegin", unsafeHtml)
	case INSERT_FIRST_CHILD:
		elem.Call("insertAdjacentHTML", "afterbegin", unsafeHtml)
	case INSERT_LAST_CHILD:
		elem.Call("insertAdjacentHTML", "beforeend", unsafeHtml)
	case INSERT_AFTER_ME:
		elem.Call("insertAdjacentHTML", "afterend", unsafeHtml)
	case INSERT_OUTER:
		elem.Set("outerHTML", unsafeHtml)
	case INSERT_BODY:
		elem.Set("innerHTML", unsafeHtml)
	}
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

func (_me *Element) InsertElement(_where INSERT_WHERE, elem *Element) {
	if !_me.IsDefined() {
		return
	}
	switch _where {
	case INSERT_BEFORE_ME:
		_me.Call("insertAdjacentElement", "beforebegin", elem.Value())
	case INSERT_FIRST_CHILD:
		_me.Call("prepend", elem.Value())
	case INSERT_LAST_CHILD:
		_me.Call("append", elem.Value())
	case INSERT_AFTER_ME:
		_me.Call("insertAdjacentElement", "afterend", elem.Value())
	case INSERT_OUTER:
		_me.Set("replaceWith", elem.Value())
	case INSERT_BODY:
		// HACK: element.InsertElement - maybe a better way to replace the content
		_me.Set("innerText", "")
		_me.Call("append", elem.Value())
	}
}

// ScrollTop gets or sets the number of pixels that an element's content is scrolled vertically.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollTop
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollWidth
func (elem *Element) ScrollRect() (_rect Rect) {
	if !elem.IsDefined() {
		return
	}
	_rect.X = elem.GetFloat("scrollLeft")
	_rect.Y = elem.GetFloat("scrollTop")
	_rect.Width = elem.GetFloat("scrollWidth")
	_rect.Height = elem.GetFloat("scrollHeight")
	return _rect
}

// ClientRect returns border coordinates of an element in pixels.
//
// Note: This property will round the value to an integer. If you need a fractional value, use element.jsValue.getBoundingClientRect().
//
//   - https://developer.mozilla.org/en-US/docs/Web/API/Element/clientTop
//   - https://developer.mozilla.org/en-US/docs/Web/API/Element/clientLeft
func (elem *Element) ClientRect() (_rect Rect) {
	if !elem.IsDefined() {
		return
	}
	_rect.X = elem.GetFloat("clientLeft")
	_rect.Y = elem.GetFloat("clientTop")
	_rect.Width = elem.GetFloat("clientWidth")
	_rect.Height = elem.GetFloat("clientHeight")
	return _rect
}

// GetBoundingClientRect eturns a DOMRect object providing information about the size of an element and its position relative to the viewport.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
func (elem *Element) BoundingClientRect() (_rect Rect) {
	if !elem.IsDefined() {
		return
	}
	jsrect := elem.Call("getBoundingClientRect")

	_rect.X = jsrect.GetFloat("x")
	_rect.Y = jsrect.GetFloat("y")
	_rect.Width = jsrect.GetFloat("width")
	_rect.Height = jsrect.GetFloat("height")
	return _rect
}

// ScrollIntoView scrolls the element's ancestor containers such that the element on which scrollIntoView() is called is visible to the user.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollIntoView
func (elem *Element) ScrollIntoView() *Element {
	if !elem.IsDefined() {
		return &Element{}
	}
	elem.Call("scrollIntoView")
	return elem
}

// AccessKey A string indicating the single-character keyboard key to give access to the button.
func (elem *Element) AccessKey() string {
	if !elem.IsDefined() {
		return UNDEFINED_NODE
	}
	return elem.GetString("accessKey")
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
func (elem *Element) AddFullscreenEvent(_evttyp event.FULLSCREEN_EVENT, _listener func(*event.Event, *Element)) {
	if !elem.IsDefined() || !elem.IsInDOM() {
		console.Warnf("AddFullscreenEvent not listening on nil Element")
		return
	}
	evh := makehandlerElement_Event(_listener)
	elem.addListener(string(_evttyp), evh)
}

// AddGenericEvent adds Event Listener for a GENERIC_EVENT  on target.
// Returns the function to call to remove and release the listener
func (_htmle *Element) AddGenericEvent(_evttyp event.GENERIC_EVENT, _listener func(*event.Event, *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		console.Warnf("AddGenericEvent failed: nil Element or not in DOM")
		return
	}
	jsevh := makehandlerElement_Event(_listener)
	_htmle.addListener(string(_evttyp), jsevh)
}

// AddClick adds Event Listener for MOUSE_EVENT on target.
// Returns the function to call to remove and release the listener
func (_htmle *Element) AddMouseEvent(_evttyp event.MOUSE_EVENT, listener func(*event.MouseEvent, *Element)) {
	if !_htmle.IsDefined() || !_htmle.IsInDOM() {
		console.Warnf("AddMouseEvent failed: nil Element or not in DOM")
		return
	}
	evh := makehandlerElement_MouseEvent(listener)
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
func makehandlerElement_Event(_listener func(*event.Event, *Element)) syscalljs.Func {
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
func makehandlerElement_MouseEvent(listener func(*event.MouseEvent, *Element)) syscalljs.Func {
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
