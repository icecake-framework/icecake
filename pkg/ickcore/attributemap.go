package ickcore

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/icecake-framework/icecake/internal/helper"
	"github.com/icecake-framework/icecake/pkg/namingpattern"
	"github.com/lolorenzo777/verbose"
	"golang.org/x/exp/slices"
)

// AttributeMap represents a list of valid HTML tag attributes, providing methods to easily set, update and extract each.
//
// Validity of attributes is checked once when setting it. Most most of setters are chainable, so they does not returns errors.
// If verbose mode is turnned on, errors are printed out otherwise the setter fails and nothing happen.
// Use CheckAttribute to check the validity of an attribute and to receive an error.
//
// HTML Attributes name are case insensitive: https://www.w3.org/TR/2010/WD-html-markup-20101019/documents.html
type AttributeMap map[string]string // map of HTML tag attributes

// Clone clones the attribute map
func (amap AttributeMap) Clone() AttributeMap {
	c := make(AttributeMap, len(amap))
	for k, v := range amap {
		c[k] = v
	}
	return amap
}

// AttributeString returns the formated list of attributes, ready to use to generate the tag element.
// always sorted the same way : 1>id 2>name 3>class 4>others sorted alpha by name
func (amap AttributeMap) AttributeString() string {
	if len(amap) == 0 {
		return ""
	}

	attrsortindex := map[string]int{"id": 0, "name": 1, "class": 2}

	strhtml := ""
	sorted := make([]string, 0, len(amap))
	for k := range amap {
		sorted = append(sorted, k)
	}
	sort.SliceStable(sorted, func(i, j int) bool {
		ii, foundi := attrsortindex[sorted[i]]
		if !foundi {
			ii = 9
		}
		jj, foundj := attrsortindex[sorted[j]]
		if !foundj {
			jj = 9
		}
		if ii == jj {
			return sorted[i] < sorted[j]
		}
		return ii < jj
	})

	for _, k := range sorted {
		v := amap[k]
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

// Attribute returns the value of the attribute identified by its name.
// Returns false if the attribute does not exist.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive
func (amap AttributeMap) Attribute(aname string) (value string, found bool) {
	aname = helper.Normalize(aname)
	value, found = amap[aname]
	return value, found
}

// TrySetAttribute creates an attribute in the map and set its value.
// If the attribute already exists in the map then does nothing.
// Returns if the attribute has been created/updated or not.
func (amap AttributeMap) TrySetAttribute(name string, value string) bool {
	name = helper.Normalize(name)
	update, err := amap.setAttribute(name, value, false)
	verbose.Error("SetAttribute", err)
	return update
}

// SetAttribute creates an attribute in the map and set its value.
// If the attribute already exists in the map it is updated.
// SetAttribute returns the map to allow chainning and so never returns an error.
// If the name or the value are not valid nothing is created and a log is printed out if verbose mode is on.
// Use CheckAttribute to check name and value validity.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive
func (amap AttributeMap) SetAttribute(name string, value string) AttributeMap {
	name = helper.Normalize(name)
	_, err := amap.setAttribute(name, value, true)
	verbose.Error("SetAttribute", err)
	return amap
}

// setAttribute creates an attribute in the map and set its value.
// update flag indicates if existing attribute must be updated or not.
// setAttribute returns error related about name and value validity.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive
func (amap AttributeMap) setAttribute(name string, value string, update bool) (updated bool, err error) {
	err = checkAttribute(name, value)
	if err != nil {
		return false, err
	}

	switch name {
	case "id", "name":
		updated = amap.saveAttribute(name, value, update)
	case "class":
		before := amap["class"]
		if update {
			amap.ResetClass()
		}
		c := strings.Fields(value)
		amap.AddClass(c...)
		updated = amap["class"] != before
	case "tabindex":
		i, _ := strconv.Atoi(value)
		updated = amap.saveAttribute(name, strconv.Itoa(i), update)
	case "style":
		updated = amap.saveAttribute("style", value, update)
	default:
		updated = amap.saveAttribute(name, value, update)
	}
	return updated, nil
}

// SetAttributeIf SetAttribute if the condition is true, otherwise remove the attribute.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive.
func (amap AttributeMap) SetAttributeIf(condition bool, name string, valueiftrue string, valueiffalse ...string) AttributeMap {
	if condition {
		amap.SetAttribute(name, valueiftrue)
	} else if len(valueiffalse) > 0 {
		amap.SetAttribute(name, valueiffalse[0])
	} else {
		amap.RemoveAttribute(name)
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
// Returns if the attributes has been created/updated
func (amap AttributeMap) saveAttribute(name string, value string, overwrite bool) bool {
	_, exists := amap[name]
	if !overwrite && exists {
		return false
	}

	var emptyvalue, nameinlist bool
	if emptyvalue = strings.Trim(value, " ") == ""; emptyvalue {
		nameinlist = slices.Contains([]string{"id", "class", "tabindex", "name", "style"}, name)
	}

	falsevalue := strings.Trim(strings.ToLower(value), " ") == "false"
	if overwrite && (falsevalue || (emptyvalue && nameinlist)) {
		delete(amap, name)
	} else {
		amap[name] = value
	}
	return true
}

// ResetAttributes deletes all attributes in the map.
func (amap AttributeMap) ResetAttributes() AttributeMap {
	for k := range amap {
		delete(amap, k)
	}
	return amap
}

// RemoveAttribute removes the attribute identified by its name.
// Does nothing if the name is not in the map.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive
func (amap AttributeMap) RemoveAttribute(name string) AttributeMap {
	name = helper.Normalize(name)
	delete(amap, name)
	return amap
}

// ToggleAttribute toggles an attribute like a boolean.
// If the attribute exists it's removed, if it does not exists it's created without value.
// id, tabindex, name, class and style can't be setup with this method.
// In verbose mode ToggleAttribute can log an error if the name is not valid.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name are case sensitive
func (amap AttributeMap) ToggleAttribute(name string) AttributeMap {
	name = helper.Normalize(name)
	_, found := amap[name]
	if !found {
		_, err := amap.setAttribute(name, "", true)
		verbose.Error("ToggleAttribute", err)
	} else {
		delete(amap, name)
	}
	return amap
}

// Id returns the id attribute if any
func (amap AttributeMap) Id() string {
	return amap["id"]
}

// SubId build an id with the id attribute followed by a dot and by the provided is string.
// Returns an empty string if the id attribute or the provided id is empty.
func (amap AttributeMap) SubId(id string) string {
	id = strings.Trim(id, " ")
	aid := amap["id"]
	if aid == "" || id == "" {
		return ""
	}
	return aid + "." + id
}

// SetId sets or overwrites the id attribute. In HTML5 id is case sensitive.
// blanks at the ends of the id are automatically trimmed.
// if id is empty, the id attribute is removed.
//
// SetId returns the map to allow chainning. SetId never returns an error.
// If the id is not valid nothing is setted and a log is printed out if verbose mode is on.
func (amap AttributeMap) SetId(id string) AttributeMap {
	_, err := amap.setAttribute("id", id, true)
	verbose.Error("SetId", err)
	return amap
}

// SetUniqueId sets or overwrites the id attribute by generating a unique id starting with the prefix.
// "ick-" is used to prefix the returned id if prefix is empty.
func (amap AttributeMap) SetUniqueId(prefix string) {
	_, uid := GetUniqueId(prefix)
	amap.saveAttribute("id", uid, true)
}

// SetURL sets a url string as a value of a given attributeName.
// if url is nil or if its value is an empty string, the attribute name is setup with an "#" value.
func (amap AttributeMap) SetURL(attributeName string, url *url.URL) AttributeMap {
	strurl := ""
	if url != nil {
		strurl = url.String()
	}
	if strurl == "" {
		strurl = "#"
	}
	_, err := amap.setAttribute(attributeName, strurl, true)
	verbose.Error("SetURL", err)
	return amap
}

// Name returns the name attribute if any
func (amap AttributeMap) Name() string {
	return amap["name"]
}

// SetName sets or overwrites the name attribute. In HTML5 name is case sensitive.
// blanks at the ends of the id are automatically trimmed.
// if name is empty, the name attribute is removed.
//
// SetName returns the map to allow chainning.
// If the name is not valid nothing is setted and a log is printed out if verbose mode is on.
func (amap AttributeMap) SetName(name string) AttributeMap {
	_, err := amap.setAttribute("name", name, true)
	verbose.Error("SetName", err)
	return amap
}

// Classes returns the class attribute as a full string.
func (amap AttributeMap) Classes() string {
	return amap["class"]
}

// HasClass returns if the class exists in the list of amap classes.
// Returns false if class is empty.
func (amap AttributeMap) HasClass(class string) bool {
	if class == "" {
		return false
	}
	actual := amap["class"]
	actualf := strings.Fields(actual)
	for _, actualc := range actualf {
		if actualc == class {
			return true
		}
	}
	return false
}

// AddClass adds each class in lists strings to the class attribute.
// Each lists class string can be either a single class or a string of multiple classes separated by spaces.
// Duplicates are not inserted twice.
func (amap AttributeMap) AddClass(lists ...string) AttributeMap {
	existstr := amap["class"]
	new := existstr
	existcs := strings.Fields(existstr)
nextlist:
	for _, list := range lists {
		cs := strings.Fields(list)
		for _, c := range cs {
			for _, existc := range existcs {
				if existc == c {
					continue nextlist
				}
			}
			if err := checkclass(c); err != nil {
				verbose.Error("AddClasses", err)
				// continue adding other classes even if error
			} else {
				new += " " + c
			}
		}
	}
	new = strings.TrimLeft(new, " ")
	amap["class"] = new
	return amap
}

// AddClassIf adds each c classe to the class attribute if the condition is true.
// Does nothing if the condition is false. See SetClassIf to also remove the class if the condition is false.
// Duplicates are not inserted twice.
func (amap AttributeMap) AddClassIf(condition bool, addlist ...string) AttributeMap {
	if condition {
		amap.AddClass(addlist...)
	}
	return amap
}

// SetClassIf adds each c classe to the class attribute if the condition is true, otherwise remove them.
// Duplicates are not inserted twice.
func (amap AttributeMap) SetClassIf(condition bool, classiftrue string, classiffalse ...string) AttributeMap {
	if condition {
		amap.AddClass(classiftrue)
		if len(classiffalse) > 0 {
			amap.RemoveClass(classiffalse[0])
		}
	} else {
		amap.RemoveClass(classiftrue)
		if len(classiffalse) > 0 {
			amap.AddClass(classiffalse[0])
		}
	}
	return amap
}

// PickClass set the picked classes and only that one, removing all others in the classlist.
// If picked is empty or not in the classlist then it's not added.
func (amap AttributeMap) PickClass(classlist string, picked string) AttributeMap {
	amap.RemoveClass(classlist)
	fs := strings.Fields(picked)
	for _, f := range fs {
		if strings.Contains(classlist, f) {
			amap.AddClass(picked)
		}
	}
	return amap
}

func checkclass(class string) error {
	if class == "" || !namingpattern.IsValidName(class) {
		return fmt.Errorf("class %q is not valid", class)
	}
	return nil
}

// ResetClass removes all classes by removing the class attribute
func (amap AttributeMap) ResetClass() AttributeMap {
	delete(amap, "class")
	return amap
}

// RemoveClass removes each class in lists strings from the class attribute.
// Each lists class string can be either a single class or a string of multiple classes separated by spaces.
func (amap AttributeMap) RemoveClass(lists ...string) AttributeMap {
	existstr := amap["class"]
	new := ""
	existcs := strings.Fields(existstr)
nextexist:
	for _, existc := range existcs {
		for _, list := range lists {
			cs := strings.Fields(list)
			for _, c := range cs {
				if existc == c {
					continue nextexist
				}
			}
		}
		new += " " + existc
	}
	new = strings.TrimLeft(new, " ")
	if new == "" {
		delete(amap, "class")
	} else {
		amap["class"] = new
	}
	return amap
}

// SwitchClass removes one class and set a new one
func (amap AttributeMap) SwitchClass(removec string, addc string) AttributeMap {
	amap.RemoveClass(removec)
	amap.AddClass(addc)
	return amap
}

// TabIndex returns the TabIndex attribute.
// Returns 0 if tabindex attribute does not exists.
func (amap AttributeMap) TabIndex() (i int) {
	sidx := amap["tabindex"]
	i, _ = strconv.Atoi(sidx)
	return i
}

// SetTabIndex sets or overwrites the tabindex attribute.
// SetTabIndex returns the map to allow chainning.
func (amap AttributeMap) SetTabIndex(idx int) AttributeMap {
	amap.saveAttribute("tabindex", strconv.Itoa(idx), true)
	return amap
}

// Style returns the style attribute
func (amap AttributeMap) Style() string {
	return amap["style"]
}

// AddStyle adds the style to the style attribute.
// AddStyle returns the map to allow chainning.
func (amap AttributeMap) AddStyle(style string) AttributeMap {
	style = strings.Trim(style, " ")
	if err := checkstyle(style); err != nil {
		verbose.Error("SetStyle", err)
	} else {
		s := amap["style"]
		amap.saveAttribute("style", s+style, true)
	}
	return amap
}

// TODO: html - checkstyle
func checkstyle(style string) error {
	return nil
}

// IsTrue returns true if the attribute exists and if it's value is not false nor 0.
//
// Blanks at the ends of the name are automatically trimmed. Attribute's name is case sensitive.
func (amap AttributeMap) IsTrue(aname string) bool {
	aname = helper.Normalize(aname)
	value, found := amap[aname]
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

// SetBool sets or overwrites a boolean attribute and returns the map to allow chainning.
// SetBool does nothing if trying to set id, style, name or class attribute to true.
func (amap AttributeMap) SetBool(aname string, f bool) AttributeMap {
	if !f {
		amap.RemoveAttribute(aname)
	} else {
		amap.SetAttribute(aname, "")
	}
	return amap
}

// IsDisabled returns true if the disabled attribute is set
func (amap AttributeMap) IsDisabled() bool {
	return amap.IsTrue("disabled")
}

// SetDisabled sets the boolean disabled attribute
func (amap AttributeMap) SetDisabled(f bool) AttributeMap {
	return amap.SetBool("disabled", f)
}

// CheckAttribute returns an error if the attribute name or its value are not valid.
// Following checks are done:
//   - checks characters allowed in name (HTML5 standard).
//   - checks characters allowed in id, class and style value.
//   - checks tabindex value is numerical.
//
// See the best practice (https://stackoverflow.com/questions/925994/what-characters-are-allowed-in-an-html-attribute-name) for rules we've applied for character allowed in Name.
//
// blanks at the ends of the name are automatically trimmed. Attributes name are case-insensitve.
func CheckAttribute(name string, value string) (err error) {
	name = helper.Normalize(name)
	return checkAttribute(name, value)
}

func checkAttribute(name string, value string) (err error) {
	switch name {
	case "id", "name":
		value := strings.Trim(value, " ")
		if value != "" {
			if !namingpattern.IsValidName(value) {
				err = fmt.Errorf("%s %q is not valid", name, value)
			}
		}
	case "class":
		for _, c := range strings.Fields(value) {
			err = checkclass(c)
			if err != nil {
				break
			}
		}
	case "tabindex":
		_, err = strconv.Atoi(value)
	case "style":
		err = checkstyle(value)
	default:
	}
	if err == nil {
		if !namingpattern.IsValidName(name) {
			err = fmt.Errorf("attribute %q is not a valid name", name)
		}
	}
	return err
}

// StringifyAttributeValue makes a valid attribute value string.
//
// Returns an empty string if value is empty or if value == "true".
// Returns an unquoted string if the value is numerical.
// Returns a quoted value in any other cases. The enclosed quotes are a double quote `"` unless the value contains one otherwise quotes are a simple one `'`.
// If the value contains both kind of quotes an error is return.
func StringifyAttributeValue(value string) (string, error) {

	tv := strings.Trim(value, " ")
	lv := strings.ToLower(tv)
	if len(tv) == 0 || lv == "true" {
		return "", nil
	}

	if lv == "false" {
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

// ParseAttributes tries to parse attributes and ignore errors.
// errors are logged out if verbose mode is on.
func ParseAttributes(alist string) AttributeMap {
	amap, err := TryParseAttributes(alist)
	verbose.Error("ParseAttributes", err)
	return amap
}

// TryParseAttribute splits alist of attributes separated by spaces and sets each to the returned AttributeMap. Always returns a not nil amap.
//
// An attribute can be either a single name, usually a boolean attribute, either a name and a value at the right of an "=" symbol.
// The value can be delimited by quotes ( " or ' ) and in that case may contains whitespaces.
// The string is processed until the end or an error occurs when invalid char is met.
//
// Use ParseAttributes to chain calls and ignore errors.
func TryParseAttributes(alist string) (amap AttributeMap, err error) {
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
				_, err = amap.setAttribute(n, "", true)
				if err != nil {
					verbose.Error("ParseAttribute", err)
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
		_, err = amap.setAttribute(name, value, true)
		if err != nil {
			verbose.Error("ParseAttribute", err)
			return make(AttributeMap), err
		}
	}
	return amap, nil
}
