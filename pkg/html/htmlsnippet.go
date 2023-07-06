package html

import (
	"io"

	"github.com/huandu/go-clone"
	"github.com/icecake-framework/icecake/pkg/ickcore"
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
	meta         ickcore.RMetaData     // Rendering MetaData.
	tag          Tag                   // HTML Element Tag with its attributes.
	contantstack []HTMLContentComposer // HTML composers to render within the enclosed tag.

	// ds *DataState // a reference to a datastate that can be used for rendering.
}

// Ensure HTMLSnippet implements HTMLComposer interface
var _ HTMLContentComposer = (*HTMLSnippet)(nil)

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
		to.contantstack = copy.([]HTMLContentComposer)
	}
	return to
}

func (snippet *HTMLSnippet) RMeta() *ickcore.RMetaData {
	return &snippet.meta
}

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

// SetIf sets the snippet id. This is a shortcut to s.Tag().AttributeMap.SetId(id)
// func (s *HTMLSnippet) SetId(id string) *HTMLSnippet {
// 	s.Tag().SetId(id)
// 	return s
// }

// AddContent adds one or many HTMLComposer to the rendering stack of this composer.
// Returns the snippet to allow chaining calls.
//
// Warning: Struct embedding HTMLSnippet should be car of AddContent returns an HTMLSnippet and not the parent stuct type.
func (snippet *HTMLSnippet) AddContent(content ...HTMLContentComposer) *HTMLSnippet {
	if snippet.contantstack == nil {
		snippet.contantstack = make([]HTMLContentComposer, 0)
	}
	if len(content) > 0 {
		for _, c := range content {
			if c != nil {
				snippet.contantstack = append(snippet.contantstack, c)
			}
		}
	}
	return snippet
}

// Clear clears the rendering stack
func (snippet *HTMLSnippet) ClearContent() {
	snippet.contantstack = make([]HTMLContentComposer, 0)
}

// InsertSnippet builds and add a single snippet at the end of the content stack.
// InsertSnippet returns the new snippet created and added to the stack.
// func (snippet *HTMLSnippet) InsertSnippet(tagname string, attrlist ...string) *HTMLSnippet {
// 	if snippet.contantstack == nil {
// 		snippet.contantstack = make([]HTMLComposer, 0)
// 	}
// 	s := NewSnippet(tagname, attrlist...)
// 	snippet.contantstack = append(snippet.contantstack, s)
// 	return s
// }

// RenderSnippet writes the HTML string the tag element and the content of the composer to the writer.
// The content is unfolded to look for sub-snippet and every sub-snippet are also written to the writer.
// If the child request an ID, RenderSnippet generates an ID by prefixing its parent id.
// In addition the child is appended into the list of sub-components.
func (parent *HTMLSnippet) RenderChild(out io.Writer, childs ...HTMLContentComposer) error {
	for _, child := range childs {
		err := Render(out, parent, child)
		if err != nil {
			return err
		}
	}
	return nil
}

// RenderSnippetIf renders the Snippet only if the condition is true otherwise does nothing.
func (parent *HTMLSnippet) RenderChildIf(condition bool, out io.Writer, childs ...HTMLContentComposer) error {
	if !condition {
		return nil
	}
	return parent.RenderChild(out, childs...)
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// The default implementation for an HTMLSnippet snippet is to render all the internal stack of composers inside an enclosed HTML tag.
func (s *HTMLSnippet) RenderContent(out io.Writer) (err error) {
	if s.contantstack != nil {
		return s.RenderChild(out, s.contantstack...)
	}
	return nil
}

// HasContent returns true is the content stack is not nil and it contains at least on item
func (s HTMLSnippet) HasContent() bool {
	return s.contantstack != nil && len(s.contantstack) > 0
}
