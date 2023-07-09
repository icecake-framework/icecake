package docs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RequireCSSStyle("docsFooter", docsFooterStyle)
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

type docFooter struct{ html.BareSnippet }

// Ensuring docFooter implements the right interface
var _ html.ElementComposer = (*docFooter)(nil)

func DocFooter() *docFooter {
	return new(docFooter)
}

func (footer *docFooter) BuildTag() html.Tag {
	footer.Tag().SetTagName("footer").AddClass("footer")
	return *footer.Tag()
}

func (footer *docFooter) RenderContent(out io.Writer) error {

	hrefMIT := `<a href="https://opensource.org/licenses/mit-license.php" rel="license">MIT</a>`
	hrefCCBY := `<a href="https://creativecommons.org/licenses/by-nc-sa/4.0/">CC BY-NC-SA 4.0</a>`
	hrefLinks := []string{
		`<a href="/">Home</a>`,
		`<a href="/overview">Docs</a>`,
		`<a href="https://github.com/icecake-framework/icecake">Contribute</a> on GitHub`,
	}

	html.RenderString(out, `<div class="container"><div class="columns">`)

	// 1st column
	html.RenderString(out, `<div class="column is-8">`)
	html.RenderString(out, `<h4 class="myfooter-title">`, `<strong>IceCake</strong> by Lolorenzo`, `</h4>`)
	html.RenderString(out, `<div class="myfooter-info">`, `Source code licences `, hrefMIT, `</div>`)
	html.RenderString(out, `<div class="myfooter-info">`, `Website content licensed `, hrefCCBY, `</div>`)
	html.RenderString(out, `<br><div class="myfooter-info">Wasm code: <span id="icecake-status"></span></div>`)
	html.RenderString(out, `</div>`)

	// 2nd column
	html.RenderString(out, `<div class="column is-4">`)
	html.RenderString(out, `<h4 class="myfooter-title">`, `<strong>Links</strong>`, `</h4>`)
	for _, hrefLink := range hrefLinks {
		html.RenderString(out, `<p class="myfooter-link">`, hrefLink, `</p>`)
	}
	html.RenderString(out, `</div>`)

	html.RenderString(out, `</div></div>`)
	return nil
}
