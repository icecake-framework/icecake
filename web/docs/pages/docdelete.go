package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaDelete     string = `<a href="https://bulma.io/documentation/elements/delete/">bulma Delete</a>`
	hrefICKDelete_Git   string = href_GitPkg + `/ick/delete.go`
	hrefICKDelete_GitUI string = href_GitPkg + `/ick/ickui/delete.go`
	hrefICKDelete_Go    string = href_GoPkg + `/ick#ICKDelete`
	hrefICKDelete_GoUI  string = href_GoPkg + `/ick/ickui#ICKDelete`
)

type SectionDocDelete struct{ SectionDocIcecake }

func (cmp *SectionDocDelete) RenderContent(out io.Writer) error {
	html.Render(out, nil, ick.Title(3, "Delete"))
	html.WriteString(out, `<div class="block">`+
		`<p>ICKDelete is an icecake snippet providing the HTML rendering for a `, linkBulmaDelete, `</p>`+
		`<p>The html rendering is a simple button with a centered cross.</p>`+
		`</div>`)

	// usages
	ux := html.Snippet("div", `id="boxusage" class="box mr-5"`).AddContent(ick.Spinner())
	btnreset := ick.Button("reset", `class="mb-3"`).
		SetId("btnreset").
		SetColor(ick.COLOR_PRIMARY).
		SetOutlined(true).
		SetDisabled(true)
	cmp.RenderChild(out, ux, btnreset)

	// apis
	html.Render(out, nil, ick.Title(3, "APIs"))
	b := ick.Button("").SetSize(ick.SIZE_SMALL).SetColor(ick.COLOR_LINK).SetOutlined(true)
	html.Render(out, nil, html.Snippet("div", "class='is-flex spaceout'").AddContent(
		b.Clone().SetTitle("ICKDelete code").ParseHRef(hrefICKDelete_Git),
		b.Clone().SetTitle("ICKDelete UI code").ParseHRef(hrefICKDelete_GitUI),
		b.Clone().SetTitle("ICKDelete Go pkg doc").ParseHRef(hrefICKDelete_Go),
		b.Clone().SetTitle("ICKDelete UI Go pkg doc").ParseHRef(hrefICKDelete_GoUI),
	))

	// rendering
	html.Render(out, nil, ick.Title(3, "Rendering"))
	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uA1 := &ick.ICKDelete{TargetId: "Idone"}
	uA2 := &ick.ICKDelete{TargetId: "Idtwo", SIZE: ick.SIZE_LARGE}
	html.Render(out, cmp, uA1, uA2)
	html.WriteString(out, `</div>`)

	return nil
}
