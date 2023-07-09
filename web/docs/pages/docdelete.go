package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaDelete string = `<a href="https://bulma.io/documentation/elements/delete/">bulma Delete</a>`
)

type SectionDocDelete struct{ SectionDocIcecake }

func (sec *SectionDocDelete) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Delete", "delete.go", "ICKDelete")

	html.RenderString(out, `<div class="block">`+
		`<p>ICKDelete is an icecake snippet providing the HTML rendering for a `, linkBulmaDelete, `</p>`+
		`<p>The html rendering is a simple button with a centered cross.</p>`+
		`</div>`)

	// usages
	html.RenderChild(out, sec,
		html.Snippet("div", `id="boxusage" class="box mr-5"`, ick.Spinner()),
		ick.Button("reset", `class="mb-3"`).
			SetId("btnreset").
			SetColor(ick.COLOR_PRIMARY).
			SetOutlined(true).
			SetDisabled(true))

	// rendering
	html.RenderChild(out, sec, ick.Title(3, "Rendering"))
	html.RenderString(out, `<div class="box spaceout">`)
	html.RenderChild(out, sec,
		&ick.ICKDelete{TargetId: "Idone"},
		&ick.ICKDelete{TargetId: "Idtwo", SIZE: ick.SIZE_LARGE})
	html.RenderString(out, `</div>`)

	return nil
}
