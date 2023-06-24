package bulma

import (
	"io"
	"net/url"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-button", &Button{})
}

// The button type can be one of the following:
//
//	BTN_TYPE_BUTTON  // <button> form buttons.
//	BTN_TYPE_A       // <a> anchor links
//	BTN_TYPE_SUBMIT  // <input type="submit"> submit inputs
//	BTN_TYPE_RESET   // <input type="reset"> reset inputs
type BUTTON_TYPE int

const (
	BTN_TYPE_BUTTON BUTTON_TYPE = iota // <button> form buttons.
	BTN_TYPE_A                         // <a> anchor link
	BTN_TYPE_SUBMIT                    // <input type="submit"> submit inputs
	BTN_TYPE_RESET                     // <input type="reset"> reset inputs
)

// Button is an UISnippet registered with the ick-tag `ick-button`.
//
// According to the ButtonType property, Button can be used either as a standard <button> element but also as an anchor link or a submit or reset form input.
// The core text is handle with the Title html property, allowing simple text or complex rendering.
//
// The IsDisabled property is directly handled by the embedded UISnippet.
type Button struct {
	html.HTMLSnippet

	// The button type.
	// If nothing is specified, the default ButtonType is BTN_TYPE_BUTTON.
	ButtonType BUTTON_TYPE

	// If the ButtonType is BTN_TYPE_A then HRef defines the associated url link. HRef has no effect on other ButtonType.
	// If HRef is defined then the button type is automatically set to BTN_TYPE_A.
	// HRef can be nil. Usually it's created calling Button.TryParseHRef
	HRef *url.URL

	// The title of the Button. Can be a simple text or a more complex html string.
	Title html.HTMLString

	IsOutlined bool // Outlined button style
	IsRounded  bool // Rounded button style

	IsDisabled bool // Disabled state
	IsLoading  bool // Loading button state
}

// Ensure Button implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*Button)(nil)

func NewButton(title html.HTMLString) *Button {
	btn := new(Button)
	btn.ButtonType = BTN_TYPE_BUTTON
	btn.Title = title
	return btn
}

func NewButtonLink(title html.HTMLString, rawUrl string) *Button {
	btn := new(Button)
	btn.ButtonType = BTN_TYPE_A
	btn.Title = title
	btn.ParseHRef(rawUrl)
	return btn
}

// ParseHRef parses _rawUrl to HRef. HRef stays nil in case of error.
func (btn *Button) ParseHRef(rawUrl string) (err error) {
	btn.HRef, err = url.Parse(rawUrl)
	return
}

// BuildTag builds the tag used to render the html element.
// The tagname depends on the button type.
func (btn *Button) BuildTag(tag *html.Tag) {

	if btn.HRef != nil && btn.HRef.String() != "" {
		btn.ButtonType = BTN_TYPE_A
	}
	switch btn.ButtonType {
	case BTN_TYPE_A:
		tag.SetTagName("a")
	case BTN_TYPE_SUBMIT:
		tag.SetTagName("input")
	case BTN_TYPE_RESET:
		tag.SetTagName("input")
	default:
		tag.SetTagName("button")
	}

	tag.AddClasses("button").
		SetClassesIf(btn.IsOutlined, "is-outlined").
		SetClassesIf(btn.IsRounded, "is-rounded").
		SetClassesIf(btn.IsLoading, "is-loading")

	href := ""
	if btn.HRef != nil {
		href = btn.HRef.String()
	}
	switch btn.ButtonType {
	case BTN_TYPE_A:
		tag.RemoveAttribute("type")
		tag.SetAttributeIf(href != "", "href", href)
	case BTN_TYPE_SUBMIT:
		tag.RemoveAttribute("href")
		tag.SetAttribute("type", "submit")
	case BTN_TYPE_RESET:
		tag.RemoveAttribute("href")
		tag.SetAttribute("type", "reset")
	default:
	}

	tag.SetDisabled(btn.IsDisabled)
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// Button rendering unfold the Title
func (btn *Button) RenderContent(out io.Writer) error {
	err := btn.RenderChilds(out, &btn.Title)
	return err
}
