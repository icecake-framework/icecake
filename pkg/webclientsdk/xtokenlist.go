package webclientsdk

import (
	"strings"
	"syscall/js"
)

/****************************************************************************
* TokenList
*****************************************************************************/

// TokenList represents a set of space-separated tokens.
// Such a set is returned by Element.classList or HTMLLinkElement.relList, and many others.
//
// # Need to call ToDOM() to update the DOM with internal value avec any change
//
// https://developer.mozilla.org/en-US/docs/Web/API/TokenList
type TokenList struct {
	jsValue js.Value
	tokens  []string
}

// JSValue returns the js.Value or js.Null() if _this is nil
// Calling JSValue updates the DOM.
func (_thisList *TokenList) JSValue() js.Value {
	if _thisList == nil {
		return js.Null()
	}
	_thisList.ToDOM()
	return _thisList.jsValue
}

// DOMTokenListFromJS is casting a js.Value into DOMTokenList.
func MakeTokenListFromJS(value js.Value) *TokenList {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &TokenList{}
	ret.jsValue = value
	ret.tokens = make([]string, 0)
	ret.SetValue(ret.jsValue.Get("value").String())
	return ret
}

// ToDOM update the DOM with the internal value.
func (_thisList *TokenList) ToDOM() {
	_thisList.jsValue.Set("value", _thisList.String())
}

/****************************************************************************
* TokenList's properties
*****************************************************************************/

// Length returns the number of tokens in the list.
func (_thisList *TokenList) Length() int {
	return len(_thisList.tokens)
}

// String returns the value of the list serialized as a string
//
// https://developer.mozilla.org/en-US/docs/Web/API/DOMTokenList/value
func (_thisList *TokenList) String() (_ret string) {
	for _, v := range _thisList.tokens {
		_ret += v + " "
	}
	_ret = strings.TrimRight(_ret, " ")
	return _ret
}

// SetValue clears and sets the list to the given value
//
// https://developer.mozilla.org/en-US/docs/Web/API/DOMTokenList/value
func (_thisList *TokenList) SetValue(value string) *TokenList {
	value = normalize(value)
	_thisList.tokens = strings.Split(value, " ")
	return _thisList
}

// Item returns an item in the list, determined by its position in the list, its index.
// Returns an empty string if the index is out of range.
func (_thisList *TokenList) Item(index int) (_result string) {
	if index >= 0 && index < len(_thisList.tokens) {
		return _thisList.tokens[index]
	}
	return ""
}

// Has return true if token is found within the list.
// Has is the alias of the webapi.Contains
func (_thisList *TokenList) Has(token string) (_result bool) {
	token = normalize(token)
	return _thisList.has(token)
}

func (_thisList *TokenList) has(token string) (_result bool) {
	for _, v := range _thisList.tokens {
		if v == token {
			return true
		}
	}
	return false
}

// SetToken adds token in the list. If a token already exist it's not added to avoid duplicate.
// Always converted in lowercase.
func (_thisList *TokenList) SetTokens(tokens ...string) *TokenList {
	for _, t := range tokens {
		t = normalize(t)
		if !_thisList.has(t) {
			_thisList.tokens = append(_thisList.tokens, t)
		}
	}
	return _thisList
}

// Remove removes tokens in the list or does nothing for the one that does not exist.
// Duplicates are removed too.
func (_thisList *TokenList) RemoveTokens(tokens ...string) *TokenList {
	for i, t := range _thisList.tokens {
		t = normalize(t)
		for _, r := range tokens {
			if t == r {
				_thisList.tokens = append(_thisList.tokens[:i], _thisList.tokens[i+1:]...)
				break
			}
		}
	}
	return _thisList
}

// Toggle removes an existing token from the list or add it if it doesn't exist in the list.
func (_thisList *TokenList) Toggle(token string) *TokenList {
	token = normalize(token)
	if _thisList.has(token) {
		_thisList.RemoveTokens(token)
	} else {
		_thisList.SetTokens(token)
	}
	return _thisList
}

// Replace chain a Remove and a Add
func (_thisList *TokenList) Replace(token string, newToken string) *TokenList {
	_thisList.RemoveTokens(token)
	_thisList.SetTokens(token)
	return _thisList
}
