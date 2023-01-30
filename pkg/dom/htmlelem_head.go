package dom

import "syscall/js"

/****************************************************************************
* HTMLHeadElement
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/HTMLHeadElement
type HTMLHeadElement struct {
	HTMLElement
}

// NewHTMLHeadElementFromJS is casting a js.Value into HTMLHeadElement.
func NewHTMLHeadElementFromJS(value js.Value) *HTMLHeadElement {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &HTMLHeadElement{}
	ret.jsValue = value
	return ret
}
