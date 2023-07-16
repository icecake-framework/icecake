package ickcore

import (
	"io"
	"reflect"
	"strings"
)

type Composer interface {
	// Meta returns a reference to render meta data
	RMetaProvider

	// NeedRendering returns true if the composer needs rendering or has something to render.
	NeedRendering() bool
}

type TagBuilder interface {
	Composer

	// SetAttribute creates a tag attribute and set its value.
	// If the attribute already exists then it is updated.
	//
	// Attribute's name is case sensitive
	SetAttribute(name string, value string)

	// BuildTag builds the tag used to render the html element.
	// The composer rendering processes call BuildTag once.
	// If the implementer builds an empty tag, only the body will be rendered.
	//
	// The returned tag can be built from scratch or on top of an embedded tag in the snippet.
	BuildTag() Tag
}

// BuildTag get a tag from the TagBuilder then set up name attribute and RMeta id
func BuildTag(tb TagBuilder) Tag {
	tag := tb.BuildTag()

	if !tag.IsEmpty() {
		// force property name ?
		if !tag.NoName {
			cmpname := reflect.TypeOf(tb).Elem().Name()
			cmpname = strings.ToLower(cmpname)
			tag.SetAttribute("name", cmpname)
		}
		tb.RMeta().TagId = tag.Id()
	}

	return tag
}

type ContentComposer interface {
	Composer

	// RenderContent writes the HTML string corresponding to the content of the HTML element.
	// Return an error to stops the rendering process.
	RenderContent(out io.Writer) error
}

type TagProvider interface {
	TagBuilder

	Tag() *Tag
}
