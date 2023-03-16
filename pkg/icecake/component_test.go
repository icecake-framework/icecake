package ick

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcomponent struct {
	Text
}

func (*testcomponent) RegisterName() string { return "ick-test" }
func (*testcomponent) Container(_compid string) (_tagname string, _contclasses string, _contattrs string, _contstyle string) {
	return "DIV", "test", "hidden tabIndex=2", "display=test;"
}

func TestComposeFlat(t *testing.T) {

	out := new(bytes.Buffer)

	// simple template execution
	txt1 := new(Text)
	txt1.Content = `<strong>{{.Root.Name}}</strong>`
	data := struct{ Name string }{
		Name: "Bob",
	}
	err := ComposeHtmlE(out, txt1, data)
	assert.NoError(t, err)
	fmt.Println("A1)", out)

	// template execution fails
	out.Reset()
	txt1.Content = `Hello {{.name}}`
	err = ComposeHtmlE(out, txt1, data)
	assert.Error(t, err)
	fmt.Println("A2)", out)

	// force id, setup classes, attributes and style
	out.Reset()
	txt1.Content = `Hello World`
	txt1.SetupId().Force("helloworld0")
	txt1.SetupClasses().AddTokens("text")
	txt1.SetupAttributes().SetTabIndex(1)
	*txt1.SetupStyle() = "color=red;"
	ComposeHtmlE(out, txt1, nil)
	assert.Equal(t, "<SPAN class='text' id='helloworld0' style='color=red;' tabIndex=1>Hello World</SPAN>", out.String())
	fmt.Println("A3)", out)

	// classes combined setup classes and container class, priority to setup class
	out.Reset()
	txt2 := new(testcomponent)
	txt2.SetupClasses().AddTokens("text")
	txt2.SetupAttributes().SetTabIndex(1)
	*txt2.SetupStyle() = "color=red;"
	ComposeHtmlE(out, txt2, nil)
	assert.Equal(t, "<DIV class='text test' hidden id='texttest-0' style='display=test;color=red;' tabIndex=1></DIV>", out.String())
	fmt.Println("A4)", out)

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
			want: `<SPAN id='ick-text-0'><SPAN id='ick-text-1'><br/></SPAN></SPAN>`,
			err:  false},
		{name: "two components",
			in:   `Hello <ick-text Content="Bob"/>, Hello <ick-text Content="Alice"/>`,
			want: `<SPAN id='ick-text-2'>Hello <SPAN id='ick-text-3'>Bob</SPAN>, Hello <SPAN id='ick-text-4'>Alice</SPAN></SPAN>`,
			err:  false},
		{name: "setup attributes",
			in:   `Hello <ick-text style='color=red;' tabIndex=1 Content="Bob" class='text'/>`,
			want: `<SPAN id='ick-text-5'>Hello <SPAN class='text' id='ick-text-6' style='color=red;' tabIndex=1>Bob</SPAN></SPAN>`,
			err:  false},
		{name: "overloadind class",
			in:   `<ick-test class='text'/>`,
			want: `<SPAN id='ick-text-7'><DIV class='test text' hidden id='ick-test-0' style='display=test;' tabIndex=2></DIV></SPAN>`,
			err:  false},
		{name: "overloadind attribute",
			in:   `Hello <ick-test tabIndex=1/>`,
			want: `<SPAN id='ick-text-8'>Hello <DIV class='test' hidden id='ick-test-1' style='display=test;' tabIndex=1></DIV></SPAN>`,
			err:  false},
		{name: "overloadind style",
			in:   `Hello <ick-test style='color=red;'/>`,
			want: `<SPAN id='ick-text-9'>Hello <DIV class='test' hidden id='ick-test-2' style='color=red;display=test;' tabIndex=2></DIV></SPAN>`,
			err:  false},
		{name: "overloadind id",
			in:   `Hello <ick-test id='forcedid'/>`,
			want: `<SPAN id='ick-text-10'>Hello <DIV class='test' hidden id='forcedid' style='display=test;' tabIndex=2></DIV></SPAN>`,
			err:  false},
		{name: "recursive",
			in:   `<ick-test Content='<ick-test />' />`,
			want: ``,
			err:  true},
	}

	TheCmpReg.RegisterComponent(testcomponent{})

	// running tests
	cmp := new(Text)
	out := new(bytes.Buffer)
	for i, tst := range tstset {
		out.Reset()
		cmp.Content = tst.in
		err := ComposeHtmlE(out, cmp, nil)
		if tst.err {
			assert.Error(t, err, tst.name)
		} else {
			if assert.NoError(t, err, tst.name) {
				assert.Equal(t, tst.want, out.String(), tst.name)
				fmt.Printf("C%v) %s\n   --->%s\n", i, tst.name, out)
			} else {
				fmt.Println(err.Error())
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
			want: ``,
			err:  false},
		{name: "icktag2",
			in:   `<ick-test />`,
			want: ``,
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
			want: ``,
			err:  false},
		{name: "attrib2",
			in:   `<ick-test a />`,
			want: ``,
			err:  false},
		{name: "attrib3",
			in:   `<ick-test  c a1  b />`,
			want: ``,
			err:  false},
		{name: "attrib4",
			in:   `<ick-test ; />`,
			want: ``,
			err:  true},
		{name: "attrib5",
			in:   `<ick-test a😀b />`,
			want: ``,
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
			want: ``,
			err:  false},
		{name: "value 2",
			in:   `<ick-test abc=1 />`,
			want: ``,
			err:  false},
		{name: "quoted value 1",
			in:   `<ick-test a='x'/>`,
			want: ``,
			err:  false},
		{name: "quoted value 2",
			in:   `<ick-test a= 'x' />`,
			want: ``,
			err:  false},
		{name: "quoted value 3",
			in:   `<ick-test a='x' />`,
			want: ``,
			err:  false},
		{name: "value with spaces",
			in:   `<ick-test a=' x '/>`,
			want: ``,
			err:  false},
		{name: "double quote value",
			in:   `<ick-test a="y"/>`,
			want: ``,
			err:  true},
		{name: "mixed quotes value",
			in:   `<ick-test a="y'z;"/>`,
			want: ``,
			err:  false},
		{name: "string with quote value",
			in:   `<ick-test a=y'z/>`,
			want: ``,
			err:  false},
		{name: "html value",
			in:   `<ick-test a="<ok></>"/>`,
			want: ``,
			err:  false},
		{name: "text + embedded + text",
			in:   ` Hello <ick-test/> folks <ick-test/> ! `,
			want: ` Hello <ick-test/> folks <ick-test/> ! `,
			err:  false},
	}

	output := new(bytes.Buffer)
	for i, tst := range tstset {
		output.Reset()
		fmt.Printf("C%v) %s\n", i, tst.name)
		err := unfoldBody(output, []byte(tst.in), nil, 1)
		if tst.err {
			if assert.Error(t, err, tst.name) {
				fmt.Println("expected error:", err.Error())
			} else {
				fmt.Println("no error!")
			}
		} else {
			if assert.NoError(t, err, tst.name) {
				assert.Equal(t, tst.want, output.String(), tst.name)
				fmt.Printf("   --**%s**--\n", output)
			} else {
				fmt.Println("unexpected error:", err.Error())
			}
		}
	}

	fmt.Println()
	fmt.Println()
}
