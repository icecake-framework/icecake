package ick

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCoreHtmlFile(t *testing.T) {

	dft := NewPage("en", "")
	out := new(bytes.Buffer)
	err := dft.RenderContent(out)
	require.NoError(t, err)
	assert.Equal(t, `<!doctype html><html lang="en"><head></head><body></body></html>`, out.String())
}
