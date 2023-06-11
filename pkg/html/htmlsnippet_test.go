package html

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/icecake-framework/icecake/pkg/registry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComposeBasics(t *testing.T) {
	out := new(bytes.Buffer)

	// empty snippet
	s00 := new(HTMLSnippet)
	_, err := WriteHTMLSnippet(out, s00, nil, true)
	require.NoError(t, err)
	require.Empty(t, out)

	// unregistered snippet with a simple tagname and an empty body
	s0 := new(testsnippet0)
	_, err = WriteHTMLSnippet(out, s0, nil, true)
	require.NoError(t, err)
	require.Equal(t, `<SPAN id="ick-2"></SPAN>`, out.String())

	// force id, setup classes, attributes and style
	out.Reset()
	s0.SetId("helloworld0").SetClasses("test").SetTabIndex(1).SetStyle("color=red;")
	WriteHTMLSnippet(out, s0, nil, true)
	require.Equal(t, `<SPAN id="helloworld0" tabIndex=1 class="test" style="color=red;"></SPAN>`, out.String())

	// The same but when registered
	out.Reset()
	registry.AddRegistryEntry("ick-test-snippet0", &testsnippet0{}, nil)
	WriteHTMLSnippet(out, s0, nil, true)
	require.Equal(t, `<SPAN id="helloworld0" tabIndex=1 class="test ick-test-snippet0" style="color=red;"></SPAN>`, out.String())

	// unregistered snippet with a simple body and no tagname
	out.Reset()
	s1 := new(testsnippet1)
	s1.Html = `Hello World`
	WriteHTMLSnippet(out, s1, nil, true)
	require.Equal(t, "Hello World", out.String())

	// snippet with a tagname and default attributes
	out.Reset()
	registry.AddRegistryEntry("ick-test-snippet2", &testsnippet2{}, nil)
	s2 := new(testsnippet2)
	WriteHTMLSnippet(out, s2, nil, true)
	require.Equal(t, `<DIV id="ick-test-snippet2-1" tabIndex=2 class="ts2a ts2b ick-test-snippet2" style="display=test;" a2></DIV>`, out.String())

	// snippet with a tagname, default attributes and custom attributes
	out.Reset()
	s2 = new(testsnippet2)
	s2.SetId("tst").SetClasses("ts2c").SetTabIndex(1).SetStyle("color=red;").SetAttribute("a3", "")
	WriteHTMLSnippet(out, s2, nil, true)
	require.Equal(t, `<DIV id="tst" tabIndex=1 class="ts2c ts2a ts2b ick-test-snippet2" style="color=red;" a2 a3></DIV>`, out.String())

	// update custom attributes on an existing component
	out.Reset()
	s2.SetClasses("ts2a ts2d").SetTabIndex(3).SetStyle("color=blue;").SetAttribute("a4", "")
	s2.RemoveClasses("ts2c")
	s2.RemoveAttribute("a3")
	WriteHTMLSnippet(out, s2, nil, true)
	require.Equal(t, `<DIV id="tst" tabIndex=3 class="ts2a ts2b ick-test-snippet2 ts2d" style="color=blue;" a2 a4></DIV>`, out.String())
}

func TestComposeDataState(t *testing.T) {
	out := new(bytes.Buffer)
	s3 := new(testsnippet3)
	ds := &DataState{
		App: "hello world",
	}
	_, err := WriteHTMLSnippet(out, s3, ds, true)
	require.NoError(t, err)
	require.Equal(t, `data.app=hello world`, out.String())
}

func TestUnfoldBody1(t *testing.T) {

	registry.ResetRegistry()
	out := new(bytes.Buffer)

	out.Reset()
	s := &testsnippet0{}
	s.TagName = "div"
	s.Body = "test"
	registry.AddRegistryEntry("ick-test-snippet0", s, nil)
	err := unfoldBody(nil, out, []byte("<ick-test-snippet0/>"), nil, 0)
	require.NoError(t, err)
	require.Equal(t, `<SPAN id="ick-test-snippet0-1" class="ick-test-snippet0"></SPAN>`, out.String())

	out.Reset()
	s = &testsnippet0{}
	unfoldBody(nil, out, []byte(`<ick-test-snippet0 empty1='' empty2="" bool/>`), nil, 0)
	require.Equal(t, `<SPAN id="ick-test-snippet0-2" class="ick-test-snippet0" bool empty1 empty2></SPAN>`, out.String())
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
		{name: "empty attributes",
			in:   `<ick-test-snippet0 empty1='' empty2="" bool/>`,
			want: `<SPAN id="ick-test-snippet0-22" class="ick-test-snippet0" bool empty1 empty2></SPAN>`,
			err:  false},
	}

	// restet the component registrey for tests
	registry.ResetRegistry()
	registry.AddRegistryEntry("ick-test-snippet0", &testsnippet0{}, nil)
	registry.AddRegistryEntry("ick-test-snippet1", &testsnippet1{}, nil)
	registry.AddRegistryEntry("ick-test-snippet2", &testsnippet2{}, nil)

	output := new(bytes.Buffer)
	for i, tst := range tstset {
		output.Reset()
		err := unfoldBody(nil, output, []byte(tst.in), nil, 0)
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
			in:   `Hello <ick-test-snippet1 Html="Bob"/>`,
			want: `Hello Bob`,
			err:  false},
		{name: "attributes with html syntax",
			in:   `Hello <ick-test-snippet1 Html="<br/>"/>`,
			want: `Hello <br/>`,
			err:  false},
		{name: "two components",
			in:   `Hello <ick-test-snippet1 Html="Bob"/>, Hello <ick-test-snippet1 Html="Alice"/>`,
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
	registry.ResetRegistry()
	registry.AddRegistryEntry("ick-test-snippet0", &testsnippet0{}, nil)
	registry.AddRegistryEntry("ick-test-snippet1", &testsnippet1{}, nil)
	registry.AddRegistryEntry("ick-test-snippet2", &testsnippet2{}, nil)
	registry.AddRegistryEntry("ick-test-infinite", &testsnippetinfinite{}, nil)

	// running tests
	cmp := new(testsnippet1)
	out := new(bytes.Buffer)
	for i, tst := range tstset {
		out.Reset()
		cmp.Html = String(tst.in)
		_, err := WriteHTMLSnippet(out, cmp, nil, true)
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

func TestUnfoldbody3(t *testing.T) {

	html := String(`<ick-button class="m-2 is-primary is-light" data-example=6 Title="Embedded" IsOutlined/>`)

	out := new(bytes.Buffer)
	embedded, err := UnfoldHTML(out, html, nil)
	require.NoError(t, err)
	require.NotNil(t, embedded)

	for _, sub := range embedded {
		_, ok := sub.(HTMLComposer)
		assert.True(t, ok)
	}
}

// TODO: TestSnippetId
func TestSnippetId(t *testing.T) {

	registry.ResetRegistry()
	registry.AddRegistryEntry("ick-test-snippet0", &testsnippet0{}, nil)

	out := new(bytes.Buffer)

	// A> setup an ID upfront, before rendering
	cmpA := new(testsnippet0)
	// A.1> withid = false
	cmpA.SetId("idA1")
	id, err := WriteHTMLSnippet(out, cmpA, nil, false)
	if err != nil {
		t.Error(err)
	}
	if id != "" {
		t.Errorf("id must be empty: %s", id)
	}

	// A.2> withid = true
	out.Reset()
	cmpA.SetId("IdA2")
	id, err = WriteHTMLSnippet(out, cmpA, nil, true)
	if err != nil {
		t.Error(err)
	}
	if id != "IdA2" {
		t.Errorf(`wrong snippet Id. Get %q, want:"IdA2"`, id)
	}

	// B> setup an ID into the template, during rendering
	cmpB := new(testsnippetid)
	// B.1> withid = false
	out.Reset()
	id, err = WriteHTMLSnippet(out, cmpB, nil, false)
	if err != nil {
		t.Error(err)
	}
	if id != "" {
		t.Errorf("id must be empty: %s", id)
	}

	// B.2> withid = true
	out.Reset()
	id, err = WriteHTMLSnippet(out, cmpB, nil, true)
	if err != nil {
		t.Error(err)
	}
	if id != "ick-1" {
		t.Errorf(`wrong snippet Id. Get %q, want:"ick-1"`, id)
	}
}
