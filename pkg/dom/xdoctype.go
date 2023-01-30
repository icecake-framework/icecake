package dom

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

// MakeDocumentTypeFromJS is casting a js.Value into DocumentType.
func MakeDocumentTypeFromJS(value js.Value) *DocumentType {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &DocumentType{}
	ret.jsValue = value

	ret.Name = (ret.jsValue.Get("name")).String()
	ret.PublicId = (ret.jsValue.Get("publicId")).String()
	ret.SystemId = (ret.jsValue.Get("SystemId")).String()

	return ret
}
