package html

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/icecake-framework/icecake/internal/helper"
	"github.com/icecake-framework/icecake/pkg/registry"
	"github.com/lolorenzo777/verbose"
)

// RegisterComposer registers a _composer with the unique _ickname.
// Registering a _composer is required to enable rendering of a composer embedded into an html string with an auto-closing ick-tag.
//
// _css is a list of required stylesheet for the component. During the first instantiation of the component,
// these stylesheets are included in the header of the html file as links entries:
//
//	<link rel="stylesheet" href="{_css[n]}">
//
// The _ickname must start by `ick-` followed by at least one character.
// An error is returned in the following cases:
//   - If the _ickname has already been registered
//   - If the _ickname does not meet the pattern "ick-*"
//   - If the _composer does not implement the HTMLComposer interface
func RegisterComposer(icktagname string, composer any) (entry *registry.RegistryEntry, err error) {
	// TODO: RegisterComposer should generate ickname automatically (see RenderSnippet), unless a same component can be registered with 2 names ?!

	// register by reference
	typ := reflect.TypeOf(composer)
	if typ.Kind() != reflect.Pointer {
		err = fmt.Errorf("RegisterComposer: %s(%v) must register by reference not by value", icktagname, typ.String())
		log.Println(err.Error())
		return nil, err
	}

	// must implement HTMLContentComposer at least
	_, iscmp := composer.(HTMLContentComposer)
	if !iscmp {
		err = fmt.Errorf("RegisterComposer: %s(%v) failed: must implement HTMLComposer interface", icktagname, typ.String())
		log.Println(err.Error())
		return nil, err
	}

	// a tagbuilder should have tag rendering
	cmptag, istagger := composer.(TagBuilder)
	if istagger {
		if !cmptag.BuildTag().HasRendering() {
			err = fmt.Errorf("RegisterComposer: %s(%v) warning: TagBuilder without rendering", icktagname, icktagname)
			log.Println(err.Error())
			return nil, err
		}
	}

	// naming prefix
	icktagname = helper.Normalize(icktagname)
	if !strings.HasPrefix(icktagname, "ick-") {
		err = fmt.Errorf("RegisterComposer: %s(%v) failed: name must start by 'ick-'", icktagname, typ.String())
		log.Println(err.Error())
		return nil, err
	}

	// non empty name
	if len(icktagname) <= 4 {
		err = fmt.Errorf("RegisterComposer: %s(%v) failed: name missing", icktagname, typ.String())
		log.Println(err.Error())
		return nil, err
	}

	// already registeredwith anoher
	if registry.IsRegistered(icktagname) && reflect.TypeOf(registry.GetRegistryEntry(icktagname).Component()).String() != typ.String() {
		err = fmt.Errorf("RegisterComposer: %s(%v) warning: already registered with another composer", icktagname, typ.String())
		log.Println(err.Error())
		return nil, err
	}

	entry = registry.AddRegistryEntry(icktagname, composer)

	verbose.Debug("RegisterComposer: %s(%v) with success", icktagname, typ.String())

	return entry, nil
}
