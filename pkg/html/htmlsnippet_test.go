package html

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/icecake-framework/icecake/pkg/ickcore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposeBasics(t *testing.T) {
	out := new(bytes.Buffer)

	// empty snippet
	s00 := new(HTMLSnippet)
	err := Render(out, nil, s00)
	require.NoError(t, err)
	require.Empty(t, out)

	// unregistered snippet with a simple tagname and an empty body
	ickcore.ResetRegistry()
	s0 := Snippet("span")
	err = Render(out, nil, s0)
	require.NoError(t, err)
	assert.Equal(t, `<span></span>`, out.String())

	// force id, setup classes, attributes and style
	out.Reset()
	s0.Tag().SetTabIndex(1).AddStyle("color=red;").SetId("helloworld").AddClass("test")
	Render(out, nil, s0)
	assert.Equal(t, `<span id="helloworld" class="test" style="color=red;" tabindex=1></span>`, out.String())

	// The same but when registered, with forced id
	out.Reset()
	ickcore.AddRegistryEntry("ick-testsnippet0", &HTMLSnippet{})
	Render(out, nil, s0)
	assert.Equal(t, `<span id="helloworld" class="test" style="color=red;" tabindex=1></span>`, out.String())

	// The same but when registered, without id
	out.Reset()
	s0.Tag().SetId("")
	Render(out, nil, s0)
	assert.Equal(t, `<span class="test" style="color=red;" tabindex=1></span>`, out.String())

	// unregistered snippet with a simple body and no tagname
	out.Reset()
	s1 := ToHTML(`Hello World`)
	Render(out, nil, s1)
	assert.Equal(t, "Hello World", out.String())

	// snippet with a tagname and default attributes
	out.Reset()
	ickcore.AddRegistryEntry("ick-testsnippet2", &testsnippet2{})
	s2 := new(testsnippet2)
	Render(out, nil, s2)
	assert.Equal(t, `<div name="testsnippet2" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`, out.String())

	// snippet with a tagname, default attributes and custom attributes
	// the tag builder of testsnippet2 overwrite custom attributes
	out.Reset()
	s2 = new(testsnippet2)
	s2.Tag().SetTabIndex(1).AddStyle("color=red;").SetAttribute("a3", "").SetId("tst").AddClass("ts2c")
	Render(out, nil, s2)
	assert.Equal(t, `<div id="tst" name="testsnippet2" class="ts2a ts2b" a2 a3 style="display=test;" tabindex=2></div>`, out.String())

	// update custom attributes on an existing component
	out.Reset()
	s2.Tag().SetTabIndex(3).AddStyle("color=blue;").SetAttribute("a4", "").AddClass("ts2a ts2d").RemoveClass("ts2c").RemoveAttribute("a3")
	Render(out, nil, s2)
	assert.Equal(t, `<div id="tst" name="testsnippet2" class="ts2a ts2b" a2 a4 style="display=test;" tabindex=2></div>`, out.String())
}

func TestUnfoldBody1(t *testing.T) {

	ickcore.ResetRegistry()
	out := new(bytes.Buffer)

	// only the composer is registered, not the the data
	// so ick-testsnippet0 is an empty snippet : no tag and no content
	// the rendering should render nothing
	out.Reset()
	ickcore.AddRegistryEntry("ick-testsnippet0", &HTMLSnippet{})
	err := renderHTML(out, nil, *ToHTML("<ick-testsnippet0/>"))
	require.NoError(t, err)
	require.Equal(t, ``, out.String())

	out.Reset()
	ickcore.AddRegistryEntry("ick-testsnippet4", &testsnippet4{})
	err = renderHTML(out, nil, *ToHTML("<ick-testsnippet4/>"))
	require.NoError(t, err)
	require.Equal(t, `<div name="testsnippet4"></div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML("<ick-testsnippet4 test/>"))
	require.NoError(t, err)
	require.Equal(t, `<div name="testsnippet4" test></div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML("<ick-testsnippet4 Unmanaged='test'/>"))
	require.NoError(t, err)
	require.Equal(t, `<!--ick-testsnippet4: "Unmanaged" attribute: unmanaged type html.Unmanaged-->`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML("<ick-testsnippet4 IsOk=true/>"))
	require.NoError(t, err)
	require.Equal(t, `<div name="testsnippet4" class="ok"></div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML("<ick-testsnippet4 Text=success/>"))
	require.NoError(t, err)
	require.Equal(t, `<div name="testsnippet4">success</div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML(`<ick-testsnippet4 HTML="<strong>STRONG</strong>"/>`))
	require.NoError(t, err)
	require.Equal(t, `<div name="testsnippet4"><strong>STRONG</strong></div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML(`<ick-testsnippet4 I=777/>`))
	require.NoError(t, err)
	require.Equal(t, `<div name="testsnippet4">777</div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML(`<ick-testsnippet4 F=777.777/>`))
	require.NoError(t, err)
	require.Equal(t, `<div name="testsnippet4">777.777</div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML(`<ick-testsnippet4 D=5h30m40s/>`))
	require.NoError(t, err)
	require.Equal(t, `<div name="testsnippet4">6h30m40s</div>`, out.String())

	out.Reset()
	err = renderHTML(out, nil, *ToHTML(`<ick-testsnippet4 U="/icecake.dev"/>`))
	require.NoError(t, err)
	require.Equal(t, `<div name="testsnippet4"><a href='/icecake.dev'></a></div>`, out.String())
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
			want: `<span name="testsnippet0"></span>`,
			err:  false},
		{name: "icktag2",
			in:   `<ick-testsnippet0 />`,
			want: `<span name="testsnippet0"></span>`,
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
			want: `<span name="testsnippet0" a></span>`,
			err:  false},
		{name: "attrib2",
			in:   `<ick-testsnippet0 a />`,
			want: `<span name="testsnippet0" a></span>`,
			err:  false},
		{name: "attrib3",
			in:   `<ick-testsnippet0  c a1  b />`,
			want: `<span name="testsnippet0" a1 b c></span>`,
			err:  false},
		{name: "attrib4",
			in:   `<ick-testsnippet0 ; />`,
			want: ``,
			err:  true},
		{name: "attrib5",
			in:   `<ick-testsnippet0 a😀b />`,
			want: `<span name="testsnippet0" a😀b></span>`,
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
			want: `<span name="testsnippet0" a=1></span>`,
			err:  false},
		{name: "value 2",
			in:   `<ick-testsnippet0 abc=1 />`,
			want: `<span name="testsnippet0" abc=1></span>`,
			err:  false},
		{name: "quoted value 1",
			in:   `<ick-testsnippet0 a='x'/>`,
			want: `<span name="testsnippet0" a="x"></span>`,
			err:  false},
		{name: "quoted value 2",
			in:   `<ick-testsnippet0 a= 'x' />`,
			want: `<span name="testsnippet0" a="x"></span>`,
			err:  false},
		{name: "quoted value 3",
			in:   `<ick-testsnippet0 a='x' />`,
			want: `<span name="testsnippet0" a="x"></span>`,
			err:  false},

		{name: "value with spaces",
			in:   `<ick-testsnippet0 a=' x '/>`,
			want: `<span name="testsnippet0" a=" x "></span>`,
			err:  false},
		{name: "double quote value",
			in:   `<ick-testsnippet0 a="y"/>`,
			want: `<span name="testsnippet0" a="y"></span>`,
			err:  false},
		{name: "mixed quotes value",
			in:   `<ick-testsnippet0 a="y'z;"/>`,
			want: `<span name="testsnippet0" a="y'z;"></span>`,
			err:  false},
		{name: "string with quote value",
			in:   `<ick-testsnippet0 a=y'z/>`,
			want: `<span name="testsnippet0" a="y'z"></span>`,
			err:  false},
		{name: "html value",
			in:   `<ick-testsnippet0 a="<ok></>"/>`,
			want: `<span name="testsnippet0" a="<ok></>"></span>`,
			err:  false},
		{name: "text + embedding + text",
			in:   ` Hello <ick-testsnippet0/> folks <ick-testsnippet0/> ! `,
			want: ` Hello <span name="testsnippet0"></span> folks <span name="testsnippet0"></span> ! `,
			err:  false},

		{name: "simple embedding with attributes",
			in:   `<ick-testsnippet0 a b=1 c="x"/>`,
			want: `<span name="testsnippet0" a b=1 c="x"></span>`,
			err:  false},
		{name: "multi embedding with attributes",
			in:   `<ick-testsnippet0 a b=1 c="x"/><ick-testsnippet0 c="y" d/>`,
			want: `<span name="testsnippet0" a b=1 c="x"></span><span name="testsnippet0" c="y" d></span>`,
			err:  false},
		{name: "setup attributes",
			in:   `<ick-testsnippet2 class='text' d='test'/>`,
			want: `<div name="testsnippet2" class="ts2a ts2b" a2 d="test" style="display=test;" tabindex=2></div>`,
			err:  false},
		{name: "empty attributes",
			in:   `<ick-testsnippet0 empty1='' empty2="" bool/>`,
			want: `<span name="testsnippet0" bool empty1 empty2></span>`,
			err:  false},
	}

	// restet the component registrey for tests
	ickcore.ResetRegistry()
	ickcore.AddRegistryEntry("ick-testsnippet0", &testsnippet0{})
	ickcore.AddRegistryEntry("ick-testsnippet1", &testsnippet1{})
	ickcore.AddRegistryEntry("ick-testsnippet2", &testsnippet2{})

	output := new(bytes.Buffer)
	for i, tst := range tstset {
		output.Reset()
		err := renderHTML(output, nil, *ToHTML(tst.in))
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
			want: `Hello <span name="testsnippet0"></span>`,
			err:  false},
		{name: "simple body, no tagnam",
			in:   `Hello <ick-testsnippet1 HTML="Bob"/>`,
			want: `Hello Bob`,
			err:  false},
		{name: "attributes with html syntax",
			in:   `Hello <ick-testsnippet1 HTML="<br/>"/>`,
			want: `Hello <br/>`,
			err:  false},
		{name: "two components",
			in:   `Hello <ick-testsnippet1 HTML="Bob"/>, Hello <ick-testsnippet1 HTML="Alice"/>`,
			want: `Hello Bob, Hello Alice`,
			err:  false},

		{name: "setup attributes",
			in:   `Hello <ick-testsnippet0 style='color=red;' tabindex=1 a=0 class='text'/>`,
			want: `Hello <span name="testsnippet0" class="text" a=0 style="color=red;" tabindex=1></span>`,
			err:  false},
		{name: "overloadind class", // not possible with testsnippet2, forced by the TagBuilder
			in:   `<ick-testsnippet2 class='text'/>`,
			want: `<div name="testsnippet2" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`,
			err:  false},
		{name: "overloadind tabindex", // not possible with testsnippet2, forced by the TagBuilder
			in:   `<ick-testsnippet2 tabindex=1/>`,
			want: `<div name="testsnippet2" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`,
			err:  false},
		{name: "overloadind style", // not possible with testsnippet2, forced by the TagBuilder
			in:   `<ick-testsnippet2 style='color=red;'/>`,
			want: `<div name="testsnippet2" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`,
			err:  false},
		{name: "overloadind id",
			in:   `<ick-testsnippet2 id='forcedid'/>`,
			want: `<div id="forcedid" name="testsnippet2" class="ts2a ts2b" a2 style="display=test;" tabindex=2></div>`,
			err:  false},

		{name: "recursive",
			in:   `<ick-testinfinite/>`,
			want: ``,
			err:  true},
	}

	// restet the component registrey for tests
	ickcore.ResetRegistry()
	ickcore.AddRegistryEntry("ick-testsnippet0", &testsnippet0{})
	ickcore.AddRegistryEntry("ick-testsnippet1", &testsnippet1{})
	ickcore.AddRegistryEntry("ick-testsnippet2", &testsnippet2{})
	ickcore.AddRegistryEntry("ick-testsnippet4", &testsnippet4{})
	ickcore.AddRegistryEntry("ick-testinfinite", &testsnippetinfinite{})

	// running tests
	out := new(bytes.Buffer)
	for i, tst := range tstset {
		out.Reset()
		err := renderHTML(out, nil, *ToHTML(tst.in))
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

func TestSnippetId_1(t *testing.T) {

	// ickcore.ResetRegistry()
	// ickcore.AddRegistryEntry("ick-testsnippet0", &testsnippet0{})

	// out := new(bytes.Buffer)

	// // A> setup an ID upfront, before rendering
	// cmpA := &testsnippet0{}
	// // A.1> with no id
	// cmpA.Tag().SetId("idA1").SetBool("noid", true)
	// err := Render(out, nil, cmpA)
	// require.NoError(t, err)
	// require.Empty(t, cmpA.Id())
	// require.NotContains(t, out.String(), `id=`)

	// // A.2> with forced id
	// out.Reset()
	// cmpA.Tag().SetId("IdA2").SetBool("noid", false)
	// err = Render(out, nil, cmpA)
	// require.NoError(t, err)
	// require.Equal(t, "IdA2", cmpA.Id())
	// require.Contains(t, out.String(), `id="IdA2"`)

	// // B> setup an ID into the tagbuilder
	// cmpB := new(testsnippetid)
	// // B.1> withid noid
	// out.Reset()
	// cmpB.Tag().SetBool("noid", true)
	// err = Render(out, nil, cmpB)
	// require.NoError(t, err)
	// require.Empty(t, cmpB.Id())
	// require.NotContains(t, out.String(), `id=`)

	// // B.2> with forced id
	// out.Reset()
	// cmpB.Tag().SetBool("noid", false)
	// err = Render(out, nil, cmpB)
	// require.NoError(t, err)
	// require.Equal(t, "IdTemplate1", cmpB.Id())
	// require.Contains(t, out.String(), `id="IdTemplate1"`)

	// // C> setup an ID into the icktag
	// ickcore.AddRegistryEntry("ick-testsnippetid", &testsnippetid{})
	// // C.1> without parent
	// out.Reset()
	// err = renderHTML(out, nil, *ToHTML(`<ick-testsnippetid/>`))
	// require.NoError(t, err)
	// require.Equal(t, `<span id="IdTemplate1" name="testsnippetid"></span>`, out.String())

	// // C.2> with a parent
	// out.Reset()
	// err = renderHTML(out, cmpA, *ToHTML(`<ick-testsnippetid/>`))
	// require.NoError(t, err)
	// require.Equal(t, `<span id="IdA2.testsnippetid-0" name="testsnippetid"></span>`, out.String())

	// // C.2> with a parent and noid
	// out.Reset()
	// err = renderHTML(out, cmpA, *ToHTML(`<ick-testsnippetid noid/>`))
	// require.NoError(t, err)
	// require.Equal(t, `<span name="testsnippetid" noid></span>`, out.String())

}

func TestSnippetId_2(t *testing.T) {

	out := new(bytes.Buffer)

	s := Snippet("div", "noid").AddContent(ToHTML("<i>test</i>"))
	err := Render(out, nil, s)
	require.NoError(t, err)
	require.Equal(t, `<div noid><i>test</i></div>`, out.String())
}
