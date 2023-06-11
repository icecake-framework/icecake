package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/icecake-framework/icecake/pkg/html"
	"github.com/icecake-framework/icecake/pkg/ui"
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
	index.AddMetaHttpEquiv("X-UA-Compatible", "IE=edge")
	index.AddMetaName("viewport", "width=device-width, initial-scale=1.0")
	index.AddMetaLink(map[string]string{"rel": "stylesheet", "href": "https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"})
	index.AddMetaLink(map[string]string{"rel": "stylesheet", "href": "https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.3/font/bootstrap-icons.css"})
	index.AddMetaScript("icecake.js")

	navbar := &ui.Navbar{}
	index.Body, _, _ = html.RenderHTMLSnippet(navbar)

	err := index.WriteHTMLFile(path, "index.html")
	if err != nil {
		verbose.Error("makedoc", err)
		fmt.Println("makedoc fails. use the verbose flag to get more info.")
		os.Exit(1)
	}
}
