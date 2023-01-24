package browser

import (
	"log"
	"syscall/js"
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

	// Name returns the qualified name of an attribute, that is the name of the attribute, with the namespace prefix, if any, in front of it.
	// For example, if the local name is lang and the namespace prefix is xml, the returned qualified name is xml:lang.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/Attr/name
	name Name

	// value contains the value of the attribute.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/Attr/value
	value string
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
	ret.name = Name(value.Get("name").String())
	ret.value = value.Get("value").String()
	ownerElement := value.Get("ownerElement")
	ret.ownerElement = NewElementFromJS(ownerElement)
	return ret
}

// SetNewAttribute create a new attribut and update the DOM
func SetNewAttribute(name Name, value string, ownerElement *Element) *Attribute {
	ret := &Attribute{}
	ret.name = Name(normalize(string(name)))
	ret.value = normalize(value)
	ret.ownerElement = ownerElement
	ret.ownerElement.JSValue().Call("setAttribute", string(ret.name), ret.value)
	return ret
}

/****************************************************************************
* Attribute's properties
*****************************************************************************/

// String returns normalized formated properties of this attribute
func (_attr *Attribute) String() (_ret string) {
	if _attr == nil {
		return ""
	}
	_ret = normalize(string(_attr.name))
	if _attr.value != "" {
		_ret += `="` + normalize(_attr.value) + `"`
	}
	return _ret
}

// OwnerElement  returns the Element the attribute belongs to.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/ownerElement
func (_attr *Attribute) OwnerElement() *Element {
	if _attr == nil {
		log.Println("OwnerElement() call on a nil Attribute")
		return &Element{}
	}
	return _attr.ownerElement
}

// Name return the name of this Attribute
func (_attr *Attribute) Name() Name {
	if _attr == nil {
		log.Println("Name() call on a nil Attribute")
		return ""
	}
	return _attr.name
}

/****************************************************************************
* Attribute's Methods
*****************************************************************************/

// Reset update the DOM of the ownerElement with this attribute
func (_attr *Attribute) Reset(_attribute string) {
	if _attr == nil {
		log.Println("ToDOM() call on a nil Attribute")
		return
	}
	if !_attr.ownerElement.IsDefined() {
		log.Println("ToDOM() call on an Attribute without ownerElement")
		return
	}

	_attr.value = _attribute
	_attr.ownerElement.JSValue().Call("setAttribute", string(_attr.name), _attr.value)
}
