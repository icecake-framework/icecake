package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionBulmaCard struct{ SectionIcecakeDoc }

func (cmp *SectionBulmaCard) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Card</h2>
	<p>welcome</p>`)

	return nil
}
