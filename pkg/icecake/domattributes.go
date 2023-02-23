package ick

import (
	"log"
	"strconv"
	"strings"

	"github.com/sunraylab/icecake/internal/helper"
)

/****************************************************************************
* Attributes
*****************************************************************************/

// Attributes represents a set of element's attributes. The subset is static.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
type Attributes struct {
	//jsValue      js.Value     // the embedded live collectio (a NamedNodeMap) of all attribute nodes registered to the specified node.
	ownerElement *Element     // the Element the attributes belongs to.
	attributes   []*Attribute // the internal slice of attributes (dom ones)
}

func NewAttributes(_ownerElement *Element) *Attributes {
	attrs := new(Attributes)
	attrs.ownerElement = _ownerElement
	attrs.attributes = make([]*Attribute, 0)
	return attrs
}

// NewAttributesFromJSNodeMap is casting a js.Value into DOMAttributes.
// func NewAttributesFromJSNodeMap(_namedNodeMap js.Value, _ownerElement *Element) (_ret *Attributes) {
// 	if typ := _namedNodeMap.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
// 		return nil
// 	}
// 	_ret = new(Attributes)
// //	_ret.jsValue = _namedNodeMap
// 	_ret.ownerElement = _ownerElement

// 	// feed the internal slice
// 	_ret.attributes = make([]*Attribute, 0)
// 	data := _namedNodeMap.Get("length")
// 	len := (uint)((data).Int())
// 	for i := uint(0); i < len; i++ {
// 		__returned := _namedNodeMap.Call("item", i)
// 		_result := CastAttribute(__returned)
// 		_ret.attributes = append(_ret.attributes, _result)
// 	}
// 	return _ret
// }

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
func (_attrs *Attributes) String() (_str string) {
	for _, v := range _attrs.attributes {
		_str += v.String() + " "
	}
	return strings.TrimRight(_str, " ")
}

// Bool returns true if the attribute is set without value or with a value which is not "false" nor "0".
// Returns false if the attribute is set with the value "false", or "0" or if not set.
func (_attrs *Attributes) IsTrue(_name string) bool {
	_name = helper.Normalize(_name)
	return _attrs.isTrue(_name)
}

func (_attrs *Attributes) isTrue(_name string) bool {
	for _, a := range _attrs.attributes {
		if a.Name() == _name {
			return a.IsTrue()
		}
	}
	return false
}

// GetAttribue returns the attribute with the given name in the list.
// Returns an nil if Attribute not found
func (_attrs *Attributes) Get(_name string) *Attribute {
	_name = helper.Normalize(_name)
	return _attrs.get(_name)
}

func (_attrs *Attributes) get(_name string) *Attribute {
	for _, a := range _attrs.attributes {
		if a.Name() == _name {
			return a
		}
	}
	return nil
}

// GetAttribue returns the attribute with the given name in the list.
// Returns an nil if Attribute not found
func (_attrs *Attributes) GetValue(_name string) string {
	_name = helper.Normalize(_name)
	return _attrs.getValue(_name)
}

func (_attrs *Attributes) getValue(_name string) string {
	for _, a := range _attrs.attributes {
		if a.Name() == _name {
			return a.value
		}
	}
	return ""
}

// SetAttribue adds an attribute in the list. If the attribute already exist it's only updated
//
// if the value is empty, a boolean attribute is set
func (_attrs *Attributes) Set(_name string, _value string) *Attribute {
	_name = helper.Normalize(_name)
	_value = strings.Trim(_value, " ")
	return _attrs.set(_name, _value)
}

func (_attrs *Attributes) set(_name string, _value string) *Attribute {
	a := _attrs.get(_name)
	if a == nil {
		a := &Attribute{name: _name, value: _value, ownerElement: _attrs.ownerElement}
		_attrs.attributes = append(_attrs.attributes, a)
	}
	a.Update(_value)
	return a
}

// Remove removes attributes in the list or does nothing for the one that does not exist.
func (_attrs *Attributes) Remove(_name string) bool {
	_name = helper.Normalize(_name)
	return _attrs.remove(_name)
}

func (_attrs *Attributes) remove(_name string) bool {
	removed := false
	for i, a := range _attrs.attributes {
		if a.Name() == _name {
			_attrs.attributes = append(_attrs.attributes[:i], _attrs.attributes[i+1:]...)
			_attrs.ownerElement.Call("removeAttribute", string(_name))
			removed = true
		}
	}
	return removed
}

// Toggle toggles a boolean attribute (removing it if it is present and adding it if it is not present).
//
// returns the set attribute.
func (_attrs *Attributes) Toggle(_name string) *Attributes {
	_name = helper.Normalize(_name)
	if _attrs.get(_name) == nil {
		_attrs.set(_name, "")
	} else {
		_attrs.remove(_name)
	}
	return _attrs
}

// returns an subset of _attrs with only "data-*" attributes
func (_attrs *Attributes) Dataset() *Attributes {
	dataset := new(Attributes)
	dataset.ownerElement = _attrs.ownerElement
	for _, a := range _attrs.attributes {
		if len(a.Name()) > 5 && strings.HasPrefix(a.Name(), "data-") {
			dataset.attributes = append(dataset.attributes, a)
		}
	}
	return dataset
}

// Title represents the title of the element: the text usually displayed in a 'tooltip' popup when the mouse is over the node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/title
func (_attrs *Attributes) Title() string {
	return _attrs.getValue("title")
}

// Title represents the title of the element: the text usually displayed in a 'tooltip' popup when the mouse is over the node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/title
func (_attrs *Attributes) SetTitle(title string) *Attributes {
	title = helper.Normalize(title)
	_attrs.Set("title", title)
	return _attrs
}

// Hidden returning attribute 'hidden'
func (_attrs *Attributes) IsHidden() bool {
	return _attrs.isTrue("hidden")
}

func (_attrs *Attributes) SetHidden(_f bool) *Attributes {
	if _f {
		_attrs.set("hidden", "")
	} else {
		_attrs.remove("hidden")
	}
	return _attrs
}

func (_attrs *Attributes) IsDraggable() bool {
	return _attrs.isTrue("draggable")
}

func (_attrs *Attributes) SetDraggable(_f bool) *Attributes {
	if _f {
		_attrs.set("draggable", "")
	} else {
		_attrs.remove("draggable")
	}
	return _attrs
}

func (_attrs *Attributes) IsSpellcheck() bool {
	return _attrs.isTrue("spellcheck")
}

func (_attrs *Attributes) SetSpellcheck(_f bool) *Attributes {
	if _f {
		_attrs.set("spellcheck", "true")
	} else {
		_attrs.set("spellcheck", "false")
	}
	return _attrs
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
func (_attrs *Attributes) Autocapitalize() AUTOCAPITALIZE {
	autocap := _attrs.getValue("autocapitalize")
	switch autocap {
	case string(AUTOCAP_OFF), string(AUTOCAP_SENTENCES), string(AUTOCAP_WORDS), string(AUTOCAP_CHARS):
		return AUTOCAPITALIZE(autocap)
	}
	return "not valid"
}

// SetAutocapitalize setting attribute 'autocapitalize' with
func (_attrs *Attributes) SetAutocapitalize(_autocap AUTOCAPITALIZE) *Attributes {
	switch _autocap {
	case AUTOCAP_OFF, AUTOCAP_SENTENCES, AUTOCAP_WORDS, AUTOCAP_CHARS:
		_attrs.Set("autocapitalize", string(_autocap))
	default:
		log.Println("SetAutocapitalize failed: not valid value")
	}
	return _attrs
}

// Controls whether and how text input is automatically capitalized as it is entered/edited by the user.
type CONTENT_EDITABLE string

const (
	CONTEDIT_FALSE   CONTENT_EDITABLE = "false"
	CONTEDIT_TRUE    CONTENT_EDITABLE = "true"
	CONTEDIT_INHERIT CONTENT_EDITABLE = "inherit"
)

// ContentEditable returns a boolean value that is true if the contents of the element are editable; otherwise it returns false.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/isContentEditable
func (_attrs *Attributes) ContentEditable() CONTENT_EDITABLE {
	editable := _attrs.GetValue("contentEditable")
	switch editable {
	case string(CONTEDIT_FALSE), string(CONTEDIT_TRUE), string(CONTEDIT_INHERIT):
		return CONTENT_EDITABLE(editable)
	}
	return "not valid"
}

// SetContentEditable setting attribute 'contentEditable' with
// type string (idl: DOMString).
func (_attrs *Attributes) SetContentEditable(_editable CONTENT_EDITABLE) *Attributes {
	switch _editable {
	case CONTEDIT_FALSE, CONTEDIT_TRUE, CONTEDIT_INHERIT:
		_attrs.Set("contentEditable", string(_editable))
	default:
		log.Println("contentEditable fails: not a valid value")
	}
	return _attrs
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
func (_attrs *Attributes) TabIndex() int {
	stri := _attrs.getValue("tabIndex")
	i, _ := strconv.Atoi(stri)
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
func (_attrs *Attributes) SetTabIndex(_index int) *Attributes {
	_attrs.Set("tabIndex", strconv.Itoa(_index))
	return _attrs
}
