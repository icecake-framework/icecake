package ick

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/ickcore"
)

func init() {
	ickcore.RegisterComposer("ick-hero", &ICKHero{})
}

// The height of the hero section
type HERO_HEIGHT string

const (
	HH_SMALL                   HERO_HEIGHT = "is-small"
	HH_STANDARD                HERO_HEIGHT = ""
	HH_MEDIUM                  HERO_HEIGHT = "is-medium"
	HH_LARGE                   HERO_HEIGHT = "is-large"
	HH_HALFHEIGHT              HERO_HEIGHT = "is-halfheight"
	HH_FULLFHEIGHT             HERO_HEIGHT = "is-fullheight"
	HH_FULLFHEIGHT_WITH_NAVBAR HERO_HEIGHT = "is-fullheight-with-navbar"
	HH_OPTIONS                 string      = string(HH_SMALL + " " + HH_MEDIUM + " " + HH_LARGE + " " + HH_HALFHEIGHT + " " + HH_FULLFHEIGHT + " " + HH_FULLFHEIGHT_WITH_NAVBAR)
)

type ICKHero struct {
	ickcore.BareSnippet

	Height HERO_HEIGHT // the height of the hero section,

	InsideHead ickcore.ContentComposer

	Title    ICKTitle
	Subtitle ICKTitle
	Centered bool
	CWidth   CONTAINER_WIDTH

	CTA ICKButton

	InsideFoot ickcore.ContentComposer
}

// Ensuring Hero implements the right interface
var _ ickcore.ContentComposer = (*ICKHero)(nil)
var _ ickcore.TagBuilder = (*ICKHero)(nil)

func Hero() *ICKHero {
	hero := new(ICKHero)
	return hero
}

// Tag Builder used by the rendering functions.
func (h *ICKHero) BuildTag() ickcore.Tag {
	h.Tag().SetTagName("section").AddClass("hero").PickClass(HH_OPTIONS, string(h.Height))
	return *h.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (h *ICKHero) RenderContent(out io.Writer) error {

	if h.InsideHead != nil {
		ickcore.RenderChild(out, h, Elem("div", `class="hero-head"`, h.InsideHead))
	}

	ickcore.RenderString(out, `<div class="hero-body">`)

	c := Container(h.CWidth)
	c.Tag().AddClassIf(h.Centered, "has-text-centered")
	c.Append(&h.Title, &h.Subtitle, &h.CTA)
	ickcore.RenderChild(out, h, c)

	ickcore.RenderString(out, `</div>`)

	if h.InsideFoot != nil {
		ickcore.RenderChild(out, h, Elem("div", `class="hero-foot"`, h.InsideFoot))
	}

	return nil
}
