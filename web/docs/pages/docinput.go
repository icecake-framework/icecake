package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

type SectionDocInput struct{ SectionDocIcecake }

func (sec *SectionDocInput) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "InputField", "input.go", "ICKInput")

	// usages
	html.RenderString(out, `<div class="box">`)
	html.RenderChild(out, sec,
		ick.InputField())
	html.RenderString(out, `</div>`)

	return nil
}
