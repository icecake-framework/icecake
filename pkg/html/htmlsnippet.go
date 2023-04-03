package html

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/sunraylab/icecake/internal/helper"
	"github.com/sunraylab/icecake/pkg/registry"
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
	attrs    map[string]string // map of all attributes of any type
	embedded map[string]any    // instantiated embedded objects
}

// Id is an htmlComposer Interface
func (s *HTMLSnippet) Id() string {
	s.makemap()
	id := s.attrs["01id"]
	return string(id)
}

// TabIndex returns the TabIndex attribute
func (s *HTMLSnippet) TabIndex() int {
	s.makemap()
	sidx := s.attrs["02tabIndex"]
	idx, _ := strconv.Atoi(string(sidx))
	return idx
}

// Classes returns the class attribute
func (s *HTMLSnippet) Classes() string {
	s.makemap()
	str := s.attrs["03class"]
	return string(str)
}

// Classes returns the class attribute
func (s *HTMLSnippet) HasClass(_class string) bool {
	_class = strings.Trim(_class, " ")
	if _class == "" {
		return false
	}
	s.makemap()
	actual := s.attrs["03class"]
	actualf := strings.Fields(string(actual))
	for _, actualc := range actualf {
		if actualc == _class {
			return true
		}
	}
	return false
}

// Style returns the style attribute
func (s *HTMLSnippet) Style() string {
	s.makemap()
	str := s.attrs["04style"]
	return string(str)
}

// IsDisabled returns the boolean attribute
func (s *HTMLSnippet) IsDisabled() bool {
	s.makemap()
	str, found := s.attrs["05disabled"]
	if !found || strings.ToLower(string(str)) == "false" || str == "0" {
		return false
	}
	return true
}

// Attributes returns the formated list of attributes used to generate the container element.
// always sorted the same way : 1.id 2.tabindex 3.class 4.style 5. other-alpha
// Attributes is an interface implementation of HtmlComposer
func (s *HTMLSnippet) Attributes() String {
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

// SetId sets or overwrites the id attribute of the html snippet
func (s *HTMLSnippet) SetId(_id String) *HTMLSnippet {
	s.makemap()
	s.attrs["01id"] = string(_id)
	return s
}

func (s *HTMLSnippet) SetUniqueId(_prefix string) *HTMLSnippet {
	s.makemap()
	s.attrs["01id"] = registry.GetUniqueId(_prefix)
	return s
}

func (s *HTMLSnippet) SetTabIndex(idx int) *HTMLSnippet {
	s.makemap()
	s.attrs["02tabIndex"] = strconv.Itoa(idx)
	return s
}

func (s *HTMLSnippet) SetDisabled(_f bool) *HTMLSnippet {
	s.makemap()
	if _f {
		s.attrs["05disabled"] = ""
	} else {
		delete(s.attrs, "05disabled")
	}
	return s
}

// ResetClasses replaces any existing classes with _clist to the class attribute
// _clist must contains strings separated by spaces.
// All classes are removed if _clist is empty.
// TODO: check validity of the class name pattern
func (s *HTMLSnippet) ResetClasses(_list String) *HTMLSnippet {
	s.makemap()
	n := ""
	f := strings.Fields(string(_list))
	for _, c := range f {
		if c != "" {
			n += c + " "
		}
	}
	n = strings.TrimRight(n, " ")
	if n == "" {
		delete(s.attrs, "03class")
	} else {
		s.attrs["03class"] = n
	}
	return s
}

// SetClasses adds the _list of classes to the class attribute
// duplicate are not inserted twice.
// TODO: check validity of the class name pattern
func (s *HTMLSnippet) SetClasses(_list String) *HTMLSnippet {
	s.makemap()
	actual := s.attrs["03class"]
	new := string(actual)
	actualf := strings.Fields(string(actual))
	listf := strings.Fields(string(_list))
nextf:
	for _, listc := range listf {
		if listc != "" {
			for _, actualc := range actualf {
				if actualc == listc {
					continue nextf
				}
			}
			new += " " + listc
		}
	}
	new = strings.TrimLeft(new, " ")
	if new != "" {
		s.attrs["03class"] = new
	}
	return s
}

// RemoveClasses removes any class in _list from the "class" attribute.
// Does nothing if c did not exist.
func (s *HTMLSnippet) RemoveClasses(_list string) *HTMLSnippet {
	s.makemap()
	actual := s.attrs["03class"]
	new := ""
	actualf := strings.Fields(string(actual))
	listf := strings.Fields(string(_list))
nexta:
	for _, actualc := range actualf {
		for _, listc := range listf {
			if actualc == listc {
				continue nexta
			}
		}
		new += " " + actualc
	}
	new = strings.TrimRight(new, " ")
	s.attrs["03class"] = new
	return s
}

// SwitchClasses removes _remove classes and set the _new one
// Does nothing if c did not exist.
func (s *HTMLSnippet) SwitchClasses(_remove string, _new String) *HTMLSnippet {
	s.RemoveClasses(_remove)
	s.SetClasses(_new)
	return s
}

// TODO: check style validity
func (s *HTMLSnippet) SetStyle(style String) *HTMLSnippet {
	s.makemap()
	s.attrs["04style"] = string(style)
	return s
}

// CreateAttribute create an attribute and set its value.
// If the attribute already exists nothing is done.
func (s *HTMLSnippet) CreateAttribute(_key string, _value any) {
	s.setAttribute(_key, _value, false)
}

// SetAttribute create an attribute and set its value.
// If the attribute already exists the value is updated.
func (s *HTMLSnippet) SetAttribute(_key string, _value any) {
	s.setAttribute(_key, _value, true)
}

func (s *HTMLSnippet) Attribute(_key string) (string, bool) {
	_key = strings.Trim(_key, " ")
	v, found := s.attrs["05"+_key]
	return string(v), found
}

func (s *HTMLSnippet) RemoveAttribute(_key string) *HTMLSnippet {
	s.makemap()
	_key = strings.Trim(_key, " ")
	switch strings.ToLower(_key) {
	case "id":
		delete(s.attrs, "01id")
	case "tabindex":
		delete(s.attrs, "02tabIndex")
	case "class":
		delete(s.attrs, "03class")
	case "style":
		delete(s.attrs, "04style")
	default:
		delete(s.attrs, "05"+_key)
	}
	return s
}

// ToggleAttribute set the attribute if unset and unset it if set.
func (s *HTMLSnippet) ToggleAttribute(_key string) {
	s.makemap()
	_key = strings.Trim(_key, " ")
	_, found := s.attrs["05"+_key]
	if !found {
		s.attrs["05"+_key] = ""
	} else {
		delete(s.attrs, "05"+_key)
	}
}

// TODO: find a way to avoid overwrite parameter
func (s *HTMLSnippet) setAttribute(key string, value any, overwrite bool) error {
	s.makemap()
	key = strings.Trim(key, " ")
	switch strings.ToLower(key) {
	case "id":
		_, found := s.attrs["01id"]
		if !found || overwrite {
			switch v := value.(type) {
			case string:
				s.SetId(String(v))
			case String:
				s.SetId(v)
			default:
				return errors.New("wrong value type for id")
			}
		}
	case "tabindex":
		_, found := s.attrs["02tabIndex"]
		if !found || overwrite {
			switch v := value.(type) {
			case string:
				idx, _ := strconv.Atoi(string(v))
				s.SetTabIndex(idx)
			case String:
				idx, _ := strconv.Atoi(string(v))
				s.SetTabIndex(idx)
			case int:
				s.SetTabIndex(v)
			case uint:
				s.SetTabIndex(int(v))
			case float32:
				s.SetTabIndex(int(v))
			case float64:
				s.SetTabIndex(int(v))
			default:
				return errors.New("wrong value type for tabindex")
			}
		}
	case "class":
		var lst String
		switch v := value.(type) {
		case string:
			lst = String(v)
		case String:
			lst = v
		default:
			return errors.New("wrong value type for class")
		}
		if overwrite {
			s.ResetClasses(lst)
		} else if value != "" {
			s.SetClasses(lst)
		}
	case "style":
		// TODO: handle style update to not overwrite
		_, found := s.attrs["04style"]
		if !found || overwrite {
			var style String
			switch v := value.(type) {
			case string:
				style = String(v)
			case String:
				style = v
			default:
				return errors.New("wrong value type for class")
			}
			s.SetStyle(style)
		}
	default:
		_, found := s.attrs["05"+key]
		if !found || overwrite {
			var strv string
			switch v := value.(type) {
			case string:
				strv = v
			case String:
				strv = string(v)
			case bool:
				if v {
					strv = ""
				} else {
					delete(s.attrs, "05"+key)
					break
				}
			case int, uint, float32, float64:
				strv = fmt.Sprintf("%v", v)
			default:
				return errors.New("wrong value type for " + key)
			}
			s.attrs["05"+key] = strv
		}
	}
	return nil
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

// Template is an interface implementation of HtmlComposer
func (s HTMLSnippet) Template(*DataState) (_t SnippetTemplate) {
	_t.TagName = s.TagName
	_t.Body = s.Body
	return
}

/******************************************************************************
 * PRIVATE
 *****************************************************************************/

func (s *HTMLSnippet) makemap() {
	if s.attrs == nil {
		s.attrs = make(map[string]string)
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
