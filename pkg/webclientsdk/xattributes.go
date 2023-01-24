package browser

import (
	"log"
	"strings"
	"syscall/js"
)

/****************************************************************************
* ElemAttributes
*****************************************************************************/

// ElemAttributes represents a set of element's attributes. The subset is static
//
// https://developer.mozilla.org/en-US/docs/Web/API/ElemAttributes
type ElemAttributes struct {
	jsValue    js.Value
	attributes []*Attribute

	ownerElement *Element // the Element the attributes belongs to.
}

// NewElemAttributesFromJSNamedNodeMap is casting a js.Value into DOMAttributes.
func NewElemAttributesFromJSNamedNodeMap(_value js.Value, _ownerElement *Element) (_ret *ElemAttributes) {
	if typ := _value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	_ret = new(ElemAttributes)
	_ret.jsValue = _value

	_ret.attributes = make([]*Attribute, 0)
	_ret.ownerElement = _ownerElement

	data := _value.Get("length")
	len := (uint)((data).Int())
	for i := uint(0); i < len; i++ {
		__returned := _value.Call("item", i)
		_result := NewAttributeFromJS(__returned)
		_ret.attributes = append(_ret.attributes, _result)
	}
	return _ret
}

/****************************************************************************
* Attributes's properties
*****************************************************************************/

// Length returns the number of attributes in the list.
func (_elemattrs *ElemAttributes) OwnerElement() *Element {
	return _elemattrs.ownerElement
}

// Length returns the number of attributes in the list.
func (_elemattrs *ElemAttributes) Count() int {
	return len(_elemattrs.attributes)
}

// Item returns an item in the list, determined by its position in the list, its index.
// Returns nil if out of range.
func (_elemattrs *ElemAttributes) Index(index int) *Attribute {
	if index >= 0 && index < len(_elemattrs.attributes) {
		return _elemattrs.attributes[index]
	}
	return nil
}

// String returns the value of the list serialized as a string
func (_elemattrs *ElemAttributes) String() (_ret string) {
	for _, v := range _elemattrs.attributes {
		_ret += v.String() + " "
	}
	_ret = strings.TrimRight(_ret, " ")
	return _ret
}

// GetAttribue returns the attribute with the given name in the list.
// returns an empty Attribute
func (_elemattrs *ElemAttributes) Get(name Name) *Attribute {
	name = Name(normalize(string(name)))
	for _, a := range _elemattrs.attributes {
		if a.name == name {
			return a
		}
	}
	return nil
}

// SetAttribue adds an attribute in the list. If the attribute already exist it's only updated
func (_elemattrs *ElemAttributes) Set(name Name, value string) *Attribute {
	name = Name(normalize(string(name)))
	value = normalize(value)
	a := _elemattrs.Get(name)
	if a == nil {
		a := SetNewAttribute(name, value, _elemattrs.ownerElement)
		_elemattrs.attributes = append(_elemattrs.attributes, a)
	} else {
		a.Reset(value)
	}
	return a
}

// Remove removes attributes in the list or does nothing for the one that does not exist.
func (_elemattrs *ElemAttributes) Remove(name Name) (_ret bool) {
	_ret = false
	name = Name(normalize(string(name)))
	for i, a := range _elemattrs.attributes {
		if a.name == name {
			_elemattrs.attributes = append(_elemattrs.attributes[:i], _elemattrs.attributes[i+1:]...)
			_elemattrs.ownerElement.jsValue.Call("removeAttribute", string(name))
			_ret = true
		}
	}
	return _ret
}

// Toggle a attribute from one to another. If none exists then one is set.
//
// returns the set attribute.
func (_elemattrs *ElemAttributes) Toggle(one Name, another Name) (_ret *Attribute) {
	one = Name(normalize(string(one)))
	another = Name(normalize(string(another)))

	onewas := _elemattrs.Remove(one)
	if onewas {
		return _elemattrs.Set(another, "")
	} else {
		_elemattrs.Remove(another)
		return _elemattrs.Set(one, "")
	}
}

/****************************************************************************
* HTMLAttributes
*****************************************************************************/

// HTMLAttributes represents a set of HTMLelement's attributes. The subset is static.
type HTMLAttributes struct {
	*ElemAttributes
}

// NewAttributesFromJSNamedNodeMap is casting a js.Value into DOMAttributes.
func NewHTMLAttributesFromJSNamedNodeMap(_value js.Value, _ownerElement *Element) (_ret *HTMLAttributes) {
	_ret = new(HTMLAttributes)
	_ret.ElemAttributes = NewElemAttributesFromJSNamedNodeMap(_value, _ownerElement)
	return _ret
}

// Hidden returning attribute 'hidden' with
// type bool (idl: boolean).
func (_this *HTMLAttributes) Hidden() bool {
	return _this.jsValue.Get("hidden").Bool()
}

// SetHidden setting attribute 'hidden' with
// type bool (idl: boolean).
func (_this *HTMLAttributes) SetHidden(value bool) *HTMLAttributes {
	if value {
		_this.Set("hidden", "")
	} else {
		_this.Remove("hidden")
	}
	return _this
}

// Draggable A boolean value indicating if the element can be dragged.
func (_this *HTMLAttributes) Draggable() bool {
	return _this.jsValue.Get("draggable").Bool()
}

// Draggable A boolean value indicating if the element can be dragged.
func (_this *HTMLAttributes) SetDraggable(value bool) *HTMLAttributes {
	if value {
		_this.Set("draggable", "")
	} else {
		_this.Remove("draggable")
	}
	return _this
}

// Spellcheck returning attribute 'spellcheck' with
// type bool (idl: boolean).
func (_this *HTMLAttributes) Spellcheck() bool {
	return _this.jsValue.Get("spellcheck").Bool()
}

// SetSpellcheck setting attribute 'spellcheck' with
// type bool (idl: boolean).
func (_this *HTMLAttributes) SetSpellcheck(value bool) *HTMLAttributes {
	if value {
		_this.Set("spellcheck", "true")
	} else {
		_this.Set("spellcheck", "false")
	}
	return _this
}

// Controls whether and how text input is automatically capitalized as it is entered/edited by the user.
type Autocapitalize string

const (
	AUTOCAP_OFF       Autocapitalize = "off"
	AUTOCAP_SENTENCES Autocapitalize = "sentences"
	AUTOCAP_WORDS     Autocapitalize = "words"
	AUTOCAP_CHARS     Autocapitalize = "characters"
)

// Autocapitalize returning attribute 'autocapitalize' with
// type string (idl: DOMString).
func (_this *HTMLAttributes) Autocapitalize() Autocapitalize {
	value := _this.jsValue.Get("autocapitalize").String()
	value = normalize(value)
	switch value {
	case string(AUTOCAP_OFF), string(AUTOCAP_SENTENCES), string(AUTOCAP_WORDS), string(AUTOCAP_CHARS):
		return Autocapitalize(value)
	}
	return "not valid"
}

// SetAutocapitalize setting attribute 'autocapitalize' with
// type string (idl: DOMString).
func (_this *HTMLAttributes) SetAutocapitalize(value Autocapitalize) *HTMLAttributes {
	switch value {
	case AUTOCAP_OFF, AUTOCAP_SENTENCES, AUTOCAP_WORDS, AUTOCAP_CHARS:
		_this.Set("autocapitalize", string(value))
	default:
		log.Println("SetAutocapitalize fails: not valid value")
	}
	return _this
}

// Controls whether and how text input is automatically capitalized as it is entered/edited by the user.
type ContentEditable string

const (
	CONTEDIT_FALSE   ContentEditable = "false"
	CONTEDIT_TRUE    ContentEditable = "true"
	CONTEDIT_INHERIT ContentEditable = "inherit"
)

// ContentEditable returns a boolean value that is true if the contents of the element are editable; otherwise it returns false.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/isContentEditable
func (_this *HTMLAttributes) ContentEditable() ContentEditable {
	value := _this.jsValue.Get("contentEditable").String()
	value = normalize(value)
	switch value {
	case string(CONTEDIT_FALSE), string(CONTEDIT_TRUE), string(CONTEDIT_INHERIT):
		return ContentEditable(value)
	}
	return "not valid"
}

// SetContentEditable setting attribute 'contentEditable' with
// type string (idl: DOMString).
func (_this *HTMLAttributes) SetContentEditable(value ContentEditable) *HTMLAttributes {
	switch value {
	case CONTEDIT_FALSE, CONTEDIT_TRUE, CONTEDIT_INHERIT:
		_this.Set("contentEditable", string(value))
	default:
		log.Println("contentEditable fails: not valid value")
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
	return _this.jsValue.Get("tabIndex").Int()
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
func (_this *HTMLAttributes) SetTabIndex(value int) *HTMLAttributes {
	_this.jsValue.Set("tabIndex", value)
	return _this
}

// Title represents the title of the element: the text usually displayed in a 'tooltip' popup when the mouse is over the node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/title
func (_this *HTMLAttributes) Title() string {
	return _this.jsValue.Get("title").String()
}

// Title represents the title of the element: the text usually displayed in a 'tooltip' popup when the mouse is over the node.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/title
func (_this *HTMLAttributes) SetTitle(value string) *HTMLAttributes {
	_this.jsValue.Set("title", value)
	return _this
}

// Lang  gets or sets the base language of an element's attribute values and text content.
//
// The language code returned by this property is defined in RFC 5646: Tags for Identifying Languages (also known as BCP 47).
// Common examples include "en" for English, "ja" for Japanese, "es" for Spanish and so on. The default value of this attribute is unknown. Note that this attribute, though valid at the individual element level described here, is most often specified for the root element of the document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/lang
func (_this *HTMLAttributes) Lang() string {
	return _this.jsValue.Get("lang").String()
}

// Lang  gets or sets the base language of an element's attribute values and text content.
//
// The language code returned by this property is defined in RFC 5646: Tags for Identifying Languages (also known as BCP 47).
// Common examples include "en" for English, "ja" for Japanese, "es" for Spanish and so on. The default value of this attribute is unknown. Note that this attribute, though valid at the individual element level described here, is most often specified for the root element of the document.
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement/lang
func (_this *HTMLAttributes) SetLang(value string) *HTMLAttributes {
	_this.jsValue.Set("lang", value)
	return _this
}
