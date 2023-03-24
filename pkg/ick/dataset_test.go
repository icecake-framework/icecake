package ick

import "fmt"

type testsnippet0 struct{ HtmlSnippet }

func (testsnippet0) Template(*DataState) (_t SnippetTemplate) { _t.TagName = "span"; return }

type testsnippet1 struct {
	HtmlSnippet
	HTML
}

func (tst testsnippet1) Template(*DataState) (_t SnippetTemplate) { _t.Body = tst.HTML; return }

type testsnippet2 struct{ testsnippet1 }

func (testsnippet2) Template(*DataState) (_t SnippetTemplate) {
	_t.TagName = "div"
	_t.Attributes = `class="ts2a ts2b" tabIndex=2 style="display=test;" a2`
	return
}

type testsnippet3 struct{ HtmlSnippet }

func (testsnippet3) Template(_data *DataState) (_t SnippetTemplate) {
	strapp, _ := _data.App.(string)
	_t.Body = HTML(fmt.Sprintf("data.app=%s", strapp))
	return
}

type testsnippetinfinite struct{ HtmlSnippet }

func (testsnippetinfinite) Template(_data *DataState) (_t SnippetTemplate) {
	_t.Body = "<ick-test-infinite/>"
	return
}
