package html

import (
	"os"
	"path/filepath"

	"github.com/icecake-framework/icecake/internal/helper"
	"github.com/sunraylab/verbose"
)

type WebSite struct {
	pages map[string]*Page

	OutPath string // output path where generated websites files will be saved

}

func NewWebSite(outpath string) *WebSite {
	w := new(WebSite)
	w.pages = make(map[string]*Page)
	w.OutPath = outpath

	return w
}

func (w *WebSite) CopyToAssets(srcs ...string) error {
	outassets := filepath.Join(w.OutPath, "/assets/")
	os.RemoveAll(outassets)
	os.MkdirAll(outassets, os.ModePerm)
	for _, src := range srcs {
		err := helper.CopyFiles(outassets, src)
		if err != nil {
			return verbose.Error("Website.CopyToAssets", err)
		}
	}
	return nil
}

func (w *WebSite) AddPage(lang string, rawUrl string) *Page {
	pg := NewPage(lang, rawUrl)
	if pg == nil {
		return nil
	}
	if w.pages == nil {
		w.pages = make(map[string]*Page)
	}
	w.pages[pg.RelURL().String()] = pg
	return pg
}

func (w *WebSite) Page(rawUrl string) *Page {
	return w.pages[rawUrl]
}

// returns the number of pages written and errors.
func (w WebSite) WriteFiles() (n int, err error) {
	if w.pages == nil {
		return 0, nil
	}

	n = 0
	for _, p := range w.pages {
		err = p.WriteFile(w.OutPath)
		if err != nil {
			return n, err
		}
		n++
	}
	return n, nil
}
