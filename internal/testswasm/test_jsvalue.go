package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunraylab/icecake/pkg/js"
)

func TestJSValue(t *testing.T) {

	// ValueOf
	i := 1
	jsi := js.ValueOf(i)
	assert.True(t, jsi.IsDefined())
	assert.True(t, jsi.Type() == js.TYPE_NUMBER)
	assert.True(t, jsi.Value().Int() == 1)

	var p *int
	assert.Panics(t, func() { js.ValueOf(p) }) // console error -->  JSValueOf: invalid value

	// // Set
	assert.Panics(t, func() { jsi.Set("test", "ok") }) // console error -->  JSSet: syscall/js: call of Value.Set on number

	browser := js.ValueOf(js.Global())
	browser.Set("test", "ok")

	// // Get
	assert.False(t, browser.Get("abc2").IsDefined()) // console warning --> get "abc2" returns an undefined js value
	assert.Equal(t, 0, browser.GetInt("abc3"))       // console warning --> get "abc3" returns an undefined js value
	assert.Greater(t, browser.GetFloat("devicePixelRatio"), 0.0)
	assert.Equal(t, "ok", browser.GetString("test"))

	// // Call
	assert.NotPanics(t, func() { browser.Call("focus") })
	assert.NotPanics(t, func() { browser.Call("scrollBy", 0, 0) })
	assert.Panics(t, func() { browser.Call("abc1") })

	e := browser.Get("document").Call("getElementById", "testid")
	assert.True(t, e.Type() == js.TYPE_OBJECT)

	res := browser.Call("testhello", 1)
	assert.Equal(t, "hello 1", res.String())

	// invoke
	f := browser.Get("testhello")
	assert.True(t, f.Type() == js.TYPE_FUNCTION)
	res = f.Invoke(2)
	assert.Equal(t, "hello 2", res.String())

	ff := js.FuncOf(func(this js.JSValue, args []js.JSValue) any {
		fmt.Println("js func called")
		return nil
	})
	browser.Call("test2", ff)

	// event
	browser.Call("addEventListener", "resize", js.FuncOf(onResize))

}

func onResize(this js.JSValue, args []js.JSValue) any {

	return nil
}
