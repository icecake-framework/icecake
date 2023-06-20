package html

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttribute(t *testing.T) {
	amap := make(AttributeMap)

	amap.SetAttribute("a", "1")
	assert.Equal(t, AttributeMap{"a": "1"}, amap)

	amap.SetAttribute("a", "2")
	assert.Equal(t, AttributeMap{"a": "1"}, amap)

	amap.SetAttribute("a", "3")
	assert.Equal(t, AttributeMap{"a": "3"}, amap)

	_, err := amap.setAttribute("4", "", false)
	assert.Error(t, err)

	amap.RemoveAttribute("a")
	assert.Equal(t, AttributeMap{}, amap)

	amap.ToggleAttribute("a")
	assert.Equal(t, AttributeMap{"a": ""}, amap)

	amap.ToggleAttribute("a")
	assert.Equal(t, AttributeMap{}, amap)

	_, err = amap.setAttribute("id", "id1.0", false)
	assert.NoError(t, err)

	_, err = amap.setAttribute("id", "1id", false)
	assert.Error(t, err)

	amap.Reset()
	amap.AddClasses("c1", "c2")
	assert.Equal(t, AttributeMap{"class": "c1 c2"}, amap)
	assert.True(t, amap.HasClass("c1") && amap.HasClass("c2"))

	amap.AddClasses(" c3  c4 ")
	assert.Equal(t, AttributeMap{"class": "c1 c2 c3 c4"}, amap)

	amap.AddClasses("  ")
	assert.Equal(t, AttributeMap{"class": "c1 c2 c3 c4"}, amap)

	amap.AddClasses("c4", "c2 c1")
	assert.Equal(t, AttributeMap{"class": "c1 c2 c3 c4"}, amap)

	amap.RemoveClasses("c1")
	assert.Equal(t, AttributeMap{"class": "c2 c3 c4"}, amap)

	amap.RemoveClasses("c1", " c4 c2")
	assert.Equal(t, AttributeMap{"class": "c3"}, amap)

	assert.True(t, amap.Is("class"))
	assert.False(t, amap.Is("disabled"))

	amap.SetBool("disabled", true)
	assert.True(t, amap.Is("disabled"))
}

func TestCheckAttribute(t *testing.T) {
	assert.NoError(t, CheckAttribute("id", "id1.0"))
	assert.Error(t, CheckAttribute("id", "123"))

	assert.NoError(t, CheckAttribute("class", "a"))
	assert.Error(t, CheckAttribute("class", "1"))
	assert.NoError(t, CheckAttribute("class", " a b c "))
	assert.Error(t, CheckAttribute("class", " a 1 c "))

	assert.NoError(t, CheckAttribute("tabindex", "1"))
	assert.NoError(t, CheckAttribute("tabindex", "-1"))
	assert.Error(t, CheckAttribute("tabindex", "a"))

	assert.NoError(t, CheckAttribute("valid", ""))
	assert.Error(t, CheckAttribute("1valid", ""))
	assert.Error(t, CheckAttribute("notvalid#", ""))
}

func TestStringifyAttributeValue(t *testing.T) {

	s, err := StringifyAttributeValue("True")
	assert.NoError(t, err)
	assert.Equal(t, "", s)

	s, err = StringifyAttributeValue("123")
	assert.NoError(t, err)
	assert.Equal(t, "123", s)

	s, err = StringifyAttributeValue("-123")
	assert.NoError(t, err)
	assert.Equal(t, "-123", s)

	s, err = StringifyAttributeValue("-123")
	assert.NoError(t, err)
	assert.Equal(t, "-123", s)

	s, err = StringifyAttributeValue("a")
	assert.NoError(t, err)
	assert.Equal(t, `"a"`, s)

	s, err = StringifyAttributeValue("a'b")
	assert.NoError(t, err)
	assert.Equal(t, `"a'b"`, s)

	s, err = StringifyAttributeValue(`a"b`)
	assert.NoError(t, err)
	assert.Equal(t, `'a"b'`, s)

	_, err = StringifyAttributeValue(`a"b'c`)
	assert.Error(t, err)
}

func TestTryParseAttributes(t *testing.T) {

	// empty
	amap, err := TryParseAttributes("")
	assert.NoError(t, err)
	assert.Zero(t, len(amap))

	// single boolean attributes
	amap, err = TryParseAttributes("attr1")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr1": ""}, amap)

	amap, err = TryParseAttributes(" attr2 ")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr2": ""}, amap)

	// multiple boolean attributes
	amap, err = TryParseAttributes("attr1 attr2 attr3")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr1": "", "attr2": "", "attr3": ""}, amap)

	amap, err = TryParseAttributes("  attr2   attr4   ")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr2": "", "attr4": ""}, amap)

	// single numerical values
	amap, err = TryParseAttributes("attr1=1")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr1": "1"}, amap)

	amap, err = TryParseAttributes("  attr2  =  2  ")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr2": "2"}, amap)

	// multiple numerical values
	amap, err = TryParseAttributes("attr1=1 attr2=2 attr3=3")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr1": "1", "attr2": "2", "attr3": "3"}, amap)

	amap, err = TryParseAttributes("attr1=10   attr3  =  30     attr4=   40  ")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attr1": "10", "attr3": "30", "attr4": "40"}, amap)

	// single alpha values
	amap, err = TryParseAttributes("attrA='a'")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "a"}, amap)

	amap, err = TryParseAttributes(`attrA="A"`)
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "A"}, amap)

	amap, err = TryParseAttributes(" attrA  =  '  AA  ' ")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "  AA  "}, amap)

	amap, err = TryParseAttributes(`attrA="1"`)
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "1"}, amap)

	amap, err = TryParseAttributes(`attrA=" 1 "`)
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": " 1 "}, amap)

	// multiple alpha values
	amap, err = TryParseAttributes("attrA='a' attrB='b'")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "a", "attrB": "b"}, amap)

	amap, err = TryParseAttributes(" attrA  =  '  a  '   attrB =  '  b  '")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "  a  ", "attrB": "  b  "}, amap)

	// quoted alpha values
	amap, err = TryParseAttributes("attrA='  a\"  '")
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrA": "  a\"  "}, amap)

	amap, err = TryParseAttributes(`attrB="  b'  "`)
	assert.NoError(t, err)
	assert.Equal(t, AttributeMap{"attrB": `  b'  `}, amap)

}

func TestString(t *testing.T) {

	amap := make(AttributeMap)
	assert.Equal(t, "", amap.String())
	assert.Equal(t, "a", amap.SetAttribute("A", "").String())
	assert.Equal(t, "a=1", amap.SetAttribute("A", "1").String())
	assert.Equal(t, `a=1 b="b"`, amap.SetAttribute("B", "b").String())
	assert.Equal(t, `a=1 b="b" c="a'b'c"`, amap.SetAttribute("C", `a'b'c`).String())
	assert.Equal(t, `a=1 b="b" c="a'b'c" d='a"b"c'`, amap.SetAttribute("D", `a"b"c`).String())

	amap.Reset()
	assert.Equal(t, `id="ID1" class="c1 c2"`, amap.SetId("ID1").AddClasses("c1 c2").String())
	assert.Equal(t, `id="ID1" class="c1 c2" disabled`, amap.SetDisabled(true).String())
}
