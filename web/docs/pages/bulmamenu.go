package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionBulmaMenu struct{ SectionIcecakeDoc }

func (cmp *SectionBulmaMenu) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Menu</h2>
	<p>welcome</p>`)

	return nil
}