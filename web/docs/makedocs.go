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
	outpathparam := flag.String("output", "", "output path where generated html files will be saved")
	flag.BoolVar(&verbose.IsOn, "verbose", false, "print out execution details")
	flag.BoolVar(&verbose.IsDebugging, "debug", false, "print out debugging info")
	flag.Parse()
	outpath := html.MustCheckOutputPath(outpathparam)

	// init pages
	var pages []html.Page
	pgindex := html.NewPage("index")
	pgindex.Title = "icecake framework"
	pgindex.Description = "Develop SPA and Static Websites in with a pure Go Web Assembly Framework"
	pages = append(pages, *pgindex)

	pgdocs := html.NewPage("docs.html")
	pgdocs.Title = "documentation - icecake framework"
	pgdocs.Description = "go Web Assembly Framework documentation"
	pages = append(pages, *pgdocs)

	// setup the common navbar
	navbar := bulma.Navbar{HasShadow: true}
	navbar.Tag().SetId("topnavbar")

	navbari0 := bulma.NavbarItem{
		ItemType: bulma.NAVBARIT_BRAND,
		Content:  html.HTML(`<span class="title pl-2">Icecake</span>`)}
	navbari0.ParseHRef("/")
	navbari0.ParseImageSrc("/assets/icecake-color.svg")

	navbarihome := bulma.NavbarItem{
		ItemType: bulma.NAVBARIT_START,
		Content:  html.HTML(`Home`),
		HRef:     pgindex.RelURL()}

	navbaridocs := bulma.NavbarItem{
		ItemType: bulma.NAVBARIT_START,
		Content:  html.HTML(`Docs`),
		HRef:     pgdocs.RelURL()}

	navbari3 := bulma.NavbarItem{
		ItemType: bulma.NAVBARIT_END,
		Content:  bulma.NewButtonLink(*html.HTML("GitHub"), "https://github.com/icecake-framework/icecake")}

	navbari4 := bulma.NavbarItem{
		ItemType: bulma.NAVBARIT_END,
		Content:  html.HTML("<small>Alpha</small>")}

	navbar.AddItems(&navbari0, &navbarihome, &navbaridocs, &navbari3, &navbari4)

	// Build the pgindex content
	hero := &bulma.Hero{
		Height:    bulma.HH_FULLFHEIGHT_WITH_NAVBAR,
		Title:     *html.HTML("Develop SPA and Static Websites in Go."),
		TitleSize: 2,
		Subtitle:  *html.HTML("Pure Go Web Assembly Framework"),
		Container: &bulma.Container{FullWidth: bulma.CFW_MAXDESKTOP},
	}
	hero.Container.Tag().AddClasses("has-text-centered")
	btnCTA := bulma.NewButtonLink(*html.HTML("Read doc"), "/docs")
	btnCTA.Tag().SetId("cta")
	hero.CTA = btnCTA
	pages[0].Stack(&navbar, hero, &docsFooter{})

	// build the pgdocs content
	pages[1].Stack(&navbar, &docsFooter{})

	// writing files
	html.RequireCSSFile("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.3/font/bootstrap-icons.css")

	dfthtmlfile := html.NewHtmlFile("en").
		AddHeadItem("meta", "charset=UTF-8").
		AddHeadItem("meta", `http-equiv="X-UA-Compatible" content="IE=edge"`).
		AddHeadItem("meta", `name="viewport" content="width=device-width, initial-scale=1.0"`).
		AddHeadItem("script", `type="text/javascript" src="icecake.js"`)

	n := 0
	for i, pg := range pages {
		dfthtmlfile.Title = pg.Title
		dfthtmlfile.Description = pg.Description
		dfthtmlfile.Body = html.NewSnippet("body").Stack(&pg)
		switch i {
		case 0:
			navbarihome.IsActive = true
			navbaridocs.IsActive = false
		case 1:
			navbarihome.IsActive = false
			navbaridocs.IsActive = true
		}
		err := dfthtmlfile.WriteHTMLFile(outpath, pg.RelURL().String())
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
