package ick

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	TheRegistry = registry{}

	c1 := new(HtmlSnippet)
	err := RegisterComposer("snippet", c1)
	assert.ErrorContains(t, err, "not a pointer")

	err = RegisterComposer("snippet", 0)
	assert.ErrorContains(t, err, "must be an HtmlComposer")

	err = RegisterComposer("ick-test-snippet1", testsnippet1{})
	assert.NoError(t, err)

	err = RegisterComposer("ick-test-snippet1", testsnippet1{})
	assert.NoError(t, err) // only log "already registered"

	e := GetRegistryEntry("ick-test-snippet1")
	assert.NotNil(t, e)

	id := GetUniqueId("ick-test-snippet1")
	assert.Equal(t, HTML("ick-test-snippet1-1"), id)
	assert.Equal(t, 1, TheRegistry.entries["ick-test-snippet1"].count)

	r := LookupRegistryEntery(testsnippet1{})
	assert.NotNil(t, r)
}
