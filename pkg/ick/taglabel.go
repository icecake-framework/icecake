package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

func init() {
	ickcore.RegisterComposer("ick-tag", &ICKTagLabel{})
}

type TAGLABEL_SIZE string

const (
	TAGLBLSZ_STD     TAGLABEL_SIZE = "is-normal"
	TAGLBLSZ_MEDIUM  TAGLABEL_SIZE = "is-medium"
	TAGLBLSZ_LARGE   TAGLABEL_SIZE = "is-large"
	TAGLBLSZ_OPTIONS string        = string(TAGLBLSZ_STD + " " + TAGLBLSZ_MEDIUM + " " + TAGLBLSZ_LARGE)
)

type ICKTagLabel struct {
	ickcore.BareSnippet

	Header string // optional text to render on the left side of the taglabel
	Text   string // the tag text

	CanDelete bool // set to true to display the delete button and allow user to delete the message

	IsRounded   bool
	HeaderColor COLOR
	TextColor   COLOR
	TAGLABEL_SIZE

	jointed bool // joint tag labels
}

// Ensuring ICKTag implements the right interface
var _ ickcore.ContentComposer = (*ICKTagLabel)(nil)
var _ ickcore.TagBuilder = (*ICKTagLabel)(nil)

func TagLabel(text string, c COLOR, attrs ...string) *ICKTagLabel {
	n := new(ICKTagLabel)
	n.Text = text
	n.TextColor = c
	n.Tag().ParseAttributes(attrs...)
	return n
}

func (t *ICKTagLabel) SetHeader(text string, c COLOR) *ICKTagLabel {
	t.Header = text
	t.HeaderColor = c
	return t
}

// SetSize set the size of the tag
func (t *ICKTagLabel) SetSize(s TAGLABEL_SIZE) *ICKTagLabel {
	t.TAGLABEL_SIZE = s
	return t
}

func (t *ICKTagLabel) SetRounded(f bool) *ICKTagLabel {
	t.IsRounded = f
	return t
}

func (t *ICKTagLabel) SetCanDelete(can bool) *ICKTagLabel {
	t.CanDelete = can
	return t
}

/******************************************************************************/

// BuildTag returns <span class="tag {classes}" {attributes}>
func (t *ICKTagLabel) BuildTag() ickcore.Tag {
	t.jointed = t.Header != "" && (t.Text != "" || t.CanDelete)
	if t.jointed {
		t.Tag().SetTagName("div").AddClass("tags", "has-addons")
		switch t.TAGLABEL_SIZE {
		case TAGLBLSZ_LARGE:
			t.Tag().PickClass("are-normal are-medium are-large", "are-large")
		case TAGLBLSZ_MEDIUM:
			t.Tag().PickClass("are-normal are-medium are-large", "are-medium")
		default:
			t.Tag().PickClass("are-normal are-medium are-large", "are-normal")
		}
	} else {
		t.Tag().SetTagName("span").AddClass("tag")
		t.Tag().
			SetClassIf(t.IsRounded, "is-rounded").
			PickClass(SIZE_OPTIONS, string(t.TAGLABEL_SIZE))
		if t.Text != "" {
			t.Tag().PickClass(COLOR_OPTIONS, string(t.TextColor))
		} else {
			t.Tag().PickClass(COLOR_OPTIONS, string(t.HeaderColor))
		}
	}

	return *t.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (t *ICKTagLabel) RenderContent(out io.Writer) error {
	if t.jointed {
		if t.Header != "" {
			cmp := Elem("span", `class="tag"`, ickcore.ToHTML(t.Header))
			cmp.Tag().PickClass(COLOR_OPTIONS, string(t.HeaderColor))
			ickcore.RenderChild(out, t, cmp)
		}
		if t.Text != "" {
			cmp := Elem("span", `class="tag"`, ickcore.ToHTML(t.Text))
			cmp.Tag().PickClass(COLOR_OPTIONS, string(t.TextColor))
			ickcore.RenderChild(out, t, cmp)
		}
		if t.CanDelete {
			// btndel := Delete(t.Tag().SubId("btndel"), t.Tag().Id()).SetType(DLTTYP_ANCHOR)
			btndel := Elem("a", `class="tag is-delete" id="`+t.Tag().SubId("btndel")+`" data-targetid="`+t.Tag().Id()+`"`)
			ickcore.RenderChild(out, t, btndel)
		}

	} else {
		if t.Text != "" {
			ickcore.RenderString(out, t.Text)
		} else if t.Header != "" {
			ickcore.RenderString(out, t.Header)
		}
		if t.CanDelete {
			btndel := Delete(t.Tag().SubId("btndel"), t.Tag().Id())
			switch t.TAGLABEL_SIZE {
			case TAGLBLSZ_LARGE:
				btndel.SetSize(SIZE_STD)
			default:
				btndel.SetSize(SIZE_SMALL)
			}
			ickcore.RenderChild(out, t, btndel)
		}
	}
	return nil
}
