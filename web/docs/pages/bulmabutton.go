package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionBulmaButton struct {
	SectionIcecakeDoc
}

func (cmp *SectionBulmaButton) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Button</h2>
	<p>welcome</p>`)

	return nil
}
