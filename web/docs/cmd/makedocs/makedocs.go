package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/icecake-framework/icecake/internal/helper"
	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/web/docs"
	webdocs "github.com/icecake-framework/icecake/web/docs/pages"
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
		CTA:       bulma.NewButton(*html.ToHTML("Read doc"), "cta", "/overview.html").SetColor(bulma.COLOR_PRIMARY),
	}
	hero.Container.Tag().AddClasses("has-text-centered")

	pgindex.Body = html.NewSnippet("body").Stack(
		docs.MyNavbar().SetActiveItem("home"),
		hero,
		docs.MyFooter())

	// menu for each pages unless home
	menu := bulma.Menu{TagName: "nav"}
	menu.Tag().SetId("docmenu").AddClasses("p-2").SetStyle("background-color:#fdfdfd;font-size:0.8rem;")
	menu.AddItem("", bulma.MENUIT_LABEL, "General")
	menu.AddItem("overview", bulma.MENUIT_LINK, "Overview").ParseHRef("/overview.html")
	menu.AddItem("", bulma.MENUIT_LABEL, "Core Snippets")
	menu.AddItem("", bulma.MENUIT_LINK, "HTMLString")
	menu.AddItem("", bulma.MENUIT_LINK, "HTMLSnippet")
	menu.AddItem("", bulma.MENUIT_LINK, "HTMLPage")
	menu.AddItem("", bulma.MENUIT_LABEL, "Bulma Snippets")
	menu.AddItem("bulmabutton", bulma.MENUIT_LINK, "Button").ParseHRef("/bulmabutton.html")
	menu.AddItem("bulmacard", bulma.MENUIT_LINK, "Card").ParseHRef("/bulmacard.html")
	menu.AddItem("bulmadelete", bulma.MENUIT_LINK, "Delete").ParseHRef("/bulmadelete.html")
	menu.AddItem("bulmahero", bulma.MENUIT_LINK, "Hero").ParseHRef("/bulmahero.html")
	menu.AddItem("bulmaimage", bulma.MENUIT_LINK, "Image").ParseHRef("/bulmaimage.html")
	menu.AddItem("bulmamenu", bulma.MENUIT_LINK, "Menu").ParseHRef("/bulmamenu.html")
	menu.AddItem("bulmamessage", bulma.MENUIT_LINK, "Message").ParseHRef("/bulmamessage.html")
	menu.AddItem("bulmanavbar", bulma.MENUIT_LINK, "Navbar").ParseHRef("/bulmanavbar.html")
	menu.AddItem("bulmanotify", bulma.MENUIT_LINK, "Notify").ParseHRef("/bulmanotify.html")

	// page docs
	addPageDoc(web, menu.Clone(), "overview")
	addPageDoc(web, menu.Clone(), "bulmabutton")
	addPageDoc(web, menu.Clone(), "bulmacard")
	addPageDoc(web, menu.Clone(), "bulmadelete")
	addPageDoc(web, menu.Clone(), "bulmahero")
	addPageDoc(web, menu.Clone(), "bulmaimage")
	addPageDoc(web, menu.Clone(), "bulmamenu")
	addPageDoc(web, menu.Clone(), "bulmamessage")
	addPageDoc(web, menu.Clone(), "bulmanavbar")
	addPageDoc(web, menu.Clone(), "bulmanotify")

	// required files
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

func addPageDoc(web *html.WebSite, menu *bulma.Menu, pgkey string) {
	pg := web.AddPage("en", pgkey)
	pg.AddHeadItem("meta", "charset=UTF-8")
	pg.AddHeadItem("meta", `http-equiv="X-UA-Compatible" content="IE=edge"`)
	pg.AddHeadItem("meta", `name="viewport" content="width=device-width, initial-scale=1.0"`)
	pg.AddHeadItem("script", `type="text/javascript" src="/assets/icecake.js"`)

	pgc := html.NewSnippet("div", `class="columns is-mobile mb-0 pb-0"`)
	pgc.InsertSnippet("div", `class="column is-narrow mb-0 pb-0"`).Stack(menu.SetActiveItem(pgkey))
	pgc.InsertSnippet("div", `class="column mb-0 pb-0"`).Stack(webdocs.NewSectionIcecakeDoc(pgkey))

	pg.Body = html.NewSnippet("body").Stack(
		docs.MyNavbar().SetActiveItem("docs"),
		pgc,
		docs.MyFooter())

}