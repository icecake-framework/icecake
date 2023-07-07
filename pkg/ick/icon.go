package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type ICKIcon struct {
	html.HTMLSnippet

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
var _ html.ElementComposer = (*ICKIcon)(nil)

func Icon(key string, attrs ...string) *ICKIcon {
	i := &ICKIcon{Key: key}
	i.Tag().ParseAttributes(attrs...)
	return i
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

// Tag Builder used by the rendering functions.
func (icon *ICKIcon) BuildTag() html.Tag {
	if icon.Key == "" {
		return *html.NewTag("", nil)
	}
	icon.Tag().SetTagName("span").
		AddClassIf(icon.Text == "", "icon").
		AddClassIf(icon.Text != "", "icon-text").
		PickClass(TXTCOLOR_OPTIONS, string(icon.Color))
	return *icon.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (icon *ICKIcon) RenderContent(out io.Writer) error {
	if icon.Key == "" {
		return nil
	}
	if icon.Text == "" {
		html.WriteString(out, `<i class="`, icon.Key, `"></i>`)
	} else {
		s := html.Snippet("span", `class="icon"`).AddContent(html.ToHTML(`<i class="` + icon.Key + `"></i>`))
		html.Render(out, nil, s)
		html.WriteString(out, `<span>`+icon.Text+`</span>`)
	}
	return nil
}
