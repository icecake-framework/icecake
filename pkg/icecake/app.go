package ick

var GData map[string]any = make(map[string]any, 0)

/*****************************************************************************
* WebApp
******************************************************************************/

var (
	app *WebApp
)

func init() {
	app = NewWebApp()
}

func App() *WebApp {
	return app
}

// WebApp
type WebApp struct {
	*Document
	win *Window
}

func NewWebApp() *WebApp {
	webapp := new(WebApp)
	webapp.win = getWindow()
	webapp.Document = getDocument()
	return webapp
}

func (_app *WebApp) Browser() *Window {
	return _app.win
}
