package dom

import (
	"github.com/sunraylab/icecake/pkg/console"
	"github.com/sunraylab/icecake/pkg/js"
)

/****************************************************************************
* HTMLHeadElement
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLHeadElement
type HeadElement struct {
	Element
}

// CastHeadElement is casting a js.Value into HTMLHeadElement.
func CastHeadElement(_jsvp js.JSValueProvider) *HeadElement {
	if _jsvp.Value().Type() != js.TYPE_OBJECT {
		console.Errorf("casting HeadElement failed")
		return new(HeadElement)
	}
	cast := new(HeadElement)
	cast.JSValue = _jsvp.Value()
	return cast
}
