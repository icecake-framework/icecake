package webclientsdk

import "syscall/js"

/****************************************************************************
* DOMTokenList
*****************************************************************************/

// DOMTokenList represents a set of space-separated tokens.
// Such a set is returned by Element.classList or HTMLLinkElement.relList, and many others.
//
// https://developer.mozilla.org/en-US/docs/Web/API/DOMTokenList
type DOMTokenList struct {
	//jsValue js.Value
	tokens []string
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *DOMTokenList) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.jsValue
}

// DOMTokenListFromJS is casting a js.Value into DOMTokenList.
func DOMTokenListFromJS(value js.Value) *DOMTokenList {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &DOMTokenList{}
	ret.jsValue = value
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *DOMTokenList) Length() uint {
	var ret uint
	value := _this.jsValue.Get("length")
	ret = (uint)((value).Int())
	return ret
}

// ToString is an alias for Value.
func (_this *DOMTokenList) ToString() string {
	return _this.Value()
}

// SetValue setting attribute 'value' with
// type string (idl: DOMString).
func (_this *DOMTokenList) SetValue(value string) {
	input := value
	_this.jsValue.Set("value", input)
}

func (_this *DOMTokenList) Index(index uint) (_result *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("item", _args[0:_end]...)
	var (
		_converted *string // javascript: DOMString _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		__tmp := (_returned).String()
		_converted = &__tmp
	}
	_result = _converted
	return
}

func (_this *DOMTokenList) Item(index uint) (_result *string) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("item", _args[0:_end]...)
	var (
		_converted *string // javascript: DOMString _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		__tmp := (_returned).String()
		_converted = &__tmp
	}
	_result = _converted
	return
}

func (_this *DOMTokenList) Contains(token string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := token
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("contains", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *DOMTokenList) Add(tokens ...string) {
	var (
		_args []interface{} = make([]interface{}, 0+len(tokens))
		_end  int
	)
	for _, __in := range tokens {
		__out := __in
		_args[_end] = __out
		_end++
	}
	_this.jsValue.Call("add", _args[0:_end]...)
	return
}

func (_this *DOMTokenList) Remove(tokens ...string) {
	var (
		_args []interface{} = make([]interface{}, 0+len(tokens))
		_end  int
	)
	for _, __in := range tokens {
		__out := __in
		_args[_end] = __out
		_end++
	}
	_this.jsValue.Call("remove", _args[0:_end]...)
	return
}

func (_this *DOMTokenList) Toggle(token string, force *bool) (_result bool) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := token
	_args[0] = _p0
	_end++
	if force != nil {

		var _p1 interface{}
		if force != nil {
			_p1 = *(force)
		} else {
			_p1 = nil
		}
		_args[1] = _p1
		_end++
	}
	_returned := _this.jsValue.Call("toggle", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *DOMTokenList) Replace(token string, newToken string) (_result bool) {
	var (
		_args [2]interface{}
		_end  int
	)
	_p0 := token
	_args[0] = _p0
	_end++
	_p1 := newToken
	_args[1] = _p1
	_end++
	_returned := _this.jsValue.Call("replace", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}

func (_this *DOMTokenList) Supports(token string) (_result bool) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := token
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("supports", _args[0:_end]...)
	var (
		_converted bool // javascript: boolean _what_return_name
	)
	_converted = (_returned).Bool()
	_result = _converted
	return
}
