package namingpattern

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// a range of runes
type charset struct {
	from rune
	to   rune
}

var (
	nameCharset0 []charset // list of range of runes valid for the first char of a name
	nameCharsetN []charset // list of range of runes valid for the following chars of a name
	styleCharset []charset // list of range of runes valid for a style string
)

// NameIsValid returns true or false if the s match allowed HTML Name pattern. https://stackoverflow.com/questions/925994/what-characters-are-allowed-in-an-html-attribute-name.
// usefull to check attribute or token name.
//
//	returns FALSE if s is empty
//
// must be trimed before, if required
func IsValidName(name string) bool {
	if name == "" {
		return false
	}

	if nameCharset0 == nil {
		new()
	}

	for i, c := range name {
		r := rune(c)
		if (i == 0 && !isValidRune(&nameCharset0, r)) || (i > 0 && !isValidRune(&nameCharsetN, r)) {
			return false
		}
	}
	return true
}

func IsValidNameRune(r rune, first bool) bool {
	if r == '\u0000' {
		return false
	}
	if nameCharset0 == nil {
		new()
	}
	if first {
		return isValidRune(&nameCharset0, r)
	} else {
		return isValidRune(&nameCharsetN, r)
	}
}

func IsValidStyleRune(r rune) bool {
	if r == '\u0000' {
		return false
	}
	if styleCharset == nil {
		new()
	}
	return isValidRune(&styleCharset, r)
}

func isValidRune(cs *[]charset, r rune) bool {
	for _, cset := range *cs {
		if r >= cset.from && r <= cset.to {
			return true
		}
	}
	return false
}

func new() {
	strNameCsA := `[a-z]|[A-Z]|_|:`
	nameCsA := mustCompileCharset(strNameCsA)

	strNameCsB := `[\xC0-\xD6]|[\xD8-\xF6]|[\x00F8-\x02FF]|[\x0370-\x037D]|[\x037F-\x1FFF]|[\x200C-\x200D]|[\x2070-\x218F]|[\x2C00-\x2FEF]|[\x3001-\xD7FF]|[\xF900-\xFDCF]|[\xFDF0-\xFFFD]|[\x10000-\xEFFFF]`
	nameCsB := mustCompileCharset(strNameCsB)

	strNameCsC := `-|[0-9]|.`
	nameCsC := mustCompileCharset(strNameCsC)

	strNameCsD := `\xB7|[\x0300-\x036F]|[\x203F-\x2040]`
	nameCsD := mustCompileCharset(strNameCsD)

	nameCharset0 = append(nameCsA, nameCsB...)

	nameCharsetN = append(nameCsA, nameCsC...)
	nameCharsetN = append(nameCharsetN, nameCsB...)
	nameCharsetN = append(nameCharsetN, nameCsD...)

	strStyleCsA := `[a-z]|[A-Z]|_|-|[0-9]|[\x00A0-\x00FF]`
	styleCsA := mustCompileCharset(strStyleCsA)
	styleCharset = append(styleCharset, styleCsA...)
}

func mustCompileCharset(_strcharsets string) (_ret []charset) {
	_ret, err := compileCharset(_strcharsets)
	if err != nil {
		panic(err)
	}
	return _ret
}

// compileCharset converts a chartset string into a slice of chartset.
//
// The chartstring can combine multiple sets separated by pipe char "|".
// Every set can be either a single char or a range with hexa rune code defined in braket [\xFF-\xFF]
func compileCharset(_strcharsets string) (ret []charset, err error) {
	startCharsets := strings.Split(_strcharsets, "|")
	for _, subset := range startCharsets {
		switch len(subset) {
		case 0: // empty subset
			err = fmt.Errorf("subset can not be empty")

		case 1: // single char subset
			cset := charset{}
			cset.from, err = parseRune(subset)
			if err == nil {
				cset.to = cset.from
				ret = append(ret, cset)
			}

		default: // range subset expected
			if subset[0] == '[' && subset[len(subset)-1] == ']' {
				subset = subset[1 : len(subset)-1]
				fromto := strings.Split(subset, "-")
				if len(fromto) == 2 {
					cset := charset{}
					cset.from, err = parseRune(fromto[0])
					if err == nil {
						cset.to, err = parseRune(fromto[1])
						if err == nil {
							ret = append(ret, cset)
						}
					}
				} else {
					err = fmt.Errorf("wrong range subset from-to format")
				}
			} else {
				err = fmt.Errorf("wrong range subset format, braket missing")
			}
		}
	}
	return ret, err
}

// parseRune Returns the rune code corresponding to the given string.
// This given string can be either a single letter or a hexa code formatted with the pattern `\x{FF...}`.
func parseRune(s string) (_r rune, _err error) {
	switch len([]rune(s)) {
	case 0:
		return 0, fmt.Errorf("empty string")
	case 1:
		_r, _ = utf8.DecodeRune([]byte(s))
		if _r == utf8.RuneError {
			return 0, fmt.Errorf("wrong rune hexa code in range")
		}
	default:
		switch s[:2] {
		case "\\x":
			r, err := strconv.ParseInt(s[2:], 16, 32)
			if err != nil {
				return 0, fmt.Errorf("wrong hexa in range")
			}
			_r = rune(r)
		default:
			// TODO: stringpattern - handle parseRune "\\u" case
			return 0, fmt.Errorf("wrong range definition")
		}
	}
	return _r, nil
}
