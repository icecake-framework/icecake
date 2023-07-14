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
	ickcore.RenderChild(out, sec,
		ick.Elem("div", `id="boxusage" class="box"`, ick.Spinner()),
		ick.Button("reset", `class="mb-5"`).
			SetId("btnreset").
			SetColor(ick.COLOR_PRIMARY).
			SetOutlined(true).
			SetDisabled(true))

	return nil
}
