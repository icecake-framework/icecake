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

	c1 := new(HTMLSnippet)
	_, err := RegisterComposer("snippet", *c1, nil)
	assert.ErrorContains(t, err, "not a component")

	i := new(int)
	_, err = RegisterComposer("snippet", i, nil)
	assert.ErrorContains(t, err, "must be an HTMLComposer")

	_, err = RegisterComposer("ick-testsnippet1", &testsnippet1{}, nil)
	assert.NoError(t, err)

	_, err = RegisterComposer("ick-testsnippet1", &testsnippet1{}, nil)
	assert.NoError(t, err) // only log "already registered"
}
