package ick

type CONTAINER_WIDTH string

const (
	CONTWIDTH_NONE          CONTAINER_WIDTH = ""
	CONTWIDTH_MAXDESKTOP    CONTAINER_WIDTH = "is-max-desktop"    // 960px in any cases
	CONTWIDTH_DESKTOP       CONTAINER_WIDTH = "is-desktop"        // 960px or 1152 px or 1344 px
	CONTWIDTH_MAXWIDESCREEN CONTAINER_WIDTH = "is-max-widescreen" // 1152 px
	CONTWIDTH_WIDESCREEN    CONTAINER_WIDTH = "is-widescreen"     // 1152 px or 1344 px
	CONTWIDTH_FULLHD        CONTAINER_WIDTH = "is-fullhd"         // 1344 px
	CONTWIDTH_FLUID         CONTAINER_WIDTH = "is-fluid"          // fullscreen + 32px margin
	CONTWIDTH_OPTIONS       CONTAINER_WIDTH = CONTWIDTH_MAXDESKTOP + " " + CONTWIDTH_DESKTOP + " " + CONTWIDTH_MAXWIDESCREEN + " " + CONTWIDTH_WIDESCREEN + " " + CONTWIDTH_FULLHD + " " + CONTWIDTH_FLUID
)

// Container allow centering element on larger viewport. See [bulma container]
//
// [bulma container]: https://bulma.io/documentation/layout/container/
func Container(w CONTAINER_WIDTH, attrs ...string) *ICKElem {
	e := Elem("div", `class="container`)
	e.Tag().PickClass(string(CONTWIDTH_OPTIONS), string(w))
	e.Tag().ParseAttributes(attrs...)
	return e
}
