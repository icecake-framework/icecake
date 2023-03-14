package ui

import (
	_ "embed"

	ick "github.com/sunraylab/icecake/pkg/icecake"
	wick "github.com/sunraylab/icecake/pkg/wicecake"
)

/******************************************************************************
* Component
******************************************************************************/

//go:embed "notify.css"
var notifycss string

func init() {
	ick.TheCmpReg.RegisterComponent(Notify{})
}

type Notify struct {
	wick.UIComponent // embedded Component, with default implementation of composer interfaces

	// the message to display within the notification
	Message string

	Delete // the delete sub-component
}

func (*Notify) RegisterName() string {
	return "ick-notify"
}

func (*Notify) RegisterCSS() string {
	return notifycss
}

func (c *Notify) Container(_compid string) (_tagname string, _contclasses string, _contattrs string, _contstyle string) {
	c.Delete.TargetID = _compid
	return "div", "ick-notify notification", "", ""
}

func (c *Notify) Body() (_html string) {

	//_, del, _ := ick.App.CreateComponent(&c.Delete)

	return "{{.Me.Delete}}" + c.Message
	//return `<button class="delete"></button>` + c.Message
}
