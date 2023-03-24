package ick0

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
	_err = executeBodyTemplate(body, _icmp, _id, _data)
	if _err == nil {
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
	c, _ := ParseClasses(sdftc + " " + string(ac))
	c.AddClasses(*_icmp.SetupClasses())
	if c.Count() > 0 {
		attre.SetAttribute("class", StringQuotes(c.String()))
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
		attre.SetAttribute("style", StringQuotes(s))
	}

	// set the attribute if not forced by an embedded id attribute
	if _, f := attre.Attribute("id"); !f {
		attre.SetAttribute("id", StringQuotes(_id))
	}

	fmt.Fprintf(_output, " %s>", attre.StringQuoted())
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
	processing int // processing operation
	fieldat    int // starting position of the current processing field
	fieldto    int // ending position of the current processing field
}

func (_st *stepway) startfield(i int) {
	_st.fieldat = i
	_st.fieldto = _st.fieldat
}
func (_st *stepway) extendfield(i int) {
	_st.fieldto = i
}
func (_st *stepway) opentxt(i int) {
	_st.processing = processing_TXT
	_st.startfield(i)
}

func (_st *stepway) openick(_pi *int) {
	_st.processing = processing_ICKTAG
	_st.startfield(*_pi + 1)
	_st.extendfield(*_pi + 4)
	*_pi += 5 - 1
}
func (_st *stepway) closeick(_pi *int) {
	_st.processing = processing_NONE
	_st.startfield(*_pi + 2)
	*_pi += 2 - 1
}

func (_st *stepway) openaname(i int) {
	_st.processing = processing_ANAME
	_st.startfield(0)
}
func (_st *stepway) openavalue(i int) {
	_st.processing = processing_AVALUE
	_st.startfield(0)
}

// unfoldBody lookups for ick-component tags in the _body htmlstring and unfold each of them recursively into _output.
// ick-component tags should be in the form <ick-{tagname} [bollattribute] [attribute=[']value[']]/> otherwise
// an error is generated and the unfolding process stops immediatly.
func unfoldBody(_output io.Writer, _body []byte, _data any, _deep int) (_err error) {

	field := func(s stepway) []byte {
		return _body[s.fieldat : s.fieldto+1]
	}

	walk := stepway{processing: processing_NONE}
	var ickname, aname, avalue string
	var bquote byte
	var attrs Attributes

	ilast := len(_body) - 1
	for i := 0; i <= ilast && _err == nil; i++ {
		b := _body[i]
		bclose_delim := string(_body[i:mini(i+2, ilast+1)]) == "/>"
		bopen_delim := string(_body[i:mini(i+5, ilast+1)]) == "<ick-"

		// decide what to do according to walk.processing and b value _</>*
		switch walk.processing {
		case processing_NONE:
			switch {
			case bopen_delim: // start processing an ick-tage
				walk.openick(&i)
			default: // start processing a text field
				walk.opentxt(i)
			}

		case processing_TXT:
			switch {
			case i == ilast: // flush processed text field and exit
				walk.extendfield(ilast)
				_output.Write(field(walk))
			case bopen_delim: // flush processed text field and start processing an ick-tage
				_output.Write(field(walk))
				walk.openick(&i)
			default: // extend the text field
				walk.extendfield(i)
			}

		case processing_ICKTAG:
			if b == ' ' || bclose_delim { // record component tagname
				ickname = string(field(walk))
				if ickname == "ick-" {
					_err = errors.New("'<ick-' tag found without name")
					break
				}
				ickname = strings.ToLower(ickname)
				aname = ""
				avalue = ""
				attrs.Clear()
			}
			switch {
			case b == ' ': // look for another aname
				walk.openaname(i)
			case bclose_delim: // process a single ick-component
				walk.closeick(&i)

				log.Println("composing embedded component:", ickname)
				// fmt.Fprintf(_output, "*** composing embedded component %q ***", ickname)
				_err = unfoldick(_output, ickname, attrs, _data, _deep)

			default: // build component ick-tagname
				r, size := utf8.DecodeRune(_body[i:mini(ilast+1, i+4)])
				if size != 0 && htmlname.IsValidRune(r, false) {
					i += size - 1
					walk.extendfield(i)
				} else {
					_err = fmt.Errorf("invalid character found in ick-tagname: %q", string(_body[walk.fieldat:i+1]))
				}
			}

		case processing_ANAME:
			switch {
			case b == ' ' && walk.fieldat == 0: // trim left spaces
				break
			case (b == ' ' || b == '=' || bclose_delim) && walk.fieldat > 0: // get and save aname
				aname = string(field(walk))
				attrs.SetAttribute(aname, "")
			}

			switch {
			case b == ' ': // look for another aname
				aname = ""
				walk.openaname(i)
			case b == '=': // look for a value
				if aname == "" {
					_err = fmt.Errorf("= symbol found without attribute name: %q", ickname)
					break
				}
				walk.openavalue(i)
				bquote = 0
			case bclose_delim: // process an ick-component
				walk.closeick(&i)

				log.Println("composing embedded component:", ickname)
				// fmt.Fprintf(_output, "*** composing embedded component %q ***", ickname)
				unfoldick(_output, ickname, attrs, _data, _deep)

			default: // build attribute name
				r, size := utf8.DecodeRune(_body[i:mini(ilast+1, i+4)])
				if size > 0 && htmlname.IsValidRune(r, walk.fieldat == 0) {
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
			if bquote == 0 { // don't know yet if a quoted or unquoted value
				switch {
				case b == ' ': // trim left spaces
				case b == '"' || b == '\'': // start a quoted value ?
					bquote = b
					walk.startfield(i + 1)
				case bclose_delim: // empty value
					_err = fmt.Errorf("attribute with empty value: %q", string(_body[walk.fieldat:i+1]))
				default: // start unquoted value
					bquote = 1
					walk.startfield(i)
				}
				break
			}

			switch {
			case bquote == 1 && (b == ' ' || bclose_delim): // process unquoted value
				avalue = string(field(walk))
				attrs.ParseAttribute(aname, avalue)
				switch {
				case bclose_delim: // process an ick-component
					walk.closeick(&i)

					log.Println("composing embedded component:", ickname)
					//fmt.Fprintf(_output, "*** composing embedded component %q ***", ickname)
					unfoldick(_output, ickname, attrs, _data, _deep)

				default: // look for another aname
					walk.openaname(i)
				}
			case bquote != 1 && b == bquote: // process a quoted value
				avalue = string(field(walk))
				attrs.SetAttribute(aname, StringQuotes(avalue))
				walk.openaname(i + 1)
			default: // extend field value
				walk.extendfield(i)
			}
		}
	}
	return _err
}

func unfoldick(_output io.Writer, _ickname string, _attrs Attributes, _data any, _deep int) (_err error) {
	// does this tag refer to a registered component ?
	if regentry := TheCmpReg.LookupComponent(_ickname); regentry != nil {

		// Instantiate the component and get a new id
		newcmpreflect := reflect.New(regentry.typ)
		newcmp := newcmpreflect.Interface().(HtmlComposer)
		newcmpid := TheCmpReg.GetUniqueId(newcmp)

		// DEBUG:
		log.Printf("instantiating %q(%s)\n", newcmpid, newcmpreflect.Type())

		anames := _attrs.Keys()
		for _, aname := range anames {
			_, found := newcmpreflect.Elem().Type().FieldByName(aname)
			if !found {
				// this attribute is not a field of the componenent
				// keep it as is unless it is the class attribute, in this case, add the tokens
				aval, _ := _attrs.Attribute(aname)
				newcmp.SetupAttributes().SetAttribute(aname, aval)
			} else {
				// feed data struct with the value
				strav, _ := _attrs.Attribute(aname)
				field := newcmpreflect.Elem().FieldByName(aname)
				if _err = updateCProperty(field, string(strav)); _err != nil {
					return _err
				}
			}
		}

		// recursively unfold the component template
		data := TemplateData{
			Id:   newcmpid,
			Me:   newcmp,
			Root: _data,
		}
		_err = composeHtmlE(_output, newcmpid, newcmp, data, _deep+1)

	} else {
		// the tag is not a registered component
		// unable to instantiate it
		htmlmsg := fmt.Sprintf("<!-- unable to unfold unregistered %s component -->", _ickname)
		_output.Write([]byte(htmlmsg))
		log.Println(htmlmsg)
		_err = errors.New(htmlmsg)
	}
	return _err
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
