package ick

import "github.com/icecake-framework/icecake/pkg/html"

type CONTAINER_FULLWIDTH string

const (
	CFW_NONE          CONTAINER_FULLWIDTH = ""
	CFW_MAXDESKTOP    CONTAINER_FULLWIDTH = "is-max-desktop"    // 960px in any cases
	CFW_DESKTOP       CONTAINER_FULLWIDTH = "is-desktop"        // 960px or 1152 px or 1344 px
	CFW_MAXWIDESCREEN CONTAINER_FULLWIDTH = "is-max-widescreen" // 1152 px
	CFW_WIDESCREEN    CONTAINER_FULLWIDTH = "is-widescreen"     // 1152 px or 1344 px
	CFW_FULLHD        CONTAINER_FULLWIDTH = "is-fullhd"         // 1344 px
	CFW_FLUID         CONTAINER_FULLWIDTH = "is-fluid"          // fullscreen + 32px margin
	CFW_OPTIONS       CONTAINER_FULLWIDTH = CFW_MAXDESKTOP + " " + CFW_DESKTOP + " " + CFW_MAXWIDESCREEN + " " + CFW_WIDESCREEN + " " + CFW_FULLHD + " " + CFW_FLUID
)

// Container allow centering element on larger viewport. See [bulma container]
//
// [bulma container]: https://bulma.io/documentation/layout/container/
type Container struct {
	html.HTMLSnippet
	FullWidth CONTAINER_FULLWIDTH
}

// Ensure Container implements HTMLComposer interface
var _ html.HTMLComposer = (*Container)(nil)

func (c *Container) BuildTag() html.Tag {
	c.Tag().SetTagName("div").AddClass("container").PickClass(string(CFW_OPTIONS), string(c.FullWidth))
	return *c.Tag()
}
