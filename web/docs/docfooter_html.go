package docs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

func init() {
	ickcore.RequireCSSStyle("docsFooter", docsFooterStyle)
}

/******************************************************************************
/* docsFooter
******************************************************************************/

const docsFooterStyle string = `.myfooter-title {
	color: rgb(122, 122, 122);
}
.myfooter-info {
	color: rgb(122, 122, 122);
	font-size: 0.9rem;
}
.myfooter-link {
	color: rgb(122, 122, 122);
}
`

type docFooter struct {
	ickcore.BareSnippet
	page *ick.Page
}

// Ensuring docFooter implements the right interface
var _ ickcore.ContentComposer = (*docFooter)(nil)
var _ ickcore.TagBuilder = (*docFooter)(nil)

func DocFooter(pg *ick.Page) *docFooter {
	return &docFooter{page: pg}
}

func (footer docFooter) NeedRendering() bool {
	return true
}

func (footer *docFooter) BuildTag() ickcore.Tag {
	footer.Tag().SetTagName("footer").AddClass("footer")
	return *footer.Tag()
}

func (footer *docFooter) RenderContent(out io.Writer) error {

	hrefMIT := `<a href="https://opensource.org/licenses/mit-license.php" rel="license">MIT</a>`
	hrefCCBY := `<a href="https://creativecommons.org/licenses/by-nc-sa/4.0/">CC BY-NC-SA 4.0</a>`
	hrefLinks := []string{
		`<a href="` + footer.page.ToAbsURLString("/") + `">Home</a>`,
		`<a href="` + footer.page.ToAbsURLString("/docoverview.html") + `">Docs</a>`,
		`<a href="https://github.com/icecake-framework/icecake">Contribute</a> on GitHub`,
	}

	ickcore.RenderString(out, `<div class="container"><div class="columns">`)

	// 1st column
	ickcore.RenderString(out, `<div class="column is-8">`)
	ickcore.RenderString(out, `<h4 class="myfooter-title">`, `<strong>IceCake</strong> by Lolorenzo`, `</h4>`)
	ickcore.RenderString(out, `<div class="myfooter-info">`, `Source code licences `, hrefMIT, `</div>`)
	ickcore.RenderString(out, `<div class="myfooter-info">`, `Website content licensed `, hrefCCBY, `</div>`)
	ickcore.RenderString(out, `<br><div class="myfooter-info">Wasm code: <span id="icecake-status"></span></div>`)
	ickcore.RenderString(out, `</div>`)

	// 2nd column
	ickcore.RenderString(out, `<div class="column is-4">`)
	ickcore.RenderString(out, `<h4 class="myfooter-title">`, `<strong>Links</strong>`, `</h4>`)
	for _, hrefLink := range hrefLinks {
		ickcore.RenderString(out, `<p class="myfooter-link">`, hrefLink, `</p>`)
	}
	ickcore.RenderString(out, `</div>`)

	ickcore.RenderString(out, `</div></div>`)
	return nil
}
