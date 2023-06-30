package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionBulmaMessage struct{ SectionIcecakeDoc }

func (cmp *SectionBulmaMessage) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Message</h2>
	<p>welcome</p>`)

	return nil
}
