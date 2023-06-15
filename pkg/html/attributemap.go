package html

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/icecake-framework/icecake/pkg/namepattern"
	"github.com/icecake-framework/icecake/pkg/registry"
)

type AttributeMap map[string]string // map of all attributes of any type

func NewAttributeMap() AttributeMap {
	return make(AttributeMap)
}

func (amap AttributeMap) Id() string {
	id := amap["01id"]
	return id
}

// SetId sets or overwrites the id attribute of the html snippet
// if id is empty, the id attribute is removed
func (amap AttributeMap) SetId(id string) AttributeMap {
	if _, found := amap["01id"]; found {
		delete(amap, "01id")
	} else {
		amap["01id"] = id
	}
	return amap
}

// SetUniqueId sets or overwrites the id attribute of the html snippet
// generating a unique id starting with the prefix.
// "ick-" is used to prefix the returned id if prefix is empty.
func (amap AttributeMap) SetUniqueId(prefix string) {
	amap["01id"] = registry.GetUniqueId(prefix)
}

func (amap AttributeMap) Reset() AttributeMap {
	for k, _ := range amap {
		delete(amap, k)
	}
	return amap
}

/******************************************************************************
* Classes
******************************************************************************/

// Classes returns the class attribute
func (amap AttributeMap) Classes() string {
	str := amap["03class"]
	return string(str)
}

// Classes returns the class attribute
func (amap AttributeMap) HasClass(class string) bool {
	class = strings.Trim(class, " ")
	if class == "" {
		return false
	}
	actual := amap["03class"]
	actualf := strings.Fields(string(actual))
	for _, actualc := range actualf {
		if actualc == class {
			return true
		}
	}
	return false
}

// ResetClasses replaces any existing classes with _clist to the class attribute
// _clist must contains strings separated by spaces.
// All classes are removed if _clist is empty.
// TODO: check validity of the class name pattern
func (amap AttributeMap) ResetClasses(list string) AttributeMap {
	n := ""
	f := strings.Fields(list)
	for _, c := range f {
		if c != "" {
			n += c + " "
		}
	}
	n = strings.TrimRight(n, " ")
	if n == "" {
		delete(amap, "03class")
	} else {
		amap["03class"] = n
	}
	return amap
}

// SetClasses adds the _list of classes to the class attribute
// duplicate are not inserted twice.
// TODO: check validity of the class name pattern
func (amap AttributeMap) SetClasses(list string) AttributeMap {
	actual := amap["03class"]
	new := string(actual)
	actualf := strings.Fields(string(actual))
	listf := strings.Fields(list)
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
		amap["03class"] = new
	}
	return amap
}

// SetClassesIf SetClasses if the _condition is true
func (amap AttributeMap) SetClassesIf(condition bool, list string) AttributeMap {
	if condition {
		amap.SetClasses(list)
	}
	return amap
}

// RemoveClasses removes any class in _list from the "class" attribute.
// Does nothing if c did not exist.
func (amap AttributeMap) RemoveClasses(list string) AttributeMap {
	actual := amap["03class"]
	new := ""
	actualf := strings.Fields(string(actual))
	listf := strings.Fields(string(list))
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
	amap["03class"] = new
	return amap
}

// SwitchClasses removes _remove classes and set the _new one
// Does nothing if c did not exist.
func (amap AttributeMap) SwitchClasses(removeclass string, newclass string) {
	amap.RemoveClasses(removeclass)
	amap.SetClasses(newclass)
}

/******************************************************************************
* Attributes
******************************************************************************/

// Attributes returns the formated list of attributes used to generate the container element.
// always sorted the same way : 1.id 2.tabindex 3.class 4.style 5. other-alpha
func (amap AttributeMap) String() string {
	if len(amap) == 0 {
		return ""
	}

	strhtml := ""
	sorted := make([]string, 0, len(amap))
	for k := range amap {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)
	for _, kx := range sorted {
		v := amap[kx]
		k := kx[2:]
		strhtml += k
		sv := stringifyValue(string(v))
		if len(sv) > 0 {
			strhtml += "=" + sv
		}
		strhtml += " "
	}
	strhtml = strings.TrimRight(strhtml, " ")
	return strhtml
}

// stringifyValue returns an empty string if v is empty or if v == "false".
// stringifyValue returns an unquoted string if v can be converted in float.
// stringifyValue returns a quoted value, choosing the right quote, in other cases.
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

// Attribute returns the value of the attribute identified by _key.
// Returns false if the attribute does not exist.
func (amap AttributeMap) Attribute(_key string) (string, bool) {
	_key = strings.Trim(_key, " ")
	v, found := amap["05"+_key]
	return string(v), found
}

// CreateAttribute create an attribute and set its value.
// If the attribute already exists nothing is done. Use SetAttribute if you want to create OR update an attribute.
// func (amap AttributeMap) CreateAttribute(_key string, _value any) HTMLComposer {
// 	err := _snippet.setAttribute(_key, _value, false)
// 	if err != nil {
// 		fmt.Println("CreateAttribute fails:", err)
// 	}
// 	return _snippet
// }

// SetAttribute create an attribute and set its value.
// update indicates if an existing attribute must be updated or not.
func (amap AttributeMap) SetAttribute(_key string, _value any, update bool) AttributeMap {
	err := amap.setAttribute(_key, _value, update)
	if err != nil {
		fmt.Println("SetAttribute fails:", err)
	}
	return amap
}

// SetAttributeIf SetAttribute if the condition is true.
// update indicates if an existing attribute must be updated or not.
func (amap AttributeMap) SetAttributeIf(condition bool, key string, value any, update bool) AttributeMap {
	if condition {
		amap.SetAttribute(key, value, update)
	}
	return amap
}

func (amap AttributeMap) setAttribute(key string, value any, overwrite bool) error {
	key = strings.Trim(key, " ")
	switch strings.ToLower(key) {
	case "id":
		_, found := amap["01id"]
		if !found || overwrite {
			switch v := value.(type) {
			case string:
				if v == "" {
					delete(amap, "01id")
				} else {
					amap.SetId(v)
				}
			case HTMLString:
				if v == "" {
					delete(amap, "01id")
				} else {
					amap.SetId(string(v))
				}
			default:
				return errors.New("wrong value type for id")
			}
		}
	case "tabindex":
		_, found := amap["02tabIndex"]
		if !found || overwrite {
			switch v := value.(type) {
			case string:
				idx, _ := strconv.Atoi(string(v))
				amap.SetTabIndex(idx)
			case HTMLString:
				idx, _ := strconv.Atoi(string(v))
				amap.SetTabIndex(idx)
			case int:
				amap.SetTabIndex(v)
			case uint:
				amap.SetTabIndex(int(v))
			case float32:
				amap.SetTabIndex(int(v))
			case float64:
				amap.SetTabIndex(int(v))
			default:
				return errors.New("wrong value type for tabindex")
			}
		}
	case "class":
		var lst string
		switch v := value.(type) {
		case string:
			lst = v
		case HTMLString:
			lst = string(v)
		default:
			return errors.New("wrong value type for class")
		}
		if overwrite {
			amap.ResetClasses(lst)
		} else if value != "" {
			amap.SetClasses(lst)
		}
	case "style":
		// TODO: handle style update to not overwrite
		_, found := amap["04style"]
		if !found || overwrite {
			var style HTMLString
			switch v := value.(type) {
			case string:
				style = HTMLString(v)
			case HTMLString:
				style = v
			default:
				return errors.New("wrong value type for class")
			}
			amap.SetStyle(style)
		}
	default:
		_, found := amap["05"+key]
		if !found || overwrite {
			var strv string
			switch v := value.(type) {
			case string:
				strv = v
			case HTMLString:
				strv = string(v)
			case bool:
				if v {
					strv = ""
				} else {
					delete(amap, "05"+key)
					break
				}
			case int, uint, float32, float64:
				strv = fmt.Sprintf("%v", v)
			default:
				return errors.New("wrong value type for " + key)
			}
			amap["05"+key] = strv
		}
	}
	return nil
}

// RemoveAttribute remove the the attribute identified by _key.
// Does nothing if the _key is not found.
func (amap AttributeMap) RemoveAttribute(_key string) AttributeMap {
	_key = strings.Trim(_key, " ")
	switch strings.ToLower(_key) {
	case "id":
		delete(amap, "01id")
	case "tabindex":
		delete(amap, "02tabIndex")
	case "class":
		delete(amap, "03class")
	case "style":
		delete(amap, "04style")
	default:
		delete(amap, "05"+_key)
	}
	return amap
}

// ToggleAttribute toggles the boolean attribute _key. Sets it when unsetted, and unset whrn setted.
func (amap AttributeMap) ToggleAttribute(_key string) AttributeMap {
	_key = strings.Trim(_key, " ")
	_, found := amap["05"+_key]
	if !found {
		amap["05"+_key] = ""
	} else {
		delete(amap, "05"+_key)
	}
	return amap
}

/******************************************************************************
* Special Attributes
******************************************************************************/

// TabIndex returns the TabIndex attribute
func (amap AttributeMap) TabIndex() int {
	sidx := amap["02tabIndex"]
	idx, _ := strconv.Atoi(string(sidx))
	return idx
}

func (amap AttributeMap) SetTabIndex(idx int) AttributeMap {
	amap["02tabIndex"] = strconv.Itoa(idx)
	return amap
}

// IsDisabled returns the boolean attribute
func (amap AttributeMap) IsDisabled() bool {
	str, found := amap["05disabled"]
	if !found || strings.ToLower(string(str)) == "false" || str == "0" {
		return false
	}
	return true
}

func (amap AttributeMap) SetDisabled(_f bool) AttributeMap {
	if _f {
		amap["05disabled"] = ""
	} else {
		delete(amap, "05disabled")
	}
	return amap
}

/******************************************************************************
* Style
******************************************************************************/

// Style returns the style attribute
func (amap AttributeMap) Style() string {
	str := amap["04style"]
	return string(str)
}

// TODO: check style validity
func (amap AttributeMap) SetStyle(style HTMLString) AttributeMap {
	amap["04style"] = string(style)
	return amap
}

// TryParseAttributes tries to ParseAttributes and ignore errors
func TryParseAttributes(alist string) (amap AttributeMap) {
	amap, _ = ParseAttributes(alist)
	return
}

// ParseAttribute split alist into attributes separated by spaces and set each to the HtmlComposer.
// An attribute can have a value at the right of an "=" symbol.
// The value can be delimited by quotes ( " or ' ) and in that case may contains whitespaces.
// The string is processed until the end or an error occurs when invalid char is met.
// Existing _cmp attributes are not overwritten.
// TODO: secure _alist ?
func ParseAttributes(alist string) (amap AttributeMap, err error) {

	amap = make(map[string]string)
	var strnames string
	unparsed := alist
	for i := 0; len(unparsed) > 0; i++ {

		// process all simple attributes until next "="
		var hasval bool
		strnames, unparsed, hasval = strings.Cut(unparsed, "=")
		names := strings.Fields(strnames)
		for i, n := range names {
			if !namepattern.IsValid(n) {
				return amap, fmt.Errorf("attribute name %q is not valid", n)
			}
			if i < len(names)-1 || !hasval {
				amap.SetAttribute(n, "", false)
			}
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
		amap.SetAttribute(name, value, false)
	}
	return amap, nil
}
