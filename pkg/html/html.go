package html

import (
	"fmt"
	"io"
	"net/url"
	"reflect"
)

// HTMLString encapsulates a known safe string document fragment.
// It should not be used for string from a third-party, or string with
// unclosed tags or comments.
//
// Use of this type presents a security risk:
// the encapsulated content should come from a trusted source,
// as it will be included verbatim in the output.
type HTMLString string

// WriteHTML write the HTMLString to out
func (strhtml HTMLString) WriteHTML(out io.Writer) {
	io.WriteString(out, string(strhtml))
}

// WriteHTML write the HTMLString to out
func (strhtml HTMLString) String() string {
	return string(strhtml)
}

type DataState struct {
	//Id   string // the id of the current processing component
	//Me   any    // the current processing component, should embedd an HtmlSnippet
	Page any // the current ick page, can be nil
	App  any // the current ick App, can be nil
}

// WriteStringsIf writes one or many strings to w only if the condition is true.
// errors comes from the writer.
func WriteStringsIf(condition bool, w io.Writer, ss ...string) (n int, err error) {
	if !condition {
		return 0, nil
	}
	return WriteStrings(w, ss...)
}

// WriteStrings writes one or many strings to w
// Returns the number of bytes written.
// errors comes from the writer.
func WriteStrings(w io.Writer, ss ...string) (n int, err error) {
	nn := 0
	for _, s := range ss {
		nn, err = WriteString(w, s)
		if err != nil {
			return
		}
		n += nn
	}
	return
}

// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
// If w implements StringWriter, its WriteString method is invoked directly.
// Otherwise, w.Write is called exactly once.
// // errors comes from the writer.
func WriteString(out io.Writer, s string) (n int, err error) {
	return io.WriteString(out, s)
}

// mini helper
func mini(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func debugValue(_v reflect.Value) {
	fmt.Printf("Type: %s\n", _v.Type().String())

	n := _v.Type().NumMethod()
	fmt.Printf("Nb Method: %v\n", n)
	for i := 0; i < n; i++ {
		m := _v.Method(i)
		name := _v.Type().Method(i).Name
		fmt.Printf("Method %v: %s %s '%v'\n", i, name, m.String(), m)
	}

	n = _v.NumField()
	fmt.Printf("Nb Field: %v\n", n)
	for i := 0; i < n; i++ {
		m := _v.Field(i)
		name := _v.Type().Field(i).Name
		fmt.Printf("Field %v: %v %v '%v'\n", i, name, m.Type().String(), m)
	}
}

func debugAny(_v any) {
	fmt.Printf("Type: %v\n", reflect.TypeOf(_v).String())
	fmt.Printf("Type: %v\n", reflect.ValueOf(_v).Interface())

	_, ok := _v.(*url.URL)
	fmt.Printf("Type url.URL: %v\n", ok)

}
