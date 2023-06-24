package html

import (
	"io"
	"net/url"
	"strings"
)

// Page represents a set of stacked HTML elements associated to an url and a set of usual properties.
// Page implements HTMLcomposer interfaces and is rendered to an output stream with the icecake Rendering functions.
// It is part of the core icecake snippets.
type Page struct {
	meta  RenderingMeta  // Rendering MetaData.
	stack []HTMLComposer // HTML composers to render without enclosed tag.

	Title       string // the html <head><title> value.
	Description string // the html <head><meta name="description"> value.

	// relative url of the page.
	url *url.URL
}

// Ensure Page implements HTMLComposer interface
var _ HTMLComposer = (*Page)(nil)

func NewPage(rawHTMLUrl string) *Page {
	pg := new(Page)
	pg.stack = make([]HTMLComposer, 0)
	pg.ParseURL(rawHTMLUrl)
	return pg
}

// ParseURL parses rawHTMLUrl to the URL of the page. The page URL stays nil in case of error.
// Only the relative path will be used.
// The path extention must be html or nothing, otherwise fails.
func (pg *Page) ParseURL(rawHTMLUrl string) (err error) {
	pg.url, err = url.Parse(rawHTMLUrl)
	if err == nil {
		extpos := strings.LastIndex(pg.url.Path, ".")
		if extpos >= 0 {
			ext := pg.url.Path[extpos:]
			if ext != ".html" {
				return ErrBadHtmlFileExtention
			}
		} else if pg.url.Path != "" {
			pg.url.Path += ".html"
		}
	}
	return
}

// RelURL returns the relative URL of the page, excluding the query and the fragments if any.
func (pg Page) RelURL() *url.URL {
	if pg.url == nil {
		return &url.URL{}
	}
	return &url.URL{Path: pg.url.Path}
}

// Meta provides a reference to the RenderingMeta object associated with this composer.
// This is required by the icecake rendering process.
func (pg *Page) Meta() *RenderingMeta {
	return &pg.meta
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// The default implementation for an Page snippet is to render the internal stack of composers.
// This can be overloaded by a custom page.
func (pg *Page) RenderContent(out io.Writer) error {
	for _, cmp := range pg.stack {
		err := Render(out, pg, cmp)
		if err != nil {
			return err
		}
	}
	return nil
}

// Stack adds one or many HTMLComposer to the rendering stack of this composer.
// Returns the page to allow chaining calls.
func (pg *Page) Stack(content ...HTMLComposer) *Page {
	if pg.stack == nil {
		pg.stack = make([]HTMLComposer, 0)
	}
	pg.stack = append(pg.stack, content...)
	return pg
}