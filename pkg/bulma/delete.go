package bulma

import (
	"github.com/icecake-framework/icecake/pkg/clock"
	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-delete", &Delete{})
}

type Delete struct {
	html.HTMLSnippet

	// the element id to remove from the DOM
	TargetID string

	// The TargetID will be automatically removed after the clock Timeout duration if not zero.
	// The timer starts when the delete button is rendered (call to addlisteners).
	clock.Clock

	// OnDelete, if set, is called when the deletion occurs and after the targetId has been removed
	OnDelete func(*Delete)
}

// Ensure Delete implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*Delete)(nil)

// BuildTag builds the tag used to render the html element.
// Delete tag is a simple <button class="delete"></delete>
func (del *Delete) BuildTag(tag *html.Tag) {
	tag.SetTagName("button").AddClass("delete").SetAttribute("aria-label", "delete")
}
