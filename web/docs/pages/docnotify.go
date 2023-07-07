package webdocs

import (
	"io"
)

type SectionDocNotify struct{ SectionDocIcecake }

func (sec *SectionDocNotify) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Notify", "notify.go", "ICKNotify")

	return nil
}
