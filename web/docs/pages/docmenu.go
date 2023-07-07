package webdocs

import (
	"io"
)

type SectionDocMenu struct{ SectionDocIcecake }

func (sec *SectionDocMenu) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Menue", "menu.go", "ICKMenu")

	return nil
}
