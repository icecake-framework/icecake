package ick0

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQName(t *testing.T) {

	qn := QualifiedName("simple")
	assert.Equal(t, "", qn.Prefix())
	assert.Equal(t, "simple", qn.LocalName())

	qn = QualifiedName("Namespace:name")
	assert.Equal(t, "namespace", qn.Prefix())
	assert.Equal(t, "name", qn.LocalName())

	qn = QualifiedName(" nameSpace : Name ")
	assert.Equal(t, "namespace", qn.Prefix())
	assert.Equal(t, "Name", qn.LocalName())

	qn = QualifiedName("ns1:n1 n2")
	assert.Equal(t, "ns1", qn.Prefix())
	assert.Equal(t, "n1", qn.LocalName())
}
