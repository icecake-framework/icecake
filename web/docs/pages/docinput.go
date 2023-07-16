package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

type SectionDocInput struct{ SectionDocIcecake }

func (sec *SectionDocInput) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "InputField", "inputfield.go", "ICKInputField")

	// usages
	ickcore.RenderChild(out, sec,
		ick.Elem("div", `id="boxusage" class="box"`, ick.Spinner()),
		ick.Button("reset", `class="mb-5"`).
			SetId("btnreset").
			SetColor(ick.COLOR_PRIMARY).
			SetOutlined(true).
			SetDisabled(true))

	// rendering
	ickcore.RenderChild(out, sec, ick.Title(4, "Rendering"))
	ickcore.RenderString(out, `<div class="box">`)
	ickcore.RenderChild(out, sec,
		ick.InputField("in-r0", "", "input"),
		ick.InputField("in-r1", "initial value", "input with initial value"),
		ick.InputField("in-r2", "readonly", "input readonly").
			SetReadOnly(true),
		ick.InputField("in-r3", "", "loading input").
			SetState(ick.INPUT_LOADING),
		ick.InputField("in-r4", "", "disabled input").
			SetState(ick.INPUT_DISABLED),
		ick.InputField("in-r5", "", "static readonly input").
			SetState(ick.INPUT_STATIC).
			SetReadOnly(true),
		ick.InputField("in-r6", "", "email").
			SetIcon(*ick.Icon("bi bi-envelope-at").SetColor(ick.TXTCOLOR_INFO_DARK), false).
			SetIcon(*ick.Icon("bi bi-info-circle").SetColor(ick.TXTCOLOR_INFO_DARK), true),
		ick.InputField("in-r7", "", "password").
			SetHidden(true),
		ick.InputField("in-r8", "", "e.g. bob").
			SetLabel("With a label").
			SetHelp("With a help text"),
		ick.InputField("in-r9", "", "in error state").
			SetHelp("there's an error in this input").
			SetState(ick.INPUT_ERROR),
		ick.InputField("in-r10", "", "in warning state").
			SetHelp("there's a warning error in this input").
			SetState(ick.INPUT_WARNING),
		ick.InputField("in-r11", "", "in success state").
			SetHelp("there's a success in this input").
			SetState(ick.INPUT_SUCCESS),
		ick.InputField("in-r12", "", "Can toggle visibility").
			SetLabel("With a label").
			SetHelp("With help").
			SetCanToggleVisibility(true))

	ickcore.RenderString(out, `</div>`)

	return nil
}
