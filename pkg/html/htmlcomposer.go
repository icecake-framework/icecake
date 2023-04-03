package html

import (
	"errors"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/namepattern"
	"github.com/sunraylab/icecake/pkg/registry"
)

type HTMLComposer interface {
	// InlineName returns the unique name of the composer.
	// This name is used to register a component to enable inline instantiation,
	// it's also used as a default class name in the component container.
	// InlineName is case unsensitive.
	// InlineName() string

	// Id Returns the unique Id of a component.
	Id() string

	// Template returns a SnippetTemplate used to render the html string of a Snippet.
	Template(_data *DataState) SnippetTemplate

	// CreateAttribute saves a single key/value attribute. String value must be unquoted.
	// If the key exists, nothing happen.
	CreateAttribute(key string, value any)

	// SetAttribute saves a single key/value attribute. String value must be unquoted.
	// If the key exists value is updated
	SetAttribute(key string, value any)

	// Attributes returns the formated list of attributes used to generate the container element,
	Attributes() String

	// Embed embeds sub-composers
	Embed(id string, cmp any)

	// Embedded return all sub-composers
	Embedded() map[string]any
}

// WriteHTMLSnippet render the HTML of the _composer, its tag element and its body, to _out.
//
// If the composer provides a tagname the output looks like this:
//
//	`<{tagname} id={xxx} class="{ick-tag} [classes]" [attributes]>[body]<tagname/>`
//
// otherwise only the body is written:
//
//	`[body]`
//
// In this case a virtual id (never in the DOM) is returned unless you've forced one before the call.
//
// WriteHTMLSnippet returns the id allocated to the _composer. This Id can be empty if nothing has been rendered when the composer doesn't have a tagname and the generated body is empty.
//
// Every ick-tag founded in the body of the _composer are unfolded and written recursively.
// Direct unfolded components feed the embedded list of the _composer if they implements also the HTMLComposer interface.
func WriteHTMLSnippet(_out io.Writer, _composer HTMLComposer, _data *DataState) (_id string, _err error) {
	return writeHtmlSnippet(_out, _composer, _data, 0)
}

// UnfoldHtml lookups for ick-tags in the _html string and unfold each of them recursively into the _output.
// ick-tags are autoclosing tags and should be in the form:
//
//	`<ick-{tag} [boolattribute] [attribute=[']value['] ...] [property=[']value['] ...]/>`
//
// otherwise an error is generated and the unfolding process stops immediatly.
//
// Direct ick-tags found and instantiated are returned in the _embedded map.
func UnfoldHtml(_out io.Writer, _html String, _data *DataState) (_embedded map[string]any, _err error) {
	virts := &HTMLSnippet{}
	if len(_html) > 0 {
		_err = unfoldBody(virts, _out, []byte(_html), _data, 0)
	}
	return virts.Embedded(), _err
}

/******************************************************************************
* PRIVATE area
******************************************************************************/

// writeHtmlSnippet render the HTML of the _composer, its tag element and its body, to _out.
// Returns the id of the _composer rendered. Id can be empty if nothing has been rendered (composer without tagname and with an empty body).
// if the composer does not have a tagname, a virtual id (never in the DOM) is returned unless you've forced an Id.
// render may be called recursively 10 times max.
func writeHtmlSnippet(_out io.Writer, _composer HTMLComposer, _data *DataState, _deep int) (_id string, _err error) {
	if _deep > 10 {
		_err = fmt.Errorf("RenderHtmlComposer stopped at level %d. Too many recursive calls", _deep)
		log.Println(_err.Error())
		return "", _err
	}

	// get ickname for this _composer
	ickname := ""
	if entry := registry.LookupRegistryEntry(_composer); entry != nil {
		ickname = entry.Name()
	}

	// get best id
	id := _composer.Id()
	if id == "" {
		id = registry.GetUniqueId(ickname)
		_composer.CreateAttribute("id", id)
	}
	_id = _composer.Id()

	// DEBUG:
	fmt.Printf("level=%d -> rendering html snippet id=%q(%s)\n", _deep, _id, reflect.TypeOf(_composer).String())

	// get the template
	t := _composer.Template(_data)
	tagname := helper.NormalizeUp(string(t.TagName))

	if tagname != "" {
		// must merge template attributes with already loaded component attributes
		// the id attribute is always ignored because already setup
		parseAttributes(t.Attributes, _composer)
		_composer.CreateAttribute("class", ickname)
		fmt.Fprintf(_out, "<%s %s>", tagname, _composer.Attributes())
	} else {
		if len(t.Body) == 0 {
			log.Printf("Warning empty html snippet, no tagname and no body: level=%d, type:%s\n", _deep, reflect.TypeOf(_composer).String())
			return "", nil
		}
	}
	// DEBUG:
	if _composer.Id() != _id {
		panic("renderHtmlSnippet: id discrepency")
	}

	if len(t.Body) > 0 {
		_err = unfoldBody(_composer, _out, []byte(t.Body), _data, _deep)
	}

	if tagname != "" {
		fmt.Fprintf(_out, "</%s>", tagname)
	}
	return _id, _err
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
func (_st *stepway) openick(i int) {
	_st.processing = processing_ICKTAG
	_st.fieldat = i + 1
	_st.fieldto = i + 4
}
func (_st *stepway) closeick(i int) {
	_st.processing = processing_NONE
	_st.startfield(i + 2)
}

// unfoldBody lookups for ick-component tags in the _body htmlstring and unfold each of them recursively into _output.
// ick-component tags are autoclosing tags and should be in the form:
//
//	`<ick-{tagname} [boolattribute] [attribute=[']value[']]/>`
//
// otherwise an error is generated and the unfolding process stops immediatly.
func unfoldBody(_parent HTMLComposer, _output io.Writer, _body []byte, _data *DataState, _deep int) (_err error) {

	field := func(s stepway) []byte {
		return _body[s.fieldat:s.fieldto]
	}

	walk := stepway{processing: processing_NONE}
	var ickname, aname, avalue string
	var bquote byte
	attrs := make(map[string]string, 0)

	ilast := len(_body) - 1
nextbyte:
	for i := 0; i <= ilast && _err == nil; i++ {
		b := _body[i]
		bclose_delim := string(_body[i:mini(i+2, ilast+1)]) == "/>"
		bopen_delim := string(_body[i:mini(i+5, ilast+1)]) == "<ick-"

		// decide what to do according to walk.processing and b value _</>*
		switch walk.processing {
		case processing_NONE:
			switch {
			case bopen_delim: // start processing an ick-tage
				walk.openick(i)
				i += 5 - 1
			default: // start processing a text field
				walk.processing = processing_TXT
				walk.startfield(i)
			}

		case processing_TXT:
			switch {
			case i == ilast: // flush processed text field and exit
				walk.fieldto = ilast + 1
				_output.Write(field(walk))
			case bopen_delim: // flush processed text field and start processing an ick-tage
				walk.fieldto = i
				_output.Write(field(walk))
				walk.openick(i)
				i += 5 - 1
			default: // extend the text field
				walk.fieldto = i
			}

		case processing_ICKTAG:
			if b == ' ' || bclose_delim { // record component tagname
				walk.fieldto = i
				ickname = string(field(walk))
				if ickname == "ick-" {
					_err = errors.New("'<ick-' tag found without name")
					break
				}
				ickname = strings.ToLower(ickname)
				aname = ""
				avalue = ""
				attrs = make(map[string]string, 0)
			}
			switch {
			case b == ' ': // look for another aname
				walk.processing = processing_ANAME
				walk.startfield(0)
			case bclose_delim: // process a single ick-component
				walk.closeick(i)
				i += 2 - 1

				//log.Println("composing embedded component:", ickname)
				_err = unfoldick(_parent, _output, ickname, attrs, _data, _deep)

			default: // build component ick-tagname
				r, size := utf8.DecodeRune(_body[i:mini(ilast+1, i+4)])
				if size != 0 && namepattern.IsValidRune(r, false) {
					i += size - 1
					walk.fieldto = i
				} else {
					_err = fmt.Errorf("invalid character found in ick-tagname: %q", string(_body[walk.fieldat:i+1]))
				}
			}

		case processing_ANAME:
			switch {
			case (b == ' ' || b == '\n' || b == '\t') && walk.fieldat == 0: // trim left spaces and \n
				continue nextbyte
			case (b == ' ' || b == '=' || b == '\n' || b == '\t' || bclose_delim) && walk.fieldat > 0: // get and save aname
				walk.fieldto = i
				aname = string(field(walk))
				attrs[aname] = ""
			}

			switch {
			case b == ' ': // look for another aname
				aname = ""
				walk.processing = processing_ANAME
				walk.startfield(0)
			case b == '=': // look for a value
				if aname == "" {
					_err = fmt.Errorf("= symbol found without attribute name: %q", ickname)
					break
				}
				walk.processing = processing_AVALUE
				walk.startfield(0)
				bquote = 0
			case bclose_delim: // process an ick-component
				walk.closeick(i)
				i += 2 - 1

				//log.Println("composing embedded component:", ickname)
				unfoldick(_parent, _output, ickname, attrs, _data, _deep)

			default: // build attribute name
				r, size := utf8.DecodeRune(_body[i:mini(ilast+1, i+4)])
				if size > 0 && namepattern.IsValidRune(r, walk.fieldat == 0) {
					if walk.fieldat == 0 {
						walk.startfield(i)
					}
					i += size - 1
					walk.fieldto = i
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
				walk.fieldto = i
				avalue = string(field(walk))
				attrs[aname] = parseQuoted(avalue)
				switch {
				case bclose_delim: // process an ick-component
					walk.closeick(i)
					i += 2 - 1

					log.Println("composing embedded component:", ickname)
					//fmt.Fprintf(_output, "*** composing embedded component %q ***", ickname)
					unfoldick(_parent, _output, ickname, attrs, _data, _deep)

				default: // look for another aname
					walk.processing = processing_ANAME
					walk.startfield(0)
				}
			case bquote != 1 && b == bquote: // process a quoted value
				walk.fieldto = i
				avalue = string(field(walk))
				attrs[aname] = avalue
				walk.processing = processing_ANAME
				walk.startfield(0)
			default: // extend field value
				walk.fieldto = i
			}
		}
	}
	return _err
}

// unfoldick render the ick-component corresponding to _ickname and its unfolded _attrs.
func unfoldick(_parent HTMLComposer, _output io.Writer, _ickname string, _attrs map[string]string, _data *DataState, _deep int) (_err error) {
	// does this tag refer to a registered component ?
	regentry := registry.GetRegistryEntry(_ickname)
	if regentry.Component() != nil {

		// clone the registered component
		newref := reflect.New(reflect.TypeOf(regentry.Component()).Elem())
		newref.Elem().Set(reflect.ValueOf(regentry.Component()).Elem())
		newcmp := newref.Interface().(HTMLComposer)

		// process unfolded attributes, set value of ickcomponent field when name of attribute matches field name,
		// otherwise set unfolded attribute to the attribute of the component.
		for aname, avalue := range _attrs {
			_, found := newref.Elem().Type().FieldByName(aname)
			if !found {
				// this attribute is not a field of the componenent
				// keep it as is unless it is the class attribute, in this case, add the tokens
				newcmp.CreateAttribute(aname, String(avalue))
			} else {
				// feed data struct with the value
				field := newref.Elem().FieldByName(aname)
				if _err = updateProperty(field, avalue); _err != nil {
					return _err
				}
			}
		}

		// recursively unfold the component snippet
		newcmpid := ""
		newcmpid, _err = writeHtmlSnippet(_output, newcmp, _data, _deep+1)

		// add it to the map of embedded components
		if newcmpid != "" && _parent != nil {
			_parent.Embed(newcmpid, newcmp)
		}

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

// updateProperty updates _cprop with the _value trying to convert the _value to the type of _cprop
func updateProperty(_cprop reflect.Value, _value string) (_erra error) {
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

// ParseQuoted returns a trimed value keeping white space inside quotes if any.
// If _value does not have quotes, the returned value is truncated at the first white space found.
func parseQuoted(_str string) string {
	trimspaces := strings.Trim(_str, " ")

	trimq1 := strings.Trim(trimspaces, "'")
	if len(trimq1) == len(trimspaces)-2 {
		return trimq1
	}

	trimq2 := strings.Trim(trimspaces, "\"")
	if len(trimq2) == len(trimspaces)-2 {
		return trimq2
	}

	s, _, _ := strings.Cut(trimspaces, " ")
	return s
}

// ParseAttribute split _alist into attributes separated by spaces and set each to the HtmlComposer.
// An attribute can have a value at the right of an "=" symbol.
// The value can be delimited by quotes ( " or ' ) and in that case may contains whitespaces.
// The string is processed until the end or an error occurs when invalid char is met.
func parseAttributes(_alist string, _cmp HTMLComposer) (_err error) {

	//pattrs = new(Attributes)
	//pattrs.amap = make(map[string]StringQuotes)
	var strnames string
	unparsed := _alist
	for i := 0; len(unparsed) > 0; i++ {

		// process all simple attributes until next "="
		var hasval bool
		strnames, unparsed, hasval = strings.Cut(unparsed, "=")
		names := strings.Fields(strnames)
		for i, n := range names {
			if !namepattern.IsValid(n) {
				return fmt.Errorf("attribute name %q is not valid", n)
			}
			if i < len(names)-1 || !hasval {
				_cmp.CreateAttribute(n, "")
			}
			//_pattrs.amap[n] = ""
		}

		// remove blanks just after "="
		unparsed = strings.TrimLeft(unparsed, " ")

		// stop if nothing else to proceed
		if len(unparsed) == 0 || len(names) == 0 {
			break
		}

		// extract attribute name with a value
		name := names[len(names)-1]

		// extract value with quotes or no quotes
		var value string
		delim := unparsed[0]
		istart := 1
		if delim != '"' && delim != '\'' {
			delim = ' '
			istart = 0
		}
		value, unparsed, _ = strings.Cut(unparsed[istart:], string(delim))
		_cmp.CreateAttribute(name, String(value))
		//		_pattrs.amap[name] = StringQuotes(value)
	}
	return nil
}

// mini helper
func mini(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
