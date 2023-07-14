package ickui

import (
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
)

type ICKCard struct {
	ick.ICKCard
	dom.UI
}

// Ensure ICKCard implements UIComposer interface
var _ dom.UIComposer = (*ICKCard)(nil)

// Card main factory
func Card(content ickcore.ContentComposer) *ICKCard {
	c := new(ICKCard)
	c.ICKCard = *ick.Card(content)
	return c
}
