package browser

import (
	"log"
	"strings"
)

type Name string

// Prefix returns the namespace prefix of the Name, or an empty string if no prefix is specified.
// For example, if the qualified name is xml:lang, the returned prefix is xml.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/prefix
func (_name *Name) Prefix() string {
	if _name == nil {
		log.Println("Previx() call on a nil Name")
		return ""
	}
	if strings.Contains(string(*_name), ":") {
		s := strings.Split(string(*_name), ":")
		return normalize(s[0])
	}
	return ""
}

// LocalName returns the local part of the qualified name of a Name, that is the name of the attribute,
// stripped from any namespace in front of it.
// For example, if the qualified name is xml:lang, the returned local name is lang.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/localName
func (_name *Name) LocalName() (_ret string) {
	if _name == nil {
		log.Println("LocalName() call on a nil Name")
		return ""
	}

	_ret = normalize(string(*_name))
	if strings.Contains(_ret, ":") {
		s := strings.Split(_ret, ":")
		_ret = s[1]
	}
	return _ret
}
