package pgbulmamenu

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type Section struct{ html.HTMLSnippet }

func (cmp *Section) BuildTag(tag *html.Tag) { tag.SetTagName("section").AddClasses("content py-3") }

func (cmp *Section) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Overview</h2>
	<p>welcome</p>`)

	return nil
}
