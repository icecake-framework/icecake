package webclientsdk

import (
	"syscall/js"
)

/****************************************************************************
* NamedAttrMap
*
* renaming of NamedNodeMap
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/NamedNodeMap
type NamedAttrMap struct {
	jsValue js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *NamedAttrMap) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.jsValue
}

// NamedNodeMapFromJS is casting a js.Value into NamedNodeMap.
func NamedNodeMapFromJS(value js.Value) *NamedAttrMap {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &NamedAttrMap{}
	ret.jsValue = value
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *NamedAttrMap) Length() uint {
	var ret uint
	value := _this.jsValue.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *NamedAttrMap) Index(index uint) (_result *Attr) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("item", _args[0:_end]...)
	var (
		_converted *Attr // javascript: Attr _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = AttrFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *NamedAttrMap) Get(qualifiedName string) (_result *Attr) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := qualifiedName
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("getNamedItem", _args[0:_end]...)
	var (
		_converted *Attr // javascript: Attr _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = AttrFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *NamedAttrMap) Item(index uint) (_result *Attr) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("item", _args[0:_end]...)
	var (
		_converted *Attr // javascript: Attr _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = AttrFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *NamedAttrMap) GetNamedItem(qualifiedName string) (_result *Attr) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := qualifiedName
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("getNamedItem", _args[0:_end]...)
	var (
		_converted *Attr // javascript: Attr _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = AttrFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *NamedAttrMap) SetNamedItem(attr *Attr) (_result *Attr) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := attr.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("setNamedItem", _args[0:_end]...)
	var (
		_converted *Attr // javascript: Attr _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = AttrFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *NamedAttrMap) SetNamedItemNS(attr *Attr) (_result *Attr) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := attr.JSValue()
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("setNamedItemNS", _args[0:_end]...)
	var (
		_converted *Attr // javascript: Attr _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = AttrFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *NamedAttrMap) RemoveNamedItem(qualifiedName string) (_result *Attr) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := qualifiedName
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("removeNamedItem", _args[0:_end]...)
	return AttrFromJS(_returned)
}
