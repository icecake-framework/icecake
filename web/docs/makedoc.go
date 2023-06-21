package main

import (
	"flag"
	"fmt"
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
		AddHeadMeta("charset=UTF-8").
		AddHeadMeta(`http-equiv="X-UA-Compatible" content="IE=edge"`).
		AddHeadMeta(`name="viewport" content="width=device-width, initial-scale=1.0"`).
		AddHeadLink(`rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"`).
		AddHeadLink(`rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.3/font/bootstrap-icons.css"`).
		AddHeadScript(`type="text/javascript" src="icecake.js"`)

	index.Title = "documentation - icecake framework"
	index.Description = "go wasm framework"
	index.HTMLFileName = "index.html"

	navbar := &bulma.Navbar{HasShadow: true}
	navbar.AddItems(
		(&bulma.NavbarItem{
			ItemType: bulma.NAVBARIT_BRAND,
			Content:  html.NewString(`<span class="title pl-2">Icecake</span>`)}).
			ParseImageSrc("/assets/icecake-color.svg").
			ParseHRef("/"),

		(&bulma.NavbarItem{
			ItemType: bulma.NAVBARIT_START,
			Content:  html.NewString(`DOCS`)}).
			ParseHRef("/"),

		(&bulma.NavbarItem{
			ItemType: bulma.NAVBARIT_END,
			Content:  bulma.NewButtonLink(html.String("GitHub"), "https://github.com/icecake-framework/icecake")}),

		(&bulma.NavbarItem{
			ItemType: bulma.NAVBARIT_END,
			Content:  html.NewString("<small>Alpha</small>")}),
	)

	body := html.NewSnippet("body", html.ParseAttributes("id=body")).
		AddContent(navbar)

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
