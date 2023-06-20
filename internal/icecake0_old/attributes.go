package ick0

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/icecake-framework/icecake/pkg/stringpattern"
)

type StringQuotes string

// ParseStringQuotes returns a trimed value keeping white space inside quotes if any
// If _value do not have quotes, the returned value is truncated at the first white space found.
func ParseStringQuotes(_str string) (_unquote StringQuotes) { //, _quote byte, _has bool) {

	trimspaces := strings.Trim(_str, " ")

	trimq1 := strings.Trim(trimspaces, "'")
	if len(trimq1) == len(trimspaces)-2 {
		_unquote = StringQuotes(trimq1)
		return
	}

	trimq2 := strings.Trim(trimspaces, "\"")
	if len(trimq2) == len(trimspaces)-2 {
		_unquote = StringQuotes(trimq2)
		return
	}

	s, _, _ := strings.Cut(trimspaces, " ")
	return StringQuotes(s)
}

func (_vq StringQuotes) StringUnquoted() string {
	return _vq.string(false)
}

func (_vq StringQuotes) StringQuoted() string {
	return _vq.string(true)
}

func (_vq StringQuotes) string(_quoted bool) string {
	if _vq == "" {
		return ""
	}

	boolv := strings.ToLower(string(_vq))
	if boolv == "true" || boolv == "false" {
		return boolv
	}

	num := strings.Trim(string(_vq), " ")
	_, err := strconv.ParseFloat(num, 64)
	if err == nil {
		return num
	}

	var delim string
	if _quoted || strings.ContainsRune(string(_vq), rune(' ')) {
		if strings.ContainsRune(string(_vq), rune('\'')) {
			delim = "\""
		} else {
			delim = "'"
		}
	}
	return delim + string(_vq) + delim
}

/****************************************************************************
* Attributes
*****************************************************************************/

// Attributes represents a set of element's attributes
type Attributes struct {
	amap map[string]StringQuotes // the internal map of attributes
}

// ParseAttribute split _str into attributes separated by spaces.
// An attribute can have a value at the right of an "=" symbol.
// The value can be delimited by quotes ( " or ' ) and in that case may contains whitespaces.
// The string is processed until the end and an error occurs if invalid char is met.
func ParseAttributes(_alist string) (_pattrs *Attributes, _err error) {

	_pattrs = new(Attributes)
	_pattrs.amap = make(map[string]StringQuotes)
	var strnames string
	unparsed := _alist
	for i := 0; len(unparsed) > 0; i++ {

		// process all simple attributes until next "="
		strnames, unparsed, _ = strings.Cut(unparsed, "=")
		names := strings.Fields(strnames)
		for _, n := range names {
			if !stringpattern.IsValidName(n) {
				return nil, fmt.Errorf("attribute name %q is not valid", n)
			}
			_pattrs.amap[n] = ""
		}

		// remove blanks just after "="
		unparsed = strings.TrimLeft(unparsed, " ")

		// stop if nothing else to proceed
		if len(unparsed) == 0 || len(names) == 0 {
			break
		}

		// extract attribute name with a value
		name := names[len(names)-1]

		// extract value with quotes or no quotes
		var value string
		delim := unparsed[0]
		istart := 1
		if delim != '"' && delim != '\'' {
			delim = ' '
			istart = 0
		}
		value, unparsed, _ = strings.Cut(unparsed[istart:], string(delim))
		_pattrs.amap[name] = StringQuotes(value)
	}
	return _pattrs, nil
}

/****************************************************************************
* Properties
*****************************************************************************/

// Keys returns a sorted slice of attributes' key
func (_attrs Attributes) Keys() []string {
	s := make([]string, 0, len(_attrs.amap))
	if _attrs.amap != nil {
		for k := range _attrs.amap {
			s = append(s, k)
		}
		sort.Strings(s)
	}
	return s
}

// String returns the value of the list serialized as a string
func (_attrs Attributes) StringQuoted() (_str string) {
	return _attrs.string(true)
}

func (_attrs Attributes) StringUnquoted() (_str string) {
	return _attrs.string(false)
}

func (_attrs Attributes) string(_quoted bool) (_str string) {
	k := _attrs.Keys()
	for _, name := range k {
		_str += name
		value := _attrs.amap[name]
		if len(value) > 0 {
			_str += "=" + value.string(_quoted)
		}
		_str += " "
	}
	return strings.TrimRight(_str, " ")
}

// Clear removes all attribute
func (_attrs *Attributes) Clear() {
	_attrs.amap = nil
}

// GetAttribue returns the attribute with the given name in the list.
// _name is case sensitive and must be trimed.
func (_attrs Attributes) Attribute(_name string) (_val StringQuotes, _found bool) {
	if _attrs.amap != nil {
		_val, _found = _attrs.amap[_name]
	}
	return _val, _found
}

// SetAttribues adds attributes in the list. If the attribute already exist it's updated.
func (_attrs *Attributes) SetAttributes(_newattrs Attributes, _overwrite bool) *Attributes {
	if _newattrs.amap != nil {
		if _attrs.amap == nil {
			_attrs.amap = make(map[string]StringQuotes)
		}
		for n, v := range _newattrs.amap {
			_, found := _attrs.amap[n]
			if !found || found && _overwrite {
				_attrs.amap[n] = v
			}
		}
	}
	return _attrs
}

// IsTrue returns true if the attribute is set and its value is not false nor 0.
// _name is case sensitive and must be trimed.
func (_attrs Attributes) IsTrue(_name string) bool {
	val, found := _attrs.Attribute(_name)
	v := strings.ToLower(string(val))
	if !found || v == "false" || v == "0" {
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
//  1. Elements with a positive tabindex. Elements that have identical tabindex values should be navigated in the order they appear. Navigation proceeds from the lowest tabindex to the highest tabindex.
//  1. Elements that do not support the tabindex attribute or support it and assign tabindex to 0, in the order they appear.
//  1. Elements that are disabled do not participate in the tabbing order.
//
// Values don't need to be sequential, nor must they begin with any particular value.
// They may even be negative, though each browser trims very large values.
//
// https://developer.mozilla.org/fr/docs/Web/HTML/Global_attributes/tabindex
func (_attrs Attributes) TabIndex() int {
	stri, _ := _attrs.Attribute("tabindex")
	i, _ := strconv.Atoi(string(stri))
	return i
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
func (_attrs *Attributes) SetTabIndex(_index int) *Attributes {
	_attrs.SetAttribute("tabindex", StringQuotes(strconv.Itoa(_index)))
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
func (_attrs Attributes) Autocapitalize() AUTOCAPITALIZE {
	autocap, _ := _attrs.Attribute("autocapitalize")
	switch autocap {
	case StringQuotes(AUTOCAP_OFF), StringQuotes(AUTOCAP_SENTENCES), StringQuotes(AUTOCAP_WORDS), StringQuotes(AUTOCAP_CHARS):
		return AUTOCAPITALIZE(autocap)
	}
	return "not valid"
}

// SetAutocapitalize setting attribute 'autocapitalize' with
func (_attrs *Attributes) SetAutocapitalize(_autocap AUTOCAPITALIZE) *Attributes {
	switch _autocap {
	case AUTOCAP_OFF, AUTOCAP_SENTENCES, AUTOCAP_WORDS, AUTOCAP_CHARS:
		_attrs.SetAttribute("autocapitalize", StringQuotes(_autocap))
	default:
		log.Printf("SetAutocapitalize failed: not a valid value %q\n", _autocap)
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
func (_attrs Attributes) ContentEditable() CONTENT_EDITABLE {
	editable, _ := _attrs.Attribute("contentEditable")
	switch editable {
	case StringQuotes(CONTEDIT_FALSE), StringQuotes(CONTEDIT_TRUE), StringQuotes(CONTEDIT_INHERIT):
		return CONTENT_EDITABLE(editable)
	}
	return "not valid"
}

// SetContentEditable setting attribute 'contentEditable' with
// type string (idl: DOMString).
func (_attrs *Attributes) SetContentEditable(_editable CONTENT_EDITABLE) *Attributes {
	switch _editable {
	case CONTEDIT_FALSE, CONTEDIT_TRUE, CONTEDIT_INHERIT:
		_attrs.SetAttribute("contentEditable", StringQuotes(_editable))
	default:
		log.Printf("contentEditable fails: not a valid value %q\n", _editable)
	}
	return _attrs
}

// SetAttribue setup an attribute in the map. If the attribute already exist it's updated.
// An empty value means a boolean attribute set to true.
//
// Name and Value are case sensitive, they will be trimed. Quotes delimiters of the value will be removed if any.
func (_attrs *Attributes) SetAttribute(_name string, _sq StringQuotes) *Attributes {
	if _attrs.amap == nil {
		_attrs.amap = make(map[string]StringQuotes)
	}
	_name = strings.Trim(_name, " ")
	_attrs.amap[_name] = _sq
	return _attrs
}

// ParseAttribute parse the _value and set it up in the map. If the attribute already exist it's updated.
// An empty value means a boolean attribute set to true.
//
// Name and Value are case sensitive, they will be trimed. Quotes delimiters of the value will be removed if any.
func (_attrs *Attributes) ParseAttribute(_name string, _value string) *Attributes {
	return _attrs.SetAttribute(_name, ParseStringQuotes(_value))
}

// RemoveAttribute removes attribute in the list or does nothing for the one that does not exist.
// _name is case sensitive and must be trimed.
func (_attrs *Attributes) RemoveAttribute(_name string) *Attributes {
	if _attrs.amap != nil {
		delete(_attrs.amap, _name)
	}
	return _attrs
}

// Toggle toggles a boolean attribute (removing it if it's present and adding it if it's not present).
// _name is case sensitive and must be trimed.
// returns true is the token is in the list after the call.
func (_attrs *Attributes) Toggle(_name string) (_isin bool) {
	_, found := _attrs.amap[_name]
	if found {
		delete(_attrs.amap, _name)
	} else {
		_attrs.amap[_name] = ""
		_isin = true
	}
	return _isin
}

// returns an subset of _attrs with only "data-*" attributes
func (_attrs Attributes) Data() *Attributes {
	dataset := new(Attributes)
	dataset.amap = make(map[string]StringQuotes)
	if _attrs.amap != nil {
		for name, value := range _attrs.amap {
			if len(name) > 5 && strings.HasPrefix(name, "data-") {
				dataset.amap[name] = value
			}
		}
	}
	return dataset
}
