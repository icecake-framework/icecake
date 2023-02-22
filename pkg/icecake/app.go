package ick

var GData map[string]any = make(map[string]any, 0)

/*****************************************************************************
* ICKWebApp
******************************************************************************/

// ICWebApp
type ICWebApp struct {
	*Window
	*Document
}

func NewWebApp() *ICWebApp {
	webapp := new(ICWebApp)
	webapp.Window = GetWindow()
	webapp.Document = GetDocument()
	return webapp
}
