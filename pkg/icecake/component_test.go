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
		out  string
		err  bool
	}{
		{name: "simple",
			in:  `Hello <ick-text Content="Bob"/>`,
			out: `<SPAN id='ick-text-0'>Hello <SPAN id='ick-text-1'>Bob</SPAN></SPAN>`,
			err: false},
		{name: "attributes with html syntax",
			in:  `<ick-text Content="<br/>"/>`,
			out: `<SPAN id='ick-text-0'><SPAN id='ick-text-1'><br/></SPAN></SPAN>`,
			err: false},
		{name: "two components",
			in:  `Hello <ick-text Content="Bob"/>, Hello <ick-text Content="Alice"/>`,
			out: `<SPAN id='ick-text-2'>Hello <SPAN id='ick-text-3'>Bob</SPAN>, Hello <SPAN id='ick-text-4'>Alice</SPAN></SPAN>`,
			err: false},
		{name: "setup attributes",
			in:  `Hello <ick-text style='color=red;' tabIndex=1 Content="Bob" class='text'/>`,
			out: `<SPAN id='ick-text-5'>Hello <SPAN class='text' id='ick-text-6' style='color=red;' tabIndex=1>Bob</SPAN></SPAN>`,
			err: false},
		{name: "overloadind class",
			in:  `<ick-test class='text'/>`,
			out: `<SPAN id='ick-text-7'><DIV class='test text' hidden id='ick-test-0' style='display=test;' tabIndex=2></DIV></SPAN>`,
			err: false},
		{name: "overloadind attribute",
			in:  `Hello <ick-test tabIndex=1/>`,
			out: `<SPAN id='ick-text-8'>Hello <DIV class='test' hidden id='ick-test-1' style='display=test;' tabIndex=1></DIV></SPAN>`,
			err: false},
		{name: "overloadind style",
			in:  `Hello <ick-test style='color=red;'/>`,
			out: `<SPAN id='ick-text-9'>Hello <DIV class='test' hidden id='ick-test-2' style='color=red;display=test;' tabIndex=2></DIV></SPAN>`,
			err: false},
		{name: "overloadind id",
			in:  `Hello <ick-test id='forcedid'/>`,
			out: `<SPAN id='ick-text-10'>Hello <DIV class='test' hidden id='forcedid' style='display=test;' tabIndex=2></DIV></SPAN>`,
			err: false},
		{name: "recursive",
			in:  `<ick-test Content='<ick-test />' />`,
			out: ``,
			err: true},
	}

	TheCmpReg.RegisterComponent(testcomponent{})

	// running tests
	cmp := new(Text)
	out := new(bytes.Buffer)
	for i, tst := range tstset {
		cmp.Content = tst.in
		err := ComposeHtmlE(out, cmp, nil)
		if tst.err {
			assert.Error(t, err, tst.name)
		} else {
			if assert.NoError(t, err, tst.name) {
				assert.Equal(t, tst.out, out.String(), tst.name)
				fmt.Printf("C%v) %s\n   --->%s\n", i, tst.name, out)
			} else {
				fmt.Println(err.Error())
			}
		}
	}

}
