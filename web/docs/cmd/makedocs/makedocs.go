package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/icecake-framework/icecake/internal/helper"
	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/web/docs"
	webdocs "github.com/icecake-framework/icecake/web/docs/pages"
	"github.com/lolorenzo777/verbose"
)

func main() {

	// get the command line parameters
	outpathparam := flag.String("output", "", "output path where generated html files will be saved")
	flag.BoolVar(&verbose.IsOn, "verbose", false, "print out execution details")
	flag.BoolVar(&verbose.IsDebugging, "debug", false, "print out debugging info")
	flag.Parse()

	start := time.Now()

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
	hero := &ick.Hero{
		Height:        ick.HH_FULLFHEIGHT_WITH_NAVBAR,
		Title:         *html.ToHTML("Develop SPA and Static Websites in Go."),
		TitleSize:     2,
		Subtitle:      *html.ToHTML("Pure Go Web Assembly Framework"),
		ContainerAttr: html.ParseAttributes(`class="has-text-centered ` + string(ick.CFW_MAXDESKTOP) + `"`),
		CTA:           ick.Button(*html.ToHTML("Read doc"), "cta", "/overview.html").SetColor(ick.COLOR_PRIMARY),
	}

	pgindex.Body().AddContent(
		docs.DocNavbar().SetActiveItem("home"),
		hero,
		docs.DocFooter())

	// menu for each pages unless home
	menu := ick.Menu{}
	menu.MenuTag().SetTagName("nav").AddClass("is-small")
	menu.Tag().SetId("docmenu").AddClass("p-2").AddStyle("background-color:#fdfdfd;")
	menu.AddItem("", ick.MENUIT_LABEL, "General")
	menu.AddItem("overview", ick.MENUIT_LINK, "Overview").ParseHRef("/overview.html")
	menu.AddItem("", ick.MENUIT_LABEL, "Core Snippets")
	menu.AddItem("", ick.MENUIT_LINK, "HTMLString")
	menu.AddItem("", ick.MENUIT_LINK, "HTMLSnippet")
	menu.AddItem("", ick.MENUIT_LINK, "HTMLPage")
	menu.AddItem("", ick.MENUIT_LABEL, "Bulma Snippets")
	menu.AddItem("bulmabutton", ick.MENUIT_LINK, "Button").ParseHRef("/bulmabutton.html")
	menu.AddItem("bulmacard", ick.MENUIT_LINK, "Card").ParseHRef("/bulmacard.html")
	menu.AddItem("bulmadelete", ick.MENUIT_LINK, "Delete").ParseHRef("/bulmadelete.html")
	menu.AddItem("bulmahero", ick.MENUIT_LINK, "Hero").ParseHRef("/bulmahero.html")
	menu.AddItem("bulmaimage", ick.MENUIT_LINK, "Image").ParseHRef("/bulmaimage.html")
	menu.AddItem("bulmamenu", ick.MENUIT_LINK, "Menu").ParseHRef("/bulmamenu.html")
	menu.AddItem("bulmamessage", ick.MENUIT_LINK, "Message").ParseHRef("/bulmamessage.html")
	menu.AddItem("bulmanavbar", ick.MENUIT_LINK, "Navbar").ParseHRef("/bulmanavbar.html")
	menu.AddItem("bulmanotify", ick.MENUIT_LINK, "Notify").ParseHRef("/bulmanotify.html")
	menu.AddItem("", ick.MENUIT_FOOTER, "Alpha 4")

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

	fmt.Println(n, "pages generated in ", time.Since(start))
}

func addPageDoc(web *html.WebSite, menu *ick.Menu, pgkey string) {
	pg := web.AddPage("en", pgkey)
	pg.AddHeadItem("meta", "charset=UTF-8")
	pg.AddHeadItem("meta", `http-equiv="X-UA-Compatible" content="IE=edge"`)
	pg.AddHeadItem("meta", `name="viewport" content="width=device-width, initial-scale=1.0"`)
	pg.AddHeadItem("script", `type="text/javascript" src="/assets/icecake.js"`)

	inside := html.Div(`class="columns is-mobile mb-0 pb-0"`).AddContent(
		html.Div(`class="column is-narrow mb-0 pb-0"`).AddContent(
			menu.SetActiveItem(pgkey)),
		html.Div(`class="column mb-0 pb-0"`).AddContent(
			webdocs.SectionDoc(pgkey)))

	pg.Body().AddContent(
		docs.DocNavbar().SetActiveItem("docs"),
		inside,
		docs.DocFooter())

}
