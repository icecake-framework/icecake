package html

type HTMLLinker interface {

	// ParseAnchor parses _rawURL and assigns the corresponding URL to an HTML Anchor.
	ParseAnchor(_rawUrl string) error
}
