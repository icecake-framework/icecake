package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionDocOverview struct{ SectionDocIcecake }

func (cmp *SectionDocOverview) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Overview</h2>
	<p>welcome</p>`)

	return nil
}
