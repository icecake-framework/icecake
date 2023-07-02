package bulmaui

import (
	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
)

// Button is an UISnippet registered with the ick-tag `ick-button`.
//
// According to the ButtonType property, Button can be used either as a standard <button> element but also as an anchor link or a submit or reset form input.
// The core text is handle with the Title html property, allowing simple text or complex rendering.
//
// The IsDisabled property is directly handled by the embedded UISnippet.
type Button struct {
	*bulma.Button
	DOM dom.Element
}

func NewButton(title html.HTMLString, id string, rawURL string, attrs ...string) *Button {
	btn := new(Button)
	btn.Button = bulma.NewButton(title, id, rawURL, attrs...)
	return btn
}

// returns nil if the id does not exists or if it's not a Navbar
//
// TODO: implement a generic wrapper to any event handlers
func WrapButton(id string) *Button {
	e := dom.RenderedId(id, "ick-button")
	if e == nil {
		return nil
	}
	n := new(Button)
	n.DOM.Wrap(e)
	return n
}

func (btn *Button) SetOutlined(f bool) *Button {
	btn.Button.SetOutlined(f)
	btn.DOM.SetClassIf(f, "is-outlined")
	return btn
}

func (btn *Button) SetRounded(f bool) *Button {
	btn.Button.SetRounded(f)
	btn.DOM.SetClassIf(f, "is-rounded")
	return btn
}

func (btn *Button) SetDisabled(disabled bool) {
	btn.Button.SetDisabled(disabled)
	btn.DOM.SetDisabled(disabled)
}

func (btn *Button) SetLoading(f bool) *Button {
	btn.Button.SetLoading(f)
	btn.DOM.SetClassIf(f, "is-loading")
	return btn
}

func (btn *Button) SetColor(c bulma.COLOR) *Button {
	btn.Button.SetColor(c)
	btn.DOM.PickClass(bulma.COLOR_OPTIONS, string(c))
	return btn
}

func (btn *Button) SetLight(f bool) *Button {
	btn.Button.SetLight(f)
	btn.DOM.SetClassIf(f, "is-light")
	return btn
}
