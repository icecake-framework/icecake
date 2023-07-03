package bulma

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-delete", &Delete{})
}

type Delete struct {
	html.HTMLSnippet

	// the element id to remove from the DOM when the delete button is clicked
	TargetID string

	// styling
	SIZE
}

// Ensure Delete implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*Delete)(nil)

// BuildTag builds the tag used to render the html element.
// Delete tag is a simple <button class="delete"></delete>
func (del *Delete) BuildTag() html.Tag {
	del.Tag().
		SetTagName("button").
		AddClass("delete").
		SetAttribute("aria-label", "delete").
		SetAttribute("data-TargetId", del.TargetID).
		PickClass(SIZE_OPTIONS, string(del.SIZE))
	return *del.Tag()
}

// Delete renders an empty content.
func (del *Delete) RenderContent(out io.Writer) error {
	return nil
}
