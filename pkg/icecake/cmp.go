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

var ComponentRegistry map[string]reflect.Type

/*****************************************************************************/

func init() {
	ComponentRegistry = make(map[string]reflect.Type, 0)
}

func RegisterComponentType(key string, typ reflect.Type) {
	// TODO: check component type and name convention with an hyphen aka "ick-XXXX"
	key = helper.Normalize(key)
	ComponentRegistry[key] = typ
}

/*****************************************************************************/

func CreateCompoundElement(_compounder HtmlCompounder) (_elem *Element, _err error) {
	// create the HTML element
	tagname, classtemplate := _compounder.Envelope()
	tagname = helper.Normalize(tagname)
	_elem = GetDocument().CreateElement(tagname)
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
			_err = tclass.Execute(buf, _compounder)
		}
		if _err == nil {
			_elem.SetClassName(buf.String())
		}
	}
	return _elem, _err
}
