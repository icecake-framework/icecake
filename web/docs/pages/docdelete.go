package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

const (
	linkBulmaDelete string = `<a href="https://bulma.io/documentation/elements/delete/">bulma Delete</a>`
)

type SectionDocDelete struct{ SectionDocIcecake }

func (sec *SectionDocDelete) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Delete", "delete.go", "ICKDelete")

	ickcore.RenderString(out, `<div class="block">`+
		`<p>ICKDelete is an icecake snippet providing the HTML rendering for a `, linkBulmaDelete, `</p>`+
		`<p>The html rendering is a simple button with a centered cross.</p>`+
		`</div>`)

	// usages
	ickcore.RenderChild(out, sec,
		ick.Elem("div", `id="boxusage" class="box mr-5"`, ick.Spinner()),
		ick.Button("reset", `class="mb-3"`).
			SetId("btnreset").
			SetColor(ick.COLOR_PRIMARY).
			SetOutlined(true).
			SetDisabled(true))

	// rendering
	ickcore.RenderChild(out, sec, ick.Title(3, "Rendering"))
	ickcore.RenderString(out, `<div class="box spaceout">`)
	ickcore.RenderChild(out, sec,
		&ick.ICKDelete{TargetId: "Idone"},
		&ick.ICKDelete{TargetId: "Idtwo", SIZE: ick.SIZE_LARGE})
	ickcore.RenderString(out, `</div>`)

	return nil
}
