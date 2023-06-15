package html

import (
	"github.com/icecake-framework/icecake/pkg/clock"
)

func init() {
	RegisterComposer("ick-delete", &Delete{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type Delete struct {
	HTMLSnippet

	// the element id to remove from the DOM
	TargetID string

	// The TargetID will be automatically removed after the clock Timeout duration if not zero.
	// The timer starts when the delete button is rendered (call to addlisteners).
	clock.Clock

	// OnDelete, if set, is called when the deletion occurs and after the targetId has been removed
	OnDelete func(*Delete)
}

// Ensure Delete implements HTMLComposer interface
var _ HTMLComposer = (*Delete)(nil)

func (del *Delete) Tag() *Tag {
	del.tag.SetName("button")
	del.tag.Attributes().SetClasses("delete").setAttribute("aria-label", "delete", true)
	return &del.tag
}
