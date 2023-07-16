package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

const (
	linkBulmaMedia string = `<a href="https://bulma.io/documentation/layout/media">bulma Media</a>`
)

type SectionDocMedia struct {
	SectionDocIcecake
}

func (sec *SectionDocMedia) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Media", "media.go", "ICKMedia")

	ickcore.RenderString(out, `<div class="block">`+
		`<p>ICKMedia is an icecake snippet providing the HTML rendering for a `, linkBulmaMedia, ` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	fake := ick.Elem("div", `class="content`,
		ickcore.ToHTML(`<p>
			<strong>John Smith</strong> <small>@johnsmith</small> <small>31m</small>
			<br>
			Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare magna eros, eu pellentesque tortor vestibulum ut. Maecenas non massa sem. Etiam finibus odio quis feugiat facilisis.
			</p>`))

	// styling
	ickcore.RenderChild(out, sec, ick.Title(4, "Styling"))
	ickcore.RenderString(out, `<div class="box">`)
	ickcore.RenderChild(out, sec,
		ick.Media(
			nil,
			fake,
			nil),
		ick.Media(
			ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_64x64),
			fake,
			nil),
		ick.Media(
			ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_64x64),
			fake,
			ick.Delete("btndel", "")),
		ick.Media(
			ick.Image("/assets/icecake.jpg", "a cake", ick.IMG_64x64),
			ick.Elem("div", "",
				fake,
				ick.Media(
					ick.Image("/assets/icecake.svg", "a cake", ick.IMG_64x64).SetNoCrop(true),
					fake,
					nil)),
			ick.Icon("bi bi-starfill")))
	ickcore.RenderString(out, `</div>`)

	return nil
}
