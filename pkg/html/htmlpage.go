package html

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// An HTML5 file with its content.
type Html5Page struct {
	Lang      string         // the html "lang" value.
	Title     HTMLString     // the html "head/title" value.
	HeadItems []*HTMLSnippet // the list of tags in the section <head>
	Body      HTMLString     // the html string embedded into the <body> of the page
}

// HtmlFile factory
func NewHtml5Page(_lang string) *Html5Page {
	f := new(Html5Page)
	f.Lang = _lang
	f.HeadItems = make([]*HTMLSnippet, 0)
	f.AddHeadMeta("charset=UTF-8")
	return f
}

// AddMeta add a meta tag
func (f *Html5Page) AddHeadMeta(attributes string) *Html5Page {
	meta := NewSnippet("meta", TryParseAttributes(attributes), "")
	f.HeadItems = append(f.HeadItems, meta)
	return f
}

func (f *Html5Page) AddHeadLink(attributes string) *Html5Page {
	link := NewSnippet("link", TryParseAttributes(attributes), "")
	f.HeadItems = append(f.HeadItems, link)
	return f
}

func (f *Html5Page) AddHeadScript(attributes string) *Html5Page {
	script := NewSnippet("script", TryParseAttributes(attributes), "")
	f.HeadItems = append(f.HeadItems, script)
	return f
}

//TODO: AddHeadStyle

// WriteHTMLFile creates or overwrites the file with _htmlfilename name, adding the html extension,
// and feeds it with f HTML content.
// If _path is provided, _htmlfilename is joint to it to make an absolute path.
// Otherwise _htmlfilename is used in the current dir unless it contains an absolute path.
func (hf Html5Page) WriteHTMLFile(_path string, htmlfilename string) (_err error) {

	// make a valid file name with _htmlfilename
	ext := filepath.Ext(htmlfilename)
	if ext != ".html" {
		htmlfilename = filepath.Join(htmlfilename, ".html")
	}
	var absfilename string
	if _path != "" {
		absfilename = filepath.Join(_path, htmlfilename)
	} else {
		absfilename, _err = filepath.Abs(htmlfilename)
		if _err != nil {
			return
		}
	}

	// generate the document
	out := new(bytes.Buffer)
	err := hf.Generate(out)
	if err != nil {
		return err
	}

	// write it to the disk
	f, err := os.OpenFile(absfilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func() {
		if err1 := f.Close(); err1 != nil && _err == nil {
			_err = err1
		}
		if _err != nil {
			fmt.Println("WriteHTMLFile fails:", _err)
		} else {
			fmt.Println(absfilename, "succesfully generated")
		}
	}()
	_, _err = f.Write(out.Bytes())
	return _err
}

// Generate turns f Html5File into a valid HTML syntax and write it to _out
func (page Html5Page) Generate(out io.Writer) (err error) {

	var htmltitle HTMLString
	if page.Title != "" {
		htmltitle = "<title>" + page.Title + "</title>"
	}

	_, err = fmt.Fprintf(out, `<!doctype html><html lang="%s"><head>%s`, page.Lang, htmltitle)
	if err != nil {
		return err
	}

	for _, headitem := range page.HeadItems {
		_, err = writeSnippet(out, headitem, nil, true, "", 0, 0)
		if err != nil {
			return err
		}
	}

	WriteString(out, "</head><body>")

	_, err = UnfoldHTML(out, page.Body, nil)
	if err != nil {
		return err
	}
	io.WriteString(out, "</body></html>")

	return nil
}
