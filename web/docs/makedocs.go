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

	// get the command line parameters
	output := flag.String("output", "", "output path where generated html files will be saved")
	flag.BoolVar(&verbose.IsOn, "verbose", false, "print out execution details")
	flag.BoolVar(&verbose.IsDebugging, "debug", false, "print out debugging info")
	flag.Parse()
	path := html.MustCheckOutputPath(output)

	html.RequireCSSFile("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.3/font/bootstrap-icons.css")

	dfthtmlfile := html.NewHtmlFile("en").
		AddHeadItem("meta", "charset=UTF-8").
		AddHeadItem("meta", `http-equiv="X-UA-Compatible" content="IE=edge"`).
		AddHeadItem("meta", `name="viewport" content="width=device-width, initial-scale=1.0"`).
		AddHeadItem("script", `type="text/javascript" src="icecake.js"`)

	navbar := &bulma.Navbar{HasShadow: true}
	navbar.Tag().SetId("topnavbar")
	navbar.AddItems(
		(&bulma.NavbarItem{
			ItemType: bulma.NAVBARIT_BRAND,
			Content:  html.HTML(`<span class="title pl-2">Icecake</span>`)}).
			ParseImageSrc("/assets/icecake-color.svg").
			ParseHRef("/"),

		(&bulma.NavbarItem{
			ItemType: bulma.NAVBARIT_START,
			Content:  html.HTML(`Docs`)}).
			ParseHRef("/docs"),

		(&bulma.NavbarItem{
			ItemType: bulma.NAVBARIT_END,
			Content:  bulma.NewButtonLink(*html.HTML("GitHub"), "https://github.com/icecake-framework/icecake")}),

		(&bulma.NavbarItem{
			ItemType: bulma.NAVBARIT_END,
			Content:  html.HTML("<small>Alpha</small>")}),
	)

	hero := &bulma.Hero{
		Height:    bulma.HH_FULLFHEIGHT_WITH_NAVBAR,
		Title:     *html.HTML("Develop SPA and Static Websites in Go."),
		TitleSize: 2,
		Subtitle:  *html.HTML("Pure Go Web Assembly Framework"),
	}
	hero.Container = &bulma.Container{FullWidth: bulma.CFW_MAXDESKTOP}
	hero.Container.Tag().AddClasses("has-text-centered")
	btnCTA := bulma.NewButtonLink(*html.HTML("Read doc"), "/docs")
	btnCTA.Tag().SetId("cta")
	hero.CTA = btnCTA

	var pages []*html.Page
	pgindex := html.NewPage("index.html").Stack(navbar, hero, &docsFooter{})
	pgindex.Title = "icecake framework"
	pgindex.Description = "Develop SPA and Static Websites in with a pure Go Web Assembly Framework"
	pages = append(pages, pgindex)

	pgdocs := html.NewPage("docs.html").Stack(navbar, &docsFooter{})
	pgdocs.Title = "documentation - icecake framework"
	pgdocs.Description = "go Web Assembly Framework documentation"
	pages = append(pages, pgdocs)

	n := 0
	for _, pg := range pages {
		dfthtmlfile.Title = pg.Title
		dfthtmlfile.Description = pg.Description
		dfthtmlfile.Body = html.NewSnippet("body").Stack(pg)
		// FIXME, convert url in file name ?!!
		err := dfthtmlfile.WriteHTMLFile(path, pg.URL().String())
		if err != nil {
			fmt.Println("makedoc fails")
			if !verbose.IsOn {
				fmt.Println("use the verbose flag to get more info")
			}
			os.Exit(1)
		}
		n++
	}
	fmt.Println(n, "files generated")
}
