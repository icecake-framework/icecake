package lib

import (
	"fmt"
	"strings"

	"github.com/sunraylab/icecake/internal/helper"
)

/****************************************************************************
* Attribute
* TODO: standardize name and value, f QuoteValue()
*****************************************************************************/

// Attribute represents one of an HTML element's attribute.
type Attribute struct {
	Name string

	//TODO: handle boolean with nil
	Value string
}

// String returns normalized formated properties of this attribute
//
//	if value is empty, the format is `{name}`
//	else the format is `{name}="{value}"`
func (_attr *Attribute) String() (_ret string) {
	if _attr == nil {
		return ""
	}
	_ret = helper.Normalize(string(_attr.Name))
	if _attr.Value != "" {
		sep := "'"
		if strings.ContainsRune(_attr.Value, rune('\'')) {
			sep = "\""
		}
		_ret += `=` + sep + _attr.Value + sep
	}
	return _ret
}

/****************************************************************************
* Attributes
*****************************************************************************/

// Attributes is a slice of attribute
type Attributes []Attribute

// String returns the value of the slice serialized as a string
func (_toks *Attributes) String() (_ret string) {
	for _, v := range *_toks {
		_ret += v.String() + " "
	}
	_ret = strings.TrimRight(_ret, " ")
	return _ret
}

// Parse splits _value into attribues separated by spaces.
// Every attribute should be in the form of a single name for a boolean one,
// or {name = " value "} or {name = value }for any other one.
//
// specifications: https://www.w3.org/TR/2012/WD-html-markup-20120329/syntax.html#syntax-attributes
func ParseAttributes(_strattrs string) (_attrs Attributes, _err error) {

	var lastattr Attribute
	eol := len(_strattrs) - 1
	from := 0
	for from <= eol {

		// proceed boolean attributes until next one with a value
		valat := strings.Index(_strattrs[from:], "=")
		if valat == from {
			return nil, fmt.Errorf("unexpected starting char '=' at pos %d", valat)
		}
		namestil := from + valat
		if valat == -1 {
			namestil = eol + 1
		}
		strnames := _strattrs[from:namestil]
		names := strings.Fields(strnames)
		for _, attrName := range names {
			attrName = helper.Normalize(attrName)
			if helper.IsValidHTMLName(attrName) {
				lastattr = Attribute{Name: attrName}
				_attrs = append(_attrs, lastattr)
			} else {
				return nil, fmt.Errorf("parseAttributes failed: attribut name %q not valid", attrName)
			}
		}
		from = namestil + 1

		// proceed with the value if any
		vallen := -1
		if valat != -1 {
			strattrs := _strattrs[from:]
			strval := strings.TrimLeft(strattrs, " ")
			shift := len(strattrs) - len(strval)

			if (len(strval) > 1) && (strval[0] == '"' || strval[0] == '\'') {
				// proceed with quoted value
				quote := strval[0]
				vallen = strings.IndexByte(strval[1:], quote)
				if vallen == -1 {
					return nil, fmt.Errorf("missing ending quote in value starting at pos %d", from+1)
				}
				if len(strval) > (vallen+2) && strval[vallen+2] != ' ' {
					return nil, fmt.Errorf("quoted value must be separated with a space with the next attribute")
				}

				lastattr.Value = strval[1 : vallen+1]
				from += shift + vallen + 2

			} else if len(strval) > 0 {
				// proceed with unquoted value
				vallen = strings.IndexByte(strval, ' ')
				if vallen == -1 {
					vallen = len(strval)
				}
				lastattr.Value = strval[:vallen]
				from += shift + vallen

			} else {
				return nil, fmt.Errorf("missing value after '='")
			}
			_attrs[len(_attrs)-1] = lastattr
		}

	}

	return _attrs, nil
}
