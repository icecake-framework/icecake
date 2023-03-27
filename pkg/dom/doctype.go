package dom

import (
	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/js"
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
func CastDocumentType(_jsvp js.JSValueProvider) *DocumentType {
	// TODO: test defer on error
	if _jsvp.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting DocumentType failed")
		return &DocumentType{}
	}
	ret := new(DocumentType)
	ret.JSValue = _jsvp.Value()
	ret.Name = ret.JSValue.GetString("name")
	ret.PublicId = ret.JSValue.GetString("publicId")
	ret.SystemId = ret.JSValue.GetString("SystemId")
	return ret
}
