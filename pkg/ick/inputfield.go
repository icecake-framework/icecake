package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

func init() {
	ickcore.RegisterComposer("ick-input", &ICKInputField{})
}

type INPUT_STATE string

const (
	INPUT_NORMAL   INPUT_STATE = "normal" // the default input state
	INPUT_SUCCESS  INPUT_STATE = "success"
	INPUT_WARNING  INPUT_STATE = "warning"
	INPUT_ERROR    INPUT_STATE = "error"
	INPUT_LOADING  INPUT_STATE = "loading"
	INPUT_READONLY INPUT_STATE = "readlonly"
)

type ICKInputField struct {
	ickcore.BareSnippet

	Label       ickcore.HTMLString // Optional
	PlaceHolder string             // Optional
	Help        ickcore.HTMLString // Optional

	// The input value
	Value string

	State INPUT_STATE
}

// Ensuring InputField implements the right interface
var _ ickcore.ContentComposer = (*ICKInputField)(nil)
var _ ickcore.TagBuilder = (*ICKInputField)(nil)

func InputField() *ICKInputField {
	n := new(ICKInputField)
	return n
}

/******************************************************************************/

func (inputfield *ICKInputField) NeedRendering() bool {
	return true
}

// BuildTag builds the tag used to render the html element.
func (inputfield *ICKInputField) BuildTag() ickcore.Tag {
	inputfield.Tag().SetTagName("div").AddClass("field")
	return *inputfield.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (cmp *ICKInputField) RenderContent(out io.Writer) error {
	// <label>
	if cmp.Label.NeedRendering() {
		ickcore.RenderString(out, `<label class="label">`)
		ickcore.RenderChild(out, cmp, &cmp.Label)
		ickcore.RenderString(out, `</label>`)
	}

	// <input>
	subinput := Elem("input", `class="input" type="text"`)
	subinput.Tag().
		SetAttributeIf(cmp.Value != "", "value", cmp.Value).
		SetAttributeIf(cmp.PlaceHolder != "", "placeholder", cmp.PlaceHolder)
	switch cmp.State {
	case "success":
		subinput.Tag().AddClass("is-success")
	case "warning":
		subinput.Tag().AddClass("is-warning")
	case "error":
		subinput.Tag().AddClass("is-danger")
	case "readonly":
		subinput.Tag().SetBool("readonly", true)
	}

	// <div control>
	subcontrol := Elem("div", `class="control"`).Append(subinput)
	subcontrol.Tag().SetClassIf(cmp.State == INPUT_LOADING, "is-loading")
	ickcore.RenderChild(out, cmp, subcontrol)

	// <p help>
	if cmp.Help.NeedRendering() {
		subhelp := Elem("p", `class="help"`, &cmp.Help)
		subhelp.Tag().
			SetClassIf(cmp.State == INPUT_SUCCESS, "is-success").
			SetClassIf(cmp.State == INPUT_WARNING, "is-warning").
			SetClassIf(cmp.State == INPUT_ERROR, "is-danger")
		ickcore.RenderChild(out, cmp, subhelp)
	}

	return nil
}
