package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type ICKIcon struct {
	html.HTMLSnippet

	Text  string // optional text
	COLOR        // icon color
	SIZE         // icon size
}

// Ensuring ICKIcon implements the right interface
var _ html.ElementComposer = (*ICKIcon)(nil)

func (icon *ICKIcon) SetColor(c COLOR) *ICKIcon {
	icon.COLOR = c
	return icon
}
func (icon *ICKIcon) SetSize(s SIZE) *ICKIcon {
	icon.SIZE = s
	return icon
}

/******************************************************************************/

// Tag Builder used by the rendering functions.
func (icon *ICKIcon) BuildTag() html.Tag {
	icon.Tag().SetTagName("span").
		AddClassIf(icon.Text == "", "icon").
		AddClassIf(icon.Text != "", "icon-text")
	return *icon.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (msg *ICKIcon) RenderContent(out io.Writer) error {

	return nil
}
