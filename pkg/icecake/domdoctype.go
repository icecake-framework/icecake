package ick

import (
	"github.com/sunraylab/icecake/pkg/errors"
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
func CastDocumentType(_jsvp JSValueProvider) *DocumentType {
	// TODO: test defer on error
	jsv := _jsvp.Value()
	if jsv.Type() != TYPE_OBJECT {
		errors.ConsoleErrorf("casting DocumentType failed")
		return &DocumentType{}
	}
	ret := new(DocumentType)
	ret.jsvalue = jsv.jsvalue
	ret.Name = jsv.GetString("name")
	ret.PublicId = jsv.GetString("publicId")
	ret.SystemId = jsv.GetString("SystemId")
	return ret
}
