package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {

	ResetRegistry()

	e := GetRegistryEntry("ick-test-snippet1")
	assert.NotNil(t, e)

	id := GetUniqueId("ick-test-snippet1")
	assert.Equal(t, "ick-test-snippet1-1", id)
	assert.Equal(t, 1, theRegistry.entries["ick-test-snippet1"].count)
}
