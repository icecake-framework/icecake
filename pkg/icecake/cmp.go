package ick

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/errors"
)

type Component struct {
	Element

	// ID string
	// Attributes
	// Classes TokenList
}

func (c *Component) Envelope() (_tagname string, _classTemplate string) { return "span", "" }

func (c *Component) Template() (_html string) { return "" }

func (c *Component) AddListeners() {}

type HtmlContainer interface {
	Wrap(JSValueProvider)
	Container() (_tagname string, _classes string, _attrs string)
}

type HtmlTemplater interface {
	Template() (_html string)
}
type HtmlListener interface {
	Wrap(JSValueProvider)
	AddListeners()
}

type StyleComposer interface {
	Style() string
}

type Composer interface {
	HtmlContainer
	HtmlTemplater
	HtmlListener
}

/*****************************************************************************/

var gComponents int

func GetNextComponentId(_prefix string) (_id string) {
	idx := gComponents + 1
	gComponents++

	_id = "c" + strconv.Itoa(idx)
	if _prefix != "" {
		_id = _prefix + "-" + _id
	}
	return _id
}

var GComponentRegistry map[string]reflect.Type

/*****************************************************************************/

func init() {
	GComponentRegistry = make(map[string]reflect.Type, 0)
}

func RegisterComponentType(key string, cmp any) error {
	key = helper.Normalize(key)
	if !strings.HasPrefix(key, "ick-") {
		return errors.ConsoleErrorf("RegisterComponentType faild: key %q does not match allowed pattern\n", key)
	}
	name := strings.TrimPrefix(key, "ick-")
	if len(name) == 0 {
		return errors.ConsoleErrorf("RegisterComponentType faild: invalid key name %q\n", key)
	}

	typ := reflect.TypeOf(cmp)
	if typ.Kind() == reflect.Pointer {
		return errors.ConsoleErrorf("RegisterComponentType faild: must register a component not a pointer to a component %q\n", typ.String())
	}

	if _, found := typ.FieldByName("Component"); !found {
		return errors.ConsoleErrorf("RegisterComponentType faild: your component must embed the ick.Component value\n")
	}

	if _, found := GComponentRegistry[key]; found {
		return errors.ConsoleErrorf("RegisterComponentType faild: %q already registered\n", key)
	}

	GComponentRegistry[key] = typ
	return errors.ConsoleLogf("RegisterComponentType: %s %q\n", key, typ.String())
}

func LookupComponent(typ reflect.Type) string {
	st := strings.TrimLeft(typ.String(), "*")
	for k, v := range GComponentRegistry {
		sv := strings.TrimLeft(v.String(), "*")
		if sv == st {
			return k
		}
	}
	return ""
}

/*****************************************************************************/

func CreateComponentElement(_composer HtmlContainer) (_elem *Element, _err error) {
	// create the HTML element
	tagname, classes, strattrs := _composer.Container()
	tagname = helper.Normalize(tagname)
	_elem = GetDocument().CreateElement(tagname)
	if !_elem.IsDefined() {
		// TODO: test HTMLUnknownElement
		return nil, fmt.Errorf("CreateComponentElement failed: invalid tagname %q", tagname)
	}

	// set the class, executing the class template
	classes = strings.Trim(classes, "")
	_elem.SetClasses(classes)

	strattrs = strings.Trim(strattrs, "")
	attrs, err := ParseAttributes(strattrs)
	if err != nil {
		// TODO: handle attribute parsing errors
	} else {
		_elem.SetAttributes(*attrs)
	}

	// if classtemplate != "" {
	// 	var tclass *template.Template
	// 	buf := new(bytes.Buffer)

	// 	tclass, _err = template.New("class").Parse(classtemplate)
	// 	if _err == nil {
	// 		data := TemplateData{
	// 			Me:     _composer,
	// 			Global: &GData,
	// 		}
	// 		_err = tclass.Execute(buf, data)
	// 	}
	// 	if _err == nil {
	// 		_elem.SetClassName(buf.String())
	// 	}
	// }

	// wrap the composer with the newly created component
	if _err == nil {
		_composer.Wrap(_elem)
	}

	return _elem, _err
}
