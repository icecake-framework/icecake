package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

const (
	href_GitPkg string = `https://github.com/icecake-framework/icecake/blob/main/pkg`
	href_GoPkg  string = `https://pkg.go.dev/github.com/icecake-framework/icecake/pkg`
)

type SectionDocIcecake struct {
	html.BareSnippet

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
	case "docicon":
		s := new(SectionDocIcon)
		s.Title = "Icon Snippet - icecake framework documentation"
		return s
	}
	s := new(SectionDocIcecake)
	s.Title = "icecake framework documentation"
	return s
}

func (cmp *SectionDocIcecake) BuildTag() html.Tag {
	cmp.Tag().SetTagName("section").AddClass("py-5 px-5")
	return *cmp.Tag()
}

func (cmp *SectionDocIcecake) RenderContent(out io.Writer) error {
	html.WriteString(out, `<p>default section</p>`)

	return nil
}

func (doc *SectionDocIcecake) RenderHead(out io.Writer, title string, gitpkg string, gostruct string) error {
	// html.Render(out, nil, ick.Title(3, "APIs"))
	b := ick.Button("").
		SetSize(ick.SIZE_SMALL).
		SetColor(ick.COLOR_LINK).
		SetOutlined(true).
		SetIcon(*ick.Icon("bi bi-box-arrow-up-right", `class="is-hidden-touch"`), true)

	hrefICK_Git := href_GitPkg + `/ick/` + gitpkg
	hrefICK_GitUI := href_GitPkg + `/ick/ickui/` + gitpkg
	hrefICK_Go := href_GoPkg + `/ick#` + gostruct
	hrefICK_GoUI := href_GoPkg + `/ick/ickui#` + gostruct

	html.WriteString(out, `<div class="is-flex is-justify-content-space-between">`)
	html.Render(out, nil, ick.Title(3, title, `style="white-space: nowrap;"`))

	html.Render(out, nil, html.Snippet("div", "class='is-flex is-justify-content-flex-end spaceout'").SetBody(
		b.Clone().SetTitle(gostruct+" code").ParseHRef(hrefICK_Git).SetIcon(*ick.Icon("bi bi-github"), false),
		b.Clone().SetTitle("UI code").ParseHRef(hrefICK_GitUI).SetIcon(*ick.Icon("bi bi-github"), false),
		b.Clone().SetTitle(gostruct+" Go pkg").ParseHRef(hrefICK_Go).SetIcon(*ick.Icon("bi bi-book"), false),
		b.Clone().SetTitle("UI Go pkg").ParseHRef(hrefICK_GoUI).SetIcon(*ick.Icon("bi bi-book"), false),
	))
	html.WriteString(out, `</div>`)
	return nil
}
