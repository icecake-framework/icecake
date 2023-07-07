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
	html.WriteString(out, `<div class="block">`+
		`<p>ICKImage is an icecake snippet providing the HTML rendering for a `, linkBulmaImage, ` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// usages
	html.WriteString(out, `<div class="box is-flex spaceout">`)
	u1 := ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_128x128)
	u2 := ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_64x64)
	u3 := ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_32x32)
	u4 := ick.Image("/assets/snow-and-ice-cake.jpg", "a cake", ick.IMG_128x128)
	u5 := ick.Image("/assets/snow-and-ice-cake.jpg", "a cake", ick.IMG_128x128, `class="has-background-info-light`).SetNoCrop(true)
	u6 := ick.Image("/assets/broken.jpg", "a cake", ick.IMG_128x128, `class="has-background-info-light`)
	html.Render(out, nil, u1, u2, u3, u4, u5, u6)
	html.WriteString(out, `</div>`)

	return nil
}
