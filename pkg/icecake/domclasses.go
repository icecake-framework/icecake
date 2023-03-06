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
// # Need to call ToDOM() to update the DOM with internal value avec any change
//
// https://developer.mozilla.org/en-US/docs/Web/API/Classes
type Classes struct {
	jslist *JSValue // the corresponding DOMTokenList. The DOM is only updated when jsList is not nil.
	tokens []string // the internal slice of tokens
}

// NewClasses init a new Classes and casts a js.Value into DOMTokenList if not nil.
func NewClasses(_jsvp JSValueProvider) *Classes {
	cast := new(Classes)
	cast.tokens = make([]string, 0)
	if _jsvp != nil {
		v := _jsvp.Value()
		cast.jslist = &v
		str := v.Get("value").String()
		split := strings.Split(str, " ")
		cast.SetTokens(split...)
	}
	return cast
}

// ParseClasses split _str into classes separated by spaces
func ParseClasses(_str string) *Classes {
	classes := NewClasses(nil)
	classes.parseTokens(_str)
	return classes
}

/****************************************************************************
* TokenList's properties
*****************************************************************************/

// Length returns the number of tokens in the list.
func (_classes Classes) Count() int {
	return len(_classes.tokens)
}

// String returns the value of the list serialized as a string
//
// https://developer.mozilla.org/en-US/docs/Web/API/DOMTokenList/value
func (_classes Classes) String() (_str string) {
	for _, v := range _classes.tokens {
		_str += v + " "
	}
	_str = strings.TrimRight(_str, " ")
	return _str
}

// Item returns an item in the list, determined by its position in the list, its index.
// Returns an empty string if the index is out of range.
func (_classes Classes) At(_index int) string {
	if _index >= 0 && _index < len(_classes.tokens) {
		return _classes.tokens[_index]
	}
	return ""
}

// Has return true if token is found within the list.
// Has is the alias of the webapi.Contains
// token is helper.Normalized before check
func (_classes Classes) Has(_token string) bool {
	for _, v := range _classes.tokens {
		if v == _token {
			return true
		}
	}
	return false
}

// Parse clears and sets the list to the given value.
// returns the chaining call value.
//
// warning: value is case sensitive
func (_classes *Classes) Parse(_value string) *Classes {
	if _classes.parseTokens(_value) && _classes.jslist != nil {
		_classes.jslist.Set("value", _classes.String())
	}
	return _classes
}

// SetTokens adds token in the list. If a token already exist it's not added to avoid duplicate.
//
// Case sensitive.
func (_classes *Classes) SetTokens(_tokens ...string) *Classes {
	if _classes.setTokens(_tokens...) && _classes.jslist != nil {
		_classes.jslist.Set("value", _classes.String())
	}
	return _classes
}

// Remove removes tokens in the list or does nothing for the one that does not exist.
// Returns the tokenlist to enable chaining calls.
func (_classes *Classes) RemoveTokens(_tokens ...string) *Classes {
	if _classes.removeTokens(_tokens...) && _classes.jslist != nil {
		_classes.jslist.Set("value", _classes.String())
	}
	return _classes
}

// Toggle removes an existing token from the list or add it if it doesn't exist in the list.
func (_classes *Classes) Toggle(_token string) bool {
	updated := false
	if _classes.Has(_token) {
		updated = _classes.removeTokens(_token)
	} else {
		updated = _classes.setTokens(_token)
	}
	if updated && _classes.jslist != nil {
		_classes.jslist.Set("value", _classes.String())
	}
	return updated
}

// Replace chain a Remove and a Add
func (_classes *Classes) Replace(_token string, _withToken string) bool {
	updated := _classes.removeTokens(_token)
	updated = _classes.setTokens(_withToken) || updated
	if updated && _classes.jslist != nil {
		_classes.jslist.Set("value", _classes.String())
	}
	return updated
}

/*****************************************************************************/

// Set adds token in the list. If a token already exist it's not added to avoid duplicate.
//
// Case sensitive.
func (_classes *Classes) setTokens(_tokens ...string) (_updated bool) {
	_updated = false
	for _, t := range _tokens {
		if !_classes.Has(t) {
			_classes.tokens = append(_classes.tokens, t)
			_updated = true
		}
	}
	return _updated
}

// Remove removes tokens in the list or does nothing for the one that does not exist.
// Returns true if at least one token has been removed
//
// Case sensitive.
func (_classes *Classes) removeTokens(_tokens ...string) (_updated bool) {
	_updated = false
	for i, t := range _classes.tokens {
		for _, r := range _tokens {
			if t == r {
				_classes.tokens = append(_classes.tokens[:i], _classes.tokens[i+1:]...)
				_updated = true
			}
		}
	}
	return _updated
}

// Parse clears and sets the list to the given value
//
// warning value is case sensitive
func (_classes *Classes) parseTokens(_value string) (_changed bool) {
	before := _classes.String()
	split := strings.Split(_value, " ")
	_classes.setTokens(split...)
	after := _classes.String()
	return after != before
}
