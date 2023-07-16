package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

const (
	linkBulmaTaglabel string = `<a href="https://bulma.io/documentation/elements/tag">bulma Tag</a>`
)

type SectionDocTagLabel struct {
	SectionDocIcecake
}

func (sec *SectionDocTagLabel) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "TagLabel", "taglabel.go", "ICKTagLabel")

	ickcore.RenderString(out, `<div class="block">`+
		`<p>ICKTagLabel is an icecake snippet providing the HTML rendering for a `, linkBulmaTaglabel, ` with extra features and usefull Go APIs.</p>`+
		`</div>`)

	// styling
	ickcore.RenderChild(out, sec, ick.Title(4, "Styling"))
	ickcore.RenderString(out, `<div class="box spaceout">`)
	ickcore.RenderChild(out, sec,
		ick.TagLabel("label", ick.COLOR_LIGHT),
		ick.TagLabel("medium", ick.COLOR_LIGHT).SetSize(ick.TAGLBLSZ_MEDIUM),
		ick.TagLabel("large", ick.COLOR_LIGHT).SetSize(ick.TAGLBLSZ_LARGE),
		ick.TagLabel("primary", ick.COLOR_PRIMARY),
		ick.TagLabel("rounder", ick.COLOR_NONE).SetRounded(true),
		ick.TagLabel("deletable", ick.COLOR_NONE).SetCanDelete(true),
		ick.TagLabel("deletable", *ick.Color(ick.COLOR_SUCCESS).SetLight(true)).
			SetCanDelete(true).
			SetSize(ick.TAGLBLSZ_MEDIUM),
		ick.TagLabel("alpha 4", *ick.Color(ick.COLOR_PRIMARY).SetLight(true)).
			SetHeader("Version", ick.COLOR_BLACK),
		ick.TagLabel("alpha 4", *ick.Color(ick.COLOR_PRIMARY).SetLight(true)).
			SetCanDelete(true).
			SetHeader("Version", ick.COLOR_BLACK),
		ick.TagLabel("", ick.COLOR("")).
			SetHeader("only header", ick.COLOR_SUCCESS),
		ick.TagLabel("", ick.COLOR("")).
			SetCanDelete(true).
			SetHeader("deletable label", ick.COLOR_WARNING))
	ickcore.RenderString(out, `</div>`)

	return nil
}
