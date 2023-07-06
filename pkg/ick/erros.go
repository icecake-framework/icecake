package ick

import "errors"

var (
	ErrBadHtmlFileExtention = errors.New("bad html file extension")
	ErrMissingFileName      = errors.New("missing file name")
)
