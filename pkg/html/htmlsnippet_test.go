package html

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/icecake-framework/icecake/pkg/registry"
	"github.com/stretchr/testify/require"
)

func TestComposeBasics(t *testing.T) {
	out := new(bytes.Buffer)

	// empty snippet
	s00 := new(HTMLSnippet)
	err := RenderSnippet(out, nil, s00)
	require.NoError(t, err)
	require.Empty(t, out)

	// unregistered snippet with a simple tagname and an empty body
	registry.ResetRegistry()
	s0 := NewSnippet("span")
	err = RenderSnippet(out, nil, s0)
	require.NoError(t, err)
	require.Equal(t, `<span id="orphan.htmlsnippet-1" name="ick-HTMLSnippet"></span>`, out.String())

	// force id, setup classes, attributes and style
	out.Reset()
	s0.Tag().SetTabIndex(1).SetStyle("color=red;").SetId("helloworld").AddClasses("test")
	RenderSnippet(out, nil, s0)
	require.Equal(t, `<span id="helloworld" name="ick-HTMLSnippet" class="test" style="color=red;" tabindex=1></span>`, out.String())

	// The same but when registered, with forced id
	out.Reset()
	registry.AddRegistryEntry("ick-testsnippet0", &HTMLSnippet{}, nil)
	RenderSnippet(out, nil, s0)
	require.Equal(t, `<span id="helloworld" name="ick-testsnippet0" class="test" style="color=red;" tabindex=1></span>`, out.String())

	// The same but when registered, without id
	out.Reset()
	s0.Tag().SetId("")
	RenderSnippet(out, nil, s0)
	require.Equal(t, `<span id="orphan.testsnippet0-2" name="ick-testsnippet0" class="test" style="color=red;" tabindex=1></span>`, out.String())

	// unregistered snippet with a simple body and no tagname
	out.Reset()
	s1 := new(testsnippet1)
	s1.Html = HTML(`Hello World`)
	RenderSnippet(out, nil, s1)
	require.Equal(t, "Hello World", out.String())

	// snippet with a tagname and default attributes
	out.Reset()
	registry.AddRegistryEntry("ick-testsnippet2", &testsnippet2{}, nil)
	s2 := new(testsnippet2)
	RenderSnippet(out, nil, s2)
	require.Equal(t, `<div id="orphan.testsnippet2-1" name="ick-testsnippet2" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`, out.String())

	// snippet with a tagname, default attributes and custom attributes
	// the tag builder of testsnippet2 overwrite custom attributes
	out.Reset()
	s2 = new(testsnippet2)
	s2.Tag().SetTabIndex(1).SetStyle("color=red;").SetAttribute("a3", "").SetId("tst").AddClasses("ts2c")
	RenderSnippet(out, nil, s2)
	require.Equal(t, `<div id="tst" name="ick-testsnippet2" class="ts2a ts2b" a2 a3 style="display=test;" tabindex=2></div>`, out.String())

	// update custom attributes on an existing component
	out.Reset()
	s2.Tag().SetTabIndex(3).SetStyle("color=blue;").SetAttribute("a4", "").AddClasses("ts2a ts2d").RemoveClasses("ts2c").RemoveAttribute("a3")
	RenderSnippet(out, nil, s2)
	require.Equal(t, `<div id="tst" name="ick-testsnippet2" class="ts2a ts2b" a2 a4 style="display=test;" tabindex=2></div>`, out.String())
}

func TestUnfoldBody1(t *testing.T) {

	registry.ResetRegistry()
	out := new(bytes.Buffer)

	// only the composer is registered, not the the data
	// so ick-testsnippet0 is an empty snippet : no tag and no content
	// the rendering should render nothing
	out.Reset()
	registry.AddRegistryEntry("ick-testsnippet0", &HTMLSnippet{}, nil)
	err := RenderHTML(out, nil, HTML("<ick-testsnippet0/>"))
	require.NoError(t, err)
	require.Equal(t, ``, out.String())

	out.Reset()
	registry.AddRegistryEntry("ick-testsnippet4", &testsnippet4{}, nil)
	err = RenderHTML(out, nil, HTML("<ick-testsnippet4/>"))
	require.NoError(t, err)
	require.Equal(t, `<div id="orphan.testsnippet4-1" name="ick-testsnippet4"></div>`, out.String())

	out.Reset()
	err = RenderHTML(out, nil, HTML("<ick-testsnippet4 test/>"))
	require.NoError(t, err)
	require.Equal(t, `<div id="orphan.testsnippet4-2" name="ick-testsnippet4" test></div>`, out.String())

	out.Reset()
	err = RenderHTML(out, nil, HTML("<ick-testsnippet4 content='test'/>"))
	require.NoError(t, err)
	require.Equal(t, `<!--ick-testsnippet4: "content" attribute: unmanaged type []html.HTMLComposer-->`, out.String())

	out.Reset()
	err = RenderHTML(out, nil, HTML("<ick-testsnippet4 IsOk=true/>"))
	require.NoError(t, err)
	require.Equal(t, `<div id="orphan.testsnippet4-3" name="ick-testsnippet4" class="ok"></div>`, out.String())

	out.Reset()
	err = RenderHTML(out, nil, HTML("<ick-testsnippet4 Text=success/>"))
	require.NoError(t, err)
	require.Equal(t, `<div id="orphan.testsnippet4-4" name="ick-testsnippet4">success</div>`, out.String())

	out.Reset()
	err = RenderHTML(out, nil, HTML(`<ick-testsnippet4 HTML="<strong>STRONG</strong>"/>`))
	require.NoError(t, err)
	require.Equal(t, `<div id="orphan.testsnippet4-5" name="ick-testsnippet4"><strong>STRONG</strong></div>`, out.String())

	out.Reset()
	err = RenderHTML(out, nil, HTML(`<ick-testsnippet4 I=777/>`))
	require.NoError(t, err)
	require.Equal(t, `<div id="orphan.testsnippet4-6" name="ick-testsnippet4">777</div>`, out.String())

	out.Reset()
	err = RenderHTML(out, nil, HTML(`<ick-testsnippet4 F=777.777/>`))
	require.NoError(t, err)
	require.Equal(t, `<div id="orphan.testsnippet4-7" name="ick-testsnippet4">777.777</div>`, out.String())

	out.Reset()
	err = RenderHTML(out, nil, HTML(`<ick-testsnippet4 D=5h30m40s/>`))
	require.NoError(t, err)
	require.Equal(t, `<div id="orphan.testsnippet4-8" name="ick-testsnippet4">6h30m40s</div>`, out.String())

	out.Reset()
	err = RenderHTML(out, nil, HTML(`<ick-testsnippet4 U="/icecake.dev"/>`))
	require.NoError(t, err)
	require.Equal(t, `<div id="orphan.testsnippet4-9" name="ick-testsnippet4"><a href='/icecake.dev'></a></div>`, out.String())
}

func TestUnfoldBody2(t *testing.T) {
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
			in:   `<ick-testsnippet0/>`,
			want: `<span id="orphan.testsnippet0-1" name="ick-testsnippet0"></span>`,
			err:  false},
		{name: "icktag2",
			in:   `<ick-testsnippet0 />`,
			want: `<span id="orphan.testsnippet0-2" name="ick-testsnippet0"></span>`,
			err:  false},
		{name: "icktag3",
			in:   `<ick-testsnippet0 >`,
			want: ``,
			err:  true},
		{name: "icktag4",
			in:   `<ick-/>`,
			want: ``,
			err:  true},
		{name: "attrib1",
			in:   `<ick-testsnippet0 a/>`,
			want: `<span id="orphan.testsnippet0-3" name="ick-testsnippet0" a></span>`,
			err:  false},
		{name: "attrib2",
			in:   `<ick-testsnippet0 a />`,
			want: `<span id="orphan.testsnippet0-4" name="ick-testsnippet0" a></span>`,
			err:  false},
		{name: "attrib3",
			in:   `<ick-testsnippet0  c a1  b />`,
			want: `<span id="orphan.testsnippet0-5" name="ick-testsnippet0" a1 b c></span>`,
			err:  false},
		{name: "attrib4",
			in:   `<ick-testsnippet0 ; />`,
			want: ``,
			err:  true},
		{name: "attrib5",
			in:   `<ick-testsnippet0 a😀b />`,
			want: `<span id="orphan.testsnippet0-6" name="ick-testsnippet0" a😀b></span>`,
			err:  false},
		{name: "attrib6",
			in:   `<ick-testsnippet0 1a />`,
			want: ``,
			err:  true},
		{name: "value without attribute 1",
			in:   `<ick-testsnippet0=/>`,
			want: ``,
			err:  true},
		{name: "value without attribute 2",
			in:   `<ick-testsnippet0 = />`,
			want: ``,
			err:  true},
		{name: "missing value 1",
			in:   `<ick-testsnippet0 a=/>`,
			want: ``,
			err:  true},
		{name: "missing value 2",
			in:   `<ick-testsnippet0 a= />`,
			want: ``,
			err:  true},
		{name: "value 1",
			in:   `<ick-testsnippet0 a=1/>`,
			want: `<span id="orphan.testsnippet0-7" name="ick-testsnippet0" a=1></span>`,
			err:  false},
		{name: "value 2",
			in:   `<ick-testsnippet0 abc=1 />`,
			want: `<span id="orphan.testsnippet0-8" name="ick-testsnippet0" abc=1></span>`,
			err:  false},
		{name: "quoted value 1",
			in:   `<ick-testsnippet0 a='x'/>`,
			want: `<span id="orphan.testsnippet0-9" name="ick-testsnippet0" a="x"></span>`,
			err:  false},
		{name: "quoted value 2",
			in:   `<ick-testsnippet0 a= 'x' />`,
			want: `<span id="orphan.testsnippet0-10" name="ick-testsnippet0" a="x"></span>`,
			err:  false},
		{name: "quoted value 3",
			in:   `<ick-testsnippet0 a='x' />`,
			want: `<span id="orphan.testsnippet0-11" name="ick-testsnippet0" a="x"></span>`,
			err:  false},

		{name: "value with spaces",
			in:   `<ick-testsnippet0 a=' x '/>`,
			want: `<span id="orphan.testsnippet0-12" name="ick-testsnippet0" a=" x "></span>`,
			err:  false},
		{name: "double quote value",
			in:   `<ick-testsnippet0 a="y"/>`,
			want: `<span id="orphan.testsnippet0-13" name="ick-testsnippet0" a="y"></span>`,
			err:  false},
		{name: "mixed quotes value",
			in:   `<ick-testsnippet0 a="y'z;"/>`,
			want: `<span id="orphan.testsnippet0-14" name="ick-testsnippet0" a="y'z;"></span>`,
			err:  false},
		{name: "string with quote value",
			in:   `<ick-testsnippet0 a=y'z/>`,
			want: `<span id="orphan.testsnippet0-15" name="ick-testsnippet0" a="y'z"></span>`,
			err:  false},
		{name: "html value",
			in:   `<ick-testsnippet0 a="<ok></>"/>`,
			want: `<span id="orphan.testsnippet0-16" name="ick-testsnippet0" a="<ok></>"></span>`,
			err:  false},
		{name: "text + embedding + text",
			in:   ` Hello <ick-testsnippet0/> folks <ick-testsnippet0/> ! `,
			want: ` Hello <span id="orphan.testsnippet0-17" name="ick-testsnippet0"></span> folks <span id="orphan.testsnippet0-18" name="ick-testsnippet0"></span> ! `,
			err:  false},

		{name: "simple embedding with attributes",
			in:   `<ick-testsnippet0 a b=1 c="x"/>`,
			want: `<span id="orphan.testsnippet0-19" name="ick-testsnippet0" a b=1 c="x"></span>`,
			err:  false},
		{name: "multi embedding with attributes",
			in:   `<ick-testsnippet0 a b=1 c="x"/><ick-testsnippet0 c="y" d/>`,
			want: `<span id="orphan.testsnippet0-20" name="ick-testsnippet0" a b=1 c="x"></span><span id="orphan.testsnippet0-21" name="ick-testsnippet0" c="y" d></span>`,
			err:  false},
		{name: "setup attributes",
			in:   `<ick-testsnippet2 class='text' d='test'/>`,
			want: `<div id="orphan.testsnippet2-1" name="ick-testsnippet2" class="ts2a ts2b" a2 d="test" style="display=test;" tabindex=2></div>`,
			err:  false},
		{name: "empty attributes",
			in:   `<ick-testsnippet0 empty1='' empty2="" bool/>`,
			want: `<span id="orphan.testsnippet0-22" name="ick-testsnippet0" bool empty1 empty2></span>`,
			err:  false},
	}

	// restet the component registrey for tests
	registry.ResetRegistry()
	registry.AddRegistryEntry("ick-testsnippet0", &testsnippet0{}, nil)
	registry.AddRegistryEntry("ick-testsnippet1", &testsnippet1{}, nil)
	registry.AddRegistryEntry("ick-testsnippet2", &testsnippet2{}, nil)

	output := new(bytes.Buffer)
	for i, tst := range tstset {
		output.Reset()
		err := RenderHTML(output, nil, HTML(tst.in))
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
			in:   `Hello <ick-testsnippet0/>`,
			want: `Hello <span id="orphan.testsnippet0-1" name="ick-testsnippet0"></span>`,
			err:  false},
		{name: "simple body, no tagnam",
			in:   `Hello <ick-testsnippet1 Html="Bob"/>`,
			want: `Hello Bob`,
			err:  false},
		{name: "attributes with html syntax",
			in:   `Hello <ick-testsnippet1 Html="<br/>"/>`,
			want: `Hello <br/>`,
			err:  false},
		{name: "two components",
			in:   `Hello <ick-testsnippet1 Html="Bob"/>, Hello <ick-testsnippet1 Html="Alice"/>`,
			want: `Hello Bob, Hello Alice`,
			err:  false},

		{name: "setup attributes",
			in:   `Hello <ick-testsnippet0 style='color=red;' tabindex=1 a=0 class='text'/>`,
			want: `Hello <span id="orphan.testsnippet0-2" name="ick-testsnippet0" class="text" a=0 style="color=red;" tabindex=1></span>`,
			err:  false},
		{name: "overloadind class", // not possible with testsnippet2, forced by the TagBuilder
			in:   `<ick-testsnippet2 class='text'/>`,
			want: `<div id="orphan.testsnippet2-1" name="ick-testsnippet2" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`,
			err:  false},
		{name: "overloadind tabindex", // not possible with testsnippet2, forced by the TagBuilder
			in:   `<ick-testsnippet2 tabindex=1/>`,
			want: `<div id="orphan.testsnippet2-2" name="ick-testsnippet2" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`,
			err:  false},
		{name: "overloadind style", // not possible with testsnippet2, forced by the TagBuilder
			in:   `<ick-testsnippet2 style='color=red;'/>`,
			want: `<div id="orphan.testsnippet2-3" name="ick-testsnippet2" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`,
			err:  false},
		{name: "overloadind id",
			in:   `<ick-testsnippet2 id='forcedid'/>`,
			want: `<div id="forcedid" name="ick-testsnippet2" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`,
			err:  false},

		{name: "recursive",
			in:   `<ick-testinfinite/>`,
			want: ``,
			err:  true},
	}

	// restet the component registrey for tests
	registry.ResetRegistry()
	registry.AddRegistryEntry("ick-testsnippet0", &testsnippet0{}, nil)
	registry.AddRegistryEntry("ick-testsnippet1", &testsnippet1{}, nil)
	registry.AddRegistryEntry("ick-testsnippet2", &testsnippet2{}, nil)
	registry.AddRegistryEntry("ick-testinfinite", &testsnippetinfinite{}, nil)

	// running tests
	out := new(bytes.Buffer)
	for i, tst := range tstset {
		out.Reset()
		err := RenderHTML(out, nil, HTML(tst.in))
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

func TestSnippetId(t *testing.T) {

	registry.ResetRegistry()
	registry.AddRegistryEntry("ick-testsnippet0", &testsnippet0{}, nil)

	out := new(bytes.Buffer)

	// A> setup an ID upfront, before rendering
	cmpA := &testsnippet0{}
	// A.1> with no id
	cmpA.Tag().SetId("idA1").SetBool("noid", true)
	err := RenderSnippet(out, nil, cmpA)
	require.NoError(t, err)
	require.Empty(t, cmpA.Id())
	require.NotContains(t, out.String(), `id=`)

	// A.2> with forced id
	out.Reset()
	cmpA.Tag().SetId("IdA2").SetBool("noid", false)
	err = RenderSnippet(out, nil, cmpA)
	require.NoError(t, err)
	require.Equal(t, "IdA2", cmpA.Id())
	require.Contains(t, out.String(), `id="IdA2"`)

	// B> setup an ID into the tagbuilder
	cmpB := new(testsnippetid)
	// B.1> withid noid
	out.Reset()
	cmpB.Tag().SetBool("noid", true)
	err = RenderSnippet(out, nil, cmpB)
	require.NoError(t, err)
	require.Empty(t, cmpB.Id())
	require.NotContains(t, out.String(), `id=`)

	// B.2> with forced id
	out.Reset()
	cmpB.Tag().SetBool("noid", false)
	err = RenderSnippet(out, nil, cmpB)
	require.NoError(t, err)
	require.Equal(t, "IdTemplate1", cmpB.Id())
	require.Contains(t, out.String(), `id="IdTemplate1"`)

	// C> setup an ID into the icktag
	registry.AddRegistryEntry("ick-testsnippetid", &testsnippetid{}, nil)
	// C.1> without parent
	out.Reset()
	err = RenderHTML(out, nil, HTML(`<ick-testsnippetid/>`))
	require.NoError(t, err)
	require.Equal(t, `<span id="IdTemplate1" name="ick-testsnippetid"></span>`, out.String())

	// C.2> with a parent
	out.Reset()
	err = RenderHTML(out, cmpA, HTML(`<ick-testsnippetid/>`))
	require.NoError(t, err)
	require.Equal(t, `<span id="IdA2.testsnippetid-0" name="ick-testsnippetid"></span>`, out.String())

	// C.2> with a parent and noid
	out.Reset()
	err = RenderHTML(out, cmpA, HTML(`<ick-testsnippetid noid/>`))
	require.NoError(t, err)
	require.Equal(t, `<span name="ick-testsnippetid" noid></span>`, out.String())

}

func TestHTMLSnippetContent(t *testing.T) {

	out := new(bytes.Buffer)

	s := NewSnippet("div", "noid").StackContent(NewHTML("<i>test</i>"))
	err := RenderSnippet(out, nil, s)
	require.NoError(t, err)
	require.Equal(t, `<div name="ick-HTMLSnippet" noid><i>test</i></div>`, out.String())
}

func TestComposeDataState(t *testing.T) {
	// out := new(bytes.Buffer)
	// s3 := new(testsnippet3)
	// ds := &DataState{
	// 	App: "hello world",
	// }
	// s3.SetDataState(ds)
	// err := RenderSnippet(out, nil, s3)
	// require.NoError(t, err)
	// require.Equal(t, `data.app=hello world`, out.String())
}
