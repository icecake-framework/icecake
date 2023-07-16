package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

func init() {
	ickcore.RegisterComposer("ick-tag", &ICKTagLabel{})
}

type ICKMedia struct {
	ickcore.BareSnippet

	Left   ickcore.TagProvider
	Middle ickcore.TagProvider
	Right  ickcore.TagProvider
}

// Ensuring ICKTag implements the right interface
var _ ickcore.ContentComposer = (*ICKMedia)(nil)
var _ ickcore.TagBuilder = (*ICKMedia)(nil)

func Media(left ickcore.TagProvider, middle ickcore.TagProvider, right ickcore.TagProvider, attrs ...string) *ICKMedia {
	n := new(ICKMedia)
	n.Left = left
	n.Middle = middle
	n.Right = right
	n.Tag().ParseAttributes(attrs...)
	return n
}

/******************************************************************************/

// BuildTag returns <span class="tag {classes}" {attributes}>
func (m *ICKMedia) BuildTag() ickcore.Tag {
	m.Tag().SetTagName("div").AddClass("media")
	return *m.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (t *ICKMedia) RenderContent(out io.Writer) error {
	if t.Left != nil {
		t.Left.Tag().AddClass("media-left")
		ickcore.RenderChild(out, t, t.Left)
	}

	if t.Middle != nil {
		t.Middle.Tag().AddClass("media-content")
		ickcore.RenderChild(out, t, t.Middle)
	}

	if t.Right != nil {
		t.Right.Tag().AddClass("media-right")
		ickcore.RenderChild(out, t, t.Right)
	}
	return nil
}
