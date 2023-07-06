package html

import (
	"testing"

	"github.com/icecake-framework/icecake/pkg/ickcore"
	"github.com/stretchr/testify/assert"
)

func TestLookupRegistryEntry(t *testing.T) {
	ickcore.ResetRegistry()
	ickcore.AddRegistryEntry("ick-snippet", &HTMLSnippet{})

	r := ickcore.LookupRegistryEntry(&HTMLSnippet{})
	assert.NotNil(t, r)

	r = ickcore.LookupRegistryEntry(HTMLSnippet{})
	assert.Nil(t, r)
}

func TestRegisterComposer(t *testing.T) {

	ickcore.ResetRegistry()

	// by reference
	c1 := new(HTMLSnippet)
	_, err := RegisterComposer("mysnippet", *c1)
	assert.ErrorContains(t, err, "not by value")

	// HTMLcomposer implementation
	i := new(int)
	_, err = RegisterComposer("mysnippet", i)
	assert.ErrorContains(t, err, "must implement HTMLComposer interface")

	// empty tag
	_, err = RegisterComposer("ick-testsnippet1", &testcustomcomposer{})
	assert.ErrorContains(t, err, "TagBuilder without rendering")

	// naming prefix
	_, err = RegisterComposer("snippet", &testsnippet0{})
	assert.ErrorContains(t, err, "name must start by 'ick-'")

	// name
	_, err = RegisterComposer("ick-", &testsnippet0{})
	assert.ErrorContains(t, err, "name missing")

	// No Error
	_, err = RegisterComposer("ick-testsnippet1", &testsnippet1{})
	assert.NoError(t, err)

	// registered twice
	_, err = RegisterComposer("ick-testsnippet1", &testsnippet0{})
	assert.Error(t, err)
	_, err = RegisterComposer("ick-testsnippet1", &testsnippet1{})
	assert.NoError(t, err)
}
