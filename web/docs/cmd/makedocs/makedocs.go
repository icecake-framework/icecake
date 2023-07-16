package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/icecake-framework/icecake/internal/helper"
	"github.com/icecake-framework/icecake/pkg/ick"
	"github.com/icecake-framework/icecake/pkg/ickcore"
	"github.com/icecake-framework/icecake/web/docs"
	webdocs "github.com/icecake-framework/icecake/web/docs/pages"
	"github.com/joho/godotenv"
	"github.com/lolorenzo777/verbose"
)

func main() {

	// get the command line parameters
	outpathparam := flag.String("output", "", "output path where generated html files will be saved")
	flag.BoolVar(&verbose.IsOn, "verbose", false, "print out execution details")
	flag.BoolVar(&verbose.IsDebugging, "debug", false, "print out debugging info")
	strenv := flag.String("env", "dev", ".env environement file to load, with the path and without the extension. dev by default.")
	flag.Parse()

	// load environment variables
	if strenv == nil {
		strenv = new(string)
		*strenv = "dev"
	}
	errenv := godotenv.Load(*strenv + ".env")
	if errenv != nil {
		log.Fatalf("Error loading .env variables: %s", errenv)
	}

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
	pgindex.AddHeadItem("script", `type="text/javascript" src="`+web.ToAbsURLString("/assets/icecake.js")+`"`)

	// ... with a hero section
	hero := &ick.ICKHero{
		Height:   ick.HH_FULLFHEIGHT_WITH_NAVBAR,
		Title:    *ick.Title(2, "Develop SPA and Static Websites in Go."),
		Subtitle: *ick.Title(4, "Pure Go Web Assembly Framework"),
		Centered: true,
		CWidth:   ick.CONTWIDTH_MAXDESKTOP,
		CTA:      *ick.Button("Read doc").SetId("cta").SetHRef(*web.ToAbsURL("/docoverview.html")).SetColor(ick.COLOR_PRIMARY),
	}

	pgindex.Body().Append(
		docs.DocNavbar(pgindex).SetActiveItem("home"),
		hero,
		docs.DocFooter(pgindex))

	// menu for each pages unless home
	menu := ick.Menu("docmenu", `class="p-2`, `style="background-color:#fdfdfd;"`).
		SetType(ick.MENUTYP_NAV).
		SetSize(ick.SIZE_SMALL)
	menu.AddItem("", ick.MENUIT_LABEL, "General")
	menu.AddItem("docoverview", ick.MENUIT_LINK, "Overview").HRef = web.ToAbsURL("/docoverview.html")
	menu.AddItem("", ick.MENUIT_LABEL, "Composers")
	menu.AddItem("docinterfaces", ick.MENUIT_LINK, "interfaces")
	menu.AddItem("dochtmlstring", ick.MENUIT_LINK, "HTMLString")
	menu.AddItem("docbaresnippet", ick.MENUIT_LINK, "BareSnippet")
	menu.AddItem("docpage", ick.MENUIT_LINK, "Page")
	menu.AddItem("", ick.MENUIT_LABEL, "Core Snippets")
	menu.AddItem("docbutton", ick.MENUIT_LINK, "Button").HRef = web.ToAbsURL("/docbutton.html")
	menu.AddItem("doccard", ick.MENUIT_LINK, "Card").HRef = web.ToAbsURL("/doccard.html")
	menu.AddItem("docdelete", ick.MENUIT_LINK, "Delete").HRef = web.ToAbsURL("/docdelete.html")
	menu.AddItem("dochero", ick.MENUIT_LINK, "Hero").HRef = web.ToAbsURL("/dochero.html")
	menu.AddItem("docimage", ick.MENUIT_LINK, "Image").HRef = web.ToAbsURL("/docimage.html")
	menu.AddItem("docmenu", ick.MENUIT_LINK, "Menu").HRef = web.ToAbsURL("/docmenu.html")
	menu.AddItem("docmessage", ick.MENUIT_LINK, "Message").HRef = web.ToAbsURL("/docmessage.html")
	menu.AddItem("docnavbar", ick.MENUIT_LINK, "Navbar").HRef = web.ToAbsURL("/docnavbar.html")
	menu.AddItem("docinput", ick.MENUIT_LINK, "Input").HRef = web.ToAbsURL("/docinput.html")
	menu.AddItem("docicon", ick.MENUIT_LINK, "Icon").HRef = web.ToAbsURL("/docicon.html")
	menu.AddItem("doctaglabel", ick.MENUIT_LINK, "Tag Label").HRef = web.ToAbsURL("/doctaglabel.html")
	menu.AddItem("docmedia", ick.MENUIT_LINK, "media").HRef = web.ToAbsURL("/docmedia.html")
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
	addPageDoc(web, menu.Clone(), "docinput")
	addPageDoc(web, menu.Clone(), "docicon")
	addPageDoc(web, menu.Clone(), "doctaglabel")
	addPageDoc(web, menu.Clone(), "docmedia")

	// required files
	ickcore.RequireCSSFile("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css")
	ickcore.RequireCSSFile(web.ToAbsURLString("/assets/docs.css"))

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

func addPageDoc(web *ick.WebSite, menu *ick.ICKMenu, pgkey string) {
	pg := web.AddPage("en", pgkey)
	pg.AddHeadItem("meta", "charset=UTF-8")
	pg.AddHeadItem("meta", `http-equiv="X-UA-Compatible" content="IE=edge"`)
	pg.AddHeadItem("meta", `name="viewport" content="width=device-width, initial-scale=1.0"`)
	pg.AddHeadItem("script", `type="text/javascript" src="/assets/icecake.js"`)

	inside := ick.Elem("div", `class="columns is-mobile mb-0 pb-0"`,
		ick.Elem("div", `class="column is-narrow mb-0 pb-0"`, menu.SetActiveItem(pgkey)),
		ick.Elem("div", `class="column mb-0 pb-0"`, webdocs.SectionDoc(pgkey)))

	pg.Body().Append(
		docs.DocNavbar(pg).SetActiveItem("docs"),
		inside,
		docs.DocFooter(pg))

}
