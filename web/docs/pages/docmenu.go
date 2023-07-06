package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionDocMenu struct{ SectionDocIcecake }

func (cmp *SectionDocMenu) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Menu</h2>`)

	return nil
}
