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
	typ := reflect.TypeOf(composer)
	if typ.Kind() != reflect.Pointer {
		err = fmt.Errorf("registering composer %q failed: must register by reference not by value", typ.String())
		log.Println(err.Error())
		return nil, err
	}

	_, iscmp := composer.(HTMLContentComposer)
	if !iscmp {
		err = fmt.Errorf("registering composer %q failed: must implement HTMLComposer interface", typ.String())
		log.Println(err.Error())
		return nil, err
	}

	cmptag, istagger := composer.(TagBuilder)
	if istagger {
		if !cmptag.BuildTag().HasRendering() {
			log.Printf("registering composer %q warning: TagBuilder without rendering", icktagname)
		}
	}

	icktagname = helper.Normalize(icktagname)
	if !strings.HasPrefix(icktagname, "ick-") {
		err = fmt.Errorf("registering composer %q failed: name must start by 'ick-'", typ.String())
		log.Println(err.Error())
		return nil, err
	}

	if len(icktagname) <= 4 {
		err = fmt.Errorf("registering composer %q failed: name missing", typ.String())
		log.Println(err.Error())
		return nil, err
	}

	if registry.IsRegistered(icktagname) {
		log.Printf("registering composer %q warning: already registered\n", icktagname)
		return nil, err
	}

	entry = registry.AddRegistryEntry(icktagname, composer)

	verbose.Debug("registering composer %s(%s) with success", icktagname, typ.String())

	return entry, nil
}
