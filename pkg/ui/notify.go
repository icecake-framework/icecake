package ui

import (
	"time"

	_ "embed"

	ick "github.com/sunraylab/icecake/pkg/icecake"
)

/******************************************************************************
* Component
******************************************************************************/

//go:embed "notify.css"
var css string

func init() {
	ick.App.RegisterComponent("ick-notify", Notify{}, css)
}

type Notify struct {
	ick.UIComponent // embedded Component, with default implementation of composer interfaces

	timer  *time.Timer  // internal timer related to the Tiemout property
	ticker *time.Ticker // internal ticker to handle time left before closing

	TickerStep time.Duration // The optional ticker step, 1s by default
	PopupTime  time.Time     // The last popup time

	// the message to display within the notification
	Message string

	// The notification will close automatically after Timeout duration.
	// The timer starts when the notification pops up.
	Timeout time.Duration
}

func (c *Notify) TimeLeft() time.Duration {
	tl := time.Until(c.PopupTime.Add(c.Timeout))
	if tl < 0 {
		tl = 0
	}
	return tl
}

func (c *Notify) Container() (_tagname string, _classes string, _attrs string) {
	return "div", "notification ick-notify", "hidden"
}

func (c *Notify) Template() (_html string) {
	return `<button class="delete"></button>` + c.Message
}

// AddListeners is called by the dispatcher after DOM rendering
func (c *Notify) AddListeners() {

	btndel := c.SelectorQueryFirst(".delete")
	btndel.AddMouseEvent(ick.MOUSE_ONCLICK, func(*ick.MouseEvent, *ick.Element) {
		//errors.ConsoleLogf("Mouse Event Fired on %s id=%q, %s\n", c.TagName(), c.Id(), c.NodeName())
		c.Stop()
		c.Remove()
	})

	if c.Timeout != 0 {
		if c.TickerStep == 0 {
			c.TickerStep = 1 * time.Second
		}
		c.PopupTime = time.Now()
		if c.UpdateUI != nil {
			go func() {
				c.ticker = time.NewTicker(c.TickerStep)
				if c.IsInDOM() {
					c.UpdateUI(c)
				}
				for _ = range c.ticker.C {
					if c.IsInDOM() {
						c.UpdateUI(c)
					}
				}
			}()
		}
		c.timer = time.AfterFunc(c.Timeout, func() {
			c.Stop()
			c.Remove()
		})
	}
}

func (c *Notify) Stop() {
	if c.timer != nil {
		c.timer.Stop()
	}
	if c.ticker != nil {
		c.ticker.Stop()
	}
}
