package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-input", &InputField{})
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

type InputField struct {
	html.HTMLSnippet

	Label       html.HTMLString // Optional
	PlaceHolder string          // Optional
	Help        html.HTMLString // Optional

	// The input value
	Value string

	// icon left

	IsRounded bool // Rounded style

	State INPUT_STATE
}

// Ensure inputfield implements HTMLComposer interface
var _ html.HTMLComposer = (*InputField)(nil)

// BuildTag builds the tag used to render the html element.
func (inputfield *InputField) BuildTag() html.Tag {
	inputfield.Tag().SetTagName("div").AddClass("field")
	return *inputfield.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (cmp *InputField) RenderContent(out io.Writer) error {
	// <label>
	if !cmp.Label.IsEmpty() {
		html.WriteString(out, `<label class="label">`)
		cmp.RenderChild(out, &cmp.Label)
		html.WriteString(out, `</label>`)
	}

	// <div control>
	subcontrol := html.Div(`class="control"`)
	subcontrol.Tag().SetClassIf(cmp.State == INPUT_LOADING, "is-loading")

	// <input>
	subinput := html.NewSnippet("input", `class="input" type="text"`)
	subinput.Tag().
		SetClassIf(cmp.IsRounded, "is-rounded").
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
	subcontrol.AddContent(subinput)
	cmp.RenderChild(out, subcontrol)

	// <p help>
	if !cmp.Help.IsEmpty() {
		subhelp := html.P(`class="help"`).AddContent(&cmp.Help)
		subhelp.Tag().
			SetClassIf(cmp.State == INPUT_SUCCESS, "is-success").
			SetClassIf(cmp.State == INPUT_WARNING, "is-warning").
			SetClassIf(cmp.State == INPUT_ERROR, "is-danger")
		cmp.RenderChild(out, subhelp)
	}

	return nil
}
