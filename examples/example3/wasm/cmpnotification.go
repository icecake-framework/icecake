package main

import (
	"reflect"
	"time"

	ick "github.com/sunraylab/icecake/pkg/icecake"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	ick.RegisterComponentType("ick-notiftoast", reflect.TypeOf(NotificationToast{}))
}

type NotificationToast struct {
	ick.Element

	timer *time.Timer

	ColorClass string        // this class tring will be added to the class of the envelop element
	Message    string        //
	Timeout    time.Duration // the notification will close automatically after Timeour duration
}

func (c *NotificationToast) Envelope() (_tagname string, _classname string) {
	return "div", "notification {{.Me.ColorClass}}"
}

func (c *NotificationToast) Template() (_html string) {
	return `<button class="delete"></button>
			{{.Me.Message}}`
}

func (c *NotificationToast) AddListeners() {
	btndel := c.SelectorQueryFirst(".delete")
	btndel.AddMouseEvent(ick.MOUSE_ONCLICK, func(*ick.MouseEvent, *ick.Element) {
		if c.timer != nil {
			c.timer.Stop()
		}
		c.Remove()
	})

	if c.Timeout != 0 {
		c.timer = time.AfterFunc(c.Timeout, func() {
			c.Remove()
		})
	}
}
