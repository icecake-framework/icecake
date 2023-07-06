package bulmaui

import (
	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/html"
)

type ICKButton struct {
	bulma.ICKButton
	dom.UI

	OnClick func() // optional OnClick function called by the default button listeners
}

func Button(title html.HTMLString, id string, rawURL string, attrs ...string) *ICKButton {
	btn := new(ICKButton)
	btn.ICKButton = *bulma.Button(title, id, rawURL, attrs...)
	return btn
}

func (btn *ICKButton) AddListeners() {
	if btn.OnClick != nil {
		btn.DOM.AddMouseEvent(event.MOUSE_ONCLICK, func(*event.MouseEvent, *dom.Element) {
			btn.OnClick()
		})
	}
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
