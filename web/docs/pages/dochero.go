package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionDocHero struct{ SectionDocIcecake }

func (cmp *SectionDocHero) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Hero</h2>`)

	return nil
}
