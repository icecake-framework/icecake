package ick

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"syscall/js"
	"text/template"

	"github.com/sunraylab/icecake/internal/helper"
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
	Wrap(js.Value)
	Envelope() (_tagname string, _classTemplate string)
}

type HtmlTemplater interface {
	Template() (_html string)
}
type HtmlListener interface {
	Wrap(js.Value)
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

func RegisterComponentType(key string, cmp any) {
	key = helper.Normalize(key)
	if !strings.HasPrefix(key, "ick-") {
		ConsoleErrorf("RegisterComponentType faild: key %q does not match allowed pattern", key)
		return
	}
	name := strings.TrimPrefix(key, "ick-")
	if len(name) == 0 {
		ConsoleErrorf("RegisterComponentType faild: invalid key name %q", key)
		return
	}

	typ := reflect.TypeOf(cmp)
	if typ.Kind() == reflect.Pointer {
		ConsoleErrorf("RegisterComponentType faild: must register a component not a pointer to a component %q", typ.String())
		return
	}

	if _, found := typ.FieldByName("Component"); !found {
		ConsoleErrorf("RegisterComponentType faild: your component must embed the ick.Component value")
		return
	}

	if _, found := GComponentRegistry[key]; found {
		ConsoleErrorf("RegisterComponentType faild: %q already registered", key)
		return
	}

	GComponentRegistry[key] = typ
	ConsoleLogf("RegisterComponentType: %s %q", key, typ.String())
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
	tagname, classtemplate := _composer.Envelope()
	tagname = helper.Normalize(tagname)
	_elem = App().CreateElement(tagname)
	if !_elem.IsDefined() {
		// TODO: test HTMLUnknownElement
		return nil, fmt.Errorf("CreateComponentElement failed: invalid tagname %q", tagname)
	}

	// set the class, executing the class template
	classtemplate = strings.Trim(classtemplate, "")
	if classtemplate != "" {
		var tclass *template.Template
		buf := new(bytes.Buffer)

		tclass, _err = template.New("class").Parse(classtemplate)
		if _err == nil {
			data := TemplateData{
				Me:     _composer,
				Global: &GData,
			}
			_err = tclass.Execute(buf, data)
		}
		if _err == nil {
			_elem.SetClassName(buf.String())
		}
	}

	// wrap the composer with the newly created component
	if _err == nil {
		_composer.Wrap(_elem.JSValue())
	}

	return _elem, _err
}
