package dom

import (
	"log"
	"strings"
	"syscall/js"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/lib"
)

/****************************************************************************
* Attribute
*****************************************************************************/

// Attr represents one of an element's attributes as an object.
// In most situations, you will directly retrieve the attribute value as a string (e.g., Element.getAttribute()),
// but certain functions (e.g., Element.getAttributeNode()) or means of iterating return Attr instances.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr
type Attribute struct {
	// the Element the attribute belongs to.
	ownerElement *Element

	// attr.Name returns the qualified name of an attribute, that is the name of the attribute, with the namespace prefix, if any, in front of it.
	// For example, if the local name is lang and the namespace prefix is xml, the returned qualified name is xml:lang.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/Attr/value
	//
	// attr.Value contains the value of the attribute.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/Attr/name
	attr lib.Attribute
}

/****************************************************************************
* Attribute's factory
*****************************************************************************/

// AttrFromJS is casting a js.Value into Attribute
func NewAttributeFromJS(value js.Value) *Attribute {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Attribute{}
	ret.attr.Name = value.Get("name").String()
	ret.attr.Value = value.Get("value").String()
	ret.ownerElement = NewElementFromJS(value.Get("ownerElement"))
	return ret
}

// NewAttributeToDOM create a new attribut and update the DOM with setattribute
func NewAttributeToDOM(_attr lib.Attribute, ownerElement *Element) *Attribute {
	ret := &Attribute{}
	ret.attr.Name = helper.Normalize(_attr.Name)
	ret.attr.Value = strings.Trim(_attr.Value, " ")
	ret.ownerElement = ownerElement
	ret.ownerElement.JSValue().Call("setAttribute", ret.attr.Name, ret.attr.Value)
	return ret
}

/****************************************************************************
* Attribute's Preperties & Methods
*****************************************************************************/

// OwnerElement returns the Element the attribute belongs to.
//
// returns an empty element if _attribute is nil
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/ownerElement
func (_attribute *Attribute) OwnerElement() *Element {
	if _attribute == nil {
		log.Println("OwnerElement() call on a nil Attribute")
		return &Element{}
	}
	return _attribute.ownerElement
}

// String returns normalized formated properties of this attribute
//
//	if value is empty, the format is `{name}`
//	else the format is `{name}="{value}"`
func (_attribute *Attribute) String() string {
	if _attribute == nil {
		return ""
	}
	return _attribute.attr.String()
}

func (_attribute *Attribute) Name() string {
	if _attribute == nil {
		return ""
	}
	return _attribute.attr.Name
}

func (_attribute *Attribute) Value() string {
	if _attribute == nil {
		return ""
	}
	return _attribute.attr.Value
}

func (_attribute *Attribute) Bool() bool {
	if _attribute == nil {
		return false
	}
	if _attribute.Value() == "false" || _attribute.Value() == "0" {
		return false
	}
	return true
}

// Update updates the DOM of the ownerElement with this attribute's value
func (_attribute *Attribute) Update(_value string) {
	if _attribute == nil {
		log.Println("ToDOM() call on a nil Attribute")
		return
	}
	if !_attribute.ownerElement.IsDefined() {
		log.Println("ToDOM() call on an Attribute without ownerElement")
		return
	}

	_attribute.attr.Value = _value
	_attribute.ownerElement.JSValue().Call("setAttribute", string(_attribute.attr.Name), _attribute.attr.Value)
}
