package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

type SectionDocOverview struct{ SectionDocIcecake }

func (cmp *SectionDocOverview) RenderContent(out io.Writer) error {
	html.Render(out, nil, ick.Title(3, "Overview"))

	return nil
}
