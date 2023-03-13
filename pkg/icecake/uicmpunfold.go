package ick

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/sunraylab/icecake/pkg/errors"
)

type TemplateData struct {
	Id  string // the id of the processing component
	Me  any    // the processing component
	App any    // the App object, can be nil
	// Page
}

// unfoldComponents lookup for component tags in htmlstring, and render each of them recursively.
//
// rendering means:
//  1. if the component does not have an ID yet, then create one and instantiate it
//  2. parse component's template, according to go html templating standards
//  3. execute this template with {{}} langage and component's data and global data
//
// NOTICE: to avoid infinite recursivity, the rendering fails at a the 10th depth
func unfoldComponents(_unfoldedCmps map[string]Composer, name string, _unsafeHtmlTemplate string, _data any, _deep int) (_rendered string, _err error) {
	if _deep >= 10 {
		return "", errors.ConsoleErrorf("unfoldComponents stopped at level %d. Recursive rendering too deep", _deep)
	}
	//cmpid := name + "-" + strconv.Itoa(_deep)
	errors.ConsoleLogf("unfolding %d:%q\n", _deep, name)

	// 1. parse
	tmpCmp := template.Must(template.New(name).Parse(_unsafeHtmlTemplate))

	// 2. execute
	bufCmp := new(bytes.Buffer)
	errTmp := tmpCmp.Execute(bufCmp, _data)
	if errTmp != nil {
		return "", errors.ConsoleErrorf("unfoldComponents stopped at level %d. %q ERROR applying data to %s", _deep, name, errTmp.Error())
	}
	htmlstring := bufCmp.String()

	// 3. lookup for components
	const (
		delim_open  = "<ick-"
		delim_close = "/>"
	)

	//n := 0
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
				return "", errors.ConsoleErrorf("unfoldComponents stopped at level %d. Close delim not found", _deep)

			} else {

				// we got a delim_close so we've a new ick-element, extract its content
				inside := htmlstring[0:to]
				tagname, leftinside, _ := strings.Cut(inside, " ")
				htmlstring = htmlstring[to+len(delim_close):] // scrap it and keep what's left

				if tagname == "" { // <ick-/> !
					continue nextdelim
				}

				// DEBUG:
				//fmt.Println("embedded ick component:'", tagname, "' leftinside:", leftinside)

				// does this tag refer to a registered component ?
				if regentry, found := App.CmpRegistry["ick-"+tagname]; found {

					// process and instantiate new component

					// Instantiate the component
					newcmpreflect := reflect.New(regentry.typ)
					newcmp := newcmpreflect.Interface().(Composer)
					//newcmpid := fmt.Sprintf("ick-%s-%d-%d", tagname, _deep, n)

					var newcmpelem *UIComponent
					var newcmpid string
					if newcmpid, newcmpelem, _err = App.CreateComponent(newcmp); _err == nil {

						errors.ConsoleLogf("unfoldComponents instantiating %s of type %s\n", newcmpid, newcmpreflect.Type())

						// add the component to the add it to the unfolded stack
						_unfoldedCmps[newcmpid] = newcmpreflect.Interface().(Composer) // newcmp

						// process embeded component's attributes
						var attrs *Attributes
						attrs, _err = ParseAttributes(leftinside)
						if _err != nil {
							continue nextdelim
						}

						// set the id, overwrite the one in the template if any
						attrs.SetAttribute("id", newcmpid)

						anames := attrs.Sort()
						// DEBUG:
						//fmt.Println(anames, " => ", attrs)

						for _, aname := range anames {
							_, found := newcmpreflect.Elem().Type().FieldByName(aname)
							if !found {
								// DEBUG:
								// fmt.Printf("attribute %v: %q is kept asis\n", i, aname)

								// this attribute is not a field of the componenent
								// keep it as is unless it is the class attribute, in this case, add the tokens
								aval, _ := attrs.Attribute(aname)
								if aname == "class" {
									newcmpelem.Classes().AddTokens(aval)
								} else {
									newcmpelem.SetAttribute(aname, aval)
								}
							} else {
								// DEBUG:
								// fmt.Printf("attribute %v: %q corresponds to a component's data\n", i, aname)

								// feed data struct with the value
								strav, _ := attrs.Attribute(aname)

								fieldvalue := newcmpreflect.Elem().FieldByName(aname)
								// DEBUG: fmt.Println("DEBUG:", fieldvalue.Type().String())
								var erra error
								switch fieldvalue.Type().String() {
								case "time.Duration":
									var d time.Duration
									d, erra = time.ParseDuration(strav)
									if erra == nil {
										fieldvalue.SetInt(int64(d))
									}
								default:
									switch fieldvalue.Kind() {
									case reflect.String:
										fieldvalue.SetString(strav)
									case reflect.Int64:
										var i int
										i, erra = strconv.Atoi(strav)
										if erra == nil {
											fieldvalue.SetInt(int64(i))
										}
									default:
										// TODO: handle other data types
										errors.ConsoleWarnf("unfoldComponents %q: unmanaged type %q for attribute %q", newcmpid, fieldvalue.Kind().String(), aname)
									}
								}

								if erra != nil {
									errors.ConsoleWarnf("unfoldComponents %q: value %q type error for attribute %q: %s", newcmpid, strav, aname, erra.Error())
								}
							}
						}

						// recursively unfold the component template
						data := TemplateData{
							Id:  newcmpid,
							Me:  newcmp,
							App: _data,
						}
						var htmlin string
						htmlin, _err = unfoldComponents(_unfoldedCmps, newcmpid, newcmp.Body(), data, _deep+1)
						newcmpelem.SetInnerHTML(htmlin)
						htmlout := newcmpelem.OuterHTML()

						// let's go deeper
						out.WriteString(htmlout)
					}

				} else {
					// the tag is not a registered component
					htmlmsg := fmt.Sprintf("<!-- unable to unfold unregistered ick-%s component -->", tagname)
					out.WriteString(htmlmsg)
					errors.ConsoleWarnf(htmlmsg)
				}
			}
		}
	}
}

// showUnfoldedComponents call addlisteners for every unfolded Components
func showUnfoldedComponents(_unfoldedCmps map[string]Composer) {
	for id, ufc := range _unfoldedCmps {
		e := GetDocument().ChildById(id)
		ufc.Wrap(e)
		ufc.Listeners()
		ufc.Show()
	}
}
