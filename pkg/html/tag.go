package html

import (
	"io"
	"reflect"
	"strings"

	"github.com/icecake-framework/icecake/internal/helper"
)

// Tag represents the tag of an HTML element with its attributes.
type Tag struct {
	AttributeMap      // map of all tag attributes of any type, including the id if there's one.
	NoName       bool // does not generate name attribute with the ick-name of the composer

	tagname     string // the name of the tag a.k.a "div", "span", "meta"...
	selfClosing bool   // specifies if this is a selfclosing tag, automatically setup by SetTagName. Use SetSelfClosing to force the value.
}

// Tag factory setting the tag named and allowing to assign an existing map of attibutes.
func NewTag(name string, amap AttributeMap) *Tag {
	tag := new(Tag)
	if amap == nil {
		tag.AttributeMap = make(AttributeMap)
	} else {
		tag.AttributeMap = amap
	}
	tag.SetTagName(name)
	return tag
}

// BuildTag get a tag from the TagBuilder then set up name attribute and RMeta id
func BuildTag(tb TagBuilder) Tag {
	tag := tb.BuildTag()

	if tag.HasRendering() {
		// force property name
		if !tag.NoName {
			cmpname := reflect.TypeOf(tb).Elem().Name()
			cmpname = strings.ToLower(cmpname)
			tag.SetAttribute("name", cmpname)
		}
		tb.RMeta().Id = tag.Id()
	}

	return tag
}

// clone clones the tag
func (src Tag) Clone() *Tag {
	to := new(Tag)
	to.AttributeMap = src.AttributeMap.Clone()
	to.NoName = src.NoName
	to.tagname = src.tagname
	to.selfClosing = src.selfClosing
	return to
}

// HasRendering returns if the tag can be rendered or not.
// A tag without name won't be rendered.
func (tag Tag) HasRendering() bool {
	return tag.tagname != ""
}

// TagName returns the name of the tag and it's selfclosing flag
func (tag *Tag) TagName() (string, bool) {
	return tag.tagname, tag.selfClosing
}

// SetTagName normalizes the name and automaticlally update the SelfClosing attribute according to standard HTML5 specifications
// https://developer.mozilla.org/en-US/docs/Glossary/Void_element
func (tag *Tag) SetTagName(name string) *Tag {
	if tag.AttributeMap == nil {
		tag.AttributeMap = make(AttributeMap)
	}
	tag.tagname = helper.Normalize(name)
	switch tag.tagname {
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

// ParseAttributes tries to parse attributes to the tag and ignore errors.
// alist will be added or will update existing tag attributes.
// errors are logged out if verbose mode is on.
func (tag *Tag) ParseAttributes(alists ...string) AttributeMap {
	if tag.AttributeMap == nil {
		tag.AttributeMap = make(AttributeMap)
	}
	for _, alist := range alists {
		amap := ParseAttributes(alist)
		for k, v := range amap {
			tag.AttributeMap[k] = v
		}
	}
	return tag.AttributeMap
}

// RenderOpening renders the HTML string of the opening tag including all the attributes.
// if the tag name is empty, nothing is rendered and there's no error.
// Returns selfclosed if the rendered opening tag as been closed too.
// errors mays occurs from the writer only.
func (tag *Tag) RenderOpening(out io.Writer) (selfclosed bool, err error) {
	if tag.tagname != "" {
		stratt := tag.AttributeString()
		_, err = RenderString(out, "<", tag.tagname)
		if err != nil {
			return
		}
		_, err = RenderStringIf(stratt != "", out, " ", tag.AttributeString())
		if err != nil {
			return
		}
		_, err = RenderString(out, ">")
		if err != nil {
			return
		}
		selfclosed = tag.selfClosing
	}
	return
}

// RenderClosing renders the closing tag HTML string to out.
// if the tag name is empty or the tag is a selclosing one, nothing is rendered and there's no error.
// errors mays occurs from the writer only.
func (tag *Tag) RenderClosing(out io.Writer) (err error) {
	if tag.tagname != "" && !tag.selfClosing {
		_, err = RenderString(out, "</", tag.tagname, ">")
	}
	return
}
