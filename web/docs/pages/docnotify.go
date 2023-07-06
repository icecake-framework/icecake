package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionDocNotify struct{ SectionDocIcecake }

func (cmp *SectionDocNotify) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Notify</h2>`)

	return nil
}
