package uicomponent

import (
	"fmt"
	"time"

	ick "github.com/sunraylab/icecake/pkg/icecake"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	ick.RegisterComponentType("ick-notify", Notify{})
}

type Notify struct {
	ick.Component

	timer *time.Timer

	ColorClass string        // class string added to the class of the envelop element
	Message    string        // message to display
	Timeout    time.Duration // the notification will close automatically after Timeour duration
}

func (c *Notify) Envelope() (_tagname string, _classes string) {
	fmt.Printf("c %p envelope\n", c)
	return "div", "notification {{.Me.ColorClass}}"
}

func (c *Notify) Template() (_html string) {
	fmt.Printf("c %p template\n", c)
	return `<button class="delete"></button>
			{{.Me.Message}}<br/><div class="pt-4">Message number:&nbsp;{{.Global.msgnumber}}</div>`
}

// the component should have been already renddered into the DOM
func (c *Notify) AddListeners() {
	// DEBUG:
	fmt.Printf("c %p listener\n", c)

	btndel := c.SelectorQueryFirst(".delete")
	//ick.App().Body().AddMouseEvent(ick.MOUSE_ONCLICK, func(*ick.MouseEvent, *ick.Element) {
	btndel.AddMouseEvent(ick.MOUSE_ONCLICK, func(*ick.MouseEvent, *ick.Element) {
		ick.ConsoleWarnf("Mouse Event Fired on %s id=%q, %s\n", c.TagName(), c.Id(), c.NodeName())
		if c.timer != nil {
			c.timer.Stop()
		}
		// DEBUG:
		c.Remove()
	})

	if c.Timeout != 0 {
		c.timer = time.AfterFunc(c.Timeout, func() {
			c.Remove()
		})
	}
}
