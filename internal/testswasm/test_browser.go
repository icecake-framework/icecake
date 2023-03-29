package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/sunraylab/icecake/pkg/browser"
)

func TestBrowser(t *testing.T) {

	t.Run("browser", func(t *testing.T) {
		win := browser.Win()
		require.True(t, win.IsDefined())
		assert.True(t, browser.OnLine())

		fmt.Println("cookie:", browser.CookieEnabled())
		fmt.Println("history len:", browser.SessionHistory().Count())
		fmt.Println("User Agent :", browser.UserAgent())
		fmt.Println("language   :", browser.Language())

		// console output must look like this:
		//
		// wasm_exec.js:22 cookie: true
		// wasm_exec.js:22 history len: 1
		// wasm_exec.js:22 User Agent : Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36
		// wasm_exec.js:22 language   : fr-FR
	})

	t.Run("window", func(t *testing.T) {

		win := browser.Win()
		assert.False(t, win.Closed())
		assert.True(t, browser.OnLine())

		fmt.Println("URL:", win.URL().String())

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

		// console output must look like this:
		//
		// URL: http://127.0.0.1:5510/
		// wasm_exec.js:22 inner size : 1200 667
		// wasm_exec.js:22 outer size : 1200 1880
		// wasm_exec.js:22 scroll pos : 0 0
	})
}
