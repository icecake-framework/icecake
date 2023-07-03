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
	"github.com/lolorenzo777/verbose"
)

// HTMLString represents a string document fragment, mixing standard HTML syntax with <ick-tags/> inline components.
// HTMLString implements HTMLcomposer interfaces and the string is rendered to an output stream with the icecake Rendering functions.
// It is part of the core icecake snippets.
type HTMLString struct {
	meta  RenderingMeta // Rendering MetaData
	bytes []byte
}

// Ensure HTMLString implements HTMLTagComposer interface
var _ HTMLComposer = (*HTMLString)(nil)

// ToHTML is the HTMLString factory allowing to convert a string into a new HTMLString reday for rendering.
// The string must contains safe string and can include icecake tags.
// TODO: ToHTML accept any types
func ToHTML(s string) *HTMLString {
	h := new(HTMLString)
	h.bytes = []byte(s)
	return h
}

// Meta provides a reference to the RenderingMeta object associated with this composer.
// This is required by the icecake rendering process.
func (h *HTMLString) Meta() *RenderingMeta {
	return &h.meta
}

// RenderContent writes the HTML string corresponding to the content of the HTML element.
// For an HTMLString snippet, RenderContent renders (unfold and generate HTML output) the internal string without enclosed tag.
// Use an HTMLSnippet snippet to renders the string inside an enclosed tag.
func (h *HTMLString) RenderContent(out io.Writer) error {
	return renderHTML(out, h, *h)
}

// String returns a copy of the content
func (h HTMLString) Bytes() []byte {
	b := h.bytes
	return b
}

// String returns the content
func (s HTMLString) IsEmpty() bool {
	return s.bytes == nil || len(s.bytes) == 0
}

// RenderHTML lookups for ick-tags in the htmlstring and unfold each of them into out.
//
// htmlstring is a string combining usual HTML text and ick-tags. HTML content is transfered to the output without control and without changes.
// ick-tags are autoclosing tags and should be in the form:
//
//	`<ick-{tagname} ...[boolattribute] ...[attribute=[']value[']]/>`
//
// If no ick-tags are found, the output is a copy of the htmlstring.
// An error is returned if the HTML string is not well formatted and the rendering process fails.
// The rendering process renders an HTML comment in the following cases:
//   - If the HTML string contains ick-tag but the ick-tagname does not correspond to a Registered composer,
//   - If the HTML string contains ick-tag with attributes but one value of these attribute is of a bad type,
//
// TODO: implement rendering of ick-tag with content inside
func renderHTML(out io.Writer, parent HTMLComposer, htmlstr HTMLString) (err error) {

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

	startfield := func(sw *stepway, i int) {
		sw.fieldat = i
		sw.fieldto = sw.fieldat
	}
	openick := func(sw *stepway, i int) {
		sw.processing = processing_ICKTAG
		sw.fieldat = i + 1
		sw.fieldto = i + 4
	}
	closeick := func(sw *stepway, i int) {
		sw.processing = processing_NONE
		startfield(sw, i+2)
	}

	htmlstring := htmlstr.bytes

	field := func(s stepway) []byte {
		return htmlstring[s.fieldat:s.fieldto]
	}

	walk := stepway{processing: processing_NONE}
	var ickname, aname, avalue string
	var bquote byte
	attrs := make(AttributeMap, 0)

	nick := 0
	ilast := len(htmlstring) - 1
nextbyte:
	for i := 0; i <= ilast && err == nil; i++ {
		b := htmlstring[i]
		bautoclose_delim := string(htmlstring[i:mini(i+2, ilast+1)]) == "/>"
		bopen_delim := string(htmlstring[i:mini(i+5, ilast+1)]) == "<ick-"

		// decide what to do according to walk.processing and b value _</>*
		funfoldick := false
		switch walk.processing {
		case processing_NONE:
			switch {
			case bopen_delim: // start processing an ick-tage
				openick(&walk, i)
				i += 5 - 1
			default: // start processing a text field
				walk.processing = processing_TXT
				startfield(&walk, i)
			}

		case processing_TXT:
			switch {
			case i == ilast: // flush processed text field and exit
				walk.fieldto = ilast + 1
				out.Write(field(walk))
			case bopen_delim: // flush processed text field and start processing an ick-tage
				walk.fieldto = i
				out.Write(field(walk))
				openick(&walk, i)
				i += 5 - 1
			default: // extend the text field
				walk.fieldto = i
			}

		case processing_ICKTAG:
			if b == ' ' || bautoclose_delim { // record component tagname
				walk.fieldto = i
				ickname = string(field(walk))
				ickname = strings.ToLower(ickname)
				if ickname == "ick-" {
					return ErrNameMissing
				}
				aname = ""
				avalue = ""
				attrs = make(AttributeMap, 0)
			}
			switch {
			case b == ' ': // look for another aname
				walk.processing = processing_ANAME
				startfield(&walk, 0)
			case bautoclose_delim: // process a single ick-component
				closeick(&walk, i)
				i += 2 - 1
				funfoldick = true

			default: // build component ick-tagname
				r, size := utf8.DecodeRune(htmlstring[i:mini(ilast+1, i+4)])
				if size != 0 && stringpattern.IsValidNameRune(r, false) {
					i += size - 1
					walk.fieldto = i
				} else {
					return &IckTagNameError{TagName: string(htmlstring[walk.fieldat : i+1]), Message: "invalid character found in tagname"}
				}
			}

		case processing_ANAME:
			switch {
			case (b == ' ' || b == '\n' || b == '\t') && walk.fieldat == 0: // trim left spaces and \n
				continue nextbyte
			case (b == ' ' || b == '=' || b == '\n' || b == '\t' || bautoclose_delim) && walk.fieldat > 0: // get and save aname
				walk.fieldto = i
				aname = string(field(walk))
				attrs[aname] = ""
			}

			switch {
			case b == ' ': // look for another aname
				aname = ""
				walk.processing = processing_ANAME
				startfield(&walk, 0)
			case b == '=': // look for a value
				if aname == "" {
					return &IckTagNameError{TagName: ickname, Message: "missing attribute's name before '='"}
				}
				walk.processing = processing_AVALUE
				startfield(&walk, 0)
				bquote = 0
			case bautoclose_delim: // process an ick-component
				closeick(&walk, i)
				i += 2 - 1
				funfoldick = true

			default: // build attribute name
				r, size := utf8.DecodeRune(htmlstring[i:mini(ilast+1, i+4)])
				if size > 0 && stringpattern.IsValidNameRune(r, walk.fieldat == 0) {
					if walk.fieldat == 0 {
						startfield(&walk, i)
					}
					i += size - 1
					walk.fieldto = i
				} else {
					return &IckTagNameError{TagName: ickname, Message: fmt.Sprintf("invalid attribute's name character %q", string(htmlstring[walk.fieldat:i+1]))}
				}
			}

		case processing_AVALUE:
			if bquote == 0 { // don't know yet if a quoted or unquoted value
				switch {
				case b == ' ': // trim left spaces
				case b == '"' || b == '\'': // start a quoted value
					bquote = b
					startfield(&walk, i+1)
				case bautoclose_delim: // empty value
					return &IckTagNameError{TagName: ickname, Message: fmt.Sprintf("missing attribute's value %q", string(htmlstring[walk.fieldat:i+1]))}
				default: // start unquoted value
					bquote = 1
					startfield(&walk, i)
				}
				break
			}

			switch {
			case bquote == 1 && (b == ' ' || bautoclose_delim): // process unquoted value
				walk.fieldto = i
				avalue = string(field(walk))
				attrs[aname] = trimfirstvalue(avalue)
				switch {
				case bautoclose_delim: // process an ick-tagname
					closeick(&walk, i)
					i += 2 - 1
					funfoldick = true
				default: // look for another aname
					walk.processing = processing_ANAME
					startfield(&walk, 0)
				}
			case bquote != 1 && b == bquote: // process a quoted value
				walk.fieldto = i
				avalue = string(field(walk))
				attrs[aname] = avalue
				walk.processing = processing_ANAME
				startfield(&walk, 0)
			default: // extend field value
				walk.fieldto = i
			}
		}

		if funfoldick {
			errunf := unfoldick(parent, out, ickname, attrs, nick)
			if errunf != nil {
				if errors.Is(errunf, ErrTooManyRecursiveRendering) {
					verbose.Printf(verbose.ALERT, errunf.Error())
					return errunf
				}
				verbose.Printf(verbose.WARNING, errunf.Error())
			}
			nick++
		}
	}
	return err
}

// unfoldick instanciates and unfolds the ick-component corresponding to ickname.
// The rendering process renders an HTML comment and return an error in the following cases:
//   - If the HTML string contains ick-tag but the ick-tagname does not correspond to a Registered composer,
//   - If the HTML string contains ick-tag with attributes but one value of these attribute is of a bad type,
func unfoldick(parent HTMLComposer, out io.Writer, ickname string, ickattrs AttributeMap, seq int) (err error) {

	verbose.Debug("unfolding component %q", ickname)

	// does this tag refer to a registered component ?
	regentry := registry.GetRegistryEntry(ickname)
	if regentry.Component() != nil {

		// clone the registered snippet (a composer)
		newref := reflect.New(reflect.TypeOf(regentry.Component()).Elem())

		// init all values with the one in the template
		// warning: pointers are copied by ref and not by value
		// newref.Elem().Set(reflect.ValueOf(regentry.Component()).Elem())

		// get the interface of the new snippet (a composer)
		newcmp := newref.Interface().(HTMLComposer)

		// process unfolded attributes, set value of ickcomponent field when name of attribute matches field name,
		// otherwise set unfolded attribute to the attribute of the component.
		for ickattname, ickattvalue := range ickattrs {
			verbose.Debug(ickattname, newref.Elem().Type())
			_, isCmpProperty := newref.Elem().Type().FieldByName(ickattname)
			if isCmpProperty {
				// feed cmp struct property with the ickattvalue
				prop := newref.Elem().FieldByName(ickattname)
				if erru := updateproperty(prop, ickattvalue); erru != nil {
					err = &IckTagNameError{TagName: ickname, Message: fmt.Sprintf("%q attribute: %s", ickattname, erru)}
					break
				}
			} else {
				// this attribute is not a field of the componenent
				// keep it as is unless it is the class attribute, in this case, add the attribute
				if tagbuilder, isbuilder := newref.Interface().(TagBuilder); isbuilder && tagbuilder != nil {
					tagbuilder.SetAttribute(ickattname, ickattvalue)
				} else {
					err = &IckTagNameError{TagName: ickname, Message: fmt.Sprintf("%q attribute: not a component property and not assignable to the composer.", ickattname)}
					break
				}
			}
		}

		// recursively unfold the component snippet
		if err == nil {
			err = render(out, parent, newcmp)
		}

	} else {
		err = &IckTagNameError{TagName: ickname, Message: "unregistered ick-tagname"}
	}

	if err != nil {
		WriteString(out, "<!--", err.Error(), "-->")
	}

	return err
}

// updateproperty updates prop with the value trying to convert the value to the type of prop.
// Returns an error if prop's type is unmannaged and its value can't be extracted.
func updateproperty(prop reflect.Value, value string) (err error) {
	switch prop.Type().String() {
	case "html.HTMLString":
		s := string(ToHTML(value).Bytes())
		prop.Set(reflect.ValueOf(s))
	case "time.Duration":
		var d time.Duration
		d, err = time.ParseDuration(value)
		if err == nil {
			prop.SetInt(int64(d))
		}
	case "*url.URL":
		pu, erru := url.Parse(value)
		if erru != nil {
			err = erru
			break
		}
		prop.Set(reflect.ValueOf(pu))

	default:
		switch prop.Kind() {
		case reflect.String:
			prop.SetString(value)
		case reflect.Int:
			var i int
			i, err = strconv.Atoi(value)
			if err == nil {
				prop.SetInt(int64(i))
			}
		case reflect.Float64:
			var f float64
			f, err = strconv.ParseFloat(value, 64)
			if err == nil {
				prop.SetFloat(f)
			}
		case reflect.Bool:
			f := true
			if s := strings.ToLower(strings.Trim(value, " ")); s == "false" || s == "0" {
				f = false
			}
			prop.SetBool(f)
		default:
			err = fmt.Errorf("unmanaged type %s", prop.Type().String())
		}
	}
	return err
}

// trimfirstvalue returns string with endings blanks trimed but keeping white space inside quotes.
// If str does not have quotes ( " or ' ) the returned string is truncated at the first white space found.
func trimfirstvalue(str string) string {
	trimspaces := strings.Trim(str, " ")

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
