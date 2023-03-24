package ick

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComposeBasics(t *testing.T) {
	out := new(bytes.Buffer)

	// empty snippet
	s00 := new(HtmlSnippet)
	err := RenderHtmlSnippet(out, s00, nil)
	require.NoError(t, err)
	require.Empty(t, out)

	// unregistered snippet with a simple tagname and an empty body
	s0 := new(testsnippet0)
	err = RenderHtmlSnippet(out, s0, nil)
	require.NoError(t, err)
	require.Equal(t, `<SPAN id="ick-1"></SPAN>`, out.String())

	// force id, setup classes, attributes and style
	out.Reset()
	s0.SetId("helloworld0").SetClasses("test").SetTabIndex(1).SetStyle("color=red;")
	RenderHtmlSnippet(out, s0, nil)
	require.Equal(t, `<SPAN id="helloworld0" tabIndex=1 class="test" style="color=red;"></SPAN>`, out.String())

	// The same but when registered
	out.Reset()
	Register("ick-testsnippet0", testsnippet0{})
	RenderHtmlSnippet(out, s0, nil)
	require.Equal(t, `<SPAN id="helloworld0" tabIndex=1 class="test ick-testsnippet0" style="color=red;"></SPAN>`, out.String())

	// unregistered snippet with a simple body and no tagname
	out.Reset()
	s1 := new(testsnippet1)
	s1.HTML = `Hello World`
	RenderHtmlSnippet(out, s1, nil)
	require.Equal(t, "Hello World", out.String())

	// snippet with a tagname and default attributes
	out.Reset()
	Register("ick-testsnippet2", testsnippet2{})
	s2 := new(testsnippet2)
	RenderHtmlSnippet(out, s2, nil)
	require.Equal(t, `<DIV id="ick-testsnippet2-1" tabIndex=2 class="ts2a ts2b ick-testsnippet2" style="display=test;" a2></DIV>`, out.String())

	// snippet with a tagname, default attributes and custom attributes
	out.Reset()
	s2 = new(testsnippet2)
	s2.SetId("tst").SetClasses("ts2c").SetTabIndex(1).SetStyle("color=red;").SetAttribute("a3", "", false)
	RenderHtmlSnippet(out, s2, nil)
	require.Equal(t, `<DIV id="tst" tabIndex=1 class="ts2c ts2a ts2b ick-testsnippet2" style="color=red;" a2 a3></DIV>`, out.String())

	// update custom attributes on an existing component
	out.Reset()
	s2.SetClasses("ts2a ts2d").SetTabIndex(3).SetStyle("color=blue;").SetAttribute("a4", "", false)
	s2.RemoveClass("ts2c")
	s2.RemoveAttribute("a3")
	RenderHtmlSnippet(out, s2, nil)
	require.Equal(t, `<DIV id="tst" tabIndex=3 class="ts2a ts2b ick-testsnippet2 ts2d" style="color=blue;" a2 a4></DIV>`, out.String())
}

func TestComposeDataState(t *testing.T) {
	out := new(bytes.Buffer)
	s3 := new(testsnippet3)
	ds := &DataState{
		App: "hello world",
	}
	err := RenderHtmlSnippet(out, s3, ds)
	require.NoError(t, err)
	require.Equal(t, `data.app=hello world`, out.String())
}

func TestUnfoldBody(t *testing.T) {
	tstset := []struct {
		name string
		in   string
		want string
		err  bool
	}{
		{name: "none",
			in:   ``,
			want: ``,
			err:  false},
		{name: "text",
			in:   `Hello`,
			want: `Hello`,
			err:  false},
		{name: "notrim",
			in:   `  Hell o  `,
			want: `  Hell o  `,
			err:  false},
		{name: "html",
			in:   `<strong>Hello</strong>`,
			want: `<strong>Hello</strong>`,
			err:  false},
		{name: "icktag1",
			in:   `<ick-test-snippet0/>`,
			want: `<SPAN id="ick-test-snippet0-1" class="ick-test-snippet0"></SPAN>`,
			err:  false},
		{name: "icktag2",
			in:   `<ick-test-snippet0 />`,
			want: `<SPAN id="ick-test-snippet0-2" class="ick-test-snippet0"></SPAN>`,
			err:  false},
		{name: "icktag3",
			in:   `<ick-test-snippet0 >`,
			want: ``,
			err:  true},
		{name: "icktag4",
			in:   `<ick-/>`,
			want: ``,
			err:  true},
		{name: "attrib1",
			in:   `<ick-test-snippet0 a/>`,
			want: `<SPAN id="ick-test-snippet0-3" class="ick-test-snippet0" a></SPAN>`,
			err:  false},
		{name: "attrib2",
			in:   `<ick-test-snippet0 a />`,
			want: `<SPAN id="ick-test-snippet0-4" class="ick-test-snippet0" a></SPAN>`,
			err:  false},
		{name: "attrib3",
			in:   `<ick-test-snippet0  c a1  b />`,
			want: `<SPAN id="ick-test-snippet0-5" class="ick-test-snippet0" a1 b c></SPAN>`,
			err:  false},
		{name: "attrib4",
			in:   `<ick-test-snippet0 ; />`,
			want: ``,
			err:  true},
		{name: "attrib5",
			in:   `<ick-test-snippet0 a😀b />`,
			want: `<SPAN id="ick-test-snippet0-6" class="ick-test-snippet0" a😀b></SPAN>`,
			err:  false},
		{name: "attrib6",
			in:   `<ick-test-snippet0 1a />`,
			want: ``,
			err:  true},
		{name: "value without attribute 1",
			in:   `<ick-test-snippet0=/>`,
			want: ``,
			err:  true},
		{name: "value without attribute 2",
			in:   `<ick-test-snippet0 = />`,
			want: ``,
			err:  true},
		{name: "missing value 1",
			in:   `<ick-test-snippet0 a=/>`,
			want: ``,
			err:  true},
		{name: "missing value 2",
			in:   `<ick-test-snippet0 a= />`,
			want: ``,
			err:  true},
		{name: "value 1",
			in:   `<ick-test-snippet0 a=1/>`,
			want: `<SPAN id="ick-test-snippet0-7" class="ick-test-snippet0" a=1></SPAN>`,
			err:  false},
		{name: "value 2",
			in:   `<ick-test-snippet0 abc=1 />`,
			want: `<SPAN id="ick-test-snippet0-8" class="ick-test-snippet0" abc=1></SPAN>`,
			err:  false},
		{name: "quoted value 1",
			in:   `<ick-test-snippet0 a='x'/>`,
			want: `<SPAN id="ick-test-snippet0-9" class="ick-test-snippet0" a="x"></SPAN>`,
			err:  false},
		{name: "quoted value 2",
			in:   `<ick-test-snippet0 a= 'x' />`,
			want: `<SPAN id="ick-test-snippet0-10" class="ick-test-snippet0" a="x"></SPAN>`,
			err:  false},
		{name: "quoted value 3",
			in:   `<ick-test-snippet0 a='x' />`,
			want: `<SPAN id="ick-test-snippet0-11" class="ick-test-snippet0" a="x"></SPAN>`,
			err:  false},

		{name: "value with spaces",
			in:   `<ick-test-snippet0 a=' x '/>`,
			want: `<SPAN id="ick-test-snippet0-12" class="ick-test-snippet0" a=" x "></SPAN>`,
			err:  false},
		{name: "double quote value",
			in:   `<ick-test-snippet0 a="y"/>`,
			want: `<SPAN id="ick-test-snippet0-13" class="ick-test-snippet0" a="y"></SPAN>`,
			err:  false},
		{name: "mixed quotes value",
			in:   `<ick-test-snippet0 a="y'z;"/>`,
			want: `<SPAN id="ick-test-snippet0-14" class="ick-test-snippet0" a="y'z;"></SPAN>`,
			err:  false},
		{name: "string with quote value",
			in:   `<ick-test-snippet0 a=y'z/>`,
			want: `<SPAN id="ick-test-snippet0-15" class="ick-test-snippet0" a="y'z"></SPAN>`,
			err:  false},
		{name: "html value",
			in:   `<ick-test-snippet0 a="<ok></>"/>`,
			want: `<SPAN id="ick-test-snippet0-16" class="ick-test-snippet0" a="<ok></>"></SPAN>`,
			err:  false},
		{name: "text + embedding + text",
			in:   ` Hello <ick-test-snippet0/> folks <ick-test-snippet0/> ! `,
			want: ` Hello <SPAN id="ick-test-snippet0-17" class="ick-test-snippet0"></SPAN> folks <SPAN id="ick-test-snippet0-18" class="ick-test-snippet0"></SPAN> ! `,
			err:  false},

		{name: "simple embedding with attributes",
			in:   `<ick-test-snippet0 a b=1 c="x"/>`,
			want: `<SPAN id="ick-test-snippet0-19" class="ick-test-snippet0" a b=1 c="x"></SPAN>`,
			err:  false},
		{name: "multi embedding with attributes",
			in:   `<ick-test-snippet0 a b=1 c="x"/><ick-test-snippet0 c="y" d/>`,
			want: `<SPAN id="ick-test-snippet0-20" class="ick-test-snippet0" a b=1 c="x"></SPAN><SPAN id="ick-test-snippet0-21" class="ick-test-snippet0" c="y" d></SPAN>`,
			err:  false},
		{name: "setup attributes",
			in:   `<ick-test-snippet2 class='text' d='test'/>`,
			want: `<DIV id="ick-test-snippet2-1" tabIndex=2 class="text ts2a ts2b ick-test-snippet2" style="display=test;" a2 d="test"></DIV>`,
			err:  false},
	}

	// restet the component registrey for tests
	TheRegistry = registry{}
	Register("ick-test-snippet0", testsnippet0{})
	Register("ick-test-snippet1", testsnippet1{})
	Register("ick-test-snippet2", testsnippet2{})

	output := new(bytes.Buffer)
	for i, tst := range tstset {
		output.Reset()
		err := unfoldBody(output, []byte(tst.in), nil, 0)
		if tst.err {
			if err == nil {
				fmt.Printf("T%v) %s\n", i, tst.name)
				fmt.Println("   no error but error expected!")
				t.FailNow()
			}
		} else {
			if err == nil {
				if tst.want != output.String() {
					fmt.Printf("T%v) %s\n", i, tst.name)
					fmt.Println("   want:", tst.want)
					fmt.Println("   out :", output.String())
					t.FailNow()
				}
			} else {
				fmt.Printf("T%v) %s\n", i, tst.name)
				fmt.Println("   unexpected error:", err.Error())
				t.FailNow()
			}
		}
	}
}

func TestComposeEmbedded(t *testing.T) {

	tstset := []struct {
		name string
		in   string
		want string
		err  bool
	}{
		{name: "empty span",
			in:   `Hello <ick-test-snippet0/>`,
			want: `Hello <SPAN id="ick-test-snippet0-1" class="ick-test-snippet0"></SPAN>`,
			err:  false},
		{name: "simple body, no tagnam",
			in:   `Hello <ick-test-snippet1 HTML="Bob"/>`,
			want: `Hello Bob`,
			err:  false},
		{name: "attributes with html syntax",
			in:   `Hello <ick-test-snippet1 HTML="<br/>"/>`,
			want: `Hello <br/>`,
			err:  false},
		{name: "two components",
			in:   `Hello <ick-test-snippet1 HTML="Bob"/>, Hello <ick-test-snippet1 HTML="Alice"/>`,
			want: `Hello Bob, Hello Alice`,
			err:  false},

		{name: "setup attributes",
			in:   `Hello <ick-test-snippet0 style='color=red;' tabIndex=1 a=0 class='text'/>`,
			want: `Hello <SPAN id="ick-test-snippet0-2" tabIndex=1 class="text ick-test-snippet0" style="color=red;" a=0></SPAN>`,
			err:  false},
		{name: "overloadind class",
			in:   `<ick-test-snippet2 class='text'/>`,
			want: `<DIV id="ick-test-snippet2-1" tabIndex=2 class="text ts2a ts2b ick-test-snippet2" style="display=test;" a2></DIV>`,
			err:  false},
		{name: "overloadind tabindex",
			in:   `<ick-test-snippet2 tabIndex=1/>`,
			want: `<DIV id="ick-test-snippet2-2" tabIndex=1 class="ts2a ts2b ick-test-snippet2" style="display=test;" a2></DIV>`,
			err:  false},
		{name: "overloadind style",
			in:   `<ick-test-snippet2 style='color=red;'/>`,
			want: `<DIV id="ick-test-snippet2-3" tabIndex=2 class="ts2a ts2b ick-test-snippet2" style="color=red;" a2></DIV>`,
			err:  false},
		{name: "overloadind id",
			in:   `<ick-test-snippet2 id='forcedid'/>`,
			want: `<DIV id="forcedid" tabIndex=2 class="ts2a ts2b ick-test-snippet2" style="display=test;" a2></DIV>`,
			err:  false},

		{name: "recursive",
			in:   `<ick-test-infinite/>`,
			want: ``,
			err:  true},
	}

	// restet the component registrey for tests
	TheRegistry = registry{}
	Register("ick-test-snippet0", testsnippet0{})
	Register("ick-test-snippet1", testsnippet1{})
	Register("ick-test-snippet2", testsnippet2{})
	Register("ick-test-infinite", testsnippetinfinite{})

	// running tests
	cmp := new(testsnippet1)
	out := new(bytes.Buffer)
	for i, tst := range tstset {
		out.Reset()
		cmp.HTML = HTML(tst.in)
		err := RenderHtmlSnippet(out, cmp, nil)
		if tst.err {
			if err == nil {
				fmt.Printf("T%v) %s\n", i, tst.name)
				fmt.Println("   no error but error expected!")
				t.FailNow()
			}
		} else {
			if err == nil {
				if tst.want != out.String() {
					fmt.Printf("T%v) %s\n", i, tst.name)
					fmt.Println("   want:", tst.want)
					fmt.Println("   out :", out.String())
					t.FailNow()
				}
			} else {
				fmt.Printf("T%v) %s\n", i, tst.name)
				fmt.Println("   unexpected error:", err.Error())
				t.FailNow()
			}
		}
	}
}