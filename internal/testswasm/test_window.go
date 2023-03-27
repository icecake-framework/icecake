package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunraylab/icecake/pkg/browser"
)

func TestWindow(t *testing.T) {

	win := browser.Win()
	fmt.Println("URL:", win.URL().String())
	assert.False(t, win.Closed())

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

	assert.True(t, browser.OnLine())

	fmt.Println("cookie:", browser.CookieEnabled())
	fmt.Println("history len:", browser.SessionHistory().Count())
	fmt.Println("User Agent :", browser.UserAgent())
	fmt.Println("language   :", browser.Language())
}
