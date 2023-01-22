package browser

import (
	"strings"
	"syscall/js"
)

/****************************************************************************
* Attributes
*****************************************************************************/

// Attributes represents a set of element's attributes. The subset is static
//
// # Need to call ToDOM() to update the DOM with internal value avec any change
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attributes
type Attributes struct {
	jsValue    js.Value
	attributes []*Attribute

	OwnerElement *Element // the Element the attributes belongs to.
}

// JSValue returns the js.Value or js.Null() if _this is nil
// Calling JSValue updates the DOM.
// func (_this *Attributes) JSValue() js.Value {
// 	if _this == nil {
// 		return js.Null()
// 	}
// 	_this.ToDOM()
// 	return _this.jsValue
// }

/****************************************************************************
* Attributes's factory
*****************************************************************************/

// MakeAttributesFromNamedNodeMapJS is casting a js.Value into DOMAttributes.
func MakeAttributesFromNamedNodeMapJS(_value js.Value, _ownerElement *Element) (_ret *Attributes) {
	if typ := _value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	_ret = &Attributes{}
	_ret.jsValue = _value

	_ret.attributes = make([]*Attribute, 0)
	_ret.OwnerElement = _ownerElement

	data := _value.Get("length")
	len := (uint)((data).Int())
	for i := uint(0); i < len; i++ {
		__returned := _value.Call("item", i)
		_result := MakeAttributeFromJS(__returned)
		_ret.attributes = append(_ret.attributes, _result)
	}
	return _ret
}

/****************************************************************************
* Attributes's properties
*****************************************************************************/

// Length returns the number of attributes in the list.
func (_this *Attributes) Length() int {
	return len(_this.attributes)
}

// Item returns an item in the list, determined by its position in the list, its index.
// Returns nil if out of range.
func (_this *Attributes) Item(index int) *Attribute {
	if index >= 0 && index < len(_this.attributes) {
		return _this.attributes[index]
	}
	return nil
}

// String returns the value of the list serialized as a string
func (_this *Attributes) String() (_ret string) {
	for _, v := range _this.attributes {
		_ret += v.String() + " "
	}
	_ret = strings.TrimRight(_ret, " ")
	return _ret
}

// GetAttribue returns the attribute with the given name in the list.
// returns an empty Attribute
func (_this *Attributes) Get(name Name) *Attribute {
	name = Name(normalize(string(name)))
	for _, a := range _this.attributes {
		if a.name == name {
			return a
		}
	}
	return nil
}

// SetAttribue adds an attribute in the list. If the attribute already exist it's only updated
func (_this *Attributes) Set(name Name, value string) *Attribute {
	name = Name(normalize(string(name)))
	value = normalize(value)
	a := _this.Get(name)
	if a == nil {
		a := CreateAttribute(name, value, _this.OwnerElement)
		_this.attributes = append(_this.attributes, a)
	} else {
		a.SetValue(value)
	}
	return a

}

// Remove removes attributes in the list or does nothing for the one that does not exist.
func (_this *Attributes) Remove(name Name) (_ret bool) {
	_ret = false
	name = Name(normalize(string(name)))
	for i, a := range _this.attributes {
		if a.name == name {
			_this.attributes = append(_this.attributes[:i], _this.attributes[i+1:]...)
			_this.OwnerElement.jsValue.Call("removeAttribute", string(name))
			_ret = true
		}
	}
	return _ret
}

// Toggle a attribute from one to another. If none exists then one is set.
//
// returns the set attribute.
func (_this *Attributes) Toggle(one Name, another Name) (_ret *Attribute) {
	one = Name(normalize(string(one)))
	another = Name(normalize(string(another)))

	onewas := _this.Remove(one)
	if onewas {
		return _this.Set(another, "")
	} else {
		_this.Remove(another)
		return _this.Set(one, "")
	}
}
