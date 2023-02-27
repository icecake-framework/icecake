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
	Document
	browser Window
}

func NewWebApp() *WebApp {
	webapp := new(WebApp)
	webapp.browser.Wrap(getWindow().JSValue())
	webapp.Document.Wrap(getDocument().JSValue())
	return webapp
}

func (_app *WebApp) Browser() Window {
	return _app.browser
}

func (_app *WebApp) Close() {
	_app.Document.RemoveListeners()
	_app.browser.RemoveListeners()
}
