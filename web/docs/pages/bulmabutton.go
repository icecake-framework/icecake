package webdocs

import (
	"io"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/html"
)

const (
	linkBulmaButton string = `<a href="https://bulma.io/documentation/components/button">bulma Button</a>`
)

type SectionBulmaButton struct {
	SectionIcecakeDoc
}

func (cmp *SectionBulmaButton) RenderContent(out io.Writer) error {
	html.WriteString(out, `<h2>Icecake Bulma Button</h2>`)
	html.WriteStrings(out, `<p>bulma.Button is an icecake snippet providing the HTML rendering for a `, linkBulmaButton, ` with extra features and usefull Go APIs.</p>`)

	html.WriteString(out, `<h3>bulma.button API</h3>`)

	html.WriteString(out, `<p><strong>Title HTMLString</strong> renders the button content.</p>`)
	html.WriteString(out, `<p><strong>IsOutlined bool</strong> renders an outlined button</p>`)
	html.WriteString(out, `<p><strong>IsRounded bool</strong> renders a rounded button</p>`)

	html.WriteString(out, `<p><strong>IsDisabled</strong> renders disabled button</p>`)

	html.WriteString(out, `<h3>Usage</h3>`)

	// usage
	html.WriteString(out, `<div class="box mr-5">`)
	us1 := &bulma.Button{Title: *html.ToHTML("Click Here")}
	us2 := &bulma.Button{Title: *html.ToHTML("Click Here")}
	html.Render(out, cmp, us1, us2)

	html.WriteString(out, `</div>`)

	return nil
}
