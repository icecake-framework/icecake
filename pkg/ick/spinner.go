package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

type ICKSpinner struct {
	ickcore.BareSnippet
}

// Ensuring ICKSpinner implements the right interface
var _ ickcore.ContentComposer = (*ICKSpinner)(nil)
var _ ickcore.TagBuilder = (*ICKSpinner)(nil)

func Spinner() *ICKSpinner {
	return new(ICKSpinner)
}

// Tag Builder used by the rendering functions.
func (spin *ICKSpinner) BuildTag() ickcore.Tag {
	spin.Tag().SetTagName("span").AddClass("button has-border-none is-loading").AddStyle("width:100%;border-width:0;")
	return *spin.Tag()
}

// RenderContent writes the ickcore string corresponding to the content of the ickcore element.
func (*ICKSpinner) RenderContent(out io.Writer) error {
	return nil
}
