package htmlname

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type charset struct {
	from rune
	to   rune
}

var (
	charset0 []charset
	charsetN []charset
)

// IsValid returns true or false if the s match allowed HTML Name pattern. https://stackoverflow.com/questions/925994/what-characters-are-allowed-in-an-html-attribute-name.
// usefull to check attribute or token name.
//
//	returns FALSE if s is empty
//
// must be trim before, if required
func IsValid(s string) (_ret bool) {
	if s == "" {
		return false
	}

	if charset0 == nil {
		new()
	}

	for i, c := range s {
		r := rune(c)
		if (i == 0 && !isValidRune(&charset0, r)) || (i > 0 && !isValidRune(&charsetN, r)) {
			return false
		}
	}
	return true
}

func IsValidRune(r rune, first bool) bool {
	if r == '\u0000' {
		return false
	}
	if charset0 == nil {
		new()
	}
	if first {
		return isValidRune(&charset0, r)
	} else {
		return isValidRune(&charsetN, r)
	}
}

/******************************************************************************
 * PRIVATE
 */

func isValidRune(cs *[]charset, r rune) (_ret bool) {
	for _, cset := range *cs {
		if r >= cset.from && r <= cset.to {
			return true
		}
	}
	return false
}

func new() {
	strcsA := `[a-z]|[A-Z]|_|:`
	csA := mustCompileCharset(strcsA)

	strcsB := `[\xC0-\xD6]|[\xD8-\xF6]|[\x00F8-\x02FF]|[\x0370-\x037D]|[\x037F-\x1FFF]|[\x200C-\x200D]|[\x2070-\x218F]|[\x2C00-\x2FEF]|[\x3001-\xD7FF]|[\xF900-\xFDCF]|[\xFDF0-\xFFFD]|[\x10000-\xEFFFF]`
	csB := mustCompileCharset(strcsB)

	strcsC := `-|[0-9]|.`
	csC := mustCompileCharset(strcsC)

	strcsD := `\xB7|[\x0300-\x036F]|[\x203F-\x2040]`
	csD := mustCompileCharset(strcsD)

	charset0 = append(csA, csB...)

	charsetN = append(csA, csC...)
	charsetN = append(charsetN, csB...)
	charsetN = append(charsetN, csD...)
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
// The chartstring can combine multiple set with pipe char "|". Every set can be either a single char or a range with hexa rune code defined in braket [\xFF-\xFF]
func compileCharset(_strcharsets string) (_ret []charset, _err error) {
	startCharsets := strings.Split(_strcharsets, "|")
	for _, subset := range startCharsets {
		switch len(subset) {
		case 0: // empty subset
			_err = fmt.Errorf("subset can not be empty")

		case 1: // single char subset
			cset := charset{}
			cset.from, _err = parseRune(subset)
			if _err == nil {
				cset.to = cset.from
				_ret = append(_ret, cset)
			}

		default: // range subset expected
			if subset[0] == '[' && subset[len(subset)-1] == ']' {
				subset = subset[1 : len(subset)-1]
				fromto := strings.Split(subset, "-")
				if len(fromto) == 2 {
					cset := charset{}
					cset.from, _err = parseRune(fromto[0])
					if _err == nil {
						cset.to, _err = parseRune(fromto[1])
						if _err == nil {
							_ret = append(_ret, cset)
						}
					}
				} else {
					_err = fmt.Errorf("wrong range subset from-to format")
				}
			} else {
				_err = fmt.Errorf("wrong range subset format, braket missing")
			}
		}
	}
	return _ret, _err
}

// parseRune Returns the rune code corresponding to the _s string.
// _S can be either a single letter or a hexa code formatted with the pattern `\x{FF...}`.
func parseRune(_s string) (_r rune, _err error) {
	switch len([]rune(_s)) {
	case 0:
		return 0, fmt.Errorf("empty string")
	case 1:
		_r, _ = utf8.DecodeRune([]byte(_s))
		if _r == utf8.RuneError {
			return 0, fmt.Errorf("wrong rune hexa code in range")
		}
	default:
		switch _s[:2] {
		case "\\x":
			r, err := strconv.ParseInt(_s[2:], 16, 32)
			if err != nil {
				return 0, fmt.Errorf("wrong hexa in range")
			}
			_r = rune(r)
		default:
			return 0, fmt.Errorf("wrong range definition")
			// TODO: case "\\u"
		}
	}
	return _r, nil
}
