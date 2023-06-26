package main

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

/******************************************************************************
/* docOverview
******************************************************************************/

type docOverview struct{ html.HTMLSnippet }

func (doc *docOverview) BuildTag(tag *html.Tag) { tag.SetTagName("div").AddClasses("content py-3") }

func (doc *docOverview) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Overview</h2>
	<p>welcome</p>`)

	return nil
}
