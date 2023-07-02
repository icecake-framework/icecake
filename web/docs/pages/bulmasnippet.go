package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionIcecakeDoc struct {
	html.HTMLSnippet

	Title       string
	Description string
}

func SectionDoc(section string) html.HTMLComposer {
	switch section {
	case "overview":
		s := new(SectionOverview)
		s.Title = "overview - icecake framework documentation"
		return s
	case "bulmabutton":
		s := new(SectionBulmaButton)
		s.Title = "bulma button - icecake framework documentation"
		return s
	case "bulmacard":
		s := new(SectionBulmaCard)
		s.Title = "bulma card - icecake framework documentation"
		return s
	case "bulmadelete":
		s := new(SectionBulmaDelete)
		s.Title = "bulma delete - icecake framework documentation"
		return s
	case "bulmahero":
		s := new(SectionBulmaHero)
		s.Title = "bulma hero - icecake framework documentation"
		return s
	case "bulmaimage":
		s := new(SectionBulmaImage)
		s.Title = "bulma image - icecake framework documentation"
		return s
	case "bulmamenu":
		s := new(SectionBulmaMenu)
		s.Title = "bulma menu - icecake framework documentation"
		return s
	case "bulmamessage":
		s := new(SectionBulmaMessage)
		s.Title = "bulma message - icecake framework documentation"
		return s
	case "bulmanavbar":
		s := new(SectionBulmaNavbar)
		s.Title = "bulma navbar - icecake framework documentation"
		return s
	case "bulmanotify":
		s := new(SectionBulmaNotify)
		s.Title = "bulma notify - icecake framework documentation"
		return s
	}
	s := new(SectionIcecakeDoc)
	s.Title = "icecake framework documentation"
	return s
}

func (cmp *SectionIcecakeDoc) BuildTag(tag *html.Tag) {
	tag.SetTagName("section").AddClass("content py-3")
}

func (cmp *SectionIcecakeDoc) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>???</h2>
	<p>welcome</p>`)

	return nil
}
