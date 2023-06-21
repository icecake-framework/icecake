package html

import (
	"errors"
	"fmt"
)

var (
	ErrTaglessParent             = errors.New("tagless parent")
	ErrTooManyRecursiveRendering = errors.New("too many recursive rendering")
	ErrIckTagNameMissing         = errors.New("'opening <ick-' tag found without name")
)

type IckTagNameError struct {
	TagName string
	Message string
}

func (e *IckTagNameError) Error() string {
	return fmt.Sprintf("%s: %s", e.TagName, e.Message)
}
