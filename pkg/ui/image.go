package ui

import (
	"net/url"

	"github.com/sunraylab/icecake/pkg/dom"
	"github.com/sunraylab/icecake/pkg/html"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	html.RegisterComposer("ick-image", &Image{}, []string{"https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
}

type IMG_SIZE string

const (
	IMG_16x16   IMG_SIZE = "is-16x16"
	IMG_24x24   IMG_SIZE = "is-24x24"
	IMG_32x32   IMG_SIZE = "is-32x32"
	IMG_48x48   IMG_SIZE = "is-48x48"
	IMG_64x64   IMG_SIZE = "is-64x64"
	IMG_96x96   IMG_SIZE = "is-96x96"
	IMG_128x128 IMG_SIZE = "is-128x128"
	IMG_SQUARE  IMG_SIZE = "is-square"
	IMG_1by1    IMG_SIZE = "is-1by1"
	IMG_5by4    IMG_SIZE = "is-5by4"
	IMG_4by3    IMG_SIZE = "is-4by3"
	IMG_3by2    IMG_SIZE = "is-3by2"
	IMG_5by3    IMG_SIZE = "is-5by3"
	IMG_16by9   IMG_SIZE = "is-16by9"
	IMG_2by1    IMG_SIZE = "is-2by1"
	IMG_3by1    IMG_SIZE = "is-3by1"
	IMG_4by5    IMG_SIZE = "is-4by5"
	IMG_3by4    IMG_SIZE = "is-3by4"
	IMG_2by3    IMG_SIZE = "is-2by3"
	IMG_3by5    IMG_SIZE = "is-3by5"
	IMG_9by16   IMG_SIZE = "is-9by16"
	IMG_1by2    IMG_SIZE = "is-1by2"
	IMG_1by3    IMG_SIZE = "is-1by3"
)

type Image struct {
	dom.UISnippet

	Src       *url.URL // the url for the source of the image
	Size      IMG_SIZE // the size or the ratio of the image
	IsRounded bool     // Rounded image style
}

func (_img *Image) Template(*html.DataState) (_t html.SnippetTemplate) {
	_t.TagName = "figure"
	_t.Attributes = html.String(`class="image ` + _img.Size + `"`)

	var htmlRounded html.String
	if _img.IsRounded {
		htmlRounded = `class="is-rounded"`
	}
	var src html.String
	if _img.Src != nil {
		src = html.String(_img.Src.String())
	}
	_t.Body = `<img ` + htmlRounded + `src="` + src + `">`
	return _t
}
