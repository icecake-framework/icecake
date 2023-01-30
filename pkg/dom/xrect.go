package dom

import (
	"syscall/js"

	"github.com/sunraylab/icecake/pkg/lib"
)

/****************************************************************************
* DOMRect
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Rect
func RectFromJS(value js.Value) *lib.Rect {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &lib.Rect{}

	x := value.Get("x")
	ret.X = (x).Float()
	y := value.Get("x")
	ret.Y = (y).Float()
	width := value.Get("x")
	ret.Width = (width).Float()
	height := value.Get("x")
	ret.Height = (height).Float()
	return ret
}
