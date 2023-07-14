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

	ContainerAttr ickcore.AttributeMap // The attributes map to setup to the hero's container, allowing text centering

	Title    ICKTitle
	Subtitle ICKTitle

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
func (msg *ICKHero) BuildTag() ickcore.Tag {
	msg.Tag().SetTagName("section").AddClass("hero").PickClass(HH_OPTIONS, string(msg.Height))
	return *msg.Tag()
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (msg *ICKHero) RenderContent(out io.Writer) error {

	if msg.InsideHead != nil {
		ickcore.RenderChild(out, msg, Elem("div", `class="hero-head"`, msg.InsideHead))
	}

	ickcore.RenderString(out, `<div class="hero-body">`)

	cont := new(Container)
	cont.Tag().AttributeMap = msg.ContainerAttr.Clone()
	contag := cont.BuildTag()
	contag.RenderOpening(out)

	ickcore.RenderChild(out, msg, &msg.Title, &msg.Subtitle, &msg.CTA)

	contag.RenderClosing(out)

	ickcore.RenderString(out, `</div>`)

	if msg.InsideFoot != nil {
		ickcore.RenderChild(out, msg, Elem("div", `class="hero-foor"`, msg.InsideFoot))
	}

	return nil
}
