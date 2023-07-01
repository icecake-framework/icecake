package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionBulmaDelete struct{ SectionIcecakeDoc }

func (cmp *SectionBulmaDelete) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Delete</h2>
	<p>welcome</p>`)

	return nil
}
