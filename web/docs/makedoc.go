package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/sunraylab/verbose"
)

func main() {

	output := flag.String("output", "", "output path where generated html files will be saved")
	flag.BoolVar(&verbose.IsOn, "verbose", false, "print out execution details")
	flag.Parse()

	path := html.MustCheckOutputPath(output)

	index := html.NewPage("en").
		AddHeadItem("meta", "charset=UTF-8").
		AddHeadItem("meta", `http-equiv="X-UA-Compatible" content="IE=edge"`).
		AddHeadItem("meta", `name="viewport" content="width=device-width, initial-scale=1.0"`).
		AddHeadItem("link", `rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"`).
		AddHeadItem("link", `rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.3/font/bootstrap-icons.css"`).
		AddHeadItem("script", `type="text/javascript" src="icecake.js"`)

	index.Title = "documentation - icecake framework"
	index.Description = "go wasm framework"
	index.HTMLFileName = "index.html"

	navbar := &bulma.Navbar{HasShadow: true}
	navbar.AddItems(
		(&bulma.NavbarItem{
			ItemType: bulma.NAVBARIT_BRAND,
			Content:  html.NewHTML(`<span class="title pl-2">Icecake</span>`)}).
			ParseImageSrc("/assets/icecake-color.svg").
			ParseHRef("/"),

		(&bulma.NavbarItem{
			ItemType: bulma.NAVBARIT_START,
			Content:  html.NewHTML(`Docs`)}).
			ParseHRef("/"),

		(&bulma.NavbarItem{
			ItemType: bulma.NAVBARIT_END,
			Content:  bulma.NewButtonLink(html.HTML("GitHub"), "https://github.com/icecake-framework/icecake")}),

		(&bulma.NavbarItem{
			ItemType: bulma.NAVBARIT_END,
			Content:  html.NewHTML("<small>Alpha</small>")}),
	)

	hero := &bulma.Hero{
		Height:    bulma.HH_FULLFHEIGHT_WITH_NAVBAR,
		Title:     html.HTML("Develop SPA and Static Websites in Go."),
		TitleSize: 2,
		Subtitle:  html.HTML("Pure Go Web Assembly Framework"),
	}
	hero.Container = &bulma.Container{FullWidth: bulma.CFW_MAXDESKTOP}
	hero.Container.Tag().AddClasses("has-text-centered")
	hero.CTA = bulma.NewButtonLink(html.HTML("Read doc"), "/docs")

	body := html.NewSnippet("body", `id="body"`).StackContent(navbar, hero, &myFooter{})

	index.Body = body

	err := index.WriteHTMLFile(path)
	if err != nil {
		verbose.Error("makedoc", err)
		fmt.Println("makedoc fails")
		if !verbose.IsOn {
			fmt.Println("use the verbose flag to get more info")
		}
		os.Exit(1)
	}
}

/******************************************************************************
/* myFooter
******************************************************************************/

type myFooter struct{ html.HTMLSnippet }

func (cmp *myFooter) BuildTag(tag *html.Tag) { tag.SetTagName("footer").AddClasses("footer") }
func (cmp *myFooter) RenderContent(out io.Writer) error {

	hrefMIT := `<a href="https://opensource.org/licenses/mit-license.php" rel="license">MIT</a>`
	hrefCCBY := `<a href="https://creativecommons.org/licenses/by-nc-sa/4.0/">CC BY-NC-SA 4.0</a>`
	hrefLinks := []string{
		`<a href="/">Home</a>`,
		`<a href="/docs">Docs</a>`}

	html.WriteString(out, `<div class="container"><div class="columns">`)

	// 1st column
	html.WriteString(out, `<div class="column is-4">`)
	html.WriteStrings(out, `<h4 class="bd-footer-title">`, `<strong>IceCake</strong> by Lolorenzo`, `</h4>`)
	html.WriteStrings(out, `<div class="bd-footer-tsp">`, `Source code licences `, hrefMIT, `</div>`)
	html.WriteStrings(out, `<div class="bd-footer-tsp">`, `Website content licensed `, hrefCCBY, `</div>`)
	html.WriteString(out, `</div>`)

	// 2nd column
	html.WriteString(out, `<div class="column is-4">`)
	html.WriteStrings(out, `<h4 class="bd-footer-title">`, `<strong>Links</strong>`, `</h4>`)
	for _, hrefLink := range hrefLinks {
		html.WriteStrings(out, `<p class="bd-footer-link">`, hrefLink, `</p>`)
	}
	html.WriteString(out, `</div>`)

	html.WriteString(out, `</div></div>`)
	return nil
}
