package framework

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/webclientsdk"
)

func RenderComponents() {
	// rechercher

	fmt.Println("Component rendering.")

	webclientsdk.GetDocument().GetElementById("id124").SetInnerHTML("This is my comp")

}
