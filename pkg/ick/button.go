package ick

import (
	"io"
	"net/url"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-button", &ICKButton{})
}

// ICKButton is an UISnippet registered with the ick-tag `ick-button`.
//
// According to the ButtonType property, ICKButton can be used either as a standard <button> element but also as an anchor link or a submit or reset form input.
// The core text is handle with the Title html property, allowing simple text or complex rendering.
//
// The IsDisabled property is directly handled by the embedded UISnippet.
type ICKButton struct {
	html.HTMLSnippet

	Title string // simple title string

	// HRef defines the associated url link. HRef can be nil. If HRef is defined then the rendered element is a <a> tag, otherwise it's a <button> tag.
	HRef *url.URL

	IsOutlined bool // Outlined button style
	IsRounded  bool // Rounded button style
	COLOR           // rendering color
	SIZE            // button size

	IsDisabled bool // Disabled state

	isLoading bool // Loading button state
}

// Ensuring ICKButton implements the right interface
var _ html.ElementComposer = (*ICKButton)(nil)

func Button(htmltitle string, attrs ...string) *ICKButton {
	btn := new(ICKButton)
	btn.SetTitle(htmltitle)
	btn.Tag().ParseAttributes(attrs...)
	return btn
}

// Clone clones the snippet, without the rendering metadata
func (btn *ICKButton) Clone() *ICKButton {
	// TODO: reset id and metadata
	to := new(ICKButton)
	*to = *btn
	to.HTMLSnippet = *btn.HTMLSnippet.Clone()
	return to
}

func (btn *ICKButton) SetId(id string) *ICKButton {
	btn.Tag().SetId(id)
	return btn
}

// ParseHRef parses rawurl to HRef. HRef stays nil in case of error.
func (btn *ICKButton) ParseHRef(rawurl string) *ICKButton {
	btn.HRef = nil
	if rawurl != "" {
		btn.HRef, _ = url.Parse(rawurl)
	}
	return btn
}

func (btn *ICKButton) SetTitle(title string) *ICKButton {
	btn.Title = title
	return btn
}

func (btn *ICKButton) SetOutlined(f bool) *ICKButton {
	btn.IsOutlined = f
	return btn
}
func (btn *ICKButton) SetRounded(f bool) *ICKButton {
	btn.IsRounded = f
	return btn
}
func (btn *ICKButton) SetColor(c COLOR) *ICKButton {
	btn.COLOR = c
	return btn
}

func (btn *ICKButton) SetSize(s SIZE) *ICKButton {
	btn.SIZE = s
	return btn
}

func (btn *ICKButton) SetDisabled(f bool) *ICKButton {
	btn.IsDisabled = f
	btn.Tag().SetDisabled(f)
	return btn
}
func (btn *ICKButton) SetLoading(f bool) *ICKButton {
	btn.isLoading = f
	return btn
}

/******************************************************************************/

// BuildTag builds the tag used to render the html element.
// The tagname depends on the button type.
func (btn *ICKButton) BuildTag() html.Tag {
	if btn.HRef != nil && btn.HRef.String() != "" {
		btn.Tag().SetTagName("a")
	} else {
		btn.Tag().SetTagName("button")
	}

	btn.Tag().AddClass("button").
		SetClassIf(btn.IsOutlined, "is-outlined").
		SetClassIf(btn.IsRounded, "is-rounded").
		SetClassIf(btn.isLoading, "is-loading").
		PickClass(COLOR_OPTIONS, string(btn.COLOR)).
		// SetClassIf(btn.LightColor, "is-light").
		PickClass(SIZE_OPTIONS, string(btn.SIZE))

	if btn.HRef != nil {
		btn.Tag().SetAttribute("href", btn.HRef.String())
	}

	btn.Tag().SetDisabled(btn.IsDisabled)
	return *btn.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (btn *ICKButton) RenderContent(out io.Writer) error {
	html.WriteString(out, btn.Title)
	return nil
}
