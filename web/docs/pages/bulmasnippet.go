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

func SectionDoc(section string) html.HTMLContentComposer {
	switch section {
	case "overview":
		s := new(SectionOverview)
		s.Title = "overview - icecake framework documentation"
		return s
	case "bulmabutton":
		s := new(SectionIckButton)
		s.Title = "bulma button - icecake framework documentation"
		return s
	case "bulmacard":
		s := new(SectionBulmaCard)
		s.Title = "bulma card - icecake framework documentation"
		return s
	case "bulmadelete":
		s := new(SectionIckDelete)
		s.Title = "bulma delete - icecake framework documentation"
		return s
	case "bulmahero":
		s := new(SectionBulmaHero)
		s.Title = "bulma hero - icecake framework documentation"
		return s
	case "bulmaimage":
		s := new(SectionIckImage)
		s.Title = "bulma image - icecake framework documentation"
		return s
	case "bulmamenu":
		s := new(SectionBulmaMenu)
		s.Title = "bulma menu - icecake framework documentation"
		return s
	case "bulmamessage":
		s := new(SectionIckMessage)
		s.Title = "bulma message - icecake framework documentation"
		return s
	case "bulmanavbar":
		s := new(SectionIckNavbar)
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

func (cmp *SectionIcecakeDoc) BuildTag() html.Tag {
	cmp.Tag().SetTagName("section").AddClass("content py-5")
	return *cmp.Tag()
}

func (cmp *SectionIcecakeDoc) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>???</h2>
	<p>welcome</p>`)

	return nil
}
