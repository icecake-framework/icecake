package bulma

import (
	"github.com/icecake-framework/icecake/pkg/clock"
	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-delete", &Delete{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
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

// Ensure Delete implements HTMLComposer interface
var _ html.HTMLComposer = (*Delete)(nil)

// BuildTag builds the tag used to render the html element.
func (del *Delete) BuildTag(tag *html.Tag) {
	tag.SetTagName("button")
	tag.AddClasses("delete").SetAttribute("aria-label", "delete")
}
