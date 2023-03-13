package ick

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/errors"
)

/*****************************************************************************
* WebApp
******************************************************************************/

var App *WebApp

func init() {
	App = NewWebApp()
}

// WebApp
type WebApp struct {
	Document // The embedded DOM document

	cmpCount    int
	CmpRegistry map[string]*componentRegEntry

	browser Window // The Global JS Window object
}

// NewWebApp is the WebApp factory. Must be call once at the begining of the wasm main code.
func NewWebApp() *WebApp {
	webapp := new(WebApp)
	webapp.browser.Wrap(GetWindow())
	webapp.Document.Wrap(GetDocument())

	webapp.CmpRegistry = make(map[string]*componentRegEntry, 0)

	return webapp
}

// Browser returns the DOM.Window object
func (_app *WebApp) Browser() *Window {
	return &_app.browser
}

type componentRegEntry struct {
	ickname string
	typ     reflect.Type
	css     string
	count   int
}

func (_cr componentRegEntry) String() string {
	return fmt.Sprintf("%s type:%s, n=%v, css:%v", _cr.ickname, _cr.typ.String(), _cr.count, _cr.css != "")
}

func (_app *WebApp) NextComponentId(_ickname string) (_id string, _first bool) {
	_app.cmpCount++

	entry := _app.CmpRegistry[_ickname]
	if entry != nil {
		cidx := entry.count + 1
		entry.count++
		_first = cidx == 1

		_id = _ickname + "-" + strconv.Itoa(cidx)
	}

	return _id, _first
}

func (_app *WebApp) RegisterComponent(_ickname string, _cmp any, _css string) error {
	_ickname = helper.Normalize(_ickname)
	if !strings.HasPrefix(_ickname, "ick-") {
		return errors.ConsoleErrorf("RegisterComponentType %q failed: key name must start by 'ick-'\n", _ickname)
	}
	name := strings.TrimPrefix(_ickname, "ick-")
	if len(name) == 0 {
		return errors.ConsoleErrorf("RegisterComponentType %q failed: name missing\n", _ickname)
	}

	typ := reflect.TypeOf(_cmp)
	if typ.Kind() == reflect.Pointer {
		return errors.ConsoleErrorf("RegisterComponentType %q failed: must register a component not a pointer to a component %q\n", _ickname, typ.String())
	}

	if _, found := typ.FieldByName("UIComponent"); !found {
		return errors.ConsoleErrorf("RegisterComponentType %q failed: your component must embed the ick.UIComponent value\n", _ickname)
	}

	if _, found := _app.CmpRegistry[_ickname]; found {
		return errors.ConsoleErrorf("RegisterComponentType %q failed: already registered\n", _ickname)
	}
	entry := componentRegEntry{
		ickname: _ickname,
		typ:     typ,
		css:     _css,
		count:   0,
	}
	_app.CmpRegistry[_ickname] = &entry
	return errors.ConsoleLogf("RegisterComponentType: %s %q\n", _ickname, typ.String())
}

func (_app *WebApp) LookupComponent(typ reflect.Type) *componentRegEntry {
	st := strings.TrimLeft(typ.String(), "*")
	for _, v := range _app.CmpRegistry {
		sv := strings.TrimLeft(v.typ.String(), "*")
		if sv == st {
			return v
		}
	}
	return nil
}

/*****************************************************************************/

func (_app *WebApp) CreateComponent(_composer Composer) (_id string, _newcmp *UIComponent, _err error) {

	// check if _composer matches with a registered component type, and get a fresh component id
	regentry := _app.LookupComponent(reflect.TypeOf(_composer))
	if regentry == nil {
		return "", nil, errors.ConsoleErrorf("CreateComponent failed: non registered component %q\n", reflect.TypeOf(_composer).String())
	}
	var first bool
	_id, first = _app.NextComponentId(regentry.ickname)

	// create the HTML element
	tagname, strclasses, strattrs := _composer.Container(_id)
	tagname = helper.Normalize(tagname)
	elem := GetDocument().CreateElement(tagname)
	if !elem.IsDefined() {
		// TODO: check HTMLUnknownElement returns
		return "", nil, errors.ConsoleErrorf("CreateComponent %q failed: invalid tagname %q\n", regentry.ickname, tagname)
	}

	// set the container classes
	_err = elem.Classes().ParseTokens(strclasses)
	if _err != nil {
		return "", nil, errors.ConsoleErrorf("CreateComponent %q failed: %s\n", regentry.ickname, _err)
	}

	// set the container attributes
	_err = elem.Attributes().ParseAttributes(strattrs)
	if _err != nil {
		return "", nil, errors.ConsoleErrorf("CreateComponent %q failed: %s\n", regentry.ickname, _err)
	}

	// wrap the composer with the newly created component
	_newcmp = CastUIComponent(elem)
	// TODO: remove wrapping and wotk with _newcomp
	_composer.Wrap(elem)

	_newcmp.SetId(_id)

	// init classes and attributes
	_newcmp.Classes().AddClasses(*_composer.Classes())
	_newcmp.Attributes().SetAttributes(*_composer.Attributes())

	// add css
	// DEBUG: fmt.Println(regentry.String(), first)

	if first && regentry.css != "" {
		style := _app.CreateElement("style")
		style.SetInnerHTML(regentry.css)
		_app.Head().AppendChild(&style.Node)
	}

	return _id, _newcmp, _err
}
