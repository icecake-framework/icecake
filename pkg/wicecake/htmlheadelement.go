package wick

import (
	"github.com/sunraylab/icecake/pkg/console"
)

/****************************************************************************
* HTMLHeadElement
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLHeadElement
type HeadElement struct {
	Element
}

// CastHeadElement is casting a js.Value into HTMLHeadElement.
func CastHeadElement(_jsvp JSValueProvider) *HeadElement {
	if _jsvp.Value().Type() != TYPE_OBJECT {
		console.Errorf("casting HeadElement failed")
		return new(HeadElement)
	}
	cast := new(HeadElement)
	cast.jsvalue = _jsvp.Value().jsvalue
	return cast
}
