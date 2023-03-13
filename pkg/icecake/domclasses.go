package ick

import (
	"strings"
)

/****************************************************************************
* TokenList
*****************************************************************************/

// Classes represents a set of space-separated tokens.
// Such a set is returned by Element.classList or HTMLLinkElement.relList, and many others.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Classes
type Classes struct {
	jslist *JSValue // the corresponding DOMTokenList. if nil methods work with the cache.
	owner  *Element // the Element the classes belongs to.
	cache  []string // the internal slice of tokens only used when the classes object does not belong to an element yet.
}

// NewClasses init a new Classes and casts a js.Value into DOMTokenList if not nil.
// func NewClasses(_jsvp JSValueProvider) *Classes {
// 	cast := new(Classes)
// 	cast.tokens = make([]string, 0)
// 	if _jsvp != nil {
// 		v := _jsvp.Value()
// 		cast.jslist = &v
// 		str := v.Get("value").String()
// 		split := strings.Split(str, " ")
// 		cast.SetTokens(split...)
// 	}
// 	return cast
// }

/****************************************************************************
* TokenList's properties
*****************************************************************************/

// String returns the value of the list serialized as a string
//
// https://developer.mozilla.org/en-US/docs/Web/API/element/classname
func (_classes Classes) String() (_str string) {
	if _classes.owner != nil {
		_str = _classes.owner.GetString("className")
	} else if _classes.cache != nil {
		for _, v := range _classes.cache {
			_str += v + " "
		}
		_str = strings.TrimRight(_str, " ")
	}
	return _str
}

// Length returns the number of tokens in the list.
func (_classes Classes) Count() int {
	if _classes.jslist != nil {
		return _classes.jslist.GetInt("length")
	}
	return len(_classes.cache)
}

// Item returns an item in the list, determined by its position in the list, its index.
// Returns an empty string if the index is out of range.
func (_classes Classes) At(_index int) string {
	if _classes.jslist != nil {
		return _classes.jslist.Call("item", _index).String()
	} else if _classes.cache != nil && _index >= 0 && _index < len(_classes.cache) {
		return _classes.cache[_index]
	}
	return ""
}

// Has return true if token is found within the list.
// Has is the alias of the webapi.Contains
// token is helper.Normalized before check
func (_classes Classes) Has(_token string) bool {
	if _classes.jslist != nil {
		return _classes.jslist.Call("contains", _token).Bool()
	} else {
		return _classes.hasCache(_token)
	}
	return false
}

// SetClasses adds token in the list. If a token already exist it's not added to avoid duplicate.
//
// Case sensitive.
func (_classes *Classes) AddClasses(_newclasses Classes) *Classes {
	if _classes.jslist != nil {
		str := _newclasses.String()
		if str != "" {
			_classes.jslist.Set("value", str)
		}
	} else {
		_classes.addCache(_newclasses.cache...)
	}
	return _classes
}

// Parse clears and sets the list to the given value.
// returns the chaining call value.
//
// warning: value is case sensitive
func (_classes *Classes) ParseTokens(_value string) (_err error) {
	_err = _classes.parseCache(_value)
	if _classes.jslist != nil {
		_classes.jslist.Set("value", _classes.String())
	}
	return _err
}

// SetTokens adds token in the list. If a token already exist it's not added to avoid duplicate.
//
// Case sensitive.
func (_classes *Classes) AddTokens(_tokens ...string) *Classes {
	if _classes.jslist != nil {
		_classes.jslist.Set("value", _classes.String())
	} else {
		_classes.addCache(_tokens...)
	}
	return _classes
}

// Remove removes tokens in the list or does nothing for the one that does not exist.
// Returns the tokenlist to enable chaining calls.
func (_classes *Classes) RemoveTokens(_tokens ...string) *Classes {
	if _classes.jslist != nil {
		_classes.jslist.Set("value", _classes.String())
	} else {
		_classes.removeCache(_tokens...)
	}
	return _classes
}

// Toggle removes an existing token from the list or add it if it doesn't exist in the list.
// returns true is the token is in the list after the call.
func (_classes *Classes) Toggle(_token string) (_isin bool) {
	if _classes.jslist != nil {
		return _classes.jslist.Call("toggle", _token).Bool()
	} else {
		if _classes.hasCache(_token) {
			_classes.removeCache(_token)
		} else {
			_classes.addCache(_token)
			_isin = true
		}
	}
	return _isin
}

// Replace chains a Remove and a Add
func (_classes *Classes) Replace(_token string, _withToken string) {
	if _classes.jslist != nil {
		_classes.jslist.Call("replace", _token, _withToken)
	} else {
		_classes.removeCache(_token)
		_classes.addCache(_withToken)
	}
}

/*****************************************************************************
* PRIVATE
*****************************************************************************/

func (_classes Classes) hasCache(_token string) bool {
	if _classes.cache != nil {
		for _, v := range _classes.cache {
			if v == _token {
				return true
			}
		}
	}
	return false
}

// Set adds token in the list. If a token already exist it's not added to avoid duplicate.
//
// Case sensitive.
func (_classes *Classes) addCache(_tokens ...string) (_err bool) {
	if _classes.cache == nil {
		_classes.cache = make([]string, 0)
	}
	for _, t := range _tokens {
		t = strings.Trim(t, " ")
		if !_classes.hasCache(t) {
			// TODO: check error
			_classes.cache = append(_classes.cache, t)
		}
	}
	return _err
}

// Remove removes tokens in the list or does nothing for the one that does not exist.
// Returns true if at least one token has been removed
//
// Case sensitive.
func (_classes *Classes) removeCache(_tokens ...string) {
	if _classes.cache != nil {
		for i, t := range _classes.cache {
			for _, r := range _tokens {
				if t == r {
					_classes.cache = append(_classes.cache[:i], _classes.cache[i+1:]...)
				}
			}
		}
	}
}

// Parse clears and sets the list to the given value
//
// warning value is case sensitive
func (_classes *Classes) parseCache(_value string) (_err error) {
	_classes.cache = nil
	split := strings.Split(_value, " ")
	_classes.addCache(split...)
	return nil
}
