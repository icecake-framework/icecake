package ickui

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

func init() {
	ickcore.RegisterComposer("ick-input", &ICKInputField{})
}

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
	dom.UI

	// Optional label above the value
	Label string

	OpeningIcon ick.ICKIcon // optional opening icon
	Value       string      // The input value
	IsHidden    bool        // Entered characters are hidden
	ClosingIcon ick.ICKIcon // optional closing icon

	// Optional PlaceHolder string
	PlaceHolder string

	// Optional help text
	Help string

	// Input state: INPUT_STD, INPUT_SUCCESS, INPUT_WARNING, INPUT_ERROR, INPUT_LOADING
	State INPUT_STATE

	// ReadOnly input field
	IsReadOnly bool

	// TODO: ickui - ICKInputField add an icon to show/hide incase of hidden field
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
func (in *ICKInputField) SetIcon(icon ick.ICKIcon, closing bool) *ICKInputField {
	if closing {
		in.ClosingIcon = icon
	} else {
		in.OpeningIcon = icon
	}
	return in
}

/******************************************************************************/

// BuildTag builds the tag used to render the html element.
func (in *ICKInputField) BuildTag() ickcore.Tag {
	in.Tag().SetTagName("div").AddClass("field")
	return *in.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (in *ICKInputField) RenderContent(out io.Writer) error {

	cmpid := in.Tag().Id()
	if cmpid == "" {
		console.Logf("ICKInputField.RenderContent: id missing")
	}

	// <label>
	if in.Label != "" {
		sublbl := ick.Elem("label", `class="label"`, ickcore.ToHTML(in.Label))
		if cmpid != "" {
			sublbl.SetId(cmpid + ".label")
		}
		ickcore.RenderChild(out, in, sublbl)
	}

	// <input>
	subinput := ick.Elem("input", `class="input"`)
	subinput.Tag().
		SetAttributeIf(!in.IsHidden, "type", "text", "password").
		SetAttributeIf(in.Value != "", "value", in.Value).
		SetAttributeIf(in.PlaceHolder != "", "placeholder", in.PlaceHolder).
		SetAttributeIf(in.IsReadOnly, "readonly", "").
		SetAttributeIf(in.State == INPUT_DISABLED, "disabled", "").
		AddClassIf(in.State == INPUT_STATIC, "is-static", "")
	if cmpid != "" {
		subinput.SetId(cmpid + ".input")
	}

	// <div control>
	subcontrol := ick.Elem("div", `class="control"`)
	subcontrol.Tag().
		SetClassIf(in.State == INPUT_LOADING, "is-loading").
		SetClassIf(in.OpeningIcon.NeedRendering(), "has-icons-left").
		SetClassIf(in.ClosingIcon.NeedRendering(), "has-icons-right")

	if in.OpeningIcon.NeedRendering() {
		in.OpeningIcon.Tag().AddClass("is-left")
		subcontrol.Append(&in.OpeningIcon)
	}
	subcontrol.Append(subinput)
	if in.ClosingIcon.NeedRendering() {
		in.ClosingIcon.Tag().AddClass("is-right")
		subcontrol.Append(&in.ClosingIcon)
	}
	ickcore.RenderChild(out, in, subcontrol)

	// <p help>
	if in.Help != "" {
		subhelp := ick.Elem("p", `class="help"`, ickcore.ToHTML(in.Help))
		subhelp.Tag().
			SetClassIf(in.State == INPUT_SUCCESS, "is-success").
			SetClassIf(in.State == INPUT_WARNING, "is-warning").
			SetClassIf(in.State == INPUT_ERROR, "is-danger")
		if cmpid != "" {
			subhelp.SetId(cmpid + ".help")
		}
		ickcore.RenderChild(out, in, subhelp)
	}

	return nil
}

func (in *ICKInputField) AddListeners() {
	in.UI.DOM.AddInputEvent(event.INPUT_ONBEFOREINPUT, in.OnBeforeInputEvent)
	in.UI.DOM.AddInputEvent(event.INPUT_ONINPUT, in.OnInputEvent)
	in.UI.DOM.AddInputEvent(event.INPUT_ONCHANGE, in.OnChangeEvent)
}

func (in *ICKInputField) OnBeforeInputEvent(*event.InputEvent, *dom.Element) {
	console.Warnf("OnBeforeInputEvent")
}
func (in *ICKInputField) OnInputEvent(*event.InputEvent, *dom.Element) {
	console.Warnf("OnInputEvent")

}
func (in *ICKInputField) OnChangeEvent(*event.InputEvent, *dom.Element) {
	console.Warnf("OnChangeEvent")

}

func (in *ICKInputField) RemoveListeners() {
	in.UI.RemoveListeners()
}
