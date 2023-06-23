package html

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type HeadItem struct {
	HTMLSnippet
}

func NewHeadItem(tagname string) *HeadItem {
	item := new(HeadItem)
	item.Tag().SetTagName(tagname).SetBool("noid", true)
	return item
}

// An HTML5 file with its content.
type HtmlFile struct {
	Lang        string       // the html "lang" value.
	Title       string       // the html "head/title" value.
	Description string       // the html "head/meta description" value.
	HeadItems   []*HeadItem  // the list of tags in the section <head>
	Body        HTMLComposer // the html snippet used during the content rendering.

	HTMLFileName string // the relative file name, should finish with the .html extension.
}

// NewHtmlFile is the HtmlFile factory, seeting up the lang for the doctype tag.
func NewHtmlFile(_lang string) *HtmlFile {
	f := new(HtmlFile)
	f.Lang = _lang
	f.HeadItems = make([]*HeadItem, 0)
	return f
}

// AddHeadItem add a line in the <head> section of the HtmlFile
func (f *HtmlFile) AddHeadItem(tagname string, attributes string) *HtmlFile {
	item := NewHeadItem(tagname)
	item.Tag().ParseAttributes(attributes)
	item.Tag().NoName = true
	f.HeadItems = append(f.HeadItems, item)
	return f
}

// WriteHTMLFile creates or overwrites the file with htmlfilename name, adding the html extension if missing,
// and feeds it with the HtmlFile content including the header and the body.
// If path is provided, htmlfilename is joint to make an absolute path,
// otherwise htmlfilename is used in the current dir (unless it contains itself an absolute path).
// returns an error if ther's no filename
func (hfile *HtmlFile) WriteHTMLFile(path string) (err error) {

	// make a valid file name with htmlfilename
	htmlfilename := hfile.HTMLFileName
	if htmlfilename == "" {
		return errors.New("WriteHTMLFile: missing file name")
	}

	ext := filepath.Ext(htmlfilename)
	if ext != ".html" {
		htmlfilename = filepath.Join(htmlfilename, ".html")
	}
	var absfilename string
	if path != "" {
		absfilename = filepath.Join(path, htmlfilename)
	} else {
		if absfilename, err = filepath.Abs(htmlfilename); err != nil {
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
		if err != nil {
			fmt.Println("WriteHTMLFile fails:", err)
		} else {
			fmt.Println(absfilename, "succesfully written")
		}
	}()

	err = hfile.Render(f)

	return err
}

// Render turns HtmlFile into a valid HTML syntax and write it to the output stream.
// Declared required CSS files and styles are automatically added.
func (hfile *HtmlFile) Render(out io.Writer) (err error) {

	// <!doctype>
	WriteStrings(out, `<!doctype html><html lang="`, hfile.Lang, `">`)

	// <head>
	WriteStrings(out, `<head>`)

	// css files

	rcssfs := RequiredCSSFile()
	for _, rcssf := range rcssfs {
		hfile.AddHeadItem("link", `rel="stylesheet" href="`+rcssf.String()+`"`)
	}

	WriteStringsIf(hfile.Title != "", out, "<title>", hfile.Title, "</title>")
	for _, headitem := range hfile.HeadItems {
		if err = RenderSnippet(out, nil, headitem); err != nil {
			return err
		}
	}

	rcssstyle := RequiredCSSStyle()
	if rcssstyle != "" {
		WriteStringsIf(rcssstyle != "", out, `<style>`, rcssstyle, `</style>`)
	}

	WriteString(out, "</head>")

	// <body>
	if hfile.Body != nil {
		hfile.Body.Tag().SetTagName("body")
		err = RenderSnippet(out, nil, hfile.Body)
		if err != nil {
			return err
		}
	}

	// <closing>
	WriteString(out, "</html>")

	return nil
}
