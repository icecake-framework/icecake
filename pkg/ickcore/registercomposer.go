package ickcore

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/icecake-framework/icecake/internal/helper"
	"github.com/lolorenzo777/verbose"
)

// RegisterComposer registers a _composer with the unique _ickname.
// Registering a _composer is required to enable rendering of a composer embedded into an html string with an auto-closing ick-tag.
//
//	<link rel="stylesheet" href="{_css[n]}">
//
// The ickname must start by `ick-` followed by at least one character.
// An error is returned in the following cases:
//   - If the ickname has already been registered
//   - If the ickname does not meet the pattern "ick-*"
//   - If the composer does not implement the ElementComposer interface
func RegisterComposer(icktagname string, composer any) (entry *RegistryEntry, err error) {
	// TODO: register - RegisterComposer should generate ickname automatically (see RenderSnippet), unless a same component can be registered with 2 names ?!

	// register by reference
	typ := reflect.TypeOf(composer)
	if typ.Kind() != reflect.Pointer {
		err = fmt.Errorf("RegisterComposer: %s(%v) must register by reference not by value", icktagname, typ.String())
		log.Println(err.Error())
		return nil, err
	}

	// must implement Composer at least
	_, iscmp := composer.(Composer)
	if !iscmp {
		err = fmt.Errorf("RegisterComposer: %s(%v) failed: must implement Composer interface", icktagname, typ.String())
		log.Println(err.Error())
		return nil, err
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
	if IsRegistered(icktagname) && reflect.TypeOf(GetRegistryEntry(icktagname).Component()).String() != typ.String() {
		err = fmt.Errorf("RegisterComposer: %s(%v) warning: already registered with another composer", icktagname, typ.String())
		log.Println(err.Error())
		return nil, err
	}

	entry = AddRegistryEntry(icktagname, composer)

	verbose.Debug("RegisterComposer: %s(%v) with success", icktagname, typ.String())

	return entry, nil
}
