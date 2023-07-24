package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

type ICKIcon struct {
	ickcore.BareSnippet

	// Key is the icon key that will be added to the <class attribute of the <i> element.
	// Format of this key depends on the icon provider:
	// 	- 'fas fa-{iconname}' for [font awesome icons]
	// 	- 'bi-{iconname}' for [bootstrap icons]
	// 	- 'mdi mdi{iconname}' for [material design icons]
	//
	// If the Key is empty nothing is rendered
	//
	// [font awesome icons]: https://fontawesome.com/icons
	// [bootstrap icons]: https://icons.getbootstrap.com/
	// [material design icons]: https://pictogrammers.com/library/mdi/
	Key string

	Text  string   // optional text
	Color TXTCOLOR // icon color
}

// Ensuring ICKIcon implements the right interface
var _ ickcore.ContentComposer = (*ICKIcon)(nil)
var _ ickcore.TagBuilder = (*ICKIcon)(nil)

func Icon(key string, attrs ...string) *ICKIcon {
	// TODO: ICKIcon - check key validity
	i := &ICKIcon{Key: key}
	i.Tag().ParseAttributes(attrs...)
	return i
}

func (icon ICKIcon) Clone() *ICKIcon {
	c := new(ICKIcon)
	c.BareSnippet = *icon.BareSnippet.Clone()
	c.Key = icon.Key
	c.Text = icon.Text
	c.Color = icon.Color
	return c
}

func (icon *ICKIcon) SetText(t string) *ICKIcon {
	icon.Text = t
	return icon
}
func (icon *ICKIcon) SetColor(c TXTCOLOR) *ICKIcon {
	icon.Color = c
	return icon
}

/******************************************************************************/

func (icon *ICKIcon) NeedRendering() bool {
	return icon.Key != ""
}

// Tag Builder used by the rendering functions.
func (icon *ICKIcon) BuildTag() ickcore.Tag {
	icon.Tag().SetTagName("span").
		AddClassIf(icon.Text == "", "icon").
		AddClassIf(icon.Text != "", "icon-text").
		PickClass(TXTCOLOR_OPTIONS, string(icon.Color))
	return *icon.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (icon *ICKIcon) RenderContent(out io.Writer) error {
	if icon.Text == "" {
		ickcore.RenderString(out, `<i class="`, icon.Key, `"></i>`)
	} else {
		ickcore.RenderChild(out, icon,
			Elem("span", `class="icon"`, ickcore.ToHTML(`<i class="`+icon.Key+`"></i>`)))
		ickcore.RenderString(out, `<span>`+icon.Text+`</span>`)
	}
	return nil
}
