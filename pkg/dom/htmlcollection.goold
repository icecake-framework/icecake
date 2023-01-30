package browser

import "syscall/js"

// class: HTMLCollection
//
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCollection
type HTMLCollection struct {
	jsValue js.Value
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *HTMLCollection) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}
	return _this.jsValue
}

// newHTMLCollectionFromJS is casting a js.Value into HTMLCollection.
func newHTMLCollectionFromJS(value js.Value) *HTMLCollection {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &HTMLCollection{}
	ret.jsValue = value
	return ret
}

// Length returning attribute 'length' with
// type uint (idl: unsigned long).
func (_this *HTMLCollection) Length() uint {
	var ret uint
	value := _this.jsValue.Get("length")
	ret = (uint)((value).Int())
	return ret
}

func (_this *HTMLCollection) Index(index uint) (_result *Element) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("item", _args[0:_end]...)
	var (
		_converted *Element // javascript: Element _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = NewElementFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *HTMLCollection) Get(name string) (_result *Element) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("namedItem", _args[0:_end]...)
	var (
		_converted *Element // javascript: Element _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = NewElementFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *HTMLCollection) Item(index uint) (_result *Element) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := index
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("item", _args[0:_end]...)
	var (
		_converted *Element // javascript: Element _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = NewElementFromJS(_returned)
	}
	_result = _converted
	return
}

func (_this *HTMLCollection) NamedItem(name string) (_result *Element) {
	var (
		_args [1]interface{}
		_end  int
	)
	_p0 := name
	_args[0] = _p0
	_end++
	_returned := _this.jsValue.Call("namedItem", _args[0:_end]...)
	var (
		_converted *Element // javascript: Element _what_return_name
	)
	if _returned.Type() != js.TypeNull && _returned.Type() != js.TypeUndefined {
		_converted = NewElementFromJS(_returned)
	}
	_result = _converted
	return
}
