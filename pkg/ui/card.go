package ui

import (
	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/dom"
)

type Card struct {
	bulma.Card
	DOM dom.Element
}
