package ick

import "strings"

type COLOR string

const (
	COLOR_NONE         COLOR  = ""
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
	COLOR_SUCCESS      COLOR  = "is-success"
	COLOR_WARNING      COLOR  = "is-warning"
	COLOR_DANGER       COLOR  = "is-danger"
	COLOR_GREY         COLOR  = "is-grey"
	COLOR_GREY_DARKER  COLOR  = "is-grey-darker"
	COLOR_GREY_DARK    COLOR  = "is-grey-dark"
	COLOR_GREY_LIGHT   COLOR  = "is-grey-light"
	COLOR_GREY_LIGHTER COLOR  = "is-grey-lighter"
	COLOR_OPTIONS      string = string(COLOR_GREY + " " + COLOR_WHITE + " " + COLOR_WHITE_BIS + " " + COLOR_WHITE_TER + " " + COLOR_BLACK + " " + COLOR_BLACK_BIS + " " + COLOR_BLACK_TER + " " + COLOR_LIGHT + " " + COLOR_DARK + " " + COLOR_PRIMARY + " " + COLOR_LINK + " " + COLOR_INFO + " " + COLOR_SUCCESS + " " + COLOR_WARNING + " " + COLOR_DANGER + " " + COLOR_GREY_DARKER + " " + COLOR_GREY_DARK + " " + COLOR_GREY_LIGHT + " " + COLOR_GREY_LIGHTER)
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

type TXTCOLOR string

const (
	TXTCOLOR_WHITE     TXTCOLOR = "has-text-white"
	TXTCOLOR_WHITE_BIS TXTCOLOR = "has-text-white-bis"
	TXTCOLOR_WHITE_TER TXTCOLOR = "has-text-white-ter"
	TXTCOLOR_BLACK     TXTCOLOR = "has-text-black"
	TXTCOLOR_BLACK_BIS TXTCOLOR = "has-text-black-bis"
	TXTCOLOR_BLACK_TER TXTCOLOR = "has-text-black-ter"
	TXTCOLOR_LIGHT     TXTCOLOR = "has-text-light"
	TXTCOLOR_DARK      TXTCOLOR = "has-text-dark"

	TXTCOLOR_PRIMARY TXTCOLOR = "has-text-primary"
	TXTCOLOR_LINK    TXTCOLOR = "has-text-link"
	TXTCOLOR_INFO    TXTCOLOR = "has-text-info"
	TXTCOLOR_SUCCESS TXTCOLOR = "has-text-success"
	TXTCOLOR_WARNING TXTCOLOR = "has-text-warning"
	TXTCOLOR_DANGER  TXTCOLOR = "has-text-danger"

	TXTCOLOR_PRIMARY_LIGHT TXTCOLOR = "has-text-primary-light"
	TXTCOLOR_LINK_LIGHT    TXTCOLOR = "has-text-link-light"
	TXTCOLOR_INFO_LIGHT    TXTCOLOR = "has-text-info-light"
	TXTCOLOR_SUCCESS_LIGHT TXTCOLOR = "has-text-success-light"
	TXTCOLOR_WARNING_LIGHT TXTCOLOR = "has-text-warning-light"
	TXTCOLOR_DANGER_LIGHT  TXTCOLOR = "has-text-danger-light"

	TXTCOLOR_PRIMARY_DARK TXTCOLOR = "has-text-primary-dark"
	TXTCOLOR_LINK_DARK    TXTCOLOR = "has-text-link-dark"
	TXTCOLOR_INFO_DARK    TXTCOLOR = "has-text-info-dark"
	TXTCOLOR_SUCCESS_DARK TXTCOLOR = "has-text-success-dark"
	TXTCOLOR_WARNING_DARK TXTCOLOR = "has-text-warning-dark"
	TXTCOLOR_DANGER_DARK  TXTCOLOR = "has-text-danger-dark"

	TXTCOLOR_GREY         TXTCOLOR = "has-text-grey"
	TXTCOLOR_GREY_DARKER  TXTCOLOR = "has-text-grey-darker"
	TXTCOLOR_GREY_DARK    TXTCOLOR = "has-text-grey-dark"
	TXTCOLOR_GREY_LIGHT   TXTCOLOR = "has-text-grey-light"
	TXTCOLOR_GREY_LIGHTER TXTCOLOR = "has-text-grey-lighter"
	TXTCOLOR_OPTIONS      string   = string(TXTCOLOR_GREY+" "+TXTCOLOR_WHITE+" "+TXTCOLOR_WHITE_BIS+" "+TXTCOLOR_WHITE_TER+" "+TXTCOLOR_BLACK+" "+TXTCOLOR_BLACK_BIS+" "+TXTCOLOR_BLACK_TER+" "+TXTCOLOR_LIGHT+" "+TXTCOLOR_DARK+" "+TXTCOLOR_PRIMARY+" "+TXTCOLOR_LINK+" "+TXTCOLOR_INFO+" "+TXTCOLOR_SUCCESS+" "+TXTCOLOR_WARNING+" "+TXTCOLOR_DANGER+" "+TXTCOLOR_GREY_DARKER+" "+TXTCOLOR_GREY_DARK+" "+TXTCOLOR_GREY_LIGHT+" "+TXTCOLOR_GREY_LIGHTER) +
		string(TXTCOLOR_PRIMARY_LIGHT+" "+TXTCOLOR_LINK_LIGHT+" "+TXTCOLOR_INFO_LIGHT+" "+TXTCOLOR_SUCCESS_LIGHT+" "+TXTCOLOR_WARNING_LIGHT+" "+TXTCOLOR_DANGER_LIGHT) +
		string(TXTCOLOR_PRIMARY_DARK+" "+TXTCOLOR_LINK_DARK+" "+TXTCOLOR_INFO_DARK+" "+TXTCOLOR_SUCCESS_DARK+" "+TXTCOLOR_WARNING_DARK+" "+TXTCOLOR_DANGER_DARK)
)

func TextColor(c TXTCOLOR) *TXTCOLOR {
	return &c
}

type BKGCOLOR string

const (
	BKGCOLOR_WHITE     BKGCOLOR = "has-background-white"
	BKGCOLOR_WHITE_BIS BKGCOLOR = "has-background-white-bis"
	BKGCOLOR_WHITE_TER BKGCOLOR = "has-background-white-ter"
	BKGCOLOR_BLACK     BKGCOLOR = "has-background-black"
	BKGCOLOR_BLACK_BIS BKGCOLOR = "has-background-black-bis"
	BKGCOLOR_BLACK_TER BKGCOLOR = "has-background-black-ter"
	BKGCOLOR_LIGHT     BKGCOLOR = "has-background-light"
	BKGCOLOR_DARK      BKGCOLOR = "has-background-dark"

	BKGCOLOR_PRIMARY BKGCOLOR = "has-background-primary"
	BKGCOLOR_LINK    BKGCOLOR = "has-background-link"
	BKGCOLOR_INFO    BKGCOLOR = "has-background-info"
	BKGCOLOR_SUCCESS BKGCOLOR = "has-background-success"
	BKGCOLOR_WARNING BKGCOLOR = "has-background-warning"
	BKGCOLOR_DANGER  BKGCOLOR = "has-background-danger"

	BKGCOLOR_PRIMARY_LIGHT BKGCOLOR = "has-background-primary-light"
	BKGCOLOR_LINK_LIGHT    BKGCOLOR = "has-background-link-light"
	BKGCOLOR_INFO_LIGHT    BKGCOLOR = "has-background-info-light"
	BKGCOLOR_SUCCESS_LIGHT BKGCOLOR = "has-background-success-light"
	BKGCOLOR_WARNING_LIGHT BKGCOLOR = "has-background-warning-light"
	BKGCOLOR_DANGER_LIGHT  BKGCOLOR = "has-background-danger-light"

	BKGCOLOR_PRIMARY_DARK BKGCOLOR = "has-background-primary-dark"
	BKGCOLOR_LINK_DARK    BKGCOLOR = "has-background-link-dark"
	BKGCOLOR_INFO_DARK    BKGCOLOR = "has-background-info-dark"
	BKGCOLOR_SUCCESS_DARK BKGCOLOR = "has-background-success-dark"
	BKGCOLOR_WARNING_DARK BKGCOLOR = "has-background-warning-dark"
	BKGCOLOR_DANGER_DARK  BKGCOLOR = "has-background-danger-dark"

	BKGCOLOR_GREY         BKGCOLOR = "has-background-grey"
	BKGCOLOR_GREY_DARKER  BKGCOLOR = "has-background-grey-darker"
	BKGCOLOR_GREY_DARK    BKGCOLOR = "has-background-grey-dark"
	BKGCOLOR_GREY_LIGHT   BKGCOLOR = "has-background-grey-light"
	BKGCOLOR_GREY_LIGHTER BKGCOLOR = "has-background-grey-lighter"
	BKGCOLOR_OPTIONS      string   = string(BKGCOLOR_GREY+" "+BKGCOLOR_WHITE+" "+BKGCOLOR_WHITE_BIS+" "+BKGCOLOR_WHITE_TER+" "+BKGCOLOR_BLACK+" "+BKGCOLOR_BLACK_BIS+" "+BKGCOLOR_BLACK_TER+" "+BKGCOLOR_LIGHT+" "+BKGCOLOR_DARK+" "+BKGCOLOR_PRIMARY+" "+BKGCOLOR_LINK+" "+BKGCOLOR_INFO+" "+BKGCOLOR_SUCCESS+" "+BKGCOLOR_WARNING+" "+BKGCOLOR_DANGER+" "+BKGCOLOR_GREY_DARKER+" "+BKGCOLOR_GREY_DARK+" "+BKGCOLOR_GREY_LIGHT+" "+BKGCOLOR_GREY_LIGHTER) +
		string(BKGCOLOR_PRIMARY_LIGHT+" "+BKGCOLOR_LINK_LIGHT+" "+BKGCOLOR_INFO_LIGHT+" "+BKGCOLOR_SUCCESS_LIGHT+" "+BKGCOLOR_WARNING_LIGHT+" "+BKGCOLOR_DANGER_LIGHT) +
		string(BKGCOLOR_PRIMARY_DARK+" "+BKGCOLOR_LINK_DARK+" "+BKGCOLOR_INFO_DARK+" "+BKGCOLOR_SUCCESS_DARK+" "+BKGCOLOR_WARNING_DARK+" "+BKGCOLOR_DANGER_DARK)
)

func BackgroundColor(c BKGCOLOR) *BKGCOLOR {
	return &c
}

type SIZE string

const (
	SIZE_SMALL   SIZE   = "is-small"
	SIZE_STD     SIZE   = ""
	SIZE_MEDIUM  SIZE   = "is-medium"
	SIZE_LARGE   SIZE   = "is-large"
	SIZE_OPTIONS string = string(SIZE_SMALL + " " + SIZE_MEDIUM + " " + SIZE_LARGE)
)
