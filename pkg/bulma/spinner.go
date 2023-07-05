package bulma

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type ICKSpinner struct{ html.HTMLSnippet }

// Ensure Hero implements HTMLTagComposer interface
var _ html.HTMLComposer = (*ICKSpinner)(nil)

func Spinner() *ICKSpinner {
	return new(ICKSpinner)
}

// Tag Builder used by the rendering functions.
func (spin *ICKSpinner) BuildTag() html.Tag {
	spin.Tag().SetTagName("span").AddClass("button has-border-none is-loading").AddStyle("width:100%;border-width:0;")
	return *spin.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (*ICKSpinner) RenderContent(out io.Writer) error {
	return nil
}
