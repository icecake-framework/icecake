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
type HtmlPage struct {
	Lang        string       // the html "lang" value.
	Title       string       // the html "head/title" value.
	Description string       // the html "head/meta description" value.
	HeadItems   []*HeadItem  // the list of tags in the section <head>
	Body        HTMLComposer // the html snippet used during the content rendering.

	HTMLFileName string // the relative file name of this page. Should finish with the .html extension.
}

// NewPage is the HtmlPage factory, seeting up the lang for the doctype tag.
func NewPage(_lang string) *HtmlPage {
	f := new(HtmlPage)
	f.Lang = _lang
	f.HeadItems = make([]*HeadItem, 0)
	return f
}

// AddHeadItem add a line in the <head> section of the page
func (f *HtmlPage) AddHeadItem(tagname string, attributes string) *HtmlPage {
	item := NewHeadItem(tagname)
	item.Tag().ParseAttributes(attributes)
	f.HeadItems = append(f.HeadItems, item)
	return f
}

// WriteHTMLFile creates or overwrites the file with htmlfilename name, adding the html extension,
// and feeds it with the page content including the dictypen the header and the body.
// If path is provided, htmlfilename is joint to make an absolute path,
// otherwise htmlfilename is used in the current dir (unless it contains itself an absolute path).
// returns an error if ther's no filename
func (page *HtmlPage) WriteHTMLFile(path string) (err error) {

	// make a valid file name with htmlfilename
	htmlfilename := page.HTMLFileName
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

	err = page.Render(f)

	return err
}

// Render turns HtmlPage into a valid HTML syntax and write it to the output stream.
func (page *HtmlPage) Render(out io.Writer) (err error) {

	WriteStrings(out, `<!doctype html><html lang="`, page.Lang, `">`)

	WriteStrings(out, `<head>`)
	WriteStringsIf(page.Title != "", out, "<title>", page.Title, "</title>")
	for _, headitem := range page.HeadItems {
		if err = RenderSnippet(out, nil, headitem); err != nil {
			return err
		}
	}
	WriteString(out, "</head>")

	if page.Body != nil {
		page.Body.Tag().SetTagName("body")
		err = RenderSnippet(out, nil, page.Body)
		if err != nil {
			return err
		}
	}

	WriteString(out, "</html>")

	return nil
}
