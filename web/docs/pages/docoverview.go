package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

type SectionDocOverview struct{ SectionDocIcecake }

func (sec *SectionDocOverview) RenderContent(out io.Writer) error {
	html.RenderChild(out, sec, ick.Title(3, "Overview"))

	return nil
}
