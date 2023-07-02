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
	HH_SMALL                   HERO_HEIGHT = "is-small"
	HH_STANDARD                HERO_HEIGHT = ""
	HH_MEDIUM                  HERO_HEIGHT = "is-medium"
	HH_LARGE                   HERO_HEIGHT = "is-large"
	HH_HALFHEIGHT              HERO_HEIGHT = "is-halfheight"
	HH_FULLFHEIGHT             HERO_HEIGHT = "is-fullheight"
	HH_FULLFHEIGHT_WITH_NAVBAR HERO_HEIGHT = "is-fullheight-with-navbar"
	HH_OPTIONS                 string      = string(HH_SMALL + " " + HH_MEDIUM + " " + HH_LARGE + " " + HH_HALFHEIGHT + " " + HH_FULLFHEIGHT + " " + HH_FULLFHEIGHT_WITH_NAVBAR)
)

type Hero struct {
	html.HTMLSnippet

	Height HERO_HEIGHT // the height of the hero section,

	InsideHead html.HTMLComposer

	Container    *Container // optional Container to render the content, allowing center
	Title        html.HTMLString
	TitleSize    int // 1 to 6
	Subtitle     html.HTMLString
	SubtitleSize int // 1 to 6

	CTA html.HTMLComposer // Call To Action

	InsideFoot html.HTMLComposer
}

// Ensure Hero implements HTMLTagComposer interface
var _ html.HTMLTagComposer = (*Hero)(nil)

// Tag Builder used by the rendering functions.
func (msg *Hero) BuildTag(tag *html.Tag) {
	tag.SetTagName("section").AddClass("hero").PickClass(HH_OPTIONS, string(msg.Height))
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
func (msg *Hero) RenderContent(out io.Writer) error {

	if msg.InsideHead != nil {
		msg.RenderChilds(out, html.Div(`class="hero-head"`).SetBody(msg.InsideHead))
	}

	html.WriteString(out, `<div class="hero-body">`)

	if msg.Container != nil {
		msg.Container.BuildTag(msg.Container.Tag())
		msg.Container.Tag().RenderOpening(out)
	}

	title := html.P(`class="title"`).SetBody(&msg.Title)
	title.Tag().SetClassIf(msg.TitleSize > 0 && msg.TitleSize <= 6, "is-"+strconv.Itoa(msg.TitleSize))
	msg.RenderChilds(out, title)

	subtitle := html.P(`class="subtitle"`).SetBody(&msg.Subtitle)
	subtitle.Tag().SetClassIf(msg.SubtitleSize > 0 && msg.SubtitleSize <= 6, "is-"+strconv.Itoa(msg.SubtitleSize))
	msg.RenderChilds(out, subtitle)

	msg.RenderChildsIf(msg.CTA != nil, out, msg.CTA)

	if msg.Container != nil {
		msg.Container.Tag().RenderClosing(out)
	}

	html.WriteString(out, `</div>`)

	if msg.InsideFoot != nil {
		msg.RenderChilds(out, html.Div(`class="hero-foor"`).SetBody(msg.InsideFoot))
	}

	return nil
}
