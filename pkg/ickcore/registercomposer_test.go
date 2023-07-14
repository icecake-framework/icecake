package ickcore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLookupRegistryEntry(t *testing.T) {
	ResetRegistry()
	AddRegistryEntry("ick-snippet", &BareSnippet{})

	r := LookupRegistryEntry(&BareSnippet{})
	assert.NotNil(t, r)

	r = LookupRegistryEntry(BareSnippet{})
	assert.Nil(t, r)
}

func TestRegisterComposer(t *testing.T) {

	ResetRegistry()

	// by reference
	c1 := new(BareSnippet)
	_, err := RegisterComposer("mysnippet", *c1)
	assert.ErrorContains(t, err, "not by value")

	// composer implementation
	i := new(int)
	_, err = RegisterComposer("mysnippet", i)
	assert.ErrorContains(t, err, "must implement Composer interface")

	// naming prefix
	_, err = RegisterComposer("snippet", &BareSnippet{})
	assert.ErrorContains(t, err, "name must start by 'ick-'")

	// name
	_, err = RegisterComposer("ick-", &BareSnippet{})
	assert.ErrorContains(t, err, "name missing")

	// No Error
	_, err = RegisterComposer("ick-testsnippet1", &HTMLString{})
	assert.NoError(t, err)

	// registered twice
	_, err = RegisterComposer("ick-testsnippet1", &BareSnippet{})
	assert.Error(t, err)
	_, err = RegisterComposer("ick-testsnippet1", &HTMLString{})
	assert.NoError(t, err)
}
