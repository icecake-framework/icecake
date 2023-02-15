package dom

import (
	"syscall/js"

	"github.com/sunraylab/icecake/pkg/lib"
)

/****************************************************************************
* DOMRect
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Rect
func CastRect(value js.Value) *lib.Rect {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	rect := new(lib.Rect)

	x := value.Get("x")
	rect.X = (x).Float()
	y := value.Get("y")
	rect.Y = (y).Float()
	width := value.Get("width")
	rect.Width = (width).Float()
	height := value.Get("height")
	rect.Height = (height).Float()

	return rect
}
