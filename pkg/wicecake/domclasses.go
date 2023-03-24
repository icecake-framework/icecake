package wick

import ick "github.com/sunraylab/icecake/pkg/icecake0"

/****************************************************************************
* TokenList
*****************************************************************************/

// JSClasses represents a set of space-separated tokens.
// Such a set is returned by Element.classList or HTMLLinkElement.relList, and many others.
//
// https://developer.mozilla.org/en-US/docs/Web/API/JSClasses
type JSClasses struct {
	jslist *JSValue // the corresponding DOMTokenList. if nil methods work with the cache.
	owner  *Element // the Element the classes belongs to.
	// cache  []string // the internal slice of tokens only used when the classes object does not belong to an element yet.
	//	cache ick.Classes // the internal slice of tokens only used when the classes object does not belong to an element yet.
}

/****************************************************************************
* TokenList's properties
*****************************************************************************/

// String returns the value of the list serialized as a string
//
// https://developer.mozilla.org/en-US/docs/Web/API/element/classname
func (_classes JSClasses) String() (_str string) {
	if _classes.owner != nil {
		_str = _classes.owner.GetString("className")
	}
	// else {
	// 	_str = _classes.cache.String()
	// }
	return _str
}

// Length returns the number of tokens in the list.
func (_classes JSClasses) Count() int {
	if _classes.jslist != nil {
		return _classes.jslist.GetInt("length")
	}
	// return _classes.cache.Count()
	return 0
}

// Item returns an item in the list, determined by its position in the list, its index.
// Returns an empty string if the index is out of range.
func (_classes JSClasses) At(_index int) string {
	if _classes.jslist != nil {
		return _classes.jslist.Call("item", _index).String()
	}
	return ""
	// return _classes.cache.At(_index)
}

// Has return true if token is found within the list.
// Has is the alias of the webapi.Contains
// token is helper.Normalized before check
func (_classes JSClasses) Has(_token string) bool {
	if _classes.jslist != nil {
		return _classes.jslist.Call("contains", _token).Bool()
	}
	return false

	// return _classes.cache.Has(_token)
}

// SetClasses adds token in the list. If a token already exist it's not added to avoid duplicate.
//
// Case sensitive.
func (_classes *JSClasses) AddClasses(_newclasses ick.Classes) *JSClasses {
	if _classes.jslist != nil {
		str := _newclasses.String()
		if str != "" {
			_classes.jslist.Set("value", str)
		}
	}
	// else {
	// 	_classes.cache.AddClasses(_newclasses.cache...)
	// }
	return _classes
}

// SetTokens adds token in the list. If a token already exist it's not added to avoid duplicate.
//
// Case sensitive.
func (_classes *JSClasses) AddTokens(_tokens ...string) *JSClasses {
	if _classes.jslist != nil {
		_classes.jslist.Set("value", _classes.String())
	}
	// else {
	// 	_classes.addCache(_tokens...)
	// }
	return _classes
}

// Remove removes tokens in the list or does nothing for the one that does not exist.
// Returns the tokenlist to enable chaining calls.
func (_classes *JSClasses) RemoveTokens(_tokens ...string) *JSClasses {
	if _classes.jslist != nil {
		_classes.jslist.Set("value", _classes.String())
	}
	// } else {
	// 	_classes.removeCache(_tokens...)
	// }
	return _classes
}

// Toggle removes an existing token from the list or add it if it doesn't exist in the list.
// returns true is the token is in the list after the call.
func (_classes *JSClasses) Toggle(_token string) (_isin bool) {
	if _classes.jslist != nil {
		return _classes.jslist.Call("toggle", _token).Bool()
	}
	// else {
	// 	if _classes.hasCache(_token) {
	// 		_classes.removeCache(_token)
	// 	} else {
	// 		_classes.addCache(_token)
	// 		_isin = true
	// 	}
	// }
	return _isin
}

// Replace chains a Remove and a Add
func (_classes *JSClasses) Replace(_token string, _withToken string) {
	if _classes.jslist != nil {
		_classes.jslist.Call("replace", _token, _withToken)
	}
	// else {
	// 	_classes.cache.RemoveCache(_token)
	// 	_classes.addCache(_withToken)
	// }
}
