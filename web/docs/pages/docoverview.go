package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

type SectionDocOverview struct{ SectionDocIcecake }

func (sec *SectionDocOverview) RenderContent(out io.Writer) error {
	ickcore.RenderChild(out, sec, ick.Title(3, "Overview"))

	return nil
}
