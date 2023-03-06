package ick

import (
	"syscall/js"

	"github.com/sunraylab/icecake/pkg/errors"
)

// Type represents the JavaScript type of a Value.
type Type int

// Constants that enumerates the JavaScript types.
const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeNumber
	TypeString
	TypeSymbol
	TypeObject
	TypeFunction
)

// JSValueProvider is implemented by types that are backed by a JavaScript value.
type JSValueProvider interface {
	Value() JSValue
}

// JSValue represents a JavaScript value. On wasm architecture,
// it wraps the JSValue from https://golang.org/pkg/syscall/js/ package.
type JSValue struct {
	jsvalue js.Value
}

func (_v JSValue) Value() JSValue {
	return _v
}

func (_v *JSValue) Wrap(_jsvp JSValueProvider) {
	if _v.jsvalue.Truthy() {
		errors.ConsoleWarnf("wrapping an already wrapped element")
	}
	_v.jsvalue = _jsvp.Value().jsvalue
}

func (_v JSValue) Type() Type {
	return Type(_v.jsvalue.Type())
}

func (_v JSValue) Call(m string, args ...any) JSValue {
	if !_v.IsDefined() {
		return null()
	}
	args = cleanArgs(args...)
	return val(_v.jsvalue.Call(m, args...))
}

func (_v JSValue) Delete(p string) {
	_v.jsvalue.Delete(p)
}

func (_v JSValue) Equal(w JSValue) bool {
	return _v.jsvalue.Equal(w.jsvalue)
}

func (_v JSValue) Get(_pname string) JSValue {
	if !_v.IsDefined() {
		return null()
	}
	return val(_v.jsvalue.Get(_pname))
}

func (_v JSValue) GetObject(_pname string) JSValue {
	get := _v.Get(_pname)
	if get.Type() != TypeObject {
		errors.ConsoleErrorf("unable to get js object for %q", _pname)
		return null()
	}
	return get
}

// TODO: tryget
func TryGet(_v js.Value, p string) (result js.Value, err error) {
	defer func() {
		if x := recover(); x != nil {
			var ok bool
			if err, ok = x.(error); !ok {
				err = errors.ConsoleErrorf("%v", x)
			}
		}
	}()
	return _v.Get(p), nil
}

func (_v JSValue) Set(p string, x any) {
	if !_v.IsDefined() {
		errors.ConsoleLogf("unable to set a property to an undefined js value")
		return
	}
	if wrapper, ok := x.(JSValue); ok {
		x = wrapper.jsvalue
	}
	_v.jsvalue.Set(p, x)
}

func (_v JSValue) Index(i int) JSValue {
	return val(_v.jsvalue.Index(i))
}

func (_v JSValue) InstanceOf(t JSValue) bool {
	return _v.jsvalue.InstanceOf(t.jsvalue)
}

func (_v JSValue) Invoke(args ...any) JSValue {
	return val(_v.jsvalue.Invoke(args...))
}

func (_v JSValue) New(args ...any) JSValue {
	args = cleanArgs(args...)
	return val(_v.jsvalue.New(args...))
}

func (_v JSValue) Then(f func(JSValue)) {
	release := func() {}

	then := FuncOf(func(this JSValue, args []JSValue) any {
		var arg JSValue
		if len(args) > 0 {
			arg = args[0]
		}

		f(arg)
		release()
		return nil
	})

	release = then.Release
	_v.jsvalue.Call("then", then)
}

// Float returns the value v as a float64.
// It panics if v is not a JavaScript number.
func (_v JSValue) Float() float64 {
	return _v.jsvalue.Float()
}

func (_v JSValue) GetFloat(_pname string) float64 {
	defer func() {
		if x := recover(); x != nil {
			errors.ConsoleErrorf("A unable to get js float for %q", _pname)
		}
	}()
	get := _v.Get(_pname)
	if get.Type() != TypeNumber {
		errors.ConsoleErrorf("B unable to get js int for %q", _pname)
		return 0
	}
	return get.Float()
}

// Int returns the value v truncated to an int.
// It panics if v is not a JavaScript number.
func (_v JSValue) Int() int {
	return _v.jsvalue.Int()
}

func (_v JSValue) GetInt(_pname string) int {
	defer func() {
		if x := recover(); x != nil {
			errors.ConsoleErrorf("A unable to get js int for %q", _pname)
		}
	}()
	get := _v.Get(_pname)
	if get.Type() != TypeNumber {
		errors.ConsoleErrorf("B unable to get js int for %q", _pname)
		return 0
	}
	return get.Int()
}

// Bool returns the value v as a bool.
// It panics if v is not a JavaScript boolean.
func (_v JSValue) Bool() bool {
	return _v.jsvalue.Bool()
}

func (_v JSValue) GetBool(_pname string) bool {
	defer func() {
		if x := recover(); x != nil {
			errors.ConsoleErrorf("A unable to get js bool for %q", _pname)
		}
	}()
	get := _v.Get(_pname)
	if get.Type() != TypeNumber {
		errors.ConsoleErrorf("B unable to get js int bool %q", _pname)
		return false
	}
	return get.Bool()
}

// String returns the value v as a string.
// String is a special case because of Go's String method convention. Unlike the other getters,
// it does not panic if v's Type is not TypeString. Instead, it returns a string of the form "<T>"
// or "<T: V>" where T is v's type and V is a string representation of v's value.
func (_v JSValue) String() string {
	return _v.jsvalue.String()
}

func (_v JSValue) GetString(_pname string) string {
	get := _v.Get(_pname)
	if get.Type() != TypeString {
		errors.ConsoleErrorf("unable to get js string for %q", _pname)
		return ""
	}
	return get.String()
}

// Truthy returns the JavaScript "truthiness" of the value v. In JavaScript,
// false, 0, "", null, undefined, and NaN are "falsy", and everything else is
// "truthy". See https://developer.mozilla.org/en-US/docs/Glossary/Truthy.
func (_v JSValue) Truthy() bool {
	return _v.jsvalue.Truthy()
}

// IsObject returns true if js value is of type Object
func (_v JSValue) IsObject() bool {
	return _v.Type() == TypeObject
}

// IsDefined returns true if js value is not null nor undefined
func (_v JSValue) IsDefined() bool {
	return _v.Type() != TypeNull && _v.Type() != TypeUndefined
}

// Length returns the JavaScript property "length" of v.
// It panics if v is not a JavaScript object.
func (_v JSValue) Length() int {
	return _v.jsvalue.Length()
}

func null() JSValue {
	return val(js.Null())
}

func undefined() JSValue {
	return val(js.Undefined())
}

type JSFunction struct {
	JSValue
	fn js.Func
}

func (f JSFunction) Release() {
	f.fn.Release()
}

// FuncOf returns a wrapped function.
//
// Invoking the JavaScript function will synchronously call the Go function fn
// with the value of JavaScript's "this" keyword and the arguments of the
// invocation. The return value of the invocation is the result of the Go
// function mapped back to JavaScript according to ValueOf.
//
// A wrapped function triggered during a call from Go to JavaScript gets
// executed on the same goroutine. A wrapped function triggered by JavaScript's
// event loop gets executed on an extra goroutine. Blocking operations in the
// wrapped function will block the event loop. As a consequence, if one wrapped
// function blocks, other wrapped funcs will not be processed. A blocking
// function should therefore explicitly start a new goroutine.
//
// Func.Release must be called to free up resources when the function will not
// be used any more.
func FuncOf(fn func(this JSValue, args []JSValue) any) JSFunction {
	f := js.FuncOf(func(this js.Value, args []js.Value) any {
		wargs := make([]JSValue, len(args))
		for i, a := range args {
			wargs[i] = val(a)
		}

		return fn(val(this), wargs)
	})

	return JSFunction{
		JSValue: JSValue{jsvalue: f.Value},
		fn:      f,
	}
}

func val(_v js.Value) JSValue {
	return JSValue{jsvalue: _v}
}

func cleanArgs(args ...any) []any {
	for i, a := range args {
		args[i] = cleanArg(a)
	}
	return args
}

func cleanArg(_v any) any {
	switch v := _v.(type) {
	case map[string]any:
		m := make(map[string]any, len(v))
		for key, val := range v {
			m[key] = cleanArg(val)
		}
		return m

	case []any:
		s := make([]any, len(v))
		for i, val := range v {
			s[i] = cleanArgs(val)
		}

	case JSValue:
		return v.jsvalue
	}

	return _v

}
