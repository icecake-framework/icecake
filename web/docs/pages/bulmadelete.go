package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/html"
)

const (
	linkBulmaDelete string = `<a href="https://bulma.io/documentation/elements/delete/">bulma Delete</a>`
)

type SectionBulmaDelete struct{ SectionIcecakeDoc }

func (cmp *SectionBulmaDelete) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Bulma Delete</h2>`)
	html.WriteString(out, `<p>bulma.Delete is an icecake snippet providing the HTML rendering for a `, linkBulmaDelete, `</p>`)
	html.WriteString(out, `<p>The html rendering is a simple button with a centered cross.</p>`)

	// API
	html.WriteString(out, `<h3>bulmaui.Delete API</h3>`)

	// usage
	html.WriteString(out, `<h3>Usage</h3>`)
	cmp.RenderChilds(out, bulma.Button(*html.ToHTML("reset"), "btnreset", "", `class="mb-3"`).SetColor(bulma.COLOR_PRIMARY).SetOutlined(true))
	html.WriteString(out, `<div id="boxusage" class="box mr-5">`)
	html.Render(out, nil, bulma.Spinner())
	html.WriteString(out, `</div>`)

	// rendering
	html.WriteString(out, `<h3>Rendering</h3>`)
	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uA1 := &bulma.Delete{TargetID: "Idone"}
	uA2 := &bulma.Delete{TargetID: "Idtwo", SIZE: bulma.SIZE_LARGE}
	html.Render(out, cmp, uA1, uA2)
	html.WriteString(out, `</div>`)

	return nil
}
