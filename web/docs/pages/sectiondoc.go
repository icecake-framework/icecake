package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
)

type SectionDocIcecake struct {
	html.HTMLSnippet

	Title       string
	Description string
}

func SectionDoc(section string) html.ContentComposer {
	switch section {
	case "docoverview":
		s := new(SectionDocOverview)
		s.Title = "overview - icecake framework documentation"
		return s
	case "docbutton":
		s := new(SectionDocButton)
		s.Title = "button snippet - icecake framework documentation"
		return s
	case "doccard":
		s := new(SectionDocCard)
		s.Title = "card snippet - icecake framework documentation"
		return s
	case "docdelete":
		s := new(SectionDocDelete)
		s.Title = "delete snippet - icecake framework documentation"
		return s
	case "dochero":
		s := new(SectionDocHero)
		s.Title = "hero snippet - icecake framework documentation"
		return s
	case "docimage":
		s := new(SectionDocImage)
		s.Title = "image Snippet - icecake framework documentation"
		return s
	case "docmenu":
		s := new(SectionDocMenu)
		s.Title = "Menu Snippet - icecake framework documentation"
		return s
	case "docmessage":
		s := new(SectionDocMessage)
		s.Title = "Message Snippet - icecake framework documentation"
		return s
	case "docnavbar":
		s := new(SectionDocNavbar)
		s.Title = "Navbar Snippet - icecake framework documentation"
		return s
	case "docnotify":
		s := new(SectionDocNotify)
		s.Title = "Notify Snippet - icecake framework documentation"
		return s
	}
	s := new(SectionDocIcecake)
	s.Title = "icecake framework documentation"
	return s
}

func (cmp *SectionDocIcecake) BuildTag() html.Tag {
	cmp.Tag().SetTagName("section").AddClass("py-5")
	return *cmp.Tag()
}

func (cmp *SectionDocIcecake) RenderContent(out io.Writer) error {
	html.WriteString(out, `<p>default section</p>`)

	return nil
}
