package ui

import (
	_ "embed"

	ick "github.com/sunraylab/icecake/pkg/icecake"
)

/******************************************************************************
* Component
******************************************************************************/

//go:embed "notify.css"
var notifycss string

func init() {
	ick.App.RegisterComponent("ick-notify", Notify{}, notifycss)
}

type Notify struct {
	ick.UIComponent // embedded Component, with default implementation of composer interfaces

	// the message to display within the notification
	Message string

	Delete // the delete sub-component
}

func (c *Notify) Container(_compid string) (_tagname string, _classes string, _attrs string) {
	c.Delete.TargetID = _compid
	return "div", "ick-notify notification", ""
}

func (c *Notify) Body() (_html string) {

	//_, del, _ := ick.App.CreateComponent(&c.Delete)

	return "{{.Me.Delete}}" + c.Message
	//return `<button class="delete"></button>` + c.Message
}

func (c *Notify) Listeners() {
}
