package ick

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/errors"
)

type UIComponent struct {
	Element

	MountClasses    *Classes    // classes added to the component during the mounting stage
	MountAttributes *Attributes // attributes addes to the component during the mounting stage
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

func (c *UIComponent) Show() {
	c.Attributes().RemoveAttribute("hidden")

	// Force a browser re-paint so the browser will realize the
	// element is no longer `hidden` and allow transitions.
	c.ClientRect()
	c.Classes().SetTokens("show")
}

func (c *UIComponent) Hide() {
	c.Classes().RemoveTokens("show")
	c.Attributes().SetAttribute("hidden", "")
}

func (c *UIComponent) UpdateUI() {
	// DEBUG:
	fmt.Printf("UIComponent.UpdateUI does nothing by default\n")
}

func (c *UIComponent) Container() (_tagname string, _classes string, _attrs string) {
	// DEBUG:
	fmt.Printf("UIComponent.Container returns default <SPAN> value\n")

	return "SPAN", "", ""
}

func (c *UIComponent) Template() (_html string) {
	// DEBUG:
	fmt.Printf("UIComponent.Template returns default empty value\n")

	return ""
}

func (c *UIComponent) AddListeners() {
	fmt.Printf("UIComponent.AddListeners is empty\n")
}

func (c *UIComponent) GetInitClasses() *Classes {
	return c.MountClasses
}

func (c *UIComponent) GetInitAttributes() *Attributes {
	return c.MountAttributes
}

/*****************************************************************************/

type UIUpdater interface {
	UpdateUI()
}

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
