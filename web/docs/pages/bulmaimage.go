package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionBulmaImage struct{ SectionIcecakeDoc }

func (cmp *SectionBulmaImage) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Image</h2>
	<p>welcome</p>`)

	return nil
}
