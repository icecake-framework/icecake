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

type Footer struct{ html.HTMLSnippet }

func DocFooter() *Footer {
	return new(Footer)
}

func (footer *Footer) BuildTag(tag *html.Tag) { tag.SetTagName("footer").AddClasses("footer") }

func (footer *Footer) RenderContent(out io.Writer) error {

	hrefMIT := `<a href="https://opensource.org/licenses/mit-license.php" rel="license">MIT</a>`
	hrefCCBY := `<a href="https://creativecommons.org/licenses/by-nc-sa/4.0/">CC BY-NC-SA 4.0</a>`
	hrefLinks := []string{
		`<a href="/">Home</a>`,
		`<a href="/overview">Docs</a>`,
		`<a href="https://github.com/icecake-framework/icecake">Contribute</a> on GitHub`,
	}

	html.WriteString(out, `<div class="container"><div class="columns">`)

	// 1st column
	html.WriteString(out, `<div class="column is-8">`)
	html.WriteStrings(out, `<h4 class="myfooter-title">`, `<strong>IceCake</strong> by Lolorenzo`, `</h4>`)
	html.WriteStrings(out, `<div class="myfooter-info">`, `Source code licences `, hrefMIT, `</div>`)
	html.WriteStrings(out, `<div class="myfooter-info">`, `Website content licensed `, hrefCCBY, `</div>`)
	html.WriteStrings(out, `<br><div class="myfooter-info">Wasm code: <span id="icecake-status"></span></div>`)
	html.WriteString(out, `</div>`)

	// 2nd column
	html.WriteString(out, `<div class="column is-4">`)
	html.WriteStrings(out, `<h4 class="myfooter-title">`, `<strong>Links</strong>`, `</h4>`)
	for _, hrefLink := range hrefLinks {
		html.WriteStrings(out, `<p class="myfooter-link">`, hrefLink, `</p>`)
	}
	html.WriteString(out, `</div>`)

	html.WriteString(out, `</div></div>`)
	return nil
}
