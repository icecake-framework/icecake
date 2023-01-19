package webclientsdk

import (
	"log"
	"syscall/js"
)

// The DocumentType interface represents a Node containing a doctype.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Document/doctype
type DocumentType struct {
	Node
}

// DocumentTypeFromJS is casting a js.Value into DocumentType.
func DocumentTypeFromJS(value js.Value) *DocumentType {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &DocumentType{}
	ret.jsValue = value
	return ret
}

// Name returns A string, eg "html" for <!DOCTYPE HTML>.
func (_doctype *DocumentType) Name() string {
	if _doctype == nil {
		log.Println("Name() call on a nil Attr")
		return ""
	}
	var ret string
	value := _doctype.jsValue.Get("name")
	ret = (value).String()
	return ret
}

// PublicId returns a string, eg "-//W3C//DTD HTML 4.01//EN", now and empty string for HTML.
func (_doctype *DocumentType) PublicId() string {
	if _doctype == nil {
		log.Println("PublicId() call on a nil Attr")
		return ""
	}
	var ret string
	value := _doctype.jsValue.Get("publicId")
	ret = (value).String()
	return ret
}

// SystemId returns a string, eg "http://www.w3.org/TR/html4/strict.dtd", now an empty string for HTML.
func (_doctype *DocumentType) SystemId() string {
	if _doctype == nil {
		log.Println("SystemId() call on a nil Attr")
		return ""
	}
	var ret string
	value := _doctype.jsValue.Get("systemId")
	ret = (value).String()
	return ret
}
