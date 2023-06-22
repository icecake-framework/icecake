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

// BuildTag builds the tag used to render the html element.
func (inputfield *InputField) BuildTag(tag *html.Tag) {
	tag.SetTagName("div").AddClasses("field")
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (cmp *InputField) RenderContent(out io.Writer) error {
	// <label>
	html.RenderHTMLIf(!cmp.Label.IsEmpty(), out, cmp, html.HTML(`<label class="label">`), cmp.Label, html.HTML(`</label>`))

	// <div control>
	subcontrol := html.NewSnippet("div", `class="control"`)
	subcontrol.Tag().AddClassesIf(cmp.State == INPUT_LOADING, "is-loading")

	// <input>
	subinput := html.NewSnippet("input", `class="input" type="text"`)
	subinput.Tag().
		AddClassesIf(cmp.IsRounded, "is-rounded").
		SetAttributeIf(cmp.Value != "", "value", cmp.Value).
		SetAttributeIf(cmp.PlaceHolder != "", "placeholder", cmp.PlaceHolder)
	switch cmp.State {
	case "success":
		subinput.Tag().AddClasses("is-success")
	case "warning":
		subinput.Tag().AddClasses("is-warning")
	case "error":
		subinput.Tag().AddClasses("is-danger")
	case "readonly":
		subinput.Tag().SetBool("readonly", true)
	}
	subcontrol.StackContent(subinput)
	cmp.RenderChildSnippet(out, subcontrol)

	// <p help>
	if !cmp.Help.IsEmpty() {
		subhelp := html.NewSnippet("p", `class="help"`).InsertHTML(cmp.Help)
		subhelp.Tag().
			AddClassesIf(cmp.State == INPUT_SUCCESS, "is-success").
			AddClassesIf(cmp.State == INPUT_WARNING, "is-warning").
			AddClassesIf(cmp.State == INPUT_ERROR, "is-danger")
		cmp.RenderChildSnippet(out, subhelp)
	}

	return nil
}
