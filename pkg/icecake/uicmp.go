package ick

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/errors"
)

type UIComponent struct {
	Element

	MountClasses    *Classes    // classes added to the component during the mounting stage
	MountAttributes *Attributes // attributes addes to the component during the mounting stage
	UpdateUI        func(any)   // Optional function called to update UI
}

// CastUIComponent is casting a js.Value into UIComponent.
func CastUIComponent(_jsv JSValueProvider) *UIComponent {
	if _jsv.Value().Type() != TypeObject {
		errors.ConsoleErrorf("casting UIComponent failed")
		return new(UIComponent)
	}
	cast := new(UIComponent)
	cast.jsvalue = _jsv.Value().jsvalue
	return cast
}

func (c *UIComponent) Show() {
	fmt.Println("show")
	c.Attributes().RemoveAttribute("hidden")

	// Force a browser re-paint so the browser will realize the
	// element is no longer `hidden` and allow transitions.
	c.ClientRect()
	c.Classes().SetTokens("show")
}

func (c *UIComponent) Hide() {
	fmt.Println("hide")
	c.Classes().RemoveTokens("show")
	c.Attributes().SetAttribute("hidden", "")
}

func (c *UIComponent) Container() (_tagname string, _classes string, _attrs string) {
	fmt.Printf("Component.Container returns default <SPAN> value\n")
	return "SPAN", "", ""
}

func (c *UIComponent) Template() (_html string) {
	fmt.Printf("Component.Template returns default empty value\n")
	return ""
}

func (c *UIComponent) AddListeners() {
	fmt.Printf("Component.AddListeners is empty\n")
}

func (c *UIComponent) GetInitClasses() *Classes {
	return c.MountClasses
}

func (c *UIComponent) GetInitAttributes() *Attributes {
	return c.MountAttributes
}

/*****************************************************************************/

type Composer interface {
	Wrap(JSValueProvider)
	Container() (_tagname string, _classes string, _attrs string)
	GetInitClasses() *Classes
	GetInitAttributes() *Attributes
	Template() (_html string)
	AddListeners()
	Show()
	Hide()
}
