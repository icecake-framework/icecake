package ick

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type ID string

func (_id *ID) Force(_forceid string) {
	*_id = ID(_forceid)
}

type HtmlComposer interface {
	RegisterName() string
	RegisterCSS() string

	//TagName() string

	SetupId() *ID
	SetupClasses() *Classes
	SetupAttributes() *Attributes
	SetupStyle() *Style

	Container(_compid string) (_tagname string, _contclasses string, _contattrs string, _contstyle string)
	Body() (_unsafehtml string)

	Embed(id string, cmp any)
}

// HtmlComponent implements most of HtmlComposer interface unless RegisterName().
// Custom components should embedd HtmlComponent and implement the RegisterName() interface.
type HtmlComponent struct {
	setupid ID // if empty the composer will generate a unique id

	setupclasses    Classes
	setupattributes Attributes
	setupstyle      Style

	embedded map[string]any // instantiated embedded objects
}

func (_cmp *HtmlComponent) SetupId() *ID {
	return &_cmp.setupid
}

func (_cmp *HtmlComponent) RegisterCSS() string {
	return ""
}

func (_cmp *HtmlComponent) Container(_compid string) (_tagname string, _contclasses string, _contattrs string, _contstyle string) {
	// DEBUG: log.Printf("HtmlComponent returns default <SPAN> container for %q\n", _compid)
	return "SPAN", "", "", ""
}

func (_cmp *HtmlComponent) Body() (_html string) {
	// DEBUG: log.Printf("HtmlComponent returns an empty body\n")
	return ""
}

// Classes provides a reference to HtmlComponent classes. By default this is the internal HtmlComponent classes.
// This reference can be used to setup custom classes for a new component.
// These classes will overwrite the container classes.
func (_cmp *HtmlComponent) SetupClasses() *Classes {
	return &_cmp.setupclasses
}

func (_cmp *HtmlComponent) SetupAttributes() *Attributes {
	return &_cmp.setupattributes
}

func (_cmp *HtmlComponent) SetupStyle() *Style {
	return &_cmp.setupstyle
}

func (_cmp *HtmlComponent) Embed(id string, cmp any) {
	if _cmp.embedded == nil {
		_cmp.embedded = make(map[string]any)
	}
	_cmp.embedded[id] = cmp
}

type TemplateData struct {
	Id   string // the id of the processing component
	Me   any    // the processing component
	Root any    // the App object, can be nil
}

// ComposeHtml writes into _wr the html string representing the HTML element of the HtmlComposer.
// HtmlComposer provides the container information. If the component does not have an ID yet, a
// unique ID is generated and assigned to the component.
// The content of the HTML Element is built upon the body template executed with _topdata
// The
func ComposeHtmlE(_wr io.Writer, _icmp HtmlComposer, _rootdata any) error {

	id := string(*_icmp.SetupId())
	if id == "" {
		id = TheCmpReg.GetUniqueId(_icmp)
	}

	data := TemplateData{
		Id:   id,
		Me:   _icmp,
		Root: _rootdata,
	}
	return composeHtmlE(_wr, id, _icmp, data, 0)
}

// composeHtml may be called recursively and if protected against infinite call with _deep
func composeHtmlE(_wr io.Writer, _id string, _icmp HtmlComposer, _data any, _deep int) error {
	if _deep >= 10 {
		err := fmt.Errorf("composeHtmlE stopped at level %d. Too many recursive calls", _deep)
		log.Println(err.Error())
		return err
	}

	// DEBUG:
	log.Printf("composing Html Element at level %d: id=%s, type:%s\n", _deep, _id, reflect.TypeOf(_icmp).String())

	tagname := openHtmlE(_wr, _icmp, _id)
	err := unfoldBody(_wr, _icmp, _id, _data, _deep)
	closeHtmlE(_wr, tagname)
	return err
}

func openHtmlE(_wr io.Writer, _icmp HtmlComposer, _id string) (_tagname string) {

	tagname, strclasses, strattrs, strstyle := _icmp.Container(_id)

	tagname = strings.ToUpper(strings.Trim(tagname, " "))
	fmt.Fprintf(_wr, "<%s", tagname)

	attrs, _ := ParseAttributes(strattrs)
	_icmp.SetupAttributes().SetAttributes(*attrs, false)

	classes, _ := ParseClasses(strclasses)
	_icmp.SetupClasses().AddClasses(*classes)
	strclasses = _icmp.SetupClasses().String()
	if strclasses != "" {
		_icmp.SetupAttributes().SetAttribute("class", strclasses)
	}

	custs := *_icmp.SetupStyle()
	*_icmp.SetupStyle() = Style(strstyle) + custs
	strstyle = string(*_icmp.SetupStyle())
	if strstyle != "" {
		_icmp.SetupAttributes().SetAttribute("style", strstyle)
	}

	_icmp.SetupAttributes().SetAttribute("id", _id)

	strattrs = _icmp.SetupAttributes().String()
	if strattrs != "" {
		fmt.Fprint(_wr, ` `, strattrs)
	}

	fmt.Fprint(_wr, ">")
	return tagname
}

func closeHtmlE(_wr io.Writer, _tagname string) {
	fmt.Fprintf(_wr, "</%s>", _tagname)
}

// unfoldBody lookup for component tags in the body htmlstring, and compose each of them recursively.
//
// unfoldBody parses body template, according to go html templating standards, and
// execute this template with {{}} langage and component's data.
func unfoldBody(_wr io.Writer, _icmp HtmlComposer, _id string, _data any, _deep int) (_err error) {

	// 1. parse body template
	tmpCmp := template.Must(template.New("").Parse(_icmp.Body()))

	// 2. execute the template with _data
	bufCmp := new(bytes.Buffer)
	errt := tmpCmp.Execute(bufCmp, _data)
	if errt != nil {
		err := fmt.Errorf("unfolding %q stopped on error: %s", _id, errt.Error())
		log.Println(err.Error())
		_wr.Write([]byte("<!-- composing template error -->"))
		return err
	}
	htmlstring := bufCmp.String()

	// 3. lookup for ick components
	const (
		delim_open  = "<ick-"
		delim_close = "/>"
	)

	//n := 0
	//out := &bytes.Buffer{}
nextdelim:
	for {
		// find next delim_open
		if from := strings.Index(htmlstring, delim_open); from == -1 || _err != nil {
			// no more delim_open = feed the output with what's left in htmlstring and return
			_wr.Write([]byte(htmlstring))
			if _err != nil {
				log.Println(_err.Error())
				_wr.Write([]byte("<!-- unfolding error -->"))
			}
			return _err
		} else {
			// we found a new delim_open
			// so it's time to feed the output with data preceding this delim
			_wr.Write([]byte(htmlstring[:from]))

			// scrap this data and keep what's left
			htmlstring = htmlstring[from+len(delim_open):]

			// look now for it's corresponding delim_close
			if to := strings.Index(htmlstring, delim_close); to == -1 {
				// not corresponding delim_close then stop and return a rendering error
				_err = fmt.Errorf("unfolding %q stopped: closing delimiter '/>' not found", _id)
				continue nextdelim
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
				if regentry := TheCmpReg.LookupComponent("ick-" + tagname); regentry != nil {

					// process and instantiate new component

					// Instantiate the component
					newcmpreflect := reflect.New(regentry.typ)
					newcmp := newcmpreflect.Interface().(HtmlComposer)
					//newcmpid := fmt.Sprintf("ick-%s-%d-%d", tagname, _deep, n)

					newcmpid := TheCmpReg.GetUniqueId(newcmp)

					//					var newcmpelem *UIComponent
					//					var newcmpid string
					//					if newcmpid, newcmpelem, _err = App.CreateComponent(newcmp); _err == nil {

					log.Printf("unfolding %q: instantiating %q(%s)\n", _id, newcmpid, newcmpreflect.Type())

					// add the component to the add it to the unfolded stack
					_icmp.Embed(newcmpid, newcmpreflect.Interface())

					// process embeded component's attributes
					var attrs *Attributes
					attrs, _err = ParseAttributes(leftinside)
					if _err != nil {
						continue nextdelim
					}

					// set the id, overwrite the one in the template if any
					//attrs.SetAttribute("id", newcmpid)

					anames := attrs.Keys()
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
							// if aname == "class" {
							// 	newcmp.Classes().AddTokens(aval)
							// } else {
							newcmp.SetupAttributes().SetAttribute(aname, aval)
							// }
						} else {
							// DEBUG:
							// fmt.Printf("attribute %v: %q corresponds to a component's data\n", i, aname)

							// feed data struct with the value
							strav, _ := attrs.Attribute(aname)
							field := newcmpreflect.Elem().FieldByName(aname)
							if _err = updateCProperty(field, strav); _err != nil {
								//log.Printf("unfolding %q body, attribute %q: %s", _id, aname, _err.Error())
								continue nextdelim
							}
						}
					}

					// recursively unfold the component template
					data := TemplateData{
						Id:   newcmpid,
						Me:   newcmp,
						Root: _data,
					}
					composeHtmlE(_wr, newcmpid, newcmp, data, _deep+1)

					// var htmlin string
					// htmlin, _err = unfoldComponents(_unfoldedCmps, newcmpid, newcmp.Body(), data, _deep+1)
					// newcmpelem.SetInnerHTML(htmlin)
					// htmlout := newcmpelem.OuterHTML()

					// let's go deeper
					//_wr.Write(htmlout)
					//					}

				} else {
					// the tag is not a registered component
					htmlmsg := fmt.Sprintf("<!-- unable to unfold unregistered 'ick-%s' component -->", tagname)
					_wr.Write([]byte(htmlmsg))
					log.Println(htmlmsg)
				}
			}
		}
	}
}

func updateCProperty(_cprop reflect.Value, _value string) (_erra error) {

	switch _cprop.Type().String() {
	case "time.Duration":
		var d time.Duration
		d, _erra = time.ParseDuration(_value)
		if _erra == nil {
			_cprop.SetInt(int64(d))
		}

	default:
		switch _cprop.Kind() {
		case reflect.String:
			_cprop.SetString(_value)

		case reflect.Int64:
			var i int
			i, _erra = strconv.Atoi(_value)
			if _erra == nil {
				_cprop.SetInt(int64(i))
			}

		// TODO: handle other data types
		default:
			return fmt.Errorf("unmanaged type %q", _cprop.Kind().String())
		}
	}
	return _erra
}
