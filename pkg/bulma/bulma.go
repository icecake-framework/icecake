package bulma

import (
	"github.com/icecake-framework/icecake/pkg/html"
)

// TODO: handle bulma properties for color, size, display

func init() {
	html.RequireCSSFile("https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css")
}
