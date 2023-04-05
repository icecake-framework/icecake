package js

import (
	"fmt"
	"reflect"
	"syscall/js"

	"github.com/sunraylab/icecake/pkg/console"
)

// TYPE represents the JavaScript type of a Value.
type TYPE int

// Constants that enumerates the JavaScript types.
const (
	TYPE_UNDEFINED TYPE = TYPE(js.TypeUndefined)
	TYPE_NULL      TYPE = TYPE(js.TypeNull)
	TYPE_BOOLEAN   TYPE = TYPE(js.TypeBoolean)
	TYPE_NUMBER    TYPE = TYPE(js.TypeNumber)
	TYPE_STRING    TYPE = TYPE(js.TypeString)
	TYPE_SYMBOL    TYPE = TYPE(js.TypeSymbol)
	TYPE_OBJECT    TYPE = TYPE(js.TypeObject)
	TYPE_FUNCTION  TYPE = TYPE(js.TypeFunction)
)

func (t TYPE) String() string {
	switch t {
	case TYPE_NULL:
		return "Null"
	case TYPE_BOOLEAN:
		return "Boolean"
	case TYPE_NUMBER:
		return "Number"
	case TYPE_STRING:
		return "String"
	case TYPE_SYMBOL:
		return "Symbol"
	case TYPE_OBJECT:
		return "Object"
	case TYPE_FUNCTION:
		return "Function"
	}
	return "Undefined"
}

// JSValueProvider is implemented by types that are backed by a JavaScript value.
type JSValueProvider interface {
	Value() JSValue
}

type JSValueWrapper interface {
	Wrap(JSValueProvider)
}

// JSValue represents a JavaScript value. On wasm architecture,
// it wraps the JSValue from https://golang.org/pkg/syscall/js/ package.
type JSValue struct {
	jsvalue js.Value
}

// Global returns the JavaScript global object, usually "window" or "global".
func Global() JSValue {
	return JSValue{jsvalue: js.Global()}
}

// New uses JavaScript's "new" operator with value v as constructor and the given arguments.
// It panics if v is not a JavaScript function.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (_v JSValue) New(args ...any) JSValue {
	args = cleanArgs(args...)
	return val(_v.jsvalue.New(args...))
}

func (_v *JSValue) Wrap(_jsvp JSValueProvider) {
	if _v.jsvalue.Truthy() {
		console.Warnf("wrapping an already wrapped jsvalue")
	}
	_v.jsvalue = _jsvp.Value().jsvalue
}

// Value returns the value itself for JSValueProvider interface
func (_v JSValue) Value() JSValue {
	return _v
}

// ValueOf returns x as a JavaScript value:
//
//	| Go                     | JavaScript             |
//	| ---------------------- | ---------------------- |
//	| js.Value               | [its value]            |
//	| js.Func                | function               |
//	| nil                    | null                   |
//	| bool                   | boolean                |
//	| integers and floats    | number                 |
//	| string                 | string                 |
//	| []interface{}          | new array              |
//	| map[string]interface{} | new object             |
//
// Panics if x is not one of the expected types.
func ValueOf(x any) (_jsv JSValue) {
	defer func() {
		if r := recover(); r != nil {
			console.Panicf("JS%s", r)
		}
	}()

	switch x := x.(type) {
	case JSValue:
		_jsv.jsvalue = js.ValueOf(x.jsvalue)
	// case JSFunc:
	// 	_jsv.jsvalue = js.ValueOf(x.Value)
	default:
		_jsv.jsvalue = js.ValueOf(x)
	}
	return _jsv
}

// Type returns the JavaScript type of the value v. It is similar to JavaScript's typeof operator,
// except that it returns TypeNull instead of TypeObject for null.
func (_v JSValue) Type() TYPE {
	return TYPE(_v.jsvalue.Type())
}

// IsDefined returns true if js value is not null nor undefined
func (_v JSValue) IsDefined() bool {
	return _v.Type() != TYPE_NULL && _v.Type() != TYPE_UNDEFINED
}

// IsObject returns true if js value is of type Object
func (_v JSValue) IsObject() bool {
	return _v.Type() == TYPE_OBJECT
}

// Equal reports whether v and w are equal according to JavaScript's === operator.
func (_v JSValue) Equal(w JSValue) bool {
	return _v.jsvalue.Equal(w.jsvalue)
}

// Delete deletes the JavaScript property p of value v.
// It panics if v is not a JavaScript object.
func (_v JSValue) Delete(p string) {
	_v.jsvalue.Delete(p)
}

// Call does a JavaScript call to the method m of value v with the given arguments.
// It panics if v has no method m. Returns js.null if _v is undefined.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (_v JSValue) Call(m string, args ...any) JSValue {
	if !_v.IsDefined() {
		console.Warnf("unable to call %q: undefined js value\n", m)
		return null()
	}
	args = cleanArgs(args...)
	defer func() {
		if r := recover(); r != nil {
			console.Panicf("JSValue: %s", r)
		}
	}()
	// DEBUG: console.Warnf("F:%s P:%+v", m, args)
	res := _v.jsvalue.Call(m, args...)
	return val(res)
}

func makeArgs(args []any) {
	for _, arg := range args {
		fmt.Println("makeargs:ValueOf =", reflect.TypeOf(arg).String())
	}
}

// Get returns the JavaScript property p of value v.
// Returns js.null if _v is undefined or unable to get p.
// Print a warning in these cases.
func (_v JSValue) Get(_pname string) JSValue {
	if !_v.IsDefined() {
		console.Warnf("unable to get %q: undefined js value\n", _pname)
		return null()
	}
	jsret := val(_v.jsvalue.Get(_pname))
	if !jsret.IsDefined() {
		console.Warnf("get %q returns an undefined js value\n", _pname)
	}
	return jsret
}

func TryGet(_param ...string) (JSValue, error) {
	tryArgs := make([]any, len(_param)+1)
	tryArgs[0] = js.Global()
	for i, v := range _param {
		tryArgs[i+1] = js.ValueOf(v)
	}
	arr := js.Global().Call("tryGet", tryArgs...)
	if arr1 := arr.Index(1); !arr1.Equal(js.Null()) {
		return JSValue{}, &js.Error{Value: arr1}
	}
	return JSValue{arr.Index(0)}, nil
}

// FIX: tryset
func TrySet(_param ...string) error {
	tryArgs := make([]any, len(_param)+1)
	tryArgs[0] = js.Global()
	for i, v := range _param {
		tryArgs[i+1] = js.ValueOf(v)
	}
	arr := js.Global().Call("trySet", tryArgs...)
	if arr1 := arr.Index(1); !arr1.Equal(js.Null()) {
		return &js.Error{Value: arr1}
	}
	return nil
}

// Check is like Get but returns an error without printing a warning in case of an error
func (_v JSValue) Check(_pname string) (JSValue, error) {
	if !_v.IsDefined() {
		return null(), fmt.Errorf("unable to get %q: undefined js value", _pname)
	}
	jsret := val(_v.jsvalue.Get(_pname))
	if !jsret.IsDefined() {
		return null(), fmt.Errorf("get %q returns an undefined js value\n", _pname)
	}
	return jsret, nil
}

// Set sets the JavaScript property p of value v to ValueOf(x).
// It panics if v is not a JavaScript object. Returns js.null if _v is undefined and print a warning.
func (_v JSValue) Set(p string, x any) {
	if !_v.IsDefined() {
		console.Warnf("unable to set a js property to an undefined js value\n")
		return
	}
	if wrapper, ok := x.(JSValue); ok {
		x = wrapper.jsvalue
	}
	defer func() {
		if r := recover(); r != nil {
			console.Panicf("JSSet: %s", r)
		}
	}()
	_v.jsvalue.Set(p, x)
}

// Index returns JavaScript index i of value v.
// It panics if v is not a JavaScript object. Returns js.null if _v is undefined and print a warning.
func (_v JSValue) Index(i int) JSValue {
	if !_v.IsDefined() {
		console.Warnf("unable to get Index: undefined js value\n")
		return null()
	}
	return val(_v.jsvalue.Index(i))
}

// Length returns the JavaScript property "length" of v.
// It panics if v is not a JavaScript object. Returns js.null if _v is undefined and print a warning.
func (_v JSValue) Length() int {
	if !_v.IsDefined() {
		console.Warnf("unable to get Length: undefined js value\n")
		return 0
	}
	return _v.jsvalue.Length()
}

func (_v JSValue) InstanceOf(t JSValue) bool {
	if !_v.IsDefined() {
		console.Warnf("unable to check InstanceOf: undefined js value\n")
		return false
	}
	return _v.jsvalue.InstanceOf(t.jsvalue)
}

func (_v JSValue) Invoke(args ...any) JSValue {
	if !_v.IsDefined() {
		console.Warnf("unable to Invoke: undefined js value\n")
		return null()
	}
	return val(_v.jsvalue.Invoke(args...))
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

	release = js.Func(then).Release
	_v.jsvalue.Call("then", then)
}

// GetObject returns the value of _v ensuring it's of type Object,
// otherwise returns null
func (_v JSValue) GetObject(_pname string) JSValue {
	get := _v.Get(_pname)
	if get.Type() != TYPE_OBJECT {
		console.Errorf("JSValue:GetObject failed: %q type is %s", _pname, get.Type().String())
		return null()
	}
	return get
}

// Float returns the value v as a float64.
// It panics if v is not a JavaScript number.
func (_v JSValue) Float() float64 {
	return _v.jsvalue.Float()
}

func (_v JSValue) GetFloat(_pname string) float64 {
	get := _v.Get(_pname)
	if get.Type() == TYPE_NUMBER {
		return get.Float()
	}
	if get.IsDefined() {
		console.Errorf("JSValue:GetFloat failed: %q type is %s", _pname, get.Type().String())
	}
	return 0.0
}

// Int returns the value v truncated to an int.
// It panics if v is not a JavaScript number.
func (_v JSValue) Int() int {
	return _v.jsvalue.Int()
}

func (_v JSValue) GetInt(_pname string) int {
	get := _v.Get(_pname)
	if get.Type() == TYPE_NUMBER {
		return get.Int()
	}
	if get.IsDefined() {
		console.Errorf("JSValue:GetInt failed: %q type is %s", _pname, get.Type().String())
	}
	return 0
}

// Bool returns the value v as a bool.
// It panics if v is not a JavaScript boolean.
func (_v JSValue) Bool() bool {
	return _v.jsvalue.Bool()
}

func (_v JSValue) GetBool(_pname string) bool {
	get := _v.Get(_pname)
	if get.Type() == TYPE_BOOLEAN {
		return get.Truthy()
	}
	if get.IsDefined() {
		console.Errorf("JSValue:GetBool failed: %q type is %s", _pname, get.Type().String())
	}
	return false
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
	if get.Type() == TYPE_STRING {
		return get.String()
	}
	if get.IsDefined() {
		console.Errorf("JSValue:GetString failed: %q type is %s", _pname, get.Type().String())
	}
	return ""
}

// Truthy returns the JavaScript "truthiness" of the value v. In JavaScript,
// false, 0, "", null, undefined, and NaN are "falsy", and everything else is
// "truthy". See https://developer.mozilla.org/en-US/docs/Glossary/Truthy.
func (_v JSValue) Truthy() bool {
	return _v.jsvalue.Truthy()
}

func null() JSValue {
	return val(js.Null())
}

/******************************************************************************
* JSFunction
******************************************************************************/

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
func FuncOf(fn func(this JSValue, args []JSValue) any) js.Func {
	f := js.FuncOf(func(this js.Value, args []js.Value) any {
		wargs := make([]JSValue, len(args))
		for i, a := range args {
			wargs[i] = val(a)
		}

		return fn(val(this), wargs)
	})
	return f
}

func val(_v js.Value) JSValue {
	return JSValue{jsvalue: _v}
}

func cleanArgs(args ...any) []any {
	for i, aa := range args {
		switch a := aa.(type) {
		case JSValue:
			args[i] = a.jsvalue
		default:
			args[i] = a
		}
	}
	return args
}

// func cleanArg(_v any) any {
// 	//DEBUG:
// 	fmt.Println(reflect.TypeOf(_v).String())

// 	switch v := _v.(type) {
// 	case map[string]any:
// 		m := make(map[string]any, len(v))
// 		for key, val := range v {
// 			m[key] = cleanArg(val)
// 		}
// 		return m

// 	case []any:
// 		// DEBUG:
// 		fmt.Println("claning a slice")
// 		s := make([]any, len(v))
// 		for i, val := range v {
// 			s[i] = cleanArgs(val)
// 		}
// 		return s

// 	case []string:
// 		// DEBUG:
// 		fmt.Println("claning a string slice")
// 		s := make([]string, len(v))
// 		for i, val := range v {
// 			s[i] = val
// 		}

// 		// case JSFunc:
// 	// 	return v.Value

// 	case JSValue:
// 		return v.jsvalue
// 	}

// 	return _v

// }
