package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/sunraylab/verbose"
)

func main() {

	output := flag.String("output", "", "output path where generated html files will be saved")
	flag.BoolVar(&verbose.IsOn, "verbose", false, "print out execution details")
	flag.Parse()

	var path string
	if output != nil && *output != "" {
		path, _ = filepath.Abs(*output)
		fileInfo, err := os.Stat(path)
		if err != nil || !fileInfo.IsDir() {
			err := fmt.Errorf("output %s is not a valid path", *output)
			verbose.Error("makedoc", err)
			fmt.Println("makedoc fails. use the verbose flag to get more info.")
			os.Exit(1)
		}
	}

	index := html.NewHtml5Page("en")
	index.Title = "documentation - icecake framework"
	index.AddHeadMeta(`http-equiv="X-UA-Compatible" content="IE=edge"`)
	index.AddHeadMeta(`name="viewport" content="width=device-width, initial-scale=1.0"`)
	index.AddHeadLink(`rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"`)
	index.AddHeadLink(`rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.3/font/bootstrap-icons.css"`)
	index.AddHeadScript(`type="text/javascript" src="icecake.js"`)

	navbar := &html.Navbar{HasShadow: true}
	navbar.AddItems(
		(&html.NavbarItem{
			ItemType: html.NAVBARIT_BRAND,
			Body:     `<span class="title pl-2">Icecake</span>`}).
			ParseImageSrc("/assets/icecake-color.svg").
			ParseHRef("/"),
		(&html.NavbarItem{
			ItemType: html.NAVBARIT_START,
			Body:     `DOCS`}),
		(&html.NavbarItem{
			ItemType: html.NAVBARIT_END,
			Body:     html.NewButtonLink("GitHub", "https://github.com/icecake-framework/icecake").String(),
		}))
	// git.Body , _, _ = html.RenderSnippet(html.NewButtonLink("GitHub", "https://github.com/icecake-framework/icecake"))

	// FIXME: confusing name "Body"
	index.Body = navbar.String()

	err := index.WriteHTMLFile(path, "index.html")
	if err != nil {
		verbose.Error("makedoc", err)
		fmt.Println("makedoc fails. use the verbose flag to get more info.")
		os.Exit(1)
	}
}
