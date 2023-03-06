package main

import (
	"fmt"

	ick "github.com/sunraylab/icecake/pkg/icecake"
)

func TestJS() {

	body := ick.NewWebApp().Body()

	i := body.GetInt("tagName")
	fmt.Println(i)
}
