package webclientsdk

import "syscall/js"

/****************************************************************************
* DOMRect
*****************************************************************************/

// https://developer.mozilla.org/en-US/docs/Web/API/Rect
type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

// JSValue returns the js.Value or js.Null() if _this is nil
func (_this *Rect) JSValue() js.Value {
	if _this == nil {
		return js.Null()
	}

	jsValue := js.Value{}
	jsValue.Set("x", _this.X)
	jsValue.Set("y", _this.Y)
	jsValue.Set("width", _this.Width)
	jsValue.Set("height", _this.Height)
	return jsValue
}

// DOMRectFromJS is casting a js.Value into DOMRect.
func DOMRectFromJS(value js.Value) *Rect {
	if typ := value.Type(); typ == js.TypeNull || typ == js.TypeUndefined {
		return nil
	}
	ret := &Rect{}

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
