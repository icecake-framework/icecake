package lib

import (
	"strings"
)

/****************************************************************************
* Tokens
*****************************************************************************/

// Tokens is a slice of case sensitive string
type Tokens []string

func MakeTokens(value string) (_new Tokens) {
	split := strings.Split(value, " ")
	_new.Set(split...)
	return _new
}

/****************************************************************************
* Tokens's properties
*****************************************************************************/

// String returns the value of the slice serialized as a string
//
// https://developer.mozilla.org/en-US/docs/Web/API/DOMTokens/value
func (_toks *Tokens) String() (_ret string) {
	for _, v := range *_toks {
		_ret += v + " "
	}
	_ret = strings.TrimRight(_ret, " ")
	return _ret
}

// Has return true if token is found within the list.
// Has is the alias of the webapi.Contains
// token is helper.Normalized before check.
//
// Check is cas sensitive.
func (_toks *Tokens) Has(token string) (_result bool) {
	return _toks.has(token)
}

// Has return true if token is found within the list.
// token is not helper.Normalized before check
func (_toks *Tokens) has(token string) (_result bool) {
	for _, v := range *_toks {
		if v == token {
			return true
		}
	}
	return false
}

// Parse clears and sets the list to the given value
//
// warning value is case sensitive
func (_toks *Tokens) Parse(value string) (_changed bool) {
	before := _toks.String()
	split := strings.Split(value, " ")
	_toks.Set(split...)
	after := _toks.String()
	return after != before
}

// SetToken adds token in the list. If a token already exist it's not added to avoid duplicate.
//
// Case sensitive.
func (_toks *Tokens) Set(tokens ...string) (_updated bool) {
	_updated = false
	for _, t := range tokens {
		if !_toks.has(t) {
			*_toks = append(*_toks, t)
			_updated = true
		}
	}
	return _updated
}

// Remove removes tokens in the list or does nothing for the one that does not exist.
// Returns true if at least one token has been removed
//
// Case sensitive.
func (_toks *Tokens) Remove(tokens ...string) (_updated bool) {
	_updated = false
	for i, t := range *_toks {
		for _, r := range tokens {
			if t == r {
				*_toks = append((*_toks)[:i], (*_toks)[i+1:]...)
				_updated = true
			}
		}
	}
	return _updated
}

// Toggle removes an existing token from the list or add it if it doesn't exist in the list.
//
// Case sensitive.
func (_toks *Tokens) Toggle(token string) (_updated bool) {
	_updated = false
	if _toks.has(token) {
		_updated = _toks.Remove(token)
	} else {
		_updated = _toks.Set(token)
	}
	return _updated
}

// Replace chains Remove and Add. Case sensitive.
func (_toks *Tokens) Replace(token string, newToken string) (_updated bool) {
	_updated = _toks.Remove(token)
	_updated = _toks.Set(newToken) || _updated
	return _updated
}
