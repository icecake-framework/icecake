package bulma

import (
	"io"
	"strconv"

	"github.com/icecake-framework/icecake/pkg/html"
)

func init() {
	html.RegisterComposer("ick-hero", &Hero{})
}

// The height of the hero section
type HERO_HEIGHT string

const (
	HH_SMALL                   HERO_HEIGHT = "small"
	HH_STANDARD                HERO_HEIGHT = ""
	HH_MEDIUM                  HERO_HEIGHT = "medium"
	HH_LARGE                   HERO_HEIGHT = "large"
	HH_HALFHEIGHT              HERO_HEIGHT = "halfheight"
	HH_FULLFHEIGHT             HERO_HEIGHT = "fullheight"
	HH_FULLFHEIGHT_WITH_NAVBAR HERO_HEIGHT = "fullheight-with-navbar"
)

type Hero struct {
	html.HTMLSnippet

	Height HERO_HEIGHT // the height of the hero section,

	InsideHead html.HTMLComposer

	*Container   // optional Container to render the content, allowing center
	Title        html.HTMLString
	TitleSize    int // 1 to 6
	Subtitle     html.HTMLString
	SubtitleSize int // 1 to 6

	CTA html.HTMLComposer // Call To Action

	InsideFoot html.HTMLComposer
}

// Ensure Hero implements HTMLComposer interface
var _ html.HTMLComposer = (*Hero)(nil)

// Tag Builder used by the rendering functions.
func (msg *Hero) BuildTag(tag *html.Tag) {
	tag.SetTagName("section").AddClasses("hero").AddClassesIf(msg.Height != "", "is-"+string(msg.Height))
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (msg *Hero) RenderContent(out io.Writer) error {

	if msg.InsideHead != nil {
		html.WriteString(out, `<div class="hero-head">`)
		msg.RenderChildSnippet(out, msg.InsideHead)
		html.WriteString(out, `</div>`)
	}

	html.WriteString(out, `<div class="hero-body">`)

	if msg.Container != nil {
		msg.Container.BuildTag(msg.Container.Tag())
		msg.Container.Tag().RenderOpening(out)
	}

	title := html.NewSnippet("p", `class="title"`).InsertHTML(msg.Title)
	title.Tag().AddClassesIf(msg.TitleSize > 0 && msg.TitleSize <= 6, "is-"+strconv.Itoa(msg.TitleSize))
	msg.RenderChildSnippet(out, title)

	subtitle := html.NewSnippet("p", `class="subtitle"`).InsertHTML(msg.Subtitle)
	subtitle.Tag().AddClassesIf(msg.SubtitleSize > 0 && msg.SubtitleSize <= 6, "is-"+strconv.Itoa(msg.SubtitleSize))
	msg.RenderChildSnippet(out, subtitle)

	msg.RenderChildSnippetIf(msg.CTA != nil, out, msg.CTA)

	if msg.Container != nil {
		msg.Container.Tag().RenderClosing(out)
	}

	html.WriteString(out, `</div>`)

	if msg.InsideFoot != nil {
		html.WriteString(out, `<div class="hero-foot">`)
		msg.RenderChildSnippet(out, msg.InsideFoot)
		html.WriteString(out, `</div>`)
	}

	return nil
}
