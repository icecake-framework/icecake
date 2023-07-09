package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

type SectionDocHero struct{ SectionDocIcecake }

func (sec *SectionDocHero) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Hero", "hero.go", "ICKHero")

	// usages
	html.RenderString(out, `<div class="box spaceout">`)
	html.RenderChild(out, sec,
		&ick.ICKHero{
			Height:   ick.HH_SMALL,
			Title:    *ick.Title(4, "Value proposition"),
			Subtitle: *ick.SubTitle(6, "Killing features"),
		},
		&ick.ICKHero{
			Title:    *ick.Title(4, "Value proposition"),
			Subtitle: *ick.SubTitle(6, "Killing features"),
			CTA:      *ick.Button("Call To Action"),
		},
		&ick.ICKHero{
			Height:        ick.HH_SMALL,
			Title:         *ick.Title(4, "Value proposition"),
			Subtitle:      *ick.SubTitle(6, "Killing features"),
			CTA:           *ick.Button("Call To Action"),
			ContainerAttr: html.ParseAttributes(`class="has-text-centered"`),
		})

	html.RenderString(out, `</div>`)

	return nil
}
