package icecakeapp

import (
	"fmt"

	"github.com/sunraylab/icecake/pkg/framework/button"
	"github.com/sunraylab/icecake/pkg/webclientsdk"
)

func RenderComponents() {
	// rechercher

	fmt.Println("Component rendering.")

	coll := webclientsdk.GetDocument().GetElementsByTagName("ic-button")
	if coll != nil {
		for i := uint(0); i < coll.Length(); i++ {
			e := coll.Item(i)
			icb := button.Cast(e.JSValue())
			icb.Render()
		}
	}

}
