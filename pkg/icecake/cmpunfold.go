package ick

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/sunraylab/icecake/pkg/errors"
)

type TemplateData struct {
	Me any
	// Page
	App any
}

// unfoldComponents lookup for component tags in htmlstring, and render each of them recursively.
//
// rendering means:
//  1. if the component does not have an ID yet, then create one and instantiate it
//  2. parse component's template, according to go html templating standards
//  3. execute this template with {{}} langage and component's data and global data
//
// NOTICE: to avoid infinite recursivity, the rendering fails at a the 10th depth
func unfoldComponents(_unfoldedCmps map[string]HtmlListener, name string, _unsafeHtmlTemplate string, _data any, _deep int) (_rendered string, _err error) {
	if _deep >= 10 {
		return "", errors.ConsoleErrorf("RenderComponents stop at level %d: recursive rendering too deep", _deep)
	}
	//cmpid := name + "-" + strconv.Itoa(_deep)
	fmt.Printf("unfolding %d:%q\n", _deep, name)

	// 1. parse
	tmpCmp := template.Must(template.New(name).Parse(_unsafeHtmlTemplate))

	// 2. execute
	bufCmp := new(bytes.Buffer)
	errTmp := tmpCmp.Execute(bufCmp, _data)
	if errTmp != nil {
		return "", errors.ConsoleErrorf("unfoldComponents stop at level %d: %q ERROR applying data to template: '%v'", _deep, name, _unsafeHtmlTemplate)
	}
	htmlstring := bufCmp.String()

	// 3. lookup for components
	const (
		delim_open  = "<ick-"
		delim_close = "/>"
	)

	n := 0
	out := &bytes.Buffer{}
nextdelim:
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
				return "", errors.ConsoleErrorf("unfoldComponents stop at level %d: close delim not found", _deep)

			} else {

				// we got a delim_close so we've a new ick element, extract its content
				inside := htmlstring[0:to]
				tagname, leftinside, _ := strings.Cut(inside, " ")
				htmlstring = htmlstring[to+len(delim_close):] // scrap it and keep what's left

				if tagname == "" { // <ick-/> !
					continue nextdelim
				}

				// DEBUG:
				//fmt.Println("embedded ick component:'", tagname, "' leftinside:", leftinside)

				// does this tag refer to a registered component ?
				if comptype, found := GComponentRegistry["ick-"+tagname]; found {

					// Instantiate the component and add it to the unfoldied stack
					newcmpreflect := reflect.New(comptype)
					newcmp := newcmpreflect.Interface().(Composer)
					newcmpid := fmt.Sprintf("ick-%s-%d-%d", tagname, _deep, n)
					_unfoldedCmps[newcmpid] = newcmpreflect.Interface().(HtmlListener) // newcmp

					// DEBUG:
					fmt.Printf("instantiating %s of type %s\n", newcmpid, newcmpreflect.Type())

					var newcmpelem *Element
					if newcmpelem, _err = CreateComponentElement(newcmp); _err == nil {

						// process component's attributes
						var attrs *Attributes
						attrs, _err = ParseAttributes(leftinside)
						if _err != nil {
							continue nextdelim
						}

						// set the id, overwrite the one in the template
						attrs.Set("id", newcmpid)

						anames := attrs.Sort()
						// DEBUG:
						//fmt.Println(anames, " => ", attrs)

						for _, aname := range anames {
							_, found := newcmpreflect.Elem().Type().FieldByName(aname)
							if !found {
								// DEBUG:
								// fmt.Printf("attribute %v: %q is kept asis\n", i, aname)
								// keep it as is
								newcmpelem.SetAttribute(aname, attrs.Get(aname))
							} else {
								// DEBUG:
								// fmt.Printf("attribute %v: %q corresponds to a component's data\n", i, aname)
								// feed data struct with the value
								fieldvalue := newcmpreflect.Elem().FieldByName(aname)
								switch fieldvalue.Kind() {
								case reflect.String:
									fieldvalue.SetString(attrs.Get(aname))
								default:
									errors.ConsoleWarnf("Unmanaged type for attribute %q of component %q", aname, newcmpid)
								}
							}
						}

						// recursively unfold the component template
						data := TemplateData{
							Me:  newcmp,
							App: _data,
						}
						var htmlin string
						htmlin, _err = unfoldComponents(_unfoldedCmps, newcmpid, newcmp.Template(), data, _deep+1)
						newcmpelem.SetInnerHTML(htmlin)
						htmlout := newcmpelem.OuterHTML()

						// TODO add style

						// let's go deeper
						out.WriteString(htmlout)
					}

				} else {
					// the tag is not a registered component
					htmlmsg := fmt.Sprintf("<!-- unable to unfold unregistered component ick-%s -->", tagname)
					out.WriteString(htmlmsg)
					errors.ConsoleWarnf(htmlmsg)
				}
			}
		}
	}
}

// addComponentslisteners call addlisteners for every unfolded Components
func addComponentslisteners(_unfoldedCmps map[string]HtmlListener) {
	for id, ufc := range _unfoldedCmps {
		e := GetDocument().ChildById(id)
		ufc.Wrap(e)
		ufc.AddListeners()
	}
}
