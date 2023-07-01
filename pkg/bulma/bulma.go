package bulma

// "github.com/icecake-framework/icecake/pkg/html"

// TODO: handle bulma properties for color, size, display

func init() {
	// html.RequireCSSFile("https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css")
}

type COLOR string

const (
	COLOR_WHITE        COLOR  = "is-white"
	COLOR_WHITE_BIS    COLOR  = "is-white-bis"
	COLOR_WHITE_TER    COLOR  = "is-white-ter"
	COLOR_BLACK        COLOR  = "is-black"
	COLOR_BLACK_BIS    COLOR  = "is-black-bis"
	COLOR_BLACK_TER    COLOR  = "is-black-ter"
	COLOR_LIGHT        COLOR  = "is-light"
	COLOR_DARK         COLOR  = "is-dark"
	COLOR_PRIMARY      COLOR  = "is-primary"
	COLOR_LINK         COLOR  = "is-link"
	COLOR_INFO         COLOR  = "is-info"
	COLOR_SUCCESS      COLOR  = "is-sucess"
	COLOR_WARNING      COLOR  = "is-warning"
	COLOR_DANGER       COLOR  = "is-danger"
	COLOR_GREY_DARKER  COLOR  = "is-grey-darker"
	COLOR_GREY_DARK    COLOR  = "is-grey-dark"
	COLOR_GREY_LIGHT   COLOR  = "is-grey-light"
	COLOR_GREY_LIGHTER COLOR  = "is-grey-lighter"
	COLOR_OPTIONS      string = string(COLOR_WHITE + " " + COLOR_WHITE_BIS + " " + COLOR_WHITE_TER + " " + COLOR_BLACK + " " + COLOR_BLACK_BIS + " " + COLOR_BLACK_TER + " " + COLOR_LIGHT + " " + COLOR_DARK + " " + COLOR_PRIMARY + " " + COLOR_LINK + " " + COLOR_INFO + " " + COLOR_SUCCESS + " " + COLOR_WARNING + " " + COLOR_DANGER + " " + COLOR_GREY_DARKER + " " + COLOR_GREY_DARK + " " + COLOR_GREY_LIGHT + " " + COLOR_GREY_LIGHTER)
)
