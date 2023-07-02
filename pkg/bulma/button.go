package bulma

import (
	"io"
	"net/url"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-button", &Button{})
}

// Button is an UISnippet registered with the ick-tag `ick-button`.
//
// According to the ButtonType property, Button can be used either as a standard <button> element but also as an anchor link or a submit or reset form input.
// The core text is handle with the Title html property, allowing simple text or complex rendering.
//
// The IsDisabled property is directly handled by the embedded UISnippet.
type Button struct {
	html.HTMLSnippet

	// The title of the Button. Can be a simple text or a more complex html string.
	Title html.HTMLString

	// HRef defines the associated url link. HRef can be nil. If HRef is defined then the rendered element is a <a> tag, otherwise it's a <button> tag.
	HRef *url.URL

	IsOutlined bool  // Outlined button style
	IsRounded  bool  // Rounded button style
	Color      COLOR // rendering color
	IsLight    bool  // light color

	IsDisabled bool // Disabled state

	isLoading bool // Loading button state
}

// Ensure Button implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*Button)(nil)

func NewButton(title html.HTMLString, id string, rawurl string, attrs ...string) *Button {
	btn := new(Button)
	btn.Tag().SetId(id)
	btn.ParseHRef(rawurl)
	btn.Title = title
	btn.Tag().ParseAttributes(attrs...)
	return btn
}

// ParseHRef parses _rawUrl to HRef. HRef stays nil in case of error.
func (btn *Button) ParseHRef(rawurl string) (err error) {
	if rawurl != "" {
		btn.HRef, err = url.Parse(rawurl)
	} else {
		btn.HRef = nil
	}
	return
}

func (btn *Button) SetOutlined(f bool) *Button {
	btn.IsOutlined = f
	return btn
}
func (btn *Button) SetRounded(f bool) *Button {
	btn.IsRounded = f
	return btn
}
func (btn *Button) SetDisabled(f bool) *Button {
	btn.IsDisabled = f
	btn.Tag().SetDisabled(f)
	return btn
}
func (btn *Button) SetLoading(f bool) *Button {
	btn.isLoading = f
	return btn
}
func (btn *Button) SetColor(c COLOR) *Button {
	btn.Color = c
	return btn
}
func (btn *Button) SetLight(f bool) *Button {
	btn.IsLight = f
	return btn
}

// BuildTag builds the tag used to render the html element.
// The tagname depends on the button type.
func (btn *Button) BuildTag(tag *html.Tag) {

	if btn.HRef != nil && btn.HRef.String() != "" {
		tag.SetTagName("a")
	} else {
		tag.SetTagName("button")
	}

	tag.AddClass("button").
		SetClassIf(btn.IsOutlined, "is-outlined").
		SetClassIf(btn.IsRounded, "is-rounded").
		SetClassIf(btn.isLoading, "is-loading").
		PickClass(COLOR_OPTIONS, string(btn.Color)).
		SetClassIf(btn.IsLight, "is-light")

	if btn.HRef != nil {
		tag.SetAttribute("href", btn.HRef.String())
	}

	tag.SetDisabled(btn.IsDisabled)
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// Button rendering unfold the Title
func (btn *Button) RenderContent(out io.Writer) error {
	err := btn.RenderChilds(out, &btn.Title)
	return err
}
