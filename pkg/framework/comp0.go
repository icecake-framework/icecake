package icecake

import (
	"reflect"

	"github.com/sunraylab/icecake/internal/helper"
)

type HtmlCompounder interface {
	Template() string
}

type HtmlListener interface {
	AddListeners()
}

// type StyleCompounder interface {
// 	Style() string
// }

var GData map[string]any = make(map[string]any, 0)

type Compounder interface {
	// State() any
	InnerTemplate() string
}

type CompoundBuilder interface {
	Mount()
}

var ComponentTypes map[string]reflect.Type

/*****************************************************************************/

func init() {
	ComponentTypes = make(map[string]reflect.Type, 0)
	RegisterComponentType("ick-ex1", reflect.TypeOf(CompEX1{}))
	RegisterComponentType("ick-ex2", reflect.TypeOf(CompEX2{}))
	RegisterComponentType("ick-ex3", reflect.TypeOf(CompEX3{}))
	RegisterComponentType("ick-ex4", reflect.TypeOf(CompEX4{}))
	RegisterComponentType("ick-ex5", reflect.TypeOf(CompEX5{}))
	RegisterComponentType("ick-ex6", reflect.TypeOf(CompEX6{}))
	// RegisterComponentType("ic-ex7", reflect.TypeOf(CompEX7{}))
}

func RegisterComponentType(key string, typ reflect.Type) {
	// TODO: check component type name convention with an hyphen aka "ic-XXXX"
	key = helper.Normalize(key)
	ComponentTypes[key] = typ
}

/*****************************************************************************/

type CompEX1 struct{}

func (c *CompEX1) InnerTemplate() string {
	return `composant1`
}

/*****************************************************************************/

type CompEX2 struct{}

func (c *CompEX2) InnerTemplate() string {
	return `composant2 <ick-ex1/>`
}

/*****************************************************************************/

type CompEX3 struct{}

func (c *CompEX3) InnerTemplate() string {
	return `composant3 {{.}}`
}

func (c *CompEX3) String() string {
	return `***`
}

/*****************************************************************************/

type CompEX4 struct {
	Count int
}

func (c *CompEX4) Mount() {

	// reference l'instance du composant, pour pouvoir appeler unmount

	// init local data

	// ajoute des events
	// le listener sera ajouté à la fin
	// AddEvent()
}

func (c *CompEX4) InnerTemplate() string {
	return `composant4 <button>count is {{.Count}}</button>`
}

/*****************************************************************************/

type StateSaver interface {
	Save()
}

var Subscriptions map[any]Compounder = make(map[any]Compounder, 0)

func Subscribe(c Compounder, pval any) {
	Subscriptions[pval] = c
}

func Signal(pval any) {
	for _, c := range Subscriptions {
		c.InnerTemplate()
	}
}

type CompEX5 struct {
	Title string
}

func (c *CompEX5) Mount() {
	c.Title = "TITLE5"
	Subscribe(c, GData["name"])
}

// affiche une donnée globale de l'app
func (c *CompEX5) InnerTemplate() string {
	return `composant5 title:{{.Me.Title}} name:{{.Global.name}}` // name:{{.Name}} count:{{.Count}}
}

type CompEX6 struct {
	Title string
}

func (c *CompEX6) Mount() {
	c.Title = "TITLE6"
}

// affiche une donnée globale de l'app
func (c *CompEX6) InnerTemplate() string {
	return `composant6 <ick-ex5/>` // name:{{.Name}} count:{{.Count}}
}

/*****************************************************************************/

// type Comp struct {
// 	Id                int
// 	InnerTemplate func() string
// 	Data              *(map[string]any)
// }

// func (c *Comp6) Mount() {
// 	c.Title = "TITLE"
// 	c.Comp.Data = &GData
// 	GData["count"] = 1
// 	c.InnerTemplate = func() string {
// 		return `composant6 {{.Comp.Id}} title:{{.Title}} count:{{index .Data "count"}}`
// 	}
// }

// var myComp6 Comp = Comp{
// 	InnerTemplate: func() string {
// 		return `composant6 {{.App.Title}}`
// 	},
// }
