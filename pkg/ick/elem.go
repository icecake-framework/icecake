package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

// ICKElem implements the Composer interface for a basic element
type ICKElem struct {
	ickcore.BareSnippet

	Body ickcore.ContentStack // HTML Element body. A stack of content composers to render.
}

// Ensuring ICKElem implements the right interface.
var _ ickcore.ContentComposer = (*ICKElem)(nil)
var _ ickcore.TagBuilder = (*ICKElem)(nil)

// Elem returns a new HTMLSnippet with a given tag name and a map of attributes.
func Elem(tagname string, attr string, body ...ickcore.ContentComposer) *ICKElem {
	s := new(ICKElem)
	s.Tag().SetTagName(tagname).ParseAttributes(attr)
	s.Body.Push(body...)
	return s
}

// Clone clones the snippet without the rendering metadata nor its id.
func (s *ICKElem) Clone() *ICKElem {
	c := new(ICKElem)
	c.BareSnippet = *s.BareSnippet.Clone()
	c.Body = s.Body.Clone()
	return c
}

// SetId sets the tag id property of the snippet.
// This is a shortcut to s.Tag().AttributeMap.SetId(id) and enables ICKSnippet call chaining.
func (s *ICKElem) SetId(id string) *ICKElem {
	s.Tag().SetId(id)
	return s
}

// Append
func (s *ICKElem) Append(content ...ickcore.ContentComposer) *ICKElem {
	s.Body.Push(content...)
	return s
}

func (s ICKElem) NeedRendering() bool {
	return !s.Tag().IsEmpty() || s.Body.NeedRendering()
}

// BuildTag builds the tag used to render the html element.
// This default implementation of BuildTag does nothing.
// So as the tag may have been preset before rendering.
func (s *ICKElem) BuildTag() ickcore.Tag {
	s.Tag().NoName = true
	return *s.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// The default implementation for an HTMLSnippet snippet is to render all the internal stack of composers inside an enclosed HTML tag.
func (s *ICKElem) RenderContent(out io.Writer) error {
	return s.Body.RenderStack(out, s)
}
