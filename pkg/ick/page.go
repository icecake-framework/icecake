package ick

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/icecake-framework/icecake/pkg/ickcore"
	"github.com/lolorenzo777/verbose"
)

type HeadItem struct {
	// TODO: HeadItem base on baresnippet
	ICKElem
}

func NewHeadItem(tagname string) *HeadItem {
	item := new(HeadItem)
	item.Tag().SetTagName(tagname)
	return item
}

// An HTML5 file with its content.
type Page struct {
	meta ickcore.RMetaData // Rendering MetaData.

	Lang        string     // the html "lang" value.
	Title       string     // the html "head/title" value.
	Description string     // the html "head/meta description" value.
	HeadItems   []HeadItem // the list of tags in the section <head>

	body ICKElem // The tagname is forced to "body" during rendering.

	url  *url.URL // relative url of the html page.
	wasm *url.URL // relative url of the html page.
}

// NewPage is the Page factory, seeting up the lang for the doctype tag, and the url of the page.
// Return nil if unable to parse the url of the page.
func NewPage(lang string, rawUrl string) *Page {
	pg := new(Page)
	pg.Lang = lang
	pg.HeadItems = make([]HeadItem, 0)
	err := pg.ParseURL(rawUrl)
	if err != nil {
		verbose.Error("NewPage", err)
		return nil
	}
	return pg
}

// Meta provides a reference to the RenderingMeta object associated with this composer.
// This is required by the icecake rendering process.
func (pg *Page) RMeta() *ickcore.RMetaData {
	return &pg.meta
}

// WasmScript returns the wasm script to be added ad the end of the page to enable loading of a wasm code
func (pg *Page) WasmScript() *ickcore.HTMLString {
	if pg.wasm == nil || pg.wasm.Path == "" {
		return nil
	}
	s := ickcore.ToHTML(`<script src="/assets/wasm_exec.js"></script>
		<script>
			const goWasm = new Go()
			WebAssembly.instantiateStreaming(fetch("` + pg.wasm.Path + `"), goWasm.importObject)
				.then((result) => {
					goWasm.run(result.instance)
				})
		</script>`)
	return s
}

// ParseURL parses rawHTMLUrl to the URL of the page. The page URL stays nil in case of error.
// Only the relative path will be used.
// The path extention must be html or nothing, otherwise fails.
func (pg *Page) ParseURL(rawHTMLUrl string) (err error) {
	pg.url, err = url.Parse(rawHTMLUrl)
	if err == nil {
		relpath := pg.url.Path
		extpos := strings.LastIndex(relpath, ".")
		if extpos >= 0 {
			ext := relpath[extpos:]
			if ext != ".html" {
				return ErrBadHtmlFileExtention
			}
			relpath = relpath[:extpos]
		}
		if relpath != "" {
			pg.url.Path += ".html"
			pg.wasm, _ = url.Parse(relpath + ".wasm")
		}
	}
	return
}

// Body returns the HTMLSnippet used to render the body tag.
// Attributes can be setup. The tag will be forced to body during rendering.
func (pg *Page) Body() *ICKElem {
	return &pg.body
}

// RelURL returns the relative URL of the page, excluding the query and the fragments if any.
func (pg Page) RelURL() *url.URL {
	if pg.url == nil {
		return &url.URL{}
	}
	return &url.URL{Path: pg.url.Path}
}

// AddHeadItem add a line in the <head> section of the HtmlFile
func (f *Page) AddHeadItem(tagname string, attributes string) *Page {
	item := NewHeadItem(tagname)
	item.Tag().ParseAttributes(attributes)
	item.Tag().NoName = true
	f.HeadItems = append(f.HeadItems, *item)
	return f
}

// WriteHTMLFile creates or overwrites the file with htmlfilename name, adding the html extension if missing,
// and feeds it with the HtmlFile content including the header and the body.
// If path is provided, htmlfilename is joint to make an absolute path,
// otherwise htmlfilename is used in the current dir (unless it contains itself an absolute path).
// returns an error if ther's no filename
func (pg Page) WriteFile(outputpath string) (err error) {

	relhtmlfile := pg.url.Path

	// make a valid file name with htmlfilename
	if relhtmlfile == "" {
		return fmt.Errorf("WriteFile: %w", ErrMissingFileName)
	}

	ext := filepath.Ext(relhtmlfile)
	if ext != ".html" {
		relhtmlfile = filepath.Join(relhtmlfile, ".html")
	}
	var absfilename string
	if outputpath != "" {
		absfilename = filepath.Join(outputpath, relhtmlfile)
	} else {
		if absfilename, err = filepath.Abs(relhtmlfile); err != nil {
			return err
		}
	}

	// write it to the disk
	f, erro := os.OpenFile(absfilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if erro != nil {
		return erro
	}
	defer func() {
		if err1 := f.Close(); err1 != nil && err == nil {
			err = err1
		}
		verbose.Error(fmt.Sprintf("WriteFile %s", outputpath), err)
		if err == nil {
			verbose.Println(verbose.INFO, absfilename, "successfully written\n")
		}
	}()

	err = pg.RenderContent(f)

	return err
}

// Render turns HtmlFile into a valid HTML syntax and write it to the output stream.
// Declared required CSS files and styles are automatically added.
func (pg *Page) RenderContent(out io.Writer) (err error) {

	// <!doctype>
	ickcore.RenderString(out, `<!doctype html><html lang="`, pg.Lang, `">`)

	// <head>
	ickcore.RenderString(out, `<head>`)
	ickcore.RenderStringIf(pg.Title != "", out, "<title>", pg.Title, "</title>")
	ickcore.RenderStringIf(pg.Description != "", out, `<meta name="description" content="`+pg.Description+`">`)
	for _, item := range pg.HeadItems {
		if err = ickcore.RenderChild(out, nil, &item); err != nil {
			return err
		}
	}

	// required css files, checking for duplicate
	rcssfs := ickcore.RequiredCSSFile()
	for _, rcssf := range rcssfs {
		strrcssf := rcssf.String()
		duplicate := false
		for _, hi := range pg.HeadItems {
			if hitn, _ := hi.Tag().TagName(); hitn == "link" {
				href, found := hi.Tag().Attribute("href")
				if found && href == strrcssf {
					duplicate = true
					break
				}
			}
		}
		ickcore.RenderStringIf(!duplicate, out, `<link rel="stylesheet" href="`+strrcssf+`">`)
	}

	// required css styles
	rcssstyle := ickcore.RequiredCSSStyle()
	ickcore.RenderStringIf(rcssstyle != "", out, `<style>`, rcssstyle, `</style>`)

	ickcore.RenderString(out, "</head>")

	// <body>
	pg.body.Tag().SetTagName("body")
	ickcore.RenderChild(out, pg, &pg.body)

	// wasm script, if any
	// must be loaded at the end of the page because the wasm code interacts with the loading/loaded )DOM
	ickcore.RenderChild(out, nil, pg.WasmScript())

	// <closing>
	ickcore.RenderString(out, "</html>")

	return nil
}
