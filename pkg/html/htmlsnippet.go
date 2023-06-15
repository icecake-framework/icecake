package html

import (
	"bytes"
	"io"

	"github.com/icecake-framework/icecake/internal/helper"
)

// HTMLSnippet enables creation of simple or complex html string based on
// an original templating system allowing embedding of other snippets.
// HTMLSnippet output is an html element:
//
//	<tagname [attributes]>[body]</tagname>
//
// It is common to embed a HTMLSnippet into a struct to define an html component.
type HTMLSnippet struct {
	tag  Tag // optional Tag
	body HTMLString
	//attrs     AttributeMap   // map of all attributes of any type
	embedded  map[string]any // instantiated embedded objects
	dataState *DataState
	// attrs     map[string]string // map of all attributes of any type
}

// Ensure HTMLSnippet implements HTMLComposer interface
var _ HTMLComposer = (*HTMLSnippet)(nil)

func NewSnippet(tagname string, amap AttributeMap, body HTMLString) *HTMLSnippet {
	snippet := new(HTMLSnippet)
	snippet.tag.SetName(tagname)
	snippet.tag.attrs = amap
	snippet.body = body
	return snippet
}

func (_snippet *HTMLSnippet) SetDataState(ds *DataState) {
	_snippet.dataState = ds
}

// makemap ensure good memory allocation of the map of attributes
// func (_snippet *HTMLSnippet) makemap() {
// 	if _snippet.attrs == nil {
// 		_snippet.attrs = make(map[string]string)
// 	}
// }

// This implementation of Tag returns the internal tag that may have been setup
// with the Snippet factory.
func (_snippet *HTMLSnippet) Tag() *Tag {
	return &_snippet.tag
}

// Id Returns the unique Id of a Snippet.
func (_snippet *HTMLSnippet) Id() string {
	return _snippet.tag.Attributes().Id()
}

// SetTag setup the Tag of the Snippet
// The HTMLComposer is returned allowing call chaining.
// func (_snippet *HTMLSnippet) SetTag(tag Tag) {
// 	_snippet.tag = tag
// }

// func (_snippet *HTMLSnippet) TagAttributes() AttributeMap {
// 	_snippet.makemap()
// 	return _snippet.attrs
// }

// BodyTemplate returns the HTML template to unfold inside the html element of the HTMLComposer.
// func (_snippet *HTMLSnippet) BodyTemplate() HTMLString {
// 	return _snippet.bodytemplate
// }

// SetBodyTemplates sets the HTML template to unfold inside the html element of the HTMLComposer.
// The HTMLComposer is returned allowing call chaining.
// func (_snippet *HTMLSnippet) SetBodyTemplate(_body HTMLString) HTMLComposer {
// 	_snippet.bodytemplate = _body
// 	return _snippet
// }

// WriteBody writes the HTML string corresponing to the body of the HTML element
// Default Snippet unfolds body property if an, and write it.
// Can be overloaded by a custom snippet embedding HTMLSnippet.
func (_parent HTMLSnippet) WriteBody(out io.Writer) (err error) {
	// FIXME: return somethng else than err
	_, err = io.WriteString(out, string(_parent.body))
	return
}

// Embedded returns the map of embedded components, keyed by their id.
func (_parent HTMLSnippet) Embedded() map[string]any {
	return _parent.embedded
}

// Embed adds subcmp to the map of embedded components within the _parent.
// If a component with the same _id has already been embedded it's replaced.
// Usually the _id is the id of the html element.
func (parent *HTMLSnippet) Embed(id string, subcmp HTMLComposer) {
	strid := helper.Normalize(id)
	if parent.embedded == nil {
		parent.embedded = make(map[string]any, 1)
	}
	parent.embedded[strid] = subcmp
	// DEBUG: fmt.Printf("embedding %q(%v) into %s\n", id, reflect.TypeOf(cmp).String(), s.Id())
}

// String renders and unfold the _snippet and returns its corresponding HTML string
// FIXME handle child html rather than writesnippet
func (_snippet *HTMLSnippet) String() HTMLString {
	out := new(bytes.Buffer)
	_, err := WriteSnippet(out, _snippet, nil, true)
	if err != nil {
		return ""
	}
	return HTMLString(out.String())
}

// RenderChildSnippet builds and unfolds the childsnippet and returns its html string.
// The _snippet is embedded into the _parent.
// RenderChildSnippet does not mount the component into the DOM and so it can't respond to events.
// func (_parent *HTMLSnippet) RenderChildSnippet(childsnippet HTMLComposer) (_html HTMLString) {
// 	out := new(bytes.Buffer)
// 	id, err := WriteSnippet(out, _childsnippet, nil, true)
// 	if err == nil {
// 		_parent.Embed(id, _childsnippet) // need to embed the snippet itself
// 		_html = HTMLString(out.String())
// 	}
// 	return _html
// }

// WriteChildSnippet writes the HTML string of the composer, its tag element and its body, to the writer.
// The body is unfolded to look for sub-snippet and ever sub-snippet are also written to the writer.
// If the child request an ID, WriteChildSnippet generates an ID by prefixing its parent id.
// In addition the child is appended into the list of sub-components.
func (parent *HTMLSnippet) WriteChildSnippet(out io.Writer, childsnippet HTMLComposer) error {
	// FIXME consider rendering a child (unfolding)
	id, err := WriteSnippet(out, childsnippet, nil, true)
	if err == nil {
		parent.Embed(id, childsnippet) // need to embed the snippet itself
	}
	return err
}

// WriteChildSnippet writes the HTML string of the composer, its tag element and its body, to the writer, only if the condition is true, otherwise does nothing.
// The body is unfolded to look for sub-snippet and ever sub-snippet are also written to the writer.
// If the child request an ID, WriteChildSnippet generates an ID by prefixing its parent id.
// In addition the child is appended into the list of sub-components.
func (parent *HTMLSnippet) WriteChildSnippetIf(condition bool, out io.Writer, childsnippet HTMLComposer) error {
	if !condition {
		return nil
	}
	return parent.WriteChildSnippet(out, childsnippet)
}

func (parent *HTMLSnippet) UnfoldHTML(out io.Writer, html HTMLString, ds *DataState) error {
	if len(html) > 0 {
		return unfoldHTML(parent, out, []byte(html), ds, 0)
	}
	return nil
}

// Template returns a SnippetTemplate used to render the html string of a Snippet.
// By default it returns any already setup _snippet properties.
// If overloaded,
// It's used by WriteSnippet, Unfold, RenderSnippet.
// func (_snippet HTMLSnippet) Template(*DataState) (_t SnippetTemplate) {
// 	_t.Tag = _snippet.Tag
// 	_t.Body = _snippet.Body
// 	_t.Attributes = _snippet.Attributes()
// 	return
// }
