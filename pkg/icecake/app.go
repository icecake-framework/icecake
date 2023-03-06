package ick

/*****************************************************************************
* WebApp
******************************************************************************/

// WebApp
type WebApp struct {
	Document // The embedded DOM document

	browser Window
}

// NewWebApp is the WebApp factory. Must be call once at the begining of the wasm main code.
func NewWebApp() *WebApp {
	webapp := new(WebApp)
	webapp.browser.Wrap(getWindow())
	webapp.Document.Wrap(GetDocument())
	return webapp
}

// Browser returns the DOM.Window object
func (_app *WebApp) Browser() Window {
	return _app.browser
}
