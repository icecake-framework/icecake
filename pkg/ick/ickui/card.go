package ickui

import (
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
)

type ICKCard struct {
	ick.ICKCard
	dom.UI
}

// Ensure ICKCard implements UIComposer interface
var _ dom.UIComposer = (*ICKCard)(nil)

// Card main factory
func Card(content html.ContentComposer) *ICKCard {
	c := new(ICKCard)
	c.ICKCard = *ick.Card(content)
	return c
}
