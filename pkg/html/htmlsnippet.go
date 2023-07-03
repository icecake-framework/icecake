package html

import (
	"io"

	"github.com/huandu/go-clone"
)

// HTMLSnippet enables creation of simple or complex html strings based on
// an original templating system. HTMLSnippet rendering is an html element string:
//
//	<tagname [attributes]>[content]</tagname>
//
// content can embed other HTMLsnippets in different ways:
//
//	content = "<ick"
//
// content can be empty. If tagname is empty only the content is rendered.
// HTMLSnippet can be instantiated by itself or it can be embedded into a struct to define a more customizable html component.
type HTMLSnippet struct {
	meta         RenderingMeta  // Rendering MetaData.
	tag          Tag            // HTML Element Tag with its attributes.
	contantstack []HTMLComposer // HTML composers to render within the enclosed tag.

	// ds *DataState // a reference to a datastate that can be used for rendering.
}

// Ensure HTMLSnippet implements HTMLComposer interface
var _ HTMLComposer = (*HTMLSnippet)(nil)

// NewSnippet returns a new HTMLSnippet with a given tag name and a map of attributes.
func NewSnippet(tagname string, attrlist ...string) *HTMLSnippet {
	snippet := new(HTMLSnippet)
	snippet.Tag().SetTagName(tagname).ParseAttributes(attrlist...)
	return snippet
}

// Div returns a new HTMLSnippet with DIV tag name and a map of attributes.
func Div(attrlist ...string) *HTMLSnippet {
	return NewSnippet("div", attrlist...)
}

// Div returns a new HTMLSnippet with DIV tag name and a map of attributes.
func Span(attrlist ...string) *HTMLSnippet {
	return NewSnippet("span", attrlist...)
}

// Div returns a new HTMLSnippet with DIV tag name and a map of attributes.
func P(attrlist ...string) *HTMLSnippet {
	return NewSnippet("p", attrlist...)
}

// Clone clones the snippet, without the rendering metadata
func (src *HTMLSnippet) Clone() *HTMLSnippet {
	to := new(HTMLSnippet)
	to.tag = *src.tag.Clone()
	if len(src.contantstack) > 0 {
		copy := clone.Clone(src.contantstack)
		to.contantstack = copy.([]HTMLComposer)
	}
	return to
}

func (snippet *HTMLSnippet) Meta() *RenderingMeta {
	return &snippet.meta
}

// Tag returns a reference to the snippet tag.
// func (s *HTMLSnippet) Tag() *Tag {
// 	if s.tag.AttributeMap == nil {
// 		s.tag.AttributeMap = make(AttributeMap)
// 	}
// 	return &s.tag
// }

// BuildTag builds the tag used to render the html element.
// This default implementation of BuildTag does nothing.
// So as the tag may have been preset before rendering.
func (s *HTMLSnippet) BuildTag() Tag {
	s.Tag().NoName = true
	return s.tag
}

func (s *HTMLSnippet) SetAttribute(aname string, value string) {
	s.Tag().SetAttribute(aname, value)
}

// return a reference to the snippet's tag. Never nil.
func (s *HTMLSnippet) Tag() *Tag {
	if s.tag.AttributeMap == nil {
		s.tag.AttributeMap = make(AttributeMap)
	}
	return &s.tag
}

// Id Returns the id of the Snippet.
// Can be empty.
func (s HTMLSnippet) Id() string {
	return s.Tag().Id()
}

// SetDataState
// func (snippet *HTMLSnippet) SetDataState(ds *DataState) *HTMLSnippet {
// 	snippet.ds = ds
// 	return snippet
// }

// AddContent adds one or many HTMLComposer to the rendering stack of this composer.
// Returns the snippet to allow chaining calls.
func (snippet *HTMLSnippet) AddContent(content ...HTMLComposer) *HTMLSnippet {
	if snippet.contantstack == nil {
		snippet.contantstack = make([]HTMLComposer, 0)
	}
	snippet.contantstack = append(snippet.contantstack, content...)
	return snippet
}

// InsertSnippet builds and add a single snippet at the end of the content stack.
// InsertSnippet returns the new snippet created and added to the stack.
func (snippet *HTMLSnippet) InsertSnippet(tagname string, attrlist ...string) *HTMLSnippet {
	if snippet.contantstack == nil {
		snippet.contantstack = make([]HTMLComposer, 0)
	}
	s := NewSnippet(tagname, attrlist...)
	snippet.contantstack = append(snippet.contantstack, s)
	return s
}

// RenderSnippet writes the HTML string the tag element and the content of the composer to the writer.
// The content is unfolded to look for sub-snippet and every sub-snippet are also written to the writer.
// If the child request an ID, RenderSnippet generates an ID by prefixing its parent id.
// In addition the child is appended into the list of sub-components.
// TODO: avoid rendering infinite loop
func (parent *HTMLSnippet) RenderChilds(out io.Writer, childs ...HTMLComposer) error {
	err := Render(out, parent, childs...)
	return err
}

// RenderSnippetIf renders the Snippet only if the condition is true otherwise does nothing.
func (parent *HTMLSnippet) RenderChildsIf(condition bool, out io.Writer, childs ...HTMLComposer) error {
	if !condition {
		return nil
	}
	return parent.RenderChilds(out, childs...)
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// The default implementation for an HTMLSnippet snippet is to render all the internal stack of composers inside an enclosed HTML tag.
func (s *HTMLSnippet) RenderContent(out io.Writer) (err error) {
	if s.contantstack != nil {
		return s.RenderChilds(out, s.contantstack...)
	}
	return nil
}

func (s HTMLSnippet) HasBody() bool {
	return s.contantstack != nil
}
