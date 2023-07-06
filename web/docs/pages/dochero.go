package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

type SectionDocHero struct{ SectionDocIcecake }

func (cmp *SectionDocHero) RenderContent(out io.Writer) error {
	html.Render(out, nil, ick.Title(3, "Hero"))

	return nil
}
