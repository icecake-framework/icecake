package ick

func init() {
	Register("ick-data", Data{})
}

type Data struct {
	HtmlSnippet
	Data HTML
}

//func (Data) InlineName() string { return "ick-data" }

// Template returns a SnippetTemplate used to render the html string of a Snippet.
func (d Data) Template() SnippetTemplate {
	return SnippetTemplate{
		Body: d.Data,
	}
}
