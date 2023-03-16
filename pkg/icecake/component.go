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
	"unicode/utf8"

	"github.com/sunraylab/icecake/pkg/htmlname"
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
	// body := new(strings.Builder)
	_err = executeBodyTemplate(body, _icmp, _id, _data)
	if _err != nil {
		_err = unfoldBody(_output, body.Bytes(), _data, _deep)
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

const (
	processing_NONE int = iota
	processing_TXT
	processing_ICKTAG
	processing_ANAME
	processing_AVALUE
)

type stepway struct {
	processing int
	fieldat    int
	fieldto    int
}

func (_st *stepway) startfield(i int) {
	_st.fieldat = i
	_st.fieldto = _st.fieldat
}
func (_st *stepway) opentxt(i int) {
	_st.processing = processing_TXT
	_st.startfield(i)
}
func (_st *stepway) openick(_pi *int) {
	_st.processing = processing_ICKTAG
	_st.fieldat = *_pi + 1
	_st.fieldto = _st.fieldat + 4
	*_pi += 5 - 1
}
func (_st *stepway) closeick(_pi *int) {
	_st.processing = processing_NONE
	_st.startfield(*_pi + 2)
	*_pi += 2 - 1
}
func (_st *stepway) extendfield(i int) {
	_st.fieldto = i
}
func (_st *stepway) openaname(i int) {
	_st.processing = processing_ANAME
	_st.startfield(0)
}
func (_st *stepway) openavalue(i int) {
	_st.processing = processing_AVALUE
	_st.startfield(0)
}

// func (_st stepway) emptyfield(_pi *int) bool {
// 	return _st.fieldat == _st.fieldto
// }

func unfoldBody(_output io.Writer, _body []byte, _data any, _deep int) (_err error) {
	// const (
	// 	delim_open  = "<ick-"
	// 	delim_close = "/>"
	// )
	field := func(s stepway) []byte {
		return _body[s.fieldat : s.fieldto+1]
	}

	walk := stepway{processing: processing_NONE}
	var ickname, aname, avalue string
	var quote byte
	var attrs Attributes

	ilast := len(_body) - 1
	for i := 0; i <= ilast && _err == nil; i++ {
		b := _body[i]

		// _</>*

		var fdelim_open, fdelim_close bool
		if i+1 <= ilast {
			fdelim_close = string(_body[i:i+2]) == "/>"
			if i+5 <= ilast {
				fdelim_open = string(_body[i:i+5]) == "<ick-"
			}
		}

		switch walk.processing {
		case processing_NONE:
			switch {
			case fdelim_open:
				walk.openick(&i)
			default:
				walk.opentxt(i)
			}

		case processing_TXT:
			switch {
			case i == ilast:
				// flush processed field and exit
				walk.fieldto = ilast
				_output.Write(field(walk))
			case fdelim_open:
				_output.Write(field(walk))
				walk.openick(&i)
			default:
				walk.extendfield(i)
			}

		case processing_ICKTAG:
			if b == ' ' || fdelim_close { // record component tagname
				ickname = string(field(walk))
				if ickname == "ick-/" || ickname == "ick- " {
					_err = errors.New("ick tag found without name")
					break
				}
				ickname = strings.ToLower(ickname)
				aname = ""
				avalue = ""
				attrs.Clear()
			}
			switch {
			case b == ' ':
				walk.openaname(i)
			case fdelim_close: // single tagname component
				walk.closeick(&i)

				log.Println("composing embedded component:", ickname)
				fmt.Fprintf(_output, "*** composing embedded component %q ***", ickname)

			default: // build component ick-tagname
				r, size := utf8.DecodeRune(_body[i : ilast+1])
				if size != 0 && htmlname.IsValidRune(r, true) {
					i += size - 1
					walk.extendfield(i)
				} else {
					_err = fmt.Errorf("invalid character found in ick-tagname: %q", string(_body[walk.fieldat:i+1]))
				}
			}

		case processing_ANAME:
			// trim left spaces
			if b == ' ' && walk.fieldat == 0 {
				break
			}
			// get and save aname
			if (b == ' ' || b == '=' || fdelim_close) && walk.fieldat > 0 {
				aname = string(field(walk))
				attrs.SetAttribute(aname, "")
			}
			switch {
			case b == ' ': // let's continue for another name
				aname = ""
				walk.openaname(i)
			case b == '=': // let's continue for a value
				if aname == "" {
					_err = fmt.Errorf("= symbol found without attribute name: %q", ickname)
					break
				}
				walk.openavalue(i)
			case fdelim_close: // time to compose the component
				walk.closeick(&i)

				log.Println("composing embedded component:", ickname)
				fmt.Fprintf(_output, "*** composing embedded component %q ***", ickname)

			default: // build attribute name
				r, size := utf8.DecodeRune(_body[i : ilast+1])
				if size != 0 && htmlname.IsValidRune(r, false) {
					if walk.fieldat == 0 {
						walk.startfield(i)
					}
					i += size - 1
					walk.extendfield(i)
				} else {
					_err = fmt.Errorf("invalid character found in attribute name: %q", string(_body[walk.fieldat:i+1]))
				}
			}

		case processing_AVALUE:
			// don't know yet if a quoted or unquoted value
			if quote == 0 {
				// trim left spaces
				if b == ' ' && quote == 0 {
					break
				}
				// start a quoted value ?
				if b == '"' || b == '\'' {
					quote = b
					walk.startfield(i + 1)
					break
				}
				// empty value
				if fdelim_close {
					_err = fmt.Errorf("attribute with empty value: %q", string(_body[walk.fieldat:i+1]))
				}
				// start unquoted value
				quote = 1
				walk.startfield(i)
				break
			}

			switch {
			// end of unquoted value
			case quote == 1 && (b == ' ' || fdelim_close):
				avalue = string(field(walk))
				attrs.SetAttribute(aname, avalue)

				if fdelim_close { // time to compose the component
					walk.closeick(&i)

					log.Println("composing embedded component:", ickname)
					fmt.Fprintf(_output, "*** composing embedded component %q ***", ickname)

				} else {
					walk.openaname(i)
				}
			// end of quoted value
			case quote != 1 && b == quote:
				avalue = string(field(walk))
				attrs.SetAttribute(aname, avalue)
				//i += 1
				walk.openaname(i + 1)

			default:
				walk.extendfield(i)
			}
		}
	}

	return _err
}

// unfoldBody lookup for component tags in the body htmlstring, and compose each of them recursively.
func unfoldBody2(_output io.Writer, _body io.Reader, _data any, _deep int) (_err error) {

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
