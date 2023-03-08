package main

import (
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
	ick "github.com/sunraylab/icecake/pkg/icecake"
)

func TestJSValue(t *testing.T) {

	var i int
	jsi := ick.ValueOf(i)
	assert.True(t, jsi.IsDefined())
	assert.True(t, jsi.Type() == ick.TYPE_NUMBER)

	var p *int
	assert.Panics(t, func() { ick.ValueOf(p) })

	browser := ick.ValueOf(js.Global())
	assert.NotPanics(t, func() { browser.Call("focus") })
	assert.Panics(t, func() { browser.Call("abc1") })

	assert.False(t, browser.Get("abc2").IsDefined()) // console warning --> get "abc2" returns an undefined js value
	assert.Equal(t, 0, browser.GetInt("abc3"))       // console warning --> get "abc3" returns an undefined js value

	assert.Greater(t, browser.GetFloat("devicePixelRatio"), 0.0)
}
