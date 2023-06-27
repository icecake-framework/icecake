package helper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPath(t *testing.T) {

	p, err := CheckOutputPath(nil)
	assert.NoError(t, err)
	assert.Empty(t, p)

	pathtemp := "./tempick/temp/"
	err = os.MkdirAll(pathtemp, os.ModePerm)
	assert.NoError(t, err, "unable to create temp dir")

	p, err = CheckOutputPath(&pathtemp)
	assert.NoError(t, err)
	assert.NotEmpty(t, p)

	err = os.RemoveAll("./tempick/")
	assert.NoError(t, err, "unable to remove temp dir")

	_, err = CheckOutputPath(&pathtemp)
	assert.Error(t, err)
}
