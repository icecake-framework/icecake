package ickcore

import (
	"net/url"
	"strings"
)

var (
	_cssFiles []url.URL // slice of required stylesheet link ref. will be added once into the head of the page
	_cssStyle string    // css string that will be added to the style tag in the head of the HTML file
)

// RequireCSSFile allows to declare a required CSS file.
// This can be call in the init function of a package defining custom snippets.
func RequireCSSFile(cssURL string) {
	if cssURL == "" {
		return
	}
	url, err := url.Parse(cssURL)
	if err != nil {
		return
	}

	// do not add it twice
	duplicate := false
	strurl := url.String()
	for _, cssf := range _cssFiles {
		if cssf.String() == strurl {
			duplicate = true
			break
		}
	}
	if !duplicate {
		_cssFiles = append(_cssFiles, *url)
	}
}

// RequireCSSFile allows to declare a required CSS style string for a specific snippet.
// This can be call in the init function of a package defining custom snippets.
func RequireCSSStyle(ickTagName string, cssStyle string) {
	cmt := `/* ` + ickTagName + " */\n"
	found := strings.Contains(_cssStyle, cmt)
	if !found {
		_cssStyle += cmt
		_cssStyle += cssStyle
	}
}

func RequiredCSSFile() []url.URL {
	return _cssFiles
}

func RequiredCSSStyle() string {
	return _cssStyle
}
