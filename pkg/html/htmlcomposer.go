package html

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/icecake-framework/icecake/pkg/registry"
	"github.com/icecake-framework/icecake/pkg/stringpattern"
	"github.com/sunraylab/verbose"
)

type HTMLComposer interface {

	// SetDataState links an optional data states to the composer.
	SetDataState(ds *DataState)

	// Id Returns the unique Id of a Snippet or empty.
	// The id is a tag attribute with specific rules.
	//Id() string

	// SetId sets or overwrites the id attribute of the html snippet
	//SetId(id string)

	// Tag returns the tag used to render the html element.
	// Composer rendering processes call Tag once for every rendering and before.
	// If the implementer returns nil or an empty tag, only the body will be rendered.
	// It's up to the implementer to persist the tag if required between multiple rendering of the same snippet.
	Tag() *Tag

	// WriteBody writes the HTML string corresponing to the body of the HTML element.
	// FIXME: RenderContent
	WriteBody(out io.Writer) error

	// Embedded returns all sub-composers
	Embedded() map[string]any

	// Embed embeds sub-composers
	Embed(id string, subcmp HTMLComposer)

	// String renders and unfold the _snippet and returns its corresponding HTML string
	//String() HTMLString
}

// WriteSnippet writes the HTML string of the composer, its tag element and its body, to out.
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
// WriteSnippet returns the id allocated to the composer.
// This Id can be empty if nothing has been rendered when the composer doesn't have a tagname and the generated body is empty.
//
// Every ick-tag founded in the body of the composer are unfolded and written recursively.
// Each unfolded component is saved in the composer embedded list (only if they implement the HTMLComposer interface.)
func WriteSnippet(out io.Writer, composer HTMLComposer, ds *DataState, noid bool) (virtualid string, err error) {
	return writeSnippet(out, composer, ds, noid, "", 0, 0)
}

// Unfold lookups for ick-tags in the _html string and unfolds each of them recursively into the _output.
// ick-tags are autoclosing tags and should be in the form:
//
//	`<ick-{tag} [boolattribute] [attribute=[']value['] ...] [property=[']value['] ...]/>`
//
// otherwise an error is generated and the unfolding process stops immediatly.
//
// Direct ick-tags found and instantiated are returned in the _embedded map.
func UnfoldHTML(out io.Writer, html HTMLString, data *DataState) (embedded map[string]any, err error) {
	tmpparent := &HTMLSnippet{}
	if len(html) > 0 {
		err = unfoldHTML(tmpparent, out, []byte(html), data, 0)
	}
	return tmpparent.Embedded(), err
}

// RenderSnippet builds and unfolds the _snippet HTMLComposer and returns its html string.
// RenderSnippet does not mount the component into the DOM and so it can't respond to events.
// FIXME: clarify use of RenderSnippet
// func RenderSnippet(_snippet HTMLComposer) (html HTMLString, id string, err error) {
// 	out := new(bytes.Buffer)
// 	id, err = WriteSnippet(out, _snippet, nil, true)
// 	if err == nil {
// 		html = HTMLString(out.String())
// 	}
// 	return html, id, err
// }

// // ParseAttribute split alist into attributes separated by spaces and set each to the HtmlComposer.
// // An attribute can have a value at the right of an "=" symbol.
// // The value can be delimited by quotes ( " or ' ) and in that case may contains whitespaces.
// // The string is processed until the end or an error occurs when invalid char is met.
// // Existing _cmp attributes are not overwritten.
// // TODO: secure _alist ?
// func ParseAttributes(alist string, cmp HTMLComposer) (err error) {

// 	var strnames string
// 	unparsed := alist
// 	for i := 0; len(unparsed) > 0; i++ {

// 		// process all simple attributes until next "="
// 		var hasval bool
// 		strnames, unparsed, hasval = strings.Cut(unparsed, "=")
// 		names := strings.Fields(strnames)
// 		for i, n := range names {
// 			if !namepattern.IsValid(n) {
// 				return fmt.Errorf("attribute name %q is not valid", n)
// 			}
// 			if i < len(names)-1 || !hasval {
// 				cmp.SetAttribute(n, "", false)
// 			}
// 		}

// 		// remove blanks just after "="
// 		unparsed = strings.TrimLeft(unparsed, " ")

// 		// stop if nothing else to proceed
// 		if len(unparsed) == 0 || len(names) == 0 {
// 			break
// 		}

// 		// extract attribute name with a value
// 		name := names[len(names)-1]

// 		// extract value with quotes or no quotes
// 		var value string
// 		delim := unparsed[0]
// 		istart := 1
// 		if delim != '"' && delim != '\'' {
// 			delim = ' '
// 			istart = 0
// 		}
// 		value, unparsed, _ = strings.Cut(unparsed[istart:], string(delim))
// 		cmp.SetAttribute(name, value, false)
// 	}
// 	return nil
// }

// maxDEEP is the maximum HTML string unfolding levels
const maxDEEP int = 25

// writeSnippet unfolds and renders the HTML of the composer to out, including its tag element its properties and its content.
// writeSnippet returns the id of the composer wtitten and an error if any.
// Rendering of sub-snippets may be called recursively maxDEEP times max to avoid infinite loop.
//
// An id can be setup up upfront (a) accessing any saved tag attribute within the snippet struct, or (b) within an html ick-tag attribute (for embedded snippet).
// In such cases these ids will be lost and the composer attributes will be overwritten with the unique id generated by the rendering process. This behaviour
// ensure that id will be unique event for multiple instanciations of the same composer.
// ids are generated by using the composer name (without "ick-" prefix) with a sequence number. sub-composer ids combine the id of the parent with the id of the sub-composer.
// if the component is not registered and so does't have a name, a unique id is generated.
//
// Composers may have no id on request. noid calling parameter must be set to true to render the composer without id.
// The special attribute noid can be defined within an ick-tag html to ensure the that no id will be rendered.
//
//  2. merge attributes
//     - some composer attributes may have already been pre-setup (a) within code instanciation, (b) during unfolding sub-composer instantiation
//     - the idea here is to merge thoses attributes with the default ones returned by the Tag method. pre-setup attributes are not overwritten in case of conflict.
//     - attributes associated with the tag are updated accordingly to ensure that composer attributes reflect last rendering.
//  3. unfold Content HTML
func writeSnippet(out io.Writer, composer HTMLComposer, ds *DataState, noid bool, parentvirtid string, seq int, deep int) (virtualid string, err error) {

	// ensure no infinite loop
	if deep > maxDEEP {
		err = fmt.Errorf("RenderHtmlComposer stopped at level %d. Too many recursive calls", deep)
		verbose.Println(verbose.ALERT, err.Error())
		return "", err
	}

	// Get a name for the composer
	var cmpname, ickcmpname string
	if entry := registry.LookupRegistryEntry(composer); entry != nil {
		ickcmpname = entry.TagName()
		cmpname = ickcmpname
		if left, has := strings.CutPrefix(ickcmpname, "ick-"); has {
			cmpname = left
		}
	} else if deep == 0 {
		ickcmpname = registry.GetUniqueId(ickcmpname)
	}
	composer.Tag().Attributes().SetAttribute("name", cmpname, true)

	// Define the virtual id and the id
	var id string
	virtualid = parentvirtid + "." + cmpname + strconv.Itoa(seq)
	_, noidatt := composer.Tag().Attributes().Attribute("noid")
	noid = noid || noidatt
	if !noid {
		id = virtualid
	}
	composer.Tag().Attributes().SetId(id)

	// verbose information
	verbose.Printf(verbose.INFO, "writting snippet (%s) vid=%q id=%q deep:%v\n", reflect.TypeOf(composer).String(), virtualid, id, deep)

	// render openingtag
	selfclosed, errtag := composer.Tag().RenderOpening(out)
	if selfclosed || errtag != nil {
		return virtualid, errtag
	}

	// Render Content HTML
	//
	// Unfold the body
	// _err = unfoldHTML(composer, out, []byte(bodytemplate), ds, deep)
	err = composer.WriteBody(out)
	if err != nil {
		return virtualid, err
	}

	// Render closingtag
	err = composer.Tag().RenderClosing(out)

	return virtualid, err
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

func (sw *stepway) startfield(i int) {
	sw.fieldat = i
	sw.fieldto = sw.fieldat
}
func (sw *stepway) openick(i int) {
	sw.processing = processing_ICKTAG
	sw.fieldat = i + 1
	sw.fieldto = i + 4
}
func (sw *stepway) closeick(i int) {
	sw.processing = processing_NONE
	sw.startfield(i + 2)
}

// unfoldHTML lookups for ick-tags in the _body htmlstring and unfold each of them recursively into _output.
//
// body is string combining HTML content and ick-tags. HTML content is transfered to the output without control and without changes.
// If no ick-tags are found, the output is a copy of the body
//
// ick-tags are autoclosing tags and should be in the form:
//
//	`<ick-{tagname} ...[boolattribute] ...[attribute=[']value[']]/>`
//
// if an error occurs the unfolding process stops immediatly.
// TODO: handle body within ickopening and ickclosing tags
func unfoldHTML(parent HTMLComposer, output io.Writer, body []byte, ds *DataState, deep int) (err error) {

	field := func(s stepway) []byte {
		return body[s.fieldat:s.fieldto]
	}

	walk := stepway{processing: processing_NONE}
	var ickname, aname, avalue string
	var bquote byte
	attrs := make(map[string]string, 0)

	nick := 0

	ilast := len(body) - 1
nextbyte:
	for i := 0; i <= ilast && err == nil; i++ {
		b := body[i]
		bclose_delim := string(body[i:mini(i+2, ilast+1)]) == "/>"
		bopen_delim := string(body[i:mini(i+5, ilast+1)]) == "<ick-"

		// decide what to do according to walk.processing and b value _</>*
		funfoldick := false
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
				output.Write(field(walk))
			case bopen_delim: // flush processed text field and start processing an ick-tage
				walk.fieldto = i
				output.Write(field(walk))
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
					err = errors.New("'<ick-' tag found without name")
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
				funfoldick = true

			default: // build component ick-tagname
				r, size := utf8.DecodeRune(body[i:mini(ilast+1, i+4)])
				if size != 0 && stringpattern.IsValidNameRune(r, false) {
					i += size - 1
					walk.fieldto = i
				} else {
					err = fmt.Errorf("invalid character found in ick-tagname: %q", string(body[walk.fieldat:i+1]))
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
					err = fmt.Errorf("= symbol found without attribute name: %q", ickname)
					break
				}
				walk.processing = processing_AVALUE
				walk.startfield(0)
				bquote = 0
			case bclose_delim: // process an ick-component
				walk.closeick(i)
				i += 2 - 1
				funfoldick = true

			default: // build attribute name
				r, size := utf8.DecodeRune(body[i:mini(ilast+1, i+4)])
				if size > 0 && stringpattern.IsValidNameRune(r, walk.fieldat == 0) {
					if walk.fieldat == 0 {
						walk.startfield(i)
					}
					i += size - 1
					walk.fieldto = i
				} else {
					err = fmt.Errorf("invalid character found in attribute name: %q", string(body[walk.fieldat:i+1]))
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
					err = fmt.Errorf("attribute with empty value: %q", string(body[walk.fieldat:i+1]))
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
					funfoldick = true
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

		if funfoldick {
			// DEBUG: unfolding embedded component
			fmt.Printf("level=%v -> unfolding embedded component: %s\n", deep, ickname)
			if warning := unfoldick(parent, output, ickname, attrs, ds, nick, deep); warning != nil {
				fmt.Printf("warning %q: %s\n", ickname, warning.Error())
				// DEBUG: fmt.Printf("embedded attributes: %v\n", attrs)
			}
			nick++
		}
	}
	return err
}

// unfoldick render the ick-component corresponding to _ickname and its unfolded _attrs.
// returns an error if the component or a sub component is not registered, or an embedded attribute type is unmannaged and it's value can't be parsed
// unfold sub components only if _deep is >= 0
func unfoldick(_parent HTMLComposer, _output io.Writer, _ickname string, _attrs map[string]string, _data *DataState, seq int, _deep int) (_err error) {
	// does this tag refer to a registered component ?
	htmlerr := ""
	regentry := registry.GetRegistryEntry(_ickname)
	if regentry.Component() != nil {

		// clone the registered component
		newref := reflect.New(reflect.TypeOf(regentry.Component()).Elem())
		newref.Elem().Set(reflect.ValueOf(regentry.Component()).Elem())
		newcmp := newref.Interface().(HTMLComposer)

		// process unfolded attributes, set value of ickcomponent field when name of attribute matches field name,
		// otherwise set unfolded attribute to the attribute of the component.
		for aname, avalue := range _attrs {
			//DEBUG:			fmt.Println(aname, newref.Elem().Type())
			_, found := newref.Elem().Type().FieldByName(aname)
			if !found {
				// this attribute is not a field of the componenent
				// keep it as is unless it is the class attribute, in this case, add the tokens
				newcmp.Tag().Attributes().SetAttribute(aname, avalue, true)
			} else {
				// feed data struct with the value
				field := newref.Elem().FieldByName(aname)
				if err := updateProperty(field, avalue); err != nil {
					htmlerr = fmt.Sprintf("<!-- unable to unfold %s component: %s for %s attribute -->", _ickname, err.Error(), aname)
					break
				}
			}
		}

		if htmlerr == "" && _deep >= 0 {
			// recursively unfold the component snippet
			// FIXME child
			newcmpid := ""
			newcmpid, _err = writeSnippet(_output, newcmp, _data, false, "", seq, _deep+1)

			// add it to the map of embedded components
			if newcmpid != "" && _parent != nil {
				_parent.Embed(newcmpid, newcmp)
			}
		}

	} else {
		htmlerr = fmt.Sprintf("<!-- unable to unfold unregistered %s component -->", _ickname)
	}

	if htmlerr != "" {
		_output.Write([]byte(htmlerr))
		_err = errors.New(htmlerr)
	}

	return _err
}

// updateProperty updates _cprop with the _value trying to convert the _value to the type of _cprop
// returns an error if _cprop type is unmannaged and it's value can't be parsed
func updateProperty(_cprop reflect.Value, _value string) (_erra error) {
	switch _cprop.Type().String() {
	case "time.Duration":
		var d time.Duration
		d, _erra = time.ParseDuration(_value)
		if _erra == nil {
			_cprop.SetInt(int64(d))
		}
	case "*url.URL":
		uu, err := url.Parse(_value)
		if err != nil {
			_erra = err
			break
		}
		_cprop.Set(reflect.ValueOf(uu))

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
		case reflect.Bool:
			f := true
			if s := strings.ToLower(strings.Trim(_value, " ")); s == "false" || s == "0" {
				f = false
			}
			_cprop.SetBool(f)

		// TODO: handle other data types
		default:
			return fmt.Errorf("unmanaged type %s", _cprop.Type().String()) // _cprop.Kind().String()
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
