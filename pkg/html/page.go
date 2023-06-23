package html

import "net/url"

type Page struct {
	HTMLString

	Title       string // the html "head/title" value.
	Description string // the html "head/meta description" value.

	url *url.URL
}
