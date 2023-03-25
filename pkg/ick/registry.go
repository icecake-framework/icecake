package ick

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/sunraylab/icecake/internal/helper"
)

// The registry is the global Registry instantiated once and used by the HtmlSnippet and other components.
var TheRegistry registry

// RegistryEntry define a component
type RegistryEntry struct {
	name string // unique name of the component
	cmp  any    // The component type that must be instantiated
	// typ   reflect.Type // The component tyoe that must be instantiated
	count int // number of time this object has already been instantiated
	//	css   string       // the unique css related to this component. will be added into the head of the page
}

// Registry stores definition of components by their name
type registry struct {
	entries map[string]RegistryEntry
}

func RegisterComposer(ickname string, _composer any) (_err error) {
	reg := &TheRegistry
	reg.init()

	typ := reflect.TypeOf(_composer)
	if typ.Kind() != reflect.Pointer {
		_err = fmt.Errorf("register component %q failed: must register a pointer to a component not a component", typ.String())
		log.Println(_err.Error())
		return _err
	}

	// _, ok := reflect.New(typ).Interface().(HtmlComposer)
	_, ok := _composer.(HtmlComposer)
	if !ok {
		_err = fmt.Errorf("register component %q failed: must be an HtmlComposer", typ.String())
		log.Println(_err.Error())
		return _err
	}

	ickname = helper.Normalize(ickname)
	if !strings.HasPrefix(ickname, "ick-") {
		_err = fmt.Errorf("register component %q failed: name must start by 'ick-'", typ.String())
		log.Println(_err.Error())
		return _err
	}
	if len(ickname) == 0 {
		_err = fmt.Errorf("registering component %q failed: name missing", typ.String())
		log.Println(_err.Error())
		return _err
	}

	if _, found := reg.entries[ickname]; found {
		log.Printf("registering component %q warning: already registered", ickname)
		return nil
	}

	entry := RegistryEntry{
		name: ickname,
		cmp:  _composer,
		// typ:  typ,
		//		css:   cmp.RegisterCSS(),
		count: 0,
	}
	reg.entries[ickname] = entry
	log.Printf("component %q(%s) registered\n", ickname, typ.String())
	return nil
}

func RegisterSnippet(ickname string, tagname HTML, body HTML) (_err error) {
	s := new(HtmlSnippet)
	s.TagName = tagname
	s.Body = body
	return RegisterComposer(ickname, s)
}

func GetRegistryEntry(name string) RegistryEntry {
	TheRegistry.init()
	name = helper.Normalize(name)
	if name == "" {
		name = "ick"
	}
	regentry, found := TheRegistry.entries[name]
	if !found {
		regentry = RegistryEntry{name: name}
	}
	return regentry
}

// _cmp must be a pointer
func LookupRegistryEntry(_cmp any) *RegistryEntry {
	TheRegistry.init()
	typ := reflect.TypeOf(_cmp)
	for _, v := range TheRegistry.entries {
		tv := reflect.TypeOf(v.cmp)
		if tv == typ {
			return &v
		}
	}
	return nil
}

// GetUniqueId.
// TODO: safe thread
func GetUniqueId(name string) string {
	regentry := GetRegistryEntry(name)
	regentry.count++
	idx := regentry.count
	TheRegistry.entries[regentry.name] = regentry
	return regentry.name + "-" + strconv.Itoa(idx)
}

/******************************************************************************
* Private
******************************************************************************/

func (_reg *registry) init() {
	if _reg.entries == nil {
		_reg.entries = make(map[string]RegistryEntry, 1)
	}
}
