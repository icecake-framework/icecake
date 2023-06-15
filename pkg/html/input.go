package html

import "io"

func init() {
	RegisterComposer("ick-input", &InputField{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
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
	HTMLSnippet

	Label       HTMLString // Optional
	PlaceHolder string     // Optional
	Help        HTMLString // Optional

	// The input value
	Value string

	// icon left

	IsRounded bool // Rounded style

	State INPUT_STATE
}

// Ensure inputfield implements HTMLComposer interface
var _ HTMLComposer = (*InputField)(nil)

func (inputfield *InputField) Tag() *Tag {
	inputfield.tag.SetName("div")
	inputfield.tag.attrs.SetClasses("field")
	return &inputfield.tag
}

func (inputfield *InputField) WriteBody(out io.Writer) error {
	// <label>
	inputfield.WriteChildSnippetIf(inputfield.Label != "", out, NewSnippet("label", TryParseAttributes(`class="label"`), inputfield.Label))

	// <input>
	subinput := NewSnippet("input", nil, "")
	inputattr := subinput.Tag().Attributes().
		SetClasses("input").
		SetClassesIf(inputfield.IsRounded, "is-rounded").
		SetAttribute("type", "text", true).
		SetAttributeIf(inputfield.PlaceHolder != "", "placeholder", inputfield.PlaceHolder, true)
	switch inputfield.State {
	case "success":
		inputattr.SetClasses("is-success")
	case "warning":
		inputattr.SetClasses("is-warning")
	case "error":
		inputattr.SetClasses("is-danger")
	case "readonly":
		inputattr.SetAttribute("readonly", "", true)
	}
	inputattr.SetAttributeIf(inputfield.Value != "", "value", inputfield.Value, true)

	// <div control>
	subcontrol := NewSnippet("div", TryParseAttributes("class=control"), "")
	// subcontrol.SetClassesIf(inputfield.State == INPUT_LOADING, "is-loading")

	// inputfield.subcontrol.SetBodyTemplate(inputfield.RenderChildSnippet(&inputfield.subinput))

	inputfield.WriteChildSnippet(out, subcontrol)

	// <p help>
	if inputfield.Help != "" {
		subhelp := NewSnippet("p", nil, inputfield.Help)
		helpattr := subinput.Tag().Attributes().SetClasses("help")
		switch inputfield.State {
		case "success":
			helpattr.SetClasses("is-success")
		case "warning":
			helpattr.SetClasses("is-warning")
		case "error":
			helpattr.SetClasses("is-danger")
		}
		inputfield.WriteChildSnippet(out, subhelp)
	}

	return nil
}
