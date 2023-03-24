package ick

import (
	"sort"
	"strconv"
	"strings"
)

// HtmlSnippet enables creation of simple or complex html string based on
// an original templating system allowing embedding of other snippets.
// HtmlSnippet output is an html element:
//
//	<tagname [attributes]>[body]</tagname>
//
// It is common to embed a HtmlSnippet into a struct to define an html component.
type HtmlSnippet struct {
	attrs map[string]HTML // map of all attributes whatever the type
}

func (s *HtmlSnippet) Id() HTML {
	s.makemap()
	id := s.attrs["01id"]
	return HTML(id)
}

// SetId sets or overwrites the id attribute of the html snippet
func (s *HtmlSnippet) SetId(_id HTML) *HtmlSnippet {
	s.makemap()
	s.attrs["01id"] = _id
	return s
}

// NewClasses replace any existing classes with c to the class attribute
// c is parsed simply
// TODO: check valididty of _c
func (s *HtmlSnippet) ResetClasses(clist HTML) *HtmlSnippet {
	s.makemap()
	n := ""
	f := strings.Fields(string(clist))
	for _, c := range f {
		if c != "" {
			// TODO check validity of class name
			n += c + " "
		}
	}
	n = strings.TrimRight(n, " ")
	if n == "" {
		delete(s.attrs, "03class")
	} else {
		s.attrs["03class"] = HTML(n)
	}
	return s
}

// AddClasses add c classes to the class attribute
// any duplicate
func (s *HtmlSnippet) SetClasses(list HTML) *HtmlSnippet {
	s.makemap()
	last := s.attrs["03class"]
	new := string(last)
	clsf := strings.Fields(string(last))
	addf := strings.Fields(string(list))
nexta:
	for _, addc := range addf {
		if addc != "" {
			// TODO check validity of class name

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
		s.attrs["03class"] = HTML(new)
	}
	return s
}

// RemoveClass removes class c within the value of "class" attribute.
// Does nothing if c did not exist.
func (s *HtmlSnippet) RemoveClass(c string) *HtmlSnippet {
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
	s.attrs["03class"] = HTML(new)
	return s
}

func (s *HtmlSnippet) SetStyle(style HTML) *HtmlSnippet {
	// TODO: check style validity
	s.makemap()
	s.attrs["04style"] = style
	return s
}

func (s *HtmlSnippet) SetTabIndex(idx uint) *HtmlSnippet {
	s.makemap()
	s.attrs["02tabIndex"] = HTML(strconv.Itoa(int(idx)))
	return s
}

func (s *HtmlSnippet) SetAttribute(key string, value HTML, overwrite bool) {
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
		// TODO handle update style to not overwrite
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

func (s *HtmlSnippet) RemoveAttribute(key string) *HtmlSnippet {
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
func (s *HtmlSnippet) SetTrue(key string) *HtmlSnippet {
	s.SetAttribute(key, "", true)
	return s
}

// False unset the boolean _key attribute in the list of attributes.
// does nothing if the key is id, style or class
func (s *HtmlSnippet) SetFalse(key string) *HtmlSnippet {
	s.RemoveAttribute(key)
	return s
}

// TagName implements the default HtmlComposer interface, a composer without container.
// func (s HtmlSnippet) TagName() HTML {
// 	return ""
// }

// Body implements the default HtmlComposer interface, an empty string.
// func (s HtmlSnippet) Body() HTML {
// 	return ""
// }

func (s HtmlSnippet) Template(*DataState) SnippetTemplate {
	return SnippetTemplate{}
}

// Attributes returns the formated list of attributes used to generate the container element.
// always sorted the same way : 1.id 2.tabindex 3.class 4.style 5. other-alpha
func (s HtmlSnippet) Attributes() HTML {
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
	return HTML(html)
}

/******************************************************************************
 * PRIVATE
 *****************************************************************************/

func (s *HtmlSnippet) makemap() {
	if s.attrs == nil {
		s.attrs = make(map[string]HTML)
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
