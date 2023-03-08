package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	ick "github.com/sunraylab/icecake/pkg/icecake"
)

func TestWindow(t *testing.T) {

	win := ick.GetWindow()

	fmt.Println("URL:", win.URL().String())

	assert.False(t, win.Closed())
	assert.True(t, win.OnLine())

	fmt.Println("cookie:", win.CookieEnabled())
	fmt.Println("history len:", win.History().Count())
	fmt.Println("User Agent :", win.UserAgent())
	fmt.Println("language   :", win.Language())

	w, h := win.InnerSize()
	assert.Greater(t, w, 0)
	assert.Greater(t, h, 0)
	fmt.Println("inner size :", w, h)
	w, h = win.OuterSize()
	assert.Greater(t, w, 0)
	assert.Greater(t, h, 0)
	fmt.Println("outer size :", w, h)

	x, y := win.ScrollPos()
	fmt.Println("scroll pos :", x, y)
}
