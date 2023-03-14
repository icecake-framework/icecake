package ick

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/sunraylab/icecake/internal/helper"
)

type componentRegEntry struct {
	ickname string
	typ     reflect.Type
	css     string
	count   int
}

func (_cr componentRegEntry) String() string {
	return fmt.Sprintf("%s type:%s, n=%v, css:%v", _cr.ickname, _cr.typ.String(), _cr.count, _cr.css != "")
}

func (_cr componentRegEntry) IsFirst() bool {
	return _cr.count == 1
}

type ComponentRegistry struct {
	count   int
	entries map[string]*componentRegEntry
}

var TheCmpReg ComponentRegistry

func (_reg *ComponentRegistry) RegisterComponent(_cmp any) error {

	typ := reflect.TypeOf(_cmp)
	if typ.Kind() == reflect.Pointer {
		err := fmt.Errorf("register component %q failed: must register a component not a pointer to a component", typ.String())
		log.Println(err.Error())
		return err
	}

	cmp, ok := reflect.New(typ).Interface().(HtmlComposer)
	if !ok {
		err := fmt.Errorf("register component %q failed: must be an HtmlComposer", typ.String())
		log.Println(err.Error())
		return err
	}

	ickname := helper.Normalize(cmp.RegisterName())
	if !strings.HasPrefix(ickname, "ick-") {
		err := fmt.Errorf("register component %q failed: name must start by 'ick-'", typ.String())
		log.Println(err.Error())
		return err
	}
	name := strings.TrimPrefix(ickname, "ick-")
	if len(name) == 0 {
		err := fmt.Errorf("registering component %q failed: name missing", typ.String())
		log.Println(err.Error())
		return err
	}

	if _, found := _reg.entries[ickname]; found {
		err := fmt.Errorf("registering component %q failed: already registered", ickname)
		log.Println(err.Error())
		return err
	}

	entry := componentRegEntry{
		ickname: ickname,
		typ:     typ,
		css:     cmp.RegisterCSS(),
		count:   0,
	}
	if _reg.entries == nil {
		_reg.entries = make(map[string]*componentRegEntry, 1)
	}
	_reg.entries[ickname] = &entry
	log.Printf("component %s %q registered\n", ickname, typ.String())
	return nil
}

func (_reg *ComponentRegistry) LookupComponentType(typ reflect.Type) *componentRegEntry {
	st := strings.TrimLeft(typ.String(), "*")
	for _, v := range _reg.entries {
		sv := strings.TrimLeft(v.typ.String(), "*")
		if sv == st {
			return v
		}
	}
	return nil
}

func (_reg *ComponentRegistry) LookupComponent(_ickname string) *componentRegEntry {
	for k, v := range _reg.entries {
		if k == _ickname {
			return v
		}
	}
	return nil
}

// func (_reg *ComponentRegistry) NextComponentId(_ickname string) (_id string, _first bool) {
// 	_reg.count++

// 	entry := _reg.entries[_ickname]
// 	if entry != nil {
// 		cidx := entry.count + 1
// 		entry.count++
// 		_first = cidx == 1

// 		_id = _ickname + "-" + strconv.Itoa(cidx)
// 	}

// 	return _id, _first
// }

func (_reg *ComponentRegistry) GetUniqueId(_composer HtmlComposer) (_id string) {
	// check if _composer matches with a registered component type
	regentry := _reg.LookupComponentType(reflect.TypeOf(_composer))
	if regentry == nil {
		log.Printf("GetUniqueId for a non registered component %q\n", reflect.TypeOf(_composer).String())
	}

	var idx int
	var name string

	// TODO: safe thread
	_reg.count++
	if regentry != nil {
		name = regentry.ickname
		idx = regentry.count + 1
		regentry.count++
	} else {
		name = reflect.TypeOf(_composer).String()
		name = strings.TrimLeft(name, "*")
		if strings.Contains(name, ".") {
			_, name, _ = strings.Cut(name, ".")
		}
		name = strings.ToLower(name)
		idx = _reg.count
	}

	_id = name + "-" + strconv.Itoa(idx)
	return _id
}

/*****************************************************************************/

// func (_reg *ComponentRegistry) CreateComponent(_composer UIComposer) (_id string, _newcmp *UIComponent, _err error) {

// 	// check if _composer matches with a registered component type, and get a fresh component id
// 	regentry := _app.LookupComponent(reflect.TypeOf(_composer))
// 	if regentry == nil {
// 		return "", nil, errors.ConsoleErrorf("CreateComponent failed: non registered component %q\n", reflect.TypeOf(_composer).String())
// 	}
// 	var first bool
// 	_id, first = _app.NextComponentId(regentry.ickname)

// 	// create the HTML element
// 	tagname, strclasses, strattrs := _composer.Container(_id)
// 	tagname = helper.Normalize(tagname)
// 	elem := GetDocument().CreateElement(tagname)
// 	if !elem.IsDefined() {
// 		// TODO: check HTMLUnknownElement returns
// 		return "", nil, errors.ConsoleErrorf("CreateComponent %q failed: invalid tagname %q\n", regentry.ickname, tagname)
// 	}

// 	// set the container classes
// 	_err = elem.Classes().ParseTokens(strclasses)
// 	if _err != nil {
// 		return "", nil, errors.ConsoleErrorf("CreateComponent %q failed: %s\n", regentry.ickname, _err)
// 	}

// 	// set the container attributes
// 	_err = elem.Attributes().ParseAttributes(strattrs)
// 	if _err != nil {
// 		return "", nil, errors.ConsoleErrorf("CreateComponent %q failed: %s\n", regentry.ickname, _err)
// 	}

// 	// wrap the composer with the newly created component
// 	_newcmp = CastUIComponent(elem)
// 	// TODO: remove wrapping and wotk with _newcomp
// 	_composer.Wrap(elem)

// 	_newcmp.SetId(_id)

// 	// init classes and attributes
// 	_newcmp.Classes().AddClasses(*_composer.Classes())
// 	_newcmp.Attributes().SetAttributes(*_composer.Attributes())

// 	// add css
// 	// DEBUG: fmt.Println(regentry.String(), first)

// 	if first && regentry.css != "" {
// 		style := _app.CreateElement("style")
// 		style.SetInnerHTML(regentry.css)
// 		_app.Head().AppendChild(&style.Node)
// 	}

// 	return _id, _newcmp, _err
// }
