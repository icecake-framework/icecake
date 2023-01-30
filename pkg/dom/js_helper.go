package dom

import "syscall/js"

type Union struct {
	jsValue js.Value
}

func (u *Union) JSValue() js.Value {
	return u.jsValue
}

func UnionFromJS(value js.Value) *Union {
	return &Union{jsValue: value}
}
