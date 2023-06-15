package html

import (
	"io"
	"net/url"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	RegisterComposer("ick-image", &Image{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type IMG_SIZE string

const (
	IMG_16x16   IMG_SIZE = "16x16"
	IMG_24x24   IMG_SIZE = "24x24"
	IMG_32x32   IMG_SIZE = "32x32"
	IMG_48x48   IMG_SIZE = "48x48"
	IMG_64x64   IMG_SIZE = "64x64"
	IMG_96x96   IMG_SIZE = "96x96"
	IMG_128x128 IMG_SIZE = "128x128"

	IMG_SQUARE IMG_SIZE = "square"
	IMG_1by1   IMG_SIZE = "1by1"
	IMG_5by4   IMG_SIZE = "5by4"
	IMG_4by3   IMG_SIZE = "4by3"
	IMG_3by2   IMG_SIZE = "3by2"
	IMG_5by3   IMG_SIZE = "5by3"
	IMG_16by9  IMG_SIZE = "16by9"
	IMG_2by1   IMG_SIZE = "2by1"
	IMG_3by1   IMG_SIZE = "3by1"
	IMG_4by5   IMG_SIZE = "4by5"
	IMG_3by4   IMG_SIZE = "3by4"
	IMG_2by3   IMG_SIZE = "2by3"
	IMG_3by5   IMG_SIZE = "3by5"
	IMG_9by16  IMG_SIZE = "9by16"
	IMG_1by2   IMG_SIZE = "1by2"
	IMG_1by3   IMG_SIZE = "1by3"
)

// Image is a typical img element embedded into a figure container specifying the image size
type Image struct {
	HTMLSnippet

	Src       *url.URL   // the url for the source of the image
	Alt       HTMLString // the alternative text
	Size      IMG_SIZE   // the size or the ratio of the image
	IsRounded bool       // Rounded image style
}

// Ensure Image implements HTMLComposer interface
var _ HTMLComposer = (*Image)(nil)

func NewImage(_rawUrl string, _size IMG_SIZE) *Image {
	img := new(Image)
	img.Size = _size
	img.ParseSrcURL(_rawUrl)
	return img
}

// TryParseSrc parses _rawurl to _img.Src and returns _img to allow chaining.
// Parsing errors are ignored and if any Src may stay nil.
func (_img *Image) ParseSrcURL(_rawUrl string) (_err error) {
	_img.Src, _err = url.Parse(_rawUrl)
	return
}

func (img *Image) Tag() *Tag {
	img.tag.SetName("figure")

	var imgsize string
	if img.Size != "" {
		imgsize = "is-" + string(img.Size)
	}
	img.tag.Attributes().SetClasses("image " + imgsize)

	return &img.tag
}

// Body returns any HTML string to unfold inside the html element
func (_img *Image) WriteBody(out io.Writer) error {
	// func (_img *Image) BodyTemplate() (body HTMLString) {
	var htmlRounded string
	if _img.IsRounded {
		htmlRounded = `class="is-rounded" `
	}
	var src string
	if _img.Src != nil {
		src = _img.Src.String()
	}
	// HACK: use NewSnippet
	WriteString(out, `<img `)
	WriteString(out, htmlRounded)
	WriteString(out, `role="img" focusable="false" `)
	WriteStringsIf(src != "", out, `src="`, src, `" `)
	WriteStringsIf(_img.Alt != "", out, `alt="`, string(_img.Alt), `" `)
	WriteString(out, `>`)
	return nil
}

// func (_img *Image) Template(*DataState) (_t SnippetTemplate) {
// 	_t.TagName = "figure"
// 	var imgsize String
// 	if _img.Size != "" {
// 		imgsize = String("is-" + _img.Size)
// 	}
// 	_t.Attributes = `class="image ` + imgsize + `"`

// 	var htmlRounded String
// 	if _img.IsRounded {
// 		htmlRounded = `class="is-rounded"`
// 	}
// 	var src String
// 	if _img.Src != nil {
// 		src = String(_img.Src.String())
// 	}
// 	_t.Body = `<img ` + htmlRounded + ` role="img" focusable="false" src="` + src + `" alt="` + src + `">`
// 	return _t
// }
