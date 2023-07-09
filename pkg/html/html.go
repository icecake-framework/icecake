package html

import (
	"io"
)

// RenderStringIf writes one or many strings to w only if the condition is true.
// Returns the number of bytes written and errors from the writer.
func RenderStringIf(condition bool, w io.Writer, ss ...string) (n int, err error) {
	if !condition {
		return 0, nil
	}
	return RenderString(w, ss...)
}

// RenderString writes one or many strings to w.
// Returns the number of bytes written and errors from the writer.
func RenderString(w io.Writer, ss ...string) (n int, err error) {
	nn := 0
	for _, s := range ss {
		nn, err = io.WriteString(w, s)
		if err != nil {
			return
		}
		n += nn
	}
	return
}

// mini helper
func mini(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// func debugValue(_v reflect.Value) {
// 	fmt.Printf("Type: %s\n", _v.Type().String())

// 	n := _v.Type().NumMethod()
// 	fmt.Printf("Nb Method: %v\n", n)
// 	for i := 0; i < n; i++ {
// 		m := _v.Method(i)
// 		name := _v.Type().Method(i).Name
// 		fmt.Printf("Method %v: %s %s '%v'\n", i, name, m.String(), m)
// 	}

// 	n = _v.NumField()
// 	fmt.Printf("Nb Field: %v\n", n)
// 	for i := 0; i < n; i++ {
// 		m := _v.Field(i)
// 		name := _v.Type().Field(i).Name
// 		fmt.Printf("Field %v: %v %v '%v'\n", i, name, m.Type().String(), m)
// 	}
// }

// func debugAny(_v any) {
// 	fmt.Printf("Type: %v\n", reflect.TypeOf(_v).String())
// 	fmt.Printf("Type: %v\n", reflect.ValueOf(_v).Interface())

// 	_, ok := _v.(*url.URL)
// 	fmt.Printf("Type url.URL: %v\n", ok)

// }
