package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {

}

func TestStringifyAttributeValue(t *testing.T) {

	//StringifyAttributeValue()

}

func TestParseAttributes(t *testing.T) {

	// empty
	amap, err := ParseAttributes("")
	assert.NoError(t, err)
	assert.Zero(t, len(amap))

	// single boolean attributes
	amap, err = ParseAttributes("attr1")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr1": ""}, amap)

	amap, err = ParseAttributes(" attr2 ")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr2": ""}, amap)

	// multiple boolean attributes
	amap, err = ParseAttributes("attr1 attr2 attr3")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr1": "", "attr2": "", "attr3": ""}, amap)

	amap, err = ParseAttributes("  attr2   attr4   ")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr2": "", "attr4": ""}, amap)

	// single numerical values
	amap, err = ParseAttributes("attr1=1")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr1": "1"}, amap)

	amap, err = ParseAttributes("  attr2  =  2  ")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr2": "2"}, amap)

	// multiple numerical values
	amap, err = ParseAttributes("attr1=1 attr2=2 attr3=3")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr1": "1", "attr2": "2", "attr3": "3"}, amap)

	amap, err = ParseAttributes("attr1=10   attr3  =  30     attr4=   40  ")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr1": "10", "attr3": "30", "attr4": "40"}, amap)

	// single alpha values
	amap, err = ParseAttributes("attrA='a'")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "a"}, amap)

	amap, err = ParseAttributes(`attrA="A"`)
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "A"}, amap)

	amap, err = ParseAttributes(" attrA  =  '  AA  ' ")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "  AA  "}, amap)

	amap, err = ParseAttributes(`attrA="1"`)
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "1"}, amap)

	amap, err = ParseAttributes(`attrA=" 1 "`)
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": " 1 "}, amap)

	// multiple alpha values
	amap, err = ParseAttributes("attrA='a' attrB='b'")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "a", "attrB": "b"}, amap)

	amap, err = ParseAttributes(" attrA  =  '  a  '   attrB =  '  b  '")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "  a  ", "attrB": "  b  "}, amap)

	// quoted alpha values
	amap, err = ParseAttributes("attrA='  a\"  '")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "  a\"  "}, amap)

	amap, err = ParseAttributes(`attrB="  b'  "`)
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrB": `  b'  `}, amap)

}
