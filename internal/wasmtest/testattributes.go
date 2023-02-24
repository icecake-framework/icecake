package main

import (
	"fmt"

	ick "github.com/sunraylab/icecake/pkg/icecake"
)

func TestAttributes() {
	attrs, _ := ick.ParseAttributes("single")
	fmt.Println(attrs.String()) // single
	attrs, _ = ick.ParseAttributes("one two")
	fmt.Println(attrs.String()) // one two
	attrs, _ = ick.ParseAttributes("zero one=1 two three=3 four five six")
	fmt.Println(attrs.String()) // zero one="1" two three="3" four five six
	attrs, _ = ick.ParseAttributes("one='one' two='two'")
	fmt.Println(attrs.String()) // one="one" two="two"
	attrs, _ = ick.ParseAttributes("  this    =   'with \"quoted sub value\"' anotherone ")
	fmt.Println(attrs.String()) // anotherone this='with "quoted sub value"'

	var err error
	_, err = ick.ParseAttributes("one t#o three")
	if err != nil {
		fmt.Println(err) // attribute name is not valid : "t#o"
	}
}
