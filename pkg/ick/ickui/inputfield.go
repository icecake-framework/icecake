package ickui

import (
	"github.com/icecake-framework/icecake/pkg/console"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/event"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
	"github.com/lolorenzo777/verbose"
)

type ICKInputField struct {
	ick.ICKInputField
	dom.UI

	OnChange func(me *ICKInputField, newvalue string)

	btnToggleVisibility ICKButton
}

// Ensuring InputField implements the right interface
var _ ickcore.ContentComposer = (*ICKInputField)(nil)
var _ ickcore.TagBuilder = (*ICKInputField)(nil)
var _ dom.UIComposer = (*ICKInputField)(nil)

func InputField(id string, value string, placeholder string, attrs ...string) *ICKInputField {
	n := new(ICKInputField)
	n.ICKInputField = *ick.InputField(id, value, placeholder, attrs...)
	return n
}

func (in *ICKInputField) SetLabel(lbl string) *ICKInputField {
	in.ICKInputField.SetLabel(lbl)
	if !in.UI.DOM.IsInDOM() {
		return in
	}

	in.RefreshLabel()
	return in
}

func (in *ICKInputField) SetHelp(help string) *ICKInputField {
	in.ICKInputField.SetHelp(help)
	if !in.UI.DOM.IsInDOM() {
		return in
	}

	in.RefreshHelp()
	return in
}

func (in *ICKInputField) SetCanToggleVisibility(can bool) *ICKInputField {
	in.ICKInputField.SetCanToggleVisibility(can)
	if !in.UI.DOM.IsInDOM() {
		return in
	}

	etggctrl := dom.Id(in.Tag().SubId("togglecontrol"))
	if can {
		// setup = add the togglecontrol if it does not exist yet otherwise does nothing
		if !etggctrl.IsDefined() {
			subtogglecontrol := ick.Elem("div", `class="control"`).SetId(in.Tag().SubId("togglecontrol"))
			if in.IsHidden {
				in.btnToggleVisibility.ICKButton = *ick.BtnVisibilityVisible.Clone()
			} else {
				in.btnToggleVisibility.ICKButton = *ick.BtnVisibilityHidden.Clone()
			}
			in.btnToggleVisibility.Tag().AddClass("is-right").SetId(in.Tag().SubId("btntoggvis"))
			in.btnToggleVisibility.OnClick = in.ToggleVisibility
			subtogglecontrol.Append(&in.btnToggleVisibility)
			dom.Id(in.Tag().SubId("input")).InsertSnippet(dom.INSERT_AFTER_ME, subtogglecontrol)
		}
	} else {
		// unset = remove the togglecontrol
		etggctrl.Remove()
	}

	return in
}

func (in *ICKInputField) SetReadOnly(ro bool) *ICKInputField {
	in.ICKInputField.SetReadOnly(ro)
	if !in.UI.DOM.IsInDOM() {
		return in
	}

	dom.Id(in.Tag().SubId("input")).
		SetAttributeIf(in.IsReadOnly, "readonly", "")
	return in
}

func (in *ICKInputField) SetHidden(h bool) *ICKInputField {
	in.ICKInputField.SetHidden(h)
	if !in.UI.DOM.IsInDOM() {
		return in
	}

	dom.Id(in.Tag().SubId("input")).
		SetAttributeIf(!in.IsHidden, "type", "text", "password")
	return in
}

func (in *ICKInputField) SetState(st ick.INPUT_STATE) *ICKInputField {
	in.ICKInputField.SetState(st)
	if !in.UI.DOM.IsInDOM() {
		return in
	}

	dom.Id(in.Tag().SubId("input")).
		SetAttributeIf(in.State == ick.INPUT_DISABLED, "disabled", "").
		AddClassIf(in.State == ick.INPUT_STATIC, "is-static", "")
	dom.Id(in.Tag().SubId("control")).
		SetClassIf(in.State == ick.INPUT_LOADING, "is-loading")
	return in
}

func (in *ICKInputField) SetDisabled(f bool) *ICKInputField {
	in.ICKInputField.SetDisabled(f)
	if !in.UI.DOM.IsInDOM() {
		return in
	}

	dom.Id(in.Tag().SubId("input")).
		SetAttributeIf(in.State == ick.INPUT_DISABLED, "disabled", "")
	return in
}

func (in *ICKInputField) SetIcon(icon ick.ICKIcon, closing bool) *ICKInputField {
	in.ICKInputField.SetIcon(icon, closing)
	if !in.UI.DOM.IsInDOM() {
		return in
	}
	var eicon *dom.Element
	var where dom.INSERT_WHERE
	if closing {
		eicon = dom.Id(in.Tag().SubId("cicon"))
		where = dom.INSERT_AFTER_ME
	} else {
		eicon = dom.Id(in.Tag().SubId("oicon"))
		where = dom.INSERT_BEFORE_ME
	}
	if eicon.IsDefined() {
		// replace
		if !icon.NeedRendering() {
			eicon.Remove()
		} else {
			eicon.InsertSnippet(dom.INSERT_OUTER, &icon)
		}
	} else if icon.NeedRendering() {
		// add
		dom.Id(in.Tag().SubId("input")).InsertSnippet(where, &icon)
	}

	return in
}

/******************************************************************************/

func (in *ICKInputField) RefreshLabel() {
	lblid := in.Tag().SubId("label")
	if in.Label == "" {
		// no label = remove it from the dom if any
		dom.Id(lblid).Remove()
	} else if !dom.Id(lblid).IsDefined() {
		// label not yet in the dom = insert it before the control
		sublbl := ick.Elem("label", `class="label"`, ickcore.ToHTML(in.Label))
		sublbl.SetId(lblid)
		dom.Id(in.Tag().SubId("control")).InsertSnippet(dom.INSERT_BEFORE_ME, sublbl)
	} else {
		// label already in the dom = update it
		isin := dom.Id(lblid).InnerHTML()
		if isin != in.Label {
			dom.Id(lblid).InsertSnippet(dom.INSERT_BODY, ickcore.ToHTML(in.Label))
		}
	}
}

func (in *ICKInputField) RefreshHelp() {
	helpid := in.Tag().SubId("help")
	if in.Help == "" {
		// no help = remove it from the dom if any
		dom.Id(helpid).Remove()
	} else if !dom.Id(helpid).IsDefined() {
		// help not yet in the dom = insert it after the control
		subhelp := ick.Elem("p", `class="help"`, ickcore.ToHTML(in.Help))
		subhelp.SetId(helpid)
		dom.Id(in.Tag().SubId("control")).InsertSnippet(dom.INSERT_AFTER_ME, subhelp)
	} else {
		// help already in the dom = update it
		isin := dom.Id(helpid).InnerHTML()
		if isin != in.Help {
			dom.Id(helpid).InsertSnippet(dom.INSERT_BODY, ickcore.ToHTML(in.Help))
		}
	}
}

func (in *ICKInputField) AddListeners() {
	// DEBUG: console.Warnf("ICKInputField.AddListeners: %q", in.DOM.Id())
	console.Warnf("ICKInputField.AddListeners: %q", in.DOM.Id())

	// in.UI.DOM.AddInputEvent(event.INPUT_ONBEFOREINPUT, in.OnBeforeInputEvent)
	// in.UI.DOM.AddInputEvent(event.INPUT_ONINPUT, in.OnInputEvent)
	dom.Id(in.Tag().SubId("input")).AddInputEvent(event.INPUT_ONCHANGE, in.OnChangeEvent)

	in.btnToggleVisibility.OnClick = in.ToggleVisibility
	dom.TryMountId(&in.btnToggleVisibility, in.Tag().SubId("btntoggvis"))
}

func (in *ICKInputField) OnBeforeInputEvent(*event.InputEvent, *dom.Element) {
	console.Warnf("OnBeforeInputEvent")
}

func (in *ICKInputField) OnInputEvent(*event.InputEvent, *dom.Element) {
	console.Warnf("OnInputEvent")

}

func (in *ICKInputField) OnChangeEvent(*event.InputEvent, *dom.Element) {
	in.Value = dom.Id(in.Tag().SubId("input")).GetString("value")
	if in.OnChange != nil {
		in.OnChange(in, in.Value)
	}
	console.Warnf("OnChangeEvent: %+v", in.Value)
}

func (in *ICKInputField) RemoveListeners() {
	in.btnToggleVisibility.RemoveListeners()
	in.UI.RemoveListeners()
}

func (in *ICKInputField) ToggleVisibility() {
	verbose.Debug("ToggleVisibility")

	if in.IsHidden {
		in.SetHidden(false)
		in.btnToggleVisibility.OpeningIcon = *ick.BtnVisibilityHidden.OpeningIcon.Clone()
		in.btnToggleVisibility.UI.RefreshContent(&in.btnToggleVisibility)
	} else {
		in.SetHidden(true)
		in.btnToggleVisibility.OpeningIcon = *ick.BtnVisibilityVisible.OpeningIcon.Clone()
		in.btnToggleVisibility.UI.RefreshContent(&in.btnToggleVisibility)
	}
	dom.Id(in.Tag().SubId("input")).Focus()
}
