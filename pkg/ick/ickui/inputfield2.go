package ickui

// type InputField struct {
// 	ick.ICKInputField
// 	DOM dom.Element
// }

// func (input *InputField) AddListeners() {

// 	dom.Id(input.Id())

// 	input.uiinput.DOM.AddInputEvent(event.INPUT_ONINPUT, input.OnInput)

// }

// func (input *InputField) RemoveListeners() {
// 	input.DOM.RemoveListeners()
// }

// func (input *InputField) OnInput(_event *event.InputEvent, _target *dom.Element) {
// 	input.Value = _target.JSValue.String()
// 	console.Warnf("InputField: %s", input.Value)
// }

// func (inputfield *InputField) SetState(_state html.INPUT_STATE) *InputField {
// 	inputfield.State = _state
// 	if inputfield.DOM.IsInDOM() {
// 		switch inputfield.State {
// 		case "success":
// 			inputfield.uicontrol.DOM.RemoveClasses("is-loading")
// 			inputfield.uiinput.DOM.SwitchClasses("is-warning is-danger", "is-success")
// 			inputfield.uiinput.DOM.RemoveAttribute("readonly")
// 			inputfield.uihelp.DOM.SwitchClasses("is-warning is-danger", "is-success")
// 		case "warning":
// 			inputfield.uicontrol.DOM.RemoveClasses("is-loading")
// 			inputfield.uiinput.DOM.SwitchClasses("is-success is-danger", "is-warning")
// 			inputfield.uiinput.DOM.RemoveAttribute("readonly")
// 			inputfield.uihelp.DOM.SwitchClasses("is-success is-danger", "is-warning")
// 		case "error":
// 			inputfield.uicontrol.DOM.RemoveClasses("is-loading")
// 			inputfield.uiinput.DOM.SwitchClasses("is-warning is-success", "is-danger")
// 			inputfield.uiinput.DOM.RemoveAttribute("readonly")
// 			inputfield.uihelp.DOM.SwitchClasses("is-warning is-success", "is-danger")
// 		case "loading":
// 			inputfield.uicontrol.DOM.SetClasses("is-loading")
// 			inputfield.uiinput.DOM.RemoveAttribute("readonly")
// 			inputfield.uiinput.DOM.RemoveClasses("is-success is-warning is-danger")
// 			inputfield.uihelp.DOM.RemoveClasses("is-success is-warning is-danger")
// 		case "readonly":
// 			inputfield.uicontrol.DOM.RemoveClasses("is-loading")
// 			inputfield.uiinput.DOM.SetAttribute("readonly", "")
// 			inputfield.uiinput.DOM.RemoveClasses("is-success is-warning is-danger")
// 			inputfield.uihelp.DOM.RemoveClasses("is-success is-warning is-danger")
// 		default:
// 			inputfield.uicontrol.DOM.RemoveClasses("is-loading")
// 			inputfield.uiinput.DOM.RemoveAttribute("readonly")
// 			inputfield.uiinput.DOM.RemoveClasses("is-success is-warning is-danger")
// 			inputfield.uihelp.DOM.RemoveClasses("is-success is-warning is-danger")
// 		}
// 	}
// 	return inputfield
// }

// func (input *InputField) SetHelp(_help html.HTMLString, _state html.INPUT_STATE) *InputField {
// 	input.Help = _help
// 	input.uihelp.DOM.InsertHTML(dom.INSERT_BODY, _help, nil)
// 	input.SetState(_state)
// 	return input
// }

// func (input *InputField) SetRounded(_f bool) *InputField {
// 	input.IsRounded = _f
// 	if _f {
// 		input.uiinput.DOM.SetClasses("is-rounded")
// 	} else {
// 		input.uiinput.DOM.RemoveClasses("is-rounded")
// 	}
// 	return input
// }
