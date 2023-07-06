package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

type SectionDocMenu struct{ SectionDocIcecake }

func (cmp *SectionDocMenu) RenderContent(out io.Writer) error {
	html.Render(out, nil, ick.Title(3, "Menu"))

	return nil
}
