package html

import (
	"testing"

	"github.com/icecake-framework/icecake/pkg/registry"
	"github.com/stretchr/testify/assert"
)

func TestLookupRegistryEntry(t *testing.T) {

	registry.ResetRegistry()
	registry.AddRegistryEntry("ick-snippet", &HTMLSnippet{}, nil)

	r := registry.LookupRegistryEntry(&HTMLSnippet{})
	assert.NotNil(t, r)

	r = registry.LookupRegistryEntry(HTMLSnippet{})
	assert.Nil(t, r)
}

func TestRegisterComposer(t *testing.T) {

	registry.ResetRegistry()

	// by reference
	c1 := new(HTMLSnippet)
	_, err := RegisterComposer("snippet", *c1, nil)
	assert.ErrorContains(t, err, "not by value")

	// HTMLcomposer implementation
	i := new(int)
	_, err = RegisterComposer("snippet", i, nil)
	assert.ErrorContains(t, err, "must implement HTMLComposer interface")

	// empty tag
	_, err = RegisterComposer("ick-testsnippet1", &testcustomcomposer{}, nil)
	assert.ErrorContains(t, err, "Tag() must return a valid reference")

	// naming prefix
	_, err = RegisterComposer("snippet", &testsnippet1{}, nil)
	assert.ErrorContains(t, err, "name must start by 'ick-'")

	// name
	_, err = RegisterComposer("ick-", &testsnippet1{}, nil)
	assert.ErrorContains(t, err, "name missing")

	// log tag builder
	_, err = RegisterComposer("ick-testsnippet1", &testsnippet1{}, nil)
	assert.NoError(t, err)

	// log "already registered"
	_, err = RegisterComposer("ick-testsnippet1", &testsnippet1{}, nil)
	assert.NoError(t, err)
}
