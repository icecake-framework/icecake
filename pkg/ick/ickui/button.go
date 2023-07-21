package ickui

import (
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/ick"
)

type ICKButton struct {
	ick.ICKButton
	dom.UI

	OnClick func() // optional OnClick function called by the default button listeners
}

// Ensure ICKButton implements UIComposer interface
var _ dom.UIComposer = (*ICKButton)(nil)

func Button(title string, attrs ...string) *ICKButton {
	btn := new(ICKButton)
	btn.ICKButton = *ick.Button(title, attrs...)
	return btn
}

func (cmp *ICKButton) SetId(id string) *ICKButton {
	cmp.Tag().SetId(id)
	return cmp
}

func (btn *ICKButton) SetTitle(title string) *ICKButton {
	btn.ICKButton.SetTitle(title)
	btn.UI.RefreshContent(btn)
	return btn
}
func (btn *ICKButton) SetIcon(icon ick.ICKIcon, closing bool) *ICKButton {
	if closing {
		btn.ClosingIcon = icon
	} else {
		btn.OpeningIcon = icon
	}
	btn.UI.RefreshContent(btn)
	return btn
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
func (btn *ICKButton) SetDisabled(disabled bool) *ICKButton {
	btn.ICKButton.SetDisabled(disabled)
	btn.DOM.SetDisabled(disabled)
	return btn
}
func (btn *ICKButton) SetSize(sz ick.SIZE) *ICKButton {
	btn.ICKButton.SetSize(sz)
	btn.DOM.PickClass(ick.SIZE_OPTIONS, string(btn.SIZE))
	return btn
}
func (btn *ICKButton) SetLoading(f bool) *ICKButton {
	btn.ICKButton.SetLoading(f)
	btn.DOM.SetClassIf(f, "is-loading")
	return btn
}
func (btn *ICKButton) SetColor(c ick.COLOR) *ICKButton {
	btn.ICKButton.SetColor(c)
	btn.DOM.PickClass(ick.COLOR_OPTIONS, string(c))
	return btn
}
func (btn *ICKButton) SetLight(f bool) *ICKButton {
	btn.ICKButton.SetLight(f)
	btn.DOM.SetClassIf(f, "is-light")
	return btn
}

/******************************************************************************/

func (btn *ICKButton) AddListeners() {
	if btn.OnClick != nil {
		btn.DOM.AddMouseEvent(event.MOUSE_ONCLICK, func(*event.MouseEvent, *dom.Element) {
			btn.OnClick()
		})
	}
}
