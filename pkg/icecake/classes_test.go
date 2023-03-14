package ick

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClasses(t *testing.T) {

	var cs1, cs2 *Classes
	var err error

	cs1, err = ParseClasses(" c2   c3 c1 ")
	assert.NoError(t, err)
	assert.Equal(t, "c2 c3 c1", cs1.String())

	assert.True(t, cs1.Has("c3"))

	assert.Equal(t, "c2", cs1.At(0))

	cs2, _ = ParseClasses("d1 d2")
	assert.Equal(t, 5, cs1.AddClasses(*cs2).Count())

	assert.True(t, cs2.RemoveTokens("c1").Toggle("c1"))
	assert.False(t, cs2.Toggle("c1"))
}
