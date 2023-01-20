package webclientsdk

import (
	"log"
	"strings"
	"syscall/js"
)

// Attr represents one of an element's attributes as an object.
// In most situations, you will directly retrieve the attribute value as a string (e.g., Element.getAttribute()),
// but certain functions (e.g., Element.getAttributeNode()) or means of iterating return Attr instances.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr
type Attr struct {

	// Name returns the qualified name of an attribute, that is the name of the attribute, with the namespace prefix, if any, in front of it.
	// For example, if the local name is lang and the namespace prefix is xml, the returned qualified name is xml:lang.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/Attr/name
	Name string

	// Value contains the value of the attribute.
	//
	// https://developer.mozilla.org/en-US/docs/Web/API/Attr/value
	Value string

	// the Element the attribute belongs to.
	ownerElement *Element
}

// AttrFromJS is casting a js.Value into Attr.
func AttrFromJS(value js.Value) *Attr {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Attr{}
	//ret.jsValue = value
	name := value.Get("name")
	ret.Name = (name).String()
	val := value.Get("value")
	ret.Value = (val).String()
	ownerElement := value.Get("ownerElement")
	ret.ownerElement = ElementFromJS(ownerElement)

	return ret
}

// JSValue returns the js.Value or js.Null() if _attr is nil
func (_attr *Attr) JSValue() js.Value {
	if _attr == nil {
		return js.Null()
	}
	jsValue := js.Value{}
	jsValue.Set("name", _attr.Name)
	jsValue.Set("value", _attr.Value)
	jsValue.Set("ownerElement", _attr.ownerElement)
	return jsValue
}

// Prefix returns the namespace prefix of the attribute, or an empty string if no prefix is specified.
// For example, if the qualified name is xml:lang, the returned prefix is xml.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/prefix
func (_attr *Attr) Prefix() string {
	if _attr == nil {
		log.Println("Prefix() call on a nil Attr")
		return ""
	}

	if strings.Contains(_attr.Name, ":") {
		s := strings.Split(_attr.Name, ":")
		return s[0]
	} else {
		return ""
	}
}

// LocalName returns the local part of the qualified name of an attribute, that is the name of the attribute,
// stripped from any namespace in front of it.
// For example, if the qualified name is xml:lang, the returned local name is lang.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/localName
func (_attr *Attr) LocalName() string {
	if _attr == nil {
		log.Println("LocalName() call on a nil Attr")
		return ""
	}

	if strings.Contains(_attr.Name, ":") {
		s := strings.Split(_attr.Name, ":")
		return s[1]
	} else {
		return _attr.Name
	}
}

// OwnerElement  returns the Element the attribute belongs to.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/ownerElement
func (_attr *Attr) OwnerElement() *Element {
	if _attr == nil {
		log.Println("OwnerElement() call on a nil Attr")
		return nil
	}
	return _attr.ownerElement
}
