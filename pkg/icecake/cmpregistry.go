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
	ickname string // unique ick-name of the component
	typ     reflect.Type
	css     string
	count   int
}

func (_cr componentRegEntry) String() string {
	return fmt.Sprintf("%s type:%s, n=%v, css:%v", _cr.ickname, _cr.typ.String(), _cr.count, _cr.css != "")
}

func (_cr componentRegEntry) IsFirst() bool {
	return _cr.count == 0
}

type ComponentRegistry struct {
	count   int
	entries map[string]*componentRegEntry
}

var TheCmpReg ComponentRegistry

func (_reg *ComponentRegistry) RegisterComponent(_cmp any) (_err error) {

	typ := reflect.TypeOf(_cmp)
	if typ.Kind() == reflect.Pointer {
		_err = fmt.Errorf("register component %q failed: must register a component not a pointer to a component", typ.String())
		log.Println(_err.Error())
		return _err
	}

	cmp, ok := reflect.New(typ).Interface().(HtmlComposer)
	if !ok {
		_err = fmt.Errorf("register component %q failed: must be an HtmlComposer", typ.String())
		log.Println(_err.Error())
		return _err
	}

	ickname := helper.Normalize(cmp.RegisterName())
	if !strings.HasPrefix(ickname, "ick-") {
		_err = fmt.Errorf("register component %q failed: name must start by 'ick-'", typ.String())
		log.Println(_err.Error())
		return _err
	}
	name := strings.TrimPrefix(ickname, "ick-")
	if len(name) == 0 {
		_err = fmt.Errorf("registering component %q failed: name missing", typ.String())
		log.Println(_err.Error())
		return _err
	}

	if _, found := _reg.entries[ickname]; found {
		log.Printf("registering component %q warning: already registered", ickname)
		return nil
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

func (_reg *ComponentRegistry) GetUniqueId(_composer HtmlComposer) (_id string) {
	// check if _composer matches with a registered component type
	regentry := _reg.LookupComponentType(reflect.TypeOf(_composer))
	if regentry == nil {
		log.Printf("GetUniqueId for a non registered component %q\n", reflect.TypeOf(_composer).String())
	}

	var idx int
	var name string

	// TODO: safe thread
	if regentry != nil {
		name = regentry.ickname
		idx = regentry.count
		regentry.count++
	} else {
		name = reflect.TypeOf(_composer).String()
		name = strings.TrimLeft(name, "*")
		if strings.Contains(name, ".") {
			_, name, _ = strings.Cut(name, ".")
		}
		name = strings.ToLower(name)
		idx = _reg.count
		_reg.count++
	}

	_id = name + "-" + strconv.Itoa(idx)
	return _id
}
