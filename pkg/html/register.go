package html

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/registry"
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
func RegisterComposer(_ickname string, _composer any, _css []string) (_err error) {
	typ := reflect.TypeOf(_composer)
	if typ.Kind() != reflect.Pointer {
		_err = fmt.Errorf("register composering %q failed: must register a pointer to a component not a component", typ.String())
		log.Println(_err.Error())
		return _err
	}

	_, ok := _composer.(HTMLComposer)
	if !ok {
		_err = fmt.Errorf("registering composer %q failed: must be an HTMLComposer", typ.String())
		log.Println(_err.Error())
		return _err
	}

	_ickname = helper.Normalize(_ickname)
	if !strings.HasPrefix(_ickname, "ick-") {
		_err = fmt.Errorf("registering composer %q failed: name must start by 'ick-'", typ.String())
		log.Println(_err.Error())
		return _err
	}
	if len(_ickname) == 0 {
		_err = fmt.Errorf("registering composer %q failed: name missing", typ.String())
		log.Println(_err.Error())
		return _err
	}

	if registry.IsRegistered(_ickname) {
		log.Printf("registering composer %q warning: already registered", _ickname)
		return _err
	}

	registry.AddRegistryEntry(_ickname, _composer, _css)

	// DEBUG:
	fmt.Printf("registering composer %s(%s) with success\n", _ickname, typ.String())
	return nil
}

// RegisterHTMLSnippet allows registering a simple HTMLSnippet with a simple line of code.
// The HTMLSnippet will be rendered every time the auto-closing ick-tag will be met within an html string.
//
// The _ickname must start by `ick-` followed by at least one character.
// An error is returned in the following cases:
//   - If the _ickname has already been registered
//   - If the _ickname does not meet tye pattern "ick-*"
//   - If the _composer does not implement the HTMLComposer interface
//   - If parsing _template attributes fails
func RegisterHTMLSnippet(_ickname string, _template SnippetTemplate) (_err error) {
	s := new(HTMLSnippet)
	s.TagName = _template.TagName
	s.Body = _template.Body
	if _err = ParseAttributes(string(_template.Attributes), s); _err != nil {
		return _err
	}
	return RegisterComposer(_ickname, s, nil)
}
