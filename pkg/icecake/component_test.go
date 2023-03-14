package ick

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type texttest struct {
	Text
}

func (*texttest) Container(_compid string) (_tagname string, _contclasses string, _contattrs string, _contstyle string) {
	return "DIV", "text2", "hidden tabIndex=2", "width=10px;"
}

func TestComposeL0(t *testing.T) {

	out := new(bytes.Buffer)

	// simple template execution
	txt1 := new(Text)
	txt1.Content = `<strong>{{.Root.Name}}</strong>`
	data := struct{ Name string }{
		Name: "Bob",
	}
	err := ComposeHtmlE(out, txt1, data)
	assert.NoError(t, err)
	fmt.Println("1)", out)

	// template execution fails
	out.Reset()
	txt1.Content = `Hello {{.name}}`
	err = ComposeHtmlE(out, txt1, data)
	assert.Error(t, err)
	fmt.Println("2)", out)

	// force id, setup classes, attributes and style
	out.Reset()
	txt1.Content = `Hello World`
	txt1.SetupId().Force("helloworld0")
	txt1.SetupClasses().AddTokens("text")
	txt1.SetupAttributes().SetTabIndex(1)
	*txt1.SetupStyle() = "color=red;"
	ComposeHtmlE(out, txt1, nil)
	assert.Equal(t, "<SPAN class='text' id='helloworld0' style='color=red;' tabIndex=1>Hello World</SPAN>", out.String())
	fmt.Println("3)", out)

	// classes combined setup classes and container class, priority to setup class
	out.Reset()
	txt2 := new(texttest)
	txt2.SetupClasses().AddTokens("text")
	txt2.SetupAttributes().SetTabIndex(1)
	*txt2.SetupStyle() = "color=red;"
	ComposeHtmlE(out, txt2, nil)
	assert.Equal(t, "<DIV class='text text2' hidden id='texttest-3' style='width=10px;color=red;' tabIndex=1></DIV>", out.String())
	fmt.Println("4)", out)

}

func TestComposeEmbedded(t *testing.T) {

	out := new(bytes.Buffer)

	txt2 := new(Text)
	txt2.Content = `Hello <ick-text content="Bob"/>, Hello <ick-text content="Alice"/>`
	err := ComposeHtmlE(out, txt2, nil)
	assert.NoError(t, err)
	assert.Equal(t, `Hello <span id="ick-text-1">Bob</span>, Hello <span id="ick-text-2">alice</span>`, out.String())
	fmt.Println("1)", out)

}
