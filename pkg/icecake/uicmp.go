package ick

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/errors"
)

type Composer interface {
	Wrap(JSValueProvider)

	Classes() *Classes
	Attributes() *Attributes

	Container(_compid string) (_tagname string, _classes string, _attrs string)
	Body() (_html string)
	Listeners()

	Show()
	Hide()
}

/*****************************************************************************/

type UIComponent struct {
	Element

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

func (c *UIComponent) Container(_compid string) (_tagname string, _classes string, _attrs string) {
	// DEBUG:
	fmt.Printf("UIComponent default <SPAN> Container for %q\n", _compid)

	return "SPAN", "", ""
}

func (c *UIComponent) Body() (_html string) {
	// DEBUG:
	fmt.Printf("UIComponent.Body returns default empty value\n")

	return ""
}

func (c *UIComponent) Listeners() {
	// DEBUG:
	fmt.Printf("UIComponent.Listeners is empty\n")
}

// func (c *UIComponent) Classes() *Classes {
// 	return c.Element.Classes()
// }

// func (c *UIComponent) Attributes() *Attributes {
// 	return &c.Element.attributes
// }
