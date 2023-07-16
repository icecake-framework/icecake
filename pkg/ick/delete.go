package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ickcore"
	"github.com/lolorenzo777/verbose"
)

func init() {
	ickcore.RegisterComposer("ick-delete", &ICKDelete{})
}

type DELETE_TYPE string

const (
	DLTTYP_BUTTON DELETE_TYPE = "button"
	DLTTYP_ANCHOR DELETE_TYPE = "anchor"
)

type ICKDelete struct {
	ickcore.BareSnippet

	DELETE_TYPE

	// The element id to remove from the DOM when the delete button is clicked
	TargetId string

	// styling
	SIZE
}

// Ensuring ICKDelete implements the right interface
var _ ickcore.ContentComposer = (*ICKDelete)(nil)
var _ ickcore.TagBuilder = (*ICKDelete)(nil)

func Delete(id string, targetid string) *ICKDelete {
	del := new(ICKDelete)
	del.Tag().SetId(id)
	del.TargetId = targetid
	if targetid == "" {
		verbose.Debug("Delete factory: TargetId missing")
	}
	return del
}

// SetSize set the size of the tag
func (t *ICKDelete) SetSize(s SIZE) *ICKDelete {
	t.SIZE = s
	return t
}

func (t *ICKDelete) SetType(typ DELETE_TYPE) *ICKDelete {
	t.DELETE_TYPE = typ
	return t
}

/******************************************************************************/

// BuildTag builds the tag used to render the html element.
// Delete tag is a simple <button class="delete"></delete>
func (del *ICKDelete) BuildTag() ickcore.Tag {
	del.Tag().
		AddClass("delete").
		SetAttribute("aria-label", "delete").
		SetAttributeIf(del.TargetId != "", "data-targetid", del.TargetId).
		PickClass(SIZE_OPTIONS, string(del.SIZE))

	if del.DELETE_TYPE == DLTTYP_ANCHOR {
		del.Tag().SetTagName("a")
	} else {
		del.Tag().SetTagName("button")
	}

	verbose.PrintfIf(del.TargetId == "", verbose.WARNING, "ICKDelete.BuildTag: missing TargetId\n")
	return *del.Tag()
}

// Delete rendering is made by the tag attributes.
func (del *ICKDelete) RenderContent(out io.Writer) error {
	return nil
}
