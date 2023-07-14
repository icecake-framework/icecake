package ickcore

type BareSnippet struct {
	rmeta RMetaData
	tag   Tag
}

func (bs BareSnippet) Clone() *BareSnippet {
	c := new(BareSnippet)
	c.tag = *bs.tag.Clone()
	return c
}

func (bs *BareSnippet) RMeta() *RMetaData {
	return &bs.rmeta
}

func (bs *BareSnippet) Tag() *Tag {
	if bs.tag.AttributeMap == nil {
		bs.tag.AttributeMap = make(AttributeMap)
	}
	return &bs.tag
}

func (bs *BareSnippet) SetAttribute(name string, value string) {
	bs.Tag().SetAttribute(name, value)
}

func (bs *BareSnippet) NeedRendering() bool {
	return true
}
