package ick

import (
	"strings"
)

/****************************************************************************
* Classes
*****************************************************************************/

// Classes represents a safe set of space-separated tokens.
// tokens are case sensitives.
type Classes struct {
	tokens []string // internal slice of tokens
}

// ParseClasses returns a new classes composed of extracted space-separated tokens.
func ParseClasses(_tokenlist string) (_c *Classes, _err error) {
	split := strings.Fields(_tokenlist)
	c := new(Classes)
	if _err = c.AddTokens(split...); _err != nil {
		c = nil
	}
	return c, _err
}

/****************************************************************************
* Properties
*****************************************************************************/

// String returns the value of the list serialized as a string
func (_classes Classes) String() (_str string) {
	if _classes.tokens != nil {
		for _, v := range _classes.tokens {
			_str += v + " "
		}
		_str = strings.TrimRight(_str, " ")
	}
	return _str
}

// Count returns the number of tokens in the list.
func (_classes Classes) Count() int {
	return len(_classes.tokens)
}

// Item returns an item in the list, determined by its position in the list, its index.
// Returns an empty string if the index is out of range.
func (_classes Classes) At(_index int) string {
	if _classes.tokens != nil && _index >= 0 && _index < len(_classes.tokens) {
		return _classes.tokens[_index]
	}
	return ""
}

// Has returns true if token is found within the list.
func (_classes Classes) Has(_token string) bool {
	if _classes.tokens != nil {
		for _, v := range _classes.tokens {
			if v == _token {
				return true
			}
		}
	}
	return false
}

// AddClasses adds all _newclasses tokens in the list.
// If a token already exist it's not added to avoid duplicate.
func (_classes *Classes) AddClasses(_newclasses Classes) *Classes {
	_classes.AddTokens(_newclasses.tokens...)
	return _classes
}

// AddTokens adds token in the list. If a token already exist it's not added to avoid duplicate.
func (_classes *Classes) AddTokens(_tokens ...string) error {
	if _classes.tokens == nil {
		_classes.tokens = make([]string, 0)
	}
	for _, t := range _tokens {
		t = strings.Trim(t, " ")
		// TODO: check token name validity

		if t != "" && !_classes.Has(t) {
			_classes.tokens = append(_classes.tokens, t)
		}
	}
	return nil
}

// Removetokens removes tokens in the list or does nothing for the one that does not exist.
// Returns true if at least one token has been removed
func (_classes *Classes) RemoveTokens(_tokens ...string) *Classes {
	if _classes.tokens != nil {
		for i, t := range _classes.tokens {
			for _, r := range _tokens {
				if t == r {
					_classes.tokens = append(_classes.tokens[:i], _classes.tokens[i+1:]...)
				}
			}
		}
	}
	return _classes
}

// Toggle removes an existing token from the list or add it if it doesn't exist in the list.
// returns true is the token is in the list after the call.
func (_classes *Classes) Toggle(_token string) (_isin bool) {
	if _classes.Has(_token) {
		_classes.RemoveTokens(_token)
	} else {
		_classes.AddTokens(_token)
		_isin = true
	}
	return _isin
}

// Replace chains a Remove and a Add
func (_classes *Classes) Replace(_token string, _withToken string) *Classes {
	_classes.RemoveTokens(_token)
	_classes.AddTokens(_withToken)
	return _classes
}
