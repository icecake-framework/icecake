package html

import (
	"fmt"
	"io"
	"net/url"
	"time"
)

// testsnippet0
type testsnippet0 struct{ HTMLSnippet }

func (s *testsnippet0) BuildTag() Tag { return *s.Tag().SetTagName("span") }

// testsnippetid
// type testsnippetid struct{ HTMLSnippet }

// func (s *testsnippetid) BuildTag() Tag {
// 	tag := NewTag("span", nil)
// 	tag.SetId("IdTemplate1")
// 	return *tag
// }

// testsnippet1
type testsnippet1 struct {
	meta RMetaData // Rendering MetaData
	HTML HTMLString
}

// Meta provides a reference to the RenderingMeta object associated with this composer.
// This is required by the icecake rendering process.
func (h *testsnippet1) RMeta() *RMetaData {
	return &h.meta
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// For an HTMLString snippet, RenderContent renders (unfold and generate HTML output) the internal string without enclosed tag.
// Use an HTMLSnippet snippet to renders the string inside an enclosed tag.
func (h *testsnippet1) RenderContent(out io.Writer) error {
	return renderHTML(out, h, h.HTML)
}

// func (tst *testsnippet1) RenderContent(out io.Writer) error {
// 	err := tst.RenderChild(out, &tst.Html)
// 	return err
// }

// testsnippet2
type testsnippet2 struct{ HTMLSnippet }

func (s *testsnippet2) BuildTag() Tag {
	s.Tag().SetTagName("div")
	s.Tag().ParseAttributes(`class="ts2a ts2b" tabindex=2 style="display=test;" a2`)
	return *s.Tag()
}

// func (tst *testsnippet2) RenderContent(out io.Writer) error {
// 	err := tst.RenderChild(out, &tst.Html)
// 	return err
// }

// testsnippet3
// type testsnippet3 struct{ HTMLSnippet }

// func (s *testsnippet3) RenderContent(out io.Writer) error {
// 	_, err := WriteString(out, fmt.Sprintf("data.app=%s", s.ds.App.(string)))
// 	return err
// }

type Unmanaged struct{}

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
	Unmanaged
}

func (s *testsnippet4) BuildTag() Tag {
	s.Tag().SetTagName("div")
	s.Tag().AddClassIf(s.IsOk, "ok")
	return *s.Tag()
}

func (s *testsnippet4) RenderContent(out io.Writer) error {
	WriteStringIf(s.Text != "", out, s.Text)
	s.RenderChild(out, &s.HTML)
	WriteStringIf(s.I != 0, out, fmt.Sprintf("%v", s.I))
	WriteStringIf(s.F != 0, out, fmt.Sprintf("%v", s.F))
	WriteStringIf(s.D != 0, out, fmt.Sprintf("%v", s.D+(time.Hour*1)))
	WriteStringIf(s.U != nil, out, fmt.Sprintf("<a href='%v'></a>", s.U))
	return nil
}

// testsnippetinfinite
type testsnippetinfinite struct{ HTMLSnippet }

func (s *testsnippetinfinite) RenderContent(out io.Writer) error {
	return s.RenderChild(out, ToHTML("<ick-testinfinite/>"))
}

// testsnippetinfinite
type testcustomcomposer struct{}

// Ensure testcustomcomposer implements HTMLComposer interface
var _ HTMLComposer = (*testcustomcomposer)(nil)

// func (s *testcustomcomposer) Tag() *Tag                                   { return nil }
func (s *testcustomcomposer) RMeta() *RMetaData                       { return &RMetaData{} }
func (s *testcustomcomposer) BuildTag() Tag                           { return Tag{} }
func (s *testcustomcomposer) SetAttribute(aname string, value string) {}
func (s *testcustomcomposer) RenderContent(out io.Writer) error       { return nil }
