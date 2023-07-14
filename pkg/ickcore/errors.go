package ickcore

import (
	"errors"
	"fmt"
)

var (
	ErrBodyTagMissing            = errors.New("<body> tag missing")
	ErrTooManyRecursiveRendering = errors.New("too many recursive rendering")
	ErrNameMissing               = errors.New("'opening <ick-' tag found without name")
)

type IckTagNameError struct {
	TagName string
	Message string
}

func (e *IckTagNameError) Error() string {
	return fmt.Sprintf("%s: %s", e.TagName, e.Message)
}
