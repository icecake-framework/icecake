package html

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCoreHtmlFile(t *testing.T) {

	dft := NewHtmlFile("en")
	out := new(bytes.Buffer)
	err := dft.Render(out)
	require.NoError(t, err)
	assert.Equal(t, `<!doctype html><html lang="en"><head></head></html>`, out.String())
}