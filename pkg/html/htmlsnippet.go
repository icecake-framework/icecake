package html

import (
	"io"
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
	content []HTMLComposer // HTML composers to render within the enclosed tag.
	tag     Tag            // HTML Element Tag with its attributes.
	meta    RenderingMeta  // Rendering MetaData

	ds *DataState // a reference to a datastate that can be used for rendering.
}

// Ensure HTMLSnippet implements HTMLComposer interface
var _ HTMLComposer = (*HTMLSnippet)(nil)

// NewSnippet returns a new HTMLSnippet with a given tag name, a map of attributes and a content.
func NewSnippet(tagname string, attrlist ...string) *HTMLSnippet {
	snippet := new(HTMLSnippet)
	snippet.tag.SetTagName(tagname)
	snippet.tag.ParseAttributes(attrlist...)
	return snippet
}

func (snippet *HTMLSnippet) Meta() *RenderingMeta {
	return &snippet.meta
}

// SetDataState
func (snippet *HTMLSnippet) SetDataState(ds *DataState) *HTMLSnippet {
	snippet.ds = ds
	return snippet
}

// StackContent add one of many HTMLComposer for rendering inside the HTML tag of the snippet
// Returns the snippet to allow call chaining.
func (snippet *HTMLSnippet) StackContent(content ...HTMLComposer) *HTMLSnippet {
	if snippet.content == nil {
		snippet.content = make([]HTMLComposer, 0)
	}
	snippet.content = append(snippet.content, content...)
	return snippet
}

// InsertSnippet builds and add a single snippet at the end of the content stack.
// InsertSnippet returns the new snippet created and added to the stack.
func (snippet *HTMLSnippet) InsertSnippet(tagname string, attrlist ...string) *HTMLSnippet {
	if snippet.content == nil {
		snippet.content = make([]HTMLComposer, 0)
	}
	s := NewSnippet(tagname, attrlist...)
	snippet.content = append(snippet.content, s)
	return s
}

// InsertHTML stacks a simple HTMLString to the snippet content.
// Returns the snippet to allow call chaining.
func (snippet *HTMLSnippet) InsertHTML(html HTMLString) *HTMLSnippet {
	if snippet.content == nil {
		snippet.content = make([]HTMLComposer, 0)
	}
	snippet.content = append(snippet.content, &html)
	return snippet
}

// Tag returns a reference to the snippet tag.
func (s *HTMLSnippet) Tag() *Tag {
	if s.tag.AttributeMap == nil {
		s.tag.AttributeMap = make(AttributeMap)
	}
	return &s.tag
}

// BuildTag builds the tag used to render the html element.
// This default implementation of BuildTag does nothing.
// So as the tag may have been preset before rendering.
func (s *HTMLSnippet) BuildTag(tag *Tag) {
}

// Id Returns the id of the Snippet.
// Can be empty.
func (s HTMLSnippet) Id() string {
	return s.tag.Id()
}

// RenderSnippet writes the HTML string the tag element and the content of the composer to the writer.
// The content is unfolded to look for sub-snippet and every sub-snippet are also written to the writer.
// If the child request an ID, RenderSnippet generates an ID by prefixing its parent id.
// In addition the child is appended into the list of sub-components.
func (parent *HTMLSnippet) RenderChilds(out io.Writer, childs ...HTMLComposer) error {
	for _, child := range childs {
		err := Render(out, parent, child)
		if err != nil {
			return err
		}
	}
	return nil
}

// RenderSnippetIf renders the Snippet only if the condition is true otherwise does nothing.
func (parent *HTMLSnippet) RenderChildsIf(condition bool, out io.Writer, childs ...HTMLComposer) error {
	if !condition {
		return nil
	}
	return parent.RenderChilds(out, childs...)
}

// RenderChildHTML renders an HTML template string into out.
// func (parent *HTMLSnippet) RenderChildHTML(out io.Writer, html HTMLString) error {
// 	return RenderHTML(out, parent, html)
// }

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// The default implementation for an HTMLSnippet snippet is to render all the internal stack of composers.
// Use an HTMLString snippet to renders (unfold and generate HTML output) a simple string without enclosed tag.
func (s *HTMLSnippet) RenderContent(out io.Writer) (err error) {
	if s.content != nil {
		return s.RenderChilds(out, s.content...)
	}
	return nil
}
