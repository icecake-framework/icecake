package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

type SectionDocHero struct{ SectionDocIcecake }

func (sec *SectionDocHero) RenderContent(out io.Writer) error {
	sec.RenderHead(out, "Hero", "hero.go", "ICKHero")

	// usages
	ickcore.RenderString(out, `<div class="box spaceout">`)
	ickcore.RenderChild(out, sec,
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
			ContainerAttr: ickcore.ParseAttributes(`class="has-text-centered"`),
		})

	ickcore.RenderString(out, `</div>`)

	return nil
}
