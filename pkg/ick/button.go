package ick

import (
	"io"
	"net/url"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

func init() {
	ickcore.RegisterComposer("ick-button", &ICKButton{})
}

// ICKButton is an UISnippet registered with the ick-tag `ick-button`.
//
// According to the ButtonType property, ICKButton can be used either as a standard <button> element but also as an anchor link or a submit or reset form input.
// The core text is handle with the Title html property, allowing simple text or complex rendering.
//
// The IsDisabled property is directly handled by the embedded UISnippet.
type ICKButton struct {
	ickcore.BareSnippet

	OpeningIcon ICKIcon // optional opening icon
	Title       string  // simple title string
	ClosingIcon ICKIcon // optional closing icon

	// HRef defines the associated url link. HRef can be nil. If HRef is defined then the rendered element is a <a> tag, otherwise it's a <button> tag.
	HRef *url.URL

	IsOutlined bool // Outlined button style
	IsRounded  bool // Rounded button style
	COLOR           // rendering color
	SIZE            // button size

	IsDisabled bool // Disabled state

	IsLoading bool // Loading button state

	// TODO: ick.ICKButton - add a feature to automatically setup the link color and a closing link icon when ther's an HREF withe a base different than the one of the website
}

// Ensuring ICKButton implements the right interface
var _ ickcore.ContentComposer = (*ICKButton)(nil)
var _ ickcore.TagBuilder = (*ICKButton)(nil)

func Button(htmltitle string, attrs ...string) *ICKButton {
	btn := new(ICKButton)
	btn.SetTitle(htmltitle)
	btn.Tag().ParseAttributes(attrs...)
	return btn
}

func (btn *ICKButton) NeedRendering() bool {
	return btn.OpeningIcon.NeedRendering() || btn.Title != "" || btn.ClosingIcon.NeedRendering()
}

// Clone clones the snippet, without the rendering metadata
func (btn *ICKButton) Clone() *ICKButton {
	c := new(ICKButton)
	*c = *btn
	c.BareSnippet = *btn.BareSnippet.Clone()
	return c
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

func (btn *ICKButton) SetIcon(icon ICKIcon, closing bool) *ICKButton {
	if closing {
		btn.ClosingIcon = icon
	} else {
		btn.OpeningIcon = icon
	}
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
	btn.IsLoading = f
	return btn
}

/******************************************************************************/

// BuildTag builds the tag used to render the html element.
// The tagname depends on the button type.
func (btn *ICKButton) BuildTag() ickcore.Tag {
	if btn.HRef != nil && btn.HRef.String() != "" {
		btn.Tag().SetTagName("a")
	} else {
		btn.Tag().SetTagName("button")
	}

	btn.Tag().AddClass("button").
		SetClassIf(btn.IsOutlined, "is-outlined").
		SetClassIf(btn.IsRounded, "is-rounded").
		SetClassIf(btn.IsLoading, "is-loading").
		PickClass(COLOR_OPTIONS, string(btn.COLOR)).
		PickClass(SIZE_OPTIONS, string(btn.SIZE))

	if btn.HRef != nil {
		btn.Tag().SetAttribute("href", btn.HRef.String())
	}

	btn.Tag().SetDisabled(btn.IsDisabled)
	return *btn.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (btn *ICKButton) RenderContent(out io.Writer) error {
	has := btn.OpeningIcon.Key != "" || btn.ClosingIcon.Key != ""

	ickcore.RenderChild(out, btn, &btn.OpeningIcon)
	ickcore.RenderStringIf(has, out, "<span>")
	ickcore.RenderString(out, btn.Title)
	ickcore.RenderStringIf(has, out, "</span>")
	ickcore.RenderChild(out, btn, &btn.ClosingIcon)
	return nil
}
