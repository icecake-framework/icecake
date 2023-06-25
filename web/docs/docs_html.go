package main

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RequireCSSStyle("docsFooter", docsFooterStyle)
}

/******************************************************************************
/* docOverview
******************************************************************************/

type docOverview struct{ html.HTMLSnippet }

func (doc *docOverview) BuildTag(tag *html.Tag) { tag.SetTagName("div").AddClasses("block") }

func (doc *docOverview) RenderContent(out io.Writer) error {
	html.WriteString(out, `
	<h2 class="title is-2">Overview</h2>
	<p>welcome</p>
	`)
	return nil
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

type docsFooter struct{ html.HTMLSnippet }

func (footer *docsFooter) BuildTag(tag *html.Tag) { tag.SetTagName("footer").AddClasses("footer") }

func (footer *docsFooter) RenderContent(out io.Writer) error {

	hrefMIT := `<a href="https://opensource.org/licenses/mit-license.php" rel="license">MIT</a>`
	hrefCCBY := `<a href="https://creativecommons.org/licenses/by-nc-sa/4.0/">CC BY-NC-SA 4.0</a>`
	hrefLinks := []string{
		`<a href="/">Home</a>`,
		`<a href="/docs">Docs</a>`,
		`<a href="https://github.com/icecake-framework/icecake">Contribute</a> on GitHub`,
	}

	html.WriteString(out, `<div class="container"><div class="columns">`)

	// 1st column
	html.WriteString(out, `<div class="column is-8">`)
	html.WriteStrings(out, `<h4 class="myfooter-title">`, `<strong>IceCake</strong> by Lolorenzo`, `</h4>`)
	html.WriteStrings(out, `<div class="myfooter-info">`, `Source code licences `, hrefMIT, `</div>`)
	html.WriteStrings(out, `<div class="myfooter-info">`, `Website content licensed `, hrefCCBY, `</div>`)
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
