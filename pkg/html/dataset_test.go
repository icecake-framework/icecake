package html

import (
	"fmt"
	"io"
	"net/url"
	"time"
)

// testsnippet0
type testsnippet0 struct{ HTMLSnippet }

func (s *testsnippet0) BuildTag(tag *Tag) { tag.SetTagName("span") }

// testsnippetid
type testsnippetid struct{ HTMLSnippet }

func (s *testsnippetid) BuildTag(tag *Tag) {
	tag.SetTagName("span").SetId("IdTemplate1")
}

// testsnippet1
type testsnippet1 struct {
	HTMLSnippet
	Html HTMLString
}

func (tst *testsnippet1) RenderContent(out io.Writer) error {
	err := tst.RenderChilds(out, &tst.Html)
	return err
}

// testsnippet2
type testsnippet2 struct{ testsnippet1 }

func (testsnippet2) BuildTag(tag *Tag) {
	tag.SetTagName("div")
	tag.ParseAttributes(`class="ts2a ts2b" tabindex=2 style="display=test;" a2`)
}

func (tst *testsnippet2) RenderContent(out io.Writer) error {
	err := tst.RenderChilds(out, &tst.Html)
	return err
}

// testsnippet3
// type testsnippet3 struct{ HTMLSnippet }

// func (s *testsnippet3) RenderContent(out io.Writer) error {
// 	_, err := WriteString(out, fmt.Sprintf("data.app=%s", s.ds.App.(string)))
// 	return err
// }

// testsnippet4
type testsnippet4 struct {
	HTMLSnippet
	IsOk bool
	Text string
	HTML HTMLString
	I    int
	F    float64
	D    time.Duration
	U    *url.URL
}

func (s *testsnippet4) BuildTag(tag *Tag) {
	tag.SetTagName("div").AddClassesIf(s.IsOk, "ok")
}

func (s *testsnippet4) RenderContent(out io.Writer) error {
	WriteStringsIf(s.Text != "", out, s.Text)
	s.RenderChilds(out, &s.HTML)
	WriteStringsIf(s.I != 0, out, fmt.Sprintf("%v", s.I))
	WriteStringsIf(s.F != 0, out, fmt.Sprintf("%v", s.F))
	WriteStringsIf(s.D != 0, out, fmt.Sprintf("%v", s.D+(time.Hour*1)))
	WriteStringsIf(s.U != nil, out, fmt.Sprintf("<a href='%v'></a>", s.U))
	return nil
}

// testsnippetinfinite
type testsnippetinfinite struct{ HTMLSnippet }

func (s *testsnippetinfinite) RenderContent(out io.Writer) error {
	return s.RenderChilds(out, ToHTML("<ick-testinfinite/>"))
}

// testsnippetinfinite
type testcustomcomposer struct{}

func (s *testcustomcomposer) Tag() *Tag                            { return nil }
func (s *testcustomcomposer) BuildTag(tag *Tag)                    {}
func (s *testcustomcomposer) RenderContent(out io.Writer) error    { return nil }
func (s *testcustomcomposer) Embedded() ComposerMap                { return nil }
func (s *testcustomcomposer) Embed(id string, subcmp HTMLComposer) {}
