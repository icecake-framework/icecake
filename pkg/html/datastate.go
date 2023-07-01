package html

type DataState struct {
	//Id   string // the id of the current processing component
	//Me   any    // the current processing component, should embedd an HtmlSnippet
	Page any // the current ick page, can be nil
	App  any // the current ick App, can be nil
}
