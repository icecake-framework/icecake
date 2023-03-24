package ick0

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type testcomponent struct {
	Text
}

func (*testcomponent) RegisterName() string { return "ick-test" }
func (*testcomponent) Container(_compid string) (_tagname string, _contclasses string, _contattrs string, _contstyle string) {
	return "DIV", "test", "hidden tabIndex=2", "display=test;"
}

type testinfinite struct {
	Text
}

func (*testinfinite) RegisterName() string { return "ick-infinite" }
func (*testinfinite) Body() string {
	return "<ick-infinite/>"
}

func TestComposeFlat(t *testing.T) {

	// restet the component registrey for tests
	TheCmpReg = ComponentRegistry{}
	TheCmpReg.RegisterComponent(Text{})

	out := new(bytes.Buffer)

	// simple template execution
	txt1 := new(Text)
	txt1.Content = `<strong>{{.Root.Name}}</strong>`
	data := struct{ Name string }{
		Name: "Bob",
	}
	err := ComposeHtmlE(out, txt1, data)
	require.NoError(t, err)

	// template execution fails
	out.Reset()
	txt1.Content = `Hello {{.name}}`
	err = ComposeHtmlE(out, txt1, data)
	require.Error(t, err)

	// force id, setup classes, attributes and style
	out.Reset()
	txt1.Content = `Hello World`
	txt1.SetupId().Force("helloworld0")
	txt1.SetupClasses().AddTokens("text")
	txt1.SetupAttributes().SetTabIndex(1)
	*txt1.SetupStyle() = "color=red;"
	ComposeHtmlE(out, txt1, nil)
	require.Equal(t, "<SPAN class='text' id='helloworld0' style='color=red;' tabIndex=1>Hello World</SPAN>", out.String())

	// classes combined setup classes and container class, priority to setup class
	out.Reset()
	txt2 := new(testcomponent)
	txt2.SetupClasses().AddTokens("text")
	txt2.SetupAttributes().SetTabIndex(1)
	*txt2.SetupStyle() = "color=red;"
	ComposeHtmlE(out, txt2, nil)
	require.Equal(t, "<DIV class='test text' hidden id='testcomponent-0' style='color=red;display=test;' tabIndex=1></DIV>", out.String())
}

func TestComposeEmbedded(t *testing.T) {

	tstset := []struct {
		name string
		in   string
		want string
		err  bool
	}{
		{name: "simple",
			in:   `Hello <ick-text Content="Bob"/>`,
			want: `<SPAN id='ick-text-0'>Hello <SPAN id='ick-text-1'>Bob</SPAN></SPAN>`,
			err:  false},
		{name: "attributes with html syntax",
			in:   `<ick-text Content="<br/>"/>`,
			want: `<SPAN id='ick-text-2'><SPAN id='ick-text-3'><br/></SPAN></SPAN>`,
			err:  false},
		{name: "two components",
			in:   `Hello <ick-text Content="Bob"/>, Hello <ick-text Content="Alice"/>`,
			want: `<SPAN id='ick-text-4'>Hello <SPAN id='ick-text-5'>Bob</SPAN>, Hello <SPAN id='ick-text-6'>Alice</SPAN></SPAN>`,
			err:  false},

		{name: "setup attributes",
			in:   `Hello <ick-text style='color=red;' tabIndex=1 Content="Bob" class='text'/>`,
			want: `<SPAN id='ick-text-7'>Hello <SPAN class='text' id='ick-text-8' style='color=red;' tabIndex=1>Bob</SPAN></SPAN>`,
			err:  false},
		{name: "overloadind class",
			in:   `<ick-test class='text'/>`,
			want: `<SPAN id='ick-text-9'><DIV class='test text' hidden id='ick-test-0' style='display=test;' tabIndex=2></DIV></SPAN>`,
			err:  false},
		{name: "overloadind attribute",
			in:   `Hello <ick-test tabIndex=1/>`,
			want: `<SPAN id='ick-text-10'>Hello <DIV class='test' hidden id='ick-test-1' style='display=test;' tabIndex=1></DIV></SPAN>`,
			err:  false},
		{name: "overloadind style",
			in:   `Hello <ick-test style='color=red;'/>`,
			want: `<SPAN id='ick-text-11'>Hello <DIV class='test' hidden id='ick-test-2' style='color=red;display=test;' tabIndex=2></DIV></SPAN>`,
			err:  false},
		{name: "overloadind id",
			in:   `Hello <ick-test id='forcedid'/>`,
			want: `<SPAN id='ick-text-12'>Hello <DIV class='test' hidden id='forcedid' style='display=test;' tabIndex=2></DIV></SPAN>`,
			err:  false},
		{name: "recursive",
			in:   `<ick-infinite/>`,
			want: ``,
			err:  true},
	}

	// restet the component registrey for tests
	TheCmpReg = ComponentRegistry{}
	TheCmpReg.RegisterComponent(testcomponent{})
	TheCmpReg.RegisterComponent(Text{})
	TheCmpReg.RegisterComponent(testinfinite{})

	// running tests
	cmp := new(Text)
	out := new(bytes.Buffer)
	for i, tst := range tstset {
		out.Reset()
		cmp.Content = tst.in
		err := ComposeHtmlE(out, cmp, nil)
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
			in:   `<ick-test/>`,
			want: `<DIV class='test' hidden id='ick-test-0' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "icktag2",
			in:   `<ick-test />`,
			want: `<DIV class='test' hidden id='ick-test-1' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "icktag3",
			in:   `<ick-test>`,
			want: ``,
			err:  true},
		{name: "icktag4",
			in:   `<ick-/>`,
			want: ``,
			err:  true},
		{name: "attrib1",
			in:   `<ick-test a/>`,
			want: `<DIV a class='test' hidden id='ick-test-2' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "attrib2",
			in:   `<ick-test a />`,
			want: `<DIV a class='test' hidden id='ick-test-3' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "attrib3",
			in:   `<ick-test  c a1  b />`,
			want: `<DIV a1 b c class='test' hidden id='ick-test-4' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "attrib4",
			in:   `<ick-test ; />`,
			want: ``,
			err:  true},
		{name: "attrib5",
			in:   `<ick-test a😀b />`,
			want: `<DIV a😀b class='test' hidden id='ick-test-5' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "attrib6",
			in:   `<ick-test 1a />`,
			want: ``,
			err:  true},
		{name: "value without attribute 1",
			in:   `<ick-test=/>`,
			want: ``,
			err:  true},
		{name: "value without attribute 2",
			in:   `<ick-test = />`,
			want: ``,
			err:  true},
		{name: "missing value 1",
			in:   `<ick-test a=/>`,
			want: ``,
			err:  true},
		{name: "missing value 2",
			in:   `<ick-test a= />`,
			want: ``,
			err:  true},
		{name: "value 1",
			in:   `<ick-test a=1/>`,
			want: `<DIV a=1 class='test' hidden id='ick-test-6' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "value 2",
			in:   `<ick-test abc=1 />`,
			want: `<DIV abc=1 class='test' hidden id='ick-test-7' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "quoted value 1",
			in:   `<ick-test a='x'/>`,
			want: `<DIV a='x' class='test' hidden id='ick-test-8' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "quoted value 2",
			in:   `<ick-test a= 'x' />`,
			want: `<DIV a='x' class='test' hidden id='ick-test-9' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "quoted value 3",
			in:   `<ick-test a='x' />`,
			want: `<DIV a='x' class='test' hidden id='ick-test-10' style='display=test;' tabIndex=2></DIV>`,
			err:  false},

		{name: "value with spaces",
			in:   `<ick-test a=' x '/>`,
			want: `<DIV a=' x ' class='test' hidden id='ick-test-11' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "double quote value",
			in:   `<ick-test a="y"/>`,
			want: `<DIV a='y' class='test' hidden id='ick-test-12' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "mixed quotes value",
			in:   `<ick-test a="y'z;"/>`,
			want: `<DIV a="y'z;" class='test' hidden id='ick-test-13' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "string with quote value",
			in:   `<ick-test a=y'z/>`,
			want: `<DIV a="y'z" class='test' hidden id='ick-test-14' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "html value",
			in:   `<ick-test a="<ok></>"/>`,
			want: `<DIV a='<ok></>' class='test' hidden id='ick-test-15' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "text + embedding + text",
			in:   ` Hello <ick-test/> folks <ick-test/> ! `,
			want: ` Hello <DIV class='test' hidden id='ick-test-16' style='display=test;' tabIndex=2></DIV> folks <DIV class='test' hidden id='ick-test-17' style='display=test;' tabIndex=2></DIV> ! `,
			err:  false},
		{name: "simple embedding with attributes",
			in:   `<ick-test a b=1 c="x"/>`,
			want: `<DIV a b=1 c='x' class='test' hidden id='ick-test-18' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "multi embedding with attributes",
			in:   `<ick-test a b=1 c="x"/><ick-test c="y" d/>`,
			want: `<DIV a b=1 c='x' class='test' hidden id='ick-test-19' style='display=test;' tabIndex=2></DIV><DIV c='y' class='test' d hidden id='ick-test-20' style='display=test;' tabIndex=2></DIV>`,
			err:  false},
		{name: "setup attributes",
			in:   `<ick-text class='text' d='test'/>`,
			want: `<SPAN class='text' d='test' id='ick-text-0'></SPAN>`,
			err:  false},
	}

	// restet the component registrey for tests
	TheCmpReg = ComponentRegistry{}
	TheCmpReg.RegisterComponent(testcomponent{})
	TheCmpReg.RegisterComponent(Text{})

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
