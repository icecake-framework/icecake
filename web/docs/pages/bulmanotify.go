package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionBulmaNotify struct{ SectionIcecakeDoc }

func (cmp *SectionBulmaNotify) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Notify</h2>
	<p>welcome</p>`)

	return nil
}
