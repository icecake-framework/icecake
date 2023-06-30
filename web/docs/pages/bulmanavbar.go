package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionBulmaNavbar struct{ SectionIcecakeDoc }

func (cmp *SectionBulmaNavbar) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Navbar</h2>
	<p>welcome</p>`)

	return nil
}
