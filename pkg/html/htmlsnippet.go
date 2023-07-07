package html

import (
	"io"

	"github.com/huandu/go-clone"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

type ContentStack struct {
	Stack []ContentComposer
}

func (c ContentStack) Clone() ContentStack {
	var n ContentStack
	if len(c.Stack) > 0 {
		copy := clone.Clone(c.Stack)
		n.Stack = copy.([]ContentComposer)
	}
	return n
}

// Clear clears the rendering stack
func (c *ContentStack) ClearContent() {
	if c != nil {
		c.Stack = c.Stack[:0]
	}
}

// HasContent returns true is the content stack is not nil and it contains at least on item
func (c *ContentStack) HasContent() bool {
	return c != nil && c.Stack != nil && len(c.Stack) > 0
}

// Push adds one or many composers to the rendering stack.
// Returns the snippet to allow chaining calls.
//
// Warning: Struct embedding ICKSnippet should be car of Push returns an ICKSnippet and not the parent stuct type.
func (c *ContentStack) Push(content ...ContentComposer) {
	if c.Stack == nil {
		c.Stack = make([]ContentComposer, 0)
	}
	if len(content) > 0 {
		for _, cmp := range content {
			if cmp != nil {
				c.Stack = append(c.Stack, cmp)
			}
		}
	}
}

// RenderStack writes the HTML string corresponding to the content of the HTML element.
// The default implementation for an HTMLSnippet snippet is to render all the internal stack of composers inside an enclosed HTML tag.
func (c *ContentStack) RenderStack(out io.Writer, parent ickcore.RMetaProvider) (err error) {
	if c.Stack != nil && len(c.Stack) > 0 {
		for _, child := range c.Stack {
			err := Render(out, parent, child)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type ICKSnippet struct {
	meta ickcore.RMetaData // Rendering MetaData.
	tag  Tag               // HTML Element Tag with its attributes.

	ContentStack // HTML composers to render within the enclosed tag.
}

// Ensuring HTMLSnippet implements the right interface
var _ ContentComposer = (*ICKSnippet)(nil)

// Snippet returns a new HTMLSnippet with a given tag name and a map of attributes.
func Snippet(tagname string, attrlist ...string) *ICKSnippet {
	snippet := new(ICKSnippet)
	snippet.Tag().SetTagName(tagname).ParseAttributes(attrlist...)
	return snippet
}

// Clone clones the snippet, without the rendering metadata, nor the id
func (s *ICKSnippet) Clone() *ICKSnippet {
	c := new(ICKSnippet)
	c.tag = *s.tag.Clone()
	c.tag.SetId("")
	c.ContentStack = s.ContentStack.Clone()
	return c
}

func (snippet *ICKSnippet) RMeta() *ickcore.RMetaData {
	return &snippet.meta
}

// return a reference to the snippet's tag. Never nil.
func (s *ICKSnippet) Tag() *Tag {
	if s.tag.AttributeMap == nil {
		s.tag.AttributeMap = make(AttributeMap)
	}
	return &s.tag
}

// BuildTag builds the tag used to render the html element.
// This default implementation of BuildTag does nothing.
// So as the tag may have been preset before rendering.
func (s *ICKSnippet) BuildTag() Tag {
	s.Tag().NoName = true
	return s.tag
}

func (s *ICKSnippet) SetAttribute(aname string, value string) {
	s.Tag().SetAttribute(aname, value)
}

// Id Returns the id of the Snippet.
// Can be empty.
func (s ICKSnippet) Id() string {
	return s.Tag().Id()
}

// SetIf sets the snippet id. This is a shortcut to s.Tag().AttributeMap.SetId(id)
func (s *ICKSnippet) SetId(id string) *ICKSnippet {
	s.Tag().SetId(id)
	return s
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// The default implementation for an HTMLSnippet snippet is to render all the internal stack of composers inside an enclosed HTML tag.
func (s *ICKSnippet) RenderContent(out io.Writer) error {
	return s.RenderStack(out, s)
}

func (s *ICKSnippet) SetBody(content ...ContentComposer) *ICKSnippet {
	s.Push(content...)
	return s
}

// BareSnippet enables creation of simple or complex html strings based on
// an original templating system. BareSnippet rendering is an html element string:
//
//	<tagname [attributes]>[content]</tagname>
//
// content can embed other HTMLsnippets in different ways:
//
//	content = "<ick"
//
// content can be empty. If tagname is empty only the content is rendered.
// BareSnippet can be instantiated by itself or it can be embedded into a struct to define a more customizable html component.
type BareSnippet struct {
	meta ickcore.RMetaData // Rendering MetaData.
	tag  Tag               // HTML Element Tag with its attributes.
	//contantstack []ContentComposer // HTML composers to render within the enclosed tag.
}

// Ensuring HTMLSnippet implements the right interface
var _ TagBuilder = (*BareSnippet)(nil)

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

// Clone clones the snippet, without the rendering metadata, nor the id
func (s *BareSnippet) Clone() *BareSnippet {
	c := new(BareSnippet)
	c.tag = *s.tag.Clone()
	c.tag.SetId("")
	return c
}

// BuildTag builds the tag used to render the html element.
// This default implementation of BuildTag does nothing.
// So as the tag may have been preset before rendering.
func (s *BareSnippet) BuildTag() Tag {
	s.Tag().NoName = true
	return s.tag
}

func (s *BareSnippet) SetAttribute(aname string, value string) {
	s.Tag().SetAttribute(aname, value)
}

// Id Returns the id of the Snippet.
// Can be empty.
func (s BareSnippet) Id() string {
	return s.Tag().Id()
}

// SetIf sets the snippet id. This is a shortcut to s.Tag().AttributeMap.SetId(id)
// func (s *HTMLSnippet) SetId(id string) *HTMLSnippet {
// 	s.Tag().SetId(id)
// 	return s
// }

// AddContent adds one or many composers to the rendering stack.
// Returns the snippet to allow chaining calls.
//
// Warning: Struct embedding HTMLSnippet should be car of AddContent returns an HTMLSnippet and not the parent stuct type.
// func (snippet *HTMLSnippet) AddContent(content ...ContentComposer) *HTMLSnippet {
// 	if snippet.contantstack == nil {
// 		snippet.contantstack = make([]ContentComposer, 0)
// 	}
// 	if len(content) > 0 {
// 		for _, c := range content {
// 			if c != nil {
// 				snippet.contantstack = append(snippet.contantstack, c)
// 			}
// 		}
// 	}
// 	return snippet
// }

// // Clear clears the rendering stack
// func (snippet *HTMLSnippet) ClearContent() {
// 	snippet.contantstack = make([]ContentComposer, 0)
// }

// RenderSnippet writes the HTML string the tag element and the content of the composer to the writer.
// The content is unfolded to look for sub-snippet and every sub-snippet are also written to the writer.
// If the child request an ID, RenderSnippet generates an ID by prefixing its parent id.
// In addition the child is appended into the list of sub-components.
func (parent *BareSnippet) RenderChild(out io.Writer, childs ...ContentComposer) error {
	for _, child := range childs {
		err := Render(out, parent, child)
		if err != nil {
			return err
		}
	}
	return nil
}

// RenderSnippetIf renders the Snippet only if the condition is true otherwise does nothing.
func (parent *BareSnippet) RenderChildIf(condition bool, out io.Writer, childs ...ContentComposer) error {
	if !condition {
		return nil
	}
	return parent.RenderChild(out, childs...)
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// The default implementation for an HTMLSnippet snippet is to render all the internal stack of composers inside an enclosed HTML tag.
// func (s *HTMLSnippet) RenderContent(out io.Writer) (err error) {
// 	if s.contantstack != nil {
// 		return s.RenderChild(out, s.contantstack...)
// 	}
// 	return nil
// }

// // HasContent returns true is the content stack is not nil and it contains at least on item
// func (s HTMLSnippet) HasContent() bool {
// 	return s.contantstack != nil && len(s.contantstack) > 0
// }
