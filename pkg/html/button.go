package html

import (
	"io"
	"net/url"
)

func init() {
	RegisterComposer("ick-button", &Button{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
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
	HTMLSnippet

	// The button type.
	// If nothing is specified, the default ButtonType is BTN_TYPE_BUTTON.
	ButtonType BUTTON_TYPE

	// If the ButtonType is BTN_TYPE_A then HRef defines the associated url link. HRef has no effect on other ButtonType.
	// If HRef is defined then the button type is automatically set to BTN_TYPE_A.
	// HRef can be nil. Usually it's created calling Button.TryParseHRef
	HRef *url.URL

	// The title of the Button. Can be a simple text or a more complex html string.
	Title HTMLString

	IsOutlined bool // Outlined button style
	IsRounded  bool // Rounded button style

	IsDisabled bool // Disabled state
	IsLoading  bool // Loading button state

	// TODO: handles buttons properties for color, size, display
}

// Ensure Button implements HTMLComposer interface
var _ HTMLComposer = (*Button)(nil)

func NewButton(title HTMLString) *Button {
	btn := new(Button)
	btn.ButtonType = BTN_TYPE_BUTTON
	btn.Title = title
	return btn
}

func NewButtonLink(title HTMLString, _rawUrl string) *Button {
	btn := new(Button)
	btn.ButtonType = BTN_TYPE_A
	btn.Title = title
	btn.ParseAnchor(_rawUrl)
	return btn
}

// ParseAnchor parses _rawUrl to HRef. HRef stays nil in case of error.
func (btn *Button) ParseAnchor(rawUrl string) (err error) {
	btn.HRef, err = url.Parse(rawUrl)
	return
}

func (btn *Button) Tag() *Tag {

	if btn.HRef != nil && btn.HRef.String() != "" {
		btn.ButtonType = BTN_TYPE_A
	}
	switch btn.ButtonType {
	case BTN_TYPE_A:
		btn.tag.SetName("a")
	case BTN_TYPE_SUBMIT:
		btn.tag.SetName("input")
	case BTN_TYPE_RESET:
		btn.tag.SetName("input")
	default:
		btn.tag.SetName("button")
	}

	amap := btn.tag.Attributes()
	amap.SetClasses("button")

	href := ""
	if btn.HRef != nil {
		href = btn.HRef.String()
	}
	switch btn.ButtonType {
	case BTN_TYPE_A:
		amap.SetAttributeIf(href != "", "href", href, true)
	case BTN_TYPE_SUBMIT:
		amap.SetAttribute("type", "submit", true)
	case BTN_TYPE_RESET:
		amap.SetAttribute("type", "reset", true)
	default:
	}

	if btn.IsOutlined {
		amap.SetClasses("is-outlined")
	}
	if btn.IsRounded {
		amap.SetClasses("is-rounded")
	}
	if btn.IsLoading {
		amap.SetClasses("is-loading")
	}
	amap.SetDisabled(btn.IsDisabled)

	return &btn.tag
}

func (btn *Button) WriteBody(out io.Writer) error {
	_, err := WriteString(out, string(btn.Title))
	return err
}
