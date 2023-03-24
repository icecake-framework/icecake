package wick

import (
	"strconv"
	"strings"

	"github.com/sunraylab/icecake/pkg/errors"
)

/****************************************************************************
* Attributes
*****************************************************************************/

// JSAttributes represents a set of element's attributes. The subset is static.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
type JSAttributes struct {
	owner *Element // the Element the attributes belongs to. The DOM is only updated when ownerElement is not nil.
	// cache map[string]string // the internal map of attributes only used when the attributes object does not belong to an element yet.
}

// func NewAttributes(_ownerElement *Element) *Attributes {
// 	attrs := new(Attributes)
// 	attrs.ownerElement = _ownerElement
// 	attrs.attributes = make(map[string]string, 0)
// 	return attrs
// }

/****************************************************************************
* Attributes's properties
*****************************************************************************/

// func (_attrs Attributes) OwnerElement() *Element {
// 	return _attrs.owner
// }

// String returns the value of the list serialized as a string
//
// https://developer.mozilla.org/fr/docs/Web/API/Element/attributes
func (_attrs JSAttributes) String() (_str string) {
	if _attrs.owner != nil {
		jsa := _attrs.owner.Get("attributes")
		len := jsa.GetInt("length")
		for i := 0; i < len; i++ {
			jsi := jsa.Call("item", i)
			_str += jsi.GetString("name")
			value := jsi.GetString("value")
			if value != "" {
				delim := "'"
				if strings.ContainsRune(value, rune('\'')) {
					delim = "\""
				}
				_str += `=` + delim + value + delim
			}
			_str += " "
		}
	}
	return strings.TrimRight(_str, " ")
}

// GetAttribue returns the attribute with the given name in the list.
// _name is case sensitive and must be trimed.
func (_attrs JSAttributes) Attribute(_name string) (_val string, _found bool) {
	if _attrs.owner != nil {
		_found = _attrs.owner.Call("hasAttribute", _name).Bool()
		if _found {
			_val = _attrs.owner.Call("getAttribute", _name).String()
		}
	}
	// else if _attrs.cache != nil {
	// 	_val, _found = _attrs.cache[_name]
	// }
	return _val, _found
}

// SetAttribues adds attributes in the list. If the attribute already exist it's updated.
// func (_attrs *JSAttributes) SetAttributes(_newattrs ick.Attributes) *JSAttributes {
// 	if _attrs.owner != nil {
// 		for _, k := range _newattrs.Keys() {
// 			v, _ := _newattrs.Attribute(k)
// 			_attrs.owner.Call("setAttribute", k, v)
// 		}
// 	}
// 	return _attrs
// }

// // ParseAttribute split _str into attributes separated by spaces
// // An attribute can have a value at the right of an "=" symbol.
// // the value can be delimited by quotes and in that case may contains whitespaces.
// // The string is processed until the ends or when a ">" symbol is encoutered out of a value.
// func (_attrs *JSAttributes) ParseAttributes(_str string) (_err error) {
// 	pattrs := make(map[string]string)
// 	var strnames string
// 	left := _str
// 	for i := 0; len(left) > 0; i++ {
// 		// process all simple attributes until next "="
// 		strnames, left, _ = strings.Cut(left, "=")
// 		names := strings.Fields(strnames)
// 		for _, n := range names {
// 			if !htmlname.IsValid(n) {
// 				return errors.ConsoleErrorf("attribute name %q is not valid\n", n)
// 			}
// 			pattrs[n] = ""
// 		}

// 		// remove blanks just after "="
// 		left = strings.Trim(left, " ")

// 		// stop if nothing else to proceed
// 		if len(left) == 0 || len(names) == 0 || left[0] == '>' {
// 			break
// 		}

// 		// extract attribute name with a value
// 		name := names[len(names)-1]

// 		// extract value with quotes or no quotes
// 		var value string
// 		delim := left[0]
// 		istart := 1
// 		if delim != '"' && delim != '\'' {
// 			delim = ' '
// 			istart = 0
// 		}
// 		value, left, _ = strings.Cut(left[istart:], string(delim))
// 		pattrs[name] = value
// 	}
// 	// now set it up
// 	if _attrs.cache != nil && len(pattrs) > 0 {
// 		_attrs.cache = make(map[string]string)
// 	}
// 	for name, value := range pattrs {
// 		_attrs.SetAttribute(name, value)
// 	}
// 	return nil
// }

// IsTrue returns true if the attribute is set and its value is not false nor 0.
// _name is case sensitive and must be trimed.
func (_attrs JSAttributes) IsTrue(_name string) bool {
	val, found := _attrs.Attribute(_name)
	val = strings.ToLower(val)
	if !found || val == "false" || val == "0" {
		return false
	}
	return true
}

// Hidden returns boolean attribute 'hidden'
func (_attrs JSAttributes) Hidden() bool {
	return _attrs.IsTrue("hidden")
}

// Draggable returns boolean attribute 'draggable'
func (_attrs JSAttributes) Draggable() bool {
	return _attrs.IsTrue("draggable")
}

// Spellcheck returns boolean attribute 'spellcheck'
func (_attrs JSAttributes) Spellcheck() bool {
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
func (_attrs JSAttributes) TabIndex() int {
	stri, _ := _attrs.Attribute("tabIndex")
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
func (_attrs *JSAttributes) SetTabIndex(_index int) *JSAttributes {
	_attrs.SetAttribute("tabIndex", strconv.Itoa(_index))
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
func (_attrs JSAttributes) Autocapitalize() AUTOCAPITALIZE {
	autocap, _ := _attrs.Attribute("autocapitalize")
	switch autocap {
	case string(AUTOCAP_OFF), string(AUTOCAP_SENTENCES), string(AUTOCAP_WORDS), string(AUTOCAP_CHARS):
		return AUTOCAPITALIZE(autocap)
	}
	return "not valid"
}

// SetAutocapitalize setting attribute 'autocapitalize' with
func (_attrs *JSAttributes) SetAutocapitalize(_autocap AUTOCAPITALIZE) *JSAttributes {
	switch _autocap {
	case AUTOCAP_OFF, AUTOCAP_SENTENCES, AUTOCAP_WORDS, AUTOCAP_CHARS:
		_attrs.SetAttribute("autocapitalize", string(_autocap))
	default:
		errors.ConsoleWarnf("SetAutocapitalize failed: not valid value %q\n", _autocap)
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
func (_attrs JSAttributes) ContentEditable() CONTENT_EDITABLE {
	editable, _ := _attrs.Attribute("contentEditable")
	switch editable {
	case string(CONTEDIT_FALSE), string(CONTEDIT_TRUE), string(CONTEDIT_INHERIT):
		return CONTENT_EDITABLE(editable)
	}
	return "not valid"
}

// SetContentEditable setting attribute 'contentEditable' with
// type string (idl: DOMString).
func (_attrs *JSAttributes) SetContentEditable(_editable CONTENT_EDITABLE) *JSAttributes {
	switch _editable {
	case CONTEDIT_FALSE, CONTEDIT_TRUE, CONTEDIT_INHERIT:
		_attrs.SetAttribute("contentEditable", string(_editable))
	default:
		errors.ConsoleWarnf("contentEditable fails: not a valid value %q\n", _editable)
	}
	return _attrs
}

// SetAttribue setup an attribute in the map. If the attribute already exist it's updated.
// An empty value means a boolean attribute set to true.
//
// Name and Value are case sensitive, they will be trimed. Quotes delimiters of the value will be removed if any.
func (_attrs *JSAttributes) SetAttribute(_name string, _value string) *JSAttributes {
	_name = strings.Trim(_name, " ")
	_value = strings.Trim(strings.Trim(strings.Trim(_value, " "), "\""), "'")
	if _attrs.owner != nil {
		_attrs.owner.Call("setAttribute", _name, _value)
	}
	// else {
	// 	if _attrs.cache != nil {
	// 		_attrs.cache = make(map[string]string)
	// 		_attrs.setCache(_name, _value)
	// 	}
	// }
	return _attrs
}

// RemoveAttribute removes attribute in the list or does nothing for the one that does not exist.
// _name is case sensitive and must be trimed.
func (_attrs *JSAttributes) RemoveAttribute(_name string) *JSAttributes {
	if _attrs.owner != nil {
		_attrs.owner.Call("removeAttribute", _name)
	}
	// else if _attrs.cache != nil {
	// 	delete(_attrs.cache, _name)
	// }
	return _attrs
}

// Toggle toggles a boolean attribute (removing it if it's present and adding it if it's not present).
// _name is case sensitive and must be trimed.
// returns true is the token is in the list after the call.
func (_attrs *JSAttributes) Toggle(_name string) (_isin bool) {
	if _attrs.owner != nil {
		return _attrs.owner.Call("toggleAttribute", _name).Bool()
	}
	// else {
	// 	_, found := _attrs.cache[_name]
	// 	if found {
	// 		delete(_attrs.cache, _name)
	// 	} else {
	// 		_attrs.setCache(_name, "")
	// 		_isin = true
	// 	}
	// }
	return _isin
}

// ParseAttribute returns the name before the = symbol or _str if no symbol.
// Rteurns the value after the = symbol or an empty value if no symbol.
// func ParseAttribute(_str string) (_name string, _value string) {
// 	_name, _value, _ = strings.Cut(_str, "=")

// 	// TODO: check validity look for helper.HTMLCheckValidity
// 	_name = strings.Trim(_name, " ")
// 	_value = strings.Trim(strings.Trim(_value, "'"), "\"")
// 	return _name, _value
// }

// func (_attrs JSAttributes) sortCache() []string {
// 	s := make([]string, 0, len(_attrs.cache))
// 	if _attrs.cache != nil {
// 		for k := range _attrs.cache {
// 			s = append(s, k)
// 		}
// 		sort.Strings(s)
// 	}
// 	return s
// }

// // returns an subset of _attrs with only "data-*" attributes
// func (_attrs JSAttributes) Dataset() map[string]string {
// 	dataset := make(map[string]string)
// 	if _attrs.cache != nil {
// 		for name, value := range _attrs.cache {
// 			if len(name) > 5 && strings.HasPrefix(name, "data-") {
// 				dataset[name] = value
// 			}
// 		}
// 	}
// 	return dataset
// }

/*****************************************************************************
* PRIVATE
*****************************************************************************/

// func (_attrs *JSAttributes) setCache(_name string, _value string) {
// 	_attrs.cache[_name] = _value
// 	if _attrs.owner != nil {
// 		_attrs.owner.Call("setAttribute", _name, _value)
// 	}
// }
