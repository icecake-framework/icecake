package ick

import (
	"io"
	"net/url"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/lolorenzo777/verbose"
)

/******************************************************************************
* Component
******************************************************************************/

func init() {
	html.RegisterComposer("ick-image", &ICKImage{})
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

	IMG_RSQUARE IMG_SIZE = "is-fullwidth is-square"
	IMG_R1by1   IMG_SIZE = "is-fullwidth is-1by1"
	IMG_R5by4   IMG_SIZE = "is-fullwidth is-5by4"
	IMG_R4by3   IMG_SIZE = "is-fullwidth is-4by3"
	IMG_R3by2   IMG_SIZE = "is-fullwidth is-3by2"
	IMG_R5by3   IMG_SIZE = "is-fullwidth is-5by3"
	IMG_R16by9  IMG_SIZE = "is-fullwidth is-16by9"
	IMG_R2by1   IMG_SIZE = "is-fullwidth is-2by1"
	IMG_R3by1   IMG_SIZE = "is-fullwidth is-3by1"
	IMG_R4by5   IMG_SIZE = "is-fullwidth is-4by5"
	IMG_R3by4   IMG_SIZE = "is-fullwidth is-3by4"
	IMG_R2by3   IMG_SIZE = "is-fullwidth is-2by3"
	IMG_R3by5   IMG_SIZE = "is-fullwidth is-3by5"
	IMG_R9by16  IMG_SIZE = "is-fullwidth is-9by16"
	IMG_R1by2   IMG_SIZE = "is-fullwidth is-1by2"
	IMG_R1by3   IMG_SIZE = "is-fullwidth is-1by3"

	IMG_SIZE_OPTIONS string = string(IMG_16x16+" "+IMG_24x24+" "+IMG_32x32+" "+IMG_48x48+" "+IMG_64x64+" "+IMG_96x96+" "+IMG_128x128) + " " +
		string(IMG_RSQUARE+" "+IMG_R1by1+" "+IMG_R5by4+" "+IMG_R4by3+" "+IMG_R3by2+" "+IMG_R5by3+" "+IMG_R16by9+" "+IMG_R2by1+" "+IMG_R3by1+" "+IMG_R4by5+" "+IMG_R3by4+" "+IMG_R2by3+" "+IMG_R3by5+" "+IMG_R9by16+" "+IMG_R1by2+" "+IMG_R1by3)
)

// ICKImage is a typical img element embedded into a figure container specifying the image size. See [bulma image]
//
// Background color can be setup by adding snipet classes.
//
// [bulma image]: https://bulma.io/documentation/elements/image/
type ICKImage struct {
	html.BareSnippet

	Src       *url.URL // the url for the source of the image
	Alt       string   // the alternative text
	Size      IMG_SIZE // the size or the ratio of the image
	IsRounded bool     // Rounded image style
	NoCrop    bool     // set to true to avoid to crop the image if its size does not fit the Size property. The image may be reduced or distord.
}

// Ensuring ICKImage implements the right interface
var _ html.ElementComposer = (*ICKImage)(nil)

func Image(rawUrl string, alt string, size IMG_SIZE, attrs ...string) *ICKImage {
	img := new(ICKImage)
	img.Alt = alt
	img.Size = size
	img.ParseSrcURL(rawUrl)
	img.Tag().ParseAttributes(attrs...)
	return img
}

// ParseSrc parses rawurl to img.Src and returns img to allow chaining.
// Parsing errors are ignored and if any, Src may stay nil.
func (img *ICKImage) ParseSrcURL(rawUrl string) *ICKImage {
	var err error
	img.Src, err = url.Parse(rawUrl)
	verbose.Error("ICKImage.ParseSrcURL", err)
	return img
}

func (img *ICKImage) SetAlt(alt string) *ICKImage {
	img.Alt = alt
	return img
}
func (img *ICKImage) SetSize(s IMG_SIZE) *ICKImage {
	img.Size = s
	return img
}
func (img *ICKImage) SetRounded(f bool) *ICKImage {
	img.IsRounded = f
	return img
}
func (img *ICKImage) SetNoCrop(f bool) *ICKImage {
	img.NoCrop = f
	return img
}

/******************************************************************************/

// BuildTag builds the tag used to render the html element.
func (fig *ICKImage) BuildTag() html.Tag {
	fig.Tag().SetTagName("figure").
		AddClass("image").
		AddClassIf(!fig.NoCrop, "ickcropimage").
		PickClass(IMG_SIZE_OPTIONS, string(fig.Size))
	return *fig.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (image *ICKImage) RenderContent(out io.Writer) error {
	img := html.Snippet("img", `role="img" focusable="false"`)
	img.Tag().SetURL("src", image.Src).
		SetClassIf(image.IsRounded, "is-rounded").
		SetAttributeIf(image.Alt != "", "alt", image.Alt)

	html.RenderChild(out, image, img)
	return nil
}
