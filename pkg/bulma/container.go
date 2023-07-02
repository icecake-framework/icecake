package bulma

import "github.com/icecake-framework/icecake/pkg/html"

type CONTAINER_FULLWIDTH string

const (
	CFW_MAXDESKTOP    CONTAINER_FULLWIDTH = "max-desktop"    // 960px in any cases
	CFW_DESKTOP       CONTAINER_FULLWIDTH = "desktop"        // 960px or 1152 px or 1344 px
	CFW_MAXWIDESCREEN CONTAINER_FULLWIDTH = "max-widescreen" // 1152 px
	CFW_WIDESCREEN    CONTAINER_FULLWIDTH = "widescreen"     // 1152 px or 1344 px
	CFW_FULLHD        CONTAINER_FULLWIDTH = "fullhd"         // 1344 px
	CFW_FLUID         CONTAINER_FULLWIDTH = "fluid"          // fullscreen + 32px margin
)

// Container allow centering element on larger viewport.
type Container struct {
	tag       html.Tag
	FullWidth CONTAINER_FULLWIDTH
}

// Ensure Container implements HTMLTagComposer interface
var _ html.TagBuilder = (*Container)(nil)

// Tag returns a reference to the snippet tag.
func (s *Container) Tag() *html.Tag {
	if s.tag.AttributeMap == nil {
		s.tag.AttributeMap = make(html.AttributeMap)
	}
	return &s.tag
}

func (c *Container) BuildTag(tag *html.Tag) {
	tag.SetTagName("div").
		AddClass("container").
		SetClassIf(c.FullWidth != "", "is-"+string(c.FullWidth))
}
