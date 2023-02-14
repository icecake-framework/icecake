package dom

import (
	"fmt"
	"syscall/js"
)

type Union struct {
	jsValue js.Value
}

func (u *Union) JSValue() js.Value {
	return u.jsValue
}

func UnionFromJS(value js.Value) *Union {
	return &Union{jsValue: value}
}

func ConsoleLog(msg string, a ...any) {
	fmt.Printf(msg, a...)
}

func ConsoleError(msg string, a ...any) {
	str := fmt.Sprintf(msg, a...)
	js.Global().Call("consoleError", str)
}

func ConsoleWarn(msg string, a ...any) {
	str := fmt.Sprintf(msg, a...)
	js.Global().Call("consoleWarn", str)
}
