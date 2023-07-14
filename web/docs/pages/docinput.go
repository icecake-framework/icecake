package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

type SectionDocInput struct{ SectionDocIcecake }

func (sec *SectionDocInput) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "InputField", "input.go", "ICKInput")

	// usages
	ickcore.RenderString(out, `<div class="box">`)
	ickcore.RenderChild(out, sec,
		ick.InputField())
	ickcore.RenderString(out, `</div>`)

	return nil
}
