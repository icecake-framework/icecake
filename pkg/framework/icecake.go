package icecake

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"reflect"
	"strconv"
	"strings"
	// "github.com/sunraylab/icecake/pkg/dom"
)

/*****************************************************************************/

// func RenderElement(htmlTemplate string, data any, elem *dom.Element) {
// 	str, _ := renderComponents("App", htmlTemplate, data, 0)
// 	elem.SetInnerHTML(str)

// ajoute tous les listeners

// }

/*****************************************************************************/

// scan the Docuemnts's Body to look for IC components and to render them.
// func renderBody() {
// 	doc := dom.GetDocument()

// 	// check doc type & validity
// 	if len(ComponentTypes) == 0 {
// 		log.Println("renderBody: no customized IC component type registered")
// 		return
// 	}
// 	htmlbody := doc.Body().InnerHTML()
// 	renderComponents("body", htmlbody, nil, 0)
// }

type tree struct {
	root  *Compounder
	folds map[string]*Compounder
}

type idtree map[string]Compounder

func renderElement(name string, _htmlstring string) (_tree tree, _rendered string, _err error) {

	_rendered, _err = renderComponents(name, _htmlstring, GData, 0)

	//tree := parse.New("after", idtree)

	return _rendered, _err
}

// renderComponents lookup for component tags in htmlstring, and render each of them recursively.
//
// rendering means:
//  1. if the component does not have an ID yet, then create one and instantiate it
//  2. parse component's template, according to go html templating standards
//  3. execute this template with {{}} langage and component's data and global data
//     3.
//
// NOTICE: to avoid infinit recursivity, the rendering fails at a the 10th depth
func renderComponents(name string, _htmlstring string, data any, _deep int) (_rendered string, _err error) {
	if _deep >= 10 {
		return "", fmt.Errorf("recursive rendering too deep")
	}
	cmpid := name + "-" + strconv.Itoa(_deep)
	log.Printf("component:%s level:%d rendering...\n", name, _deep)

	// 1. parse
	tmpCmp := template.Must(template.New(name).Parse(_htmlstring))

	// 2. execute
	bufCmp := new(bytes.Buffer)
	errTmp := tmpCmp.Execute(bufCmp, data)
	if errTmp != nil {
		log.Printf("component:%s level:%d ERROR applying data to template: '%v'", name, _deep, _htmlstring)
		return "", _err
	}
	htmlstring := bufCmp.String()

	// 3. lookup for components
	const (
		delim_open  = "<ic-"
		delim_close = "/>"
	)

	n := 0
	out := &bytes.Buffer{}
	for {
		// find next delim_open
		if from := strings.Index(htmlstring, delim_open); from == -1 || _err != nil {
			// no more delim_open = feed the output with what's left in htmlstring and return
			out.WriteString(htmlstring)
			return out.String(), _err
		} else {
			// we found a new delim_open
			// so it's time to feed the output with data preceding this delim
			out.WriteString(htmlstring[:from])

			// scrap this data and keep what's left
			htmlstring = htmlstring[from+len(delim_open):]

			// look now for it's corresponding delim_close
			if to := strings.Index(htmlstring, delim_close); to == -1 {
				// not corresponding delim_close then stop and return a rendering error
				return "", fmt.Errorf("close delim not found")

			} else {
				// we got a delim_close so we've an opening element
				// extract the element's content
				inside := htmlstring[0:to]

				// scrap it and keep what's left
				htmlstring = htmlstring[to+len(delim_close):]

				// split this opening element in fields
				m := strings.Fields(inside)

				// first field is the tagName, does nothing if empty
				var tagName string
				if len(m) > 0 {
					tagName = m[0]

					// is this tag a registered component ?
					if comptype, found := ComponentTypes["ic-"+tagName]; found {
						newcmpid := cmpid + "-" + tagName + "-" + strconv.Itoa(n)

						// does it have an id ?
						// Attr := m[1:]

						var str string

						c := reflect.New(comptype)
						log.Printf("component:%s level:%d instantiating %s with name:%s", name, _deep, c.Type(), newcmpid)

						switch t := c.Interface().(type) {
						case CompoundBuilder:
							t.Mount()
						}

						switch t := c.Interface().(type) {
						case Compounder:
							type DATA struct {
								Me     any
								Owner  *any
								Global *map[string]any
							}
							d := DATA{
								Me:     t,
								Owner:  &data,
								Global: &GData,
							}
							str, _err = renderComponents(newcmpid, t.InnerHtmlTemplate(), d, _deep+1)
						}

						// let's go deeper
						out.WriteString(fmt.Sprintf(`<span id="%s">%s</span>`, newcmpid, str))

					} else {
						// the tag is empty or is not a registered component
						out.WriteString(fmt.Sprintf("<!-- component ic-%q not registered -->", tagName))
						log.Printf("component not registered")
					}
				}
			}
		}
	}

}

/*****************************************************************************/
