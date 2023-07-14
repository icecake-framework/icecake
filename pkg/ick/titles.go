package ick

import (
	"io"
	"strconv"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

// ICKTitle is an icecake snippet providing the HTML rendering for a [bulma title].
// The title is an HTMLSnippet. Use AddContent to setup the content of the body of the title element.
//
// [bulma title]: https://bulma.io/documentation/elements/title/
type ICKTitle struct {
	ickcore.BareSnippet

	Title string

	// IsSubtitle allows to render the title with the class subtitle rather than title.
	IsSubtitle bool

	// Heading allows to render the title within a <h> tag, otherwise it's in a <p> tag.
	Heading bool

	// Size define the size of the title. Bounded between 1 and 6.
	Size int
}

// Ensuring ICKTitle implements the right interface
var _ ickcore.ContentComposer = (*ICKTitle)(nil)
var _ ickcore.TagBuilder = (*ICKTitle)(nil)

func Title(size int, title string, attrs ...string) *ICKTitle {
	msg := new(ICKTitle)
	msg.Tag().ParseAttributes(attrs...)
	msg.Heading = true
	msg.Size = size
	msg.IsSubtitle = false
	msg.Title = title
	return msg
}

func SubTitle(size int, title string, attrs ...string) *ICKTitle {
	msg := Title(size, title, attrs...)
	msg.IsSubtitle = true
	return msg
}

func (t *ICKTitle) BuildTag() ickcore.Tag {
	if t.Size < 1 {
		t.Size = 1
	} else if t.Size > 6 {
		t.Size = 6
	}
	ssiz := strconv.Itoa(t.Size)

	if t.Heading {
		t.Tag().SetTagName("h" + ssiz)
	} else {
		t.Tag().SetTagName("p")
	}

	if !t.IsSubtitle {
		t.Tag().AddClass("title")
	} else {
		t.Tag().AddClass("subtitle")
	}
	t.Tag().AddClass("is-" + ssiz)

	return *t.Tag()
}

func (t *ICKTitle) RenderContent(out io.Writer) error {
	_, err := ickcore.RenderString(out, t.Title)
	return err
}
