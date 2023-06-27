package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/icecake-framework/icecake/internal/helper"
	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/web/docs"
	docoverview "github.com/icecake-framework/icecake/web/docs/pages/overview"
	"github.com/sunraylab/verbose"
)

func main() {

	// get the command line parameters
	outpathparam := flag.String("output", "", "output path where generated html files will be saved")
	flag.BoolVar(&verbose.IsOn, "verbose", false, "print out execution details")
	flag.BoolVar(&verbose.IsDebugging, "debug", false, "print out debugging info")
	flag.Parse()

	// init new website
	outpath := helper.MustCheckOutputPath(outpathparam)
	web := html.NewWebSite(outpath)

	// page index
	pgindex := web.AddPage("en", "index")
	pgindex.Title = "icecake framework"
	pgindex.Description = "Develop SPA and Static Websites in with a pure Go Web Assembly Framework"
	pgindex.AddHeadItem("meta", "charset=UTF-8")
	pgindex.AddHeadItem("meta", `http-equiv="X-UA-Compatible" content="IE=edge"`)
	pgindex.AddHeadItem("meta", `name="viewport" content="width=device-width, initial-scale=1.0"`)
	pgindex.AddHeadItem("script", `type="text/javascript" src="/assets/icecake.js"`)

	// ... with a hero section
	hero := &bulma.Hero{
		Height:    bulma.HH_FULLFHEIGHT_WITH_NAVBAR,
		Title:     *html.ToHTML("Develop SPA and Static Websites in Go."),
		TitleSize: 2,
		Subtitle:  *html.ToHTML("Pure Go Web Assembly Framework"),
		Container: &bulma.Container{FullWidth: bulma.CFW_MAXDESKTOP},
		CTA:       bulma.NewButton(*html.ToHTML("Read doc"), "cta", "/docs.html").SetColor(bulma.COLOR_PRIMARY),
	}
	hero.Container.Tag().AddClasses("has-text-centered")

	pgindex.Body = html.NewSnippet("body").Stack(
		docs.MyNavbar().SetActiveItem("home"),
		hero,
		docs.MyFooter())

	// page docs
	pgdocs := web.AddPage("en", "docs")
	pgdocs.Title = "documentation - icecake framework"
	pgdocs.Description = "go Web Assembly Framework documentation"
	pgdocs.AddHeadItem("meta", "charset=UTF-8")
	pgdocs.AddHeadItem("meta", `http-equiv="X-UA-Compatible" content="IE=edge"`)
	pgdocs.AddHeadItem("meta", `name="viewport" content="width=device-width, initial-scale=1.0"`)
	pgdocs.AddHeadItem("script", `type="text/javascript" src="/assets/icecake.js"`)

	// ... with a menu
	menu := bulma.Menu{TagName: "nav"}
	menu.Tag().SetId("docmenu").AddClasses("p-2").SetStyle("background-color:#fdfdfd;font-size:0.8rem;")
	menu.AddItem("", bulma.MENUIT_LABEL, "General")
	menu.AddItem("OVERVIEW", bulma.MENUIT_LINK, "Overview").ParseHRef("/docs.html")
	menu.AddItem("", bulma.MENUIT_LABEL, "Core Snippets")
	menu.AddItem("", bulma.MENUIT_LINK, "HTMLString")
	menu.AddItem("", bulma.MENUIT_LINK, "HTMLSnippet")
	menu.AddItem("", bulma.MENUIT_LINK, "HTMLPage")
	menu.AddItem("", bulma.MENUIT_LABEL, "Bulma Snippets")
	menu.AddItem("", bulma.MENUIT_LINK, "Button")
	menu.AddItem("", bulma.MENUIT_LINK, "Message")
	menu.AddItem("", bulma.MENUIT_LINK, "Navbar")
	menu.Item("OVERVIEW").IsActive = true

	p1bodyc := html.NewSnippet("div", `class="columns is-mobile mb-0 pb-0"`)
	p1bodyc.InsertSnippet("div", `class="column is-narrow mb-0 pb-0"`).Stack(&menu)
	p1bodyc.InsertSnippet("div", `class="column mb-0 pb-0"`).Stack(&docoverview.Section{})

	pgdocs.Body = html.NewSnippet("body").Stack(
		docs.MyNavbar().SetActiveItem("docs"),
		p1bodyc,
		docs.MyFooter())

	// files

	html.RequireCSSFile("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.3/font/bootstrap-icons.css")
	html.RequireCSSFile("/assets/docs.css")

	// copy assets
	err := web.CopyToAssets("./web/docs/assets/", "./web/docs/sass/docs.css", "./web/docs/sass/docs.css.map")
	if err != nil {
		fmt.Println("makedoc fails: ", err.Error())
		os.Exit(1)
	}

	// writing files
	n, err := web.WriteFiles()
	if err != nil {
		fmt.Println("makedoc fails")
		if !verbose.IsOn {
			fmt.Println("use the verbose flag to get more info")
		}
		os.Exit(1)
	}
	fmt.Println(n, "website pages generated")
}
