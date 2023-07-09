package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

type SectionDocMenu struct{ SectionDocIcecake }

func (sec *SectionDocMenu) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Menu", "menu.go", "ICKMenu")

	// usages
	html.RenderString(out, `<div class="box spaceout">`)
	u1 := ick.Menu("u1")
	u1.AddItem("", ick.MENUIT_LABEL, "Label")
	u1.AddItem("", ick.MENUIT_LINK, "link 1")
	u1.AddItem("", ick.MENUIT_LINK, "link 2")
	u1.AddItem("", ick.MENUIT_NESTEDLINK, "Nested Link 1")
	u1.AddItem("", ick.MENUIT_NESTEDLINK, "Nested Link 2")
	u1.AddItem("", ick.MENUIT_LINK, "link 3")
	html.RenderChild(out, sec, u1)
	html.RenderString(out, `</div>`)

	return nil
}
