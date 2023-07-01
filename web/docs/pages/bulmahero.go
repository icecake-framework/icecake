package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionBulmaHero struct{ SectionIcecakeDoc }

func (cmp *SectionBulmaHero) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Hero</h2>
	<p>welcome</p>`)

	return nil
}
