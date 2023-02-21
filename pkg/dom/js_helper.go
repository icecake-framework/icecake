package dom

import (
	"fmt"
	"syscall/js"
)

// type Union struct {
// 	jsValue js.Value
// }

// func (u *Union) JSValue() js.Value {
// 	return u.jsValue
// }

// func UnionFromJS(value js.Value) *Union {
// 	return &Union{jsValue: value}
// }

type JSWrapper interface {
	JSValue() js.Value
	Wrap(js.Value)
}

func ICKLog(msg string, a ...any) {
	fmt.Printf(msg, a...)
}

func ICKError(msg string, a ...any) {
	str := fmt.Sprintf(msg, a...)
	js.Global().Call("ickError", str)
}

func ICKWarn(msg string, a ...any) {
	str := fmt.Sprintf(msg, a...)
	js.Global().Call("ickWarn", str)
}

func tryGet(v js.Value, p string) (result js.Value, err error) {
	defer func() {
		if x := recover(); x != nil {
			var ok bool
			if err, ok = x.(error); !ok {
				err = fmt.Errorf("%v", x)
			}
		}
	}()
	return v.Get(p), nil
}
