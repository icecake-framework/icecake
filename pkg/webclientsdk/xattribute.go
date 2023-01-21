package webclientsdk

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

	// Value contains the value of the attribute.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/Attr/value
	Value string
}

/****************************************************************************
* Attribute's factory
*****************************************************************************/

// AttrFromJS is casting a js.Value into Attribute
func MakeAttributeFromJS(value js.Value) *Attribute {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Attribute{}
	ret.name = Name(value.Get("name").String())
	ret.Value = value.Get("value").String()
	ownerElement := value.Get("ownerElement")
	ret.ownerElement = MakeElementFromJS(ownerElement)
	return ret
}

func NewAttribute(name Name, value string, ownerElement *Element) *Attribute {
	ret := &Attribute{}
	ret.name = Name(normalize(string(name)))
	ret.Value = normalize(value)
	ret.ownerElement = ownerElement
	return ret
}

/****************************************************************************
* Attribute's properties
*****************************************************************************/

// String returns normalized formated properties of this attribute
func (_attr *Attribute) String() string {
	if _attr == nil {
		return ""
	}
	return normalize(string(_attr.name)) + `="` + normalize(_attr.Value) + `"`
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

// JSValue returns the js.Value or js.Null() if _attr is nil
// func (_attr *Attribute) JSValue() js.Value {
// 	if _attr == nil {
// 		return js.Null()
// 	}
// 	jsValue := js.Value{}
// 	jsValue.Set("name", _attr.Name)
// 	jsValue.Set("value", _attr.Value)
// 	jsValue.Set("ownerElement", _attr.ownerElement)
// 	return jsValue
// }

// ToDOM update the DOM of the ownerElement with this attribute
func (_attr *Attribute) ToDOM() {
	if _attr == nil {
		log.Println("ToDOM() call on a nil Attribute")
		return
	}
	if !_attr.ownerElement.IsDefined() {
		log.Println("ToDOM() call on an Attribute without ownerElement")
		return
	}

	var _args [2]interface{}
	_args[0] = string(_attr.name)
	_args[1] = _attr.Value
	_attr.ownerElement.JSValue().Call("setAttribute", _args[0:2]...)
}
