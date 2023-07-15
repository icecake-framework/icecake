package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

func init() {
	ickcore.RegisterComposer("ick-input", &ICKInputField{})
}

var (
	// IconVisibilityHidden  *ICKIcon = Icon("bi bi-eye-slash-fill").SetColor(TXTCOLOR_PRIMARY)
	// IconVisibilityVisible *ICKIcon = Icon("bi bi-eye-fill").SetColor(TXTCOLOR_PRIMARY)
	BtnVisibilityHidden  *ICKButton = Button("").SetIcon(*Icon("bi bi-eye-slash-fill").SetColor(TXTCOLOR_PRIMARY), false)
	BtnVisibilityVisible *ICKButton = Button("").SetIcon(*Icon("bi bi-eye-fill").SetColor(TXTCOLOR_PRIMARY), false)
)

type INPUT_STATE string

const (
	INPUT_STD      INPUT_STATE = "" // the default input state
	INPUT_SUCCESS  INPUT_STATE = "success"
	INPUT_WARNING  INPUT_STATE = "warning"
	INPUT_ERROR    INPUT_STATE = "error"
	INPUT_LOADING  INPUT_STATE = "loading"
	INPUT_DISABLED INPUT_STATE = "disabled"
	INPUT_STATIC   INPUT_STATE = "static"
)

type ICKInputField struct {
	ickcore.BareSnippet

	// Optional label above the value
	Label string

	OpeningIcon ICKIcon // optional opening icon
	Value       string  // The input value
	IsHidden    bool    // Entered characters are hidden
	ClosingIcon ICKIcon // optional closing icon

	// Optional PlaceHolder string
	PlaceHolder string

	// Optional help text
	Help string

	// Input state: INPUT_STD, INPUT_SUCCESS, INPUT_WARNING, INPUT_ERROR, INPUT_LOADING
	State INPUT_STATE

	// ReadOnly input field
	IsReadOnly bool

	// TODO: ickui - ICKInputField add an icon to show/hide incase of hidden field
	// CanToggleVisibility enables a toggle icon button right to the input.
	// This button toggles the IsHidden status of the input field.
	// ClosingIcon is ignored if CanToggleVisibility is enabled.
	CanToggleVisibility bool
}

// Ensuring InputField implements the right interface
var _ ickcore.ContentComposer = (*ICKInputField)(nil)
var _ ickcore.TagBuilder = (*ICKInputField)(nil)

func InputField(id string, value string, placeholder string, attrs ...string) *ICKInputField {
	n := new(ICKInputField)
	n.Tag().SetId(id)
	n.Value = value
	n.PlaceHolder = placeholder
	n.Tag().ParseAttributes(attrs...)
	return n
}

func (in *ICKInputField) SetLabel(lbl string) *ICKInputField {
	in.Label = lbl
	return in
}

func (in *ICKInputField) SetHelp(help string) *ICKInputField {
	in.Help = help
	return in
}

func (in *ICKInputField) SetReadOnly(ro bool) *ICKInputField {
	in.IsReadOnly = ro
	return in
}

func (in *ICKInputField) SetHidden(h bool) *ICKInputField {
	in.IsHidden = h
	return in
}
func (in *ICKInputField) SetState(st INPUT_STATE) *ICKInputField {
	in.State = st
	return in
}
func (in *ICKInputField) SetDisabled(f bool) *ICKInputField {
	if f {
		in.State = INPUT_DISABLED
	} else {
		in.State = INPUT_STD
	}
	in.Tag().SetDisabled(f)
	return in
}
func (in *ICKInputField) SetIcon(icon ICKIcon, closing bool) *ICKInputField {
	if closing {
		in.ClosingIcon = icon
	} else {
		in.OpeningIcon = icon
	}
	return in
}

func (in *ICKInputField) SetCanToggleVisibility(can bool) *ICKInputField {
	in.CanToggleVisibility = can
	return in
}

/******************************************************************************/

// BuildTag builds the tag used to render the html element.
func (in *ICKInputField) BuildTag() ickcore.Tag {
	in.Tag().SetTagName("div").
		AddClass("field")
	return *in.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (in *ICKInputField) RenderContent(out io.Writer) error {

	// <label>
	if in.Label != "" {
		sublbl := Elem("label", `class="label mb-1"`, ickcore.ToHTML(in.Label))
		sublbl.SetId(in.Tag().SubId("label"))
		ickcore.RenderChild(out, in, sublbl)
	}

	// <div control> input field with icons
	ickcore.RenderString(out, `<span class="field has-addons mb-0">`)

	subincontrol := Elem("div", `class="control is-expanded"`)
	subincontrol.Tag().
		SetId(in.Tag().SubId("inputcontrol")).
		SetClassIf(in.State == INPUT_LOADING, "is-loading").
		SetClassIf(in.OpeningIcon.NeedRendering(), "has-icons-left").
		SetClassIf(in.ClosingIcon.NeedRendering(), "has-icons-right")

	// opening icon
	if in.OpeningIcon.NeedRendering() {
		in.OpeningIcon.Tag().AddClass("is-left").SetId(in.Tag().SubId("oicon"))
		subincontrol.Append(&in.OpeningIcon)
	}

	// <input>
	subinput := Elem("input", `class="input is-fullwidth"`)
	subinput.Tag().
		SetId(in.Tag().SubId("input")).
		SetAttributeIf(!in.IsHidden, "type", "text", "password").
		SetAttributeIf(in.Value != "", "value", in.Value).
		SetAttributeIf(in.PlaceHolder != "", "placeholder", in.PlaceHolder).
		SetAttributeIf(in.IsReadOnly, "readonly", "").
		SetAttributeIf(in.State == INPUT_DISABLED, "disabled", "").
		AddClassIf(in.State == INPUT_STATIC, "is-static", "").
		SetClassIf(in.State == INPUT_SUCCESS, "is-success").
		SetClassIf(in.State == INPUT_WARNING, "is-warning").
		SetClassIf(in.State == INPUT_ERROR, "is-danger")
	subincontrol.Append(subinput)

	// closing icon
	if in.ClosingIcon.NeedRendering() {
		in.ClosingIcon.Tag().AddClass("is-right").SetId(in.Tag().SubId("cicon"))
		subincontrol.Append(&in.ClosingIcon)
	}

	ickcore.RenderChild(out, in, subincontrol)

	// <div control> button toggle visibility
	if in.CanToggleVisibility {
		subtogglecontrol := Elem("div", `class="control"`).SetId(in.Tag().SubId("togglecontrol"))
		var btn *ICKButton
		if in.IsHidden {
			btn = BtnVisibilityVisible.Clone()
		} else {
			btn = BtnVisibilityHidden.Clone()
		}
		btn.Tag().AddClass("is-right").SetId(in.Tag().SubId("btntoggvis"))
		subtogglecontrol.Append(btn)
		ickcore.RenderChild(out, in, subtogglecontrol)
	}

	ickcore.RenderString(out, `</span>`)

	// <p help>
	if in.Help != "" {
		subhelp := Elem("p", `class="help"`, ickcore.ToHTML(in.Help))
		subhelp.Tag().
			SetId(in.Tag().SubId("help")).
			SetClassIf(in.State == INPUT_SUCCESS, "is-success").
			SetClassIf(in.State == INPUT_WARNING, "is-warning").
			SetClassIf(in.State == INPUT_ERROR, "is-danger")
		ickcore.RenderChild(out, in, subhelp)
	}

	return nil
}
