package ick

import "strings"

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

func Color(c COLOR) *COLOR {
	return &c
}

func (c *COLOR) SetLight(f bool) *COLOR {
	if c.IsLight() != f {
		if f {
			*c += COLOR(" is-light")
		} else {
			*c = COLOR(strings.Replace(string(*c), "is-light", "", -1))
		}
	}
	return c
}

func (c COLOR) IsLight() bool {
	return strings.Contains(string(c), "is-light")
}

type SIZE string

const (
	SIZE_SMALL   SIZE   = "is-small"
	SIZE_STD     SIZE   = ""
	SIZE_MEDIUM  SIZE   = "is-medium"
	SIZE_LARGE   SIZE   = "is-large"
	SIZE_OPTIONS string = string(SIZE_SMALL + " " + SIZE_MEDIUM + " " + SIZE_LARGE)
)
