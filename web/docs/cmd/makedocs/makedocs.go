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
	web := ick.NewWebSite(outpath)

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
		Title:         *ick.Title(2, "Develop SPA and Static Websites in Go."),
		Subtitle:      *ick.Title(4, "Pure Go Web Assembly Framework"),
		ContainerAttr: html.ParseAttributes(`class="has-text-centered ` + string(ick.CFW_MAXDESKTOP) + `"`),
		CTA:           ick.Button("Read doc", "cta").ParseHRef("/docoverview.html").SetColor(ick.COLOR_PRIMARY),
	}

	pgindex.Body().AddContent(
		docs.DocNavbar().SetActiveItem("home"),
		hero,
		docs.DocFooter())

	// menu for each pages unless home
	menu := ick.IckMenu{}
	menu.MenuTag().SetTagName("nav").AddClass("is-small")
	menu.Tag().SetId("docmenu").AddClass("p-2").AddStyle("background-color:#fdfdfd;")
	menu.AddItem("", ick.MENUIT_LABEL, "General")
	menu.AddItem("docoverview", ick.MENUIT_LINK, "Overview").ParseHRef("/docoverview.html")
	menu.AddItem("", ick.MENUIT_LABEL, "Composers")
	menu.AddItem("", ick.MENUIT_LINK, "HTMLString")
	menu.AddItem("", ick.MENUIT_LINK, "HTMLSnippet")
	menu.AddItem("", ick.MENUIT_LINK, "HTMLPage")
	menu.AddItem("", ick.MENUIT_LABEL, "Core Snippets")
	menu.AddItem("docbutton", ick.MENUIT_LINK, "Button").ParseHRef("/docbutton.html")
	menu.AddItem("doccard", ick.MENUIT_LINK, "Card").ParseHRef("/doccard.html")
	menu.AddItem("docdelete", ick.MENUIT_LINK, "Delete").ParseHRef("/docdelete.html")
	menu.AddItem("dochero", ick.MENUIT_LINK, "Hero").ParseHRef("/dochero.html")
	menu.AddItem("docimage", ick.MENUIT_LINK, "Image").ParseHRef("/docimage.html")
	menu.AddItem("docmenu", ick.MENUIT_LINK, "Menu").ParseHRef("/docmenu.html")
	menu.AddItem("docmessage", ick.MENUIT_LINK, "Message").ParseHRef("/docmessage.html")
	menu.AddItem("docnavbar", ick.MENUIT_LINK, "Navbar").ParseHRef("/docnavbar.html")
	menu.AddItem("docnotify", ick.MENUIT_LINK, "Notify").ParseHRef("/docnotify.html")
	menu.AddItem("", ick.MENUIT_FOOTER, "Alpha 4")

	// page docs
	addPageDoc(web, menu.Clone(), "docoverview")
	addPageDoc(web, menu.Clone(), "docbutton")
	addPageDoc(web, menu.Clone(), "doccard")
	addPageDoc(web, menu.Clone(), "docdelete")
	addPageDoc(web, menu.Clone(), "dochero")
	addPageDoc(web, menu.Clone(), "docimage")
	addPageDoc(web, menu.Clone(), "docmenu")
	addPageDoc(web, menu.Clone(), "docmessage")
	addPageDoc(web, menu.Clone(), "docnavbar")
	addPageDoc(web, menu.Clone(), "docnotify")

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

func addPageDoc(web *ick.WebSite, menu *ick.IckMenu, pgkey string) {
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
