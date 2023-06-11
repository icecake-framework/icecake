package html

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// The <meta> tag embedded into the section <head> of the document.
type HtmlHeadItem struct{ HTMLSnippet }

func NewHeadItem(_tagname String, _tagselfclosing bool) *HtmlHeadItem {
	item := new(HtmlHeadItem)
	item.TagName = _tagname
	item.TagSelfClosing = _tagselfclosing
	return item
}

// An HTML5 file with its content.
type Html5Page struct {
	Lang      string         // the html "lang" value.
	Title     String         // the html "head/title" value.
	HeadItems []HtmlHeadItem // the list of tags in the section <head>
	Body      String         // the html string embedded into the <body> tag
}

// HtmlFile factory
func NewHtml5Page(_lang string) *Html5Page {
	f := new(Html5Page)
	f.Lang = _lang
	f.HeadItems = make([]HtmlHeadItem, 0)
	f.AddMetaCharset("utf-8")
	return f
}

func (f *Html5Page) AddMetaCharset(_value string) *Html5Page {
	meta := NewHeadItem("meta", true)
	meta.SetAttribute("charset", _value)
	f.HeadItems = append(f.HeadItems, *meta)
	return f
}

func (f *Html5Page) AddMetaHttpEquiv(_name string, _contentvalue string) *Html5Page {
	meta := NewHeadItem("meta", true)
	meta.SetAttribute("http-equiv", _name)
	meta.SetAttribute("content", _contentvalue)
	f.HeadItems = append(f.HeadItems, *meta)
	return f
}

func (f *Html5Page) AddMetaName(_name string, _contentvalue string) *Html5Page {
	meta := NewHeadItem("meta", true)
	meta.SetAttribute("name", _name)
	meta.SetAttribute("content", _contentvalue)
	f.HeadItems = append(f.HeadItems, *meta)
	return f
}

// HACK: better to use html.ParseAttributes
func (f *Html5Page) AddMetaLink(_attrs map[string]string) *Html5Page {
	link := NewHeadItem("link", true)
	for k, v := range _attrs {
		link.SetAttribute(k, v)
	}
	f.HeadItems = append(f.HeadItems, *link)
	return f
}

func (f *Html5Page) AddMetaScript(_src string) *Html5Page {
	script := NewHeadItem("script", false)
	script.SetAttribute("src", _src)
	f.HeadItems = append(f.HeadItems, *script)
	return f
}

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
	hf.Generate(out)

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
func (hf Html5Page) Generate(_out io.Writer) (_err error) {

	var htmltitle String
	if hf.Title != "" {
		htmltitle = "<title>" + hf.Title + "</title>"
	}

	_, _err = fmt.Fprintf(_out, `<!doctype html><html lang="%s"><head>%s`, hf.Lang, htmltitle)
	if _err != nil {
		return _err
	}

	for _, headitem := range hf.HeadItems {
		_, _err = writeHTMLSnippet(_out, &headitem, nil, false, -1)
		if _err != nil {
			return _err
		}
	}

	io.WriteString(_out, "</head><body>")

	_, _err = UnfoldHTML(_out, hf.Body, nil)
	if _err != nil {
		return _err
	}
	io.WriteString(_out, "</body></html>")

	return nil
}
