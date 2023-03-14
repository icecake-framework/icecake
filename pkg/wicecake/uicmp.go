package wick

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/errors"
	ick "github.com/sunraylab/icecake/pkg/icecake"
)

type UIComposer interface {
	ick.HtmlComposer

	Wrap(JSValueProvider)
	AddListeners()
	UpdateUI()
}

/*****************************************************************************/

type UIComponent struct {
	Element

	ick.HtmlComponent

	//MountClasses    *Classes    // classes added to the component during the mounting stage
	//MountAttributes *Attributes // attributes addes to the component during the mounting stage
	//UpdateUI        func(any)   // Optional function called to update UI
}

// CastUIComponent is casting a js.Value into UIComponent.
func CastUIComponent(_jsv JSValueProvider) *UIComponent {
	if _jsv.Value().Type() != TYPE_OBJECT {
		errors.ConsoleErrorf("casting UIComponent failed")
		return new(UIComponent)
	}
	cast := new(UIComponent)
	cast.jsvalue = _jsv.Value().jsvalue
	return cast
}

// func (c *UIComponent) Show() {
// 	c.Attributes.RemoveAttribute("hidden")

// 	// Force a browser re-paint so the browser will realize the
// 	// element is no longer `hidden` and allow transitions.
// 	c.ClientRect()
// 	c.Classes().AddTokens("show")
// }

// func (c *UIComponent) Hide() {
// 	c.Classes().RemoveTokens("show")
// 	c.Attributes().SetAttribute("hidden", "")
// }

func (c *UIComponent) UpdateUI() {
	// DEBUG:
	fmt.Printf("UIComponent.UpdateUI does nothing by default\n")
}

func (c *UIComponent) AddListeners() {
	// DEBUG:
	fmt.Printf("UIComponent.Listeners is empty\n")
}

// func (c *UIComponent) Classes() *Classes {
// 	return c.Element.Classes()
// }

// func (c *UIComponent) Attributes() *Attributes {
// 	return &c.Element.attributes
// }
