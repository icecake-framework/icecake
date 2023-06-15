package html

import (
	"fmt"
	"io"
)

// testsnippetid
type testsnippetid struct{ HTMLSnippet }

func (s testsnippetid) Tag() *Tag {
	s.Tag().SetName("span")
	s.Tag().Attributes().SetId("IdTemplate1")
	return s.Tag()
}

// testsnippet1
type testsnippet1 struct {
	HTMLSnippet
	Html HTMLString
}

func (tst *testsnippet1) WriteBody(out io.Writer) error {
	err := tst.UnfoldHTML(out, tst.Html, nil)
	return err
}

// testsnippet2
type testsnippet2 struct{ testsnippet1 }

func (testsnippet2) Tag() *Tag {
	return NewTag("div", TryParseAttributes(`class="ts2a ts2b" tabIndex=2 style="display=test;" a2`))
}

// testsnippet3
type testsnippet3 struct{ HTMLSnippet }

func (s testsnippet3) WriteBody(out io.Writer) error {
	_, err := WriteString(out, fmt.Sprintf("data.app=%s", s.dataState.App.(string)))
	return err
}

// testsnippetinfinite
type testsnippetinfinite struct{ HTMLSnippet }

func (s testsnippetinfinite) WriteBody(out io.Writer) error {
	return s.UnfoldHTML(out, "<ick-test-infinite/>", nil)
}
