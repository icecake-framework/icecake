package ick

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunraylab/icecake/pkg/html"
	"github.com/sunraylab/icecake/pkg/registry"
)

type testsnippet1 struct {
	html.HTMLSnippet
	Html html.String
}

func (tst testsnippet1) Template(*html.DataState) (_t html.SnippetTemplate) {
	_t.Body = tst.Html
	return
}

func TestRegisterB(t *testing.T) {

	registry.ResetRegistry()

	r := registry.LookupRegistryEntry(&testsnippet1{})
	assert.NotNil(t, r)

	r = registry.LookupRegistryEntry(testsnippet1{})
	assert.Nil(t, r)

	c1 := new(html.HTMLSnippet)
	err := RegisterComposer("snippet", *c1)
	assert.ErrorContains(t, err, "not a component")

	i := new(int)
	err = RegisterComposer("snippet", i)
	assert.ErrorContains(t, err, "must be an HtmlComposer")

	err = RegisterComposer("ick-test-snippet1", &testsnippet1{})
	assert.NoError(t, err)

	err = RegisterComposer("ick-test-snippet1", &testsnippet1{})
	assert.NoError(t, err) // only log "already registered"
}
