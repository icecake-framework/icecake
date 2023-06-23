package bulma

import (
	"io"
	"net/url"

	"github.com/icecake-framework/icecake/pkg/html"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	html.RegisterComposer("ick-image", &Image{})
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

// Bulma Image is a typical img element embedded into a figure container specifying the image size
type Image struct {
	html.HTMLSnippet

	Src       *url.URL // the url for the source of the image
	Alt       string   // the alternative text
	Size      IMG_SIZE // the size or the ratio of the image
	IsRounded bool     // Rounded image style
}

// Ensure Image implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*Image)(nil)

func NewImage(rawUrl string, size IMG_SIZE) *Image {
	img := new(Image)
	img.Size = size
	img.ParseSrcURL(rawUrl)
	return img
}

// TryParseSrc parses rawurl to _img.Src and returns _img to allow chaining.
// Parsing errors are ignored and if any Src may stay nil.
func (image *Image) ParseSrcURL(rawUrl string) (_err error) {
	image.Src, _err = url.Parse(rawUrl)
	return
}

// BuildTag builds the tag used to render the html element.
func (fig *Image) BuildTag(tag *html.Tag) {
	tag.SetTagName("figure").
		AddClasses("image").
		AddClassesIf(fig.Size != "", "is-"+string(fig.Size))
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (image *Image) RenderContent(out io.Writer) error {
	img := html.NewSnippet("img", `role="img" focusable="false"`)
	img.Tag().
		SetURL("src", image.Src).
		AddClassesIf(image.IsRounded, "is-rounded").
		SetAttributeIf(image.Alt != "", "alt", image.Alt)

	image.RenderChilds(out, img)
	return nil
}
