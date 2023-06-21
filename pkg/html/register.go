package html

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/icecake-framework/icecake/internal/helper"
	"github.com/icecake-framework/icecake/pkg/registry"
	"github.com/sunraylab/verbose"
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
func RegisterComposer(ickname string, composer any, css []string) (entry *registry.RegistryEntry, err error) {
	typ := reflect.TypeOf(composer)
	if typ.Kind() != reflect.Pointer {
		err = fmt.Errorf("registering composer %q failed: must register by reference not by value", typ.String())
		log.Println(err.Error())
		return nil, err
	}

	cmp, ok := composer.(HTMLComposer)
	if !ok {
		err = fmt.Errorf("registering composer %q failed: must implement HTMLComposer interface", typ.String())
		log.Println(err.Error())
		return nil, err
	}

	tag := cmp.Tag()
	if tag == nil {
		err = fmt.Errorf("registering composer %q failed: HTMLComposer Tag() must return a valid reference", typ.String())
		log.Println(err.Error())
		return nil, err
	}

	ickname = helper.Normalize(ickname)
	if !strings.HasPrefix(ickname, "ick-") {
		err = fmt.Errorf("registering composer %q failed: name must start by 'ick-'", typ.String())
		log.Println(err.Error())
		return nil, err
	}

	if len(ickname) <= 4 {
		err = fmt.Errorf("registering composer %q failed: name missing", typ.String())
		log.Println(err.Error())
		return nil, err
	}

	cmp.BuildTag(tag)
	if !tag.HasName() {
		log.Printf("registering composer %q warning: HTMLComposer without tag Builder", ickname)
	}

	if registry.IsRegistered(ickname) {
		log.Printf("registering composer %q warning: already registered\n", ickname)
		return nil, err
	}

	entry = registry.AddRegistryEntry(ickname, composer, css)

	verbose.Debug("registering composer %s(%s) with success", ickname, typ.String())

	return entry, nil
}

// TODO: RegisterHTMLString
// RegisterHTMLSnippet allows registering a simple HTMLSnippet with a simple line of code.
// The HTMLSnippet will be rendered every time the auto-closing ick-tag will be met within an html string.
//
// The _ickname must start by `ick-` followed by at least one character.
// An error is returned in the following cases:
//   - If the _ickname has already been registered
//   - If the _ickname does not meet tye pattern "ick-*"
//   - If the _composer does not implement the HTMLComposer interface
//   - If parsing _template attributes fails
// func RegisterHTMLSnippet(_ickname string, _template SnippetTemplate) (_err error) {
// 	s := new(HTMLSnippet)
// 	s.TagName = _template.TagName
// 	s.Body = _template.Body
// 	if _err = ParseAttributes(string(_template.Attributes), s); _err != nil {
// 		return _err
// 	}
// 	return RegisterComposer(_ickname, s, nil)
// }
