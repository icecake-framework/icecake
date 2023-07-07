package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/lolorenzo777/verbose"
)

func init() {
	html.RegisterComposer("ick-delete", &ICKDelete{})
}

type ICKDelete struct {
	html.BareSnippet

	// The element id to remove from the DOM when the delete button is clicked
	TargetId string

	// styling
	SIZE
}

// Ensuring ICKDelete implements the right interface
var _ html.ElementComposer = (*ICKDelete)(nil)

func Delete(targetid string) *ICKDelete {
	del := new(ICKDelete)
	del.TargetId = targetid
	return del
}

// BuildTag builds the tag used to render the html element.
// Delete tag is a simple <button class="delete"></delete>
func (del *ICKDelete) BuildTag() html.Tag {
	del.Tag().
		SetTagName("button").
		AddClass("delete").
		SetAttribute("aria-label", "delete").
		SetAttributeIf(del.TargetId != "", "data-targetid", del.TargetId).
		PickClass(SIZE_OPTIONS, string(del.SIZE))
	verbose.PrintfIf(del.TargetId == "", verbose.WARNING, "ICKDelete.BuildTag: missing TargetId\n")
	return *del.Tag()
}

// Delete rendering is made by the tag attributes.
func (del *ICKDelete) RenderContent(out io.Writer) error {
	return nil
}
