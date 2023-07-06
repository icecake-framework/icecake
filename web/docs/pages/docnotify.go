package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

type SectionDocNotify struct{ SectionDocIcecake }

func (cmp *SectionDocNotify) RenderContent(out io.Writer) error {
	html.Render(out, nil, ick.Title(3, "Notify"))

	return nil
}
