package icecake

import (
	"log"
	"testing"
)

func TestRenderComponent(t *testing.T) {

	data := struct{ Name string }{
		Name: "Bob",
	}
	out, err := RenderComponents("example0", `html0 Hello <strong>{{.Name}}</strong>!`, data, 0)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println("------>", out)

	GData["name"] = "Alice"
	out, err = renderElement("example00", `html00 Hello <strong>{{.name}}</strong>!`, GData)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println("------>", out)

	// out, err = renderComponents("example1", `html1 : <ic-ex1 />`, data, 0)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// }
	// log.Println("------>", out)

	// out, err = renderComponents("example2", `html2 : <ic-ex2 />`, data, 0)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// }
	// log.Println("------>", out)

	// out, err = renderComponents("example3", `html3 : <ic-ex3 />`, data, 0)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// }
	// log.Println("------>", out)

	// out, err = renderComponents("example4", `hmlt4 : <ic-ex4 />`, data, 0)
	// if err != nil {
	// 	t.Errorf(err.Error())
	// }
	// log.Println("------>", out)

	out, err = renderElement("example5", `hmlt5 <ic-ex5 />`, GData)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println("------>", out)

	out, err = renderElement("example6", `hmlt6 <ic-ex6 />`, GData)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println("------>", out)

}
