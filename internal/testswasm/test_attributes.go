package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	ick "github.com/sunraylab/icecake/pkg/icecake"
)

func TestAttributes(t *testing.T) {
	// cache
	var attrs ick.Attributes
	attrs.ParseAttributes("single")
	require.Equal(t, "single", attrs.String())

	attrs.ParseAttributes("one two")
	require.Equal(t, "one two", attrs.String())

	attrs.ParseAttributes("zero one=1 two three=3 four five six")
	require.Equal(t, `five four one='1' six three='3' two zero`, attrs.String())

	attrs.ParseAttributes("one='one' two='two'")
	require.Equal(t, `one='one' two='two'`, attrs.String())

	attrs.ParseAttributes("  this    =   'with \"quoted sub value\"' anotherone ")
	require.Equal(t, `anotherone this='with "quoted sub value"'`, attrs.String())

	var err error
	err = attrs.ParseAttributes("one t#o three")
	require.Error(t, err)
}
