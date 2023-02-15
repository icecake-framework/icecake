package dom

import "syscall/js"

/****************************************************************************
* HTMLHeadElement
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLHeadElement
type HeadElement struct {
	HTMLElement
}

// CastHeadElement is casting a js.Value into HTMLHeadElement.
func CastHeadElement(value js.Value) *HeadElement {
	if value.Type() != js.TypeObject {
		ConsoleError("casting HeadElement failed")
		return nil
	}
	ret := new(HeadElement)
	ret.jsValue = value
	return ret
}
