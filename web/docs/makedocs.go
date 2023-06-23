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

	index := html.NewHtmlFile("en").
		AddHeadItem("meta", "charset=UTF-8").
		AddHeadItem("meta", `http-equiv="X-UA-Compatible" content="IE=edge"`).
		AddHeadItem("meta", `name="viewport" content="width=device-width, initial-scale=1.0"`).
		AddHeadItem("script", `type="text/javascript" src="icecake.js"`)

	index.Title = "documentation - icecake framework"
	index.Description = "go wasm framework"
	index.HTMLFileName = "index.html"

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
			ParseHRef("/"),

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

	body := html.NewSnippet("body").StackContent(navbar, hero, &docsFooter{}) // `id="body"`

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
