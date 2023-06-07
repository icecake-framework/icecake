package ick0

func init() {
	TheCmpReg.RegisterComponent(Text{})
}

type Text struct {
	HtmlComponent

	Content string // html string
}

func (*Text) RegisterName() string {
	return "ick-text"
}

func (_t *Text) Body() string {
	return _t.Content
}
