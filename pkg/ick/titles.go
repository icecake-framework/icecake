package ick

import (
	"io"
	"strconv"

	"github.com/icecake-framework/icecake/pkg/html"
)

// ICKTitle is an icecake snippet providing the HTML rendering for a [bulma title].
// The title is an HTMLSnippet. Use AddContent to setup the content of the body of the title element.
//
// [bulma title]: https://bulma.io/documentation/elements/title/
type ICKTitle struct {
	html.BareSnippet

	Title string

	// IsSubtitle allows to render the title with the class subtitle rather than title.
	IsSubtitle bool

	// Heading allows to render the title within a <h> tag, otherwise it's in a <p> tag.
	Heading bool

	// Size define the size of the title. Bounded between 1 and 6.
	Size int
}

// Ensuring ICKTitle implements the right interface
var _ html.ElementComposer = (*ICKTitle)(nil)

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

// BuildTag returns tag <div class="message {classes}" {attributes}>
func (msg *ICKTitle) BuildTag() html.Tag {
	if msg.Size < 1 {
		msg.Size = 1
	} else if msg.Size > 6 {
		msg.Size = 6
	}
	ssiz := strconv.Itoa(msg.Size)

	if msg.Heading {
		msg.Tag().SetTagName("h" + ssiz)
	} else {
		msg.Tag().SetTagName("p")
	}

	if !msg.IsSubtitle {
		msg.Tag().AddClass("title")
	} else {
		msg.Tag().AddClass("subtitle")
	}
	msg.Tag().AddClass("is-" + ssiz)

	return *msg.Tag()
}

// BuildTag returns tag <div class="message {classes}" {attributes}>
func (msg *ICKTitle) RenderContent(out io.Writer) error {
	_, err := html.RenderString(out, msg.Title)
	return err
}
