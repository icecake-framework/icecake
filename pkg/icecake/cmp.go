package ick

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/sunraylab/icecake/internal/helper"
)

type HtmlCompounder interface {
	Envelope() (_tagname string, _classTemplate string)
	Template() (_html string)
}

type HtmlListener interface {
	AddListeners()
}

// type StyleCompounder interface {
// 	Style() string
// }

var gComponents int

func GetNextComponentIndex() (_index int) {
	_index = gComponents + 1
	gComponents++
	return _index
}

var GComponentRegistry map[string]reflect.Type

/*****************************************************************************/

func init() {
	GComponentRegistry = make(map[string]reflect.Type, 0)
}

func RegisterComponentType(key string, typ reflect.Type) {
	// TODO: check component type and name convention with an hyphen aka "ick-XXXX"
	key = helper.Normalize(key)
	GComponentRegistry[key] = typ
	ConsoleWarnf("RegisterComponentType: %s %q", key, typ.String())
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

func CreateCompoundElement(_compounder HtmlCompounder) (_elem *Element, _err error) {
	// create the HTML element
	tagname, classtemplate := _compounder.Envelope()
	tagname = helper.Normalize(tagname)
	_elem = App().CreateElement(tagname)
	if !_elem.IsDefined() {
		// TODO: test HTMLUnknownElement
		return nil, fmt.Errorf("CreateCompoundElement failed: invalid tagname %q", tagname)
	}

	// set the class, executing the class template
	classtemplate = strings.Trim(classtemplate, "")
	if classtemplate != "" {
		var tclass *template.Template
		buf := new(bytes.Buffer)

		tclass, _err = template.New("class").Parse(classtemplate)
		if _err == nil {
			data := TemplateData{
				Me:     _compounder,
				Global: &GData,
			}
			_err = tclass.Execute(buf, data)
		}
		if _err == nil {
			_elem.SetClassName(buf.String())
		}
	}
	return _elem, _err
}
