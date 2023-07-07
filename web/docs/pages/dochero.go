package webdocs

import (
	"io"
)

type SectionDocHero struct{ SectionDocIcecake }

func (sec *SectionDocHero) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Hero", "hero.go", "ICKHero")

	return nil
}
