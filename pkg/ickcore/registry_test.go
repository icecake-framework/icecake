package ickcore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistry(t *testing.T) {

	ResetRegistry()

	e := GetRegistryEntry("ickTEST")
	assert.NotNil(t, e)

	_, id := GetUniqueId("ickTEST")
	assert.Equal(t, "icktest-1", id)
	assert.Equal(t, 1, theRegistry.entries["icktest"].count)
}
