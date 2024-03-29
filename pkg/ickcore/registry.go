package ickcore

import (
	"reflect"
	"strconv"
	"sync"

	"github.com/icecake-framework/icecake/internal/helper"
)

// The registry is the global Registry instantiated once and used by the HtmlSnippet and other components.
var theRegistry registry

// RegistryEntry defines a component
type RegistryEntry struct {
	mu         sync.Mutex
	icktagname string // unique name of the component
	cmp        any    // The component type that must be instantiated. This must be a reference.
	count      int    // number of time this cmp has already been instantiated

	// csslinkref     []string // slice of required stylesheet link ref for this component. will be added once into the head of the page
	// csslinkmounted bool
}

func (_r *RegistryEntry) Count() int {
	_r.mu.Lock()
	_r.count++
	i := _r.count
	_r.mu.Unlock()
	return i
}

// Name returns the unique name of the component
// starting with the `ick-` prefix.
func (_r RegistryEntry) IckTagName() string {
	return _r.icktagname
}

// Component returns the component type that must be instantiated
func (_r RegistryEntry) Component() any {
	return _r.cmp
}

// IsRegistered returns if the _name has already been regystered
func IsRegistered(_name string) bool {
	theRegistry.init()
	_, found := theRegistry.entries[_name]
	return found
}

func (_reg *registry) init() {
	if _reg.entries == nil {
		_reg.entries = make(map[string]*RegistryEntry, 0)
	}
}

// Registry stores definition of components in a map, by unique name
type registry struct {
	entries map[string]*RegistryEntry
}

func Map() map[string]*RegistryEntry {
	theRegistry.init()
	return theRegistry.entries
}

// AddRegistryEntry create a new RegistryEntry and add it to the global and private registry.
// No check is done on name. If name is already registered a new registryentry overwrites the existing one.
// css is optional and can be nil.
func AddRegistryEntry(name string, cmp any) *RegistryEntry {
	theRegistry.init()
	name = helper.Normalize(name)

	entry := RegistryEntry{
		icktagname: name,
		cmp:        cmp,
		count:      0,
	}
	theRegistry.entries[name] = &entry
	return &entry
}

// GetRegistryEntry returns the RegistryEntry corresponding to the _name.
// If _name is empty GetRegistryEntry returns the RegistryEntry for "ick".
// If _name does not correspond to an existing entry in the resitry, then
// GetRegistryEntry create a default entry in the registry with that name.
// Also GetRegistryEntry always returns a RegistryEntry.
func GetRegistryEntry(name string) *RegistryEntry {
	theRegistry.init()
	name = helper.Normalize(name)
	if name == "" {
		name = "ick"
	}
	regentry, found := theRegistry.entries[name]
	if !found {
		regentry = &RegistryEntry{icktagname: name}
	}
	return regentry
}

// LookupRegistryEntry lookup for _cmp in the registry and returns the first RegistryEntry with matching type.
// _cmp must be a pointer, like it was registered with AddRegistryEntry.
// Return nil if nothing is found.
func LookupRegistryEntry(cmp any) *RegistryEntry {
	theRegistry.init()
	typ := reflect.TypeOf(cmp)
	for _, v := range theRegistry.entries {
		tv := reflect.TypeOf(v.cmp)
		if tv == typ {
			return v
		}
	}
	return nil
}

// GetUniqueId returns a unique id starting with prefix.
// if prefix is empty "ick-" is used to prefix the returned id.
// The returned id is always lowercase.
// GetUniqueId is thread safe.
func GetUniqueId(prefix string) (idx int, uid string) {
	regentry := GetRegistryEntry(prefix)
	idx = regentry.Count()
	theRegistry.entries[regentry.icktagname] = regentry
	return idx, regentry.icktagname + "-" + strconv.Itoa(idx)
}

// ResetRegistry is only used for testing
func ResetRegistry() {
	theRegistry.entries = make(map[string]*RegistryEntry, 1)
}
