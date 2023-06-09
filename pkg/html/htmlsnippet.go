package html

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/icecake-framework/icecake/internal/helper"
	"github.com/icecake-framework/icecake/pkg/registry"
)

// String encapsulates a known safe String document fragment.
// It should not be used for String from a third-party, or String with
// unclosed tags or comments.
//
// Use of this type presents a security risk:
// the encapsulated content should come from a trusted source,
// as it will be included verbatim in the output.
type String string

// Tag
type Tag struct {
	TagName        String // optional TagName
	TagSelfClosing bool
}

// SnippetTemplate
type SnippetTemplate struct {
	// Tag is used to render the html container element of this composer.
	// If tagname returns an empety string, the rendering does not generates the container element,
	// in such case snippet's attributes are ignored.
	Tag

	// if a TagName is defined, Attributes will be rendered with the html container element of this composer.
	Attributes String

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
	Tag                        // optional TagName
	Body     String            // optional Body
	attrs    map[string]string // map of all attributes of any type
	embedded map[string]any    // instantiated embedded objects
}

// Id is an htmlComposer Interface
func (_snippet *HTMLSnippet) Id() string {
	_snippet.makemap()
	id := _snippet.attrs["01id"]
	return string(id)
}

// TabIndex returns the TabIndex attribute
func (_snippet *HTMLSnippet) TabIndex() int {
	_snippet.makemap()
	sidx := _snippet.attrs["02tabIndex"]
	idx, _ := strconv.Atoi(string(sidx))
	return idx
}

// Classes returns the class attribute
func (_snippet *HTMLSnippet) Classes() string {
	_snippet.makemap()
	str := _snippet.attrs["03class"]
	return string(str)
}

// Classes returns the class attribute
func (_snippet *HTMLSnippet) HasClass(_class string) bool {
	_class = strings.Trim(_class, " ")
	if _class == "" {
		return false
	}
	_snippet.makemap()
	actual := _snippet.attrs["03class"]
	actualf := strings.Fields(string(actual))
	for _, actualc := range actualf {
		if actualc == _class {
			return true
		}
	}
	return false
}

// Style returns the style attribute
func (_snippet *HTMLSnippet) Style() string {
	_snippet.makemap()
	str := _snippet.attrs["04style"]
	return string(str)
}

// IsDisabled returns the boolean attribute
func (_snippet *HTMLSnippet) IsDisabled() bool {
	_snippet.makemap()
	str, found := _snippet.attrs["05disabled"]
	if !found || strings.ToLower(string(str)) == "false" || str == "0" {
		return false
	}
	return true
}

// Attributes returns the formated list of attributes used to generate the container element.
// always sorted the same way : 1.id 2.tabindex 3.class 4.style 5. other-alpha
// Attributes is an interface implementation of HtmlComposer
func (_snippet *HTMLSnippet) Attributes() String {
	_snippet.makemap()
	if len(_snippet.attrs) == 0 {
		return ""
	}

	html := ""
	sorted := make([]string, 0, len(_snippet.attrs))
	for k := range _snippet.attrs {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	for _, kx := range sorted {
		v := _snippet.attrs[kx]
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
func (_snippet *HTMLSnippet) SetId(_id String) *HTMLSnippet {
	_snippet.makemap()
	_snippet.attrs["01id"] = string(_id)
	return _snippet
}

// SetUniqueId sets or overwrites the id attribute of the html snippet
// generating a unique id starting with the _prefix.
// "ick-" is used to prefix the returned id if _prefix is empty.
func (_snippet *HTMLSnippet) SetUniqueId(_prefix string) *HTMLSnippet {
	_snippet.makemap()
	_snippet.attrs["01id"] = registry.GetUniqueId(_prefix)
	return _snippet
}

func (_snippet *HTMLSnippet) SetTabIndex(idx int) *HTMLSnippet {
	_snippet.makemap()
	_snippet.attrs["02tabIndex"] = strconv.Itoa(idx)
	return _snippet
}

func (_snippet *HTMLSnippet) SetDisabled(_f bool) *HTMLSnippet {
	_snippet.makemap()
	if _f {
		_snippet.attrs["05disabled"] = ""
	} else {
		delete(_snippet.attrs, "05disabled")
	}
	return _snippet
}

// ResetClasses replaces any existing classes with _clist to the class attribute
// _clist must contains strings separated by spaces.
// All classes are removed if _clist is empty.
// TODO: check validity of the class name pattern
func (_snippet *HTMLSnippet) ResetClasses(_list String) *HTMLSnippet {
	_snippet.makemap()
	n := ""
	f := strings.Fields(string(_list))
	for _, c := range f {
		if c != "" {
			n += c + " "
		}
	}
	n = strings.TrimRight(n, " ")
	if n == "" {
		delete(_snippet.attrs, "03class")
	} else {
		_snippet.attrs["03class"] = n
	}
	return _snippet
}

// SetClasses adds the _list of classes to the class attribute
// duplicate are not inserted twice.
// TODO: check validity of the class name pattern
func (_snippet *HTMLSnippet) SetClasses(_list String) *HTMLSnippet {
	_snippet.makemap()
	actual := _snippet.attrs["03class"]
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
		_snippet.attrs["03class"] = new
	}
	return _snippet
}

// SetClassesIf SetClasses if the _condition is true
func (_snippet *HTMLSnippet) SetClassesIf(_condition bool, _list String) *HTMLSnippet {
	if _condition {
		_snippet.SetClasses(_list)
	}
	return _snippet
}

// RemoveClasses removes any class in _list from the "class" attribute.
// Does nothing if c did not exist.
func (_snippet *HTMLSnippet) RemoveClasses(_list string) *HTMLSnippet {
	_snippet.makemap()
	actual := _snippet.attrs["03class"]
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
	_snippet.attrs["03class"] = new
	return _snippet
}

// SwitchClasses removes _remove classes and set the _new one
// Does nothing if c did not exist.
func (_snippet *HTMLSnippet) SwitchClasses(_remove string, _new String) *HTMLSnippet {
	_snippet.RemoveClasses(_remove)
	_snippet.SetClasses(_new)
	return _snippet
}

// TODO: check style validity
func (_snippet *HTMLSnippet) SetStyle(style String) *HTMLSnippet {
	_snippet.makemap()
	_snippet.attrs["04style"] = string(style)
	return _snippet
}

// CreateAttribute create an attribute and set its value.
// If the attribute already exists nothing is done.
func (_snippet *HTMLSnippet) CreateAttribute(_key string, _value any) {
	_snippet.setAttribute(_key, _value, false)
}

// SetAttribute create an attribute and set its value.
// If the attribute already exists the value is updated.
func (_snippet *HTMLSnippet) SetAttribute(_key string, _value any) {
	_snippet.setAttribute(_key, _value, true)
}

// SetAttributeIf SetAttribute if the _condition is true.
func (_snippet *HTMLSnippet) SetAttributeIf(_condition bool, _key string, _value any) *HTMLSnippet {
	if _condition {
		_snippet.setAttribute(_key, _value, true)
	}
	return _snippet
}

// Attribute returns the value of the attribute identified by _key.
// Returns false if the attribute does not exist.
func (_snippet *HTMLSnippet) Attribute(_key string) (string, bool) {
	_key = strings.Trim(_key, " ")
	v, found := _snippet.attrs["05"+_key]
	return string(v), found
}

// RemoveAttribute remove the the attribute identified by _key.
// Does nothing if the _key is not found.
func (_snippet *HTMLSnippet) RemoveAttribute(_key string) *HTMLSnippet {
	_snippet.makemap()
	_key = strings.Trim(_key, " ")
	switch strings.ToLower(_key) {
	case "id":
		delete(_snippet.attrs, "01id")
	case "tabindex":
		delete(_snippet.attrs, "02tabIndex")
	case "class":
		delete(_snippet.attrs, "03class")
	case "style":
		delete(_snippet.attrs, "04style")
	default:
		delete(_snippet.attrs, "05"+_key)
	}
	return _snippet
}

// ToggleAttribute toggles the boolean attribute _key. Sets it when unsetted, and unset whrn setted.
func (_snippet *HTMLSnippet) ToggleAttribute(_key string) {
	_snippet.makemap()
	_key = strings.Trim(_key, " ")
	_, found := _snippet.attrs["05"+_key]
	if !found {
		_snippet.attrs["05"+_key] = ""
	} else {
		delete(_snippet.attrs, "05"+_key)
	}
}

// TODO: find a way to avoid overwrite parameter
func (_snippet *HTMLSnippet) setAttribute(key string, value any, overwrite bool) error {
	_snippet.makemap()
	key = strings.Trim(key, " ")
	switch strings.ToLower(key) {
	case "id":
		_, found := _snippet.attrs["01id"]
		if !found || overwrite {
			switch v := value.(type) {
			case string:
				_snippet.SetId(String(v))
			case String:
				_snippet.SetId(v)
			default:
				return errors.New("wrong value type for id")
			}
		}
	case "tabindex":
		_, found := _snippet.attrs["02tabIndex"]
		if !found || overwrite {
			switch v := value.(type) {
			case string:
				idx, _ := strconv.Atoi(string(v))
				_snippet.SetTabIndex(idx)
			case String:
				idx, _ := strconv.Atoi(string(v))
				_snippet.SetTabIndex(idx)
			case int:
				_snippet.SetTabIndex(v)
			case uint:
				_snippet.SetTabIndex(int(v))
			case float32:
				_snippet.SetTabIndex(int(v))
			case float64:
				_snippet.SetTabIndex(int(v))
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
			_snippet.ResetClasses(lst)
		} else if value != "" {
			_snippet.SetClasses(lst)
		}
	case "style":
		// TODO: handle style update to not overwrite
		_, found := _snippet.attrs["04style"]
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
			_snippet.SetStyle(style)
		}
	default:
		_, found := _snippet.attrs["05"+key]
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
					delete(_snippet.attrs, "05"+key)
					break
				}
			case int, uint, float32, float64:
				strv = fmt.Sprintf("%v", v)
			default:
				return errors.New("wrong value type for " + key)
			}
			_snippet.attrs["05"+key] = strv
		}
	}
	return nil
}

// Embed adds _cmp to the map of embedded components within the _parent.
// If a component with the same _id has already been embedded it's replaced.
// Usually the _id is the id of the html element.
func (_parent *HTMLSnippet) Embed(_id string, _cmp any) {
	_id = helper.Normalize(_id)
	if _parent.embedded == nil {
		_parent.embedded = make(map[string]any, 1)
	}
	_parent.embedded[_id] = _cmp
	// DEBUG: fmt.Printf("embedding %q(%v) into %s\n", id, reflect.TypeOf(cmp).String(), s.Id())
}

// Embedded returns the map of embedded components, keyed by their id.
func (_parent HTMLSnippet) Embedded() map[string]any {
	return _parent.embedded
}

// Template is an interface implementation of HtmlComposer
func (_snippet HTMLSnippet) Template(*DataState) (_t SnippetTemplate) {
	_t.Tag = _snippet.Tag
	_t.Body = _snippet.Body
	_t.Attributes = _snippet.Attributes()
	return
}

/******************************************************************************
 * PRIVATE
 *****************************************************************************/

func (_snippet *HTMLSnippet) makemap() {
	if _snippet.attrs == nil {
		_snippet.attrs = make(map[string]string)
	}
}

// stringifyValue returns an empty string if v is empty or "false".
// returns an unquoted string if v can be converted in float.
func stringifyValue(_val string) string {
	tv := strings.Trim(_val, " ")
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
	return delim + _val + delim

}
