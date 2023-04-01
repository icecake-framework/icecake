package ick

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/html"
	"github.com/sunraylab/icecake/pkg/registry"
)

// Register a composer
func RegisterComposer(ickname string, _composer any) (_err error) {
	typ := reflect.TypeOf(_composer)
	if typ.Kind() != reflect.Pointer {
		_err = fmt.Errorf("register component %q failed: must register a pointer to a component not a component", typ.String())
		log.Println(_err.Error())
		return _err
	}

	_, ok := _composer.(html.HTMLComposer)
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

	if registry.IsRegistered(ickname) {
		log.Printf("registering component %q warning: already registered", ickname)
		return nil
	}

	registry.AddRegistryEntry(ickname, _composer)
	log.Printf("component %q(%s) registered\n", ickname, typ.String())
	return nil
}

// register a default snippet
func RegisterDefaultSnippet(ickname string, tagname html.String, body html.String) (_err error) {
	s := new(html.HTMLSnippet)
	s.TagName = tagname
	s.Body = body
	return RegisterComposer(ickname, s)
}
