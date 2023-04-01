package html

import (
	"sort"
	"strconv"
	"strings"

	"github.com/sunraylab/icecake/internal/helper"
)

// String encapsulates a known safe String document fragment.
// It should not be used for String from a third-party, or String with
// unclosed tags or comments.
//
// Use of this type presents a security risk:
// the encapsulated content should come from a trusted source,
// as it will be included verbatim in the output.
type String string

type SnippetTemplate struct {
	// The tagname used to render the html container element of this composer.
	// If tagname returns an empety string, the rendering does not generates the container element,
	// in such case snippet's attributes are ignored.
	TagName String

	Attributes string

	// Body returns the html template used to generate the content inside the container html element.
	Body String
}

type DataState struct {
	//Id   string // the id of the current processing component
	//Me   any    // the current processing component, should embedd an HtmlSnippet
	Page any // the current ick page, can be nil
	App  any // the current ick App, can be nil
}

// HTMLSnippet enables creation of simple or complex html string based on
// an original templating system allowing embedding of other snippets.
// HTMLSnippet output is an html element:
//
//	<tagname [attributes]>[body]</tagname>
//
// It is common to embed a HTMLSnippet into a struct to define an html component.
type HTMLSnippet struct {
	TagName  String            // optional TagName
	Body     String            // optional Body
	attrs    map[string]String // map of all attributes of any type
	embedded map[string]any    // instantiated embedded objects
}

// Id is an htmlComposer Interface
func (s *HTMLSnippet) Id() string {
	s.makemap()
	id := s.attrs["01id"]
	return string(id)
}

// SetId sets or overwrites the id attribute of the html snippet
func (s *HTMLSnippet) SetId(_id String) *HTMLSnippet {
	s.makemap()
	s.attrs["01id"] = _id
	return s
}

// ResetClasses replaces any existing classes with _clist to the class attribute
// _clist is parsed simply
func (s *HTMLSnippet) ResetClasses(_clist String) *HTMLSnippet {
	s.makemap()
	n := ""
	f := strings.Fields(string(_clist))
	for _, c := range f {
		if c != "" {
			// TODO: check validity of the class name pattern
			n += c + " "
		}
	}
	n = strings.TrimRight(n, " ")
	if n == "" {
		delete(s.attrs, "03class")
	} else {
		s.attrs["03class"] = String(n)
	}
	return s
}

// AddClasses add c classes to the class attribute
// any duplicate
func (s *HTMLSnippet) SetClasses(list String) *HTMLSnippet {
	s.makemap()
	last := s.attrs["03class"]
	new := string(last)
	clsf := strings.Fields(string(last))
	addf := strings.Fields(string(list))
nexta:
	for _, addc := range addf {
		if addc != "" {
			// TODO: check validity of the class name pattern

			for _, cls := range clsf {
				if cls == addc {
					continue nexta
				}
			}
			new += " " + addc
		}
	}
	new = strings.TrimLeft(new, " ")
	if new != "" {
		s.attrs["03class"] = String(new)
	}
	return s
}

// RemoveClass removes class c within the value of "class" attribute.
// Does nothing if c did not exist.
func (s *HTMLSnippet) RemoveClass(c string) *HTMLSnippet {
	s.makemap()
	last := s.attrs["03class"]
	new := ""
	clsf := strings.Fields(string(last))
	for _, cls := range clsf {
		if cls != c {
			new += cls + " "
		}
	}
	new = strings.TrimRight(new, " ")
	s.attrs["03class"] = String(new)
	return s
}

// SwitchClass removes class _remove within the value of "class" attribute and set the _new one
// Does nothing if c did not exist.
func (s *HTMLSnippet) SwitchClass(_remove string, _new String) *HTMLSnippet {
	s.RemoveClass(_remove)
	s.SetClasses(_new)
	return s
}

func (s *HTMLSnippet) SetStyle(style String) *HTMLSnippet {
	// TODO: check style validity
	s.makemap()
	s.attrs["04style"] = style
	return s
}

func (s *HTMLSnippet) SetTabIndex(idx uint) *HTMLSnippet {
	s.makemap()
	s.attrs["02tabIndex"] = String(strconv.Itoa(int(idx)))
	return s
}

// TODO: find a way to avoid overwrite parameter
func (s *HTMLSnippet) SetAttribute(key string, value String, overwrite bool) {
	s.makemap()
	key = strings.Trim(key, " ")
	switch strings.ToLower(key) {
	case "id":
		_, found := s.attrs["01id"]
		if !found || overwrite {
			s.SetId(value)
		}
	case "tabindex":
		_, found := s.attrs["02tabIndex"]
		if !found || overwrite {
			idx, _ := strconv.Atoi(string(value))
			s.SetTabIndex(uint(idx))
		}
	case "class":
		if overwrite {
			s.ResetClasses(value)
		} else if value != "" {
			s.SetClasses(value)
		}
	case "style":
		// TODO: handle style update to not overwrite
		_, found := s.attrs["04style"]
		if !found || overwrite {
			s.SetStyle(value)
		}
	default:
		_, found := s.attrs["05"+key]
		if !found || overwrite {
			s.attrs["05"+key] = value
		}
	}
}

func (s *HTMLSnippet) RemoveAttribute(key string) *HTMLSnippet {
	s.makemap()
	key = strings.Trim(key, " ")
	switch strings.ToLower(key) {
	case "id":
		delete(s.attrs, "01id")
	case "tabindex":
		delete(s.attrs, "02tabIndex")
	case "class":
		delete(s.attrs, "03class")
	case "style":
		delete(s.attrs, "04style")
	default:
		delete(s.attrs, "05"+key)
	}
	return s
}

// True set the boolean _key attribute in the list of attributes.
// does nothing if the key is id, style or class
func (s *HTMLSnippet) SetTrue(key string) *HTMLSnippet {
	s.SetAttribute(key, "", true)
	return s
}

// False unset the boolean _key attribute in the list of attributes.
// does nothing if the key is id, style or class
func (s *HTMLSnippet) SetFalse(key string) *HTMLSnippet {
	s.RemoveAttribute(key)
	return s
}

// Template is an interface implementation of HtmlComposer
func (s HTMLSnippet) Template(*DataState) (_t SnippetTemplate) {
	_t.TagName = s.TagName
	_t.Body = s.Body
	return
}

// Attributes returns the formated list of attributes used to generate the container element.
// always sorted the same way : 1.id 2.tabindex 3.class 4.style 5. other-alpha
// Attributes is an interface implementation of HtmlComposer
func (s HTMLSnippet) Attributes() String {
	s.makemap()
	if len(s.attrs) == 0 {
		return ""
	}

	html := ""
	sorted := make([]string, 0, len(s.attrs))
	for k := range s.attrs {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	for _, kx := range sorted {
		v := s.attrs[kx]
		k := kx[2:]
		html += k
		sv := stringifyValue(string(v))
		if len(sv) > 0 {
			html += "=" + sv
		}
		html += " "
	}
	html = strings.TrimRight(html, " ")
	return String(html)
}

func (s *HTMLSnippet) Embed(id string, cmp any) {
	id = helper.Normalize(id)
	if s.embedded == nil {
		s.embedded = make(map[string]any, 1)
	}
	s.embedded[id] = cmp
	// DEBUG: fmt.Printf("embedding %q(%v) into %s\n", id, reflect.TypeOf(cmp).String(), s.Id())
}

func (s HTMLSnippet) Embedded() map[string]any {
	return s.embedded
}

/******************************************************************************
 * PRIVATE
 *****************************************************************************/

func (s *HTMLSnippet) makemap() {
	if s.attrs == nil {
		s.attrs = make(map[string]String)
	}
}

// stringifyValue returns an empty string if v is empty or "false".
// returns an unquoted string if v can be converted in float.
func stringifyValue(v string) string {
	tv := strings.Trim(v, " ")
	lv := strings.ToLower(tv)
	if len(tv) == 0 || lv == "false" {
		return ""
	}

	if lv == "true" {
		return lv
	}

	_, err := strconv.ParseFloat(tv, 64)
	if err == nil {
		return tv
	}

	var delim string
	if strings.ContainsRune(tv, rune('"')) {
		delim = "'"
	} else {
		delim = "\""
	}
	return delim + v + delim

}
