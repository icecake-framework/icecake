package icecake

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"text/template"

	"github.com/sunraylab/icecake/pkg/dom"
)

/*****************************************************************************/

// unfoldComponents lookup for component tags in htmlstring, and render each of them recursively.
//
// rendering means:
//  1. if the component does not have an ID yet, then create one and instantiate it
//  2. parse component's template, according to go html templating standards
//  3. execute this template with {{}} langage and component's data and global data
//
// NOTICE: to avoid infinite recursivity, the rendering fails at a the 10th depth
func unfoldComponents(name string, _unsafeHtmlTemplate string, data any, _deep int) (_rendered string, _err error) {
	if _deep >= 10 {
		strerr := fmt.Sprintf("RenderComponents stop at level %d: recursive rendering too deep", _deep)
		dom.ConsoleError(strerr)
		return "", fmt.Errorf(strerr)
	}
	cmpid := name + "-" + strconv.Itoa(_deep)
	fmt.Printf("unfolding %d:%q\n", _deep, name)

	// 1. parse
	tmpCmp := template.Must(template.New(name).Parse(_unsafeHtmlTemplate))

	// 2. execute
	bufCmp := new(bytes.Buffer)
	errTmp := tmpCmp.Execute(bufCmp, data)
	if errTmp != nil {
		strerr := fmt.Sprintf("unfoldComponents stop at level %d: %q ERROR applying data to template: '%v'", _deep, name, _unsafeHtmlTemplate)
		dom.ConsoleError(strerr)
		return "", fmt.Errorf(strerr)
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
				strerr := fmt.Sprintf("unfoldComponents stop at level %d: close delim not found", _deep)
				dom.ConsoleError(strerr)
				return "", fmt.Errorf(strerr)

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
						fmt.Printf("component:%s level:%d instantiating %s with name:%s", name, _deep, c.Type(), newcmpid)

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
							str, _err = unfoldComponents(newcmpid, t.InnerHtmlTemplate(), d, _deep+1)
						}

						// let's go deeper
						out.WriteString(fmt.Sprintf(`<span id="%s">%s</span>`, newcmpid, str))

					} else {
						// the tag is empty or is not a registered component
						out.WriteString(fmt.Sprintf("<!-- component ic-%q not registered -->", tagName))
						fmt.Printf("component not registered")
					}
				}
			}
		}
	}

}
