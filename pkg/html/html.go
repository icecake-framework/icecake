package html

import (
	"fmt"
	"io"
	"net/url"
	"reflect"

	"github.com/icecake-framework/icecake/internal/helper"
)

// HTMLString encapsulates a known safe string document fragment.
// It should not be used for string from a third-party, or string with
// unclosed tags or comments.
//
// Use of this type presents a security risk:
// the encapsulated content should come from a trusted source,
// as it will be included verbatim in the output.
type HTMLString string

// WriteHTML write the HTMLString to out
func (strhtml HTMLString) WriteHTML(out io.Writer) {
	io.WriteString(out, string(strhtml))
}

// WriteHTML write the HTMLString to out
func (strhtml HTMLString) String() string {
	return string(strhtml)
}

// Tag of an HTML element
type Tag struct {
	name        string       // internal name, use the factory or SetName to update it
	selfClosing bool         // specify if this is a selfclosing tag, automatically setup by SetName. Use SetSelfClosing to force your value.
	attrs       AttributeMap // map of all tag attributes of any type, including the id if there's one.
	virtualid   string       // used internally by the rendering process
}

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

func (tag Tag) HasName() bool {
	return tag.name == ""
}

// SetName normalizes the name and automaticlally update the SelfClosing according to HTML specifications
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

func (tag *Tag) SetSelfClosing(sc bool) *Tag {
	tag.selfClosing = sc
	return tag
}

func (tag *Tag) Attributes() AttributeMap {
	if tag.attrs == nil {
		tag.attrs = make(AttributeMap)
	}
	return tag.attrs
}

func (tag Tag) RenderOpening(out io.Writer) (selfclosed bool, err error) {
	if tag.name != "" {
		_, err = WriteStrings(out, "<", tag.name, " ", tag.Attributes().String(), ">")
		if err == nil {
			selfclosed = tag.selfClosing
		}
	}
	return
}

func (tag Tag) RenderClosing(out io.Writer) (err error) {
	if tag.name != "" && !tag.selfClosing {
		_, err = WriteStrings(out, "</", tag.name, ">")
	}
	return
}

type DataState struct {
	//Id   string // the id of the current processing component
	//Me   any    // the current processing component, should embedd an HtmlSnippet
	Page any // the current ick page, can be nil
	App  any // the current ick App, can be nil
}

// WriteStringsIf writes one or many strings to w only if the condition is true
func WriteStringsIf(condition bool, w io.Writer, ss ...string) (n int, err error) {
	if !condition {
		return 0, nil
	}
	return WriteStrings(w, ss...)
}

// WriteStrings writes one or many strings to w
func WriteStrings(w io.Writer, ss ...string) (n int, err error) {
	nn := 0
	for _, s := range ss {
		nn, err = WriteString(w, s)
		if err != nil {
			return
		}
		n += nn
	}
	return
}

// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
// If w implements StringWriter, its WriteString method is invoked directly.
// Otherwise, w.Write is called exactly once.
func WriteString(out io.Writer, s string) (n int, err error) {
	return io.WriteString(out, s)
}

// mini helper
func mini(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func debugValue(_v reflect.Value) {
	fmt.Printf("Type: %s\n", _v.Type().String())

	n := _v.Type().NumMethod()
	fmt.Printf("Nb Method: %v\n", n)
	for i := 0; i < n; i++ {
		m := _v.Method(i)
		name := _v.Type().Method(i).Name
		fmt.Printf("Method %v: %s %s '%v'\n", i, name, m.String(), m)
	}

	n = _v.NumField()
	fmt.Printf("Nb Field: %v\n", n)
	for i := 0; i < n; i++ {
		m := _v.Field(i)
		name := _v.Type().Field(i).Name
		fmt.Printf("Field %v: %v %v '%v'\n", i, name, m.Type().String(), m)
	}
}

func debugAny(_v any) {
	fmt.Printf("Type: %v\n", reflect.TypeOf(_v).String())
	fmt.Printf("Type: %v\n", reflect.ValueOf(_v).Interface())

	_, ok := _v.(*url.URL)
	fmt.Printf("Type url.URL: %v\n", ok)

}
