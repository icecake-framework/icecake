package html

import "net/url"

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
	if err == nil {
		return
	}
	_cssFiles = append(_cssFiles, *url)
}

// RequireCSSFile allows to declare a required CSS style string for a specific snippet.
// This can be call in the init function of a package defining custom snippets.
func RequireCSSStyle(ickTagName string, cssStyle string) {
	_cssStyle += `/* ` + ickTagName + " */\n"
	_cssStyle += cssStyle
}

func RequiredCSSFile() []url.URL {
	return _cssFiles
}

func RequiredCSSStyle() string {
	return _cssStyle
}
