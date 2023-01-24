package browser

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

// // JSValue returns the js.Value or js.Null() if _this is nil
// // Calling JSValue updates the DOM.
// func (_thisList *TokenList) JSValue() js.Value {
// 	if _thisList == nil {
// 		return js.Null()
// 	}
// 	_thisList.ToDOM()
// 	return _thisList.jsValue
// }

// DOMTokenListFromJS is casting a js.Value into DOMTokenList.
func NewTokenListFromJS(value js.Value) (_ret *TokenList) {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	_ret = new(TokenList)
	_ret.jsValue = value

	_ret.tokens = make([]string, 0)
	list := _ret.jsValue.Get("value").String()
	list = normalize(list)
	_ret.tokens = strings.Split(list, " ")
	return _ret
}

// ToDOM update the DOM with the internal value.
// func (_thisList *TokenList) ToDOM() {
// 	_thisList.jsValue.Set("value", _thisList.String())
// }

/****************************************************************************
* TokenList's properties
*****************************************************************************/

// Length returns the number of tokens in the list.
func (_thisList *TokenList) Count() int {
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

// Item returns an item in the list, determined by its position in the list, its index.
// Returns an empty string if the index is out of range.
func (_thisList *TokenList) Index(index int) (_result string) {
	if index >= 0 && index < len(_thisList.tokens) {
		return _thisList.tokens[index]
	}
	return ""
}

// Has return true if token is found within the list.
// Has is the alias of the webapi.Contains
// token is normalized before check
func (_thisList *TokenList) Has(token string) (_result bool) {
	token = normalize(token)
	return _thisList.has(token)
}

// Has return true if token is found within the list.
// token is not normalized before check
func (_thisList *TokenList) has(token string) (_result bool) {
	for _, v := range _thisList.tokens {
		if v == token {
			return true
		}
	}
	return false
}

// SetValue clears and sets the list to the given value
//
// https://developer.mozilla.org/en-US/docs/Web/API/DOMTokenList/value
func (_thisList *TokenList) ParseList(value string) *TokenList {
	value = normalize(value)
	before := _thisList.String()
	_thisList.tokens = strings.Split(value, " ")
	after := _thisList.String()
	if after != before {
		_thisList.jsValue.Set("value", after)
	}
	return _thisList
}

// SetToken adds token in the list. If a token already exist it's not added to avoid duplicate.
// Always converted in lowercase.
func (_thisList *TokenList) Set(tokens ...string) *TokenList {
	if _thisList.setTokens(tokens...) {
		_thisList.jsValue.Set("value", _thisList.String())
	}
	return _thisList
}

// SetToken adds token in the list. If a token already exist it's not added to avoid duplicate.
// Always converted in lowercase.
func (_thisList *TokenList) setTokens(tokens ...string) (_updated bool) {
	_updated = false
	for _, t := range tokens {
		t = normalize(t)
		if !_thisList.has(t) {
			_thisList.tokens = append(_thisList.tokens, t)
			_updated = true
		}
	}
	return _updated
}

// Remove removes tokens in the list or does nothing for the one that does not exist.
// Returns the tokenlist to enable chaining calls.
func (_thisList *TokenList) Remove(tokens ...string) *TokenList {
	if _thisList.removeTokens(tokens...) {
		_thisList.jsValue.Set("value", _thisList.String())
	}
	return _thisList
}

// Remove removes tokens in the list or does nothing for the one that does not exist.
// Returns true if at least one token has been removed
func (_thisList *TokenList) removeTokens(tokens ...string) (_updated bool) {
	_updated = false
	for i, t := range _thisList.tokens {
		t = normalize(t)
		for _, r := range tokens {
			if t == r {
				_thisList.tokens = append(_thisList.tokens[:i], _thisList.tokens[i+1:]...)
				_updated = true
				break
			}
		}
	}
	return _updated
}

// Toggle removes an existing token from the list or add it if it doesn't exist in the list.
func (_thisList *TokenList) Toggle(token string) *TokenList {
	updated := false
	token = normalize(token)
	if _thisList.has(token) {
		updated = _thisList.removeTokens(token)
	} else {
		updated = _thisList.setTokens(token)
	}
	if updated {
		_thisList.jsValue.Set("value", _thisList.String())
	}
	return _thisList
}

// Replace chain a Remove and a Add
func (_thisList *TokenList) Replace(token string, newToken string) *TokenList {
	updated := _thisList.removeTokens(token)
	updated = _thisList.setTokens(token) || updated
	if updated {
		_thisList.jsValue.Set("value", _thisList.String())
	}
	return _thisList
}
