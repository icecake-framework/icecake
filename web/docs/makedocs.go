package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/otiai10/copy"
	"github.com/sunraylab/verbose"
)

func main() {

	// get the command line parameters
	outpathparam := flag.String("output", "", "output path where generated html files will be saved")
	flag.BoolVar(&verbose.IsOn, "verbose", false, "print out execution details")
	flag.BoolVar(&verbose.IsDebugging, "debug", false, "print out debugging info")
	flag.Parse()
	outpath := html.MustCheckOutputPath(outpathparam)

	// copy all assets
	outassets := filepath.Join(outpath, "/assets/")
	os.RemoveAll(outassets)
	err := copy.Copy("./web/docs/assets/", outassets)
	if err == nil {
		err = copy.Copy("./web/docs/sass/docs.css", filepath.Join(outassets, "/docs.css"))
	}
	if err == nil {
		err = copy.Copy("./web/docs/sass/docs.css.map", filepath.Join(outassets, "/docs.css.map"))
	}
	if err != nil {
		fmt.Println("makedoc fails: ", err.Error())
		os.Exit(1)
	}

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
	navbar.Tag().SetId("topbar")
	navbar.AddItem("", bulma.NAVBARIT_BRAND, html.ToHTML(`<span class="title pl-2">Icecake</span>`)).ParseHRef("/").ParseImageSrc("/assets/icecake-color.svg")
	navbar.AddItem("home", bulma.NAVBARIT_START, html.ToHTML(`Home`)).HRef = pgindex.RelURL()
	navbar.AddItem("docs", bulma.NAVBARIT_START, html.ToHTML(`Docs`)).HRef = pgdocs.RelURL()

	btngit := &bulma.Button{Title: *html.ToHTML("GitHub")}
	btngit.ParseHRef("https://github.com/icecake-framework/icecake")
	btngit.SetColor(bulma.COLOR_PRIMARY).SetOutlined(true)

	navbar.AddItem("", bulma.NAVBARIT_END, btngit)
	navbar.AddItem("", bulma.NAVBARIT_END, html.ToHTML("<small>Alpha</small>"))

	// Build the pgindex content
	hero := &bulma.Hero{
		Height:    bulma.HH_FULLFHEIGHT_WITH_NAVBAR,
		Title:     *html.ToHTML("Develop SPA and Static Websites in Go."),
		TitleSize: 2,
		Subtitle:  *html.ToHTML("Pure Go Web Assembly Framework"),
		Container: &bulma.Container{FullWidth: bulma.CFW_MAXDESKTOP},
	}
	hero.Container.Tag().AddClasses("has-text-centered")

	btnCTA := &bulma.Button{Title: *html.ToHTML("Read doc")}
	btnCTA.HRef = pgdocs.RelURL()
	btnCTA.SetColor(bulma.COLOR_PRIMARY)
	btnCTA.Tag().SetId("cta")

	hero.CTA = btnCTA

	navbar0 := navbar.Clone()
	navbar0.Item("home").IsActive = true
	pages[0].Stack(navbar0, hero, &docsFooter{})

	// build the pgdocs menu
	menu := bulma.Menu{TagName: "nav"}
	menu.Tag().SetId("docmenu").AddClasses("p-2").SetStyle("background-color:#fdfdfd;font-size:0.8rem;")
	menu.AddItem("", bulma.MENUIT_LABEL, "General")
	menu.AddItem("OVERVIEW", bulma.MENUIT_LINK, "Overview").ParseHRef("/docs")
	menu.AddItem("", bulma.MENUIT_LABEL, "Core Snippets")
	menu.AddItem("", bulma.MENUIT_LINK, "HTMLString")
	menu.AddItem("", bulma.MENUIT_LINK, "HTMLSnippet")
	menu.AddItem("", bulma.MENUIT_LINK, "HTMLPage")
	menu.AddItem("", bulma.MENUIT_LABEL, "Bulma Snippets")
	menu.AddItem("", bulma.MENUIT_LINK, "Button")
	menu.AddItem("", bulma.MENUIT_LINK, "Message")
	menu.AddItem("", bulma.MENUIT_LINK, "Navbar")
	menu.Item("OVERVIEW").IsActive = true

	navbar1 := navbar.Clone()
	navbar1.Item("docs").IsActive = true

	p1bodyc := html.NewSnippet("div", `class="columns is-mobile mb-0 pb-0"`)
	p1bodyc.InsertSnippet("div", `class="column is-narrow mb-0 pb-0"`).Stack(&menu)
	p1bodyc.InsertSnippet("div", `class="column mb-0 pb-0"`).Stack(&docOverview{})

	pages[1].Stack(navbar1, p1bodyc, &docsFooter{})

	// files

	// set default header for all files
	dfthtmlfile := html.NewHtmlFile("en").
		AddHeadItem("meta", "charset=UTF-8").
		AddHeadItem("meta", `http-equiv="X-UA-Compatible" content="IE=edge"`).
		AddHeadItem("meta", `name="viewport" content="width=device-width, initial-scale=1.0"`).
		AddHeadItem("script", `type="text/javascript" src="/assets/icecake.js"`)

	html.RequireCSSFile("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.3/font/bootstrap-icons.css")
	html.RequireCSSFile("/assets/docs.css")

	// writing files
	n := 0
	for _, pg := range pages {
		dfthtmlfile.Title = pg.Title
		dfthtmlfile.Description = pg.Description
		dfthtmlfile.Body = html.NewSnippet("body").Stack(&pg)
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
