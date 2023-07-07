package ick

import (
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

	// HRef defines the associated url link. HRef can be nil. If HRef is defined then the rendered element is a <a> tag, otherwise it's a <button> tag.
	HRef *url.URL

	IsOutlined bool  // Outlined button style
	IsRounded  bool  // Rounded button style
	Color      COLOR // rendering color
	IsLight    bool  // light color

	IsDisabled bool // Disabled state

	isLoading bool // Loading button state
}

// Ensuring ICKButton implements the right interface
var _ html.ElementComposer = (*ICKButton)(nil)

func Button(htmltitle string, id string, attrs ...string) *ICKButton {
	btn := new(ICKButton)
	btn.SetTitle(htmltitle)
	btn.Tag().SetId(id)
	btn.Tag().ParseAttributes(attrs...)
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

func (btn *ICKButton) SetTitle(htmltitle string) *ICKButton {
	btn.ClearContent()
	btn.AddContent(html.ToHTML(htmltitle))
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
func (btn *ICKButton) SetDisabled(f bool) *ICKButton {
	btn.IsDisabled = f
	btn.Tag().SetDisabled(f)
	return btn
}
func (btn *ICKButton) SetLoading(f bool) *ICKButton {
	btn.isLoading = f
	return btn
}
func (btn *ICKButton) SetColor(c COLOR) *ICKButton {
	btn.Color = c
	return btn
}
func (btn *ICKButton) SetLight(f bool) *ICKButton {
	btn.IsLight = f
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
		PickClass(COLOR_OPTIONS, string(btn.Color)).
		SetClassIf(btn.IsLight, "is-light")

	if btn.HRef != nil {
		btn.Tag().SetAttribute("href", btn.HRef.String())
	}

	btn.Tag().SetDisabled(btn.IsDisabled)
	return *btn.Tag()
}

// // RenderContent writes the HTML string corresponding to the content of the HTML element.
// // Button rendering unfold the Title
// func (btn *ICKButton) RenderContent(out io.Writer) error {
// 	err := btn.RenderChild(out, &btn.Title)
// 	return err
// }
