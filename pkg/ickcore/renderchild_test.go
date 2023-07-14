package ickcore

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type snip0 struct {
	RMetaData
}

func (snip0) NeedRendering() bool                { return true }
func (*snip0) RenderContent(out io.Writer) error { return nil }

type snip1 struct {
	BareSnippet
	Test int
}

func (s *snip1) BuildTag() Tag {
	switch s.Test {
	case 1:
		s.Tag().SetTagName("span")
	case 2:
		s.Tag().SetTagName("div")
		s.Tag().ParseAttributes(`class="ts2a ts2b" tabindex=2 style="display=test;" a2`)
	default:
	}
	return *s.Tag()
}

func TestRenderChild(t *testing.T) {

	ResetRegistry()
	out := new(bytes.Buffer)

	// empty snippet
	s0 := new(snip0)
	err := RenderChild(out, nil, s0)
	require.NoError(t, err)
	require.Empty(t, out)

	// tag + empty body
	out.Reset()
	s1 := &snip1{Test: 1}
	err = RenderChild(out, nil, s1)
	require.NoError(t, err)
	assert.Equal(t, `<span name="snip1"></span>`, out.String())

	// tag with attributes
	out.Reset()
	s1.Tag().SetTabIndex(1).AddStyle("color=red;").SetId("helloworld").AddClass("test").SetAttribute("myatt", "myval")
	RenderChild(out, nil, s1)
	assert.Equal(t, `<span id="helloworld" name="snip1" class="test" myatt="myval" style="color=red;" tabindex=1></span>`, out.String())

	// without id
	out.Reset()
	s1.Tag().SetId("")
	RenderChild(out, nil, s1)
	assert.Equal(t, `<span name="snip1" class="test" myatt="myval" style="color=red;" tabindex=1></span>`, out.String())

	// simple body
	out.Reset()
	sA := ToHTML(`Hello World`)
	RenderChild(out, nil, sA)
	assert.Equal(t, "Hello World", out.String())

	// snippet with a tagname and default attributes
	out.Reset()
	s2 := &snip1{Test: 2}
	RenderChild(out, nil, s2)
	assert.Equal(t, `<div name="snip1" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`, out.String())

	// snippet with a tagname, default attributes and custom attributes
	// the tag builder of testsnippet2 overwrite custom attributes
	out.Reset()
	s2.Tag().SetTabIndex(1).AddStyle("color=red;").SetAttribute("a3", "").SetId("tst").AddClass("ts2c")
	RenderChild(out, nil, s2)
	assert.Equal(t, `<div id="tst" name="snip1" class="ts2a ts2b" a2 a3 style="display=test;" tabindex=2></div>`, out.String())
}
