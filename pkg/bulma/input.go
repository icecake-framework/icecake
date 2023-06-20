package bulma

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-input", &InputField{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
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

func (inputfield *InputField) BuildTag(tag *html.Tag) {
	tag.SetName("div").Attributes().AddClasses("field")
}

func (inputfield *InputField) RenderContent(out io.Writer) error {
	// <label>
	inputfield.RenderChildSnippetIf(!inputfield.Label.IsEmpty(), out, html.NewSnippet("label", html.ParseAttributes(`class="label"`)).SetContent(&inputfield.Label))

	// <input>
	subinput := html.NewSnippet("input", nil)
	inputattr := subinput.Tag().Attributes().
		AddClasses("input").
		AddClassesIf(inputfield.IsRounded, "is-rounded").
		SetAttribute("type", "text").
		SetAttributeIf(inputfield.PlaceHolder != "", "placeholder", inputfield.PlaceHolder, true)
	switch inputfield.State {
	case "success":
		inputattr.AddClasses("is-success")
	case "warning":
		inputattr.AddClasses("is-warning")
	case "error":
		inputattr.AddClasses("is-danger")
	case "readonly":
		inputattr.SetAttribute("readonly", "")
	}
	inputattr.SetAttributeIf(inputfield.Value != "", "value", inputfield.Value, true)

	// <div control>
	subcontrol := html.NewSnippet("div", html.ParseAttributes("class=control"))
	// subcontrol.SetClassesIf(inputfield.State == INPUT_LOADING, "is-loading")

	// inputfield.subcontrol.SetBodyTemplate(inputfield.RenderChildSnippet(&inputfield.subinput))

	inputfield.RenderChildSnippet(out, subcontrol)

	// <p help>
	if !inputfield.Help.IsEmpty() {
		subhelp := html.NewSnippet("p", nil).SetContent(&inputfield.Help)
		helpattr := subinput.Tag().Attributes().AddClasses("help")
		switch inputfield.State {
		case "success":
			helpattr.AddClasses("is-success")
		case "warning":
			helpattr.AddClasses("is-warning")
		case "error":
			helpattr.AddClasses("is-danger")
		}
		inputfield.RenderChildSnippet(out, subhelp)
	}

	return nil
}
