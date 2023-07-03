package bulmaui

import (
	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
)

type ICKButton struct {
	*bulma.ICKButton
	DOM dom.Element
}

func Button(title html.HTMLString, id string, rawURL string, attrs ...string) *ICKButton {
	btn := new(ICKButton)
	btn.ICKButton = bulma.Button(title, id, rawURL, attrs...)
	return btn
}

// returns nil if the id does not exists or if it's not a Navbar
//
// TODO: implement a generic wrapper to any event handlers
func WrapButton(id string) *ICKButton {
	e := dom.RenderedId(id, "ick-button")
	if e == nil {
		return nil
	}
	n := new(ICKButton)
	n.DOM.Wrap(e)
	return n
}

func (btn *ICKButton) SetOutlined(f bool) *ICKButton {
	btn.ICKButton.SetOutlined(f)
	btn.DOM.SetClassIf(f, "is-outlined")
	return btn
}

func (btn *ICKButton) SetRounded(f bool) *ICKButton {
	btn.ICKButton.SetRounded(f)
	btn.DOM.SetClassIf(f, "is-rounded")
	return btn
}

func (btn *ICKButton) SetDisabled(disabled bool) {
	btn.ICKButton.SetDisabled(disabled)
	btn.DOM.SetDisabled(disabled)
}

func (btn *ICKButton) SetLoading(f bool) *ICKButton {
	btn.ICKButton.SetLoading(f)
	btn.DOM.SetClassIf(f, "is-loading")
	return btn
}

func (btn *ICKButton) SetColor(c bulma.COLOR) *ICKButton {
	btn.ICKButton.SetColor(c)
	btn.DOM.PickClass(bulma.COLOR_OPTIONS, string(c))
	return btn
}

func (btn *ICKButton) SetLight(f bool) *ICKButton {
	btn.ICKButton.SetLight(f)
	btn.DOM.SetClassIf(f, "is-light")
	return btn
}
