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

func (cmp *InputField) RenderContent(out io.Writer) error {
	// <label>
	html.RenderHTMLIf(!cmp.Label.IsEmpty(), out, cmp, html.HTML(`<label class="label">`), cmp.Label, html.HTML(`</label>`))

	// <input>
	subinput := html.NewSnippet("input", `class="input" type="text"`)
	subinput.Tag().Attributes().
		AddClassesIf(cmp.IsRounded, "is-rounded").
		SetAttributeIf(cmp.PlaceHolder != "", "placeholder", cmp.PlaceHolder)
	switch cmp.State {
	case "success":
		subinput.Tag().Attributes().AddClasses("is-success")
	case "warning":
		subinput.Tag().Attributes().AddClasses("is-warning")
	case "error":
		subinput.Tag().Attributes().AddClasses("is-danger")
	case "readonly":
		subinput.Tag().Attributes().SetAttribute("readonly", "")
	}
	subinput.Tag().Attributes().SetAttributeIf(cmp.Value != "", "value", cmp.Value)

	// <div control>
	subcontrol := html.NewSnippet("div", `class="control"`)
	// subcontrol.SetClassesIf(inputfield.State == INPUT_LOADING, "is-loading")

	cmp.RenderChildSnippet(out, subcontrol)

	// <p help>
	if !cmp.Help.IsEmpty() {
		subhelp := html.NewSnippet("p", `class="help"`).AddContent(&cmp.Help)
		subhelpa := subinput.Tag().Attributes()
		switch cmp.State {
		case "success":
			subhelpa.AddClasses("is-success")
		case "warning":
			subhelpa.AddClasses("is-warning")
		case "error":
			subhelpa.AddClasses("is-danger")
		}
		cmp.RenderChildSnippet(out, subhelp)
	}

	return nil
}
