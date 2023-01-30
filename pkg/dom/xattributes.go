package dom

import (
	"log"
	"strconv"
	"strings"
	"syscall/js"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/lib"
)

/****************************************************************************
* Attributes
*****************************************************************************/

// Attributes represents a set of element's attributes. The subset is static.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
type Attributes struct {
	jsValue      js.Value     // the embedded live collection of all attribute nodes registered to the specified node.
	attributes   []*Attribute // the internal slice of attributes (dom ones)
	ownerElement *Element     // the Element the attributes belongs to.
}

// NewAttributesFromJSNodeMap is casting a js.Value into DOMAttributes.
func NewAttributesFromJSNodeMap(_namedNodeMap js.Value, _ownerElement *Element) (_ret *Attributes) {
	if typ := _namedNodeMap.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	_ret = new(Attributes)
	_ret.jsValue = _namedNodeMap
	_ret.ownerElement = _ownerElement

	// feed the internal slice
	_ret.attributes = make([]*Attribute, 0)
	data := _namedNodeMap.Get("length")
	len := (uint)((data).Int())
	for i := uint(0); i < len; i++ {
		__returned := _namedNodeMap.Call("item", i)
		_result := NewAttributeFromJS(__returned)
		_ret.attributes = append(_ret.attributes, _result)
	}
	return _ret
}

/****************************************************************************
* Attributes's properties
*****************************************************************************/

// Length returns the number of attributes in the list.
func (_attrs *Attributes) Count() int {
	return len(_attrs.attributes)
}

// Item returns an item in the list, determined by its position in the list, its index.
// Returns nil if out of range.
func (_attrs *Attributes) At(_index int) *Attribute {
	if _index >= 0 && _index < len(_attrs.attributes) {
		return _attrs.attributes[_index]
	}
	return nil
}

// Length returns the number of attributes in the list.
func (_attrs *Attributes) OwnerElement() *Element {
	return _attrs.ownerElement
}

// // String returns the value of the list serialized as a string
func (_attrs *Attributes) String() (_ret string) {
	for _, v := range _attrs.attributes {
		_ret += v.String() + " "
	}
	_ret = strings.TrimRight(_ret, " ")
	return _ret
}

// GetAttribue returns the attribute with the given name in the list.
// Returns an nil if Attribute not found
func (_attrs *Attributes) Get(_name string) *Attribute {
	_name = helper.Normalize(_name)
	for _, a := range _attrs.attributes {
		if a.Name() == _name {
			return a
		}
	}
	return nil
}

// Bool returns true if the attribute is set without value or with a value which is not "false" nor "0".
// Returns false if the attribute is set with the value "false", or "0" or if not set.
func (_attrs *Attributes) Bool(_name string) bool {
	_name = helper.Normalize(_name)
	for _, a := range _attrs.attributes {
		if a.Name() == _name {
			return a.Bool()
		}
	}
	return false
}

// SetAttribue adds an attribute in the list. If the attribute already exist it's only updated
//
// if the value is empty, a boolean attribute is set
func (_attrs *Attributes) Set(_name string, _value string) *Attribute {
	_name = helper.Normalize(_name)
	_value = strings.Trim(_value, " ")
	a := _attrs.Get(_name)
	if a == nil {
		a := NewAttributeToDOM(lib.Attribute{Name: _name, Value: _value}, _attrs.ownerElement)
		_attrs.attributes = append(_attrs.attributes, a)
	} else {
		a.Update(_value)
	}
	return a
}

// Remove removes attributes in the list or does nothing for the one that does not exist.
func (_attrs *Attributes) Remove(_name string) (_ret bool) {
	_ret = false
	_name = helper.Normalize(_name)
	for i, a := range _attrs.attributes {
		if a.Name() == _name {
			_attrs.attributes = append(_attrs.attributes[:i], _attrs.attributes[i+1:]...)
			_attrs.ownerElement.jsValue.Call("removeAttribute", string(_name))
			_ret = true
		}
	}
	return _ret
}

// Toggle toggles a boolean attribute (removing it if it is present and adding it if it is not present).
//
// returns the set attribute.
func (_attrs *Attributes) Toggle(_attr string) (_ret *Attributes) {
	_attr = helper.Normalize(_attr)

	if _attrs.Get(_attr) == nil {
		_attrs.Set(_attr, "")
	} else {
		_attrs.Remove(_attr)
	}
	return _attrs
}

/****************************************************************************
* HTMLAttributes
*****************************************************************************/

// HTMLAttributes represents a set of HTMLelement's attributes. The subset is static.
type HTMLAttributes struct {
	Attributes
}

// NewHTMLAttributesFromJSNodeMap is casting a js.Value into DOMAttributes.
func NewHTMLAttributesFromJSNodeMap(_namedNodeMap js.Value, _ownerElement *Element) (_ret *HTMLAttributes) {
	_ret = new(HTMLAttributes)
	_ret.Attributes = *NewAttributesFromJSNodeMap(_namedNodeMap, _ownerElement)
	return _ret
}

// Hidden returning attribute 'hidden'
func (_this *HTMLAttributes) Hidden() bool {
	return _this.Bool("hidden")
}

func (_this *HTMLAttributes) SetHidden(_f bool) *HTMLAttributes {
	if _f {
		_this.Set("hidden", "")
	} else {
		_this.Remove("hidden")
	}
	return _this
}

func (_this *HTMLAttributes) IsDraggable() bool {
	return _this.Bool("draggable")
}

func (_this *HTMLAttributes) SetDraggable(_f bool) *HTMLAttributes {
	if _f {
		_this.Set("draggable", "")
	} else {
		_this.Remove("draggable")
	}
	return _this
}

func (_this *HTMLAttributes) IsSpellcheck() bool {
	return _this.Bool("spellcheck")
}

func (_this *HTMLAttributes) SetSpellcheck(_f bool) *HTMLAttributes {
	if _f {
		_this.Set("spellcheck", "true")
	} else {
		_this.Set("spellcheck", "false")
	}
	return _this
}

// Controls whether and how text input is automatically capitalized as it is entered/edited by the user.
type AUTOCAPITALIZE string

const (
	AUTOCAP_OFF       AUTOCAPITALIZE = "off"
	AUTOCAP_SENTENCES AUTOCAPITALIZE = "sentences"
	AUTOCAP_WORDS     AUTOCAPITALIZE = "words"
	AUTOCAP_CHARS     AUTOCAPITALIZE = "characters"
)

// Autocapitalize returning attribute 'autocapitalize' with
func (_this *HTMLAttributes) Autocapitalize() AUTOCAPITALIZE {
	value := _this.jsValue.Get("autocapitalize").String()
	value = helper.Normalize(value)
	switch value {
	case string(AUTOCAP_OFF), string(AUTOCAP_SENTENCES), string(AUTOCAP_WORDS), string(AUTOCAP_CHARS):
		return AUTOCAPITALIZE(value)
	}
	return "not valid"
}

// SetAutocapitalize setting attribute 'autocapitalize' with
func (_this *HTMLAttributes) SetAutocapitalize(_autocap AUTOCAPITALIZE) *HTMLAttributes {
	switch _autocap {
	case AUTOCAP_OFF, AUTOCAP_SENTENCES, AUTOCAP_WORDS, AUTOCAP_CHARS:
		_this.Set("autocapitalize", string(_autocap))
	default:
		log.Println("SetAutocapitalize fails: not valid value")
	}
	return _this
}

// Controls whether and how text input is automatically capitalized as it is entered/edited by the user.
type CONTENTEDITABLE string

const (
	CONTEDIT_FALSE   CONTENTEDITABLE = "false"
	CONTEDIT_TRUE    CONTENTEDITABLE = "true"
	CONTEDIT_INHERIT CONTENTEDITABLE = "inherit"
)

// ContentEditable returns a boolean value that is true if the contents of the element are editable; otherwise it returns false.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/isContentEditable
func (_this *HTMLAttributes) ContentEditable() CONTENTEDITABLE {
	value := _this.jsValue.Get("contentEditable").String()
	value = helper.Normalize(value)
	switch value {
	case string(CONTEDIT_FALSE), string(CONTEDIT_TRUE), string(CONTEDIT_INHERIT):
		return CONTENTEDITABLE(value)
	}
	return "not valid"
}

// SetContentEditable setting attribute 'contentEditable' with
// type string (idl: DOMString).
func (_this *HTMLAttributes) SetContentEditable(_editable CONTENTEDITABLE) *HTMLAttributes {
	switch _editable {
	case CONTEDIT_FALSE, CONTEDIT_TRUE, CONTEDIT_INHERIT:
		_this.Set("contentEditable", string(_editable))
	default:
		log.Println("contentEditable fails: not a valid value")
	}
	return _this
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
func (_this *HTMLAttributes) TabIndex() int {
	stri := _this.Get("tabIndex")
	i, _ := strconv.Atoi(stri.Value())
	return i
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
func (_this *HTMLAttributes) SetTabIndex(_index int) *HTMLAttributes {
	_this.Set("tabIndex", strconv.Itoa(_index))
	return _this
}

// Title represents the title of the element: the text usually displayed in a 'tooltip' popup when the mouse is over the node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/title
func (_this *HTMLAttributes) Title() string {
	return _this.Get("title").String()
}

// Title represents the title of the element: the text usually displayed in a 'tooltip' popup when the mouse is over the node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/title
func (_this *HTMLAttributes) SetTitle(value string) *HTMLAttributes {
	_this.Set("title", value)
	return _this
}
