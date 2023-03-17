package ick

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAttributes1(t *testing.T) {
	// cache
	var as1, as2 *Attributes
	var err error

	as1, err = ParseAttributes("single")
	require.NoError(t, err)
	require.Equal(t, "single", as1.StringUnquoted())

	as1, err = ParseAttributes("one two")
	require.NoError(t, err)
	require.Equal(t, "one two", as1.StringUnquoted())

	as2, err = ParseAttributes("zero=0 bool=False one=1 two three=3 four five six")
	require.NoError(t, err)
	require.Equal(t, `bool=false five four one=1 six three=3 two zero=0`, as2.StringUnquoted())

	as1, err = ParseAttributes("one='one' two='two'")
	require.NoError(t, err)
	require.Equal(t, `one=one two=two`, as1.StringUnquoted())

	as1, err = ParseAttributes("  this    =   'with \"quoted sub value\"' anotherone ")
	require.NoError(t, err)
	require.Equal(t, `anotherone this='with "quoted sub value"'`, as1.StringUnquoted())

	as1, err = ParseAttributes(`a1="o'connor"`)
	require.NoError(t, err)
	require.Equal(t, `a1=o'connor`, as1.StringUnquoted())

	as1, err = ParseAttributes(`a1=" o ' connor "`)
	require.NoError(t, err)
	require.Equal(t, `a1=" o ' connor "`, as1.StringUnquoted())

	as1, err = ParseAttributes(`a1="o 'connor"`)
	require.NoError(t, err)
	require.Equal(t, `a1="o 'connor"`, as1.StringUnquoted())

	v, _ := as1.Attribute("a1")
	require.Equal(t, StringQuotes(`o 'connor`), v)

	as1.SetAttributes(*as2, false)
	require.Equal(t, 9, len(as1.Keys()))

	require.True(t, as1.IsTrue("one"))
	require.False(t, as1.IsTrue("ten"))
	require.False(t, as1.IsTrue("zero"))
	require.False(t, as1.IsTrue("bool"))
	require.False(t, as1.Hidden())

	require.Equal(t, 0, as1.TabIndex())
	require.Equal(t, 2, as1.SetTabIndex(2).TabIndex())

	_, f := as1.RemoveAttribute("bool").Attribute("bool")
	require.False(t, f)

	require.True(t, as1.Toggle("bool"))
	require.False(t, as1.Toggle("bool"))

	_, err = ParseAttributes("one t#o three")
	require.Error(t, err)

	as3, err := ParseAttributes("data-a data-s='ok' data-v=10")
	require.NoError(t, err)
	as2.SetAttributes(*as3, false)
	require.Equal(t, "data-a data-s=ok data-v=10", as2.Data().StringUnquoted())

}

func TestAttributes2(t *testing.T) {

	as, err := ParseAttributes("a='<br/>'")
	require.NoError(t, err)
	require.Equal(t, "a=<br/>", as.StringUnquoted())

	as, err = ParseAttributes(`a='<br/>' b="<br/>"`)
	require.NoError(t, err)
	require.Equal(t, "a=<br/> b=<br/>", as.StringUnquoted())

	as, err = ParseAttributes("a=<br/> b=<br/>")
	require.NoError(t, err)
	require.Equal(t, "a=<br/> b=<br/>", as.StringUnquoted())

	as, err = ParseAttributes("a='< < <> </> > />'")
	require.NoError(t, err)
	require.Equal(t, "a='< < <> </> > />'", as.StringUnquoted())

	_, err = ParseAttributes("a b=0 ><something else")
	require.Error(t, err)

	_, err = ParseAttributes("a/><something else")
	require.Error(t, err)
}

func TestAttributes3(t *testing.T) {

	as := Attributes{}
	as.SetAttribute("a", "1")
	require.Equal(t, "a=1", as.StringUnquoted())

	as.Clear()
	as.SetAttribute("a", "x")
	require.Equal(t, "a=x", as.StringUnquoted())

	as.Clear()
	as.SetAttribute("a", "'x'")
	require.Equal(t, "a='x'", as.StringUnquoted())

	as.Clear()
	as.ParseAttribute("a", "'x'")
	require.Equal(t, "a=x", as.StringUnquoted())

	as.Clear()
	as.SetAttribute("a", "' x '")
	require.Equal(t, `a="' x '"`, as.StringUnquoted())

	as.Clear()
	as.SetAttribute("a", " x ")
	require.Equal(t, `a=' x '`, as.StringUnquoted())
}

func TestAttributes4(t *testing.T) {

	s := ParseStringQuotes("hello")
	require.Equal(t, "hello", string(s))

	s = ParseStringQuotes(" hello ")
	require.Equal(t, "hello", string(s))

	s = ParseStringQuotes(" ' hello ' ")
	require.Equal(t, " hello ", string(s))

	s = ParseStringQuotes(` " hello " `)
	require.Equal(t, " hello ", string(s))
}
