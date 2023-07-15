package ick

import (
	"net/url"
	"os"
	"path/filepath"

	"github.com/icecake-framework/icecake/internal/helper"
	"github.com/lolorenzo777/verbose"
)

type WebSite struct {
	pages map[string]*Page

	OutPath string   // output path where generated websites files will be saved
	WebURL  *url.URL // website URL
}

func NewWebSite(outpath string) *WebSite {
	w := new(WebSite)
	w.pages = make(map[string]*Page)
	w.OutPath = outpath

	uenv := os.Getenv("WEB_URL")
	if uenv != "" {
		if uu, err := url.Parse(uenv); err != nil {
			verbose.Error("WEB_URL parameter", err)
		} else {
			w.WebURL = uu
		}
	}

	return w
}

func (w WebSite) ToAbsURL(rawurl string) *url.URL {
	if rawurl == "" {
		return nil
	}
	u, err := url.Parse(rawurl)
	if err != nil {
		verbose.Error("MakeAbsURL:", err)
		return nil
	}
	if w.WebURL != nil {
		u = w.WebURL.JoinPath(u.String())
	}
	return u
}

func (w WebSite) ToAbsURLString(rawurl string) string {
	u := w.ToAbsURL(rawurl)
	if u == nil {
		return ""
	}
	return u.String()
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
	pg := NewPage(w, lang, rawUrl)
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

	var warn error
	n = 0
	for _, p := range w.pages {

		// write file with its rendered content
		err = p.WriteFile(w.OutPath)
		if err != nil {
			return n, err
		}
		n++
	}
	return n, warn
}
