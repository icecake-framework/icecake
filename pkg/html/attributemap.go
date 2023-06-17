package html

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/icecake-framework/icecake/pkg/namepattern"
	"github.com/icecake-framework/icecake/pkg/registry"
	"github.com/sunraylab/verbose"
)

// AttributeMap represents a list of safe HTML attributes, providing method to easily set, update and extract attributes.
//
// Validity of attributes is checked once when setting it. Most most of settin methods are chainable they does not returns errors.
// If verbose mode is turnned on, setting errors are printed out otherwise the setting fails and nothing happen.
// Use CheckAttribute to check the validity of an attribute and to receive an error.
type AttributeMap map[string]string // map of all attributes of any type

// Reset deletes all attributes in the map
func (amap AttributeMap) Reset() AttributeMap {
	for k := range amap {
		delete(amap, k)
	}
	return amap
}

// Attribute returns the value of the attribute identified by its name.
// Returns false if the attribute does not exist.
// Blanks must be trimed before.
func (amap AttributeMap) Attribute(name string) (value string, found bool) {
	_, key := key(name)
	value, found = amap[key]
	return value, found
}

func key(name string) (index string, key string) {
	seqmap := map[string]string{"id": "1", "name": "2", "class": "3", "tabIndex": "4", "style": "5"}
	index, seqfound := seqmap[name]
	if !seqfound {
		index = "9"
	}
	key = index + name
	return index, key
}

// SetAttribute creates an attribute in the map and set its value.
// update flag indicates if existing attribute must be updated or not.
// SetAttribute returns the map to allow chainning and never returns an error.
// If the name or the value are not valid nothing is created and a log is printed out if verbose mode is on.
// Use CheckAttribute to check name and value validity.
//
// NOTE: attribute's name are case sensitive https://www.w3.org/TR/2010/WD-html-markup-20101019/documents.html#:~:text=Attribute%20names%20for%20HTML%20elements%20must%20exactly%20match%20the%20names,attribute%20names%20are%20case%2Dsensitive.
func (amap AttributeMap) SetAttribute(name string, value string, update bool) AttributeMap {
	name = strings.Trim(name, " ")
	switch name {
	case "id":
		amap.SetId(value)
	case "class":
		amap.SetClass(value)
	case "tabIndex":
		amap.SetTabIndex(value)
	case "style":
		amap.SetStyle()
	}

	err := amap.setAttribute(name, value, update)
	verbose.Error("SetAttribute", err)
	return amap
}

// SetAttributeIf SetAttribute if the condition is true.
// update flag indicates if existing attribute must be updated or not.
// SetAttribute returns the map to allow chainning and never returns an error.
// If the name or the value are not valid nothing is created and a log is printed out if verbose mode is on.
// Use CheckAttribute to check name and value validity.
//
// NOTE: attribute's name are case sensitive https://www.w3.org/TR/2010/WD-html-markup-20101019/documents.html#:~:text=Attribute%20names%20for%20HTML%20elements%20must%20exactly%20match%20the%20names,attribute%20names%20are%20case%2Dsensitive.
func (amap AttributeMap) SetAttributeIf(condition bool, name string, value string, update bool) AttributeMap {
	if condition {
		amap.SetAttribute(name, value, update)
	}
	return amap
}

// setAttribute set the attibute and its value within the map. name must be trimed by the caller.
// The validity of the value must be checked by the caller.
//
// setAttributes does nothing if the name already exists in the map and overwrite parameter is false.
//
// lowercase value equal false is considered as a boolean and the attribute is deleted according to overwrite parameter.
//
// Returns an error if the name does not meet HTML5 valid pattern.
func (amap AttributeMap) setAttribute(name string, value string, overwrite bool) (err error) {

	index, key := key(name)
	nameindexed := index != "9"

	_, exists := amap[key]
	if !overwrite && exists {
		return nil
	}

	falsevalue := strings.Trim(strings.ToLower(value), " ") == "false"
	emptyvalue := strings.Trim(value, " ") == ""
	if overwrite && (falsevalue || (nameindexed && emptyvalue)) {
		delete(amap, key)
		return nil
	}

	if !nameindexed && !namepattern.IsValid(name) {
		return fmt.Errorf("not valid attribute name %q", name)
	}

	// 			if !namepattern.IsValid(value) {
	// 				return fmt.Errorf("not valid id %q", value)
	// 			}

	// 		case "tabIndex":
	// 		idx, errx := strconv.Atoi(value)
	// 		if errx != nil {
	// 			return fmt.Errorf("not valid tabIndex %q", value)
	// 		}
	// 		amap["02tabIndex"] = strconv.Itoa(idx)
	// case "class":
	// 	if overwrite {
	// 		amap.ResetClasses(value)
	// 	} else if value != "" {
	// 		amap.SetClasses(lst)
	// 	}

	amap[key] = value
	return nil
}

// RemoveAttribute removes the attribute identified by its name.
// Does nothing if the name is not in the map.
// Blanks must be trimed before.
func (amap AttributeMap) RemoveAttribute(name string) AttributeMap {
	_, key := key(name)
	delete(amap, key)
	return amap
}

// ToggleAttribute toggles an attribute like a boolean.
// id, tabindex, name, class and style can't be setup with this method.
// if the attribute exists it's removed, if it does not exists it's created without value.
// In verbose mode, ToggleAttribute can log an error if the name is not valid.
// Blanks must be trimed before.
func (amap AttributeMap) ToggleAttribute(name string) AttributeMap {
	_, found := amap.Attribute(name)
	if !found {
		amap.SetAttribute(name, "", true)
	} else {
		amap.RemoveAttribute(name)
	}
	return amap
}

// Strings returns the formated list of attributes, ready to use to generate the container element.
// always sorted the same way : 1>id 2>tabindex 3>class 4>style 5>others sorted alpha by name
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
		sv, err := StringifyAttributeValue(v)
		if err != nil {
			strhtml += "='" + err.Error() + "'"
		} else if len(sv) > 0 {
			strhtml += "=" + sv
		}
		strhtml += " "
	}
	strhtml = strings.TrimRight(strhtml, " ")
	return strhtml
}

// StringifyAttributeValue converts the value in string and add quotes if required.
// Returns an empty string if value is empty or if value == "false".
// Returns an unquoted string if v can be converted in float.
// Returns a quoted value in any other cases. The quote is a double quote " unless the value contains one. In that case a simple quote if choosen '.
// If the value contains both kind of quotes an error is return.
func StringifyAttributeValue(value string) (string, error) {

	tv := strings.Trim(value, " ")
	lv := strings.ToLower(tv)
	if len(tv) == 0 || lv == "false" {
		return "", nil
	}

	if lv == "true" {
		return lv, nil
	}

	_, err := strconv.ParseFloat(tv, 64)
	if err == nil {
		return tv, nil
	}

	var delim string
	if strings.ContainsRune(tv, rune('"')) {
		delim = "'"
		if strings.ContainsRune(tv, rune('\'')) {
			return "", errors.New("ambiguous quotes in the value")
		}
	} else {
		delim = "\""
	}
	return delim + value + delim, nil
}

// return the id attribute if any
func (amap AttributeMap) Id() string {
	id := amap["01id"]
	return id
}

// SetId sets or overwrites the id attribute.
// if id is empty, the id attribute is removed.
// NOTA: in HTML5 id are case sensitive.
//
// SetId returns the map to allow chainning. SetId never returns an error.
// If the id is not valid nothing is setted and a log is printed out if verbose mode is on.
// Use CheckAttribute to check name and value validity.
func (amap AttributeMap) SetId(id string) AttributeMap {
	if !namepattern.IsValid(id) {
		verbose.Error("SetId", errors.New("id does not match a valid pattern"))
	} else {
		amap.setAttribute("id", id, true)
	}
	return amap
}

// SetUniqueId sets or overwrites the id attribute by generating a unique id starting with the prefix.
// "ick-" is used to prefix the returned id if prefix is empty.
func (amap AttributeMap) SetUniqueId(prefix string) {
	amap["01id"] = registry.GetUniqueId(prefix)
}

// Classes returns the class attribute as a full string
func (amap AttributeMap) Classes() string {
	return amap["03class"]
}

// HasClass returns if the class exists in the class attribute of the map.
// Blanks must be trimed before.
func (amap AttributeMap) HasClass(class string) bool {
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

// SetClasses adds the list of classes to the class attribute
// Duplicates are not inserted twice.
// TODO: check validity of the class name pattern
func (amap AttributeMap) AddClasses(list string) AttributeMap {
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
func (amap AttributeMap) AddClassesIf(condition bool, list string) AttributeMap {
	if condition {
		amap.AddClasses(list)
	}
	return amap
}

// ResetClasses replaces any existing classes with _clist to the class attribute
// _clist must contains strings separated by spaces.
// All classes are removed if _clist is empty.
// TODO: check validity of the class name pattern
func (amap AttributeMap) SetClasses(list string) AttributeMap {
	delete(amap, "3class")
	amap.AddClasses(list)
	return amap
}

// RemoveClasses removes any class in _list from the "class" attribute.
// Does nothing if c did not exist.
func (amap AttributeMap) RemoveClasses(list string) AttributeMap {
	actual := amap["03class"]
	new := ""
	actualf := strings.Fields(actual)
	listf := strings.Fields(list)
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

// SwitchClasses removes classes and set the new ones
// Does nothing if c did not exist.
func (amap AttributeMap) SwitchClasses(removeclass string, newclasses string) AttributeMap {
	amap.RemoveClasses(removeclass)
	amap.AddClasses(newclass)
	return amap
}

// TabIndex returns the TabIndex attribute
func (amap AttributeMap) TabIndex() int {
	sidx, found := amap.Attribute("tabIndex")
	idx, _ := strconv.Atoi(sidx)
	return idx
}

func (amap AttributeMap) SetTabIndex(idx int) AttributeMap {
	err := amap.setAttribute("tabIndex", strconv.Itoa(idx), true)
	verbose.Error("SetTabIndex", err)
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

// TryParseAttributes tries to ParseAttributes and ignore errors.
// if alist is empty, an empty but not nil AttributeMap is returned.
func TryParseAttributes(alist string) (amap AttributeMap) {
	amap, _ = ParseAttributes(alist)
	return amap
}

// ParseAttribute splits alist into attributes separated by spaces and sets each to the AttributeMap.
//
// An attribute can have a value at the right of an "=" symbol.
// The value can be delimited by quotes ( " or ' ) and in that case may contains whitespaces.
// The string is processed until the end or an error occurs when invalid char is met.
// Always returns a not nil amap.
// Returns error when an attribute name does not match the valid HTML pattern (https://stackoverflow.com/questions/925994/what-characters-are-allowed-in-an-html-attribute-name).
// Returns error when
// TODO: ParseAttributes must returns an error for not valid attribute value
func ParseAttributes(alist string) (amap AttributeMap, err error) {
	amap = make(AttributeMap)
	var strnames string
	unparsed := alist
	for i := 0; len(unparsed) > 0; i++ {

		// process all simple attributes until next "="
		var hasval bool
		strnames, unparsed, hasval = strings.Cut(unparsed, "=")
		names := strings.Fields(strnames)
		for i, n := range names {
			// validity checked on the fly
			// HACK: better to test validity of name and value within SetAttribute
			if !namepattern.IsValid(n) {
				return nil, fmt.Errorf("attribute name %q is not valid", n)
			}
			if i < len(names)-1 || !hasval {
				amap.SetAttribute(n, "", true)
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
		amap.SetAttribute(name, value, true)
	}
	return amap, nil
}

func CheckAttributeName(name string) error {

}

func CheckAttributeValue(name string, value any) error {

}
