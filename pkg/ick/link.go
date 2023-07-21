package ick

import (
	"io"
	"net/url"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

// ICKLink represents an HTML anchor link.
// It is part of the core icecake snippets.
type ICKLink struct {
	ickcore.BareSnippet

	// HRef defines the associated url link.
	// if nil the <a> tag is rendered without href attribute.
	// Usually HRef is created calling TryParseHRef
	HRef *url.URL

	Body ickcore.ContentStack // HTML Element body. A stack of content composers to render.
}

// Ensuring ICKLink implements the right interface
var _ ickcore.ContentComposer = (*ICKLink)(nil)
var _ ickcore.TagBuilder = (*ICKLink)(nil)

// Link returns an HTML anchor link
func Link(child ickcore.ContentComposer, attrlist ...string) *ICKLink {
	lnk := new(ICKLink)
	lnk.Tag().ParseAttributes(attrlist...)
	lnk.Body.Push(child)
	return lnk
}

// ParseHRef tries to parse rawUrl to HRef ignoring error.
func (lnk *ICKLink) ParseHRef(rawUrl string) *ICKLink {
	lnk.HRef, _ = url.Parse(rawUrl)
	return lnk
}

// SetHRef sets the href url
func (lnk *ICKLink) SetHRef(href *url.URL) *ICKLink {
	if href == nil {
		lnk.HRef = nil
	} else {
		h := *href
		lnk.HRef = &h
	}
	return lnk
}

/******************************************************************************/

// BuildTag builds the tag used to render the html element.
func (lnk *ICKLink) BuildTag() ickcore.Tag {
	lnk.Tag().SetTagName("a")
	if lnk.HRef != nil {
		lnk.Tag().SetAttribute("href", lnk.HRef.String())
	}
	return *lnk.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// The default implementation for an HTMLSnippet snippet is to render all the internal stack of composers inside an enclosed HTML tag.
func (lnk *ICKLink) RenderContent(out io.Writer) error {
	return lnk.Body.RenderStack(out, lnk)
}
