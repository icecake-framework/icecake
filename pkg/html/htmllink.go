package html

import (
	"net/url"
)

// Link represents an HTML anchor link.
// It is part of the core icecake snippets.
type Link struct {
	HTMLSnippet

	// HRef defines the associated url link.
	// if nil the <a> tag is rendered without href attribute.
	// Usually HRef is created calling TryParseHRef
	HRef *url.URL
}

// Ensure HTMLString implements HTMLTagComposer interface
var _ HTMLTagComposer = (*Link)(nil)

// A returns an HTML anchor link
func A(attrlist ...string) *Link {
	lnk := new(Link)
	lnk.Tag().SetTagName("a")
	lnk.Tag().ParseAttributes(attrlist...)
	return lnk
}

// ParseHRef tries to parse rawUrl to HRef ignoring error.
func (lnk *Link) ParseHRef(rawUrl string) *Link {
	lnk.HRef, _ = url.Parse(rawUrl)
	return lnk
}

// SetHRef sets the href url
func (lnk *Link) SetHRef(href *url.URL) *Link {
	if href == nil {
		lnk.HRef = nil
	} else {
		h := *href
		lnk.HRef = &h
	}
	return lnk
}

// BuildTag builds the tag used to render the html element.
func (lnk *Link) BuildTag() Tag {
	if lnk.HRef != nil {
		lnk.Tag().SetAttribute("href", lnk.HRef.String())
	}
	return *lnk.Tag()
}

// SetBody adds one or many HTMLComposer to the rendering stack of this composer.
// Returns the snippet to allow chaining calls.
func (lnk *Link) SetBody(content ...HTMLComposer) *Link {
	lnk.HTMLSnippet.AddContent(content...)
	return lnk
}
