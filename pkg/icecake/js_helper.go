package ick

import (
	"fmt"
	"syscall/js"
)

type JSWrapper interface {
	JSValue() js.Value
	Wrap(js.Value)
}

func ConsoleLogf(msg string, a ...any) {
	fmt.Printf(msg, a...)
}

func ConsoleErrorf(msg string, a ...any) {
	str := fmt.Sprintf(msg, a...)
	js.Global().Call("ickError", str)
}

func ConsoleWarnf(msg string, a ...any) {
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
