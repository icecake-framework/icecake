package webclientsdk

import (
	"strings"
	"syscall/js"
)

/****************************************************************************
* DOMTokenList
*****************************************************************************/

// DOMTokenList represents a set of space-separated tokens.
// Such a set is returned by Element.classList or HTMLLinkElement.relList, and many others.
//
// https://developer.mozilla.org/en-US/docs/Web/API/DOMTokenList
type DOMTokenList struct {
	jsValue js.Value
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
	ret.tokens = make([]string, 0)
	ret.jsValue = value
	return ret
}

// Length returns the number of tokens in the list.
func (_this *DOMTokenList) Length() int {
	return len(_this.tokens)
}

// String returns the value of the list serialized as a string
//
// https://developer.mozilla.org/en-US/docs/Web/API/DOMTokenList/value
func (_this *DOMTokenList) String() (_ret string) {
	for _, v := range _this.tokens {
		_ret += v + " "
	}
	_ret = strings.TrimRight(_ret, " ")
	return _ret
}

// SetValue clears and sets the list to the given value
//
// https://developer.mozilla.org/en-US/docs/Web/API/DOMTokenList/value
func (_this *DOMTokenList) SetValue(value string) {
	value = strings.Trim(value, " ")
	value = strings.ToLower(value)
	_this.tokens = strings.Split(value, " ")
	// TODO must update the DOM
}

// Item returns an item in the list, determined by its position in the list, its index.
func (_this *DOMTokenList) Item(index int) (_result string) {
	if index >= 0 && index < len(_this.tokens) {
		return _this.tokens[index]
	}
	return ""
}

func (_this *DOMTokenList) Contains(token string) (_result bool) {
	token = strings.Trim(token, " ")
	token = strings.ToLower(token)
	for _, v := range _this.tokens {
		if v == token {
			return true
		}
	}
	return false
}

func (_this *DOMTokenList) Add(tokens ...string) {

	panic("TODO must update the DOM")
	// TODO must update the DOM
}

func (_this *DOMTokenList) Remove(tokens ...string) {
	for _, r := range tokens {
		for i, t := range _this.tokens {
			if t == r {
				_this.tokens = append(_this.tokens[:i], _this.tokens[i+1:]...)
				break
			}
		}
	}
	// TODO must update the DOM
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
	return (_returned).Bool()
	// TODO must update the DOM
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
	return (_returned).Bool()
	// TODO must update the DOM
}
