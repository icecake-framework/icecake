package ick

import (
	"syscall/js"
)

// The DocumentType interface represents a Node containing a doctype.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/doctype
type DocumentType struct {
	Node

	Name     string // eg "html" for <!DOCTYPE HTML>.
	PublicId string // eg "-//W3C//DTD HTML 4.01//EN", now and empty string for HTML.
	SystemId string // eg "http://www.w3.org/TR/html4/strict.dtd", now an empty string for HTML.
}

// CastDocumentType is casting a js.Value into DocumentType.
func CastDocumentType(value js.Value) *DocumentType {
	if value.Type() != js.TypeObject {
		ConsoleErrorf("casting DocumentType failed")
		return nil
	}
	ret := new(DocumentType)
	ret.jsValue = value
	ret.Name = (value.Get("name")).String()
	ret.PublicId = (value.Get("publicId")).String()
	ret.SystemId = (value.Get("SystemId")).String()
	return ret
}
