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
	html.WriteString(out, `<div class="block">`+
		`<p>ICKDelete is an icecake snippet providing the HTML rendering for a `, linkBulmaDelete, `</p>`+
		`<p>The html rendering is a simple button with a centered cross.</p>`+
		`</div>`)

	// usages
	ux := html.Snippet("div", `id="boxusage" class="box mr-5"`).SetBody(ick.Spinner())
	btnreset := ick.Button("reset", `class="mb-3"`).
		SetId("btnreset").
		SetColor(ick.COLOR_PRIMARY).
		SetOutlined(true).
		SetDisabled(true)
	sec.RenderChild(out, ux, btnreset)

	// rendering
	html.Render(out, nil, ick.Title(3, "Rendering"))
	html.WriteString(out, `<div class="box spaceout">`)
	uA1 := &ick.ICKDelete{TargetId: "Idone"}
	uA2 := &ick.ICKDelete{TargetId: "Idtwo", SIZE: ick.SIZE_LARGE}
	html.Render(out, sec, uA1, uA2)
	html.WriteString(out, `</div>`)

	return nil
}
