package html

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

// BareSnippet provides a Rendering MetataData and a Tag required by quite every custom snippet.
// It aims to be embedded into custom snippet.
type BareSnippet struct {
	meta ickcore.RMetaData // Rendering MetaData.
	tag  Tag               // HTML Element Tag with its attributes.
}

// Clone clones the snippet, without the rendering metadata, nor the id
func (s *BareSnippet) Clone() *BareSnippet {
	c := new(BareSnippet)
	c.tag = *s.tag.Clone()
	c.tag.SetId("")
	return c
}

// Id Returns the id of the Snippet.
// Can be empty.
func (s BareSnippet) Id() string {
	return s.Tag().Id()
}

func (snippet *BareSnippet) RMeta() *ickcore.RMetaData {
	return &snippet.meta
}

// return a reference to the snippet's tag. Never nil.
func (s *BareSnippet) Tag() *Tag {
	if s.tag.AttributeMap == nil {
		s.tag.AttributeMap = make(AttributeMap)
	}
	return &s.tag
}

func (s *BareSnippet) SetAttribute(aname string, value string) {
	s.Tag().SetAttribute(aname, value)
}

/******************************************************************************/

type ICKSnippet struct {
	BareSnippet  // Provides a Rendering MetataData and a Tag
	ContentStack // Stack of content composers to render.
}

// Ensuring HTMLSnippet implements the right interface.
var _ ContentComposer = (*ICKSnippet)(nil)

// Snippet returns a new HTMLSnippet with a given tag name and a map of attributes.
func Snippet(tagname string, attr string, body ...ContentComposer) *ICKSnippet {
	s := new(ICKSnippet)
	s.Tag().SetTagName(tagname).ParseAttributes(attr)
	s.Push(body...)
	return s
}

// Clone clones the snippet without the rendering metadata nor its id.
func (s *ICKSnippet) Clone() *ICKSnippet {
	c := new(ICKSnippet)
	c.tag = *s.tag.Clone()
	c.tag.SetId("")
	c.ContentStack = s.ContentStack.Clone()
	return c
}

// SetId sets the tag id property of the snippet.
// This is a shortcut to s.Tag().AttributeMap.SetId(id) and enables ICKSnippet call chaining.
func (s *ICKSnippet) SetId(id string) *ICKSnippet {
	s.Tag().SetId(id)
	return s
}

// SetBody
func (s *ICKSnippet) AddBody(content ...ContentComposer) *ICKSnippet {
	s.Push(content...)
	return s
}

// BuildTag builds the tag used to render the html element.
// This default implementation of BuildTag does nothing.
// So as the tag may have been preset before rendering.
func (s *ICKSnippet) BuildTag() Tag {
	s.Tag().NoName = true
	return s.tag
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// The default implementation for an HTMLSnippet snippet is to render all the internal stack of composers inside an enclosed HTML tag.
func (s *ICKSnippet) RenderContent(out io.Writer) error {
	return s.RenderStack(out, s)
}
