package html

import (
	"io"

	"github.com/icecake-framework/icecake/internal/helper"
)

// Tag represents the tag of an HTML element with its attributes.
type Tag struct {
	name        string       // internal name, use the factory or SetName to update it
	selfClosing bool         // specify if this is a selfclosing tag, automatically setup by SetName. Use SetSelfClosing to force your value.
	attrs       AttributeMap // map of all tag attributes of any type, including the id if there's one.
	virtualid   string       // used internally by the rendering process but never rendered itself.
	deep        int          // deepness of the tag in a tag tree
	seq         int          // sequence of the tag in a tag tree
}

// Tag factory setting the tag named and allowing to assign a map of attibutes.
func NewTag(name string, amap AttributeMap) *Tag {
	tag := new(Tag)
	tag.SetName(name)
	if amap == nil {
		tag.attrs = make(AttributeMap)
	} else {
		tag.attrs = amap
	}
	return tag
}

// HasName returns if the tag has a name or not.
// A tag without name won't be rendered.
func (tag Tag) HasName() bool {
	return tag.name != ""
}

// SetName normalizes the name and automaticlally update the SelfClosing attribute according to standard HTML5 specifications
// https://developer.mozilla.org/en-US/docs/Glossary/Void_element
func (tag *Tag) SetName(name string) *Tag {
	tag.name = helper.Normalize(name)
	switch tag.name {
	case "area", "base", "br", "col", "embed", "hr", "img", "input", "link", "meta", "param", "source", "track", "wbr":
		tag.selfClosing = true
	default:
		tag.selfClosing = false
	}
	return tag
}

// SetSelfClosing forces the selfclosing flag of the tag.
// Should be called after SetName because SetName update the selfClosing flag.
func (tag *Tag) SetSelfClosing(sc bool) *Tag {
	tag.selfClosing = sc
	return tag
}

// Attributes Returns the reference of the AttributeMap of the tag.
func (tag *Tag) Attributes() AttributeMap {
	if tag.attrs == nil {
		tag.attrs = make(AttributeMap)
	}
	return tag.attrs
}

// ParseAttributes tries to parse attributes to the tag and ignore errors.
// alist will be added or will update existing tag attributes.
// errors are logged out if verbose mode is on.
func (tag *Tag) ParseAttributes(alists ...string) AttributeMap {
	if tag.attrs == nil {
		tag.attrs = make(AttributeMap)
	}
	for _, alist := range alists {
		amap := ParseAttributes(alist)
		for k, v := range amap {
			tag.attrs[k] = v
		}
	}
	return tag.attrs
}

// RenderOpening renders the HTML string of the opening tag including all the attributes.
// if the tag name is empty, nothing is rendered and there's no error.
// Returns selfclosed if the rendered opening tag as been closed too.
// errors mays occurs from the writer only.
func (tag *Tag) RenderOpening(out io.Writer) (selfclosed bool, err error) {
	if tag.name != "" {
		_, err = WriteStrings(out, "<", tag.name, " ", tag.Attributes().String(), ">")
		if err == nil {
			selfclosed = tag.selfClosing
		}
	}
	return
}

// RenderClosing renders the closing tag HTML string to out.
// if the tag name is empty or the tag is a selclosing one, nothing is rendered and there's no error.
// errors mays occurs from the writer only.
func (tag *Tag) RenderClosing(out io.Writer) (err error) {
	if tag.name != "" && !tag.selfClosing {
		_, err = WriteStrings(out, "</", tag.name, ">")
	}
	return
}
