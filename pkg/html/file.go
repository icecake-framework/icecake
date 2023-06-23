package html

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sunraylab/verbose"
)

type HeadItem struct {
	HTMLSnippet
}

func NewHeadItem(tagname string) *HeadItem {
	item := new(HeadItem)
	item.Tag().SetTagName(tagname)
	return item
}

// An HTML5 file with its content.
type HtmlFile struct {
	Lang        string      // the html "lang" value.
	Title       string      // the html "head/title" value.
	Description string      // the html "head/meta description" value.
	HeadItems   []*HeadItem // the list of tags in the section <head>

	Body HTMLTagComposer // the html composer used to render the body tag. The tagname must be "body" otherwise will not render.

	//HTMLFileName string // the relative file name, should finish with the .html extension.
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
func (hfile *HtmlFile) WriteHTMLFile(outputpath string, relhtmlfile string) (err error) {

	// make a valid file name with htmlfilename
	if relhtmlfile == "" {
		return fmt.Errorf("WriteHTMLFile: %w", ErrMissingFileName)
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
		verbose.Error(fmt.Sprintf("WriteHTMLFile %s", outputpath), err)
		if err == nil {
			verbose.Println(verbose.INFO, absfilename, "succesfully written")
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
	WriteStringsIf(hfile.Title != "", out, "<title>", hfile.Title, "</title>")
	if hfile.Description != "" {
		hfile.AddHeadItem("meta", `name="description" content="`+hfile.Description+`"`)
	}

	// css files
	rcssfs := RequiredCSSFile()
	for _, rcssf := range rcssfs {
		strrcssf := rcssf.String()
		duplicate := false
		for _, hi := range hfile.HeadItems {
			if hi.tag.tagname == "link" {
				href, found := hi.tag.Attribute("href")
				if found && href == strrcssf {
					duplicate = true
					break
				}
			}
		}
		if !duplicate {
			hfile.AddHeadItem("link", `rel="stylesheet" href="`+strrcssf+`"`)
		}
	}

	for _, headitem := range hfile.HeadItems {
		if err = Render(out, nil, headitem); err != nil {
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
		tg, _ := hfile.Body.Tag().TagName()
		if tg != "body" {
			err = ErrBodyTagMissing
			WriteStrings(out, "<!-- ", err.Error(), " -->")
		} else {
			err = Render(out, nil, hfile.Body)
		}
		if err != nil {
			return err
		}
	}

	// <closing>
	WriteString(out, "</html>")

	return nil
}
