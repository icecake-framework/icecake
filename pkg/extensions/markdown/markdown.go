package markdown

import (
	"bytes"

	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/yuin/goldmark"
)

// RenderMarkdown process _mdtxt markdown source file and convert it to an HTML string,
// then use it as an HTML template to render it with data and components.
//
// Returns an error if the markdown processor fails.
func RenderIn(_elem *dom.Element, _mdtxt string, _options ...goldmark.Option) error {
	if !_elem.IsDefined() {
		return nil
	}
	md := goldmark.New(_options...)
	var buf bytes.Buffer
	err := md.Convert([]byte(_mdtxt), &buf)
	if err != nil {
		console.Warnf("RenderMarkdown has error: %s", err.Error())
		return err
	}
	_elem.InsertSnippet(dom.INSERT_BODY, html.ToHTML(buf.String()))
	return nil
}
