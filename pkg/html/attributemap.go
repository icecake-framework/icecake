package html

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/icecake-framework/icecake/pkg/registry"
	"github.com/icecake-framework/icecake/pkg/stringpattern"
	"github.com/sunraylab/verbose"
	"golang.org/x/exp/slices"
)

// AttributeMap represents a list of safe HTML attributes, providing method to easily set, update and extract attributes.
//
// Validity of attributes is checked once when setting it. Most most of settin methods are chainable they does not returns errors.
// If verbose mode is turnned on, setting errors are printed out otherwise the setting fails and nothing happen.
// Use CheckAttribute to check the validity of an attribute and to receive an error.
//
// Blanks must be trimed in calling parameters before calling methods.
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
// blanks at the ends of the name should be trimmed before call.
func (amap AttributeMap) Attribute(name string) (value string, found bool) {
	value, found = amap[name]
	return value, found
}

// SetAttribute creates an attribute in the map and set its value.
// Blanks at the ends of the name are automatically trimmed.
// update flag indicates if existing attribute must be updated or not.
// SetAttribute returns the map to allow chainning and so never returns an error.
// If the name or the value are not valid nothing is created and a log is printed out if verbose mode is on.
// Use CheckAttribute to check name and value validity.
//
// NOTE: attribute's name are case sensitive https://www.w3.org/TR/2010/WD-html-markup-20101019/documents.html#:~:text=Attribute%20names%20for%20HTML%20elements%20must%20exactly%20match%20the%20names,attribute%20names%20are%20case%2Dsensitive.
func (amap AttributeMap) SetAttribute(name string, value string, update bool) AttributeMap {
	err := amap.setAttribute(name, value, update)
	verbose.Error("SetAttribute", err)
	return amap
}

// setAttribute creates an attribute in the map and set its value.
// Blanks at the ends of the name are automatically trimmed.
// update flag indicates if existing attribute must be updated or not.
// setAttribute returns error related about name and value validity.
func (amap AttributeMap) setAttribute(name string, value string, update bool) (err error) {
	name = strings.Trim(name, " ")

	err = CheckAttribute(name, value)
	if err != nil {
		return err
	}

	switch name {
	case "id":
		id := strings.Trim(value, " ")
		amap.saveAttribute(name, id, update)
	case "class":
		if update {
			amap.ResetClasses()
		}
		c := strings.Fields(value)
		amap.AddClasses(c...)
	case "tabIndex":
		i, _ := strconv.Atoi(value)
		amap.saveAttribute(name, strconv.Itoa(i), update)
	case "style":
		amap.saveAttribute("style", value, update)
	default:
		amap.saveAttribute(name, value, update)
	}
	return nil
}

func checkid(id string) error {
	if id == "" || !stringpattern.IsValidName(id) {
		return fmt.Errorf("id %q is not valid", id)
	}
	return nil
}

// SetAttributeIf SetAttribute if the condition is true.
// Blanks at the ends of the name are automatically trimmed.
// update flag indicates if existing attribute must be updated or not.
// SetAttribute returns the map to allow chainning and so never returns an error.
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

// saveAttribute set the attibute and its value within the map. blanks ate the ends of the name must be trimed by the caller.
// The validity of the value must be checked by the caller.
//
// setAttributes does nothing if the name already exists in the map and overwrite parameter is false.
// lowercase value equal "false" is considered as a boolean and the attribute is deleted according to overwrite parameter.
// setting a blank value to attributes id, class, style, name or tabindex will remove the attribute according to overwrite parameter.
//
// Returns if the attributes has been overwritten (updated)
func (amap AttributeMap) saveAttribute(name string, value string, overwrite bool) bool {

	_, exists := amap[name]
	if !overwrite && exists {
		return false
	}

	var emptyvalue, nameinlist bool
	if emptyvalue = strings.Trim(value, " ") == ""; emptyvalue {
		nameinlist = slices.Contains([]string{"id", "class", "tabIndex", "name", "style"}, name)
	}

	falsevalue := strings.Trim(strings.ToLower(value), " ") == "false"
	if overwrite && (falsevalue || (emptyvalue && nameinlist)) {
		delete(amap, name)
	} else {
		amap[name] = value
	}
	return exists
}

// CheckAttribute check validity of the attribute name and its value.
// blanks at the ends of the name should be trimmed before call.
func CheckAttribute(name string, value string) (err error) {
	name = strings.Trim(name, " ")
	switch name {
	case "id":
		id := strings.Trim(value, " ")
		err = checkid(id)
	case "class":
		for _, c := range strings.Fields(value) {
			err = checkclass(c)
			if err != nil {
				break
			}
		}
	case "tabIndex":
		_, err = strconv.Atoi(value)
	case "style":
		err = checkstyle(value)
	default:
	}
	if err == nil {
		if !stringpattern.IsValidName(name) {
			err = fmt.Errorf("attribute %q is not a valid name", name)
		}
	}
	return err
}

// RemoveAttribute removes the attribute identified by its name.
// Does nothing if the name is not in the map.
// blanks at the ends of the name should be trimmed before call.
func (amap AttributeMap) RemoveAttribute(name string) AttributeMap {
	delete(amap, name)
	return amap
}

// ToggleAttribute toggles an attribute like a boolean.
// if the attribute exists it's removed, if it does not exists it's created without value.
// id, tabindex, name, class and style can't be setup with this method.
// In verbose mode ToggleAttribute can log an error if the name is not valid.
// blanks at the ends of the name should be trimmed before call.
func (amap AttributeMap) ToggleAttribute(name string) AttributeMap {
	_, found := amap[name]
	if !found {
		amap.SetAttribute(name, "", true)
	} else {
		amap.RemoveAttribute(name)
	}
	return amap
}

// Id returns the id attribute if any
func (amap AttributeMap) Id() string {
	return amap["id"]
}

// SetId sets or overwrites the id attribute.
// blanks at the ends of the id are automatically trimmed.
// if id is empty, the id attribute is removed.
// NOTA: in HTML5 id are case sensitive.
//
// SetId returns the map to allow chainning. SetId never returns an error.
// If the id is not valid nothing is setted and a log is printed out if verbose mode is on.
// Use CheckAttribute to check name and value validity.
func (amap AttributeMap) SetId(id string) AttributeMap {
	amap.SetAttribute("id", id, true)
	return amap
}

// SetUniqueId sets or overwrites the id attribute by generating a unique id starting with the prefix.
// "ick-" is used to prefix the returned id if prefix is empty.
func (amap AttributeMap) SetUniqueId(prefix string) {
	amap.saveAttribute("id", registry.GetUniqueId(prefix), true)
}

// Classes returns the class attribute as a full string.
func (amap AttributeMap) Classes() string {
	return amap["class"]
}

// HasClass returns if the class exists in the list of classes.
// return false if class is empty
// blanks at the ends of the name should be trimmed before call.
func (amap AttributeMap) HasClass(class string) bool {
	if class == "" {
		return false
	}
	actual := amap["class"]
	actualf := strings.Fields(string(actual))
	for _, actualc := range actualf {
		if actualc == class {
			return true
		}
	}
	return false
}

// AddClasses adds the list of classes to the class attribute
// Duplicates are not inserted twice.
func (amap AttributeMap) AddClasses(cs ...string) AttributeMap {
	actual := amap["class"]
	new := actual
	actualf := strings.Fields(actual)
nextf:
	for _, addc := range cs {
		if addc != "" {
			for _, actualc := range actualf {
				if actualc == addc {
					continue nextf
				}
			}
			if err := checkclass(addc); err != nil {
				verbose.Error("AddClasses", err)
			}
			new += " " + addc
		}
	}
	new = strings.TrimLeft(new, " ")
	if new != "" {
		amap["class"] = new
	}
	return amap
}

// SetClassesIf SetClasses if the _condition is true
func (amap AttributeMap) AddClassesIf(condition bool, c ...string) AttributeMap {
	if condition {
		amap.AddClasses(c...)
	}
	return amap
}

func checkclass(class string) error {
	if class == "" || !stringpattern.IsValidName(class) {
		return fmt.Errorf("class %q is not valid", class)
	}
	return nil
}

// ResetClasses removes all classes by removing the class attribute?
func (amap AttributeMap) ResetClasses() AttributeMap {
	delete(amap, "class")
	return amap
}

// RemoveClasses removes any class in _list from the "class" attribute.
// Does nothing if c did not exist.
func (amap AttributeMap) RemoveClasses(rc ...string) AttributeMap {
	actual := amap["class"]
	new := ""
	actualf := strings.Fields(actual)
nexta:
	for _, actualc := range actualf {
		for _, remc := range rc {
			if actualc == remc {
				continue nexta
			}
		}
		new += " " + actualc
	}
	new = strings.TrimRight(new, " ")
	amap["class"] = new
	return amap
}

// SwitchClass removes one class and set a new one
func (amap AttributeMap) SwitchClass(removec string, addc string) AttributeMap {
	amap.RemoveClasses(removec)
	amap.AddClasses(addc)
	return amap
}

// TabIndex returns the TabIndex attribute
func (amap AttributeMap) TabIndex() (i int) {
	sidx := amap["tabIndex"]
	i, _ = strconv.Atoi(sidx)
	return i
}

// SetTabIndex sets or overwrites the tabIndex attribute.
// SetTabIndex returns the map to allow chainning.
func (amap AttributeMap) SetTabIndex(idx int) AttributeMap {
	amap.saveAttribute("tabIndex", strconv.Itoa(idx), true)
	return amap
}

// Style returns the style attribute
func (amap AttributeMap) Style() string {
	return amap["style"]
}

// SetStyle sets or overwrites the style attribute.
// SetStyle returns the map to allow chainning.
func (amap AttributeMap) SetStyle(style string) AttributeMap {
	style = strings.Trim(style, " ")
	if err := checkstyle(style); err != nil {
		verbose.Error("SetStyle", err)
	} else {
		amap.saveAttribute("style", style, true)
	}
	return amap
}

func checkstyle(style string) error {
	//TODO check style string
	return nil
}

// Is returns the true is the attribute exists and if it's value is not false nor 0.
// blanks at the ends of the name should be trimmed before call.
func (amap AttributeMap) Is(name string) bool {
	value, found := amap[name]
	if !found {
		return false
	}

	v := strings.Trim(strings.ToLower(value), " ")
	if v == "false" {
		return false
	}
	if v != "" {
		n, err := strconv.Atoi(v)
		if n == 0 && err == nil {
			return false
		}
	}
	return true
}

// SetBool set or overwrite a boolean attribute.
// SetBool returns the map to allow chainning.
// SetBool does nothing if trying to set id, style, name or class attribute
func (amap AttributeMap) SetBool(name string, f bool) AttributeMap {
	if !f {
		amap.RemoveAttribute(name)
	} else {
		amap.SetAttribute(name, "", true)
	}
	return amap
}

// IsDisabled returns true if the disabled attribute is set
func (amap AttributeMap) IsDisabled() bool {
	return amap.Is("disabled")
}

// SetDisabled set the boolean disabled attribute
func (amap AttributeMap) SetDisabled(f bool) AttributeMap {
	amap.SetBool("disabled", true)
	return amap
}

// Strings returns the formated list of attributes, ready to use to generate the container element.
// always sorted the same way : 1>id 2>name 3>class 4>others sorted alpha by name
func (amap AttributeMap) String() string {
	if len(amap) == 0 {
		return ""
	}

	strhtml := ""
	sorted := make([]string, 0, len(amap))
	for k := range amap {
		sorted = append(sorted, k)
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i] == "id" {
			return true
		}
		if sorted[i] == "name" {
			return true
		}
		if sorted[i] == "class" {
			return true
		}
		return sorted[i] < sorted[j]
	})

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
			if i < len(names)-1 || !hasval {
				err = amap.setAttribute(n, "", true)
				if err != nil {
					return make(AttributeMap), err
				}
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
		err = amap.setAttribute(name, value, true)
		if err != nil {
			return make(AttributeMap), err
		}
	}
	return amap, nil
}
