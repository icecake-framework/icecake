package registry

import (
	"reflect"
	"strconv"

	"github.com/sunraylab/icecake/internal/helper"
)

// The registry is the global Registry instantiated once and used by the HtmlSnippet and other components.
var theRegistry registry

// Registry stores definition of components in a map, by unique name
type registry struct {
	entries map[string]RegistryEntry
}

// RegistryEntry defines a component
type RegistryEntry struct {
	name  string // unique name of the component
	cmp   any    // The component type that must be instantiated
	count int    // number of time this object has already been instantiated
	//	css   string       // TODO: handle the unique css related to this component. will be added into the head of the page
}

func (_r RegistryEntry) Name() string {
	return _r.name
}

func (_r RegistryEntry) Component() any {
	return _r.cmp
}

func IsRegistered(_name string) bool {
	theRegistry.init()
	_, found := theRegistry.entries[_name]
	return found
}

func AddRegistryEntry(_name string, _cmp any) {
	theRegistry.init()
	_name = helper.Normalize(_name)

	entry := RegistryEntry{
		name: _name,
		cmp:  _cmp,
		//		css:   cmp.RegisterCSS(),
		count: 0,
	}
	theRegistry.entries[_name] = entry
}

func GetRegistryEntry(name string) RegistryEntry {
	theRegistry.init()
	name = helper.Normalize(name)
	if name == "" {
		name = "ick"
	}
	regentry, found := theRegistry.entries[name]
	if !found {
		regentry = RegistryEntry{name: name}
	}
	return regentry
}

// _cmp must be a pointer
func LookupRegistryEntry(_cmp any) *RegistryEntry {
	theRegistry.init()
	typ := reflect.TypeOf(_cmp)
	for _, v := range theRegistry.entries {
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
	theRegistry.entries[regentry.name] = regentry
	return regentry.name + "-" + strconv.Itoa(idx)
}

func ResetRegistry() {
	theRegistry.entries = make(map[string]RegistryEntry, 1)
}

/******************************************************************************
* Private
******************************************************************************/

func (_reg *registry) init() {
	if _reg.entries == nil {
		_reg.entries = make(map[string]RegistryEntry, 0)
	}
}
