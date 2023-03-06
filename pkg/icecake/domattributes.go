package ick

import (
	"fmt"
	"log"
	"sort"
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
	ownerElement *Element          // the Element the attributes belongs to. The DOM is only updated when ownerElement is not nil.
	attributes   map[string]string // the internal map of attributes
}

func NewAttributes(_ownerElement *Element) *Attributes {
	attrs := new(Attributes)
	attrs.ownerElement = _ownerElement
	attrs.attributes = make(map[string]string, 0)
	return attrs
}

func (_attrs Attributes) Sort() []string {
	s := make([]string, 0, len(_attrs.attributes))
	for k := range _attrs.attributes {
		s = append(s, k)
	}
	sort.Strings(s)
	return s
}

/****************************************************************************
* Attributes's properties
*****************************************************************************/

// Length returns the number of attributes in the list.
func (_attrs Attributes) OwnerElement() *Element {
	return _attrs.ownerElement
}

// returns an subset of _attrs with only "data-*" attributes
func (_attrs Attributes) Dataset() map[string]string {
	dataset := make(map[string]string)
	for name, value := range _attrs.attributes {
		if len(name) > 5 && strings.HasPrefix(name, "data-") {
			dataset[name] = value
		}
	}
	return dataset
}

// Length returns the number of attributes in the list.
func (_attrs Attributes) Count() int {
	return len(_attrs.attributes)
}

// GetAttribue returns the attribute with the given name in the list.
// Returns an nil if Attribute not found.
// Case sensitive and must be trimed.
func (_attrs Attributes) Get(_name string) string {
	return _attrs.attributes[_name]
}

// SetAttribue adds an attribute in the list. If the attribute already exist it's updated.
// An empty value means a boolean attribute set to true
//
// case sensitive.
//
// Name and Value will be trimed. Quotes delimiters of the value will be removed if any.
func (_attrs *Attributes) Set(_name string, _value string) {
	_name = strings.Trim(_name, " ")
	_value = strings.Trim(strings.Trim(_value, "'"), "\"")
	_attrs.attributes[_name] = _value
	if _attrs.ownerElement != nil {
		_attrs.ownerElement.Call("setAttribute", _name, _value)
	}
}

// Remove removes attributes in the list or does nothing for the one that does not exist.
// Case sensitive.
func (_attrs *Attributes) Remove(_name string) {
	_name = strings.Trim(_name, " ")
	delete(_attrs.attributes, _name)
	if _attrs.ownerElement != nil {
		_attrs.ownerElement.Call("removeAttribute", string(_name))
	}
}

// String returns the value of the list serialized as a string
func (_attrs Attributes) String() (_str string) {
	for name, value := range _attrs.attributes {
		_str += strings.Trim(name, " ")
		if value != "" {
			delim := "'"
			if strings.ContainsRune(value, rune('\'')) {
				delim = "\""
			}
			_str += `=` + delim + value + delim
		}
		_str += " "
	}
	return strings.TrimRight(_str, " ")
}

// Toggle toggles a boolean attribute (removing it if it's present and adding it if it's not present).
//
// Case sensitive.
func (_attrs *Attributes) Toggle(_name string) {
	_name = strings.Trim(_name, " ")
	_, found := _attrs.attributes[_name]
	if found {
		_attrs.Remove(_name)
	} else {
		_attrs.Set(_name, "")
	}
}

// returns true if the attribute is set and its value is not false nor 0.
// case sensitive, must be trimes.
func (_attrs Attributes) IsTrue(_name string) bool {
	val, found := _attrs.attributes[_name]
	val = strings.ToLower(val)
	if !found || val == "false" || val == "0" {
		return false
	}
	return true
}

// Hidden returns boolean attribute 'hidden'
func (_attrs Attributes) Hidden() bool {
	return _attrs.IsTrue("hidden")
}

// Draggable returns boolean attribute 'draggable'
func (_attrs Attributes) Draggable() bool {
	return _attrs.IsTrue("draggable")
}

// Spellcheck returns boolean attribute 'spellcheck'
func (_attrs Attributes) Spellcheck() bool {
	return _attrs.IsTrue("spellcheck")
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
func (_attrs Attributes) TabIndex() int {
	stri := _attrs.attributes["tabIndex"]
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
func (_attrs *Attributes) SetTabIndex(_index int) {
	_attrs.Set("tabIndex", strconv.Itoa(_index))
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
func (_attrs Attributes) Autocapitalize() AUTOCAPITALIZE {
	autocap := _attrs.attributes["autocapitalize"]
	switch autocap {
	case string(AUTOCAP_OFF), string(AUTOCAP_SENTENCES), string(AUTOCAP_WORDS), string(AUTOCAP_CHARS):
		return AUTOCAPITALIZE(autocap)
	}
	return "not valid"
}

// SetAutocapitalize setting attribute 'autocapitalize' with
func (_attrs *Attributes) SetAutocapitalize(_autocap AUTOCAPITALIZE) {
	switch _autocap {
	case AUTOCAP_OFF, AUTOCAP_SENTENCES, AUTOCAP_WORDS, AUTOCAP_CHARS:
		_attrs.Set("autocapitalize", string(_autocap))
	default:
		log.Println("SetAutocapitalize failed: not valid value")
	}
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
func (_attrs Attributes) ContentEditable() CONTENT_EDITABLE {
	editable := _attrs.attributes["contentEditable"]
	switch editable {
	case string(CONTEDIT_FALSE), string(CONTEDIT_TRUE), string(CONTEDIT_INHERIT):
		return CONTENT_EDITABLE(editable)
	}
	return "not valid"
}

// SetContentEditable setting attribute 'contentEditable' with
// type string (idl: DOMString).
func (_attrs *Attributes) SetContentEditable(_editable CONTENT_EDITABLE) {
	switch _editable {
	case CONTEDIT_FALSE, CONTEDIT_TRUE, CONTEDIT_INHERIT:
		_attrs.Set("contentEditable", string(_editable))
	default:
		log.Println("contentEditable fails: not a valid value")
	}
}

/****************************************************************************
* Parsing
*****************************************************************************/

// ParseAttribute returns the name before the = symbol or _str if no symbol.
// Rteurns the value after the = symbol or an empty value if no symbol.
func ParseAttribute(_str string) (_name string, _value string) {
	_name, _value, _ = strings.Cut(_str, "=")

	// TODO: check validity look for helper.HTMLCheckValidity
	_name = strings.Trim(_name, " ")
	_value = strings.Trim(strings.Trim(_value, "'"), "\"")
	return _name, _value
}

// ParseAttribute split _str into attributes separated by spaces
// An attribute can have a value at the right of a = symbol.
// the value can be delimited by quotes and in that case mau contains whitespaces.
// ends when a > symbol is encoutered and is not within a value
func ParseAttributes(_str string) (*Attributes, error) {

	attrs := NewAttributes(nil)

	var strnames string
	left := _str
	for i := 0; len(left) > 0; i++ {
		strnames, left, _ = strings.Cut(left, "=")
		names := strings.Fields(strnames)
		for _, n := range names {
			// TODO: check name validity
			if !helper.IsValidHTMLName(n) {
				return nil, fmt.Errorf("attribute name is not valid : %q", n)
			}
			attrs.attributes[n] = ""
			// fmt.Printf("attribute %v:%q\n", i, n)
		}

		left = strings.Trim(left, " ")
		if len(left) == 0 || len(names) == 0 || left[0] == '>' {
			break
		}

		name := names[len(names)-1]
		// fmt.Printf(" = ")

		delim := left[0]
		istart := 1
		if delim != '"' && delim != '\'' {
			delim = ' '
			istart = 0
		}

		var value string
		value, left, _ = strings.Cut(left[istart:], string(delim))
		// fmt.Printf(" %q\n", value)
		attrs.attributes[name] = value
	}
	// fmt.Println("len=", len(attrs.attributes))
	return attrs, nil
}
