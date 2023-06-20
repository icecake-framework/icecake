package html

import (
	"io"
)

// testsnippet0
type testsnippet0 struct{ HTMLSnippet }

func (s *testsnippet0) BuildTag(tag *Tag) { tag.SetName("span") }

// testsnippetid
type testsnippetid struct{ HTMLSnippet }

func (s *testsnippetid) BuildTag(tag *Tag) {
	tag.SetName("span").Attributes().SetId("IdTemplate1")
}

// testsnippet1
type testsnippet1 struct {
	HTMLSnippet
	Html HTMLString
}

func (tst *testsnippet1) RenderContent(out io.Writer) error {
	err := tst.RenderChildHTML(out, tst.Html)
	return err
}

// testsnippet2
type testsnippet2 struct{ testsnippet1 }

func (testsnippet2) BuildTag(tag *Tag) {
	tag.SetName("div")
	tag.ParseAttributes(`class="ts2a ts2b" tabindex=2 style="display=test;" a2`)
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
}

func (s *testsnippet4) BuildTag(tag *Tag) {
	tag.SetName("div")
	tag.Attributes().AddClassesIf(s.IsOk, "ok")
}

func (s *testsnippet4) RenderContent(out io.Writer) error {
	WriteStringsIf(s.Text != "", out, s.Text)
	s.RenderChildHTML(out, s.HTML)
	return nil
}

// testsnippetinfinite
type testsnippetinfinite struct{ HTMLSnippet }

func (s *testsnippetinfinite) RenderContent(out io.Writer) error {
	return s.RenderChildHTML(out, *NewString("<ick-testinfinite/>"))
}
