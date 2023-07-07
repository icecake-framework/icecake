package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	linkBulmaDelete string = `<a href="https://ick.io/documentation/elements/delete/">bulma Delete</a>`
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
	btnreset := ick.Button("reset", "btnreset", `class="mb-3"`).
		SetColor(ick.COLOR_PRIMARY).
		SetOutlined(true).
		SetDisabled(true)
	cmp.RenderChild(out, ux, btnreset)

	// apis
	html.Render(out, nil, ick.Title(3, "ICKDelete APIs"))
	html.WriteString(out, `<div class="block">`+
		`<p><code>Delete(targetid string) *ICKDelete</code> is the only one Delete factory.</p>`+
		`<p><code>TargetId string</code> The element id to remove from the DOM when the delete button is clicked.</p>`+
		`</div>`)

	html.Render(out, nil, ick.Title(3, "UI APIs"))
	html.WriteString(out, `<div class="block">`+
		`<p><code>Delete(targetid string) *ICKDelete</code> is the only one Delete factory.</p>`+
		`<p><code>clock.Clock</code> The TargetID will be automatically removed after the clock Timeout duration if not zero. The timer starts when the delete button is rendered (call to addlisteners).</p>`+
		`<p><code>OnDelete func(*ICKDelete)</code> if it is set, it's called when the deletion occurs and after the targetId has been removed.</p>`+
		`</div>`)

	// rendering
	html.Render(out, nil, ick.Title(3, "Rendering"))
	html.WriteString(out, `<div class="box spaceout mr-5">`)
	uA1 := &ick.ICKDelete{TargetId: "Idone"}
	uA2 := &ick.ICKDelete{TargetId: "Idtwo", SIZE: ick.SIZE_LARGE}
	html.Render(out, cmp, uA1, uA2)
	html.WriteString(out, `</div>`)

	return nil
}
