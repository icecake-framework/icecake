package ick

import (
	"bytes"
	"errors"
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

// SetupClasses provides a reference to HtmlComponent classes. By default this is the internal HtmlComponent classes.
// This reference can be used to setup custom classes for a new component.
// These classes will overwrite the default container classes of the component.
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
func composeHtmlE(_output io.Writer, _id string, _icmp HtmlComposer, _data any, _deep int) (_err error) {
	if _deep >= 10 {
		_err = fmt.Errorf("composeHtmlE stopped at level %d. Too many recursive calls", _deep)
		log.Println(_err.Error())
		return _err
	}

	// DEBUG:
	log.Printf("composing Html Element at level %d: id=%s, type:%s\n", _deep, _id, reflect.TypeOf(_icmp).String())

	// add the component to the unfolded stack
	_icmp.Embed(_id, _icmp)

	tagname := openHtmlE(_output, _icmp, _id)

	body := new(bytes.Buffer)
	_err = executeBodyTemplate(body, _icmp, _id, _data)
	if _err != nil {
		_err = unfoldBody(_output, body, _data, _deep)
	}
	closeHtmlE(_output, tagname)
	return _err
}

func openHtmlE(_output io.Writer, _icmp HtmlComposer, _id string) (_tagname string) {

	tagname, sdftc, sdfta, sdfts := _icmp.Container(_id)

	tagname = strings.ToUpper(strings.Trim(tagname, " "))
	fmt.Fprintf(_output, "<%s", tagname)

	// classes, attributes and styles are all attributes
	attre := new(Attributes)

	// update setupattributes with default container ones, do not overwrite
	dfta, _ := ParseAttributes(sdfta)
	attre.SetAttributes(*_icmp.SetupAttributes(), false)
	attre.SetAttributes(*dfta, false)

	// merge setup classes with default container ones and class defined as an attribute if any
	ac, _ := attre.Attribute("class")
	c, _ := ParseClasses(sdftc + " " + ac)
	c.AddClasses(*_icmp.SetupClasses())
	if c.Count() > 0 {
		attre.SetAttribute("class", c.String())
	}

	// merge setup style with default container ones and style defined as an attribute if any
	as, _ := attre.Attribute("style")
	s := *_icmp.SetupStyle()
	if as != "" {
		s += Style(as)
	}
	if sdfts != "" {
		s += Style(sdfts)
	}
	if s != "" {
		attre.SetAttribute("style", string(s))
	}

	// set the attribute if not forced by an embedded id attribute
	if _, f := attre.Attribute("id"); !f {
		attre.SetAttribute("id", _id)
	}

	fmt.Fprintf(_output, " %s>", attre.String())
	return tagname
}

func closeHtmlE(_output io.Writer, _tagname string) {
	fmt.Fprintf(_output, "</%s>", _tagname)
}

// executeBodyTemplate parses _icmp.body template according to the go html templating standards,
// and execute this template with {{}} syntax and _data.
func executeBodyTemplate(_output io.Writer, _icmp HtmlComposer, _id string, _data any) (_err error) {
	t := template.Must(template.New("").Parse(_icmp.Body()))
	_err = t.Execute(_output, _data)
	if _err != nil {
		_output.Write([]byte("<!-- composing html template error -->"))
	}
	return _err
}

// unfoldBody lookup for component tags in the body htmlstring, and compose each of them recursively.
func unfoldBody(_output io.Writer, _body io.Reader, _data any, _deep int) (_err error) {

	// 3. lookup for ick components
	const (
		delim_open  = "<ick-"
		delim_close = "/>"
	)

	htmlstring := ""

nextdelimo:
	for {
		// find next delim_open
		if from := strings.Index(htmlstring, delim_open); from == -1 || _err != nil {
			// no more delim_open = feed the output with what's left in htmlstring and return
			_output.Write([]byte(htmlstring))
			if _err != nil {
				log.Println(_err.Error())
				_output.Write([]byte("<!-- unfolding error -->"))
			}
			return _err
		} else {
			// we found a new delim_open, so feed the output with data preceding this delim
			_output.Write([]byte(htmlstring[:from]))

			// scrap this data and keep what's left
			htmlstring = htmlstring[from+len(delim_open):]

			// look now for it's corresponding delim_close
			if to := strings.Index(htmlstring, delim_close); to == -1 {
				// not corresponding delim_close then stop and return a rendering error
				_err = errors.New("closing delimiter '/>' not found")
				continue nextdelimo
			} else {

				// we got a delim_close so we've a new ick-element, extract its content
				inside := htmlstring[0:to]
				ickname, leftinside, _ := strings.Cut(inside, " ")
				htmlstring = htmlstring[to+len(delim_close):] // scrap it and keep what's left

				if ickname == "" { // <ick-/> !
					continue nextdelimo
				}

				// does this tag refer to a registered component ?
				if regentry := TheCmpReg.LookupComponent("ick-" + ickname); regentry != nil {

					// Instantiate the component and get a new id
					newcmpreflect := reflect.New(regentry.typ)
					newcmp := newcmpreflect.Interface().(HtmlComposer)
					newcmpid := TheCmpReg.GetUniqueId(newcmp)

					// DEBUG:
					log.Printf("instantiating %q(%s)\n", newcmpid, newcmpreflect.Type())

					// // add the component to the add it to the unfolded stack
					// _icmp.Embed(newcmpid, newcmpreflect.Interface())

					// process embeded component's attributes
					var attrs *Attributes
					attrs, _err = ParseAttributes(leftinside)
					if _err != nil {
						continue nextdelimo
					}

					anames := attrs.Keys()
					for _, aname := range anames {
						_, found := newcmpreflect.Elem().Type().FieldByName(aname)
						if !found {
							// this attribute is not a field of the componenent
							// keep it as is unless it is the class attribute, in this case, add the tokens
							aval, _ := attrs.Attribute(aname)
							newcmp.SetupAttributes().SetAttribute(aname, aval)
						} else {
							// feed data struct with the value
							strav, _ := attrs.Attribute(aname)
							field := newcmpreflect.Elem().FieldByName(aname)
							if _err = updateCProperty(field, strav); _err != nil {
								continue nextdelimo
							}
						}
					}

					// recursively unfold the component template
					data := TemplateData{
						Id:   newcmpid,
						Me:   newcmp,
						Root: _data,
					}
					composeHtmlE(_output, newcmpid, newcmp, data, _deep+1)

				} else {
					// the tag is not a registered component
					htmlmsg := fmt.Sprintf("<!-- unable to unfold unregistered 'ick-%s' component -->", ickname)
					_output.Write([]byte(htmlmsg))
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
