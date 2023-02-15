package icecake

import (
	"reflect"

	"github.com/sunraylab/icecake/internal/helper"
)

var GData map[string]any = make(map[string]any, 0)

type Compounder interface {
	// State() any
	InnerHtmlTemplate() string
}

type CompoundBuilder interface {
	Mount()
}

var ComponentTypes map[string]reflect.Type

/*****************************************************************************/

func init() {
	ComponentTypes = make(map[string]reflect.Type, 0)
	RegisterComponentType("ic-ex1", reflect.TypeOf(CompEX1{}))
	RegisterComponentType("ic-ex2", reflect.TypeOf(CompEX2{}))
	RegisterComponentType("ic-ex3", reflect.TypeOf(CompEX3{}))
	RegisterComponentType("ic-ex4", reflect.TypeOf(CompEX4{}))
	RegisterComponentType("ic-ex5", reflect.TypeOf(CompEX5{}))
	RegisterComponentType("ic-ex6", reflect.TypeOf(CompEX6{}))
	// RegisterComponentType("ic-ex7", reflect.TypeOf(CompEX7{}))
}

func RegisterComponentType(key string, typ reflect.Type) {
	// TODO: check component type name convention with an hyphen aka "ic-XXXX"
	key = helper.Normalize(key)
	ComponentTypes[key] = typ
}

/*****************************************************************************/

type CompEX1 struct{}

func (c *CompEX1) InnerHtmlTemplate() string {
	return `composant1`
}

/*****************************************************************************/

type CompEX2 struct{}

func (c *CompEX2) InnerHtmlTemplate() string {
	return `composant2 <ic-ex1/>`
}

/*****************************************************************************/

type CompEX3 struct{}

func (c *CompEX3) InnerHtmlTemplate() string {
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

func (c *CompEX4) InnerHtmlTemplate() string {
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
		c.InnerHtmlTemplate()
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
func (c *CompEX5) InnerHtmlTemplate() string {
	return `composant5 title:{{.Me.Title}} name:{{.Global.name}}` // name:{{.Name}} count:{{.Count}}
}

type CompEX6 struct {
	Title string
}

func (c *CompEX6) Mount() {
	c.Title = "TITLE6"
}

// affiche une donnée globale de l'app
func (c *CompEX6) InnerHtmlTemplate() string {
	return `composant6 <ic-ex5/>` // name:{{.Name}} count:{{.Count}}
}

/*****************************************************************************/

// type Comp struct {
// 	Id                int
// 	InnerHtmlTemplate func() string
// 	Data              *(map[string]any)
// }

// func (c *Comp6) Mount() {
// 	c.Title = "TITLE"
// 	c.Comp.Data = &GData
// 	GData["count"] = 1
// 	c.InnerHtmlTemplate = func() string {
// 		return `composant6 {{.Comp.Id}} title:{{.Title}} count:{{index .Data "count"}}`
// 	}
// }

// var myComp6 Comp = Comp{
// 	InnerHtmlTemplate: func() string {
// 		return `composant6 {{.App.Title}}`
// 	},
// }
