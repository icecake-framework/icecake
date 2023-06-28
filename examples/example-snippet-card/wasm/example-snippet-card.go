package main

import (
	"fmt"

	"github.com/icecake-framework/icecake/pkg/bulma"
	"github.com/icecake-framework/icecake/pkg/bulmaui"
	"github.com/icecake-framework/icecake/pkg/dom"
	"github.com/icecake-framework/icecake/pkg/html"
)

// This main package contains the web assembly source code for the icecake example.
// It's compiled into a '.wasm' file with the build_ex1 task
func main() {

	c := make(chan struct{})
	fmt.Println("Go/WASM loaded.")

	card1 := &bulmaui.Card{Card: bulma.Card{
		Title:   *html.ToHTML("Hello World"),
		Image:   bulma.NewImage("/icecake.jpg", bulma.IMG_SQUARE),
		Content: html.ToHTML("Nice cake"),
	}}
	card1.FooterItem = append(card1.FooterItem, *html.ToHTML("<a href='/'>home</a>"))
	card1.Tag().SetStyle("width: 350px;").AddClasses("mr-5")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, card1, nil)

	card2 := &bulmaui.Card{Card: bulma.Card{
		Content: html.ToHTML("Nice cake")}}
	card2.Tag().SetStyle("width: 128px;").AddClasses("mr-5")
	card2.Image = bulma.NewImage("/icecake.jpg", bulma.IMG_SQUARE)
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, card2, nil)

	card3 := &bulmaui.Card{Card: bulma.Card{
		Content: html.ToHTML("Very Nice cake")}}
	card3.Tag().AddClasses("mr-5")
	dom.Id("content").InsertSnippet(dom.INSERT_LAST_CHILD, card3, nil)

	// let's go
	fmt.Println("Go/WASM listening browser events")
	<-c
}
