package html

import (
	"net/url"
)

// ICKLink represents an HTML anchor link.
// It is part of the core icecake snippets.
type ICKLink struct {
	HTMLSnippet

	// HRef defines the associated url link.
	// if nil the <a> tag is rendered without href attribute.
	// Usually HRef is created calling TryParseHRef
	HRef *url.URL
}

// Ensure HTMLString implements HTMLTagComposer interface
var _ HTMLTagComposer = (*ICKLink)(nil)

// A returns an HTML anchor link
func A(attrlist ...string) *ICKLink {
	lnk := new(ICKLink)
	lnk.Tag().SetTagName("a").ParseAttributes(attrlist...)
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

// BuildTag builds the tag used to render the html element.
func (lnk *ICKLink) BuildTag() Tag {
	if lnk.HRef != nil {
		lnk.Tag().SetAttribute("href", lnk.HRef.String())
	}
	return *lnk.Tag()
}

// // SetBody adds one or many HTMLComposer to the rendering stack of this composer.
// // Returns the snippet to allow chaining calls.
// func (lnk *Link) SetBody(content ...HTMLComposer) *Link {
// 	lnk.HTMLSnippet.AddContent(content...)
// 	return lnk
// }
