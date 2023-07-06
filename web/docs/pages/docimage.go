package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaImage string = `<a href="https://ick.io/documentation/elements/image/">bulma Image</a>`
)

type SectionDocImage struct{ SectionDocIcecake }

func (cmp *SectionDocImage) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Image</h2>
	<p>ICKImage is an icecake snippet providing the HTML rendering for a `, linkBulmaImage, ` with extra features and usefull Go APIs.</p>`)

	// usages
	html.WriteString(out, `<div class="box is-flex mr-5">`)

	u1 := ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_128x128, `class="m-0 mr-2"`)
	u2 := ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_64x64, `class="m-0 mr-2"`)
	u3 := ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_32x32, `class="m-0 mr-2"`)
	u4 := ick.Image("/assets/snow-and-ice-cake.jpg", "a cake", ick.IMG_128x128, `class="m-0 mr-2"`)
	u5 := ick.Image("/assets/snow-and-ice-cake.jpg", "a cake", ick.IMG_128x128, `class="m-0 mr-2 has-background-info-light`).SetNoCrop(true)
	u6 := ick.Image("/assets/broken.jpg", "a cake", ick.IMG_128x128, `class="m-0 mr-2 has-background-info-light`)

	html.Render(out, nil, u1, u2, u3, u4, u5, u6)
	html.WriteString(out, `</div>`)

	// apis
	html.WriteString(out, `<h3>ick.ICKImage API</h3>`+
		`<p><code>Image(rawUrl string, size IMG_SIZE) *ICKImage</code> is the main Image factory.</p>`+
		`<p><code>.ParseSrcURL(rawUrl string)</code> parses rawurl to img.Src.</p>`+
		`<p><code>.SetAlt(alt string) *ICKImage</code> sets alternate text.</p>`+
		`<p><code>.SetSize(s IMG_SIZE) *ICKImage</code> sets image size.</p>`+
		`<p><code>.SetRounded(f bool) *ICKImage</code> sets rounded style.</p>`)

	return nil
}
