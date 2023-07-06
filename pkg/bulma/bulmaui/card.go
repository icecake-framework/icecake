package bulmaui

import (
	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
)

type ICKCard struct {
	bulma.ICKCard
	dom.UI
}

// Ensure ICKCard implements HTMLComposer interface
var _ html.HTMLComposer = (*ICKCard)(nil)

// Card main factory
func Card(content html.HTMLContentComposer) *ICKCard {
	c := new(ICKCard)
	c.ICKCard = *bulma.Card(content)
	return c
}
