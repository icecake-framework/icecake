package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaImage string = `<a href="https://bulma.io/documentation/elements/image/">bulma Image</a>`
)

type SectionDocImage struct{ SectionDocIcecake }

func (sec *SectionDocImage) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Image", "image.go", "ICKImage")

	html.RenderString(out, `<div class="block">`+
		`<p>ICKImage is an icecake snippet providing the HTML rendering for a `, linkBulmaImage, ` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// usages
	html.RenderString(out, `<div class="box is-flex spaceout">`)
	html.RenderChild(out, sec,
		ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_128x128),
		ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_64x64),
		ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_32x32),
		ick.Image("/assets/snow-and-ice-cake.jpg", "a cake", ick.IMG_128x128),
		ick.Image("/assets/snow-and-ice-cake.jpg", "a cake", ick.IMG_128x128, `class="has-background-info-light`).SetNoCrop(true),
		ick.Image("/assets/broken.jpg", "a cake", ick.IMG_128x128, `class="has-background-info-light`))
	html.RenderString(out, `</div>`)

	return nil
}
