package webclientsdk

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
func MakeAttributesFromNamedNodeMapJS(_value js.Value) *Attributes {
	if typ := _value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Attributes{}
	ret.jsValue = _value
	ret.attributes = make([]*Attribute, 0)

	var _args [1]interface{}
	data := _value.Get("length")
	len := (uint)((data).Int())
	for i := uint(0); i < len; i++ {
		_args[0] = i
		_returned := _value.Call("item", _args[0:1]...)
		_result := MakeAttributeFromJS(_returned)
		ret.attributes = append(ret.attributes, _result)
	}
	return ret
}

/****************************************************************************
* Attributes's properties
*****************************************************************************/

// ToDOM update the DOM with the internal value and returns the DOM's js.Value.
func (_this *Attributes) ToDOM() js.Value {
	// TODO
	_this.jsValue.Set("value", _this.String())

	return _this.jsValue
}

// Length returns the number of attributes in the list.
func (_this *Attributes) Length() int {
	return len(_this.attributes)
}

// Item returns an item in the list, determined by its position in the list, its index.
// Returns an empty Attribute if the index is out of range.
func (_this *Attributes) Item(index int) *Attribute {
	if index >= 0 && index < len(_this.attributes) {
		return _this.attributes[index]
	}
	return &Attribute{}
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
		a = NewAttribute(name, value, MakeElementFromJS(_this.jsValue))
	} else {
		a.Value = value
	}
	return a
}

// Remove removes attributes in the list or does nothing for the one that does not exist.
// Duplicates are removed too
func (_this *Attributes) Remove(name Name) {
	name = Name(normalize(string(name)))
	for i, a := range _this.attributes {
		if a.name == name {
			_this.attributes = append(_this.attributes[:i], _this.attributes[i+1:]...)
			var _args [1]interface{}
			_args[0] = string(name)
			_this.jsValue.Call("removeNamedItem", _args[0:1]...)
			return
		}
	}
}
