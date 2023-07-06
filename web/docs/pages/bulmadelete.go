package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaDelete string = `<a href="https://ick.io/documentation/elements/delete/">bulma Delete</a>`
)

type SectionIckDelete struct{ SectionIcecakeDoc }

func (cmp *SectionIckDelete) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Delete</h2>`)
	html.WriteString(out, `<p>ick.Delete is an icecake snippet providing the HTML rendering for a `, linkBulmaDelete, `</p>`)
	html.WriteString(out, `<p>The html rendering is a simple button with a centered cross.</p>`)

	// API
	html.WriteString(out, `<h3>bulmaui.Delete API</h3>`)

	// usage
	html.WriteString(out, `<h3>Usage</h3>`)
	cmp.RenderChild(out, ick.Button(*html.ToHTML("reset"), "btnreset", "", `class="mb-3"`).SetColor(ick.COLOR_PRIMARY).SetOutlined(true))
	html.WriteString(out, `<div id="boxusage" class="box mr-5">`)
	html.Render(out, nil, ick.Spinner())
	html.WriteString(out, `</div>`)

	// rendering
	html.WriteString(out, `<h3>Rendering</h3>`)
	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uA1 := &ick.ICKDelete{TargetId: "Idone"}
	uA2 := &ick.ICKDelete{TargetId: "Idtwo", SIZE: ick.SIZE_LARGE}
	html.Render(out, cmp, uA1, uA2)
	html.WriteString(out, `</div>`)

	return nil
}
